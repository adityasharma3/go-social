package handlers

import (
	"net/http"

	"github.com/adityasharma3/go-social/internal/store"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

var users = []store.User{}

func (h *UserHandler) GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}
