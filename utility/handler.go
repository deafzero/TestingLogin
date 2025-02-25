package utility

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

// index handler
func Index(ctx *fiber.Ctx) error {
	return ctx.Redirect("/login")
}

// login hanlders
func LoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("login", flash.Get(ctx))
}

func LoginAuth(ctx *fiber.Ctx) error {
	// System Requiremment
	type Requirements struct {
		Username string `form:"username" validate:"required"`
		Password string `form:"password" validate:"required,min=8"`
	}
	// extract request to body
	var body Requirements
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}

	// pesan ketika validasi error
	errMessage := map[string]string{
		"Username.required": "Username harus diisi!",
		"Password.required": "Passowrd harus diisi!",
		"Password.min":      "Passowrd minimal 8 karakter!",
	}

	validate := validator.New()
	if err := validate.Struct(body); err != nil {
		errors := make(map[string]any)
		for _, verr := range err.(validator.ValidationErrors) {
			errors[verr.Field()] = errMessage[verr.Field()+"."+verr.Tag()]
		}

		return flash.WithData(ctx, errors).Redirect("/login")
	}

	// validasi Users
	if body.Username == "lucas" && body.Password == "ganteng123" {
		return ctx.Redirect("/dashboard")
	}
	return flash.WithData(ctx, fiber.Map{"notvalid": "username atau password salah"}).Redirect("/login")
}

func Dashboard(ctx *fiber.Ctx) error {
	return ctx.Render("dashboard", fiber.Map{})
}
