package models

import (
	"crud-golang/database"
)

type Car struct {
    ID     int
    Marca  string
    Modelo string
    Ano    int
}

func GetAllCars() ([]Car, error) {
    rows, err := database.DB.Query("SELECT id, marca, modelo, ano FROM cars")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var cars []Car
    for rows.Next() {
        var car Car
        if err := rows.Scan(&car.ID, &car.Marca, &car.Modelo, &car.Ano); err != nil {
            return nil, err
        }
        cars = append(cars, car)
    }

    return cars, nil
}

func CreateCar(car Car) error {
    _, err := database.DB.Exec("INSERT INTO cars (marca, modelo, ano) VALUES (?, ?, ?)", car.Marca, car.Modelo, car.Ano)
    return err
}

func GetCarByID(id int) (Car, error) {
    var car Car
    err := database.DB.QueryRow("SELECT id, marca, modelo, ano FROM cars WHERE id = ?", id).Scan(&car.ID, &car.Marca, &car.Modelo, &car.Ano)
    return car, err
}

func UpdateCar(car Car) error {
    _, err := database.DB.Exec("UPDATE cars SET marca = ?, modelo = ?, ano = ? WHERE id = ?", car.Marca, car.Modelo, car.Ano, car.ID)
    return err
}

func DeleteCar(id int) error {
    _, err := database.DB.Exec("DELETE FROM cars WHERE id = ?", id)
    return err
}
