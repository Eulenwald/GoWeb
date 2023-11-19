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
var pRepository *Repository

func GetRepository() *Repository {
	return pRepository
}

// param: pointer to a app-modul
// return: pointer to 1 Repository
// creates a new repository include the pointer of the AppConfig
func NewRepository(pAppConfig *config.AppConfig) *Repository {
	return &Repository{
		App: pAppConfig,
	}
}

// param: pointer to an Repository
// return: none
// linked the pointer to the package gloabl var PRepositorz
func NewHandler(pR *Repository) {
	pRepository = pR
}

// handler for the startpage
func (pR *Repository) Home(w http.ResponseWriter, r *http.Request) {	
	//fmt.Fprintf(w, "Hello, 世界")
	render.RenderATemplate(w, "home-page.tmpl", &models.TmplData{})
}

// handler for the about-page
func (pR *Repository) About(w http.ResponseWriter, r *http.Request) {
	sideKickMap := make(map[string]string)
	sideKickMap["morty"] = "oh my god"
	render.RenderATemplate(w, "about-page.tmpl", &models.TmplData{
		StringMap: sideKickMap,
	})
}
