package main

import (
	"testing"
)

func Test_gen_should_produce_file(t *testing.T) {

}

func Test_nvv_insert(t *testing.T) {
	db := database{name: "Test_DB"}
	query := "insert groups (5 4)"
	_, err := execute_line(&db, query, 0)

	if err != nil {
		t.Fatalf("Couldn't transform query")
	}

}
