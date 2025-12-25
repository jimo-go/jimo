package controllers

import (
	"net/http"
	"strconv"

	jimo "github.com/jimo-go/framework"
	"github.com/jimo-go/jimo/app/models"
)

type UserController struct{}

func NewUserController() *UserController { return &UserController{} }

func (c *UserController) Home(ctx *jimo.Context) {
	ctx.View("index", map[string]any{
		"Message": "Welcome to JIMO",
	})
}

func (c *UserController) Profile(ctx *jimo.Context) {
	id := ctx.Param("id")
	n, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]any{"message": "Invalid user id"})
		return
	}

	user, ok, err := models.Users().Find(n)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]any{"message": "Failed to load user"})
		return
	}
	if !ok {
		ctx.JSON(http.StatusNotFound, map[string]any{"message": "User not found"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type apiTestRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (c *UserController) APITest(ctx *jimo.Context) {
	var body apiTestRequest
	ctx.MustBind(&body)
	ctx.JSON(200, body)
}
