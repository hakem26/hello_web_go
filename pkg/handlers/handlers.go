package handlers

import (
	"net/http"

	"github.com/hakem26/hello_web_go/pkg/config"
	"github.com/hakem26/hello_web_go/pkg/models"
	"github.com/hakem26/hello_web_go/pkg/render"
)

// the repository used by handlers
var Repo *Repository

// the repository type
type Repository struct {
	App *config.AppConfig
}

// create a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
    stringMap := make(map[string]string)
    stringMap["test"] = "Hello, Test."

    render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
        StringMap: stringMap,
    })
}

