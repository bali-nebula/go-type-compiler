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

package assembler

import (
	fmt "fmt"
	ins "github.com/bali-nebula/go-bali-instructions/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func AssemblerClass() AssemblerClassLike {
	return assemblerClass()
}

// Constructor Methods

func (c *assemblerClass_) Assembler() AssemblerLike {
	var instance = &assembler_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ins.Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

func (c *assemblerClass_) FormatInstructions(
	instructions com.Sequential[uint16],
) string {
	var result sts.Builder
	var header = `
┌─────────┬───────┬─────────────┬───────────────────────────────────────┐
│ Address │  Hex  │ Instruction │               Mnemonic                │
├─────────┼───────┼─────────────┼───────────────────────────────────────┤
`
	result.WriteString(header[1:])
	var iterator = instructions.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var address = fmt.Sprintf("[x%03x]", iterator.GetSlot())
		var bytes = fmt.Sprintf("x%04x", instruction)
		var operation = (c.operation_ & instruction) >> 13
		var modifier = (c.modifier_ & instruction) >> 11
		var operand = c.operand_ & instruction
		var operandAsString = fmt.Sprintf(" %4d ", operand)
		if operation == 0 || (operation == 1 && modifier == 3) {
			operandAsString = fmt.Sprintf("[x%03x]", operand)
		}
		var bytecode = fmt.Sprintf("%d %d  %s", operation, modifier, operandAsString)
		var description = c.generateDescription(operation, modifier, operandAsString)
		var line = fmt.Sprintf(
			"│ %s: │ %s │ %s │ %s │\n",
			address,
			bytes,
			bytecode,
			description,
		)
		result.WriteString(line)
	}
	var footer = `
└─────────┴───────┴─────────────┴───────────────────────────────────────┘
`
	result.WriteString(footer[1:])
	return result.String()
}

// INSTANCE INTERFACE

// Principal Methods

func (v *assembler_) GetClass() AssemblerClassLike {
	return assemblerClass()
}

func (v *assembler_) AssembleMethod(
	literals com.Accessible[string],
	constants com.Accessible[string],
	method ins.MethodLike,
) com.Sequential[uint16] {
	if uti.IsUndefined(literals) {
		panic("The \"literals\" attribute is required by this class.")
	}
	if uti.IsUndefined(constants) {
		panic("The \"constants\" attribute is required by this class.")
	}
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	v.analyzer_ = AnalyzerClass().Analyzer(literals, constants, method)
	v.instructions_ = com.List[uint16]()
	ins.VisitorClass().Visitor(v).VisitMethod(method)
	return v.instructions_
}

// Attribute Methods

// ins.Methodical Methods

func (v *assembler_) PostprocessArgument(
	argument ins.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	var symbol = argument.GetSymbol()
	var modifier = assemblerClass().argument_
	var operand = uint16(v.analyzer_.LookupArgument(symbol))
	var instruction = assemblerClass().push_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessCall(
	call ins.CallLike,
	index_ uint,
	count_ uint,
) {
	var symbol = call.GetSymbol()
	var argumentCount = call.GetArgumentCount()
	var modifier = uint16(argumentCount) << 11
	var operand = uint16(v.analyzer_.LookupIntrinsic(symbol))
	if operand == 0 {
		var message = fmt.Sprintf(
			"Found an unknown intrinsic function: %s",
			symbol,
		)
		panic(message)
	}
	var instruction = assemblerClass().call_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessConstant(
	constant ins.ConstantLike,
	index_ uint,
	count_ uint,
) {
	var symbol = constant.GetSymbol()
	var modifier = assemblerClass().constant_
	var operand = uint16(v.analyzer_.LookupConstant(symbol))
	if operand == 0 {
		var message = fmt.Sprintf(
			"Found an unknown defined constant: %s",
			symbol,
		)
		panic(message)
	}
	var instruction = assemblerClass().push_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessDrop(
	drop ins.DropLike,
	index_ uint,
	count_ uint,
) {
	var symbol = drop.GetSymbol()
	var component = drop.GetComponent()
	var modifier = uint16(component) << 11
	var operand = uint16(v.analyzer_.LookupVariable(symbol))
	var instruction = assemblerClass().drop_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessHandler(
	handler ins.HandlerLike,
	index_ uint,
	count_ uint,
) {
	var label = handler.GetLabel()
	var modifier = assemblerClass().handler_
	var operand = uint16(v.analyzer_.LookupAddress(label))
	if operand == 0 {
		var message = fmt.Sprintf(
			"Found an unknown handler label: %s",
			label,
		)
		panic(message)
	}
	var instruction = assemblerClass().push_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessJump(
	jump ins.JumpLike,
	index_ uint,
	count_ uint,
) {
	var label = jump.GetLabel()
	var condition = jump.GetCondition()
	var modifier = uint16(condition) << 11
	var operand = uint16(v.analyzer_.LookupAddress(label))
	if operand == 0 {
		var message = fmt.Sprintf(
			"Found an unknown jump label: %s",
			label,
		)
		panic(message)
	}
	var instruction = assemblerClass().jump_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessLiteral(
	literal ins.LiteralLike,
	index_ uint,
	count_ uint,
) {
	var quoted = literal.GetQuoted()
	var modifier = assemblerClass().literal_
	var operand = uint16(v.analyzer_.LookupLiteral(quoted))
	if operand == 0 {
		var message = fmt.Sprintf(
			"Found an unknown quoted literal: %s",
			quoted,
		)
		panic(message)
	}
	var instruction = assemblerClass().push_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessLoad(
	load ins.LoadLike,
	index_ uint,
	count_ uint,
) {
	var symbol = load.GetSymbol()
	var component = load.GetComponent()
	var modifier = uint16(component) << 11
	var operand = uint16(v.analyzer_.LookupVariable(symbol))
	var instruction = assemblerClass().load_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessPull(
	pull ins.PullLike,
	index_ uint,
	count_ uint,
) {
	var value = pull.GetValue()
	var modifier = uint16(value) << 11
	var instruction = assemblerClass().pull_ ^ modifier
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessSave(
	save ins.SaveLike,
	index_ uint,
	count_ uint,
) {
	var symbol = save.GetSymbol()
	var component = save.GetComponent()
	var modifier = uint16(component) << 11
	var operand = uint16(v.analyzer_.LookupVariable(symbol))
	var instruction = assemblerClass().save_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessSend(
	send ins.SendLike,
	index_ uint,
	count_ uint,
) {
	var symbol = send.GetSymbol()
	var destination = send.GetDestination()
	var modifier = uint16(destination) << 11
	var operand = uint16(v.analyzer_.LookupMessage(symbol))
	var instruction = assemblerClass().send_ ^ modifier ^ operand
	v.instructions_.AppendValue(instruction)
}

func (v *assembler_) PostprocessSkip(
	skip ins.SkipLike,
	index_ uint,
	count_ uint,
) {
	var instruction = assemblerClass().jump_
	v.instructions_.AppendValue(instruction)
}

// PROTECTED INTERFACE

// Private Methods

func (c *assemblerClass_) generateDescription(
	operation uint16,
	modifier uint16,
	operandAsString string,
) string {
	var description string
	var operand = sts.ReplaceAll(operandAsString, " ", "")
	switch operation {
	case 0:
		description = "SKIP"
		if operand != "[x000]" {
			description = "JUMP TO " + operand
			switch modifier {
			case 1:
				description += " ON EMPTY"
			case 2:
				description += " ON NONE"
			case 3:
				description += " ON FALSE"
			}
		}
	case 1:
		description = "PUSH "
		switch modifier {
		case 0:
			description += "LITERAL "
		case 1:
			description += "CONSTANT "
		case 2:
			description += "ARGUMENT "
		case 3:
			description += "HANDLER "
		}
		description += operand
	case 2:
		description = "PULL "
		switch modifier {
		case 0:
			description += "COMPONENT"
		case 1:
			description += "RESULT"
		case 2:
			description += "EXCEPTION"
		case 3:
			description += "HANDLER"
		}
	case 3:
		description = "LOAD "
		switch modifier {
		case 0:
			description += "VARIABLE "
		case 1:
			description += "DRAFT "
		case 2:
			description += "DOCUMENT "
		case 3:
			description += "MESSAGE "
		}
		description += operand
	case 4:
		description = "SAVE "
		switch modifier {
		case 0:
			description += "VARIABLE "
		case 1:
			description += "DRAFT "
		case 2:
			description += "DOCUMENT "
		case 3:
			description += "MESSAGE "
		}
		description += operand
	case 5:
		description = "DROP "
		switch modifier {
		case 0:
			description += "VARIABLE "
		case 1:
			description += "DRAFT "
		case 2:
			description += "DOCUMENT "
		case 3:
			description += "MESSAGE "
		}
		description += operand
	case 6:
		description = "CALL " + operand
		switch modifier {
		case 1:
			description += " WITH 1 ARGUMENT"
		case 2:
			description += " WITH 2 ARGUMENTS"
		case 3:
			description += " WITH 3 ARGUMENTS"
		}
	case 7:
		description = "SEND " + operand
		switch modifier {
		case 0:
			description += " TO COMPONENT"
		case 1:
			description += " TO COMPONENT WITH ARGUMENTS"
		case 2:
			description += " TO DOCUMENT"
		case 3:
			description += " TO DOCUMENT WITH ARGUMENTS"
		}
	}
	for len(description) < 37 {
		description += " "
	}
	return description
}

// Instance Structure

type assembler_ struct {
	// Declare the instance attributes.
	analyzer_     AnalyzerLike
	instructions_ com.ListLike[uint16]

	// Declare the inherited aspects.
	ins.Methodical
}

// Class Structure

type assemblerClass_ struct {
	// Declare the class constants.
	operation_ uint16
	jump_      uint16
	push_      uint16
	pull_      uint16
	load_      uint16
	save_      uint16
	drop_      uint16
	call_      uint16
	send_      uint16
	modifier_  uint16
	literal_   uint16
	constant_  uint16
	argument_  uint16
	handler_   uint16
	operand_   uint16
}

// Class Reference

func assemblerClass() *assemblerClass_ {
	return assemblerClassReference_
}

var assemblerClassReference_ = &assemblerClass_{
	// Initialize the class constants.
	operation_: 0b1110000000000000,
	jump_:      0b0000000000000000,
	push_:      0b0010000000000000,
	pull_:      0b0100000000000000,
	load_:      0b0110000000000000,
	save_:      0b1000000000000000,
	drop_:      0b1010000000000000,
	call_:      0b1100000000000000,
	send_:      0b1110000000000000,
	modifier_:  0b0001100000000000,
	literal_:   0b0000000000000000,
	constant_:  0b0000100000000000,
	argument_:  0b0001000000000000,
	handler_:   0b0001100000000000,
	operand_:   0b0000011111111111,
}
