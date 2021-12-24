package handlers

import (
	"net/http"

	"github.com/Holilamd/module-pemeliharaan/app/models"
	"github.com/Holilamd/module-pemeliharaan/app/services"
	"github.com/Holilamd/module-pemeliharaan/helper"
	"github.com/gin-gonic/gin"
)

type handlerUser struct {
	userservice services.Serviceuser
}

func NewHandlerUser(userservice services.Serviceuser) *handlerUser {
	return &handlerUser{userservice}
}
func (h *handlerUser) All(g *gin.Context) {
	result, err := h.userservice.All()
	if err != nil {
		respond := helper.Error("Data User", http.StatusOK, err)
		g.JSON(http.StatusOK, respond)
		return
	}
	respond := helper.Success("Data User", http.StatusOK, result)
	g.JSON(http.StatusOK, respond)
}

func (h *handlerUser) Create(g *gin.Context) {
	var input models.Userinput
	err := g.ShouldBindJSON(&input)
	if err != nil {
		if err != nil {
			errors := helper.FormatValidationError(err)
			errorMessage := errors
			response := helper.Error("Create Data User", http.StatusOK, errorMessage)
			g.JSON(http.StatusOK, response)
			return
		}
	}
	result, err := h.userservice.Create(input)
	if err != nil {
		response := helper.Error(err.Error(), http.StatusOK, nil)
		g.JSON(http.StatusOK, response)
		return
	}
	response := helper.Success("Create Data User", http.StatusOK, result)
	g.JSON(http.StatusOK, response)
}
