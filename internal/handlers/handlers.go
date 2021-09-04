package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Aryan-mn/go_web_app/internal/config"
	"github.com/Aryan-mn/go_web_app/internal/forms"
	"github.com/Aryan-mn/go_web_app/internal/model"
	"github.com/Aryan-mn/go_web_app/internal/render"
	"log"
	"net/http"
)



// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}


//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers (r *Repository){
	Repo = r
}



func (m *Repository) Index (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r, "index.gohtml" , &model.TemplateData{})
}

func (m *Repository) About (w http.ResponseWriter, r *http.Request){

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again!!!"

	render.RenderTemplate(w,r, "about.gohtml", &model.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation (w http.ResponseWriter, r *http.Request){
	var emptyReservation model.Reservation
	data := make(map[string]interface{})
	data["reservation"] = emptyReservation
	render.RenderTemplate(w,r, "make-reservation.gohtml" , &model.TemplateData{
		Form: forms.New(nil),
		Data: data,
	})

}

func (m *Repository) PostReservation (w http.ResponseWriter, r *http.Request){
	err:= r.ParseForm()
	if err != nil{
		log.Println(err)
		return
	}
	reservation := model.Reservation{
		FirstName: r.Form.Get("first_name"),
		LastName:  r.Form.Get("last_name"),
		Email:     r.Form.Get("email"),
		Phone:     r.Form.Get("phone"),
	}
	form := forms.New(r.PostForm)

	form.Required("first_name","last_name","email")
	form.MinLength("first_name",3,r)
	form.IsEmail("email")
	if !form.Valid(){
		data := make(map[string]interface{})
		data["reservation"] = reservation
		render.RenderTemplate(w,r, "make-reservation.gohtml" , &model.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}
	m.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w,r, "/reservation-summary", http.StatusSeeOther)
}

func (m *Repository) Generals (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r, "generals.gohtml" , &model.TemplateData{})
}

func (m *Repository) Majors (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r, "majors.gohtml" , &model.TemplateData{})
}

func (m *Repository) Availability (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r, "search-availability.gohtml" , &model.TemplateData{})
}

func (m *Repository) PostAvailability (w http.ResponseWriter, r *http.Request){
	start := r.Form.Get("start")
	end := r.Form.Get("end")
	w.Write([]byte(fmt.Sprintf("The startDate is %s and the endDate is %s", start, end)))
}

type jsonResponse struct {
	Ok bool `jason:"ok"`
	Message string `json:"message"`
}
func (m *Repository) AvailabilityJson (w http.ResponseWriter, r *http.Request){

	resp := jsonResponse{
		Ok:      true,
		Message: "Hello motherFucker!",
	}

	bs, err := json.MarshalIndent(resp, "", "   ")
	if err !=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bs)
}

func (m *Repository) Contact (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w,r, "contact.gohtml" , &model.TemplateData{})
}

func (m *Repository) ReservationSummary (w http.ResponseWriter, r *http.Request){
	reservation, ok := m.App.Session.Get(r.Context(),"reservation").(model.Reservation)
	if !ok{
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "cant get reservation from session")
		http.Redirect(w,r,"/", http.StatusTemporaryRedirect)
		return
	}
	m.App.Session.Remove(r.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation
	render.RenderTemplate(w,r, "reservation-summary.gohtml" , &model.TemplateData{
		Data: data,
	})
}