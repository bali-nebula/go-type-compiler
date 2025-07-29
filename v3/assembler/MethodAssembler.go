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

package assembler

import (
	lan "github.com/bali-nebula/go-assembly-language/v3"
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	stc "strconv"
)

// CLASS INTERFACE

// Access Function

func MethodAssemblerClass() MethodAssemblerClassLike {
	return methodAssemblerClass()
}

// Constructor Methods

func (c *methodAssemblerClass_) MethodAssembler(
	type_ not.DocumentLike,
) MethodAssemblerLike {
	if uti.IsUndefined(type_) {
		panic("The \"type_\" attribute is required by this class.")
	}
	var constants = fra.Set[string]()
	var literals = fra.Set[string]()
	var instance = &methodAssembler_{
		// Initialize the instance attributes.
		type_:      type_,
		constants_: constants,
		literals_:  literals,

		// Initialize the inherited aspects.
		Methodical: lan.Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *methodAssembler_) GetClass() MethodAssemblerClassLike {
	return methodAssemblerClass()
}

func (v *methodAssembler_) AssembleMethod(
	method not.DocumentLike,
) {
	// Reset the instance attributes.
	v.address_ = 1
	v.arguments_ = fra.Set[string]()
	v.variables_ = fra.Set[string]()
	v.messages_ = fra.Set[string]()
	v.addresses_ = fra.Catalog[string, uint16]()
	v.instructions_ = fra.List[InstructionLike]()

	// Assemble the method into assembly instructions.
	var assembly lan.AssemblyLike
	lan.Visitor(v).VisitAssembly(assembly)

	// Convert the assembly instructions into bytecode.
	BytecodeClass().Bytecode(v.instructions_)
}

// Attribute Methods

func (v *methodAssembler_) GetType() not.DocumentLike {
	return v.type_
}

// Methodical Methods

func (v *methodAssembler_) ProcessLabel(
	label string,
) {
	v.addresses_.SetValue(label, v.address_)
}

func (v *methodAssembler_) PreprocessCall(
	call lan.CallLike,
	index_ uint,
	count_ uint,
) {
	var operation = Call
	var modifier Modifier
	var cardinality = call.GetOptionalCardinality()
	if uti.IsDefined(cardinality) {
		var count, _ = stc.Atoi(cardinality.GetCount())
		modifier = Modifier(count)
	}
	var symbol = call.GetSymbol()
	var index = methodAssemblerClass().intrinsics_.GetIndex(symbol)
	var operand = Operand(index)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PostprocessCall(
	call lan.CallLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessDrop(
	drop lan.DropLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessDrop(
	drop lan.DropLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessJump(
	jump lan.JumpLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessJump(
	jump lan.JumpLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessLoad(
	load lan.LoadLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessLoad(
	load lan.LoadLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessPull(
	pull lan.PullLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessPull(
	pull lan.PullLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessPush(
	push lan.PushLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessPush(
	push lan.PushLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessSave(
	save lan.SaveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessSave(
	save lan.SaveLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessSend(
	send lan.SendLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *methodAssembler_) PostprocessSend(
	send lan.SendLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

func (v *methodAssembler_) PreprocessSkip(
	skip lan.SkipLike,
	index_ uint,
	count_ uint,
) {
	var instruction = InstructionClass().Instruction(0, 0, 0)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PostprocessSkip(
	skip lan.SkipLike,
	index_ uint,
	count_ uint,
) {
	v.address_++
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type methodAssembler_ struct {
	// Declare the instance attributes.
	type_         not.DocumentLike
	literals_     fra.SetLike[string]
	constants_    fra.SetLike[string]
	arguments_    fra.SetLike[string]
	variables_    fra.SetLike[string]
	messages_     fra.SetLike[string]
	address_      uint16
	addresses_    fra.CatalogLike[string, uint16]
	instructions_ fra.ListLike[InstructionLike]

	// Declare the inherited aspects.
	lan.Methodical
}

// Class Structure

type methodAssemblerClass_ struct {
	// Declare the class constants.
	intrinsics_ fra.SetLike[string]
}

// Class Reference

func methodAssemblerClass() *methodAssemblerClass_ {
	return methodAssemblerClassReference_
}

var methodAssemblerClassReference_ = &methodAssemblerClass_{
	// Initialize the class constants.
	intrinsics_: fra.SetFromArray[string](
		[]string{
			"$invalid",
			"$addItem",
			"$ancestry",
			"$and",
			"$arccosine",
			"$arcsine",
			"$arctangent",
			"$areEqual",
			"$areSame",
			"$association",
			"$attribute",
			"$authority",
			"$base02",
			"$base16",
			"$base32",
			"$base64",
			"$binary",
			"$bytes",
			"$catalog",
			"$chain",
			"$citation",
			"$code",
			"$coinToss",
			"$comparator",
			"$complement",
			"$component",
			"$conjugate",
			"$connector",
			"$cosine",
			"$day",
			"$days",
			"$default",
			"$degrees",
			"$difference",
			"$document",
			"$doesMatch",
			"$duplicate",
			"$duration",
			"$earlier",
			"$effective",
			"$emptyCollection",
			"$exponential",
			"$factorial",
			"$first",
			"$format",
			"$fragment",
			"$hasNext",
			"$hasPrevious",
			"$hash",
			"$head",
			"$hour",
			"$hours",
			"$html",
			"$imaginary",
			"$insertItem",
			"$insertItems",
			"$integer",
			"$interfaces",
			"$inverse",
			"$ior",
			"$isEnumerable",
			"$isLess",
			"$isMore",
			"$isNegative",
			"$isSignificant",
			"$item",
			"$iterator",
			"$key",
			"$keys",
			"$last",
			"$later",
			"$levels",
			"$list",
			"$logarithm",
			"$magnitude",
			"$matchesText",
			"$millisecond",
			"$milliseconds",
			"$minute",
			"$minutes",
			"$month",
			"$months",
			"$nextItem",
			"$nextVersion",
			"$node",
			"$not",
			"$now",
			"$parameters",
			"$path",
			"$phase",
			"$previousItem",
			"$procedure",
			"$product",
			"$query",
			"$queue",
			"$quotient",
			"$radians",
			"$random",
			"$range",
			"$ranking",
			"$real",
			"$reciprocal",
			"$remainder",
			"$removeAttribute",
			"$removeHead",
			"$removeIndex",
			"$removeIndices",
			"$removeItem",
			"$removeTop",
			"$reverseItems",
			"$sans",
			"$scaled",
			"$scheme",
			"$second",
			"$seconds",
			"$set",
			"$setAttribute",
			"$setFirst",
			"$setItem",
			"$setLast",
			"$setParameter",
			"$setValue",
			"$shuffleItems",
			"$sine",
			"$size",
			"$sortItems",
			"$sorter",
			"$source",
			"$stack",
			"$sum",
			"$supplement",
			"$tag",
			"$tangent",
			"$toEnd",
			"$toSlot",
			"$toStart",
			"$top",
			"$value",
			"$weeks",
			"$xor",
			"$year",
			"$years",
		},
	),
}
