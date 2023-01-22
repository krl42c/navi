package main

import (
	"testing"
)

func Test_gen_should_produce_file(t *testing.T) {

}

func Test_nvv_insert(t *testing.T) {
	/*
		Excepted bytecode:
		ADDR 	 OP				P1		P2		P3
		0 		 TRANSACTION 	0 		 0 		 0
		1 		 OPEN_WRITE 	0 		 0 		 groups
		2 		 INSERT 		0 		 0 		 val1
		3 		 INSERT 		0 		 0 		 val2
		4 		 INSERT 		0 		 0 		 val3
		5 		 INSERT 		0 		 0 		 val4
		6 		 CLOSE 		 	0 		 0 		 0
		7 		 COMMIT 		0 		 0 		 0
		8 		 HALT 		 	0 		 0 		 0
	*/

	db := database{name: "Test_DB"}
	query := "insert groups (val1 val2 val3 val4)"
	_, err := execute_line(&db, query, 0)

	if err != nil {
		t.Fatalf("Couldn't transform query into bytecode: %s", err.Error())
	}

}
