// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package s390xasm

import (
	"encoding/binary"
	"fmt"
	"log"
)

const debugDecode = false

const prefixOpcode = 1

// instFormat is a decoding rule for one specific instruction form.
// an instruction ins matches the rule if ins&Mask == Value
// DontCare bits should be zero, but the machine might not reject
// ones in those bits, they are mainly reserved for future expansion
// of the instruction set.
// The Args are stored in the same order as the instruction manual.
//
// Prefixed instructions are stored as:
//
//	prefix << 32 | suffix,
//
// Regular instructions are:
//
//	inst << 32
type instFormat struct {
	Op       Op
	Mask     uint64
	Value    uint64
//	DontCare uint64
	Args     [6]*argField
}

// argField indicate how to decode an argument to an instruction.
// First parse the value from the BitFields, shift it left by Shift
// bits to get the actual numerical value.
type argField struct {
	Type  ArgType
	BitField
}

// Parse parses the Arg out from the given binary instruction i.
func (a argField) Parse(i uint64) Arg {
	switch a.Type {
	default:
		return nil
	case TypeUnknown:
		return nil
	case TypeReg, TypeBaseReg, TypeIndexReg:
		return R0 + Reg(a.BitField.Parse(i))
	case TypeFPReg:
		return F0 + Reg(a.BitField.Parse(i))
	case TypeCReg:
		return C0 + Reg(a.BitField.Parse(i))
	case TypeACReg:
		return A0 + Reg(a.BitField.Parse(i))
	case TypeBaseReg:
		return B0 + Base(a.BitField.Parse(i))
	case TypeIndexReg:
		return X0 + Index(a.BitField.Parse(i))
	case TypeDisp:
		return Disp(a.BitField.Parse(i))
	case TypeVecReg:
		return V0 + Reg(a.BitField.Parse(i))
	case TypeVecpReg:
		return V0 + Reg(a.BitField.Parse(i))
	case TypeImmSigned:
		return Imm(a.BitField.ParseSigned(i))
	case TypeImmUnsigned, TypeLen:
		return Imm(a.BitField.Parse(i))
	case TypeOffset:
		return Offset(a.BitField.ParseSigned(i))
	}
}

type ArgType int8

const (
	TypeUnknown      ArgType = iota
	TypeReg                  // integer register
	TypeFPReg                // floating point register
	TypeACReg                // access register
        TypeCReg                 // control register
        TypeVecReg               // vector register
        TypeVecpReg              // vector register
        TypeImmSigned            // signed immediate
        TypeImmUnsigned          // unsigned immediate/flag/mask, this is the catch-all type
	TypeBaseReg		// Base Register for accessing memory
	TypeIndexReg		// Index Register
	TypeDisp		// Displacement for memory address
	TypeLen			// Length
        TypeOffset               // signed offset in load/store
        TypeNegOffset            // A negative 16 bit value 0b1111111xxxxx000 encoded as 0bxxxxx (e.g in the hashchk instruction)

)

func (t ArgType) String() string {
	switch t {
	default:
		return fmt.Sprintf("ArgType(%d)", int(t))
	case TypeUnknown:
		return "Unknown"
	case TypeReg:
		return "Reg"
	case TypeFPReg:
		return "FPReg"
	case TypeACReg:
		return "ACReg"
	case TypeCReg:
		return "CReg"
	case TypeDisp:
		return "Disp"
	case TypeBaseReg:
		return "BaseReg"
	case TypeIndexReg:
		return "IndexReg"
	case TypeLen:
		return "Len"
	case TypeVecReg:
		return "VecReg"
	case TypeVecSpReg:
		return "VecSpReg"
	case TypeImmSigned:
		return "ImmSigned"
	case TypeImmUnsigned:
		return "ImmUnsigned"
	case TypeOffset:
		return "Offset"
	case TypeNegOffset:
		return "NegOffset"
	}
}

func (t ArgType) GoString() string {
	s := t.String()
	if t > 0 && t < TypeLast {
		return "Type" + s
	}
	return s
}

var (
	// Errors
	errShort   = fmt.Errorf("truncated instruction")
	errUnknown = fmt.Errorf("unknown instruction")
)

var decoderCover []bool

// Decode decodes the leading bytes in src as a single instruction using
// byte order ord.
func Decode(src []byte ) (inst Inst, err error) {
	if len(src) < 2 {
		return inst, errShort
	}
	if decoderCover == nil {
		decoderCover = make([]bool, len(instFormats))
	}
	bit_check :=  binary.BigEndian.Uint8(src[0:1]) >> 6
	l := uint8(0)
	if bit_check & uint8(0x03) == 0 {
		l = 2
	} else if (bit_check & uint8(0x01) || bit_check & uint8(0x02) ) {
		l = 4
	} else if bit_check & uint8(0x03) {
		l = 6
	}
	inst.Len = l
	switch l {
	case 2:
		ui_extn := uint64(binary.BigEndian.Uint16(src[:inst.Len]))
	case 4:
		ui_extn := uint64(binary.BigEndian.Uint32(src[:inst.Len]))
	case 6:
		ui_extn := uint64(binary.BigEndian.Uint48(src[:inst.Len]))
	}
	inst.Enc = ui_extn
	ui_extn =  ui_extn << (8-l)*8
	for i, iform := range instFormats {
		if ui_extn&iform.Mask != iform.Value {
			continue
		}
		/*if ui&iform.DontCare != 0 {
			if debugDecode {
				log.Printf("Decode(%#x): unused bit is 1 for Op %s", ui, iform.Op)
			}
			// to match GNU objdump (libopcodes), we ignore don't care bits
		}*/
		for i, argfield := range iform.Args {
			if argfield == nil {
				break
			}
			inst.Args[i] = argfield.Parse(ui_extn)
		}
		inst.Op = iform.Op
		if debugDecode {
			log.Printf("%#x: search entry %d", ui, i)
			continue
		}
		break
	}
	if inst.Op == 0 && inst.Enc != 0 {
		return inst, errUnknown
	}
	return inst, nil
}
