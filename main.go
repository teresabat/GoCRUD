package main

import (
	"crud-golang/database"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

type Car struct {
	ID     int
	Marca  string
	Modelo string
	Ano    int
}

func main() {
	// Conectar ao banco de dados
	database.Connect()

	// Configurar o roteador
	r := mux.NewRouter()

	// Definir as rotas no roteador
	r.HandleFunc("/", ListCars).Methods("GET")
	r.HandleFunc("/create", CreateCar).Methods("GET", "POST")
	r.HandleFunc("/edit/{id}", EditCar).Methods("GET", "POST")
	r.HandleFunc("/delete/{id}", DeleteCar).Methods("POST")

	// Servir arquivos estáticos (CSS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	// Iniciar o servidor
	http.ListenAndServe(":8080", r)
}

// Função para listar carros
func ListCars(w http.ResponseWriter, r *http.Request) {
	var cars []Car
	rows, err := database.DB.Query("SELECT id, marca, modelo, ano FROM cars")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var car Car
		err := rows.Scan(&car.ID, &car.Marca, &car.Modelo, &car.Ano)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		cars = append(cars, car)
	}

	tmpl.ExecuteTemplate(w, "index.html", cars)
}

// Função para criar um carro
func CreateCar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		marca := r.FormValue("marca")
		modelo := r.FormValue("modelo")
		ano := r.FormValue("ano")

		_, err := database.DB.Exec("INSERT INTO cars(marca, modelo, ano) VALUES(?, ?, ?)", marca, modelo, ano)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		tmpl.ExecuteTemplate(w, "create.html", nil)
	}
}

// Função para editar um carro
func EditCar(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if r.Method == "POST" {
		marca := r.FormValue("marca")
		modelo := r.FormValue("modelo")
		ano := r.FormValue("ano")

		_, err := database.DB.Exec("UPDATE cars SET marca=?, modelo=?, ano=? WHERE id=?", marca, modelo, ano, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		var car Car
		err := database.DB.QueryRow("SELECT id, marca, modelo, ano FROM cars WHERE id=?", id).Scan(&car.ID, &car.Marca, &car.Modelo, &car.Ano)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		tmpl.ExecuteTemplate(w, "edit.html", car)
	}
}

// Função para deletar um carro
func DeleteCar(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := database.DB.Exec("DELETE FROM cars WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
