package admin

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/sukenda/starter/apps/backend/internal/auth"
	"github.com/sukenda/starter/apps/backend/internal/httpx"
)

type Handler struct{ service *Service }
func NewHandler(s *Service) *Handler { return &Handler{service:s} }
type userRequest struct { Email string `json:"email"`; Name string `json:"name"`; Password string `json:"password"`; Status string `json:"status"`; Roles []string `json:"roles"` }
type roleRequest struct { Code string `json:"code"`; Name string `json:"name"`; Description string `json:"description"`; Permissions []string `json:"permissions"` }
type menuRequest struct { ParentID *string `json:"parent_id"`; Code string `json:"code"`; Name string `json:"name"`; Route *string `json:"route"`; Icon *string `json:"icon"`; Permission *string `json:"permission"`; SortOrder int `json:"sort_order"`; Active bool `json:"active"` }
type userResponse struct { ID string `json:"id"`; Email string `json:"email"`; Name string `json:"name"`; Status string `json:"status"`; Roles []string `json:"roles"` }
type roleResponse struct { Code string `json:"code"`; Name string `json:"name"`; Description string `json:"description"`; System bool `json:"system"`; Permissions []string `json:"permissions"` }
type permissionResponse struct { Code string `json:"code"`; Name string `json:"name"`; Description string `json:"description"` }
type menuResponse struct { ID string `json:"id"`; ParentID *string `json:"parent_id"`; Code string `json:"code"`; Name string `json:"name"`; Route *string `json:"route"`; Icon *string `json:"icon"`; Permission *string `json:"permission"`; SortOrder int `json:"sort_order"`; Active bool `json:"active"` }
func nullable(v sql.NullString)*string{if !v.Valid{return nil};s:=v.String;return &s}
func usersDTO(v []User)[]userResponse{out:=make([]userResponse,0,len(v));for _,x:=range v{out=append(out,userResponse{ID:x.PublicID,Email:x.Email,Name:x.Name,Status:x.Status,Roles:x.Roles})};return out}
func rolesDTO(v []Role)[]roleResponse{out:=make([]roleResponse,0,len(v));for _,x:=range v{out=append(out,roleResponse{Code:x.Code,Name:x.Name,Description:x.Description,System:x.IsSystem,Permissions:x.Permissions})};return out}
func permissionsDTO(v []Permission)[]permissionResponse{out:=make([]permissionResponse,0,len(v));for _,x:=range v{out=append(out,permissionResponse{Code:x.Code,Name:x.Name,Description:x.Description})};return out}
func menusDTO(v []Menu)[]menuResponse{out:=make([]menuResponse,0,len(v));for _,x:=range v{out=append(out,menuResponse{ID:x.PublicID,ParentID:nullable(x.ParentPublicID),Code:x.Code,Name:x.Name,Route:nullable(x.Route),Icon:nullable(x.Icon),Permission:nullable(x.PermissionCode),SortOrder:x.SortOrder,Active:x.Active})};return out}
func fail(c fiber.Ctx,err error)error{if errors.Is(err,ErrNotFound){return httpx.Failure(c,404,"not_found","Resource was not found.",nil)};if errors.Is(err,ErrInvalidHierarchy){return httpx.Failure(c,422,"invalid_hierarchy","Menu hierarchy would contain a cycle.",nil)};return err}
func(h *Handler)ListUsers(c fiber.Ctx)error{v,e:=h.service.ListUsers(c);if e!=nil{return e};return httpx.OK(c,usersDTO(v))}
func(h *Handler)CreateUser(c fiber.Ctx)error{var r userRequest;if c.Bind().Body(&r)!=nil||strings.TrimSpace(r.Email)==""||strings.TrimSpace(r.Name)==""||len(r.Password)<12{return httpx.Failure(c,422,"validation_error","Email, name, and a password of at least 12 characters are required.",nil)};e:=h.service.CreateUser(c,UserInput{Email:r.Email,Name:r.Name,Password:r.Password,Status:r.Status,Roles:r.Roles});if e!=nil{return fail(c,e)};return c.SendStatus(204)}
func(h *Handler)UpdateUser(c fiber.Ctx)error{var r userRequest;if c.Bind().Body(&r)!=nil||r.Email==""||r.Name==""{return httpx.Failure(c,422,"validation_error","Email and name are required.",nil)};return fail(c,h.service.UpdateUser(c,c.Params("id"),UserInput{Email:r.Email,Name:r.Name,Status:r.Status,Roles:r.Roles}))}
func(h *Handler)ListRoles(c fiber.Ctx)error{v,e:=h.service.ListRoles(c);if e!=nil{return e};return httpx.OK(c,rolesDTO(v))}
func(h *Handler)Permissions(c fiber.Ctx)error{v,e:=h.service.ListPermissions(c);if e!=nil{return e};return httpx.OK(c,permissionsDTO(v))}
func(h *Handler)CreateRole(c fiber.Ctx)error{var r roleRequest;if c.Bind().Body(&r)!=nil||r.Code==""||r.Name==""{return httpx.Failure(c,422,"validation_error","Code and name are required.",nil)};return createNoContent(c,h.service.CreateRole(c,RoleInput{Code:r.Code,Name:r.Name,Description:r.Description,Permissions:r.Permissions}))}
func(h *Handler)UpdateRole(c fiber.Ctx)error{var r roleRequest;if c.Bind().Body(&r)!=nil||r.Name==""{return httpx.Failure(c,422,"validation_error","Name is required.",nil)};return fail(c,h.service.UpdateRole(c,c.Params("code"),RoleInput{Name:r.Name,Description:r.Description,Permissions:r.Permissions}))}
func(h *Handler)ListMenus(c fiber.Ctx)error{v,e:=h.service.ListMenus(c);if e!=nil{return e};return httpx.OK(c,menusDTO(v))}
func(h *Handler)CreateMenu(c fiber.Ctx)error{var r menuRequest;if c.Bind().Body(&r)!=nil||r.Code==""||r.Name==""{return httpx.Failure(c,422,"validation_error","Code and name are required.",nil)};return createNoContent(c,h.service.CreateMenu(c,MenuInput{ParentID:r.ParentID,Code:r.Code,Name:r.Name,Route:r.Route,Icon:r.Icon,Permission:r.Permission,SortOrder:r.SortOrder,Active:r.Active}))}
func(h *Handler)UpdateMenu(c fiber.Ctx)error{var r menuRequest;if c.Bind().Body(&r)!=nil||r.Name==""{return httpx.Failure(c,422,"validation_error","Name is required.",nil)};return fail(c,h.service.UpdateMenu(c,c.Params("id"),MenuInput{ParentID:r.ParentID,Name:r.Name,Route:r.Route,Icon:r.Icon,Permission:r.Permission,SortOrder:r.SortOrder,Active:r.Active}))}
func createNoContent(c fiber.Ctx,err error)error{if err!=nil{return fail(c,err)};return c.SendStatus(204)}
func RegisterRoutes(api fiber.Router,authHandler *auth.Handler,h *Handler){g:=api.Group("/admin",authHandler.Authenticate);g.Get("/users",auth.RequirePermission("users.read"),h.ListUsers);g.Post("/users",auth.RequirePermission("users.write"),h.CreateUser);g.Put("/users/:id",auth.RequirePermission("users.write"),h.UpdateUser);g.Get("/roles",auth.RequirePermission("roles.read"),h.ListRoles);g.Post("/roles",auth.RequirePermission("roles.write"),h.CreateRole);g.Put("/roles/:code",auth.RequirePermission("roles.write"),h.UpdateRole);g.Get("/permissions",auth.RequirePermission("roles.read"),h.Permissions);g.Get("/menus",auth.RequirePermission("menus.read"),h.ListMenus);g.Post("/menus",auth.RequirePermission("menus.write"),h.CreateMenu);g.Put("/menus/:id",auth.RequirePermission("menus.write"),h.UpdateMenu)}
