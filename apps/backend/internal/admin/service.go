package admin

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"strings"

	"github.com/sukenda/starter/apps/backend/internal/auth"
)

// Store is the persistence boundary consumed by the admin application service.
// Keeping the interface here makes business behavior independently testable.
type Store interface {
	ListUsers(context.Context) ([]User, error)
	CreateUser(context.Context, string, string, string, string, string, []string) error
	UpdateUser(context.Context, string, string, string, string, []string) error
	ListRoles(context.Context) ([]Role, error)
	ListPermissions(context.Context) ([]Permission, error)
	CreateRole(context.Context, string, string, string, []string) error
	UpdateRole(context.Context, string, string, string, []string) error
	ListMenus(context.Context) ([]Menu, error)
	CreateMenu(context.Context, string, *string, string, string, *string, *string, *string, int, bool) error
	UpdateMenu(context.Context, string, *string, string, *string, *string, *string, int, bool) error
}

type Service struct{ repo Store }

func NewService(repo Store) *Service { return &Service{repo: repo} }

type UserInput struct {
	Email, Name, Password, Status string
	Roles                         []string
}
type RoleInput struct {
	Code, Name, Description string
	Permissions             []string
}
type MenuInput struct {
	ParentID                *string
	Code, Name              string
	Route, Icon, Permission *string
	SortOrder                int
	Active                   bool
}

func (s *Service) ListUsers(ctx context.Context) ([]User, error) { return s.repo.ListUsers(ctx) }
func (s *Service) CreateUser(ctx context.Context, in UserInput) error {
	hash, err := auth.HashPassword(in.Password)
	if err != nil { return err }
	id, err := newUUID()
	if err != nil { return err }
	status := in.Status
	if status == "" { status = "active" }
	return s.repo.CreateUser(ctx, id, strings.ToLower(strings.TrimSpace(in.Email)), strings.TrimSpace(in.Name), hash, status, in.Roles)
}
func (s *Service) UpdateUser(ctx context.Context, id string, in UserInput) error {
	status := in.Status
	if status == "" { status = "active" }
	return s.repo.UpdateUser(ctx, id, strings.ToLower(strings.TrimSpace(in.Email)), strings.TrimSpace(in.Name), status, in.Roles)
}
func (s *Service) ListRoles(ctx context.Context) ([]Role, error) { return s.repo.ListRoles(ctx) }
func (s *Service) ListPermissions(ctx context.Context) ([]Permission, error) { return s.repo.ListPermissions(ctx) }
func (s *Service) CreateRole(ctx context.Context, in RoleInput) error {
	return s.repo.CreateRole(ctx, strings.TrimSpace(in.Code), strings.TrimSpace(in.Name), strings.TrimSpace(in.Description), in.Permissions)
}
func (s *Service) UpdateRole(ctx context.Context, code string, in RoleInput) error {
	return s.repo.UpdateRole(ctx, code, strings.TrimSpace(in.Name), strings.TrimSpace(in.Description), in.Permissions)
}
func (s *Service) ListMenus(ctx context.Context) ([]Menu, error) { return s.repo.ListMenus(ctx) }
func (s *Service) CreateMenu(ctx context.Context, in MenuInput) error {
	id, err := newUUID()
	if err != nil { return err }
	return s.repo.CreateMenu(ctx, id, in.ParentID, strings.TrimSpace(in.Code), strings.TrimSpace(in.Name), in.Route, in.Icon, in.Permission, in.SortOrder, in.Active)
}
func (s *Service) UpdateMenu(ctx context.Context, id string, in MenuInput) error {
	if in.ParentID != nil && *in.ParentID == id { return ErrInvalidHierarchy }
	return s.repo.UpdateMenu(ctx, id, in.ParentID, strings.TrimSpace(in.Name), in.Route, in.Icon, in.Permission, in.SortOrder, in.Active)
}

func newUUID() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil { return "", err }
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	buf := make([]byte, 36)
	hex.Encode(buf[0:8], b[0:4]); buf[8] = '-'
	hex.Encode(buf[9:13], b[4:6]); buf[13] = '-'
	hex.Encode(buf[14:18], b[6:8]); buf[18] = '-'
	hex.Encode(buf[19:23], b[8:10]); buf[23] = '-'
	hex.Encode(buf[24:36], b[10:16])
	return string(buf), nil
}
