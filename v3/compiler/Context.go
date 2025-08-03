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

package compiler

import (
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func ContextClass() ContextClassLike {
	return contextClass()
}

// Constructor Methods

func (c *contextClass_) Context(
	labelPrefix string,
	labels fra.CatalogLike[string, string],
	statement not.StatementLike,
	statementNumber uint,
	statementCount uint,
	blockNumber uint,
	blockCount uint,
) ContextLike {
	if uti.IsUndefined(labelPrefix) {
		panic("The \"labelPrefix\" attribute is required by this class.")
	}
	if uti.IsUndefined(labels) {
		panic("The \"labels\" attribute is required by this class.")
	}
	if uti.IsUndefined(statement) {
		panic("The \"statement\" attribute is required by this class.")
	}
	if uti.IsUndefined(statementNumber) {
		panic("The \"statementNumber\" attribute is required by this class.")
	}
	if uti.IsUndefined(statementCount) {
		panic("The \"statementCount\" attribute is required by this class.")
	}
	if uti.IsUndefined(blockNumber) {
		panic("The \"blockNumber\" attribute is required by this class.")
	}
	if uti.IsUndefined(blockCount) {
		panic("The \"blockCount\" attribute is required by this class.")
	}
	var instance = &context_{
		// Initialize the instance attributes.
		labelPrefix_: labelPrefix,
		labels_: labels,
		statement_: statement,
		statementNumber_: statementNumber,
		statementCount_: statementCount,
		blockNumber_: blockNumber,
		blockCount_: blockCount,
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *context_) GetClass() ContextClassLike {
	return contextClass()
}

// Attribute Methods

func (v *context_) GetLabelPrefix() string {
	return v.labelPrefix_
}

func (v *context_) GetLabels() fra.CatalogLike[string, string] {
	return v.labels_
}

func (v *context_) GetStatement() not.StatementLike {
	return v.statement_
}

func (v *context_) GetStatementNumber() uint {
	return v.statementNumber_
}

func (v *context_) GetStatementCount() uint {
	return v.statementCount_
}

func (v *context_) GetBlockNumber() uint {
	return v.blockNumber_
}

func (v *context_) GetBlockCount() uint {
	return v.blockCount_
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type context_ struct {
	// Declare the instance attributes.
	labelPrefix_ string
	labels_ fra.CatalogLike[string, string]
	statement_ not.StatementLike
	statementNumber_ uint
	statementCount_ uint
	blockNumber_ uint
	blockCount_ uint
}

// Class Structure

type contextClass_ struct {
	// Declare the class constants.
}

// Class Reference

func contextClass() *contextClass_ {
	return contextClassReference_
}

var contextClassReference_ = &contextClass_{
	// Initialize the class constants.
}
