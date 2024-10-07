package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
)

type Etudiant struct {
	Nom    string
	Prenom string
	Age    int
	Sexe   bool
}

type Promo struct {
	Name         string
	Filière      string
	Niveau       string
	NbrEtudiant  int
	ListEtudiant []Etudiant
}

type Change struct {
	Pair     bool
	Compteur int
}

type StockageForm struct {
	CheckValue bool
	Value      string
}

type PageAffiche struct {
	CheckValue bool
	Value      string
	IsEmpty    bool
}

var stockageForm = StockageForm{false, ""}

func main() {
	temp, err := template.ParseGlob("./templates/*.html")
	Compteur := Change{true, 0}
	if err != nil {
		fmt.Printf("Erreur => %s\n", err.Error())
		os.Exit(02)
	}
	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		LE := []Etudiant{{"Lecomte", "Adrien", 20, true}, {"Petitfrere", "Alexandre", 20, true}, {"Rodrigues", "Cyril", 24, false}}
		data := Promo{" B1 Informatique", " Informatique", "Bachelor 1", len(LE), LE}
		temp.ExecuteTemplate(w, "promo", data)
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) {
		Compteur.Compteur += 1
		Compteur.Pair = Compteur.Compteur%2 == 0
		temp.ExecuteTemplate(w, "change", Compteur)
	})

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "AffichageFormulaire", nil)
	})

	http.HandleFunc("/user/traitement", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorecte", http.StatusMovedPermanently)
			return
		}

		checkValue, _ := regexp.MatchString("^[\\p{L}-]{1,64}$", r.FormValue("name"))
		if !checkValue {
			stockageForm = StockageForm{false, ""}
			http.Redirect(w, r, "/erreur?code=400&message=Oups les données sont invalides", http.StatusMovedPermanently)
			return
		}

		stockageForm = StockageForm{true, r.FormValue("name")}
		http.Redirect(w, r, "/user/display", http.StatusSeeOther)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) {
		data := PageAffiche{stockageForm.CheckValue, stockageForm.Value, (!stockageForm.CheckValue && stockageForm.Value == "")}
		temp.ExecuteTemplate(w, "FormulaireResultat", data)
	})

	http.HandleFunc("/erreur", func(w http.ResponseWriter, r *http.Request) {
		code, message := r.FormValue("code"), r.FormValue("message")
		if code != "" && message != "" {
			fmt.Fprintf(w, "Erreur %s - %s", code, message)
			return
		}
		fmt.Fprint(w, "Oups une erreur est survenue")
	})
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fileserver))

	http.ListenAndServe("localhost:8000", nil)
}
