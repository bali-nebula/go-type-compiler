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

// INSTANCE INTERFACE

// Principal Methods

func (v instruction_) GetClass() InstructionClassLike {
	return instructionClass()
}

func (v instruction_) AsIntrinsic() uint16 {
	return uint16(v)
}

func (v instruction_) AsString() string {
	var result string
	if v.AsIntrinsic() == 0 {
		return "SKIP"
	}
	var operation = v.GetOperation()
	var modifier = v.GetModifier()
	var operand = v.GetOperand()
	var index, address string
	if operation == Jump || (operation == Push && modifier == Handler) {
		// Treat the operand as an address "[xHEX]".
		address = fmt.Sprintf("[x%03x]", operand)
	} else {
		// Treat the operand as an index "DECI".
		index = fmt.Sprintf("%d", operand)
	}
	switch operation {
	case Jump:
		result = "JUMP TO " + address
		switch modifier {
		case Any:
		case Empty:
			result += " ON EMPTY"
		case None:
			result += " ON NONE"
		case False:
			result += " ON FALSE"
		}
	case Push:
		result = "PUSH "
		switch modifier {
		case Handler:
			result += "HANDLER " + address
		case Literal:
			result += "LITERAL " + index
		case Constant:
			result += "CONSTANT " + index
		case Argument:
			result += "ARGUMENT " + index
		}
	case Pull:
		result = "PULL "
		switch modifier {
		case Handler:
			result += "HANDLER"
		case Exception:
			result += "EXCEPTION"
		case Component:
			result += "COMPONENT"
		case Result:
			result += "RESULT"
		}
	case Load:
		result = "LOAD "
		switch modifier {
		case Contract:
			result += "CONTRACT " + index
		case Draft:
			result += "DRAFT " + index
		case Message:
			result += "MESSAGE " + index
		case Variable:
			result += "VARIABLE " + index
		}
	case Save:
		result = "SAVE "
		switch modifier {
		case Contract:
			result += "CONTRACT " + index
		case Draft:
			result += "DRAFT " + index
		case Message:
			result += "MESSAGE " + index
		case Variable:
			result += "VARIABLE " + index
		}
	case Drop:
		result = "DROP "
		switch modifier {
		case Contract:
			result += "CONTRACT " + index
		case Draft:
			result += "DRAFT " + index
		case Message:
			result += "MESSAGE " + index
		case Variable:
			result += "VARIABLE " + index
		}
	case Call:
		result = "CALL " + index
		switch modifier {
		case With0Arguments:
		case With1Argument:
			result += " WITH 1 ARGUMENT"
		case With2Arguments:
			result += " WITH 2 ARGUMENTS"
		case With3Arguments:
			result += " WITH 3 ARGUMENTS"
		}
	case Send:
		result = "SEND " + index + " TO "
		switch modifier {
		case Component:
			result += "COMPONENT"
		case ComponentWithArguments:
			result += "COMPONENT WITH ARGUMENTS"
		case Contract:
			result += "CONTRACT"
		case ContractWithArguments:
			result += "CONTRACT WITH ARGUMENTS"
		}
	}
	return result
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
