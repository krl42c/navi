package main

import (
	"testing"
)

func Test_create_table(t *testing.T) {
	db := database{}
	tbl, err := create_table(&db, "test")

	if err != nil {
		t.Fatalf("Failed to create table %s", tbl.name)
	}

	tbl, err = create_table(&db, "test")

	if err == nil {
		t.Fatalf("Test failed, duplicated table was created")
	}
}

func Test_remove_table_by_name(t *testing.T) {
	db := database{}
	tbl, err := create_table(&db, "test")

	if err != nil {
		t.Fatalf("Failed to create table %s", tbl.name)
	}

	if !delete_table_by_name(&db, "test") {
		t.Fatalf("Test failed, couldn't delete table")
	}
}

func Test_get_table(t *testing.T) {
	db := database{}
	tbl, err := create_table(&db, "test")

	if err != nil {
		t.Fatalf("Failed to create table %s", tbl.name)
	}

	find_tbl, err := get_table(&db, "test")

	if err != nil {
		t.Fatalf("Failed to find table %s", find_tbl.name)
	}
}

func Test_creation(t *testing.T) {
	db := database{name: "navi"}
	users_tbl, err := create_table(&db, "users")
	if err != nil {
		t.Fatalf("Failed to create table %s", users_tbl.name)
	}
	col := column[string]{name: "username"}
	if !insert_column_str(&db, "users", col) {
		t.Fatalf("Failed to insert column %s", col.name)
	}
}
