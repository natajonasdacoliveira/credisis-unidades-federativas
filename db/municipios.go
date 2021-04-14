package db

import (
	"errors"
	"log"

	"github.com/natajonasdacoliveira/credisis-unidades-federativas/model"
)

func GetAllMunicipios() ([]model.Municipio, error) {
	db := InitDb()

	municipios:= []model.Municipio{}
	
	err:= db.Select(&municipios, "SELECT * FROM municipios")
	if err != nil {
		log.Println(err)
		return []model.Municipio{}, err
	}

	return municipios, err
}

func CreateMunicipio(municipio model.Municipio) (err error) {
	db := InitDb()
	tx := db.MustBegin()

	estado:= model.Estado{}

	err = db.Get(&estado, "SELECT * FROM estados WHERE id= ? LIMIT 1", municipio.IDEstado)
	if(estado.ID < 1 || err != nil) {
		return errors.New("estado não cadastrado")
	}
	
	tx.MustExec("INSERT INTO municipios(nome, prefeito, populacao, id_estado_fk) VALUES(?, ?, ?, ?)", municipio.Nome, municipio.Prefeito, municipio.Populacao, municipio.IDEstado)
	tx.Commit()

	return err
}

func UpdateMunicipio(municipio model.Municipio) (err error) {
	db := InitDb()
	tx := db.MustBegin()

	estado:= model.Estado{}
	oldMunicipio:= model.Municipio{}

	err = db.Get(&oldMunicipio, "SELECT * FROM municipios WHERE id= ? LIMIT 1", municipio.ID)
	if(oldMunicipio.ID < 1 || err != nil) {
		return errors.New("município não cadastrado")
	}

	err = db.Get(&estado, "SELECT * FROM estados WHERE id= ? LIMIT 1", municipio.IDEstado)
	if(estado.ID < 1 || err != nil) {
		return errors.New("estado não cadastrado")
	}

	
	tx.MustExec("UPDATE municipios SET nome = ?, prefeito = ?, populacao = ?, id_estado_fk = ? WHERE id = ?", municipio.Nome, municipio.Prefeito, municipio.Populacao, municipio.IDEstado, municipio.ID)
	tx.Commit()

	return err
}

func DeleteMunicipio(municipio model.Municipio) (err error) {
	db := InitDb()
	tx := db.MustBegin()

	oldMunicipio:= model.Municipio{}

	err = db.Get(&oldMunicipio, "SELECT * FROM municipios WHERE id= ? LIMIT 1", municipio.ID)
	if(oldMunicipio.ID < 1 || err != nil) {
		return errors.New("município não cadastrado")
	}

	tx.MustExec("DELETE FROM municipios WHERE id = ?", municipio.ID)
	tx.Commit()

	return err
}