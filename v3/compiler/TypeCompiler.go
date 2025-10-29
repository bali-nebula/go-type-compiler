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
	ass "github.com/bali-nebula/go-type-compiler/v3/assembler"
	com "github.com/craterdog/go-essential-composites/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	stc "strconv"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func TypeCompilerClass() TypeCompilerClassLike {
	return typeCompilerClass()
}

// Constructor Methods

func (c *typeCompilerClass_) TypeCompiler() TypeCompilerLike {
	var instance = &typeCompiler_{
		// Initialize the instance attributes.
		literals_:  com.Set[string](),
		constants_: com.Set[string](),

		// Initialize the inherited aspects.
		Methodical: not.Processor(),
	}
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *typeCompiler_) GetClass() TypeCompilerClassLike {
	return typeCompilerClass()
}

func (v *typeCompiler_) CompileType(
	type_ not.DocumentLike,
) {
	v.compileMethods(type_)
	v.assembleMethods(type_)
}

// Attribute Methods

// Methodical Methods

func (v *typeCompiler_) ProcessAcceptClauseSlot(
	checkoutClause not.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendNote("Retrieve the message to be accepted.")
	}
}

func (v *typeCompiler_) PostprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	var message = v.createVariable("message")
	v.appendSave("VARIABLE", message)
	v.appendNote("Extract and save the name of the message bag.")
	v.appendLoad("VARIABLE", message)
	v.appendPush("LITERAL", "$bag")
	v.appendCall("$getAttribute", 2) // getAttribute(message, "$bag")
	var bag = v.createVariable("bag")
	v.appendSave("VARIABLE", bag)
	v.appendNote("Drop the message from the named message bag.")
	v.appendLoad("VARIABLE", message)
	v.appendDrop("MESSAGE", bag)
}

func (v *typeCompiler_) PostprocessAssociation(
	association not.AssociationLike,
	index_ uint,
	count_ uint,
) {
	// The key and value of the association are already on the stack.
	v.appendCall("$association", 2) // association(key, value)
}

func (v *typeCompiler_) ProcessAttributesSlot(
	attributes not.AttributesLike,
	slot_ uint,
) {
	switch slot {
	case 1:
		v.appendNote("Place a catalog of attributes on the stack.")
		v.appendCall("$catalog", 0) // catalog()
	}
}

func (v *typeCompiler_) PostprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	var iterator = v.context_.GetIterator()
	for iterator.HasNext() {
		var context = iterator.GetNext()
		var loopLabel = context.GetLabel("$loopLabel")
		if uti.IsDefined(loopLabel) {
			var doneLabel = context.GetLabel("$doneLabel")
			v.appendJump(doneLabel, "")
			return
		}
	}
	var message = "A break statement was found with no enclosing loop."
	panic(message)
}

func (v *typeCompiler_) ProcessCheckoutClauseSlot(
	checkoutClause not.CheckoutClauseLike,
	slot_ uint,
) {
	switch slot_ {
	case 1:
		v.appendNote("Determine the recipient of the draft document.")
	case 2:
		v.appendNote("Determine the version level to be incremented.")
	case 3:
		var atLevel = checkoutClause.GetOptionalAtLevel()
		if uti.IsUndefined(atLevel) {
			v.appendPush("LITERAL", "0")
		}
		v.appendNote("Determine the citation to the contract to be checked out.")
	}
}

func (v *typeCompiler_) PostprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	var citation = v.createVariable("citation")
	v.appendSave("VARIABLE", citation)
	v.appendNote("Save a draft copy of the named contract from the repository.")
	v.appendLoad("CONTRACT", citation)
	v.appendPush("LITERAL", "$draft")
	v.appendCall("$attribute", 2) // attribute(contract, "$draft")
	v.appendCall("$duplicate", 1) // duplicate(draft)
	var draft = v.createVariable("draft")
	v.appendSave("VARIABLE", draft)
	v.appendLoad("VARIABLE", draft)
	v.appendNote("Determine the next version of the draft document.")
	v.appendCall("$parameters", 1) // parameters(draft)
	v.appendPush("LITERAL", "$version")
	v.appendCall("$attribute", 2)   // attribute(parameters, "$version")
	v.appendCall("$nextVersion", 2) // nextVersion(version, level)
	var version = v.createVariable("version")
	v.appendSave("VARIABLE", version)
	v.appendNote("Set the new version for the draft document.")
	v.appendLoad("VARIABLE", draft)
	v.appendPush("LITERAL", "$version")
	v.appendLoad("VARIABLE", version)
	v.appendCall("$setParameter", 3) // setParameter(draft, "$version", version)
	v.appendLoad("VARIABLE", draft)
	var recipient = checkoutClause.GetRecipient()
	v.setRecipient(recipient)
}

func (v *typeCompiler_) PreprocessCited(
	cited not.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
	v.appendPull("")
}

func (v *typeCompiler_) PostprocessCited(
	cited not.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessCitedSlot(
	cited not.CitedLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessCollection(
	collection not.CollectionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessCollection(
	collection not.CollectionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessCollectionSlot(
	collection not.CollectionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessComplement(
	complement not.ComplementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessComplement(
	complement not.ComplementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessComplementSlot(
	complement not.ComplementLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessComponent(
	component not.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessComponent(
	component not.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessComponentSlot(
	component not.ComponentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessContinueClauseSlot(
	continueClause not.ContinueClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	v.appendNote("Save the citation to the document.")
}

func (v *typeCompiler_) PostprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	var citation = v.createVariable("citation")
	v.appendSave("VARIABLE", citation)
	v.appendNote("Drop the cited document from the repository.")
	v.appendDrop("DRAFT", citation)
}

func (v *typeCompiler_) PreprocessDoClause(
	doClause not.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessDoClause(
	doClause not.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessDoClauseSlot(
	doClause not.DoClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// The component, parameters and note are already on the stack.
	v.appendCall("$document", 3)
}

func (v *typeCompiler_) ProcessDocumentSlot(
	document not.DocumentLike,
	slot_ uint,
) {
	var parameters = document.GetOptionalParameters()
	var note = document.GetOptionalNote()
	switch {
	case uti.IsUndefined(parameters):
		v.appendPush("LITERAL", "none")
	case uti.IsUndefined(note):
		v.appendPush("LITERAL", "none")
	}
}

func (v *typeCompiler_) PreprocessDraft(
	draft not.DraftLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessDraft(
	draft not.DraftLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessDraftSlot(
	draft not.DraftLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessElement(
	element not.ElementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessElement(
	element not.ElementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessElementSlot(
	element not.ElementLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessEntity(
	entity not.EntityLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessEntity(
	entity not.EntityLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessEntitySlot(
	entity not.EntityLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessEvent(
	event not.EventLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessEvent(
	event not.EventLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessEventSlot(
	event not.EventLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessException(
	exception not.ExceptionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessException(
	exception not.ExceptionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessExceptionSlot(
	exception not.ExceptionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessExpressionSlot(
	expression not.ExpressionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessFailure(
	failure not.FailureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessFailure(
	failure not.FailureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessFailureSlot(
	failure not.FailureLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessFlowControl(
	flowControl not.FlowControlLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessFlowControl(
	flowControl not.FlowControlLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessFlowControlSlot(
	flowControl not.FlowControlLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessFunctionSlot(
	function not.FunctionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
	v.setLabel("")
}

func (v *typeCompiler_) PostprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessIfClauseSlot(
	ifClause not.IfClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessIndex(
	index not.IndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessIndex(
	index not.IndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessIndexSlot(
	index not.IndexLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessIndirect(
	indirect not.IndirectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessIndirect(
	indirect not.IndirectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessIndirectSlot(
	indirect not.IndirectLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessInverse(
	inverse not.InverseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessInverse(
	inverse not.InverseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessInverseSlot(
	inverse not.InverseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessInversion(
	inversion not.InversionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessInversion(
	inversion not.InversionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessInversionSlot(
	inversion not.InversionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessInvocation(
	invocation not.InvocationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessInvocation(
	invocation not.InvocationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessInvocationSlot(
	invocation not.InvocationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessInvoke(
	invoke not.InvokeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessInvoke(
	invoke not.InvokeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessInvokeSlot(
	invoke not.InvokeLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessItems(
	items not.ItemsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessItems(
	items not.ItemsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessItemsSlot(
	items not.ItemsLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// At this point, the recipient and expression are already on the stack.
	var assignment = letClause.GetAssignment()
	switch assignment.GetAny().(string) {
	case ":=":
		v.appendCall("$assign", 2)
	case "?=":
		v.appendCall("$default", 2)
	case "+=":
		v.appendCall("$sum", 2)
	case "-=":
		v.appendCall("$difference", 2)
	case "*=":
		v.appendCall("$product", 2)
	case "/=":
		v.appendCall("$quotient", 2)
	case "&=":
		v.appendCall("$chain", 2)
	}
}

func (v *typeCompiler_) PreprocessLine(
	line not.LineLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessLine(
	line not.LineLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessLineSlot(
	line not.LineLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessLogical(
	logical not.LogicalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessLogical(
	logical not.LogicalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessLogicalSlot(
	logical not.LogicalLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMagnitudeSlot(
	magnitude not.MagnitudeLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessMainClause(
	mainClause not.MainClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessMainClause(
	mainClause not.MainClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMainClauseSlot(
	mainClause not.MainClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessMatchingClause(
	matchingClause not.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessMatchingClause(
	matchingClause not.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMatchingClauseSlot(
	matchingClause not.MatchingClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessMessage(
	message not.MessageLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
	v.appendSend("", "", false)
}

func (v *typeCompiler_) PostprocessMessage(
	message not.MessageLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMessageSlot(
	message not.MessageLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessMessageHandling(
	messageHandling not.MessageHandlingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessMessageHandling(
	messageHandling not.MessageHandlingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMessageHandlingSlot(
	messageHandling not.MessageHandlingLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMethodSlot(
	method not.MethodLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessNotarizeClauseSlot(
	notarizeClause not.NotarizeClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessNumerical(
	numerical not.NumericalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessNumerical(
	numerical not.NumericalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessNumericalSlot(
	numerical not.NumericalLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessOnClauseSlot(
	onClause not.OnClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessOperation(
	operation not.OperationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessOperation(
	operation not.OperationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessOperationSlot(
	operation not.OperationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessParameters(
	parameters not.ParametersLike,
	index_ uint,
	count_ uint,
) {
	v.appendNote("Place a catalog of parameters on the stack.")
	v.appendCall("$catalog", 0)
}

func (v *typeCompiler_) PostprocessParameters(
	parameters not.ParametersLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessParametersSlot(
	parameters not.ParametersLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessPostClause(
	postClause not.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessPostClause(
	postClause not.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPostClauseSlot(
	postClause not.PostClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPrecedenceSlot(
	precedence not.PrecedenceLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// At this point, the two operands are already on the stack.
	var operation = predicate.GetOperation()
	switch actual := operation.GetAny().(type) {
	case not.LexicalOperatorLike:
		switch actual.GetAny().(string) {
		case "&":
			v.appendCall("$chain", 2)
		}
	case not.LogicalOperatorLike:
		switch actual.GetAny().(string) {
		case "and":
			v.appendCall("$and", 2)
		case "san":
			v.appendCall("$san", 2)
		case "ior":
			v.appendCall("$ior", 2)
		case "xor":
			v.appendCall("$xor", 2)
		}
	case not.ArithmeticOperatorLike:
		switch actual.GetAny().(string) {
		case "+":
			v.appendCall("$sum", 2)
		case "-":
			v.appendCall("$difference", 2)
		case "*":
			v.appendCall("$product", 2)
		case "/":
			v.appendCall("$quotient", 2)
		case "%":
			v.appendCall("$remainder", 2)
		case "^":
			v.appendCall("$exponential", 2)
		}
	case not.ComparisonOperatorLike:
		switch actual.GetAny().(string) {
		case "<":
			v.appendCall("$isLess", 2)
		case "=":
			v.appendCall("$areEqual", 2)
		case ">":
			v.appendCall("$isMore", 2)
		case "is":
			v.appendCall("$areSame", 2)
		case "matches":
			v.appendCall("$doesMatch", 2)
		}
	}
}

func (v *typeCompiler_) PreprocessPrimitive(
	primitive not.PrimitiveLike,
	index_ uint,
	count_ uint,
) {
	var literal string
	switch actual := primitive.GetAny().(type) {
	case not.ElementLike:
		literal = actual.GetAny().(string)
	case not.StringLike:
		literal = actual.GetAny().(string)
	}
	literal = "`<\n    " + literal + "\n>`"
	v.appendPush("LITERAL", literal)
}

func (v *typeCompiler_) PostprocessPrimitive(
	primitive not.PrimitiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPrimitiveSlot(
	primitive not.PrimitiveLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
	v.pushContext(nil)
}

func (v *typeCompiler_) PostprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessProcedureSlot(
	procedure not.ProcedureLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPublishClauseSlot(
	publishClause not.PublishClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessRangeSlot(
	range_ not.RangeLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessRecipient(
	recipient not.RecipientLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessRecipient(
	recipient not.RecipientLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessRecipientSlot(
	recipient not.RecipientLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessReferent(
	referent not.ReferentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessReferent(
	referent not.ReferentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessReferentSlot(
	referent not.ReferentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessRejectClauseSlot(
	rejectClause not.RejectClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessRepositoryAccess(
	repositoryAccess not.RepositoryAccessLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessRepositoryAccess(
	repositoryAccess not.RepositoryAccessLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessRepositoryAccessSlot(
	repositoryAccess not.RepositoryAccessLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessResult(
	result not.ResultLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessResult(
	result not.ResultLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessResultSlot(
	result not.ResultLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessRetrieveClause(
	retrieveClause not.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessRetrieveClause(
	retrieveClause not.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessRetrieveClauseSlot(
	retrieveClause not.RetrieveClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessReturnClauseSlot(
	returnClause not.ReturnClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSaveClauseSlot(
	saveClause not.SaveClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSelectClauseSlot(
	selectClause not.SelectClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessSequence(
	sequence not.SequenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessSequence(
	sequence not.SequenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSequenceSlot(
	sequence not.SequenceLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessStatementSlot(
	statement not.StatementLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessString(
	string_ not.StringLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessString(
	string_ not.StringLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessStringSlot(
	string_ not.StringLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSubcomponentSlot(
	subcomponent not.SubcomponentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessSubject(
	subject not.SubjectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessSubject(
	subject not.SubjectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSubjectSlot(
	subject not.SubjectLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessTarget(
	target not.TargetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessTarget(
	target not.TargetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessTargetSlot(
	target not.TargetLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessTemplate(
	template not.TemplateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessTemplate(
	template not.TemplateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessTemplateSlot(
	template not.TemplateLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessThrowClauseSlot(
	throwClause not.ThrowClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessValue(
	value not.ValueLike,
	index_ uint,
	count_ uint,
) {
	var identifier = value.GetIdentifier()
	v.appendLoad("VARIABLE", identifier)
}

func (v *typeCompiler_) PreprocessVariable(
	variable not.VariableLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessVariable(
	variable not.VariableLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessVariableSlot(
	variable not.VariableLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessWhileClauseSlot(
	whileClause not.WhileClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessWithClauseSlot(
	withClause not.WithClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

// PROTECTED INTERFACE

// Private Methods

func (v *typeCompiler_) appendNote(
	note string,
) {
	if uti.IsDefined(v.label_) {
		v.assembly_.WriteString("\n" + v.label_ + ":" + "\n")
		v.label_ = ""
	}
	var instruction = "NOTE --" + note
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendSkip() {
	var instruction = "SKIP"
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendJump(
	label string,
	modifier string,
) {
	var instruction = "JUMP TO " + label
	if uti.IsDefined(modifier) {
		instruction += " " + modifier
	}
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendPush(
	modifier string,
	value string,
) {
	if modifier == "LITERAL" {
		value = "`" + value + "`"
	}
	var instruction = "PUSH " + modifier + " " + value
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendPull(
	modifier string,
) {
	var instruction = "PULL " + modifier
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendLoad(
	modifier string,
	symbol string,
) {
	v.variables_.AddValue(symbol)
	var instruction = "LOAD " + modifier + " " + symbol
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendSave(
	modifier string,
	symbol string,
) {
	v.variables_.AddValue(symbol)
	var instruction = "SAVE " + modifier + " " + symbol
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendDrop(
	modifier string,
	symbol string,
) {
	v.variables_.AddValue(symbol)
	var instruction = "DROP " + modifier + " " + symbol
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendCall(
	intrinsic string,
	arguments uint8,
) {
	var instruction = "CALL " + intrinsic
	switch arguments {
	case 0:
	case 1:
		instruction += " WITH 1 ARGUMENT"
	case 2:
		instruction += " WITH 2 ARGUMENTS"
	case 3:
		instruction += " WITH 3 ARGUMENTS"
	}
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) appendSend(
	message string,
	destination string,
	arguments bool,
) {
	var instruction = "SEND " + message + " TO " + destination
	if arguments {
		instruction += " WITH ARGUMENTS"
	}
	v.assembly_.WriteString(instruction + "\n")
}

func (v *typeCompiler_) assembleMethods(
	type_ not.DocumentLike,
) {
	var assembler = ass.MethodAssemblerClass().MethodAssembler(type_)
	var associations = v.getAssociations(type_, "$methods")
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var method = association.GetDocument()
		assembler.AssembleMethod(method)
	}
}

func (v *typeCompiler_) createVariable(
	variable string,
) string {
	v.counter_++
	return "$" + variable + "-" + stc.Itoa(int(v.counter_))
}

func (v *typeCompiler_) compileMethod(
	method not.DocumentLike,
) {
	v.counter_ = 1
	v.context_ = com.Stack[ContextLike]()
	v.arguments_ = com.Catalog[string, not.DocumentLike]()
	v.variables_ = com.Set[string]()
	v.messages_ = com.Set[string]()
	v.assembly_.Reset()
	not.Visitor(v).VisitDocument(method)
}

func (v *typeCompiler_) compileMethods(
	type_ not.DocumentLike,
) {
	var associations = v.getAssociations(type_, "$methods")
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var method = association.GetDocument()
		v.compileMethod(method)
	}
}

func (v *typeCompiler_) getAssociations(
	type_ not.DocumentLike,
	symbol string,
) com.Sequential[not.AssociationLike] {
	var primitive = not.Primitive(not.Element(symbol))
	var document = not.GetAttribute(type_, primitive)
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	return associations
}

func (v *typeCompiler_) pushContext(
	procedure not.ProcedureLike,
) {
	var labelPrefix string
	var statementCount = uint(procedure.GetLines().GetSize())
	if statementCount > 0 {
		var parent = v.context_.RemoveLast()
		v.context_.AddValue(parent) // HACK until GetLast() is supported.
		labelPrefix = parent.GetLabelPrefix()
		labelPrefix += stc.Itoa(int(parent.GetStatementNumber())) + "."
		labelPrefix += stc.Itoa(int(parent.GetBlockNumber())) + "."
		parent.IncrementBlockNumber()
	}
	var currentStatement not.StatementLike

	var context = ContextClass().Context(
		labelPrefix,
		statementCount,
		currentStatement,
	)
	v.context_.AddValue(context)
}

func (v *typeCompiler_) setLabel(
	label string,
) {
	if uti.IsDefined(v.label_) {
		v.appendSkip()
	}
	v.label_ = label
}

func (v *typeCompiler_) setRecipient(
	recipient not.RecipientLike,
) {
	/*
	switch actual := recipient.GetAny().(type) {
	case not.VariableLike:
		var symbol = actual.GetSymbol()
		v.appendSave("VARIABLE", symbol)
	case not.SubcomponentLike:
		var identifier = actual.GetIdentifier()
		v.appendLoad("VARIABLE", identifier)
		var indexes = actual.GetIndexes()
		var iterator = indexes.GetIterator()
		if iterator.GetSize() > 1 {
			for iterator.HasNext() {
				var index = iterator.GetNext()
	
				v.appendCall("$getAttribute", 2) // getAttribute(component, index)
			}
			return
		}
		v.appendCall("$setAttribute", 3) // setAttribute(component, index, value)
	}
	*/
}

// Instance Structure

type typeCompiler_ struct {
	// Declare the instance attributes.
	literals_  com.SetLike[string]
	constants_ com.SetLike[string]
	context_   com.StackLike[ContextLike]
	arguments_ com.CatalogLike[string, not.DocumentLike]
	variables_ com.SetLike[string]
	messages_  com.SetLike[string]
	assembly_  sts.Builder
	counter_   uint16
	label_     string

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type typeCompilerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func typeCompilerClass() *typeCompilerClass_ {
	return typeCompilerClassReference_
}

var typeCompilerClassReference_ = &typeCompilerClass_{
	// Initialize the class constants.
}
