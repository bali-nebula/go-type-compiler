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

package assembler

import (
	fmt "fmt"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	str "strings"
)

// CLASS INTERFACE

// Access Function

func BytecodeClass() BytecodeClassLike {
	return bytecodeClass()
}

// Constructor Methods

func (c *bytecodeClass_) Bytecode(
	instructions fra.Sequential[InstructionLike],
) BytecodeLike {
	if uti.IsUndefined(instructions) {
		panic("The \"instructions\" attribute is required by this class.")
	}
	var instance = &bytecode_{
		// Initialize the instance attributes.
		instructions_: instructions,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *bytecode_) GetClass() BytecodeClassLike {
	return bytecodeClass()
}

func (v *bytecode_) AsString() string {
	var result str.Builder
	result.WriteString(`
 Addr   Bytes   Bytecode               Instruction
---------------------------------------------------------------
`,
	)
	var counter uint16
	var iterator = v.GetIterator()
	for counter = 1; iterator.HasNext(); counter++ {
		var instruction = iterator.GetNext()
		var address = fmt.Sprintf("[%03x]", counter)
		var bytes = fmt.Sprintf("x%04x", instruction.AsIntrinsic())
		var operation = instruction.GetOperation() >> 13
		var modifier = instruction.GetModifier() >> 11
		var operand = v.operandAsString(
			instruction.GetOperation(),
			instruction.GetModifier(),
			instruction.GetOperand(),
		)
		var bytecode = fmt.Sprintf("%d%d %s", operation, modifier, operand)
		var description = instruction.AsString()
		var line = fmt.Sprintf(
			"%s:  %s   %s  %s\n",
			address,
			bytes,
			bytecode,
			description,
		)
		result.WriteString(line)
	}
	return result.String()
}

// Attribute Methods

// Sequential[InstructionLike] Methods

func (v *bytecode_) IsEmpty() bool {
	return v.instructions_.IsEmpty()
}

func (v *bytecode_) GetSize() uti.Cardinal {
	return v.instructions_.GetSize()
}

func (v *bytecode_) AsArray() []InstructionLike {
	return v.instructions_.AsArray()
}

func (v *bytecode_) GetIterator() fra.IteratorLike[InstructionLike] {
	return v.instructions_.GetIterator()
}

// PROTECTED INTERFACE

// Private Methods

func (v *bytecode_) operandAsString(
	operation Operation,
	modifier Modifier,
	operand Operand,
) string {
	var result string
	if operation == Jump || (operation == Push && modifier == Handler) {
		// Treat the operand as an address "[xxx]".
		result = fmt.Sprintf("[%03x]", operand)
	} else {
		// Treat the operand as an index "dddd ".
		result = fmt.Sprintf("%4d ", operand)
	}
	return result
}

// Instance Structure

type bytecode_ struct {
	// Declare the instance attributes.
	instructions_ fra.Sequential[InstructionLike]
}

// Class Structure

type bytecodeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func bytecodeClass() *bytecodeClass_ {
	return bytecodeClassReference_
}

var bytecodeClassReference_ = &bytecodeClass_{
	// Initialize the class constants.
}
