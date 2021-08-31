package render

import (
	"bytes"
	"github.com/Aryan-mn/go_web_app/internal/config"
	"github.com/Aryan-mn/go_web_app/internal/model"
	"github.com/justinas/nosurf"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

var app *config.AppConfig
func NewTemplate(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *model.TemplateData, r *http.Request) *model.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}


func RenderTemplate(w http.ResponseWriter,r *http.Request, tpl string, td *model.TemplateData){
	var tc map[string]*template.Template

	if app.UseCache{
		tc = app.TemplateCache
	}else{
		tc,_ = CreatTemplateCache()
	}


	t, ok := tc[tpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)
	_ = t.Execute(buf, td)

	_,err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

// Creat a template cache as a map
func CreatTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	pages,err := filepath.Glob("./templates/*.gohtml")
	if err !=nil{
		return myCache, err
	}

	for _,page := range pages {
		name :=filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err !=nil{
			return myCache,err
		}

		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err !=nil{
			return myCache,err
		}
		if len(matches) >0 {
			ts,err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err !=nil{
				return myCache,err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
