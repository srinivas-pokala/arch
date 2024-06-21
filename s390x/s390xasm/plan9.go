// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390xasm

import (
	"fmt"
	"strings"
	"strconv"
)

// GoSyntax returns the Go assembler syntax for the instruction.
// The syntax was originally defined by Plan 9.
// The pc is the program counter of the instruction, used for
// expanding PC-relative addresses into absolute ones.
// The symname function queries the symbol table for the program
// being disassembled. Given a target address it returns the name
// and base address of the symbol containing the target, if any;
// otherwise it returns "", 0.
// The reader text should read from the text segment using text addresses
// as offsets; it is used to display pc-relative loads as constant loads.
func GoSyntax(inst Inst, pc uint64, symname func(uint64) (string, uint64)) string {
	if symname == nil {
		symname = func(uint64) (string, uint64) { return "", 0 }
	}

	var args []string
	for _, a := range inst.Args {
		if a == nil {
			break
		}
		args = append(args, plan9Arg(&inst, pc, symname, a))
	}

	op := inst.Op.String()
	switch inst.Op {
	case BRASL:
		return "CALL " + args[1]
	case LCGR:
		return "NEG " + args[1] + args[0]
	case SRAG:
		return "SRAD " + args[0]
	case LG,LGF,LLGF,LGH, LLGH, LGB, LLGC, LDY, LEY, LRVG, LRV, LRVH:
		if args[2] != "" && args[3] != "" {
			args[2]=fmt.Sprintf("(%s, %s)",args[2], args[3])
		} else if args[2] != "" {
			args[2]=fmt.Sprintf("(%s)",args[2])
		} else if args[3] != "" {
			args[2]=fmt.Sprintf("(%s)",args[3])
		}
		switch inst.Op {
			case LG:
				op ="MOVD"
			case LGF:
				op ="MOVW"
			case LLGF:
				op ="MOVWZ"
			case LGH:
				op ="MOVH"
			case LLGH:
				op ="MOVHZ"
			case LGB:
				op ="MOVB"
			case LLGC:
				op ="MOVBZ"
			case LDY:
				op ="FMOVD"
			case LEY:
				op ="FMOVS"
			case LRVG:
				op ="MOVDBR"
			case LRV:
				op ="MOVWBR"
			case LRVH:
				op ="MOVHBR"
		}
	}

	if args != nil {
		op += " " + strings.Join(args, ", ")
	}

	return op
}


// plan9Arg formats arg (which is the argIndex's arg in inst) according to Plan 9 rules.
//
// NOTE: because Plan9Syntax is the only caller of this func, and it receives a copy
// of inst, it's ok to modify inst.Args here.
func plan9Arg(inst *Inst,  pc uint64, symname func(uint64) (string, uint64), arg Arg) string {
	switch a := arg.(type) {
	case Reg:
		if arg == R13 {
			return "g"
		}
		return strings.ToUpper(arg.String(pc)[1:])
	case Base:
		n := uint8(a)
		if n == 0 {
			return ""
		}
		return strings.ToUpper(arg.String(pc)[1:])
	case Index:
		n := uint8(a)
		if n == 0 {
			return ""
		}
		return strings.ToUpper(arg.String(pc)[1:])
	case VReg:
		return strings.ToUpper(arg.String(pc)[1:])
	case Disp20, Disp12:
		numstr := arg.String(pc)
		num, err := strconv.Atoi(numstr[:len(numstr)-1])
		if err != nil {
			return fmt.Sprintf("plan9Arg: error in converting Atoi:%s", err)
		}
		if num == 0 {
			return ""
		} else {
			return strconv.Itoa(num)
		}
	case RegIm12, RegIm16, RegIm24, RegIm32:
		addr, err := strconv.ParseUint(arg.String(pc)[2:], 16, 64)
		if err != nil {
			return fmt.Sprintf("plan9Arg: error in converting ParseUint:%s", err)
		}
		s, base := symname(addr)
		if s != "" && addr == base {
			return fmt.Sprintf("%s(SB)", s)
		}
		return fmt.Sprintf("%#x", addr)
	case Imm, Sign8, Sign16, Sign32:
		numImm := arg.String(pc)
		return fmt.Sprintf("$%s", numImm)
	case Mask, Len:
		num := arg.String(pc)
		return fmt.Sprintf("$%s", num)
	}
	return fmt.Sprintf("???(%v)", arg)
}

// Convert a general-purpose register to plan9 assembly format.
func plan9gpr(r Reg) string {
	regno := uint16(r) & 31
	if regno == 31 {
		return "ZR"
	}
	return fmt.Sprintf("R%d", regno)
}
