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
var Repo *Repository

// Vreates a new repositorz
func NewRepo(pAppConfig *config.AppConfig) *Repository {
	return &Repository{
		App: pAppConfig,
	}
}

func NewHandler(pRepository *Repository) {
	Repo = pRepository
}

// Handler for the startpage
func (pR *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "Hello, 世界")
	//fmt.Fprintf(w, "Hello, 世界")
	render.RenderATemplate(w, "home-page.tmpl", &models.TmplData{})
}

// Handler for the about-page
func (pR *Repository) About(w http.ResponseWriter, r *http.Request) {
	sideKickMap := make(map[string]string)
	sideKickMap["morty"] = "oh my god"
	render.RenderATemplate(w, "about-page.tmpl", &models.TmplData{
		StringMap: sideKickMap,
	})
}
