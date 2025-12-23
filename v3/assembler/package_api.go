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
	ins "github.com/bali-nebula/go-bali-instructions/v3"
	com "github.com/craterdog/go-essential-composites/v8"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AnalyzerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete analyzer-like class.
*/
type AnalyzerClassLike interface {
	// Constructor Methods
	Analyzer(
		literals com.Accessible[string],
		constants com.Accessible[string],
		method ins.MethodLike,
	) AnalyzerLike
}

/*
AssemblerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete assembler-like class.
*/
type AssemblerClassLike interface {
	// Constructor Methods
	Assembler() AssemblerLike

	// Function Methods
	FormatInstructions(
		instructions com.Sequential[uint16],
	) string
}

// INSTANCE DECLARATIONS

/*
AnalyzerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete analyzer-like class.
*/
type AnalyzerLike interface {
	// Principal Methods
	GetClass() AnalyzerClassLike
	LookupLiteral(
		literal string,
	) uint16
	LookupConstant(
		constant string,
	) uint16
	LookupArgument(
		argument string,
	) uint16
	LookupVariable(
		variable string,
	) uint16
	LookupIntrinsic(
		intrinsic string,
	) uint16
	LookupMessage(
		message string,
	) uint16
	LookupAddress(
		label string,
	) uint16

	// Aspect Interfaces
	ins.Methodical
}

/*
AssemblerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete assembler-like class.
*/
type AssemblerLike interface {
	// Principal Methods
	GetClass() AssemblerClassLike
	AssembleMethod(
		literals com.Accessible[string],
		constants com.Accessible[string],
		method ins.MethodLike,
	) com.Sequential[uint16]

	// Aspect Interfaces
	ins.Methodical
}

// ASPECT DECLARATIONS
