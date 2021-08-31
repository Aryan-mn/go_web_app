package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Aryan-mn/go_web_app/internal/config"
	"github.com/Aryan-mn/go_web_app/internal/model"
	"github.com/Aryan-mn/go_web_app/internal/render"
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
	render.RenderTemplate(w,r, "make-reservation.gohtml" , &model.TemplateData{})
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