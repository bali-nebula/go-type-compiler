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
	not "github.com/bali-nebula/go-document-notation/v3"
	fra "github.com/craterdog/go-component-framework/v7"
)

// CLASS INTERFACE

// Access Function

func TypeAssemblerClass() TypeAssemblerClassLike {
	return typeAssemblerClass()
}

// Constructor Methods

func (c *typeAssemblerClass_) TypeAssembler() TypeAssemblerLike {
	var instance = &typeAssembler_{
		// Initialize the instance attributes.
		literals_: fra.List[not.EntityLike](),

		// Initialize the inherited aspects.
		Methodical: not.Processor(),
	}
	instance.visitor_ = not.Visitor(instance)
	return instance
}

// Constant Methods

// Function Methods

// INSTANCE INTERFACE

// Principal Methods

func (v *typeAssembler_) GetClass() TypeAssemblerClassLike {
	return typeAssemblerClass()
}

func (v *typeAssembler_) AssembleType(
	document not.DocumentLike,
) {
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var key = association.GetPrimitive().GetAny().(string)
		if key == "$methods" {
			var methods = association.GetDocument()
			v.assembleMethods(methods)
		}
	}
}

// Attribute Methods

// Methodical Methods

func (v *typeAssembler_) ProcessAngle(
	angle string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessBinary(
	binary string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessBoolean(
	boolean string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessBytecode(
	bytecode string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessComment(
	comment string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessDelimiter(
	delimiter string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessDuration(
	duration string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessGlyph(
	glyph string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessIdentifier(
	identifier string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMoment(
	moment string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessName(
	name string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessNarrative(
	narrative string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessNewline(
	newline string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessNote(
	note string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessNumber(
	number string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPattern(
	pattern string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPercentage(
	percentage string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessProbability(
	probability string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessQuote(
	quote string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessResource(
	resource string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSpace(
	space string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSymbol(
	symbol string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessTag(
	tag string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessVersion(
	version string,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessAcceptClauseSlot(
	acceptClause not.AcceptClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessActionInduction(
	actionInduction not.ActionInductionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessActionInduction(
	actionInduction not.ActionInductionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessActionInductionSlot(
	actionInduction not.ActionInductionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessAnnotation(
	annotation not.AnnotationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessAnnotation(
	annotation not.AnnotationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessAnnotationSlot(
	annotation not.AnnotationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessArgument(
	argument not.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessArgument(
	argument not.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessArgumentSlot(
	argument not.ArgumentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessArithmeticOperator(
	arithmeticOperator not.ArithmeticOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessArithmeticOperator(
	arithmeticOperator not.ArithmeticOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessArithmeticOperatorSlot(
	arithmeticOperator not.ArithmeticOperatorLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessAssignmentSlot(
	assignment not.AssignmentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessAssociation(
	association not.AssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessAssociation(
	association not.AssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessAssociationSlot(
	association not.AssociationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessAtLevel(
	atLevel not.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessAtLevel(
	atLevel not.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessAtLevelSlot(
	atLevel not.AtLevelLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessAttributesSlot(
	attributes not.AttributesLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessBag(
	bag not.BagLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessBag(
	bag not.BagLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessBagSlot(
	bag not.BagLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessBra(
	bra not.BraLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessBra(
	bra not.BraLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessBraSlot(
	bra not.BraLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessBreakClauseSlot(
	breakClause not.BreakClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessCheckoutClauseSlot(
	checkoutClause not.CheckoutClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessCited(
	cited not.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessCited(
	cited not.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessCitedSlot(
	cited not.CitedLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessCollection(
	collection not.CollectionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessCollection(
	collection not.CollectionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessCollectionSlot(
	collection not.CollectionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessComparisonOperator(
	comparisonOperator not.ComparisonOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessComparisonOperator(
	comparisonOperator not.ComparisonOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessComparisonOperatorSlot(
	comparisonOperator not.ComparisonOperatorLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessComplement(
	complement not.ComplementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessComplement(
	complement not.ComplementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessComplementSlot(
	complement not.ComplementLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessComponent(
	component not.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessComponent(
	component not.ComponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessComponentSlot(
	component not.ComponentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessCondition(
	condition not.ConditionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessCondition(
	condition not.ConditionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessConditionSlot(
	condition not.ConditionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessContinueClause(
	continueClause not.ContinueClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessContinueClauseSlot(
	continueClause not.ContinueClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessDiscardClauseSlot(
	discardClause not.DiscardClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessDoClause(
	doClause not.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessDoClause(
	doClause not.DoClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessDoClauseSlot(
	doClause not.DoClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessDocument(
	document not.DocumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessDocumentSlot(
	document not.DocumentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessDraft(
	draft not.DraftLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessDraft(
	draft not.DraftLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessDraftSlot(
	draft not.DraftLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessElement(
	element not.ElementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessElement(
	element not.ElementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessElementSlot(
	element not.ElementLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessEmpty(
	empty not.EmptyLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessEmpty(
	empty not.EmptyLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessEmptySlot(
	empty not.EmptyLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessEntity(
	entity not.EntityLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessEntity(
	entity not.EntityLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessEntitySlot(
	entity not.EntityLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessEvent(
	event not.EventLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessEvent(
	event not.EventLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessEventSlot(
	event not.EventLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessException(
	exception not.ExceptionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessException(
	exception not.ExceptionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessExceptionSlot(
	exception not.ExceptionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessExpression(
	expression not.ExpressionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessExpressionSlot(
	expression not.ExpressionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessFailure(
	failure not.FailureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessFailure(
	failure not.FailureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessFailureSlot(
	failure not.FailureLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessFlowControl(
	flowControl not.FlowControlLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessFlowControl(
	flowControl not.FlowControlLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessFlowControlSlot(
	flowControl not.FlowControlLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessFunction(
	function not.FunctionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessFunctionSlot(
	function not.FunctionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessIfClause(
	ifClause not.IfClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessIfClauseSlot(
	ifClause not.IfClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessIndex(
	index not.IndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessIndex(
	index not.IndexLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessIndexSlot(
	index not.IndexLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessIndirect(
	indirect not.IndirectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessIndirect(
	indirect not.IndirectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessIndirectSlot(
	indirect not.IndirectLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessInverse(
	inverse not.InverseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessInverse(
	inverse not.InverseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessInverseSlot(
	inverse not.InverseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessInversion(
	inversion not.InversionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessInversion(
	inversion not.InversionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessInversionSlot(
	inversion not.InversionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessInvocation(
	invocation not.InvocationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessInvocation(
	invocation not.InvocationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessInvocationSlot(
	invocation not.InvocationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessInvoke(
	invoke not.InvokeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessInvoke(
	invoke not.InvokeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessInvokeSlot(
	invoke not.InvokeLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessItems(
	items not.ItemsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessItems(
	items not.ItemsLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessItemsSlot(
	items not.ItemsLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessKet(
	ket not.KetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessKet(
	ket not.KetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessKetSlot(
	ket not.KetLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessLetClauseSlot(
	letClause not.LetClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessLexicalOperator(
	lexicalOperator not.LexicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessLexicalOperator(
	lexicalOperator not.LexicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessLexicalOperatorSlot(
	lexicalOperator not.LexicalOperatorLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessLine(
	line not.LineLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessLine(
	line not.LineLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessLineSlot(
	line not.LineLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessLogical(
	logical not.LogicalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessLogical(
	logical not.LogicalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessLogicalSlot(
	logical not.LogicalLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessLogicalOperator(
	logicalOperator not.LogicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessLogicalOperator(
	logicalOperator not.LogicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessLogicalOperatorSlot(
	logicalOperator not.LogicalOperatorLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessMagnitude(
	magnitude not.MagnitudeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMagnitudeSlot(
	magnitude not.MagnitudeLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessMainClause(
	mainClause not.MainClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessMainClause(
	mainClause not.MainClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMainClauseSlot(
	mainClause not.MainClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessMatchingClause(
	matchingClause not.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessMatchingClause(
	matchingClause not.MatchingClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMatchingClauseSlot(
	matchingClause not.MatchingClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessMessage(
	message not.MessageLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessMessage(
	message not.MessageLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMessageSlot(
	message not.MessageLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessMessageHandling(
	messageHandling not.MessageHandlingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessMessageHandling(
	messageHandling not.MessageHandlingLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMessageHandlingSlot(
	messageHandling not.MessageHandlingLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessMethod(
	method not.MethodLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessMethodSlot(
	method not.MethodLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessNotarizeClause(
	notarizeClause not.NotarizeClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessNotarizeClauseSlot(
	notarizeClause not.NotarizeClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessNumerical(
	numerical not.NumericalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessNumerical(
	numerical not.NumericalLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessNumericalSlot(
	numerical not.NumericalLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessOnClause(
	onClause not.OnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessOnClauseSlot(
	onClause not.OnClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessOperation(
	operation not.OperationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessOperation(
	operation not.OperationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessOperationSlot(
	operation not.OperationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessParameters(
	parameters not.ParametersLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessParameters(
	parameters not.ParametersLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessParametersSlot(
	parameters not.ParametersLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessPostClause(
	postClause not.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessPostClause(
	postClause not.PostClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPostClauseSlot(
	postClause not.PostClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessPrecedence(
	precedence not.PrecedenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPrecedenceSlot(
	precedence not.PrecedenceLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPredicateSlot(
	predicate not.PredicateLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessPrimitive(
	primitive not.PrimitiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessPrimitive(
	primitive not.PrimitiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPrimitiveSlot(
	primitive not.PrimitiveLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessProcedure(
	procedure not.ProcedureLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessProcedureSlot(
	procedure not.ProcedureLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessPublishClause(
	publishClause not.PublishClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessPublishClauseSlot(
	publishClause not.PublishClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessRange(
	range_ not.RangeLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessRangeSlot(
	range_ not.RangeLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessRecipient(
	recipient not.RecipientLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessRecipient(
	recipient not.RecipientLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessRecipientSlot(
	recipient not.RecipientLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessReferent(
	referent not.ReferentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessReferent(
	referent not.ReferentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessReferentSlot(
	referent not.ReferentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessRejectClause(
	rejectClause not.RejectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessRejectClauseSlot(
	rejectClause not.RejectClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessRepositoryAccess(
	repositoryAccess not.RepositoryAccessLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessRepositoryAccess(
	repositoryAccess not.RepositoryAccessLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessRepositoryAccessSlot(
	repositoryAccess not.RepositoryAccessLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessResult(
	result not.ResultLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessResult(
	result not.ResultLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessResultSlot(
	result not.ResultLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessRetrieveClause(
	retrieveClause not.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessRetrieveClause(
	retrieveClause not.RetrieveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessRetrieveClauseSlot(
	retrieveClause not.RetrieveClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessReturnClause(
	returnClause not.ReturnClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessReturnClauseSlot(
	returnClause not.ReturnClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessSaveClause(
	saveClause not.SaveClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSaveClauseSlot(
	saveClause not.SaveClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessSelectClause(
	selectClause not.SelectClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSelectClauseSlot(
	selectClause not.SelectClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessSequence(
	sequence not.SequenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessSequence(
	sequence not.SequenceLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSequenceSlot(
	sequence not.SequenceLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessStatement(
	statement not.StatementLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessStatementSlot(
	statement not.StatementLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessString(
	string_ not.StringLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessString(
	string_ not.StringLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessStringSlot(
	string_ not.StringLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessSubcomponent(
	subcomponent not.SubcomponentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSubcomponentSlot(
	subcomponent not.SubcomponentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessSubject(
	subject not.SubjectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessSubject(
	subject not.SubjectLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessSubjectSlot(
	subject not.SubjectLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessTarget(
	target not.TargetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessTarget(
	target not.TargetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessTargetSlot(
	target not.TargetLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessTemplate(
	template not.TemplateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessTemplate(
	template not.TemplateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessTemplateSlot(
	template not.TemplateLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessThrowClause(
	throwClause not.ThrowClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessThrowClauseSlot(
	throwClause not.ThrowClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessValue(
	value not.ValueLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessValue(
	value not.ValueLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessValueSlot(
	value not.ValueLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessVariable(
	variable not.VariableLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessVariable(
	variable not.VariableLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessVariableSlot(
	variable not.VariableLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessWhileClause(
	whileClause not.WhileClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessWhileClauseSlot(
	whileClause not.WhileClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PreprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) PostprocessWithClause(
	withClause not.WithClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeAssembler_) ProcessWithClauseSlot(
	withClause not.WithClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

// PROTECTED INTERFACE

// Private Methods

func (v *typeAssembler_) assembleMethod(
	document not.DocumentLike,
) {
}

func (v *typeAssembler_) assembleMethods(
	document not.DocumentLike,
) {
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var iterator = associations.GetIterator()
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var method = association.GetDocument()
		v.assembleMethod(method)
	}
}

// Instance Structure

type typeAssembler_ struct {
	// Declare the instance attributes.
	visitor_  not.VisitorLike
	literals_ fra.ListLike[not.EntityLike]

	// Declare the inherited aspects.
	not.Methodical
}

// Class Structure

type typeAssemblerClass_ struct {
	// Declare the class constants.
}

// Class Reference

func typeAssemblerClass() *typeAssemblerClass_ {
	return typeAssemblerClassReference_
}

var typeAssemblerClassReference_ = &typeAssemblerClass_{
	// Initialize the class constants.
}
