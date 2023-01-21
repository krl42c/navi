package main

import (
	"encoding/gob"
	"os"
)

/* File encoder for nvv bytecode */

func write_instruction(instruction ins) {
}

func write(stack []ins) {
	for _, instruction := range stack {
		write_instruction(instruction)
	}
}

// https://pkg.go.dev/encoding/gob
func generate_code(file_path string, stack []ins) error {
	err := write_gob(file_path, stack)
	if err != nil {
		return err
	}

	return nil
}

func write_gob(filePath string, object interface{}) error {
	file, err := os.Create(filePath)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(object)
	}
	file.Close()
	return err
}
