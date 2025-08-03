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
Package "compiler" provides an implementation of a compiler that can be used to
compile types defined using Bali Document Notation™.

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
package compiler

import (
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
ContextClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete context-like class.
*/
type ContextClassLike interface {
	// Constructor Methods
	Context(
		labelPrefix string,
		labels fra.CatalogLike[string, string],
		statement not.StatementLike,
		statementNumber uint,
		statementCount uint,
		blockNumber uint,
		blockCount uint,
	) ContextLike
}

/*
TypeCompilerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete type-compiler-like class.
*/
type TypeCompilerClassLike interface {
	// Constructor Methods
	TypeCompiler() TypeCompilerLike
}

// INSTANCE DECLARATIONS

/*
ContextLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete type-compiler-like class.
*/
type ContextLike interface {
	// Principal Methods
	GetClass() ContextClassLike

	// Attribute Methods
	GetLabelPrefix() string
	GetLabels() fra.CatalogLike[string, string]
	GetStatement() not.StatementLike
	GetStatementNumber() uint
	GetStatementCount() uint
	GetBlockNumber() uint
	GetBlockCount() uint
}

/*
TypeCompilerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each instance
of a concrete type-compiler-like class.
*/
type TypeCompilerLike interface {
	// Principal Methods
	GetClass() TypeCompilerClassLike
	CompileType(
		type_ not.DocumentLike,
	)

	// Aspect Interfaces
	not.Methodical
}

// ASPECT DECLARATIONS
