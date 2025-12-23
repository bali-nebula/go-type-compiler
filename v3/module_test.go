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

package module_test

import (
	ins "github.com/bali-nebula/go-bali-instructions/v3"
	typ "github.com/bali-nebula/go-type-compiler/v3"
	com "github.com/craterdog/go-essential-composites/v8"
	pri "github.com/craterdog/go-essential-primitives/v8"
	uti "github.com/craterdog/go-essential-utilities/v8"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const directory = "./test/"

var literals = com.SetFromArray[string]([]string{
	"`none`",
	"`" + `
    ">
        This is a literal text string
        containing an ` + "\\`" + ` and spanning multiple lines.
    <"
` + "`",
	"`" + `
    {
        $foo := bar
        $bar := baz
        $baz := foo
    }($bar: 5)
` + "`",
})

var constants = com.SetFromArray[string]([]string{
	"$constant",
})

func TestMethodAssembler(t *tes.T) {
	var filename = directory + "instructions.basm"
	var source = uti.ReadFile(filename)
	ass.True(t, len(source) > 0)
	var method = ins.ParseMethod(source)
	var assembler = typ.Assembler()
	var instructions = assembler.AssembleMethod(literals, constants, method)
	var formatted = typ.AssemblerClass().FormatInstructions(instructions)
	filename = directory + "instructions.dbug"
	uti.WriteFile(filename, formatted)
	var bytecode = pri.Bytecode(instructions.AsArray())
	filename = directory + "bytecode.dbug"
	uti.WriteFile(filename, bytecode.AsSource())
}
