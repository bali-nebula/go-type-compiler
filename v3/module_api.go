/*
................................................................................
.    Copyright (c) 2009-2026 Crater Dog Technologies™.  All Rights Reserved.   .
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
	ins "github.com/bali-nebula/go-bali-instructions/v3"
	ass "github.com/bali-nebula/go-type-compiler/v3/assembler"
	com "github.com/craterdog/go-essential-composites/v8"
)

// TYPE ALIASES

// Assembler

type (
	AnalyzerClassLike  = ass.AnalyzerClassLike
	AssemblerClassLike = ass.AssemblerClassLike
)

type (
	AnalyzerLike  = ass.AnalyzerLike
	AssemblerLike = ass.AssemblerLike
)

// CLASS ACCESSORS

// Assembler

func AnalyzerClass() AnalyzerClassLike {
	return ass.AnalyzerClass()
}

func Analyzer(
	literals com.Accessible[string],
	constants com.Accessible[string],
	method ins.MethodLike,
) AnalyzerLike {
	return AnalyzerClass().Analyzer(
		literals,
		constants,
		method,
	)
}

func AssemblerClass() AssemblerClassLike {
	return ass.AssemblerClass()
}

func Assembler() AssemblerLike {
	return AssemblerClass().Assembler()
}

// GLOBAL FUNCTIONS
