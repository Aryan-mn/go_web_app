package handlers

import (
	"github.com/Aryan-mn/go_web_app/pkg/config"
	"github.com/Aryan-mn/go_web_app/pkg/model"
	"github.com/Aryan-mn/go_web_app/pkg/render"
	"net/http"
)



// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct{
	App *config.AppConfig
}


//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App: a,
	}
}
// NewHandlers sets the repository for the handlers
func NewHandlers (r *Repository){
	Repo = r
}



func (m *Repository) Index (w http.ResponseWriter, r *http.Request){
	render.RenderTemplate(w, "index.gohtml" , &model.TemplateData{})
}

func (m *Repository) About (w http.ResponseWriter, r *http.Request){

	stringMap := make(map[string]string)
	stringMap["test"] = "hello, again!!!"

	render.RenderTemplate(w, "about.gohtml", &model.TemplateData{
		StringMap: stringMap,
	})
}

