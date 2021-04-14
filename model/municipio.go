package model

type Municipio struct {
	ID        uint64 `db:"id" json:"id"`
	Nome      string `db:"nome" json:"nome"`
	Prefeito  string `db:"prefeito" json:"prefeito"`
	Populacao uint64 `db:"populacao" json:"populacao"`
	IDEstado  uint64 `db:"id_estado_fk" json:"id_estado_fk"`
}
