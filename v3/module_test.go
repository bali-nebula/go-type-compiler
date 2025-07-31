/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package module_test

import (
	fmt "fmt"
	not "github.com/bali-nebula/go-document-notation/v3"
	com "github.com/bali-nebula/go-type-compiler/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const directory = "./test/"

func TestFormattingBytecode(t *tes.T) {
	var instructions = fra.List[com.InstructionLike]()
	var instruction = com.Instruction(0, 0, 0) // SKIP instruction.
	instructions.AppendValue(instruction)
	var operation uint16
	var modifier uint16
	var operand uint16 = 100
	for operation = 0; operation < 8; operation++ {
		for modifier = 0; modifier < 4; modifier++ {
			if operation == 2 {
				// The PULL instruction has no operand.
				instruction = com.Instruction(
					com.Operation(operation<<13),
					com.Modifier(modifier<<11),
					0,
				)
			} else {
				instruction = com.Instruction(
					com.Operation(operation<<13),
					com.Modifier(modifier<<11),
					com.Operand(operand),
				)
			}
			instructions.AppendValue(instruction)
			operand++
		}
	}
	ass.Equal(t, 33, int(instructions.GetSize()))
	var _ = com.Bytecode(instructions)
	var formatted = com.FormatInstructions(instructions)
	var filename = directory + "instructions.txt"
	uti.WriteFile(filename, formatted)
}

func TestMethodAssembler(t *tes.T) {
	var filename = directory + "type.bali"
	var source = uti.ReadFile(filename)
	var type_ = not.ParseSource(source)
	var key1 = not.Primitive(not.Element("$methods"))
	var key2 = not.Primitive(not.Element("$example"))
	var method = not.GetAttribute(type_, key1, key2)
	var assembler = com.MethodAssembler(type_)
	assembler.AssembleMethod(method)
	source = not.FormatDocument(type_)
	uti.WriteFile(filename, source)
	var key = not.Primitive(not.Element("$bytecode"))
	var document = not.GetAttribute(method, key)
	var component = document.GetComponent()
	source = component.GetAny().(not.StringLike).GetAny().(string)
	var bytecode = com.BytecodeFromString(source)
	var instructions = bytecode.GetInstructions()
	fmt.Println(com.FormatInstructions(instructions))
}
