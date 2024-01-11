package s390xasm

func GNUSyntax(inst Inst, pc uint64) string {
        if inst.Enc == 0 { 
                return ".long 0x0"
        } else if inst.Op == 0 { 
                return "error: unknown instruction"
        }
        return inst.String(pc)
}
