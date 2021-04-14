package db

import (
	"errors"
	"log"

	"github.com/natajonasdacoliveira/credisis-unidades-federativas/model"
)

func GetAllEstados() ([]model.Estado, error) {
	db := InitDb()

	estados:= []model.Estado{}
	
	err:= db.Select(&estados, "SELECT * FROM estados")
	if err != nil {
		log.Println(err)
		return []model.Estado{}, err
	}

	return estados, err
}

func CreateEstado(estado model.Estado) (err error) {
	db := InitDb()
	tx := db.MustBegin()

	tx.MustExec("INSERT INTO estados(nome, sigla) VALUES(?, ?)", estado.Nome, estado.Sigla)
	tx.Commit()

	return err
}

func UpdateEstado(estado model.Estado) (err error) {
	db := InitDb()
	tx := db.MustBegin()

	oldEstado:= model.Estado{}

	err = db.Get(&oldEstado, "SELECT * FROM estados WHERE id= ? LIMIT 1", estado.ID)
	if(oldEstado.ID < 1 || err != nil) {
		return errors.New("estado não cadastrado")
	}
	
	tx.MustExec("UPDATE estados SET nome = ?, sigla = ? WHERE id = ?", estado.Nome, estado.Sigla, estado.ID)
	tx.Commit()

	return err
}

func DeleteEstado(estado model.Estado) (err error) {
	db := InitDb()
	tx := db.MustBegin()

	oldEstado:= model.Estado{}

	err = db.Get(&oldEstado, "SELECT * FROM estados WHERE id= ? LIMIT 1", estado.ID)
	if(oldEstado.ID < 1 || err != nil) {
		return errors.New("estado não cadastrado")
	}

	tx.MustExec("DELETE FROM estados WHERE id = ?", estado.ID)
	tx.Commit()

	return err
}