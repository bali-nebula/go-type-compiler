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
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
MethodAssemblerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete method-assembler-like class.
*/
type MethodAssemblerClassLike interface {
	// Constructor Methods
	MethodAssembler(
		literals fra.ListLike[string],
		constants fra.CatalogLike[string, string],
	) MethodAssemblerLike
}

// INSTANCE DECLARATIONS

/*
MethodAssemblerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete method-assembler-like class.
*/
type MethodAssemblerLike interface {
	// Principal Methods
	GetClass() MethodAssemblerClassLike
	AssembleInstructions(
		assembly lan.AssemblyLike,
	)

	// Aspect Interfaces
	lan.Methodical
}

// ASPECT DECLARATIONS
