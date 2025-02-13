package handlers

import (
	"myGoWebApplication/pkg/config"
	"myGoWebApplication/pkg/models"
	render "myGoWebApplication/pkg/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	render.RenderTemplate(w, "home-page.tpml", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// some data or calculation takes place

	sidekickMap := make(map[string]string)
	sidekickMap["morty"] = "Ooh, wee!"

	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	sidekickMap["remote_ip"] = remoteIp

	// send the result or any prepared data to the template
	render.RenderTemplate(w, "about-page.tpml", &models.TemplateData{
		StringMap: sidekickMap,
	})
}
