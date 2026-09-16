package admin

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound         = errors.New("resource not found")
	ErrInvalidHierarchy = errors.New("invalid menu hierarchy")
)

type Repository struct{ db *sql.DB }
func NewRepository(db *sql.DB) *Repository { return &Repository{db: db} }

type User struct { ID uint64; PublicID, Email, Name, Status string; Roles []string }
type Role struct { ID uint64; Code, Name, Description string; IsSystem bool; Permissions []string }
type Permission struct { Code, Name, Description string }
type Menu struct {
	ID uint64
	PublicID string
	ParentPublicID, Route, Icon, PermissionCode sql.NullString
	Code, Name string
	SortOrder int
	Active bool
}

func splitCodes(v string) []string {
	if v == "" { return []string{} }
	return strings.Split(v, ",")
}

func (r *Repository) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := r.db.QueryContext(ctx, `
		SELECT u.id,u.public_id,u.email,u.name,u.status,
		       COALESCE(GROUP_CONCAT(r.code ORDER BY r.code SEPARATOR ','),'')
		FROM users u
		LEFT JOIN user_roles ur ON ur.user_id=u.id
		LEFT JOIN roles r ON r.id=ur.role_id
		GROUP BY u.id,u.public_id,u.email,u.name,u.status
		ORDER BY u.id DESC`)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []User
	for rows.Next() {
		var u User; var roles string
		if err := rows.Scan(&u.ID,&u.PublicID,&u.Email,&u.Name,&u.Status,&roles); err != nil { return nil, err }
		u.Roles = splitCodes(roles); out = append(out,u)
	}
	return out, rows.Err()
}

func (r *Repository) CreateUser(ctx context.Context, publicID,email,name,passwordHash,status string, roleCodes []string) error {
	return r.withTx(ctx, func(tx *sql.Tx) error {
		res,err:=tx.ExecContext(ctx,`INSERT INTO users(public_id,email,name,password_hash,status) VALUES(?,?,?,?,?)`,publicID,email,name,passwordHash,status)
		if err!=nil{return err}; id,err:=res.LastInsertId(); if err!=nil{return err}
		return assignRoles(ctx,tx,uint64(id),roleCodes)
	})
}
func (r *Repository) UpdateUser(ctx context.Context, publicID,email,name,status string, roleCodes []string) error {
	return r.withTx(ctx,func(tx *sql.Tx)error{
		res,err:=tx.ExecContext(ctx,`UPDATE users SET email=?,name=?,status=? WHERE public_id=?`,email,name,status,publicID); if err!=nil{return err}
		n,_:=res.RowsAffected(); if n==0{return ErrNotFound}
		var id uint64; if err:=tx.QueryRowContext(ctx,`SELECT id FROM users WHERE public_id=?`,publicID).Scan(&id);err!=nil{return err}
		if _,err:=tx.ExecContext(ctx,`DELETE FROM user_roles WHERE user_id=?`,id);err!=nil{return err}
		return assignRoles(ctx,tx,id,roleCodes)
	})
}
func assignRoles(ctx context.Context,tx *sql.Tx,userID uint64,codes []string)error{
	for _,code:=range codes { res,err:=tx.ExecContext(ctx,`INSERT INTO user_roles(user_id,role_id) SELECT ?,id FROM roles WHERE code=?`,userID,code);if err!=nil{return err};n,_:=res.RowsAffected();if n==0{return fmt.Errorf("%w: role %s",ErrNotFound,code)} }
	return nil
}

func (r *Repository) ListRoles(ctx context.Context) ([]Role,error) {
	rows,err:=r.db.QueryContext(ctx,`
		SELECT r.id,r.code,r.name,COALESCE(r.description,''),r.is_system,
		       COALESCE(GROUP_CONCAT(p.code ORDER BY p.code SEPARATOR ','),'')
		FROM roles r
		LEFT JOIN role_permissions rp ON rp.role_id=r.id
		LEFT JOIN permissions p ON p.id=rp.permission_id
		GROUP BY r.id,r.code,r.name,r.description,r.is_system
		ORDER BY r.name`)
	if err!=nil{return nil,err};defer rows.Close();var out []Role
	for rows.Next(){var x Role;var permissions string;if err:=rows.Scan(&x.ID,&x.Code,&x.Name,&x.Description,&x.IsSystem,&permissions);err!=nil{return nil,err};x.Permissions=splitCodes(permissions);out=append(out,x)}
	return out,rows.Err()
}
func(r *Repository)ListPermissions(ctx context.Context)([]Permission,error){rows,err:=r.db.QueryContext(ctx,`SELECT code,name,COALESCE(description,'') FROM permissions ORDER BY code`);if err!=nil{return nil,err};defer rows.Close();var out []Permission;for rows.Next(){var x Permission;if err:=rows.Scan(&x.Code,&x.Name,&x.Description);err!=nil{return nil,err};out=append(out,x)};return out,rows.Err()}
func(r *Repository)CreateRole(ctx context.Context,code,name,description string,permissions []string)error{return r.withTx(ctx,func(tx *sql.Tx)error{res,err:=tx.ExecContext(ctx,`INSERT INTO roles(code,name,description) VALUES(?,?,?)`,code,name,description);if err!=nil{return err};id,err:=res.LastInsertId();if err!=nil{return err};return assignPermissions(ctx,tx,uint64(id),permissions)})}
func(r *Repository)UpdateRole(ctx context.Context,code,name,description string,permissions []string)error{return r.withTx(ctx,func(tx *sql.Tx)error{var id uint64;err:=tx.QueryRowContext(ctx,`SELECT id FROM roles WHERE code=?`,code).Scan(&id);if errors.Is(err,sql.ErrNoRows){return ErrNotFound};if err!=nil{return err};if _,err:=tx.ExecContext(ctx,`UPDATE roles SET name=?,description=? WHERE id=?`,name,description,id);err!=nil{return err};if _,err:=tx.ExecContext(ctx,`DELETE FROM role_permissions WHERE role_id=?`,id);err!=nil{return err};return assignPermissions(ctx,tx,id,permissions)})}
func assignPermissions(ctx context.Context,tx *sql.Tx,roleID uint64,codes []string)error{for _,code:=range codes{res,err:=tx.ExecContext(ctx,`INSERT INTO role_permissions(role_id,permission_id) SELECT ?,id FROM permissions WHERE code=?`,roleID,code);if err!=nil{return err};n,_:=res.RowsAffected();if n==0{return fmt.Errorf("%w: permission %s",ErrNotFound,code)}};return nil}

func(r *Repository)ListMenus(ctx context.Context)([]Menu,error){
	rows,err:=r.db.QueryContext(ctx,`SELECT m.id,m.public_id,parent.public_id,m.code,m.name,m.route,m.icon,p.code,m.sort_order,m.is_active FROM menus m LEFT JOIN menus parent ON parent.id=m.parent_id LEFT JOIN permissions p ON p.id=m.permission_id ORDER BY COALESCE(m.parent_id,0),m.sort_order,m.id`);if err!=nil{return nil,err};defer rows.Close();var out []Menu
	for rows.Next(){var x Menu;if err:=rows.Scan(&x.ID,&x.PublicID,&x.ParentPublicID,&x.Code,&x.Name,&x.Route,&x.Icon,&x.PermissionCode,&x.SortOrder,&x.Active);err!=nil{return nil,err};out=append(out,x)};return out,rows.Err()
}
func resolveMenuParent(ctx context.Context,q interface{QueryRowContext(context.Context,string,...any)*sql.Row}, publicID *string)(*uint64,error){if publicID==nil{return nil,nil};var id uint64;if err:=q.QueryRowContext(ctx,`SELECT id FROM menus WHERE public_id=?`,*publicID).Scan(&id);errors.Is(err,sql.ErrNoRows){return nil,ErrNotFound}else if err!=nil{return nil,err};return &id,nil}
func resolvePermission(ctx context.Context,q interface{QueryRowContext(context.Context,string,...any)*sql.Row}, code *string)(*uint64,error){if code==nil||strings.TrimSpace(*code)==""{return nil,nil};var id uint64;if err:=q.QueryRowContext(ctx,`SELECT id FROM permissions WHERE code=?`,strings.TrimSpace(*code)).Scan(&id);errors.Is(err,sql.ErrNoRows){return nil,ErrNotFound}else if err!=nil{return nil,err};return &id,nil}
func(r *Repository)CreateMenu(ctx context.Context,publicID string,parentPublicID *string,code,name string,route,icon,permission *string,sort int,active bool)error{return r.withTx(ctx,func(tx *sql.Tx)error{parentID,err:=resolveMenuParent(ctx,tx,parentPublicID);if err!=nil{return err};permissionID,err:=resolvePermission(ctx,tx,permission);if err!=nil{return err};_,err=tx.ExecContext(ctx,`INSERT INTO menus(public_id,parent_id,code,name,route,icon,permission_id,sort_order,is_active) VALUES(?,?,?,?,?,?,?,?,?)`,publicID,parentID,code,name,route,icon,permissionID,sort,active);return err})}
func(r *Repository)UpdateMenu(ctx context.Context,publicID string,parentPublicID *string,name string,route,icon,permission *string,sort int,active bool)error{return r.withTx(ctx,func(tx *sql.Tx)error{parentID,err:=resolveMenuParent(ctx,tx,parentPublicID);if err!=nil{return err};if parentID!=nil{var selfID uint64;if err:=tx.QueryRowContext(ctx,`SELECT id FROM menus WHERE public_id=?`,publicID).Scan(&selfID);errors.Is(err,sql.ErrNoRows){return ErrNotFound}else if err!=nil{return err};if *parentID==selfID{return ErrInvalidHierarchy};var cycle int;err=tx.QueryRowContext(ctx,`WITH RECURSIVE ancestors AS (SELECT id,parent_id FROM menus WHERE id=? UNION ALL SELECT m.id,m.parent_id FROM menus m JOIN ancestors a ON m.id=a.parent_id) SELECT COUNT(*) FROM ancestors WHERE id=?`,*parentID,selfID).Scan(&cycle);if err!=nil{return err};if cycle>0{return ErrInvalidHierarchy}}
		permissionID,err:=resolvePermission(ctx,tx,permission);if err!=nil{return err};res,err:=tx.ExecContext(ctx,`UPDATE menus SET parent_id=?,name=?,route=?,icon=?,permission_id=?,sort_order=?,is_active=? WHERE public_id=?`,parentID,name,route,icon,permissionID,sort,active,publicID);if err!=nil{return err};n,_:=res.RowsAffected();if n==0{return ErrNotFound};return nil})}
func(r *Repository)withTx(ctx context.Context,fn func(*sql.Tx)error)error{tx,err:=r.db.BeginTx(ctx,nil);if err!=nil{return err};if err:=fn(tx);err!=nil{_ = tx.Rollback();return err};return tx.Commit()}
