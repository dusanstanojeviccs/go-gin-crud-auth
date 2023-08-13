package users

import (
	"go-gin-crud-auth/security"
	"go-gin-crud-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func signup(ctx *gin.Context) {
	signUpRequest := new(SignUpRequest)

	ctx.BindJSON(signUpRequest)

	saltedPassword, error := bcrypt.GenerateFromPassword([]byte(signUpRequest.Password), 13)

	if error != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := &User{
		Name:     signUpRequest.Name,
		Email:    signUpRequest.Email,
		Password: string(saltedPassword[:]),
	}

	validations := &[]*utils.ValidationMessage{}

	error = validate(validations, user)

	if error != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if len(*validations) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, validations)
		return
	}

	UserRepository.create(user)

	utils.Session.SetUserId(ctx, user.Id)
	utils.Cookies.SetSessionCookie(
		ctx,
		utils.Jwt.GenerateSessionCookie(user.Id),
	)

	current(ctx)
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func login(ctx *gin.Context) {
	// if the user exists logs them in
	// sends the set-cookie header back

	loginRequest := new(LoginRequest)

	ctx.BindJSON(loginRequest)

	foundUser, error := UserRepository.findByEmail(loginRequest.Email)
	if error != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if foundUser != nil {

		error := bcrypt.CompareHashAndPassword(
			[]byte(foundUser.Password),
			[]byte(loginRequest.Password),
		)

		if error == nil {
			// we need to login and write the cookie header
			utils.Session.SetUserId(ctx, foundUser.Id)
			utils.Cookies.SetSessionCookie(
				ctx,
				utils.Jwt.GenerateSessionCookie(foundUser.Id),
			)

			current(ctx)
		}
	}

	ctx.AbortWithStatus(http.StatusBadRequest)
}

type CurrentUser struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func current(ctx *gin.Context) {
	userId := utils.Session.GetUserId(ctx)

	user, error := UserRepository.findById(userId)

	if error != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if user == nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	ctx.JSON(200, &CurrentUser{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
	})
}

func RegisterEndpoints(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users/current", security.LoggedInFilter, current)
}
