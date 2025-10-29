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
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	mat "math"
	reg "regexp"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func BytecodeClass() BytecodeClassLike {
	return bytecodeClass()
}

// Constructor Methods

func (c *bytecodeClass_) Bytecode(
	instructions com.Sequential[InstructionLike],
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

func (c *bytecodeClass_) BytecodeFromString(
	source string,
) BytecodeLike {
	var matches = c.matcher_.FindStringSubmatch(source)
	if uti.IsUndefined(matches) {
		var message = fmt.Sprintf(
			"An illegal string was passed to the bytecode constructor method: %s",
			source,
		)
		panic(message)
	}
	var base16 = matches[0]
	base16 = base16[2 : len(base16)-2]        // Strip off the delimiters.
	base16 = sts.ReplaceAll(base16, "\n", "") // Remove all newlines.
	base16 = sts.ReplaceAll(base16, " ", "")  // Remove all spaces.
	var strings = sts.Split(base16, ":")[1:]  // Extract the instructions.
	var instructions = com.List[InstructionLike]()
	for _, hex := range strings {
		var integer, _ = stc.ParseUint(hex, 16, 16)
		var instruction = InstructionClass().InstructionFromInteger(uint16(integer))
		instructions.AppendValue(instruction)
	}
	return c.Bytecode(instructions)
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *bytecode_) GetClass() BytecodeClassLike {
	return bytecodeClass()
}

func (v *bytecode_) AsString() string {
	var source = "'>"
	var newline = "\n    "
	var iterator = v.instructions_.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		if mat.Mod(float64(iterator.GetSlot()-1), 12.0) == 0 {
			source += newline
		}
		source += fmt.Sprintf(":%04x", instruction.AsIntrinsic())
	}
	source += "\n<'"
	return source
}

// Attribute Methods

func (v *bytecode_) GetInstructions() com.Sequential[InstructionLike] {
	return v.instructions_
}

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

func (v *bytecode_) GetIterator() com.Ratcheted[InstructionLike] {
	return v.instructions_.GetIterator()
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type bytecode_ struct {
	// Declare the instance attributes.
	instructions_ com.Sequential[InstructionLike]
}

// Class Structure

type bytecodeClass_ struct {
	// Declare the class constants.
	matcher_ *reg.Regexp
}

// Class Reference

func bytecodeClass() *bytecodeClass_ {
	return bytecodeClassReference_
}

var bytecodeClassReference_ = &bytecodeClass_{
	// Initialize the class constants.
	matcher_: reg.MustCompile(
		"^'>\n(?:(?: )*(?:(?::(?:[0-9]|[a-f]){4})){1,12}\n)+(?: )*<'",
	),
}
