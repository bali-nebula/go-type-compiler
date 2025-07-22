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
)

// TYPE ALIASES

// Assembler

type (
	MethodAssemblerClassLike = ass.MethodAssemblerClassLike
)

type (
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

func MethodAssemblerClass() MethodAssemblerClassLike {
	return ass.MethodAssemblerClass()
}

func MethodAssembler(
	method not.DocumentLike,
) MethodAssemblerLike {
	return MethodAssemblerClass().MethodAssembler(method)
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
