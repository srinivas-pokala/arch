// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390xasm

import (
	"fmt"
	"strconv"
	"strings"
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

var vectorSize = map[int]string{0: "B", 1: "H", 2: "F", 3: "G", 4: "Q"}
var vectorCS = map[int]string{0: "BS", 1: "HS", 2: "FS", 3: "GS"}

func GoSyntax(inst Inst, pc uint64, symname func(uint64) (string, uint64)) string {
	if symname == nil {
		symname = func(uint64) (string, uint64) { return "", 0 }
	}

	var args []string
	opString := inst.Op.String()
	op := strings.ToUpper(opString)
	//for i, a := range inst.Args {
	for i := 0; i < len(inst.Args); i++ {
		if inst.Args[i] == nil {
			break
		}
		switch inst.Args[i].(type) {
		case Disp12, Disp20:
			var temp []string
			if _, ok := inst.Args[i+1].(Index); ok { //D(X,B)
				for j := 0; j < 3; j++ {
					temp = append(temp, plan9Arg(&inst, pc, symname, inst.Args[i+j]))
				}
				args = append(args, mem_operandx(temp))
				i = i + 2
			} else { //D(B)
				for j := 0; j < 2; j++ {
					temp = append(temp, plan9Arg(&inst, pc, symname, inst.Args[i+j]))
				}
				args = append(args, mem_operand(temp))
				i = i + 1
			}
		default:
			args = append(args, plan9Arg(&inst, pc, symname, inst.Args[i]))
		}
	}
	if strings.HasPrefix(op, "V") || strings.Contains(op, "WFC") || strings.Contains(op, "WFK") {
		args = args[:len(args)-1]
	}

	switch inst.Op {
	default:
		switch len(args) {
		case 0:
			return op
		case 1:
			return fmt.Sprintf("%s %s", op, args[0])
		case 2:
			if reverseOperandOrder(inst.Op) {
				args[0], args[1] = args[1], args[0]
			}
			op += " " + strings.Join(args, ", ")
			return op
		case 3:
			if reverseOperandOrder(inst.Op) {
				args[0], args[2] = args[2], args[0]
			} else if reverseOperand3(inst.Op) {
				args[0], args[1], args[2] = args[2], args[0], args[1]
			}
			op += " " + strings.Join(args, ", ")
			return op
		case 4:
			if reverseOperandOrder(inst.Op) {
				args[0], args[1], args[2], args[3] = args[1], args[2], args[3], args[0]
			} else if reverseEndOperand(inst.Op) {
				args[0], args[3] = args[3], args[0]
			}
		}
		return op + " " + strings.Join(args, ",")
	case LCGR, LCGFR:
		switch inst.Op {
		case LCGR:
			op = "NEG"
		case LCGFR:
			op = "NEGW"
		}
		if args[0] == args[1] {
			args = args[:1]
		} else {
			args[0], args[1] = args[1], args[0]
		}
		return op + " " + strings.Join(args, ", ")
	case LD, LE, LG, LGF, LLGF, LGH, LLGH, LGB, LLGC, LDY, LEY, LRVG, LRV, LRVH:
		//args[1] = mem_operandx(args[1:4]) //D(X,B)
		//args = args[:2]
		args[0], args[1] = args[1], args[0]
		switch inst.Op {
		case LG:
			op = "MOVD"
		case LGF:
			op = "MOVW"
		case LLGF:
			op = "MOVWZ"
		case LGH:
			op = "MOVH"
		case LLGH:
			op = "MOVHZ"
		case LGB:
			op = "MOVB"
		case LLGC:
			op = "MOVBZ"
		case LDY, LD:
			op = "FMOVD"
		case LEY, LE:
			op = "FMOVS"
		case LRVG:
			op = "MOVDBR"
		case LRV:
			op = "MOVWBR"
		case LRVH:
			op = "MOVHBR"
		}
	case LA, LAY:
		//args[1] = mem_operandx(args[1:4]) //D(X,B)
		args[0], args[1] = args[1], args[0]
		//args = args[:2]
		op = "MOVD"
		return op + " " + strings.Join(args, ", ")

	case LAA, LAAG, LAAL, LAALG, LAN, LANG, LAX, LAXG, LAO, LAOG:
		//args[2] = mem_operand(args[2:4]) //D(B)
		args[0], args[1] = args[1], args[0]
		//args = args[:3]
		return op + " " + strings.Join(args, ", ")
	case LM, LMY, LMG: // Load Multiple
		switch inst.Op {
		case LM, LMY:
			op = "LMY"
		}
		//args[2] = mem_operand(args[2:4]) //D(B)
		args[0], args[1], args[2] = args[2], args[0], args[1]
		//args = args[:3]
		return op + " " + strings.Join(args, ", ")

	case TCEB, TCDB: // Test Data class
		//args[1] = mem_operandx(args[1:4]) //D(X,B)
		//args = args[:2]
		return op + " " + strings.Join(args, ", ")
	case STM, STMY, STMG: // Store Multiple
		switch inst.Op {
		case STM, STMY:
			op = "STMY"
		}
		//args[2] = mem_operand(args[2:4]) //D(B)
		//args = args[:3]
		return op + " " + strings.Join(args, ", ")
	case ST, STY, STG:
		switch inst.Op {
		case ST, STY:
			op = "MOVW"
		case STG:
			op = "MOVD"
		}
		//args[1] = mem_operandx(args[1:4]) //D(X,B)
		//args = args[:2]
		return op + " " + strings.Join(args, ", ")
	case LGR, LGFR, LGHR, LGBR, LLGFR, LLGHR, LLGCR, LRVGR, LRVR, LDR:
		switch inst.Op {
		case LGR:
			op = "MOVD"
		case LGFR:
			op = "MOVW"
		case LGHR:
			op = "MOVH"
		case LGBR:
			op = "MOVB"
		case LLGFR:
			op = "MOVWZ"
		case LLGHR:
			op = "MOVHZ"
		case LLGCR:
			op = "MOVBZ"
		case LRVGR:
			op = "MOVDBR"
		case LRVR:
			op = "MOVWBR"
		case LDR:
			op = "FMOVD"
		}
		args[0], args[1] = args[1], args[0]
	case LZDR:
		op = "FMOVD"
		return op + " " + "$0" + ", " + args[0]
	case LZER:
		op = "FMOVS"
		return op + " " + "$0" + ", " + args[0]
	case STD, STDY, STE, STEY:
		switch inst.Op {
		case STD, STDY:
			op = "FMOVD"
		case STE, STEY:
			op = "FMOVS"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
		return op + " " + strings.Join(args, ", ")

	case LGHI, LLILH, LLIHL, LLIHH, LGFI, LLILF, LLIHF:
		op = "MOVD"
		args[0], args[1] = args[1], args[0]
	case ARK, AGRK, ALGRK:
		switch inst.Op {
		case ARK:
			op = "ADDW"
		case AGRK:
			op = "ADD"
		case ALGRK:
			op = "ADDC"
		}
		if args[0] == args[1] {
			args[0], args[1] = args[2], args[0]
			args = args[:2]
		} else {
			args[0], args[1], args[2] = args[2], args[1], args[0]
		}
		return op + " " + strings.Join(args, ", ")
	case AGHIK, AHIK, ALGHSIK:
		switch inst.Op {
		case AGHIK:
			op = "ADD"
		case AHIK:
			op = "ADDW"
		case ALGHSIK:
			op = "ADDC"
		}
		args[0], args[1], args[2] = args[2], args[1], args[0]
		return op + " " + strings.Join(args, ", ")
	case AGHI, AHI, AGFI, AFI, AR, ALCGR:
		switch inst.Op {
		case AGHI, AGFI:
			op = "ADD"
		case AHI, AFI, AR:
			op = "ADDW"
		case ALCGR:
			op = "ADDE"
		}
		args[0], args[1] = args[1], args[0]
		return op + " " + strings.Join(args, ", ")
	case AEBR, ADBR, DDBR, DEBR, MDBR, MEEBR, SDBR, SEBR, LPDBR, LNDBR, LPDFR, LNDFR, LCDFR, LCEBR, LEDBR, LDEBR, SQDBR, SQEBR:
		switch inst.Op {
		case AEBR:
			op = "FADDS"
		case ADBR:
			op = "FADD"
		case DDBR:
			op = "FDIV"
		case DEBR:
			op = "FDIVS"
		case MDBR:
			op = "FMUL"
		case MEEBR:
			op = "FMULS"
		case SDBR:
			op = "FSUB"
		case SEBR:
			op = "FSUBS"
		case LPDBR:
			op = "FABS"
		case LNDBR:
			op = "FNABS"
		case LCDFR:
			op = "FNEG"
		case LCEBR:
			op = "FNEGS"
		case SQDBR:
			op = "FSQRT"
		case SQEBR:
			op = "FSQRTS"
		}
		args[0], args[1] = args[1], args[0]
		return op + " " + strings.Join(args, ", ")
	case SR, SGR, SLGR, SLFI:
		switch inst.Op {
		case SR, SLFI:
			op = "SUBW"
		case SGR:
			op = "SUB"
		case SLGR:
			op = "SUBC"
		}
		args[0], args[1] = args[1], args[0]
		return op + " " + strings.Join(args, ", ")
	case SGRK, SLGRK, SRK:
		switch inst.Op {
		case SGRK:
			op = "SUB"
		case SLGRK:
			op = "SUBC"
		case SRK:
			op = "SUBW"
		}
		if args[0] == args[1] {
			args[0], args[1] = args[2], args[0]
			args = args[:2]
		} else {
			args[0], args[1], args[2] = args[2], args[1], args[0]
		}
		return op + " " + strings.Join(args, ", ")
	case SLBGR:
		op = "SUBE"
		args[0], args[1] = args[1], args[0]
		return op + " " + strings.Join(args, ", ")
	case MSGFR, MHI, MSFI, MSGFI:
		switch inst.Op {
		case MSGFR, MHI, MSFI:
			op = "MULLW"
		case MSGFI:
			op = "MULLD"
		}
		args[0], args[1] = args[1], args[0]
		return op + " " + strings.Join(args, ", ")

	case NGR, NR, NILL, NILF, NILH, OGR, OR, OILL, OILF, OILH, XGR, XR, XILF:
		op = bitwise_op(inst.Op)
		args[0], args[1] = args[1], args[0]
		return op + " " + strings.Join(args, ", ")

	case NGRK, NRK, OGRK, ORK, XGRK, XRK: // opcode R1, R2, R3
		op = bitwise_op(inst.Op)
		args[0], args[1], args[2] = args[1], args[2], args[0]
		return op + " " + strings.Join(args, ", ")
	case SLLG, SRLG, SLLK, SRLK, RLL, RLLG, SRAK, SRAG:
		switch inst.Op {
		case SLLG:
			op = "SLD"
		case SRLG:
			op = "SRD"
		case SLLK:
			op = "SLW"
		case SRLK:
			op = "SRW"
		case SRAK:
			op = "SRAW"
		case SRAG:
			op = "SRAD"
		}
		//args[2] = mem_operand(args[2:])
		//args = args[:3]
		args[0], args[1], args[2] = args[2], args[1], args[0]
		return op + " " + strings.Join(args, ", ")
	case TRAP2, SVC:
		op = "SYSALL"
	case CR, CLR, CGR, CLGR, KDBR, CDBR, CEBR, CGHI, CHI, CGFI, CLGFI, CFI, CLFI:
		switch inst.Op {
		case CGHI, CGFI, CGR:
			op = "CMP"
		case CHI, CFI, CR:
			op = "CMPW"
		case CLGFI, CLGR:
			op = "CMPU"
		case CLFI, CLR:
			op = "CMPWU"
		case CDBR:
			op = "FCMPU"
		case KDBR:
			op = "FCMPO"
		}
		return op + " " + strings.Join(args, ", ")
	case CEFBRA, CDFBRA, CEGBRA, CDGBRA, CELFBR, CDLFBR, CELGBR, CDLGBR:
		return op + " " + args[2] + ", " + args[0]
	case CFEBRA, CFDBRA, CGEBRA, CGDBRA, CLFEBR, CLFDBR, CLGEBR, CLGDBR:
		return op + " " + args[2] + ", " + args[0]
	case CGRJ, CGIJ:
		mask, err := strconv.Atoi(args[2][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		var check bool
		switch mask & 0xf {
		case 2:
			op = "CMPBGT"
			check = true
		case 4:
			op = "CMPBLT"
			check = true
		case 7:
			op = "CMPBNE"
			check = true
		case 8:
			op = "CMPBEQ"
			check = true
		case 10:
			op = "CMPBGE"
			check = true
		case 12:
			op = "CMPBLE"
			check = true
		}
		if check {
			args[2] = args[3]
			args = args[:3]
		}
		return op + " " + strings.Join(args, ", ")
	case CLGRJ, CLGIJ:
		mask, err := strconv.Atoi(args[2][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		var check bool
		switch mask & 0xf {
		case 2:
			op = "CMPUBGT"
			check = true
		case 4:
			op = "CMPUBLT"
			check = true
		case 7:
			op = "CMPUBNE"
			check = true
		case 8:
			op = "CMPUBEQ"
			check = true
		case 10:
			op = "CMPUBGE"
			check = true
		case 12:
			op = "CMPUBLE"
			check = true
		}
		if check {
			args[2] = args[3]
			args = args[:3]
		}
		return op + " " + strings.Join(args, ", ")
	case CLRJ, CRJ, CIJ, CLIJ:
		args[0], args[1], args[2], args[3] = args[2], args[0], args[1], args[3]
		return op + " " + strings.Join(args, ", ")
	case BRC, BRCL:
		mask, err := strconv.Atoi(args[0][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		opStr, check := branch_relative_op(mask, inst.Op)
		if opStr != "" {
			op = opStr
		}
		if check {
			return op + " " + args[1]
		} else {
			return op + " " + strings.Join(args, ", ")
		}
	case BCR:
		mask, err := strconv.Atoi(args[0][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		opStr, check := branchOnConditionOp(mask, inst.Op)
		if opStr != "" {
			op = opStr
		}
		if op == "SYNC" || op == "NOPH" {
			args = args[:0]
			return op
		}
		if check {
			return op + " " + args[1]
		} else {
			return op + " " + strings.Join(args, ", ")
		}
	case LOCGR:
		mask, err := strconv.Atoi(args[2][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		var check bool
		switch mask & 0xf {
		case 2: //Greaterthan (M=2)
			op = "MOVDGT"
			check = true
		case 4: //Lessthan (M=4)
			op = "MOVDLT"
			check = true
		case 7: // Not Equal (M=7)
			op = "MOVDNE"
			check = true
		case 8: // Equal (M=8)
			op = "MODEQ"
			check = true
		case 10: // Greaterthan or Equal (M=10)
			op = "MOVDGE"
			check = true
		case 12: // Lessthan or Equal (M=12)
			op = "MOVDLE"
			check = true
		}
		if check {
			args[0], args[1] = args[1], args[0]
			args = args[:2]
		} else {
			args[0], args[1], args[2] = args[2], args[1], args[0]
		}

		return op + " " + strings.Join(args, ", ")
	case BRASL:
		op = "CALL" // BL
		return op + " " + args[1]
	case X, XY, XG:
		switch inst.Op {
		case X, XY:
			op = "XORW"
		case XG:
			op = "XOR"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
	case XC, NC, OC, MVC, MVCIN, CLC: //D(L,B)
		args[0], args[1] = args[1], args[0]
		//args[1] = mem_operand(args[1:3])
		//args[2] = mem_operand(args[3:])
		args[1], args[2] = args[2], args[1]
		//args = args[:3]
		return op + " " + strings.Join(args, ", ")
	case O, OY, OG:
		switch inst.Op {
		case O, OY:
			op = "ORW"
		case OG:
			op = "OR"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
	case N, NY, NG:
		switch inst.Op {
		case N, NY:
			op = "ANDW"
		case NG:
			op = "AND"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
	case S, SY, SLBG, SLG, SG:
		switch inst.Op {
		case S, SY:
			op = "SUBW"
		case SLBG:
			op = "SUBE"
		case SLG:
			op = "SUBC"
		case SG:
			op = "SUB"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
	case MSG, MSY, MS:
		switch inst.Op {
		case MSG:
			op = "MULLD"
		case MSY, MS:
			op = "MULLW"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
	case A, AY, ALCG, ALG, AG:
		switch inst.Op {
		case A, AY:
			op = "ADDW"
		case ALCG:
			op = "ADDE"
		case ALG:
			op = "ADDC"
		case AG:
			op = "ADD"
		}
		//args[1] = mem_operandx(args[1:])
		//args = args[:2]
	case RISBG, RISBGN, RISBHG, RISBLG, RNSBG, RXSBG, ROSBG:
		switch inst.Op {
		case RNSBG, RXSBG, ROSBG:
			num, err := strconv.Atoi(args[2][1:])
			if err != nil {
				return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
			}
			if ((num >> 7) & 0x1) != 0 {
				op = op + "T"
			}
		case RISBG, RISBGN, RISBHG, RISBLG:
			num, err := strconv.Atoi(args[3][1:])
			if err != nil {
				return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
			}
			if ((num >> 7) & 0x1) != 0 {
				op = op + "Z"
			}
		}
		if len(args) == 5 {
			args[0], args[1], args[2], args[3], args[4] = args[2], args[3], args[4], args[1], args[0]
		} else {
			args[0], args[1], args[2], args[3] = args[2], args[3], args[1], args[0]
		}
		return op + " " + strings.Join(args, ", ")

	case VEC, VECL, VCLZ, VCTZ, VREPI, VPOPCT: //mnemonic V1, V2, M3
		mask, err := strconv.Atoi(args[2][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi for %q:%s", op, err)
		}
		val := mask & 0x7
		if val >= 0 && val < 4 {
			op = op + vectorSize[val]
			args = args[:2]
		} else {
			return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
		}
		switch inst.Op {
		case VCLZ, VCTZ, VREPI, VPOPCT:
			args[0], args[1] = args[1], args[0]
		default:
		}
		//Mnemonic V1, V2, V3, M4 or Mnemonic V1, I2, I3, M4 or Mnemonic V1, V3, I2, M4
	case VA, VS, VACC, VAVG, VAVGL, VMX, VMXL, VMN, VMNL, VGFM, VGM, VREP, VERLLV, VESLV, VSCBI, VSUM, VSUMG, VSUMQ:
		mask, err := strconv.Atoi(args[3][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		val := mask & 0x7
		switch inst.Op {
		case VA, VS, VACC, VSCBI:
			if val >= 0 && val < 5 {
				op = op + vectorSize[val]
				args[0], args[1], args[2] = args[1], args[2], args[0]
				args = args[:3]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
			}
		case VAVG, VAVGL, VMX, VMXL, VMN, VMNL, VGFM, VGM:
			if val >= 0 && val < 4 {
				op = op + vectorSize[val]
				args[0], args[1], args[2] = args[1], args[2], args[0]
				args = args[:3]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
			}
		case VREP, VERLLV, VESLV:
			if val >= 0 && val < 4 {
				op = op + vectorSize[val]
				args[0], args[1], args[2] = args[2], args[1], args[0]
				args = args[:3]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
			}
		case VSUM, VSUMG, VSUMQ:
			if val >= 0 && val < 2 {
				op = op + vectorSize[val]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
			}
		case VSUMG:
			if val > 0 && val < 3 {
				op = op + vectorSize[val]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
			}
		case VSUMQ:
			if val > 1 && val < 4 {
				op = op + vectorSize[val]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
			}
		}

	case VGFMA, VERIM: // Mnemonic V1, V2, V3, V4/I4, M5
		mask, err := strconv.Atoi(args[4][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		val := mask & 0x7
		args = args[:4]
		if val >= 0 && val < 4 {
			op = op + vectorSize[val]
		} else {
			return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
		}
		switch inst.Op {
		case VGFMA:
			args[0], args[1], args[2], args[3] = args[1], args[2], args[3], args[0]
		default:
			args[0], args[3] = args[3], args[0]
		}
	case VSTRC, VFAE, VFEE, VFENE:
		var off uint8
		switch inst.Op {
		case VSTRC:
			off = uint8(1)
		default:
			off = uint8(0)
		}
		m1, err := strconv.Atoi(args[3+off][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		m2, err := strconv.Atoi(args[4+off][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		index := m1 & 0x3
		if index < 0 || index > 2 {
			return fmt.Sprintf("Specefication exception is recognized for %q with mask values: %v, %v \n", op, m1, m2)
		}
		switch m2 {
		case 0:
			op = op + vectorSize[index]
		case 1:
			op = op + vectorCS[index]
		case 2:
			op = op + "Z" + vectorSize[index]
		case 3:
			op = op + "Z" + vectorCS[index]
		default:
			return fmt.Sprintf("Specefication exception is recognized for %q with mask values: %v, %v \n", op, m1, m2)
		}
		switch inst.Op {
		case VSTRC:
			args[0], args[1], args[2], args[3] = args[1], args[2], args[3], args[0]
		default:
			args[0], args[1], args[2] = args[1], args[2], args[0]
		}
		args = args[:3+off]

	case VCEQ, VCH, VCHL: // Mnemonic V1, V2, V3, M4, M5
		m4, err := strconv.Atoi(args[3][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		m5, err := strconv.Atoi(args[4][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		val := (m4 & 0x7)
		if m5 == 0 {
			if val >= 0 && val < 4 {
				op = op + vectorSize[val]
				args[0], args[1], args[2] = args[1], args[2], args[0]
				args = args[:3]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask(m4) value: %v \n", op, m4)
			}
		} else if m5 == 1 {
			if val >= 0 && val < 4 {
				op = op + vectorCS[val]
				args[0], args[1], args[2] = args[1], args[2], args[0]
				args = args[:3]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask(m4) value: %v \n", op, m4)
			}
		} else {
			return fmt.Sprintf("Specefication exception is recognized for %q with mask(m5) value: %v \n", op, m5)
		}
	case VFMA, VFMS, VMSL: //Mnemonic V1, V2, V3, V4, M5, M6
		m5, err := strconv.Atoi(args[4][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		m6, err := strconv.Atoi(args[5][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		switch inst.Op {
		case VMSL:
			if m5 == 3 && m6 == 8 {
				op = op + "EG"
			} else if m5 == 3 && m6 == 4 {
				op = op + "OG"
			} else if m5 == 3 && m6 == 12 {
				op = op + "EOG"
			} else if m5 == 3 {
				op = op + "G"
			}
		default:
			if m5 == 0 && m6 == 3 {
				op = op + "DB"
			} else if m5 == 8 && m6 == 3 {
				op = "W" + op[1:] + "DB"
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with m5: %v m6: %v \n", op, m5, m6)
			}
		}
		args[0], args[1], args[2], args[3] = args[1], args[2], args[3], args[0]
		args = args[:4]

	case VFCE, VFCH, VFCHE: //Mnemonic V1,V2,V3,M4,M5,M6
		m4, err := strconv.Atoi(args[3][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		m5, err := strconv.Atoi(args[4][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		m6, err := strconv.Atoi(args[5][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		if m5 == 0 {
			if m4 == 3 && m6 == 0 {
				op = op + "DB"
			} else if m4 == 3 && m6 == 1 {
				op = op + "DBS"
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with m4: %v, m6: %v \n", op, m4, m6)
			}

		} else if m5 == 8 {
			if m4 == 3 && m6 == 0 {
				op = "W" + op[1:] + "DB"
			} else if m4 == 3 && m6 == 1 {
				op = "W" + op[1:] + "DBS"
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with m4: %v, m6: %v \n", op, m4, m6)
			}
		} else {
			return fmt.Sprintf("Specefication exception is recognized for %q with m5: %v \n", op, m5)
		}
		args[0], args[1], args[2] = args[1], args[2], args[0]
		args = args[:3]

	case VFTCI:
		m4, err := strconv.Atoi(args[3][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		m5, err := strconv.Atoi(args[4][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: %q error in converting Atoi:%s", op, err)
		}
		val := (m4 & 0x7)
		if m5 == 0 {
			switch val {
			case 2:
				op = op + "SB"
			case 3:
				op = op + "DB"
			default:
				return fmt.Sprintf("Specefication exception is recognized for %q with mask(m4) value: %v \n", op, m4)
			}
		} else if m5 == 8 {
			switch val {
			case 2:
				op = "W" + op[1:] + "SB"
			case 3:
				op = "W" + op[1:] + "DB"
			case 4:
				op = "W" + op[1:] + "XB"
			default:
				return fmt.Sprintf("Specefication exception is recognized for %q with mask(m4) value: %v \n", op, m4)
			}
		} else {
			return fmt.Sprintf("Specefication exception is recognized for %q with mask(m5) value: %v \n", op, m5)
		}
		args[0], args[1], args[2] = args[2], args[1], args[0]
		args = args[:3]
	case VAC, VACCC:
		mask, err := strconv.Atoi(args[4][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		if mask&0x04 == 0 {
			return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, mask)
		}
		op = op + "Q"
		args[0], args[1], args[2], args[3] = args[1], args[2], args[3], args[0]
		args = args[:5]
	case VL, VLREP:
		//args[1] = mem_operandx(args[1:4]) // D(X,B)
		switch inst.Op {
		case VL:
			args[0], args[1] = args[1], args[0]
			args = args[:2]
		case VLREP:
			args[0], args[1] = args[1], args[0]
			mask, err := strconv.Atoi(args[2][1:])
			if err != nil {
				return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
			}
			if mask >= 0 && mask < 4 {
				op = op + vectorSize[mask]
			}
		}
	case VST, VSTEB, VSTEH, VSTEF, VSTEG, VLEB, VLEH, VLEF, VLEG: //Mnemonic V1, D2(X2,B2), M3
		//args[1] = mem_operandx(args[1:4]) // D(X,B)
		m, err := strconv.Atoi(args[2][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		//args[2] = args[4]
		//args = args[:3]
		switch inst.Op {
		case VST:
			if m == 0 || (m > 2 && m < 5) {
				args = args[:2]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, m)
			}
		case VLEB, VLEH, VLEF, VLEG:
			args[0], args[2] = args[2], args[0]
		default:
			args[0], args[1], args[2] = args[2], args[0], args[1]
		}
	case VSTM, VSTL, VESL, VESRA, VLM, VERLL: //Mnemonic V1, V3, D2(B2)[,M4] or V1, R3,D2(B2)
		//args[2] = mem_operand(args[2:4]) // D(B)
		switch inst.Op {
		case VSTM, VLM:
			m, err := strconv.Atoi(args[3][1:])
			if err != nil {
				return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
			}
			if !(m == 0 || (m > 2 && m < 5)) {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, m)
			}
			if inst.Op == VLM {
				args[0], args[1], args[2] = args[2], args[0], args[1]
			}
			args = args[:3]
		case VESL, VESRA, VERLL:
			m, err := strconv.Atoi(args[3][1:])
			if err != nil {
				return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
			}
			if m >= 0 && m < 4 {
				op = op + vectorSize[m]
			} else {
				return fmt.Sprintf("Specefication exception is recognized for %q with mask value: %v \n", op, m)
			}
			if args[0] == args[1] {
				args[0], args[1] = args[2], args[1]
				args = args[:2]
				break
			}
			args[0], args[1], args[2] = args[2], args[1], args[0]
			args = args[:3]
		case VSTL:
			args[0], args[1] = args[1], args[0]
			args = args[:3]
		}
	case VGBM:
		val, err := strconv.Atoi(args[1][1:])
		if err != nil {
			return fmt.Sprintf("GoSyntax: error in converting Atoi:%s", err)
		}
		if val == 0 {
			op = "VZERO"
			args = args[:1]
		} else if val == 0xffff {
			op = "VONE"
			args = args[:1]
		} else {
			args[0], args[1] = args[1], args[0]
			args = args[:2]
		}
	case VN, VNC, VO, VX, VNO: //mnemonic V1, V2, V3
		if args[0] == args[2] {
			args = args[:2]
			args[0], args[1] = args[1], args[0]
		} else {
			args[0], args[1], args[2] = args[1], args[2], args[0]
		}
		if op == "VNO" {
			op = op + "T"
		}

	}
	if args != nil {
		op += " " + strings.Join(args, ", ")
	}

	return op
}

func branch_relative_op(mask int, opconst Op) (op string, check bool) {
	switch mask & 0xf {
	case 2:
		op = "BGT"
		check = true
	case 4:
		op = "BLT"
		check = true
	case 5:
		op = "BLTU"
		check = true
	case 7:
		op = "BNE"
		check = true
	case 8:
		op = "BEQ"
		check = true
	case 10:
		op = "BGE"
		check = true
	case 12:
		op = "BLE"
		check = true
	case 13:
		op = "BLEU"
		check = true
	case 15:
		op = "JMP" // BR
		check = true
	}
	return op, check
}

func branchOnConditionOp(mask int, opconst Op) (op string, check bool) {
	switch mask & 0xf {
	case 0:
		op = "NOPH"
	case 14:
		op = "SYNC"
	case 15:
		op = "JMP"
		check = true
	}
	return op, check
}

func bitwise_op(op Op) string {
	var ret string
	switch op {
	case NGR, NGRK, NILL, NILF:
		ret = "AND"
	case NR, NRK, NILH:
		ret = "ANDW"
	case OGR, OGRK, OILL, OILF:
		ret = "OR"
	case OR, ORK, OILH:
		ret = "ORW"
	case XGR, XGRK, XILF:
		ret = "XOR"
	case XR, XRK:
		ret = "XORW"
	}
	return ret
}

func mem_operand(args []string) string { //D(B)
	if args[0] != "" && args[1] != "" {
		args[0] = fmt.Sprintf("%s(%s)", args[0], args[1])
	} else if args[0] != "" {
		args[0] = fmt.Sprintf("$%s", args[0])
	} else if args[1] != "" {
		args[0] = fmt.Sprintf("(%s)", args[1])
	} else {
		args[0] = ""
	}
	return args[0]
}

func mem_operandx(args []string) string { //D(X,B)
	if args[1] != "" && args[2] != "" {
		args[1] = fmt.Sprintf("(%s, %s)", args[1], args[2])
	} else if args[1] != "" {
		args[1] = fmt.Sprintf("(%s)", args[1])
	} else if args[2] != "" {
		args[1] = fmt.Sprintf("(%s)", args[2])
	} else if args[0] != "" {
		args[1] = ""
	}
	if args[0] != "" && args[1] != "" {
		args[0] = fmt.Sprintf("%s%s", args[0], args[1])
	} else if args[0] != "" {
		args[0] = fmt.Sprintf("$%s", args[0])
	} else if args[1] != "" {
		args[0] = fmt.Sprintf("%s", args[1])
	} else {
		args[0] = ""
	}
	return args[0]
}

// plan9Arg formats arg (which is the argIndex's arg in inst) according to Plan 9 rules.
//
// NOTE: because Plan9Syntax is the only caller of this func, and it receives a copy
// of inst, it's ok to modify inst.Args here.
func plan9Arg(inst *Inst, pc uint64, symname func(uint64) (string, uint64), arg Arg) string {
	switch arg.(type) {
	case Reg:
		if arg == R13 {
			return "g"
		}
		return strings.ToUpper(arg.String(pc)[1:])
	case Base:
		if arg == R13 {
			return "g"
		}
		s := arg.String(pc)
		if s != "" {
			return strings.ToUpper(s[1 : len(s)-1])
		}
		return "R0"
	case Index:
		if arg == R13 {
			return "g"
		}
		s := arg.String(pc)
		if s != "" {
			return strings.ToUpper(s[1:])
		}
		return ""
	case VReg:
		return strings.ToUpper(arg.String(pc)[1:])
	case Disp20, Disp12:
		numstr := arg.String(pc)
		num, err := strconv.Atoi(numstr[:len(numstr)])
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
		off := int32(addr - pc)
		s, base := symname(addr)
		if s != "" && addr == base {
			return fmt.Sprintf("%s(SB)", s)
		}
		off = off / 4
		return fmt.Sprintf("%v(PC)", off)
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

func reverseOperandOrder(op Op) bool {
	switch op {
	case LOCR:
		return true
	case LTEBR, LTDBR:
		return true
	case VLEIB, VLEIH, VLEIF, VLEIG, VSEL, VPERM:
		return true
	}
	return false
}
func reverseOperand3(op Op) bool {
	switch op {
	case VLVGP:
		return true
	}
	return false
}

func reverseEndOperand(op Op) bool {
	switch op {
	case VPDI:
		return true
	}
	return false
}
