package handler

import (
	"net/http"

	"github.com/Eulenwald/GoWeb/pkg/config"
	"github.com/Eulenwald/GoWeb/pkg/models"
	"github.com/Eulenwald/GoWeb/pkg/render"
)

type Repository struct {
	App *config.AppConfig
}

// used by handlers
var PRepository *Repository

// Vreates a new repositorz
func NewRepository(pAppConfig *config.AppConfig) *Repository {
	return &Repository{
		App: pAppConfig,
	}
}

func NewHandler(pRepository *Repository) {
	PRepository = pRepository
}

// Handler for the startpage
func (pRepository *Repository) Home(w http.ResponseWriter, r *http.Request) {	
	//fmt.Fprintf(w, "Hello, 世界")
	render.RenderATemplate(w, "home-page.tmpl", &models.TmplData{})
}

// Handler for the about-page
func (pRepository *Repository) About(w http.ResponseWriter, r *http.Request) {
	sideKickMap := make(map[string]string)
	sideKickMap["morty"] = "oh my god"
	render.RenderATemplate(w, "about-page.tmpl", &models.TmplData{
		StringMap: sideKickMap,
	})
}
