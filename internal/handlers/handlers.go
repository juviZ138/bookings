package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/juviZ138/bookings/internal/config"
	"github.com/juviZ138/bookings/internal/models"
	"github.com/juviZ138/bookings/internal/render"
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
	render.RenderTemplate(w,r,"home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "hello , friend"

	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")

	stringMap["remote_ip"] = remoteIP

	// send the data to the template
	render.RenderTemplate(w,r,"about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Reservation render  the make a reservation page and display form
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"make-reservation.page.tmpl",  &models.TemplateData{})
}

// Generals render  the room Generals page
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"generals.page.tmpl", &models.TemplateData{})
}

// Majors render  the room Majors page
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"majors.page.tmpl", &models.TemplateData{})
}

// Availiability render  the room Availiability page
func (m *Repository) Availiability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"search-availability.page.tmpl", &models.TemplateData{})
}


// Post Availiability render  the room Availiability page
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("start date is: %s and end date is: %s",start,end)))
}


type jsonResponse struct {
	OK bool `json:"ok"`
	Message string `json:"message"`
}
//  AvailabilityJSON 
func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK: false,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type","application/json")

	w.Write(out)
}


// Contact render  the room Contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r,"contact.page.tmpl", &models.TemplateData{})
}