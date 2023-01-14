package main

import (
	"errors"
	"math/rand"
)

type database struct {
	tables []table
	index  []string
}

type table struct {
	id   uint16
	name string
	cols column[string]
}

type column[T int32 | string] struct {
	name      string
	value     T
	refers_to *table
}

/* Core functionalities */

func create_table(db *database, name string) (tbl table, err error) {
	ret := table{}
	for _, ref := range db.index {
		if ref == name {
			return ret, errors.New("EXISTS")
		}
	}
	ret.name = name
	ret.id = uint16(rand.Uint32())
	db.index = append(db.index, name)
	db.tables = append(db.tables, ret)
	return ret, nil
}

func delete_table_by_name(db *database, tbl_name string) bool {
	for i, table := range db.tables {
		if table.name == tbl_name {
			remove(db.tables, i)
			return true
		}
	}
	return false
}

func get_table(db *database, tbl_name string) (tbl *table, err error) {
	for _, table := range db.tables {
		if table.name == tbl_name {
			return &table, nil
		}
	}
	return nil, err
}

func insert_column_str(db *database, tbl_name string, col column[string]) {

}

func insert_column_int(db *database, tbl_name string, col column[int32]) {

}

func remove(s []table, i int) []table {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
