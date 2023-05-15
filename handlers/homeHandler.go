package handlers

import (
	"net/http"

	"github.com/Louiservelle/findme/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	models.RenderTemplate(w, "index")
}
