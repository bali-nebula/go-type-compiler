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
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│         This "module_api.go" file was automatically generated using:         │
│            https://github.com/craterdog/go-development-tools/wiki            │
│                                                                              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-type-compiler/wiki
*/
package module

import (
	not "github.com/bali-nebula/go-document-notation/v3"
	rep "github.com/bali-nebula/go-document-repository/v3"
	ass "github.com/bali-nebula/go-type-compiler/v3/assembler"
	com "github.com/bali-nebula/go-type-compiler/v3/compiler"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE ALIASES

// Assembler

type (
	Operation = ass.Operation
	Modifier  = ass.Modifier
	Operand   = ass.Operand
)

const (
	// Operations
	Jump = ass.Jump
	Push = ass.Push
	Pull = ass.Pull
	Load = ass.Load
	Save = ass.Save
	Drop = ass.Drop
	Call = ass.Call
	Send = ass.Send

	// Jump Modifiers
	Any   = ass.Any
	Empty = ass.Empty
	None  = ass.None
	False = ass.False

	// Push Modifiers
	Handler  = ass.Handler
	Literal  = ass.Literal
	Constant = ass.Constant
	Argument = ass.Argument

	// Pull Modifiers
	Exception = ass.Exception
	Component = ass.Component
	Result    = ass.Result

	// Load, Save and Drop Modifiers
	Contract = ass.Contract
	Draft    = ass.Draft
	Message  = ass.Message
	Variable = ass.Variable

	// Call Modifiers
	With0Arguments = ass.With0Arguments
	With1Arguments = ass.With1Arguments
	With2Arguments = ass.With2Arguments
	With3Arguments = ass.With3Arguments

	// Send Modifiers
	ComponentWithArguments = ass.ComponentWithArguments
	ContractWithArguments  = ass.ContractWithArguments
)

type (
	BytecodeClassLike        = ass.BytecodeClassLike
	InstructionClassLike     = ass.InstructionClassLike
	MethodAssemblerClassLike = ass.MethodAssemblerClassLike
)

type (
	BytecodeLike        = ass.BytecodeLike
	InstructionLike     = ass.InstructionLike
	MethodAssemblerLike = ass.MethodAssemblerLike
)

// Compiler

type (
	TypeCompilerClassLike = com.TypeCompilerClassLike
)

type (
	TypeCompilerLike = com.TypeCompilerLike
)

// CLASS ACCESSORS

// Assembler

func BytecodeClass() BytecodeClassLike {
	return ass.BytecodeClass()
}

func Bytecode(
	instructions fra.Sequential[ass.InstructionLike],
) BytecodeLike {
	return BytecodeClass().Bytecode(
		instructions,
	)
}

func BytecodeFromString(
	source string,
) BytecodeLike {
	return BytecodeClass().BytecodeFromString(
		source,
	)
}

func InstructionClass() InstructionClassLike {
	return ass.InstructionClass()
}

func Instruction(
	operation ass.Operation,
	modifier ass.Modifier,
	operand ass.Operand,
) InstructionLike {
	return InstructionClass().Instruction(
		operation,
		modifier,
		operand,
	)
}

func InstructionFromInteger(
	integer uint16,
) InstructionLike {
	return InstructionClass().InstructionFromInteger(
		integer,
	)
}

func FormatInstructions(
	instructions fra.Sequential[InstructionLike],
) string {
	return InstructionClass().FormatInstructions(instructions)
}

func MethodAssemblerClass() MethodAssemblerClassLike {
	return ass.MethodAssemblerClass()
}

func MethodAssembler(
	type_ not.DocumentLike,
) MethodAssemblerLike {
	return MethodAssemblerClass().MethodAssembler(type_)
}

// Compiler

func TypeCompilerClass() TypeCompilerClassLike {
	return com.TypeCompilerClass()
}

func TypeCompiler(
	repository rep.DocumentRepositoryLike,
) TypeCompilerLike {
	return TypeCompilerClass().TypeCompiler(
		repository,
	)
}

// GLOBAL FUNCTIONS
