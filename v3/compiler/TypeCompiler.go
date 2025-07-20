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
	rep "github.com/bali-nebula/go-document-repository/v3"
	fra "github.com/craterdog/go-component-framework/v7"
	uti "github.com/craterdog/go-missing-utilities/v7"
)

// CLASS INTERFACE

// Access Function

func TypeCompilerClass() TypeCompilerClassLike {
	return typeCompilerClass()
}

// Constructor Methods

func (c *typeCompilerClass_) TypeCompiler(
	repository rep.DocumentRepositoryLike,
) TypeCompilerLike {
	if uti.IsUndefined(repository) {
		panic("The \"repository\" attribute is required by this class.")
	}
	var instance = &typeCompiler_{
		// Initialize the instance attributes.
		repository_: repository,
		literals_:   fra.List[not.EntityLike](),

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

func (v *typeCompiler_) GetClass() TypeCompilerClassLike {
	return typeCompilerClass()
}

func (v *typeCompiler_) CompileType(
	document not.DocumentLike,
) {
	v.cleanType(document)
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
			v.compileMethods(methods)
			v.assembleMethods(methods) // This must happen after all compilation.
		}
	}
	v.addLiterals(associations)
}

// Attribute Methods

func (v *typeCompiler_) GetRepository() rep.DocumentRepositoryLike {
	return v.repository_
}

// Methodical Methods

func (v *typeCompiler_) ProcessAngle(
	angle string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessBinary(
	binary string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessBoolean(
	boolean string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessBytecode(
	bytecode string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessComment(
	comment string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessDelimiter(
	delimiter string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessDuration(
	duration string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessGlyph(
	glyph string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessIdentifier(
	identifier string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessMoment(
	moment string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessName(
	name string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessNarrative(
	narrative string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessNewline(
	newline string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessNote(
	note string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessNumber(
	number string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPattern(
	pattern string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPercentage(
	percentage string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessProbability(
	probability string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessQuote(
	quote string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessResource(
	resource string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSpace(
	space string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessSymbol(
	symbol string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessTag(
	tag string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessVersion(
	version string,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessAcceptClause(
	acceptClause not.AcceptClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessAcceptClauseSlot(
	acceptClause not.AcceptClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessActionInduction(
	actionInduction not.ActionInductionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessActionInduction(
	actionInduction not.ActionInductionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessActionInductionSlot(
	actionInduction not.ActionInductionLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessAnnotation(
	annotation not.AnnotationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessAnnotation(
	annotation not.AnnotationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessAnnotationSlot(
	annotation not.AnnotationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessArgument(
	argument not.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessArgument(
	argument not.ArgumentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessArgumentSlot(
	argument not.ArgumentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessArithmeticOperator(
	arithmeticOperator not.ArithmeticOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessArithmeticOperator(
	arithmeticOperator not.ArithmeticOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessArithmeticOperatorSlot(
	arithmeticOperator not.ArithmeticOperatorLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessAssignment(
	assignment not.AssignmentLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessAssignmentSlot(
	assignment not.AssignmentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessAssociation(
	association not.AssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessAssociation(
	association not.AssociationLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessAssociationSlot(
	association not.AssociationLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessAtLevel(
	atLevel not.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessAtLevel(
	atLevel not.AtLevelLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessAtLevelSlot(
	atLevel not.AtLevelLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessAttributes(
	attributes not.AttributesLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessAttributesSlot(
	attributes not.AttributesLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessBag(
	bag not.BagLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessBag(
	bag not.BagLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessBagSlot(
	bag not.BagLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessBra(
	bra not.BraLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessBra(
	bra not.BraLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessBraSlot(
	bra not.BraLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessBreakClause(
	breakClause not.BreakClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessBreakClauseSlot(
	breakClause not.BreakClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessCheckoutClause(
	checkoutClause not.CheckoutClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessCheckoutClauseSlot(
	checkoutClause not.CheckoutClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessCited(
	cited not.CitedLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
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

func (v *typeCompiler_) PreprocessComparisonOperator(
	comparisonOperator not.ComparisonOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessComparisonOperator(
	comparisonOperator not.ComparisonOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessComparisonOperatorSlot(
	comparisonOperator not.ComparisonOperatorLike,
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

func (v *typeCompiler_) PreprocessCondition(
	condition not.ConditionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessCondition(
	condition not.ConditionLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessConditionSlot(
	condition not.ConditionLike,
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
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessDiscardClause(
	discardClause not.DiscardClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessDiscardClauseSlot(
	discardClause not.DiscardClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
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
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessDocumentSlot(
	document not.DocumentLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
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

func (v *typeCompiler_) PreprocessEmpty(
	empty not.EmptyLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessEmpty(
	empty not.EmptyLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessEmptySlot(
	empty not.EmptyLike,
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

func (v *typeCompiler_) PreprocessKet(
	ket not.KetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessKet(
	ket not.KetLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessKetSlot(
	ket not.KetLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessLetClause(
	letClause not.LetClauseLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessLetClauseSlot(
	letClause not.LetClauseLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessLexicalOperator(
	lexicalOperator not.LexicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessLexicalOperator(
	lexicalOperator not.LexicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessLexicalOperatorSlot(
	lexicalOperator not.LexicalOperatorLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
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

func (v *typeCompiler_) PreprocessLogicalOperator(
	logicalOperator not.LogicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessLogicalOperator(
	logicalOperator not.LogicalOperatorLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessLogicalOperatorSlot(
	logicalOperator not.LogicalOperatorLike,
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
	// TBD - Add the method implementation.
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

func (v *typeCompiler_) PreprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessPredicate(
	predicate not.PredicateLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessPredicateSlot(
	predicate not.PredicateLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PreprocessPrimitive(
	primitive not.PrimitiveLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
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
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) PostprocessValue(
	value not.ValueLike,
	index_ uint,
	count_ uint,
) {
	// TBD - Add the method implementation.
}

func (v *typeCompiler_) ProcessValueSlot(
	value not.ValueLike,
	slot_ uint,
) {
	// TBD - Add the method implementation.
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

func (v *typeCompiler_) addLiterals(
	associations fra.ListLike[not.AssociationLike],
) {
	var items = not.Items("[", v.literals_, "]")
	var collection = not.Collection(items)
	var component = not.Component(collection)
	var document = not.Document(component, nil, "")
	var association = not.Association(
		not.Primitive(not.Element("$literals")),
		":",
		document,
	)
	associations.AppendValue(association)
}

func (v *typeCompiler_) assembleMethod(
	document not.DocumentLike,
) {
}

func (v *typeCompiler_) assembleMethods(
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

func (v *typeCompiler_) cleanMethod(
	document not.DocumentLike,
) {
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var index uti.Index
	var iterator = associations.GetIterator()
	iterator.ToEnd() // Must work backwards when removing multiple associations.
	for index = uti.Index(iterator.GetSize()); iterator.HasPrevious(); index-- {
		var association = iterator.GetPrevious()
		var key = association.GetPrimitive().GetAny().(string)
		switch key {
		case "$instructions", "$bytecode", "$arguments", "$variables",
			"$messages", "$addresses":
			associations.RemoveValue(index)
		}
	}
}

func (v *typeCompiler_) compileMethod(
	document not.DocumentLike,
) {
	v.cleanMethod(document)
}

func (v *typeCompiler_) compileMethods(
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
		v.compileMethod(method)
	}
}

func (v *typeCompiler_) cleanType(
	document not.DocumentLike,
) {
	var component = document.GetComponent()
	var collection = component.GetAny().(not.CollectionLike)
	var attributes = collection.GetAny().(not.AttributesLike)
	var associations = attributes.GetAssociations()
	var index uti.Index
	var iterator = associations.GetIterator()
	for index = 1; iterator.HasNext(); index++ {
		var association = iterator.GetNext()
		var key = association.GetPrimitive().GetAny().(string)
		if key == "$literals" {
			associations.RemoveValue(index)
		}
	}
}

// Instance Structure

type typeCompiler_ struct {
	// Declare the instance attributes.
	repository_ rep.DocumentRepositoryLike
	visitor_    not.VisitorLike
	literals_   fra.ListLike[not.EntityLike]

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
