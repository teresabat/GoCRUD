package controllers

import (
	"crud-golang/models"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
    cars, err := models.GetAllCars()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    tmpl.ExecuteTemplate(w, "Index", cars)
}

func Create(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        car := models.Car{
            Marca:  r.FormValue("marca"),
            Modelo: r.FormValue("modelo"),
            Ano:    atoi(r.FormValue("ano")),
        }
        err := models.CreateCar(car)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
    tmpl.ExecuteTemplate(w, "Create", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    car, err := models.GetCarByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if r.Method == "POST" {
        car.Marca = r.FormValue("marca")
        car.Modelo = r.FormValue("modelo")
        car.Ano = atoi(r.FormValue("ano"))
        err := models.UpdateCar(car)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        http.Redirect(w, r, "/", http.StatusSeeOther)
    }
    tmpl.ExecuteTemplate(w, "Edit", car)
}

func Delete(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    err := models.DeleteCar(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func atoi(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}
