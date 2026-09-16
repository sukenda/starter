package auth

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"strings"
	"time"
)

var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUnauthorized = errors.New("unauthorized")

type Principal struct {
	UserID      uint64
	PublicID    string
	Email       string
	Name        string
	Permissions map[string]struct{}
	SessionID   uint64
}

type Tokens struct {
	AccessToken      string
	RefreshToken     string
	AccessExpiresAt  time.Time
	RefreshExpiresAt time.Time
}

type Service struct {
	repo       *Repository
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewService(repo *Repository, accessTTL, refreshTTL time.Duration) *Service {
	return &Service{repo: repo, accessTTL: accessTTL, refreshTTL: refreshTTL}
}

func (s *Service) Login(ctx context.Context, email, password string) (User, Tokens, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	user, err := s.repo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) { return User{}, Tokens{}, ErrInvalidCredentials }
		return User{}, Tokens{}, err
	}
	if user.Status != "active" { return User{}, Tokens{}, ErrInvalidCredentials }
	valid, err := VerifyPassword(user.PasswordHash, password)
	if err != nil || !valid { return User{}, Tokens{}, ErrInvalidCredentials }

	access, err := NewOpaqueToken(); if err != nil { return User{}, Tokens{}, err }
	refresh, err := NewOpaqueToken(); if err != nil { return User{}, Tokens{}, err }
	publicID, err := newUUID(); if err != nil { return User{}, Tokens{}, err }
	now := time.Now().UTC()
	tokens := Tokens{AccessToken: access, RefreshToken: refresh, AccessExpiresAt: now.Add(s.accessTTL), RefreshExpiresAt: now.Add(s.refreshTTL)}
	if err := s.repo.CreateSession(ctx, publicID, user.ID, HashToken(access), HashToken(refresh), tokens.AccessExpiresAt, tokens.RefreshExpiresAt); err != nil { return User{}, Tokens{}, err }
	return user, tokens, nil
}

func (s *Service) Authenticate(ctx context.Context, token string) (Principal, error) {
	if token == "" { return Principal{}, ErrUnauthorized }
	session, user, err := s.repo.FindSessionByAccessToken(ctx, HashToken(token))
	if err != nil { if errors.Is(err, ErrSessionNotFound) { return Principal{}, ErrUnauthorized }; return Principal{}, err }
	if session.RevokedAt.Valid || time.Now().UTC().After(session.AccessExpiresAt) || user.Status != "active" { return Principal{}, ErrUnauthorized }
	codes, err := s.repo.PermissionsForUser(ctx, user.ID); if err != nil { return Principal{}, err }
	permissions := make(map[string]struct{}, len(codes)); for _, code := range codes { permissions[code] = struct{}{} }
	return Principal{UserID: user.ID, PublicID: user.PublicID, Email: user.Email, Name: user.Name, Permissions: permissions, SessionID: session.ID}, nil
}

func (s *Service) Logout(ctx context.Context, principal Principal) error { return s.repo.RevokeSession(ctx, principal.SessionID) }

func newUUID() (string, error) {
	b := make([]byte, 16); if _, err := rand.Read(b); err != nil { return "", err }
	b[6] = (b[6] & 0x0f) | 0x40; b[8] = (b[8] & 0x3f) | 0x80
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], b[0:4]); buf[8] = '-'; hex.Encode(buf[9:13], b[4:6]); buf[13] = '-'; hex.Encode(buf[14:18], b[6:8]); buf[18] = '-'; hex.Encode(buf[19:23], b[8:10]); buf[23] = '-'; hex.Encode(buf[24:36], b[10:16])
	return string(buf), nil
}
