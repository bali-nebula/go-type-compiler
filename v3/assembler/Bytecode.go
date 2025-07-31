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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	mat "math"
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
	var source = "'>"
	var newline = "\n    "
	var iterator = v.instructions_.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		if mat.Mod(float64(iterator.GetSlot()), 12.0) == 0 {
			source += newline
		}
		source += instruction.AsString()
	}
	source += "\n<'"
	return source
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
