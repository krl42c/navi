package main

import (
	"testing"
)

func Test_gen_should_produce_bytecode(t *testing.T) {
	db := database{name: "navi"}
	ins := nvv_insert(&db, "test", statement{})
	if ins != 0 {
		t.Fatalf("Test not passed")
	}
}
