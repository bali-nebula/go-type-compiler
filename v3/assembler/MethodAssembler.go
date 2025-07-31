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
	var constants = c.extractKeys(type_, "$constants")
	var literals = c.extractItems(type_, "$literals")
	var instance = &methodAssembler_{
		// Initialize the instance attributes.
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
	// Extract the instance attributes from the method.
	v.extractAttributes(method)

	// Assemble the method into bytecode.
	v.assembleBytecode(method)

	// Add the assembled bytecode and address labels to the method.
	v.setBytecode(method)
	v.setAddresses(method)
}

// Attribute Methods

// Methodical Methods

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
		modifier = Modifier(count << 11)
	}
	var symbol = call.GetSymbol()
	var index = methodAssemblerClass().intrinsics_.GetIndex(symbol)
	var operand = Operand(index)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessDrop(
	drop lan.DropLike,
	index_ uint,
	count_ uint,
) {
	var operation = Drop
	var modifier Modifier
	switch drop.GetComponent().GetAny().(string) {
	case "CONTRACT":
		modifier = Contract
	case "DRAFT":
		modifier = Draft
	case "MESSAGE":
		modifier = Message
	case "VARIABLE":
		modifier = Variable
	}
	var symbol = drop.GetSymbol()
	var index = v.variables_.GetIndex(symbol)
	var operand = Operand(index)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessJump(
	jump lan.JumpLike,
	index_ uint,
	count_ uint,
) {
	var operation = Jump
	var modifier Modifier
	var conditionally = jump.GetOptionalConditionally()
	if uti.IsDefined(conditionally) {
		var condition = conditionally.GetCondition()
		switch condition.GetAny().(string) {
		case "EMPTY":
			modifier = Empty
		case "NONE":
			modifier = None
		case "FALSE":
			modifier = False
		}
	}
	var label = jump.GetLabel()
	var address = v.addresses_.GetValue(label)
	var operand = Operand(address)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessLoad(
	load lan.LoadLike,
	index_ uint,
	count_ uint,
) {
	var operation = Load
	var modifier Modifier
	switch load.GetComponent().GetAny().(string) {
	case "CONTRACT":
		modifier = Contract
	case "DRAFT":
		modifier = Draft
	case "MESSAGE":
		modifier = Message
	case "VARIABLE":
		modifier = Variable
	}
	var symbol = load.GetSymbol()
	var index = v.arguments_.GetIndex(symbol)
	if index == 0 {
		index = v.constants_.GetIndex(symbol)
	}
	if index == 0 {
		index = v.variables_.GetIndex(symbol)
	}
	var operand = Operand(index)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessPull(
	pull lan.PullLike,
	index_ uint,
	count_ uint,
) {
	var operation = Pull
	var modifier Modifier
	switch pull.GetValue().GetAny().(string) {
	case "HANDLER":
		modifier = Handler
	case "EXCEPTION":
		modifier = Exception
	case "COMPONENT":
		modifier = Component
	case "RESULT":
		modifier = Result
	}
	var operand Operand
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessPush(
	push lan.PushLike,
	index_ uint,
	count_ uint,
) {
	var operation = Push
	var modifier Modifier
	var operand Operand
	switch source := push.GetSource().GetAny().(type) {
	case lan.HandlerLike:
		modifier = Handler
		var address = v.addresses_.GetValue(source.GetLabel())
		operand = Operand(address)
	case lan.LiteralLike:
		modifier = Literal
		var literal = source.GetQuoted()
		literal = literal[3 : len(literal)-3] // Remove the delimiters.
		literal = not.FormatDocument(not.ParseSource(literal))
		literal = literal[:len(literal)-1] // Remove the trailing newline.
		var index = v.literals_.GetIndex(literal)
		operand = Operand(index)
	case lan.ConstantLike:
		modifier = Constant
		var index = v.constants_.GetIndex(source.GetSymbol())
		operand = Operand(index)
	case lan.ArgumentLike:
		modifier = Argument
		var index = v.arguments_.GetIndex(source.GetSymbol())
		operand = Operand(index)
	}
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessSave(
	save lan.SaveLike,
	index_ uint,
	count_ uint,
) {
	var operation = Save
	var modifier Modifier
	switch save.GetComponent().GetAny().(string) {
	case "CONTRACT":
		modifier = Contract
	case "DRAFT":
		modifier = Draft
	case "MESSAGE":
		modifier = Message
	case "VARIABLE":
		modifier = Variable
	}
	var symbol = save.GetSymbol()
	var index = v.variables_.GetIndex(symbol)
	var operand = Operand(index)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessSend(
	send lan.SendLike,
	index_ uint,
	count_ uint,
) {
	var operation = Send
	var modifier Modifier
	var withArguments = uti.IsDefined(send.GetOptionalParameterized())
	switch send.GetDestination().GetAny().(string) {
	case "CONTRACT":
		modifier = Contract
		if withArguments {
			modifier = ContractWithArguments
		}
	case "COMPONENT":
		modifier = Component
		if withArguments {
			modifier = ComponentWithArguments
		}
	}
	var symbol = send.GetSymbol()
	var index = v.messages_.GetIndex(symbol)
	var operand = Operand(index)
	var instruction = InstructionClass().Instruction(operation, modifier, operand)
	v.instructions_.AppendValue(instruction)
}

func (v *methodAssembler_) PreprocessSkip(
	skip lan.SkipLike,
	index_ uint,
	count_ uint,
) {
	var instruction = InstructionClass().Instruction(0, 0, 0)
	v.instructions_.AppendValue(instruction)
}

// PROTECTED INTERFACE

// Private Methods

func (c *methodAssemblerClass_) extractItems(
	document not.DocumentLike,
	symbol string,
) fra.SetLike[string] {
	var items = fra.Set[string]()
	var primitive = not.Primitive(not.Element(symbol))
	document = not.GetAttribute(document, primitive)
	switch component := document.GetComponent().GetAny().(type) {
	case not.CollectionLike:
		switch collection := component.GetAny().(type) {
		case not.ItemsLike:
			var entities = collection.GetEntities()
			var iterator = entities.GetIterator()
			for iterator.HasNext() {
				var entity = iterator.GetNext()
				var document = entity.GetDocument()
				var item = not.FormatDocument(document)
				items.AddValue(item[:len(item)-1]) // Remove the trailing newline.
			}
		}
	}
	return items
}

func (c *methodAssemblerClass_) extractKeys(
	document not.DocumentLike,
	symbol string,
) fra.SetLike[string] {
	var keys = fra.Set[string]()
	var primitive = not.Primitive(not.Element(symbol))
	document = not.GetAttribute(document, primitive)
	switch component := document.GetComponent().GetAny().(type) {
	case not.CollectionLike:
		switch collection := component.GetAny().(type) {
		case not.AttributesLike:
			var associations = collection.GetAssociations()
			var iterator = associations.GetIterator()
			for iterator.HasNext() {
				var association = iterator.GetNext()
				var element = association.GetPrimitive().GetAny().(not.ElementLike)
				var key = element.GetAny().(string)
				keys.AddValue(key)
			}
		}
	}
	return keys
}

func (v *methodAssembler_) assembleBytecode(
	method not.DocumentLike,
) {
	var assembly = v.extractAssembly(method)
	var instructions = assembly.GetInstructions()
	v.addresses_ = v.extractAddresses(instructions)
	lan.Visitor(v).VisitAssembly(assembly)
}

func (v *methodAssembler_) extractAddresses(
	instructions fra.Sequential[lan.InstructionLike],
) fra.CatalogLike[string, uint16] {
	var addresses = fra.Catalog[string, uint16]()
	var address uint16 = 1
	var iterator = instructions.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var prefix = instruction.GetOptionalPrefix()
		if uti.IsDefined(prefix) {
			var label = prefix.GetLabel()
			addresses.SetValue(label, address)
		}
		var action = instruction.GetAction()
		switch action.GetAny().(type) {
		case lan.NoteLike:
			// Notes don't generate a bytecode instruction.
		default:
			// All other actions do.
			address++
		}
	}
	return addresses
}

func (v *methodAssembler_) extractAssembly(
	method not.DocumentLike,
) lan.AssemblyLike {
	var primitive = not.Primitive(not.Element("$instructions"))
	var document = not.GetAttribute(method, primitive)
	var source = not.FormatDocument(document)
	source = source[3 : len(source)-3] // Remove the delimiters and their newlines.
	return lan.ParseSource(source)
}

func (v *methodAssembler_) extractAttributes(
	method not.DocumentLike,
) {
	var class = methodAssemblerClass()
	v.instructions_ = fra.List[InstructionLike]()
	v.arguments_ = class.extractItems(method, "$arguments")
	v.variables_ = class.extractItems(method, "$variables")
	v.messages_ = class.extractItems(method, "$messages")
}

func (v *methodAssembler_) setBytecode(
	method not.DocumentLike,
) {
	// Add the bytecode to the method.
	var source = BytecodeClass().Bytecode(v.instructions_).AsString()
	var bytecode = not.Document(not.Component(not.String(source)), nil, "")
	var primitive = not.Primitive(not.Element("$bytecode"))
	not.SetAttribute(method, bytecode, primitive)
}

func (v *methodAssembler_) setAddresses(
	method not.DocumentLike,
) {
	// Add the label addresses to the method.
	var list = fra.List[not.AssociationLike]()
	var iterator = v.addresses_.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var label = `"` + association.GetKey() + `"`
		var address = association.GetValue()
		var primitive = not.Primitive(not.String(label))
		var document = not.ParseSource(stc.Itoa(int(address)))
		list.AppendValue(not.Association(primitive, ":", document))
	}
	var attributes = not.Attributes("[", list, "]")
	var collection = not.Collection(attributes)
	var addresses = not.Document(not.Component(collection), nil, "")
	var primitive = not.Primitive(not.Element("$addresses"))
	not.SetAttribute(method, addresses, primitive)

}

// Instance Structure

type methodAssembler_ struct {
	// Declare the instance attributes.
	literals_     fra.SetLike[string]
	constants_    fra.SetLike[string]
	arguments_    fra.SetLike[string]
	variables_    fra.SetLike[string]
	messages_     fra.SetLike[string]
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
