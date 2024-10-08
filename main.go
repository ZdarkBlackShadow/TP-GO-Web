package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

// structures
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
	Nom        string
	Prenom     string
	Date       string
	Sexe       string
	IsEmpty    bool
}

// variable globales pour récuperer les données de l'utlisateur dans le formulaires
var stockageFormNom = StockageForm{false, ""}
var stockageFormPrenom = StockageForm{false, ""}
var stockageFormDate = StockageForm{false, ""}
var stockageFormSexe = StockageForm{false, ""}

func main() {
	temp, err := template.ParseGlob("./Templates/*.html")
	Compteur := Change{true, 0} //initialisation du compteur
	if err != nil {
		fmt.Printf("Erreur => %s\n", err.Error())
		os.Exit(02)
	}
	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) { //route pour voir la promo
		LE := []Etudiant{
			{"Lecomte", "Adrien", 20, true},
			{"Perez", "Jonathan", 19, true},
			{"Lili", "Rosello", 19, false}}
		data := Promo{" B1 Cybersécurité", "Cybersécurité", "Bachelor 1", len(LE), LE}
		err1 := temp.ExecuteTemplate(w, "promo", data)
		if err1 != nil {
			log.Fatal(err1)
		}
	})
	http.HandleFunc("/change", func(w http.ResponseWriter, r *http.Request) { //route pour change
		Compteur.Compteur += 1
		Compteur.Pair = Compteur.Compteur%2 == 0
		temp.ExecuteTemplate(w, "change", Compteur)
	})

	http.HandleFunc("/user/form", func(w http.ResponseWriter, r *http.Request) { //route pour afficher le formulaire
		temp.ExecuteTemplate(w, "AffichageFormulaire", nil)
	})

	http.HandleFunc("/user/traitement", func(w http.ResponseWriter, r *http.Request) { //route pour le traitement des données
		if r.Method != http.MethodPost {
			http.Redirect(w, r, "/erreur?code=400&message=Oups méthode incorecte", http.StatusMovedPermanently)
			return
		}
		//Verification pour le nom
		checkValueNom, _ := regexp.MatchString("^[\\p{L}-]{1,32}$", r.FormValue("nom"))
		if !checkValueNom {
			stockageFormNom = StockageForm{false, ""}
			http.Redirect(w, r, "/erreur?code=400&message=Oups les données de Nom sont invalides", http.StatusMovedPermanently)
			return
		}
		stockageFormNom = StockageForm{true, r.FormValue("nom")}
		//verification pour le prenom
		checkValuePrenom, _ := regexp.MatchString("^[\\p{L}-]{1,32}$", r.FormValue("prenom"))
		if !checkValuePrenom {
			stockageFormPrenom = StockageForm{false, ""}
			http.Redirect(w, r, "/erreur?code=400&message=Oups les données de prenom sont invalides", http.StatusMovedPermanently)
			return
		}
		stockageFormPrenom = StockageForm{true, r.FormValue("prenom")}
		//Date de naissance
		stockageFormDate = StockageForm{true, r.FormValue("date")}
		//verification du sexe
		sexe := r.FormValue("sexe")
		if sexe != "masculin" && sexe != "feminin" && sexe != "autre" {
			http.Redirect(w, r, "/erreur?code=400&message=Valeur du sexe invalide", http.StatusMovedPermanently)
			return
		}
		stockageFormSexe = StockageForm{true, sexe}
		http.Redirect(w, r, "/user/display", http.StatusSeeOther)
	})

	http.HandleFunc("/user/display", func(w http.ResponseWriter, r *http.Request) { //route pour l'affichage du résultat du formulaires
		data := PageAffiche{stockageFormNom.CheckValue && stockageFormPrenom.CheckValue && stockageFormDate.CheckValue && stockageFormSexe.CheckValue, stockageFormNom.Value, stockageFormPrenom.Value, stockageFormDate.Value, stockageFormSexe.Value, (!stockageFormNom.CheckValue && stockageFormNom.Value == "" && !stockageFormPrenom.CheckValue && stockageFormPrenom.Value == "" && !stockageFormDate.CheckValue && stockageFormDate.Value == "" && !stockageFormSexe.CheckValue && stockageFormSexe.Value == "")}
		temp.ExecuteTemplate(w, "FormulaireResultat", data)
	})

	//gestion d'erreur
	http.HandleFunc("/erreur", func(w http.ResponseWriter, r *http.Request) {
		code, message := r.FormValue("code"), r.FormValue("message")
		if code != "" && message != "" {
			fmt.Fprintf(w, "Erreur %s - %s", code, message)
			return
		}
		fmt.Fprint(w, "Oups une erreur est survenue")
	})
	fileserver := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	http.ListenAndServe("localhost:8000", nil)
}
