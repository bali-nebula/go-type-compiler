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
	fmt "fmt"
	ins "github.com/bali-nebula/go-bali-instructions/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
)

// CLASS INTERFACE

// Access Function

func AnalyzerClass() AnalyzerClassLike {
	return analyzerClass()
}

// Constructor Methods

func (c *analyzerClass_) Analyzer(
	literals com.Accessible[string],
	constants com.Accessible[string],
	method ins.AssemblyLike,
) AnalyzerLike {
	if uti.IsUndefined(literals) {
		panic("The \"literals\" attribute is required by this class.")
	}
	if uti.IsUndefined(constants) {
		panic("The \"constants\" attribute is required by this class.")
	}
	var instance = &analyzer_{
		// Initialize the instance attributes.
		literals_:  literals,
		constants_: constants,
		arguments_: com.Set[string](),
		variables_: com.Set[string](),
		messages_:  com.Set[string](),
		addresses_: com.Catalog[string, uint16](),

		// Initialize the inherited aspects.
		Methodical: ins.Processor(),
	}
	ins.VisitorClass().Visitor(instance).VisitAssembly(method)
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *analyzer_) GetClass() AnalyzerClassLike {
	return analyzerClass()
}

func (v *analyzer_) LookupLiteral(
	literal string,
) uint16 {
	return uint16(v.literals_.GetIndex(literal))
}

func (v *analyzer_) LookupConstant(
	constant string,
) uint16 {
	return uint16(v.constants_.GetIndex(constant))
}

func (v *analyzer_) LookupArgument(
	argument string,
) uint16 {
	return uint16(v.arguments_.GetIndex(argument))
}

func (v *analyzer_) LookupVariable(
	variable string,
) uint16 {
	return uint16(v.variables_.GetIndex(variable))
}

func (v *analyzer_) LookupIntrinsic(
	intrinsic string,
) uint16 {
	return uint16(analyzerClass().intrinsics_.GetIndex(intrinsic))
}

func (v *analyzer_) LookupMessage(
	message string,
) uint16 {
	return uint16(v.messages_.GetIndex(message))
}

func (v *analyzer_) LookupAddress(
	label string,
) uint16 {
	return v.addresses_.GetValue(label)
}

// Attribute Methods

// ins.Methodical Methods

func (v *analyzer_) PreprocessArgument(
	argument ins.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	var symbol = argument.GetSymbol()
	v.arguments_.AddValue(symbol)
}

func (v *analyzer_) PreprocessAssembly(
	assembly ins.AssemblyLike,
	index_ uint,
	count_ uint,
) {
	var address uint16
	var instructions = assembly.GetInstructions().GetIterator()
	for instructions.HasNext() {
		address++
		var instruction = instructions.GetNext()
		var prefix = instruction.GetOptionalPrefix()
		if uti.IsDefined(prefix) {
			var label = prefix.GetLabel()
			if v.addresses_.GetValue(label) > 0 {
				var message = fmt.Sprintf(
					"Found a duplicate prefix label: %s",
					label,
				)
				panic(message)
			}
			v.addresses_.SetValue(label, address)
		}
		switch instruction.GetAction().(type) {
		case ins.NoteLike:
			address--
		}
	}
}

func (v *analyzer_) PreprocessCall(
	call ins.CallLike,
	index_ uint,
	count_ uint,
) {
	var symbol = call.GetSymbol()
	if uint16(analyzerClass().intrinsics_.GetIndex(symbol)) == 0 {
		var message = fmt.Sprintf(
			"Found an unknown intrinsic function: %s",
			symbol,
		)
		panic(message)
	}
}

func (v *analyzer_) PreprocessConstant(
	constant ins.ConstantLike,
	index_ uint,
	count_ uint,
) {
	var symbol = constant.GetSymbol()
	if uint16(v.constants_.GetIndex(symbol)) == 0 {
		var message = fmt.Sprintf(
			"Found an unknown predefined constant: %s",
			symbol,
		)
		panic(message)
	}
}

func (v *analyzer_) PreprocessDrop(
	drop ins.DropLike,
	index_ uint,
	count_ uint,
) {
	var symbol = drop.GetSymbol()
	v.variables_.AddValue(symbol)
}

func (v *analyzer_) PreprocessHandler(
	handler ins.HandlerLike,
	index_ uint,
	count_ uint,
) {
	var label = handler.GetLabel()
	if uint16(v.addresses_.GetValue(label)) == 0 {
		var message = fmt.Sprintf(
			"Found an unknown handler label: %s",
			label,
		)
		panic(message)
	}
}

func (v *analyzer_) PreprocessJump(
	jump ins.JumpLike,
	index_ uint,
	count_ uint,
) {
	var label = jump.GetLabel()
	if uint16(v.addresses_.GetValue(label)) == 0 {
		var message = fmt.Sprintf(
			"Found an unknown jump label: %s",
			label,
		)
		panic(message)
	}
}

func (v *analyzer_) PreprocessLiteral(
	literal ins.LiteralLike,
	index_ uint,
	count_ uint,
) {
	var quoted = literal.GetQuoted()
	if uint16(v.literals_.GetIndex(quoted)) == 0 {
		var message = fmt.Sprintf(
			"Found an unknown quoted literal: %s",
			quoted,
		)
		panic(message)
	}
}

func (v *analyzer_) PreprocessLoad(
	load ins.LoadLike,
	index_ uint,
	count_ uint,
) {
	var symbol = load.GetSymbol()
	v.variables_.AddValue(symbol)
}

func (v *analyzer_) PreprocessSave(
	save ins.SaveLike,
	index_ uint,
	count_ uint,
) {
	var symbol = save.GetSymbol()
	v.variables_.AddValue(symbol)
}

func (v *analyzer_) PreprocessSend(
	send ins.SendLike,
	index_ uint,
	count_ uint,
) {
	var symbol = send.GetSymbol()
	v.messages_.AddValue(symbol)
}

// PROTECTED INTERFACE

// Private Methods

// Instance Structure

type analyzer_ struct {
	// Declare the instance attributes.
	literals_  com.Accessible[string]
	constants_ com.Accessible[string]
	arguments_ com.SetLike[string]
	variables_ com.SetLike[string]
	messages_  com.SetLike[string]
	addresses_ com.CatalogLike[string, uint16]

	// Declare the inherited aspects.
	ins.Methodical
}

// Class Structure

type analyzerClass_ struct {
	// Declare the class constants.
	intrinsics_ com.Accessible[string]
}

// Class Reference

func analyzerClass() *analyzerClass_ {
	return analyzerClassReference_
}

var analyzerClassReference_ = &analyzerClass_{
	// Initialize the class constants.
	intrinsics_: com.SetFromArray[string](
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
