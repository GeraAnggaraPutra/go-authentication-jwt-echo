package handlers

import (
	"go-auth-jwt/entity"
	"go-auth-jwt/helpers"
	"go-auth-jwt/services"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type biodataHandler struct {
	biodataService services.Service
}

type loginHandler struct {
	loginService services.ServiceLogin
}

func NewBiodataHandler(biodataService services.Service) *biodataHandler {
	return &biodataHandler{biodataService}
}

func NewLoginHandler(loginService services.ServiceLogin) *loginHandler {
	return &loginHandler{loginService}
}

func (h *biodataHandler) GetAll(c echo.Context) error {
	biodatas, err := h.biodataService.GetAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	var biodatasResponse []entity.Response

	for _, b := range biodatas {
		biodataResponse := convertToBiodataResponse(b)
		biodatasResponse = append(biodatasResponse, biodataResponse)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": biodatasResponse,
	})
}

func (h *loginHandler) CheckLogin(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	res, err := h.loginService.CheckLogin(username,password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	if !res {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"messages": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)
	return c.JSON(http.StatusOK, hash)
}

func convertToBiodataResponse(b *entity.Biodata) entity.Response {
	return entity.Response{
		ID:        b.ID,
		NAME:      b.NAME,
		AGE:       b.AGE,
		ADDRESS:   b.ADDRESS,
		CreatedAt: b.CreatedAt,
		UpdatedAt: b.UpdatedAt,
	}
}
