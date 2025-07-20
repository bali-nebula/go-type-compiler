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
to assemble compiled type methods into bytecode.

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
	not "github.com/bali-nebula/go-document-notation/v3"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
TypeAssemblerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete type-assembler-like class.
*/
type TypeAssemblerClassLike interface {
	// Constructor Methods
	TypeAssembler() TypeAssemblerLike
}

// INSTANCE DECLARATIONS

/*
TypeAssemblerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete type-assembler-like class.
*/
type TypeAssemblerLike interface {
	// Principal Methods
	GetClass() TypeAssemblerClassLike
	AssembleType(
		document not.DocumentLike,
	)

	// Aspect Interfaces
	not.Methodical
}

// ASPECT DECLARATIONS
