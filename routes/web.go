package routes

import (
	jimo "github.com/jimo-go/framework"
	"github.com/jimo-go/framework/auth"
	"github.com/jimo-go/jimo/app/http/controllers"
)

// Web registers browser (HTML) routes.
func Web(app *jimo.App) {
	user := controllers.NewUserController()
	authc := controllers.NewAuthController()

	app.Get("/", user.Home)
	app.Get("/user/{id}", user.Profile)
	app.Post("/api/test", user.APITest)

	app.Get("/login", authc.ShowLogin, jimo.Named("login"))
	app.Post("/login", authc.Login)
	app.Get("/dashboard", authc.Dashboard, jimo.WithMiddleware(auth.RequireAuth()))
	app.Get("/logout", authc.Logout)
}
