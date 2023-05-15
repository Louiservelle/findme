package handlers

import (
	"net/http"

	"github.com/Louiservelle/findme/models"
)

func About(w http.ResponseWriter, r *http.Request) {
	models.RenderTemplate(w, "about")
}
