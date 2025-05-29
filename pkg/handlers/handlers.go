package handlers

import (
	"net/http"

	"github.com/juviZ138/bookings/pkg/config"
	"github.com/juviZ138/bookings/pkg/models"
	"github.com/juviZ138/bookings/pkg/render"
)

// Repo the reposiotry used by handlers
var Repo *Repository

// Repository is repostiry type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repositry
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers set the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

//Home
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip",remoteIP)
	render.RenderTemplate(w,"home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello , friend"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w,"about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

