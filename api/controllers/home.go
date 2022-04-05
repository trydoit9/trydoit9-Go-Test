package controllers


import (
	"net/http"

	"github.com/trydoit9/Go-Test/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Geewonny Test API")

}