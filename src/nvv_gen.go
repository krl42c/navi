package main

type opcode byte

// Opcodes taken from https://www.sqlite.org/opcode.html
const (
	ADD       opcode = iota // Add the value in register P1 to the value in register P2 and store the result in register P3. If either input is NULL, the result is NULL.
	ADDIMM                  // Add the constant P2 to the value in register P1. The result is always an integer.
	AND                     // Add the value in register P1 to the value in register P2 and store the result in register P3. If either input is NULL, the result is NULL.
	CAST                    // Force the value in register P1 to be the type defined by P2.
	CLEAR                   // Delete all contents of the database table or index whose root page in the database file is given by P1. But, unlike Destroy, do not remove the table or index from the database file.
	DELETE                  // Delete the record at which the P1 cursor is currently pointing.
	DESTROY                 // Delete an entire database table or index whose root page in the database file is given by P1.
	DROPTABLE               // Remove the internal (in-memory) data structures that describe the table named P4 in database P1. This is called after a table is dropped from disk (using the Destroy opcode) in order to keep the internal representation of the schema consistent with what is on disk.
	GOTO                    // An unconditional jump to address P2. The next instruction executed will be the one at index P2 from the beginning of the program.
	IF                      //Jump to P2 if the value in register P1 is true. The value is considered true if it is numeric and non-zero. If the value in P1 is NULL then take the jump if and only if P3 is non-zero.
	IFNOT                   // Jump to P2 if the value in register P1 is False. The value is considered false if it has a numeric value of zero. If the value in P1 is NULL then take the jump if and only if P3 is non-zero.
	_INSERT                 // Write an entry into the table of cursor P1. A new entry is created if it doesn't already exist or the data for an existing entry is overwritten. The data is the value MEM_Blob stored in register number P2. The key is stored in register P3. The key must be a MEM_Int.
	INTEGER                 // The 32-bit integer value P1 is written into register P2.
	STRING8                 // The string value P4 of length P1 (bytes) is stored in register P2.
	ISNULL                  // Jump to P2 if the value in register P1 is NULL.
	ISTRUE                  // This opcode implements the IS TRUE, IS FALSE, IS NOT TRUE, and IS NOT FALSE operators.
	ISTYPE                  // Jump to P2 if the type of a column in a btree is one of the types specified by the P5 bitmask.
	JUMP                    // Jump to the instruction at address P1, P2, or P3 depending on whether in the most recent Compare instruction the P1 vector was less than equal to, or greater than the P2 vector, respectively.
	OPEN_WRITE
	TRANSACTION
	CLOSE
	COMMIT
	HALT
)

type ins struct {
	addr uint16
	op   opcode
	p1   int16
	p2   int16
	p3   string
}

/* Utility functions to generate commonly used instructions with fixed register data */

func _commit(addr uint16) ins {
	return ins{addr: addr, op: COMMIT, p1: 0, p2: 0, p3: "0"}
}

func _halt(addr uint16) ins {
	return ins{addr: addr, op: HALT, p1: 0, p2: 0, p3: "0"}
}

func _close(addr uint16) ins {
	return ins{addr: addr, op: CLOSE, p1: 0, p2: 0, p3: "0"}
}

func _open_write(addr uint16, tbl_name string) ins {
	return ins{addr: addr, op: OPEN_WRITE, p1: 0, p2: 0, p3: tbl_name}
}

func nvv_start_transaction() ins {
	transaction := ins{addr: 0, op: TRANSACTION, p1: 0, p2: 0, p3: "0"}
	return transaction
}

/* Code translation into instruction sets */

func nvv_insert(db *database, st statement) int32 {
	tbl_name := st.tokens[1].tvalue
	var addr uint16 = 1
	instruction_stack := []ins{}
	if st.stype == SINSERT {
		// Prepare before insert
		instruction_stack = append(instruction_stack, nvv_start_transaction())
		addr++
		instruction_stack = append(instruction_stack, _open_write(addr, tbl_name))

		// TODO: Insert operations here
		params := st.tokens[3:]
		for _, param := range params {
			if param.ttype != CLOSED_PAREN && param.ttype != ENDLINE {
				insert := ins{addr: addr, op: _INSERT, p1: 0, p2: 0, p3: param.tvalue}
				instruction_stack = append(instruction_stack, insert)
				addr++
			}
		}
		// Close
		close := ins{addr: addr, op: CLOSE, p1: 0, p2: 0, p3: "0"}
		addr++
		commit := ins{addr: addr, op: COMMIT, p1: 0, p2: 0, p3: "0"}
		addr++
		halt := ins{addr: addr, op: HALT, p1: 0, p2: 0, p3: "0"}
		addr++
		instruction_stack = append(instruction_stack, close, commit, halt)
	}
	generate_code("insert.nvvbc", instruction_stack)
	return 0
}

func nvv_create_table(db *database, tbl_name string, addr uint16) {
	// Table creation setup
	instruction_stack := []ins{}
	open_write := ins{addr: addr, op: OPEN_WRITE, p1: 0, p2: 0, p3: tbl_name}
	addr++
	str_len := int16(len(tbl_name))
	string8 := ins{addr: addr, op: STRING8, p1: 0, p2: str_len, p3: tbl_name}
	addr++

	// TODO: Creation operation

	halt := ins{addr: addr, op: HALT, p1: 0, p2: 0, p3: "0"}
	instruction_stack = append(instruction_stack, open_write, string8)
	instruction_stack = append(instruction_stack, halt)
}
