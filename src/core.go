package main

import (
	"errors"
	"math/rand"
)

type database struct {
	name   string
	tables []table
	index  []string
}

type table struct {
	id       uint16
	name     string
	cols_str []column[string]
	cols_int []column[int32]
}

type column[T int32 | string] struct {
	name      string
	refers_to *table
	rows      []row[T]
}

type row[T int32 | string] struct {
	index int32
	value T
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

func get_column[T int32 | string](tbl *table, name string) (col *column[T], err error) {
	for _, column := range tbl.cols_int {
		if column.name == name {
			return col, nil
		}
	}
	return nil, err
}

func insert_column_str(db *database, tbl_name string, col column[string]) bool {
	tbl, err := get_table(db, tbl_name)
	if err != nil {
		return false
	}
	tbl.cols_str = append(tbl.cols_str, col)
	return true
}

func insert_column_int(db *database, tbl_name string, col column[int32]) bool {
	tbl, err := get_table(db, tbl_name)
	if err != nil {
		return false
	}
	tbl.cols_int = append(tbl.cols_int, col)
	return true
}

func insert_row[T int32 | string](db *database, tbl_name string, col_name string, value row[T]) bool {
	tbl, err := get_table(db, tbl_name)
	if err != nil {
		return false
	}
	col, err := get_column[T](tbl, col_name)

	if err != nil {
		return false
	}
	col.rows = append(col.rows, value)
	return true
}

func remove(s []table, i int) []table {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
