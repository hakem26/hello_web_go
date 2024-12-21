package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//WITHOUT CACHE TEMPLATE
// func RenderTemplate(w http.ResponseWriter, tmpl string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
// 	err := parsedTemplate.Execute(w, nil)
// 	if err != nil {
// 		fmt.Println("An Error about:", err)
// 		return
// 	}
// }

// SIMPLE CACHE TEMPLATE
// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	//check if already have template in cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		//need to create the template
// 		log.Println("Creating the template and adding to cache")
// 		err = createTemplateChache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		//we have template in cache
// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateChache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	//parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	//add template to cache
// 	tc[t] = tmpl
// 	return nil
// }

//COMPLEX CACHE TEMPLATE
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	//create the template cache

	//get requested template from cache

	//render the template
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("An Error about:", err)
		return
	}
}

func createTemplateChache() (map[string]*template.Template) {
	
}