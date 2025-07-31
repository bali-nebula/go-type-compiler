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

/*
Package "assembler" provides an implementation of an assembler that can be used
to assemble compiled methods into bytecode for the Bali Virtual Machine™.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-type-compiler/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-development-tools/wiki/Coding-Conventions

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package assembler

import (
	lan "github.com/bali-nebula/go-assembly-language/v3"
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

/*
Operation is a constrained type representing the possible BVM operations.
*/
type Operation uint16

const (
	Jump Operation = 0b0000000000000000
	Push Operation = 0b0010000000000000
	Pull Operation = 0b0100000000000000
	Load Operation = 0b0110000000000000
	Save Operation = 0b1000000000000000
	Drop Operation = 0b1010000000000000
	Call Operation = 0b1100000000000000
	Send Operation = 0b1110000000000000
)

/*
Modifier is a constrained type representing the possible BVM modifiers.
*/
type Modifier uint16

const (
	// Jump Operation
	Any   Modifier = 0b0000000000000000
	Empty Modifier = 0b0000100000000000
	None  Modifier = 0b0001000000000000
	False Modifier = 0b0001100000000000

	// Push Operation
	Handler  Modifier = 0b0000000000000000
	Literal  Modifier = 0b0000100000000000
	Constant Modifier = 0b0001000000000000
	Argument Modifier = 0b0001100000000000

	// Pull Operation
	// Handler is defined above.
	Exception Modifier = 0b0000100000000000
	Component Modifier = 0b0001000000000000
	Result    Modifier = 0b0001100000000000

	// Load, Save and Drop Operations
	Contract Modifier = 0b0000000000000000
	Draft    Modifier = 0b0000100000000000
	Message  Modifier = 0b0001000000000000
	Variable Modifier = 0b0001100000000000

	// Call Operation
	With0Arguments Modifier = 0b0000000000000000
	With1Arguments Modifier = 0b0000100000000000
	With2Arguments Modifier = 0b0001000000000000
	With3Arguments Modifier = 0b0001100000000000

	// Send Operation
	// Contract and Component are defined above.
	ContractWithArguments  Modifier = 0b0000100000000000
	ComponentWithArguments Modifier = 0b0001100000000000
)

/*
Operand is a constrained type representing the possible BVM operands.
*/
type Operand uint16

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
BytecodeClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete bytecode-like class.
*/
type BytecodeClassLike interface {
	// Constructor Methods
	Bytecode(
		instructions fra.Sequential[InstructionLike],
	) BytecodeLike
	BytecodeFromString(
		source string,
	) BytecodeLike
}

/*
InstructionClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete instruction-like class.
*/
type InstructionClassLike interface {
	// Constructor Methods
	Instruction(
		operation Operation,
		modifier Modifier,
		operand Operand,
	) InstructionLike
	InstructionFromInteger(
		integer uint16,
	) InstructionLike

	// Constant Methods
	OperationMask() uint16
	ModifierMask() uint16
	OperandMask() uint16

	// Function Methods
	FormatInstructions(
		instructions fra.Sequential[InstructionLike],
	) string
}

/*
MethodAssemblerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete method-assembler-like class.
*/
type MethodAssemblerClassLike interface {
	// Constructor Methods
	MethodAssembler(
		type_ not.DocumentLike,
	) MethodAssemblerLike
}

// INSTANCE DECLARATIONS

/*
BytecodeLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete bytecode-like class.
*/
type BytecodeLike interface {
	// Principal Methods
	GetClass() BytecodeClassLike
	AsString() string

	// Attribute Methods
	GetInstructions() fra.Sequential[InstructionLike]

	// Aspect Interfaces
	fra.Sequential[InstructionLike]
}

/*
InstructionLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete instruction-like class.
*/
type InstructionLike interface {
	// Principal Methods
	GetClass() InstructionClassLike
	AsIntrinsic() uint16
	AsString() string
	OperationAsString() string
	ModifierAsString() string
	OperandAsString() string

	// Attribute Methods
	GetOperation() Operation
	GetModifier() Modifier
	GetOperand() Operand
}

/*
MethodAssemblerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete method-assembler-like class.
*/
type MethodAssemblerLike interface {
	// Principal Methods
	GetClass() MethodAssemblerClassLike
	AssembleMethod(
		method not.DocumentLike,
	)

	// Aspect Interfaces
	lan.Methodical
}

// ASPECT DECLARATIONS
