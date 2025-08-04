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
	statementCount uint,
	currentStatement not.StatementLike,
) ContextLike {
	if uti.IsUndefined(labelPrefix) {
		panic("The \"labelPrefix\" attribute is required by this class.")
	}
	if uti.IsUndefined(statementCount) {
		panic("The \"statementCount\" attribute is required by this class.")
	}
	if uti.IsUndefined(currentStatement) {
		panic("The \"currentStatement\" attribute is required by this class.")
	}
	var instance = &context_{
		// Initialize the instance attributes.
		labelPrefix_:      labelPrefix,
		labels_:           fra.Catalog[string, string](),
		currentStatement_: currentStatement,
		statementNumber_:  1,
		statementCount_:   statementCount,
		blockNumber_:      1,
		blockCount_:       c.countBlocks(currentStatement),
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

func (v *context_) SetLabel(
	name string,
	label string,
) {
	v.labels_.SetValue(name, label)
}

func (v *context_) GetLabel(
	name string,
) string {
	return v.labels_.GetValue(name)
}

func (v *context_) IncrementStatementNumber() {
	v.statementNumber_++
}

func (v *context_) IncrementBlockNumber() {
	v.blockNumber_++
}

// Attribute Methods

func (v *context_) GetLabelPrefix() string {
	return v.labelPrefix_
}

func (v *context_) GetCurrentStatement() not.StatementLike {
	return v.currentStatement_
}

func (v *context_) SetCurrentStatement(
	currentStatement not.StatementLike,
) {
	v.currentStatement_ = currentStatement
	v.blockNumber_ = 1
	v.blockCount_ = contextClass().countBlocks(currentStatement)
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

func (c *contextClass_) countBlocks(
	statement not.StatementLike,
) uint {
	return 0
}

// Instance Structure

type context_ struct {
	// Declare the instance attributes.
	labelPrefix_      string
	labels_           fra.CatalogLike[string, string]
	currentStatement_ not.StatementLike
	statementNumber_  uint
	statementCount_   uint
	blockNumber_      uint
	blockCount_       uint
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
