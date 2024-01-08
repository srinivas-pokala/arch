// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390xasm

import (
	"bytes"
	"fmt"
	"strings"
)

type Inst struct {
	Op   Op     // Opcode mnemonic
	Enc  uint64 // Raw encoding bits (if Len == 8, this is the prefix word)
	Len  int    // Length of encoding in bytes.
	Args Args   // Instruction arguments, in Power ISA manual order.
}

func (i Inst) String() string {
	var buf bytes.Buffer
	//buf.WriteString(i.Op.String())
	buf.WriteString(fmt.Sprintf("%-8s",strings.ToLower(i.Op.String())))
	for j, arg := range i.Args {
		if arg == nil {
			break
		}
		if j == 0 {
			buf.WriteString(" ")
		} else {
			switch arg.(type) {
			case Index, Base:
			default:
				buf.WriteString(", ")
			}
		}
		switch arg.(type) {
			case Index,Base,Reg:
				buf.WriteString("%")
			default:
		}
		//buf.WriteString(arg.String())
		buf.WriteString(arg.String())
	}
	return buf.String()
}

// An Op is an instruction operation.
type Op uint16

func (o Op) String() string {
	if int(o) >= len(opstr) || opstr[o] == "" {
		return fmt.Sprintf("Op(%d)", int(o))
	}
	return opstr[o]
}

// An Arg is a single instruction argument, one of these types: Reg, CondReg, SpReg, Imm, PCRel, Label, or Offset.
type Arg interface {
	IsArg()
	String() string
}

// An Args holds the instruction arguments.
// If an instruction has fewer than 6 arguments,
// the final elements in the array are nil.
type Args [6]Arg

// Base Register(B)
type Base uint8

const (
	B0 Base = iota
	B1
	B2
	B3
	B4
	B5
	B6
	B7
	B8
	B9
	B10
	B11
	B12
	B13
	B14
	B15
)

func (Base) IsArg() {}
func (r Base) String() string {
	switch {
	case B0 <= r && r <= B15:
		return fmt.Sprintf("r%d)", int(r-B0))
	default:
		return fmt.Sprintf("Base(%d)", int(r))
	}
}

// Index register(X)
type Index uint8

const (
	X0 Index = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	X8
	X9
	X10
	X11
	X12
	X13
	X14
	X15
)

func (Index) IsArg() {}
func (r Index) String() string {
	switch {
	case X0 <= r && r <= X15:
		return fmt.Sprintf("r%d,", int(r-X0))
	default:
		return fmt.Sprintf("Base(%d)", int(r))
	}
}

type Disp uint64

func (Disp) IsArg() {}
func (r Disp) String() string {
	return fmt.Sprintf("%d(", int32(r | 0xfff<<20))
}

// A Reg is a single register. The zero value means R0, not the absence of a register.
// It also includes special registers.
type Reg uint16

const (
	R0 Reg = iota
	R1
	R2
	R3
	R4
	R5
	R6
	R7
	R8
	R9
	R10
	R11
	R12
	R13
	R14
	R15
	F0
	F1
	F2
	F3
	F4
	F5
	F6
	F7
	F8
	F9
	F10
	F11
	F12
	F13
	F14
	F15
	V0
	V1
	V2
	V3
	V4
	V5
	V6
	V7
	V8
	V9
	V10
	V11
	V12
	V13
	V14
	V15
	V16
	V17
	V18
	V19
	V20
	V21
	V22
	V23
	V24
	V25
	V26
	V27
	V28
	V29
	V30
	V31
	A0
	A1
	A2
	A3
	A4
	A5
	A6
	A7
	A8
	A9
	A10
	A11
	A12
	A13
	A14
	A15
	C0
	C1
	C2
	C3
	C4
	C5
	C6
	C7
	C8
	C9
	C10
	C11
	C12
	C13
	C14
	C15
)

func (Reg) IsArg() {}
func (r Reg) String() string {
	switch {
	case R0 <= r && r <= R15:
		return fmt.Sprintf("r%d", int(r-R0))
	case F0 <= r && r <= F15:
		return fmt.Sprintf("f%d", int(r-F0))
	case V0 <= r && r <= V31:
		return fmt.Sprintf("v%d", int(r-V0))
	case A0 <= r && r <= A15:
		return fmt.Sprintf("a%d", int(r-A0))
	case C0 <= r && r <= C15:
		return fmt.Sprintf("c%d", int(r-C0))
	default:
		return fmt.Sprintf("Reg(%d)", int(r))
	}
}

// Imm represents an immediate number.
type Imm uint64

func (Imm) IsArg() {}
func (i Imm) String() string {
	return fmt.Sprintf("%x", uint64(i))
}

type Sign8 int8

func (Sign8) IsArg() {}
func (i Sign8) String() string {
	//x := int16(ff<<8) | int16(i)
	return fmt.Sprintf("%d", i)
}

type Sign16 int16

func (Sign16) IsArg() {}
func (i Sign16) String() string {
	return fmt.Sprintf("%d", i)
}

type Sign32 int32

func (Sign32) IsArg() {}
func (i Sign32) String() string {
	return fmt.Sprintf("%d", i)
}
// Offset represents a memory offset immediate.
type Offset int64

func (Offset) IsArg() {}
func (o Offset) String() string {
	return fmt.Sprintf("%+d", int32(o))
}
