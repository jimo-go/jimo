package controllers

import (
	"net/http"

	jimo "github.com/jimo-go/framework"
	"github.com/jimo-go/framework/auth"
	"github.com/jimo-go/framework/validation"
	"github.com/jimo-go/jimo/app/models"
)

type AuthController struct{}

func NewAuthController() *AuthController { return &AuthController{} }

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (c *AuthController) ShowLogin(ctx *jimo.Context) {
	flash := ""
	if s := ctx.Session(); s != nil {
		if v := s.PullFlash("error"); v != nil {
			if str, ok := v.(string); ok {
				flash = str
			}
		}
	}

	ctx.View("login", map[string]any{
		"CSRF":  ctx.CSRFToken(),
		"Error": flash,
	})
}

func (c *AuthController) Login(ctx *jimo.Context) {
	req := loginRequest{
		Email:    ctx.Request.FormValue("email"),
		Password: ctx.Request.FormValue("password"),
	}

	ctx.MustValidate(req, validation.Rules{
		"email":    "required|email",
		"password": "required|min:6",
	})

	user, ok, err := models.FindUserByEmail(req.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "Auth failed"})
		return
	}
	if !ok || !auth.CheckPassword(req.Password, user.PasswordHash) {
		if s := ctx.Session(); s != nil {
			s.Flash("error", "Invalid credentials")
		}
		ctx.ResponseWriter.Header().Set("Location", "/login")
		ctx.ResponseWriter.WriteHeader(http.StatusFound)
		return
	}

	auth.Login(ctx, user.ID)
	ctx.ResponseWriter.Header().Set("Location", "/dashboard")
	ctx.ResponseWriter.WriteHeader(http.StatusFound)
}

func (c *AuthController) Dashboard(ctx *jimo.Context) {
	id, _ := auth.UserID(ctx)
	ctx.View("dashboard", map[string]any{
		"UserID": id,
	})
}

func (c *AuthController) Logout(ctx *jimo.Context) {
	auth.Logout(ctx)
	ctx.ResponseWriter.Header().Set("Location", "/")
	ctx.ResponseWriter.WriteHeader(http.StatusFound)
}
