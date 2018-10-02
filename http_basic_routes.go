package main

import (
	"fmt"
	"net/http"
	"log"
)

/* VOTRE MISSION :

   => Ajouter une route /recherche/ingredient/ avec son propre handler (fonction de réponse HTTP)

   => Renseigner les routes non plus manuellement mais avec une boucle for

   => Trouver un moyen d'afficher le nom de l'ingrédient dans la réponse produite
   	  ex.: => quand on accède à /recherche/ingredient/coco, la page affiche : "Cocktails contenant de la coco"

	  Piste: cf. fonction filepath.Base
*/

type Route struct {
	Name    string
	Pattern string
	Handler http.HandlerFunc
}

var (
	RoutesCoktails = []Route{
		Route{"Accès par nom",
			"/cocktail/",
			handlerCocktail},
	}
	RoutesCSP = []Route{
		Route{"Accès",
			"/csp-violation-report-endpoint/",
			handlerCSP},
	}
	RoutesForm = []Route{
		Route{"Accès",
			"/form/",
			handlerForm},
	}
	RoutesFormInformations = []Route{
		Route{"Accès",
			"/formInformations/",
			handlerFormInformations},
	}

	messages = []string{}
)

func handlerCocktail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Strict-Transport-Security","max-age=63072000")
	w.Header().Set("Content-Security-Policy","script-src 'self'; report-uri /csp-violation-report-endpoint/;")
	fmt.Fprintf(w, "<html><h2>Recette du \"Panashake\"</h2><script>alert(\"Hello\")</script></html>")
}

func handlerCSP(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Fprintf(w, "report-uri")
}

func handlerForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<html><form method='post' action='/formInformations'> Username:<input type=\"text\" name=\"username\"> <input type=\"submit\"></html>")
}

func handlerFormInformations(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
        // logic part of log in
    fmt.Println("Champ 1:", r.Form)
}

func main() {
	http.HandleFunc(RoutesCoktails[0].Pattern, RoutesCoktails[0].Handler)
	http.HandleFunc(RoutesCSP[0].Pattern, RoutesCSP[0].Handler)
	http.HandleFunc(RoutesForm[0].Pattern, RoutesForm[0].Handler)
	http.HandleFunc(RoutesFormInformations[0].Pattern, RoutesFormInformations[0].Handler)

	fmt.Println("Serveur à l'écoute sur le port 8080")
	//http.ListenAndServe(":8080", nil)
	err := http.ListenAndServeTLS(":8080", "localhost.pem", "localhost-key.pem", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
	}
}
