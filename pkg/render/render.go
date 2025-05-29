package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/juviZ138/bookings/pkg/config"
	"github.com/juviZ138/bookings/pkg/models"
)

var functions = template.FuncMap{

}

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

//
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}	

	buf := new(bytes.Buffer)
	
	td = AddDefaultData(td)


	err := t.Execute(buf,td)

	if err != nil {
		log.Println(err)
	}

	// render the template
	_,err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache , err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts,	err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache , err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache , err
		}

		if len(matches) > 0 {
			ts,err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache , err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}


// // tc = templates cache
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	// check to see if we already has a template in our cache
// 	_, inMap := tc[t]
// 	// we dont have that key in our map
// 	if !inMap {
// 		// need to create the template
// 		log.Println("create templates and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the key in the cache
// 		log.Println("using cached template")
// 	}	

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w,nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string {
// 		fmt.Sprintf("./templates/%s",t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	// parse the  templates
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}
	
// 	// add templates to map
// 	tc[t] = tmpl

// 	return nil
// }