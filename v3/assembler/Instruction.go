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
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func InstructionClass() InstructionClassLike {
	return instructionClass()
}

// Constructor Methods

func (c *instructionClass_) Instruction(
	operation Operation,
	modifier Modifier,
	operand Operand,
) InstructionLike {
	c.validateInstruction(operation, modifier, operand)
	var instruction = uint16(operation) ^ uint16(modifier) ^ uint16(operand)
	return instruction_(instruction)
}

func (c *instructionClass_) InstructionFromInteger(
	integer uint16,
) InstructionLike {
	var operation = Operation(integer & c.operationMask_)
	var modifier = Modifier(integer & c.modifierMask_)
	var operand = Operand(integer & c.operandMask_)
	return c.Instruction(operation, modifier, operand)
}

// Constant Methods

func (c *instructionClass_) OperationMask() uint16 {
	return c.operationMask_
}

func (c *instructionClass_) ModifierMask() uint16 {
	return c.modifierMask_
}

func (c *instructionClass_) OperandMask() uint16 {
	return c.operandMask_
}

// Function Methods

func (c *instructionClass_) FormatInstructions(
	instructions fra.Sequential[InstructionLike],
) string {
	var result sts.Builder
	result.WriteString(
		`Address   Bytes   Instruction                Mnemonic
--------------------------------------------------------------------
`,
	)
	var iterator = instructions.GetIterator()
	for iterator.HasNext() {
		var instruction = iterator.GetNext()
		var address = fmt.Sprintf("[x%03x]", iterator.GetSlot())
		var bytes = fmt.Sprintf("x%04x", instruction.AsIntrinsic())
		var operation = instruction.GetOperation() >> 13
		var modifier = instruction.GetModifier() >> 11
		var operand = instruction.OperandAsString()
		for len(operand) < 5 {
			operand = " " + operand
		}
		if len(operand) < 6 {
			operand += " "
		}
		var bytecode = fmt.Sprintf("%d %d %s", operation, modifier, operand)
		var description = instruction.AsString()
		var line = fmt.Sprintf(
			"%s:   %s   %s   %s\n",
			address,
			bytes,
			bytecode,
			description,
		)
		result.WriteString(line)
	}
	return result.String()
}

// INSTANCE INTERFACE

// Principal Methods

func (v instruction_) GetClass() InstructionClassLike {
	return instructionClass()
}

func (v instruction_) AsIntrinsic() uint16 {
	return uint16(v)
}

func (v instruction_) AsString() string {
	var instruction = v.OperationAsString()
	switch instruction {
	case "SKIP":
	case "JUMP":
		instruction += " TO " + v.OperandAsString()
		var modifier = v.ModifierAsString()
		if uti.IsDefined(modifier) {
			instruction += " " + modifier
		}
	case "CALL", "SEND":
		instruction += " " + v.OperandAsString()
		var modifier = v.ModifierAsString()
		if uti.IsDefined(modifier) {
			instruction += " " + modifier
		}
	case "PULL":
		instruction += " " + v.ModifierAsString()
	case "PUSH", "LOAD", "SAVE", "DROP":
		instruction += " " + v.ModifierAsString()
		instruction += " " + v.OperandAsString()
	}
	return instruction
}

func (v instruction_) OperationAsString() string {
	var operation = v.GetOperation()
	switch operation {
	case Jump:
		if v.GetOperand() == 0 {
			return "SKIP"
		}
		return "JUMP"
	case Push:
		return "PUSH"
	case Pull:
		return "PULL"
	case Load:
		return "LOAD"
	case Save:
		return "SAVE"
	case Drop:
		return "DROP"
	case Call:
		return "CALL"
	case Send:
		return "SEND"
	default:
		var message = fmt.Sprintf(
			"Found an unknown operation type: %v",
			operation,
		)
		panic(message)
	}
}

func (v instruction_) ModifierAsString() string {
	var operation = v.GetOperation()
	var modifier = v.GetModifier()
	switch operation {
	case Jump:
		switch modifier {
		case Any:
			return ""
		case Empty:
			return "ON EMPTY"
		case None:
			return "ON NONE"
		case False:
			return "ON FALSE"
		}
	case Push:
		switch modifier {
		case Handler:
			return "HANDLER"
		case Literal:
			return "LITERAL"
		case Constant:
			return "CONSTANT"
		case Argument:
			return "ARGUMENT"
		}
	case Pull:
		switch modifier {
		case Handler:
			return "HANDLER"
		case Exception:
			return "EXCEPTION"
		case Component:
			return "COMPONENT"
		case Result:
			return "RESULT"
		}
	case Load, Save, Drop:
		switch modifier {
		case Contract:
			return "CONTRACT"
		case Draft:
			return "DRAFT"
		case Message:
			return "MESSAGE"
		case Variable:
			return "VARIABLE"
		}
	case Call:
		switch modifier {
		case With0Arguments:
			return ""
		case With1Argument:
			return "WITH 1 ARGUMENTS"
		case With2Arguments:
			return "WITH 2 ARGUMENTS"
		case With3Arguments:
			return "WITH 3 ARGUMENTS"
		}
	case Send:
		switch modifier {
		case Component:
			return "TO COMPONENT"
		case ComponentWithArguments:
			return "TO COMPONENT WITH ARGUMENTS"
		case Contract:
			return "TO CONTRACT"
		case ContractWithArguments:
			return "TO CONTRACT WITH ARGUMENTS"
		}
	default:
		var message = fmt.Sprintf(
			"Found an unknown operation type: %v",
			operation,
		)
		panic(message)
	}
	var message = fmt.Sprintf(
		"Found an unknown modifier type: %v",
		modifier,
	)
	panic(message)
}

func (v instruction_) OperandAsString() string {
	var operation = v.GetOperation()
	var modifier = v.GetModifier()
	var operand = v.GetOperand()
	if operation == Jump || (operation == Push && modifier == Handler) {
		// Treat the operand as an address "[xHEX]".
		return fmt.Sprintf("[x%03x]", operand)
	} else {
		// Treat the operand as an index "DECI".
		return fmt.Sprintf("%d", operand)
	}
}

// Attribute Methods

func (v instruction_) GetOperation() Operation {
	return Operation(v.AsIntrinsic() & instructionClass().operationMask_)
}

func (v instruction_) GetModifier() Modifier {
	return Modifier(v.AsIntrinsic() & instructionClass().modifierMask_)
}

func (v instruction_) GetOperand() Operand {
	return Operand(v.AsIntrinsic() & instructionClass().operandMask_)
}

// PROTECTED INTERFACE

// Private Methods

func (c instructionClass_) validateInstruction(
	operation Operation,
	modifier Modifier,
	operand Operand,
) {
	if (operation & Operation(c.operationMask_)) != operation {
		var message = fmt.Sprintf(
			"An invalid operation was passed to the instruction constructor: %016b",
			operation,
		)
		panic(message)
	}
	if (modifier & Modifier(c.modifierMask_)) != modifier {
		var message = fmt.Sprintf(
			"An invalid modifier was passed to the instruction constructor: %016b",
			modifier,
		)
		panic(message)
	}
	if (operand & Operand(c.operandMask_)) != operand {
		var message = fmt.Sprintf(
			"An invalid operand was passed to the instruction constructor: %016b",
			operand,
		)
		panic(message)
	}
	switch operation {
	case Jump:
		if modifier > 0 && operand == 0 {
			var message = "A JUMP operation with a modifier requires a non-zero offset."
			panic(message)
		}
	case Push:
		if operand == 0 {
			var message = "A PUSH operation requires an operand."
			panic(message)
		}
	case Pull:
		if operand > 0 {
			var message = "A PULL operation does not take an operand."
			panic(message)
		}
	case Load:
		if operand == 0 {
			var message = "A LOAD operation requires an operand."
			panic(message)
		}
	case Save:
		if operand == 0 {
			var message = "A SAVE operation requires an operand."
			panic(message)
		}
	case Drop:
		if operand == 0 {
			var message = "A DROP operation requires an operand."
			panic(message)
		}
	case Call:
		if operand == 0 {
			var message = "A CALL operation requires an operand."
			panic(message)
		}
	case Send:
		if operand == 0 {
			var message = "A SEND operation requires an operand."
			panic(message)
		}
	default:
		var message = fmt.Sprintf(
			"An invalid operand was passed to the instruction constructor: %016b",
			operand,
		)
		panic(message)
	}
}

// Instance Structure

type instruction_ uint16

// Class Structure

type instructionClass_ struct {
	// Declare the class constants.
	operationMask_ uint16
	modifierMask_  uint16
	operandMask_   uint16
}

// Class Reference

func instructionClass() *instructionClass_ {
	return instructionClassReference_
}

var instructionClassReference_ = &instructionClass_{
	// Initialize the class constants.
	operationMask_: 0b1110000000000000,
	modifierMask_:  0b0001100000000000,
	operandMask_:   0b0000011111111111,
}
