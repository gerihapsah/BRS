package model

import (
	"db"
)

type Brs struct {
	ID       	 int    `json:"id"`
	Judul        string `json:"judul"`
	Tanggal      string `json:"tanggal"`
	Abstraksi    string `json:"abstraksi"`
	Link         string `json:"link"`
	Prov		 string `json:"prov"`
	Website		 string `json:"website"`	 
}

func BrsAll() ([]Brs, error) {
	var err error
	result := []Brs{}
	sql := "SELECT * FROM tb_brs"
	err = db.SQL.Select(&result, sql)
	return result, standardizeError(err)
}

func BrsById(id string) (Brs, error) {
	var err error
	result := Brs{}
	sql := "SELECT * FROM tb_brs  WHERE id = ?"
	err = db.SQL.Get(&result, sql, id)
	return result, standardizeError(err)
}
func BrsByProv(id string) ([]Brs, error) {
	var err error
	result := []Brs{}
	sql := "SELECT * FROM tb_brs  WHERE prov = ?"
	err = db.SQL.Select(&result, sql, id)
	return result, standardizeError(err)
}
func BrsByTanggal(id string) ([]Brs, error) {
	var err error
	result := []Brs{}
	sql := "SELECT * FROM tb_brs  WHERE tanggal = ?"
	err = db.SQL.Select(&result, sql, id)
	return result, standardizeError(err)
}
func BrsCreate(brs Brs) error {
	var err error
	sql := "INSERT INTO tb_brs (judul, tanggal, abstraksi, link, prov, website) VALUES (?,?,?,?,?,?)"
	_, err = db.SQL.Exec(sql, brs.Judul, brs.Tanggal, brs.Abstraksi, brs.Link, brs.Prov, brs.Website)
	return err
}

func BrsUpdate(id string, brs Brs) error {
	var err error
	sql := "UPDATE tb_brs SET judul=?, tanggal=?, abstraksi=?, link=?, prov=?, website=? WHERE id=?"
	_, err = db.SQL.Exec(sql, brs.Judul, brs.Tanggal, brs.Abstraksi, brs.Link, brs.Prov, brs.Website, id)
	return err
}

func BrsDelete(id string) error {
	var err error
	sql := "DELETE FROM tb_brs WHERE id=?"
	_, err = db.SQL.Exec(sql, id)
	return err
}
