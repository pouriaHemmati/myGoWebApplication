package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"myGoWebApplication/internal/config"
	"myGoWebApplication/internal/forms"
	"myGoWebApplication/internal/models"
	render "myGoWebApplication/internal/render"
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

	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, r, "home-page.tpml", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "about-page.tpml", &models.TemplateData{})
}

// Contact is the handler for the caontact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact-page.tpml", &models.TemplateData{})
}

// Eremite is the handler for the eremite page
func (m *Repository) Eremite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "eremite-page.tpml", &models.TemplateData{})
}

// Couple is the handler for the couple page
func (m *Repository) Couple(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "couple-page.tpml", &models.TemplateData{})
}

// Family is the handler for the family page
func (m *Repository) Family(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "family-page.tpml", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "check-availability-page.tpml", &models.TemplateData{})
}

// PostReservation is the handler for the reservation page and POST requests
func (m *Repository) PostReservation(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("startingDate")
	end := r.Form.Get("endingDate")
	w.Write([]byte(fmt.Sprintf("Arrival date value is set to %s, departure date value to %s", start, end)))
}

type jsonResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) ReservationJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok:      true,
		Message: "It is available!",
	}
	output, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

// MakeReservation is the handler for the make-reservatio page
func (m *Repository) MakeReservation(w http.ResponseWriter, r *http.Request) {
	var emptyReservation models.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(w, r, "make-reservation-page.tpml", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})
}
func (m *Repository) PostMakeReservation(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}
	reservation := models.Reservation{
		Name:  r.Form.Get("full_name"),
		Email: r.Form.Get("email"),
		Phone: r.Form.Get("phone"),
	}

	form := forms.New(r.PostForm)

	//form.Has("full_name", r)
	form.Required("full_name", "email")
	form.MinLength("full_name", 2, r)
	form.IsEmail("email")
	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(w, r, "make-reservation-page.tpml", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-overview", http.StatusSeeOther)
}

func (m *Repository) ReservationOverview(w http.ResponseWriter, r *http.Request) {
	reservation, ok := m.App.Session.Get(r.Context(), "reservation").(models.Reservation)
	if !ok {
		log.Println("reservation not found")
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(w, r, "reservation-overview-page.tpml", &models.TemplateData{
		Data: data,
	})
}
