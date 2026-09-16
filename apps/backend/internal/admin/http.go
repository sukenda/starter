package admin

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/sukenda/starter/apps/backend/internal/auth"
	"github.com/sukenda/starter/apps/backend/internal/httpx"
)

type Handler struct{ service *Service }
func NewHandler(s *Service)*Handler{return &Handler{service:s}}

type userRequest struct{ Email string `json:"email"`; Name string `json:"name"`; Password string `json:"password"`; Status string `json:"status"`; Roles []string `json:"roles"` }
type roleRequest struct{ Code string `json:"code"`; Name string `json:"name"`; Description string `json:"description"`; Permissions []string `json:"permissions"` }
type menuRequest struct{ ParentID *uint64 `json:"parent_id"`; Code string `json:"code"`; Name string `json:"name"`; Route *string `json:"route"`; Icon *string `json:"icon"`; Permission *string `json:"permission"`; SortOrder int `json:"sort_order"`; Active bool `json:"active"` }
func fail(c fiber.Ctx,err error)error{if errors.Is(err,ErrNotFound){return httpx.Failure(c,404,"not_found","Resource was not found.",nil)};return err}
func(h *Handler)ListUsers(c fiber.Ctx)error{v,e:=h.service.ListUsers(c);if e!=nil{return e};return httpx.OK(c,v)}
func(h *Handler)CreateUser(c fiber.Ctx)error{var r userRequest;if c.Bind().Body(&r)!=nil||strings.TrimSpace(r.Email)==""||strings.TrimSpace(r.Name)==""||len(r.Password)<12{return httpx.Failure(c,422,"validation_error","Email, name, and a password of at least 12 characters are required.",nil)};e:=h.service.CreateUser(c,UserInput{Email:r.Email,Name:r.Name,Password:r.Password,Status:r.Status,Roles:r.Roles});if e!=nil{return e};return c.SendStatus(204)}
func(h *Handler)UpdateUser(c fiber.Ctx)error{var r userRequest;if c.Bind().Body(&r)!=nil||r.Email==""||r.Name==""{return httpx.Failure(c,422,"validation_error","Email and name are required.",nil)};return fail(c,h.service.UpdateUser(c,c.Params("id"),UserInput{Email:r.Email,Name:r.Name,Status:r.Status,Roles:r.Roles}))}
func(h *Handler)ListRoles(c fiber.Ctx)error{v,e:=h.service.ListRoles(c);if e!=nil{return e};return httpx.OK(c,v)}
func(h *Handler)Permissions(c fiber.Ctx)error{v,e:=h.service.ListPermissions(c);if e!=nil{return e};return httpx.OK(c,v)}
func(h *Handler)CreateRole(c fiber.Ctx)error{var r roleRequest;if c.Bind().Body(&r)!=nil||r.Code==""||r.Name==""{return httpx.Failure(c,422,"validation_error","Code and name are required.",nil)};if e:=h.service.CreateRole(c,RoleInput{Code:r.Code,Name:r.Name,Description:r.Description,Permissions:r.Permissions});e!=nil{return e};return c.SendStatus(204)}
func(h *Handler)UpdateRole(c fiber.Ctx)error{var r roleRequest;if c.Bind().Body(&r)!=nil||r.Name==""{return httpx.Failure(c,422,"validation_error","Name is required.",nil)};return fail(c,h.service.UpdateRole(c,c.Params("code"),RoleInput{Name:r.Name,Description:r.Description,Permissions:r.Permissions}))}
func(h *Handler)ListMenus(c fiber.Ctx)error{v,e:=h.service.ListMenus(c);if e!=nil{return e};return httpx.OK(c,v)}
func(h *Handler)CreateMenu(c fiber.Ctx)error{var r menuRequest;if c.Bind().Body(&r)!=nil||r.Code==""||r.Name==""{return httpx.Failure(c,422,"validation_error","Code and name are required.",nil)};if e:=h.service.CreateMenu(c,MenuInput{ParentID:r.ParentID,Code:r.Code,Name:r.Name,Route:r.Route,Icon:r.Icon,Permission:r.Permission,SortOrder:r.SortOrder,Active:r.Active});e!=nil{return e};return c.SendStatus(204)}
func(h *Handler)UpdateMenu(c fiber.Ctx)error{var r menuRequest;if c.Bind().Body(&r)!=nil||r.Name==""{return httpx.Failure(c,422,"validation_error","Name is required.",nil)};return fail(c,h.service.UpdateMenu(c,c.Params("id"),MenuInput{ParentID:r.ParentID,Name:r.Name,Route:r.Route,Icon:r.Icon,Permission:r.Permission,SortOrder:r.SortOrder,Active:r.Active}))}

func RegisterRoutes(api fiber.Router,authHandler *auth.Handler,h *Handler){g:=api.Group("/admin",authHandler.Authenticate);g.Get("/users",auth.RequirePermission("users.read"),h.ListUsers);g.Post("/users",auth.RequirePermission("users.write"),h.CreateUser);g.Put("/users/:id",auth.RequirePermission("users.write"),h.UpdateUser);g.Get("/roles",auth.RequirePermission("roles.read"),h.ListRoles);g.Post("/roles",auth.RequirePermission("roles.write"),h.CreateRole);g.Put("/roles/:code",auth.RequirePermission("roles.write"),h.UpdateRole);g.Get("/permissions",auth.RequirePermission("roles.read"),h.Permissions);g.Get("/menus",auth.RequirePermission("menus.read"),h.ListMenus);g.Post("/menus",auth.RequirePermission("menus.write"),h.CreateMenu);g.Put("/menus/:id",auth.RequirePermission("menus.write"),h.UpdateMenu)}
