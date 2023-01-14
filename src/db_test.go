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

	if tbl != *find_tbl {
		t.Fatalf("Err, table retrieved is diferent")
	}
}
