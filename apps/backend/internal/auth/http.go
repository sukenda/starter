package auth

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/sukenda/starter/apps/backend/internal/httpx"
)

const principalLocal = "auth.principal"

type Handler struct{ service *Service }
func NewHandler(service *Service)*Handler{return &Handler{service:service}}

type loginRequest struct { Email string `json:"email"`; Password string `json:"password"` }
type refreshRequest struct { RefreshToken string `json:"refresh_token"` }
type tokenResponse struct { AccessToken string `json:"access_token"`; RefreshToken string `json:"refresh_token"`; TokenType string `json:"token_type"`; AccessExpiresAt string `json:"access_expires_at"`; RefreshExpiresAt string `json:"refresh_expires_at"` }
type userResponse struct { ID string `json:"id"`; Email string `json:"email"`; Name string `json:"name"`; Permissions []string `json:"permissions,omitempty"` }

func tokenDTO(t Tokens)tokenResponse{return tokenResponse{AccessToken:t.AccessToken,RefreshToken:t.RefreshToken,TokenType:"Bearer",AccessExpiresAt:t.AccessExpiresAt.Format("2006-01-02T15:04:05.999999999Z07:00"),RefreshExpiresAt:t.RefreshExpiresAt.Format("2006-01-02T15:04:05.999999999Z07:00")}}

func (h *Handler) Login(c fiber.Ctx)error{
	var req loginRequest;if err:=c.Bind().Body(&req);err!=nil{return httpx.Failure(c,400,"invalid_request","Invalid request body.",nil)};if strings.TrimSpace(req.Email)==""||req.Password==""{return httpx.Failure(c,422,"validation_error","Email and password are required.",nil)}
	user,tokens,err:=h.service.Login(c,req.Email,req.Password);if err!=nil{if errors.Is(err,ErrInvalidCredentials){return httpx.Failure(c,401,"invalid_credentials","Invalid email or password.",nil)};return err}
	return httpx.OK(c,fiber.Map{"user":userResponse{ID:user.PublicID,Email:user.Email,Name:user.Name},"tokens":tokenDTO(tokens)})
}
func(h *Handler)Refresh(c fiber.Ctx)error{var req refreshRequest;if err:=c.Bind().Body(&req);err!=nil||req.RefreshToken==""{return httpx.Failure(c,400,"invalid_request","Refresh token is required.",nil)};tokens,err:=h.service.Refresh(c,req.RefreshToken);if err!=nil{if errors.Is(err,ErrUnauthorized){return httpx.Failure(c,401,"invalid_refresh_token","Refresh token is invalid or expired.",nil)};return err};return httpx.OK(c,tokenDTO(tokens))}
func(h *Handler)Me(c fiber.Ctx)error{p,ok:=PrincipalFromContext(c);if !ok{return httpx.Failure(c,401,"unauthorized","Authentication is required.",nil)};codes:=make([]string,0,len(p.Permissions));for code:=range p.Permissions{codes=append(codes,code)};return httpx.OK(c,userResponse{ID:p.PublicID,Email:p.Email,Name:p.Name,Permissions:codes})}
func(h *Handler)Logout(c fiber.Ctx)error{p,ok:=PrincipalFromContext(c);if !ok{return httpx.Failure(c,401,"unauthorized","Authentication is required.",nil)};if err:=h.service.Logout(c,p);err!=nil{return err};return httpx.OK(c,fiber.Map{"logged_out":true})}

func(h *Handler)Authenticate(c fiber.Ctx)error{header:=strings.TrimSpace(c.Get("Authorization"));parts:=strings.Fields(header);if len(parts)!=2||!strings.EqualFold(parts[0],"Bearer"){return httpx.Failure(c,401,"unauthorized","A valid bearer token is required.",nil)};p,err:=h.service.Authenticate(c,parts[1]);if err!=nil{if errors.Is(err,ErrUnauthorized){return httpx.Failure(c,401,"unauthorized","Authentication is required.",nil)};return err};c.Locals(principalLocal,p);return c.Next()}
func RequirePermission(code string)fiber.Handler{return func(c fiber.Ctx)error{p,ok:=PrincipalFromContext(c);if !ok{return httpx.Failure(c,401,"unauthorized","Authentication is required.",nil)};if _,ok:=p.Permissions[code];!ok{return httpx.Failure(c,403,"forbidden","You do not have permission to perform this action.",nil)};return c.Next()}}
func PrincipalFromContext(c fiber.Ctx)(Principal,bool){value:=c.Locals(principalLocal);p,ok:=value.(Principal);return p,ok}

func RegisterRoutes(api fiber.Router,h *Handler){authGroup:=api.Group("/auth");authGroup.Post("/login",h.Login);authGroup.Post("/refresh",h.Refresh);protected:=authGroup.Group("",h.Authenticate);protected.Get("/me",h.Me);protected.Post("/logout",h.Logout)}
