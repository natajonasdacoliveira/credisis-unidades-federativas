package model

type Estado struct {
	ID    uint64 `db:"id" json:"id"`
	Nome  string `db:"nome" json:"nome"`
	Sigla string `db:"sigla" json:"sigla"`
}
