package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var ErrUserNotFound = errors.New("user not found")
var ErrSessionNotFound = errors.New("session not found")
var ErrRefreshTokenReused = errors.New("refresh token already rotated")

type User struct {
	ID uint64
	PublicID, Email, Name, PasswordHash, Status string
}

type Session struct {
	ID uint64
	PublicID string
	UserID uint64
	AccessExpiresAt, RefreshExpiresAt time.Time
	RevokedAt sql.NullTime
}

type Repository struct{ db *sql.DB }
func NewRepository(db *sql.DB) *Repository { return &Repository{db: db} }

func (r *Repository) FindUserByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := r.db.QueryRowContext(ctx, `SELECT id, public_id, email, name, password_hash, status FROM users WHERE email=? LIMIT 1`, email).Scan(&u.ID,&u.PublicID,&u.Email,&u.Name,&u.PasswordHash,&u.Status)
	if errors.Is(err, sql.ErrNoRows) { return User{}, ErrUserNotFound }; return u, err
}

func (r *Repository) CreateSession(ctx context.Context, publicID string, userID uint64, accessHash, refreshHash [32]byte, accessExpiry, refreshExpiry time.Time) error {
	_, err := r.db.ExecContext(ctx, `INSERT INTO auth_sessions (public_id,user_id,access_token_hash,refresh_token_hash,access_expires_at,refresh_expires_at) VALUES (?,?,?,?,?,?)`, publicID,userID,accessHash[:],refreshHash[:],accessExpiry,refreshExpiry); return err
}

func scanSession(row *sql.Row) (Session, User, error) {
	var s Session; var u User
	err := row.Scan(&s.ID,&s.PublicID,&s.UserID,&s.AccessExpiresAt,&s.RefreshExpiresAt,&s.RevokedAt,&u.ID,&u.PublicID,&u.Email,&u.Name,&u.PasswordHash,&u.Status)
	if errors.Is(err, sql.ErrNoRows) { return Session{}, User{}, ErrSessionNotFound }; return s,u,err
}

func (r *Repository) FindSessionByAccessToken(ctx context.Context, hash [32]byte) (Session, User, error) {
	return scanSession(r.db.QueryRowContext(ctx, `SELECT s.id,s.public_id,s.user_id,s.access_expires_at,s.refresh_expires_at,s.revoked_at,u.id,u.public_id,u.email,u.name,u.password_hash,u.status FROM auth_sessions s JOIN users u ON u.id=s.user_id WHERE s.access_token_hash=? LIMIT 1`, hash[:]))
}

func (r *Repository) FindSessionByRefreshToken(ctx context.Context, hash [32]byte) (Session, User, error) {
	return scanSession(r.db.QueryRowContext(ctx, `SELECT s.id,s.public_id,s.user_id,s.access_expires_at,s.refresh_expires_at,s.revoked_at,u.id,u.public_id,u.email,u.name,u.password_hash,u.status FROM auth_sessions s JOIN users u ON u.id=s.user_id WHERE s.refresh_token_hash=? LIMIT 1`, hash[:]))
}

func (r *Repository) RotateSession(ctx context.Context, id uint64, previousRefreshHash, accessHash, refreshHash [32]byte, accessExpiry, refreshExpiry time.Time) error {
	result, err := r.db.ExecContext(ctx, `UPDATE auth_sessions SET access_token_hash=?,refresh_token_hash=?,access_expires_at=?,refresh_expires_at=? WHERE id=? AND refresh_token_hash=? AND revoked_at IS NULL`, accessHash[:],refreshHash[:],accessExpiry,refreshExpiry,id,previousRefreshHash[:])
	if err != nil { return err }; n, err := result.RowsAffected(); if err != nil { return err }; if n != 1 { return ErrRefreshTokenReused }; return nil
}

func (r *Repository) PermissionsForUser(ctx context.Context, userID uint64) ([]string, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT DISTINCT p.code FROM permissions p JOIN role_permissions rp ON rp.permission_id=p.id JOIN user_roles ur ON ur.role_id=rp.role_id WHERE ur.user_id=? ORDER BY p.code`, userID); if err != nil { return nil,err }; defer rows.Close()
	var out []string; for rows.Next(){ var code string; if err:=rows.Scan(&code);err!=nil{return nil,err}; out=append(out,code) }; return out,rows.Err()
}

func (r *Repository) RevokeSession(ctx context.Context, id uint64) error { _,err:=r.db.ExecContext(ctx,`UPDATE auth_sessions SET revoked_at=CURRENT_TIMESTAMP(6) WHERE id=? AND revoked_at IS NULL`,id); return err }
