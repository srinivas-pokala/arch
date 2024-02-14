package s390xasm


type typ1_ExtndMnics struct {
	BaseOpStr string
	Value     uint8
	Offset    uint8
	ExtnOpStr string
}

type typ2_ExtndMnics struct {
	Value     uint8
	Offset    uint8
	ExtnOpStr string
}

type typ3_ExtndMnics struct {
	Value1    uint8
	Value2    uint8
	Offset1   uint8
	Offset2   uint8
	ExtnOpStr string
}

type typ4_ExtndMnics struct {
	BaseOpStr string
	Value1    uint8
	Value2    uint8
	Offset1   uint8
	Offset2   uint8
	ExtnOpStr string
}

type typ5_ExtndMnics struct {
	BaseOpStr string
	Value1    uint8
	Value2    uint8
	Value3    uint8
	Offset1   uint8
	Offset2   uint8
	Offset3   uint8
	ExtnOpStr string
}

type typ6_ExtndMnics struct {
	Value     uint16
	Offset    uint8
	ExtnOpStr string
}

func HandleExtndMnemonic(inst *Inst) string {

	brnchInstrExtndMnics := []typ1_ExtndMnics{
		//BIC - BRANCH INDIRECT ON CONDITION instruction
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 1, Offset: 0, ExtnOpStr: "BIO"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 2, Offset: 0, ExtnOpStr: "BIH"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 4, Offset: 0, ExtnOpStr: "BIL"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 7, Offset: 0, ExtnOpStr: "BINE"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 8, Offset: 0, ExtnOpStr: "BIE"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 11, Offset: 0, ExtnOpStr: "BINL"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 13, Offset: 0, ExtnOpStr: "BINH"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 14, Offset: 0, ExtnOpStr: "BINO"},
		typ1_ExtndMnics{BaseOpStr: "BIC", Value: 15, Offset: 0, ExtnOpStr: "BI"},

		//BCR - BRANCH ON CONDITION instruction
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 0, Offset: 0, ExtnOpStr: "NOPR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 1, Offset: 0, ExtnOpStr: "BOR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 2, Offset: 0, ExtnOpStr: "BHR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 4, Offset: 0, ExtnOpStr: "BLR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 7, Offset: 0, ExtnOpStr: "BNER"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 8, Offset: 0, ExtnOpStr: "BER"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 11, Offset: 0, ExtnOpStr: "BNLR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 13, Offset: 0, ExtnOpStr: "BNHR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 14, Offset: 0, ExtnOpStr: "BNOR"},
		typ1_ExtndMnics{BaseOpStr: "BCR", Value: 15, Offset: 0, ExtnOpStr: "BR"},

		//BC - BRANCH ON CONDITION instruction
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 0, Offset: 0, ExtnOpStr: "NOPR"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 1, Offset: 0, ExtnOpStr: "BO"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 2, Offset: 0, ExtnOpStr: "BH"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 4, Offset: 0, ExtnOpStr: "BL"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 7, Offset: 0, ExtnOpStr: "BNE"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 8, Offset: 0, ExtnOpStr: "BE"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 11, Offset: 0, ExtnOpStr: "BNL"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 13, Offset: 0, ExtnOpStr: "BNH"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 14, Offset: 0, ExtnOpStr: "BNO"},
		typ1_ExtndMnics{BaseOpStr: "BC", Value: 15, Offset: 0, ExtnOpStr: "B"},

		//BRC - BRANCH RELATIVE ON CONDITION instruction
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 0, Offset: 0, ExtnOpStr: "JNOP"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 1, Offset: 0, ExtnOpStr: "JO"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 2, Offset: 0, ExtnOpStr: "JH"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 4, Offset: 0, ExtnOpStr: "JL"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 7, Offset: 0, ExtnOpStr: "JNE"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 8, Offset: 0, ExtnOpStr: "JE"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 11, Offset: 0, ExtnOpStr: "JNL"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 13, Offset: 0, ExtnOpStr: "JNH"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 14, Offset: 0, ExtnOpStr: "JNO"},
		typ1_ExtndMnics{BaseOpStr: "BRC", Value: 15, Offset: 0, ExtnOpStr: "J"},

		//BRCL - BRANCH RELATIVE ON CONDITION LONG instruction
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 0, Offset: 0, ExtnOpStr: "JGNOP"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 1, Offset: 0, ExtnOpStr: "JGO"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 2, Offset: 0, ExtnOpStr: "JGH"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 4, Offset: 0, ExtnOpStr: "JGL"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 7, Offset: 0, ExtnOpStr: "JGNE"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 8, Offset: 0, ExtnOpStr: "JGE"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 11, Offset: 0, ExtnOpStr: "JGNL"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 13, Offset: 0, ExtnOpStr: "JGNH"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 14, Offset: 0, ExtnOpStr: "JGNO"},
		typ1_ExtndMnics{BaseOpStr: "BRCL", Value: 15, Offset: 0, ExtnOpStr: "JG"},
	}

	//Compare instructions
	cmpInstrExtndMnics := []typ2_ExtndMnics{
		typ2_ExtndMnics{Value: 2, Offset: 2, ExtnOpStr: "H"},
		typ2_ExtndMnics{Value: 4, Offset: 2, ExtnOpStr: "L"},
		typ2_ExtndMnics{Value: 6, Offset: 2, ExtnOpStr: "NE"},
		typ2_ExtndMnics{Value: 8, Offset: 2, ExtnOpStr: "E"},
		typ2_ExtndMnics{Value: 10, Offset: 2, ExtnOpStr: "NL"},
		typ2_ExtndMnics{Value: 12, Offset: 2, ExtnOpStr: "NH"},
	}

	//Load and Store instructions
	ldSt_InstrExtndMnics := []typ2_ExtndMnics{
		typ2_ExtndMnics{Value: 1, Offset: 2, ExtnOpStr: "O"},
		typ2_ExtndMnics{Value: 2, Offset: 2, ExtnOpStr: "H"},
		typ2_ExtndMnics{Value: 3, Offset: 2, ExtnOpStr: "NLE"},
		typ2_ExtndMnics{Value: 4, Offset: 2, ExtnOpStr: "L"},
		typ2_ExtndMnics{Value: 5, Offset: 2, ExtnOpStr: "NHE"},
		typ2_ExtndMnics{Value: 6, Offset: 2, ExtnOpStr: "LH"},
		typ2_ExtndMnics{Value: 7, Offset: 2, ExtnOpStr: "NE"},
		typ2_ExtndMnics{Value: 8, Offset: 2, ExtnOpStr: "E"},
		typ2_ExtndMnics{Value: 9, Offset: 2, ExtnOpStr: "NLH"},
		typ2_ExtndMnics{Value: 10, Offset: 2, ExtnOpStr: "HE"},
		typ2_ExtndMnics{Value: 11, Offset: 2, ExtnOpStr: "NL"},
		typ2_ExtndMnics{Value: 12, Offset: 2, ExtnOpStr: "LE"},
		typ2_ExtndMnics{Value: 13, Offset: 2, ExtnOpStr: "NH"},
		typ2_ExtndMnics{Value: 14, Offset: 2, ExtnOpStr: "NO"},
	}

	vecInstrExtndMnics := []typ2_ExtndMnics{
		typ2_ExtndMnics{Value: 0, Offset: 3, ExtnOpStr: "B"},
		typ2_ExtndMnics{Value: 1, Offset: 3, ExtnOpStr: "H"},
		typ2_ExtndMnics{Value: 2, Offset: 3, ExtnOpStr: "F"},
		typ2_ExtndMnics{Value: 3, Offset: 3, ExtnOpStr: "G"},
		typ2_ExtndMnics{Value: 4, Offset: 3, ExtnOpStr: "Q"},
		typ2_ExtndMnics{Value: 6, Offset: 3, ExtnOpStr: "LF"},
	}

	//VCEQ, VCH, VCHL
	vec2InstrExtndMnics := []typ3_ExtndMnics{
		typ3_ExtndMnics{Value1: 0, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "B"},
		typ3_ExtndMnics{Value1: 1, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "H"},
		typ3_ExtndMnics{Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "F"},
		typ3_ExtndMnics{Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "G"},
		typ3_ExtndMnics{Value1: 0, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "BS"},
		typ3_ExtndMnics{Value1: 1, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "HS"},
		typ3_ExtndMnics{Value1: 2, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "FS"},
		typ3_ExtndMnics{Value1: 3, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "GS"},
	}

	//VFAE, VFEE, VFENE
	vec21InstrExtndMnics := []typ3_ExtndMnics{
		typ3_ExtndMnics{Value1: 0, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "B"},
		typ3_ExtndMnics{Value1: 1, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "H"},
		typ3_ExtndMnics{Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "F"},
		typ3_ExtndMnics{Value1: 0, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "BS"},
		typ3_ExtndMnics{Value1: 1, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "HS"},
		typ3_ExtndMnics{Value1: 2, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "FS"},
		typ3_ExtndMnics{Value1: 0, Value2: 2, Offset1: 3, Offset2: 4, ExtnOpStr: "ZB"},
		typ3_ExtndMnics{Value1: 1, Value2: 2, Offset1: 3, Offset2: 4, ExtnOpStr: "ZH"},
		typ3_ExtndMnics{Value1: 2, Value2: 2, Offset1: 3, Offset2: 4, ExtnOpStr: "ZF"},
		typ3_ExtndMnics{Value1: 0, Value2: 3, Offset1: 3, Offset2: 4, ExtnOpStr: "ZBS"},
		typ3_ExtndMnics{Value1: 1, Value2: 3, Offset1: 3, Offset2: 4, ExtnOpStr: "ZHS"},
		typ3_ExtndMnics{Value1: 2, Value2: 3, Offset1: 3, Offset2: 4, ExtnOpStr: "ZFS"},
	}

	vec3InstrExtndMnics := []typ3_ExtndMnics{
		typ3_ExtndMnics{Value1: 2, Value2: 0, Offset1: 2, Offset2: 3, ExtnOpStr: "SB"},
		typ3_ExtndMnics{Value1: 3, Value2: 0, Offset1: 2, Offset2: 3, ExtnOpStr: "DB"},
		typ3_ExtndMnics{Value1: 4, Value2: 0, Offset1: 2, Offset2: 3, ExtnOpStr: "XB"},
	}

	vec4InstrExtndMnics := []typ4_ExtndMnics{
		// VFA - VECTOR FP ADD
		typ4_ExtndMnics{BaseOpStr: "VFA", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFASB"},
		typ4_ExtndMnics{BaseOpStr: "VFA", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFADB"},
		typ4_ExtndMnics{BaseOpStr: "VFA", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFASB"},
		typ4_ExtndMnics{BaseOpStr: "VFA", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFADB"},
		typ4_ExtndMnics{BaseOpStr: "VFA", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFAXB"},

		// VFD - VECTOR FP DIVIDE
		typ4_ExtndMnics{BaseOpStr: "VFD", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFDSB"},
		typ4_ExtndMnics{BaseOpStr: "VFD", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFDDB"},
		typ4_ExtndMnics{BaseOpStr: "VFD", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFDSB"},
		typ4_ExtndMnics{BaseOpStr: "VFD", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFDDB"},
		typ4_ExtndMnics{BaseOpStr: "VFD", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFDXB"},

		// VFLL - VECTOR FP LOAD LENGTHENED
		typ4_ExtndMnics{BaseOpStr: "VFLL", Value1: 2, Value2: 0, Offset1: 2, Offset2: 3, ExtnOpStr: "VFLFS"},
		typ4_ExtndMnics{BaseOpStr: "VFLL", Value1: 2, Value2: 8, Offset1: 2, Offset2: 3, ExtnOpStr: "WFLLS"},
		typ4_ExtndMnics{BaseOpStr: "VFLL", Value1: 3, Value2: 8, Offset1: 2, Offset2: 3, ExtnOpStr: "WFLLD"},

		// VFMAX - VECTOR FP MAXIMUM
		typ4_ExtndMnics{BaseOpStr: "VFMAX", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFMAXSB"},
		typ4_ExtndMnics{BaseOpStr: "VFMAX", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFMAXDB"},
		typ4_ExtndMnics{BaseOpStr: "VFMAX", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMAXSB"},
		typ4_ExtndMnics{BaseOpStr: "VFMAX", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMAXDB"},
		typ4_ExtndMnics{BaseOpStr: "VFMAX", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMAXXB"},

		// VFMIN - VECTOR FP MINIMUM
		typ4_ExtndMnics{BaseOpStr: "VFMIN", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFMINSB"},
		typ4_ExtndMnics{BaseOpStr: "VFMIN", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFMINDB"},
		typ4_ExtndMnics{BaseOpStr: "VFMIN", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMINSB"},
		typ4_ExtndMnics{BaseOpStr: "VFMIN", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMINDB"},
		typ4_ExtndMnics{BaseOpStr: "VFMIN", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMINXB"},

		// VFM - VECTOR FP MULTIPLY
		typ4_ExtndMnics{BaseOpStr: "VFM", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFMSB"},
		typ4_ExtndMnics{BaseOpStr: "VFM", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFMDB"},
		typ4_ExtndMnics{BaseOpStr: "VFM", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMSB"},
		typ4_ExtndMnics{BaseOpStr: "VFM", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMDB"},
		typ4_ExtndMnics{BaseOpStr: "VFM", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFMXB"},

		// VFSQ - VECTOR FP SQUARE ROOT
		typ4_ExtndMnics{BaseOpStr: "VFSQ", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFSQSB"},
		typ4_ExtndMnics{BaseOpStr: "VFSQ", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFSQDB"},
		typ4_ExtndMnics{BaseOpStr: "VFSQ", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFSQSB"},
		typ4_ExtndMnics{BaseOpStr: "VFSQ", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFSQDB"},
		typ4_ExtndMnics{BaseOpStr: "VFSQ", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFSQXB"},

		// VFS - VECTOR FP SUBTRACT
		typ4_ExtndMnics{BaseOpStr: "VFS", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFSSB"},
		typ4_ExtndMnics{BaseOpStr: "VFS", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFSDB"},
		typ4_ExtndMnics{BaseOpStr: "VFS", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFSSB"},
		typ4_ExtndMnics{BaseOpStr: "VFS", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFSDB"},
		typ4_ExtndMnics{BaseOpStr: "VFS", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFSXB"},

		// VFTCI - VECTOR FP TEST DATA CLASS IMMEDIATE
		typ4_ExtndMnics{BaseOpStr: "VFTCI", Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFTCISB"},
		typ4_ExtndMnics{BaseOpStr: "VFTCI", Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "VFTCIDB"},
		typ4_ExtndMnics{BaseOpStr: "VFTCI", Value1: 2, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFTCISB"},
		typ4_ExtndMnics{BaseOpStr: "VFTCI", Value1: 3, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFTCIDB"},
		typ4_ExtndMnics{BaseOpStr: "VFTCI", Value1: 4, Value2: 8, Offset1: 3, Offset2: 4, ExtnOpStr: "WFTCIXB"},
	}

	vec6InstrExtndMnics := []typ5_ExtndMnics{
		// VFCE - VECTOR FP COMPARE EQUAL
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 0, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 0, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 0, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 4, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCEXB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 4, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCEXBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 4, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 4, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 4, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 4, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 2, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 3, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 4, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKEXB"},
		typ5_ExtndMnics{BaseOpStr: "VFCE", Value1: 4, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKEXBS"},

		// VFCH - VECTOR FP COMPARE HIGH
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 0, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHSB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 0, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHSBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 0, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 0, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHSB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHSBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 4, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHXB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 4, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHXBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 4, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHSB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 4, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHSBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 4, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 4, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHSB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 2, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHSBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 3, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 4, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHXB"},
		typ5_ExtndMnics{BaseOpStr: "VFCH", Value1: 4, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHXBS"},

		// VFCHE - VECTOR FP COMPARE HIGH OR EQUAL
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 0, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 0, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 0, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 0, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFCHEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 4, Value2: 8, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHEXB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 4, Value2: 8, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFCHEXBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 4, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 4, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 4, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 4, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "VFKHEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHESB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 2, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHESBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHEDB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 3, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHEDBS"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 4, Value2: 12, Value3: 0, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHEXB"},
		typ5_ExtndMnics{BaseOpStr: "VFCHE", Value1: 4, Value2: 12, Value3: 1, Offset1: 3, Offset2: 4, Offset3: 5, ExtnOpStr: "WFKHEXBS"},

		// VFPSO - VECTOR FP PERFORM SIGN OPERATION
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 2, Value2: 0, Value3: 0, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "VFLCSB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 2, Value2: 8, Value3: 0, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLCSB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 2, Value2: 0, Value3: 1, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "VFLNSB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 2, Value2: 8, Value3: 1, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLNSB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 2, Value2: 0, Value3: 2, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "VFLPSB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 2, Value2: 8, Value3: 2, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLPSB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 3, Value2: 0, Value3: 0, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "VFLCDB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 3, Value2: 8, Value3: 0, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLCDB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 3, Value2: 0, Value3: 1, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "VFLNDB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 3, Value2: 8, Value3: 1, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLNDB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 3, Value2: 0, Value3: 2, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "VFLPDB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 3, Value2: 8, Value3: 2, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLPDB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 4, Value2: 8, Value3: 0, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLCXB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 4, Value2: 8, Value3: 1, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLNXB"},
		typ5_ExtndMnics{BaseOpStr: "VFPSO", Value1: 4, Value2: 8, Value3: 2, Offset1: 2, Offset2: 3, Offset3: 4, ExtnOpStr: "WFLPXB"},
	}

	vec7InstrExtndMnics := []typ4_ExtndMnics{
		// VFMA - VECTOR FP MULTIPLY AND ADD
		typ4_ExtndMnics{BaseOpStr: "VFMA", Value1: 0, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "VFMASB"},
		typ4_ExtndMnics{BaseOpStr: "VFMA", Value1: 0, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "VFMADB"},
		typ4_ExtndMnics{BaseOpStr: "VFMA", Value1: 8, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "WFMASB"},
		typ4_ExtndMnics{BaseOpStr: "VFMA", Value1: 8, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "WFMADB"},
		typ4_ExtndMnics{BaseOpStr: "VFMA", Value1: 8, Value2: 4, Offset1: 4, Offset2: 5, ExtnOpStr: "WFMAXB"},

		// VFMS - VECTOR FP MULTIPLY AND SUBTRACT
		typ4_ExtndMnics{BaseOpStr: "VFMS", Value1: 0, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "VFMSSB"},
		typ4_ExtndMnics{BaseOpStr: "VFMS", Value1: 0, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "VFMSDB"},
		typ4_ExtndMnics{BaseOpStr: "VFMS", Value1: 8, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "WFMSSB"},
		typ4_ExtndMnics{BaseOpStr: "VFMS", Value1: 8, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "WFMSDB"},
		typ4_ExtndMnics{BaseOpStr: "VFMS", Value1: 8, Value2: 4, Offset1: 4, Offset2: 5, ExtnOpStr: "WFMSXB"},

		// VFNMA - VECTOR FP NEGATIVE MULTIPLY AND ADD
		typ4_ExtndMnics{BaseOpStr: "VFNMA", Value1: 0, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "VFNMASB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMA", Value1: 0, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "VFNMADB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMA", Value1: 8, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "WFNMASB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMA", Value1: 8, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "WFNMADB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMA", Value1: 8, Value2: 4, Offset1: 4, Offset2: 5, ExtnOpStr: "WFNMAXB"},

		// VFNMS - VECTOR FP NEGATIVE MULTIPLY AND SUBTRACT
		typ4_ExtndMnics{BaseOpStr: "VFNMS", Value1: 0, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "VFNMSSB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMS", Value1: 0, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "VFNMSDB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMS", Value1: 8, Value2: 2, Offset1: 4, Offset2: 5, ExtnOpStr: "WFNMSSB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMS", Value1: 8, Value2: 3, Offset1: 4, Offset2: 5, ExtnOpStr: "WFNMSDB"},
		typ4_ExtndMnics{BaseOpStr: "VFNMS", Value1: 8, Value2: 4, Offset1: 4, Offset2: 5, ExtnOpStr: "WFNMSXB"},
	}

	opString := inst.Op.String()
	newOpStr := opString

	if inst.Enc == 0 {
		return ".long 0x0"
	} else if inst.Op == 0 {
		return "error: unknown instruction"
	}

	switch opString {
	// Case to handle all "Branch" instructions with one M-field operand
	case "BIC", "BCR", "BC", "BRC", "BRCL":

		for i := 0; i < len(brnchInstrExtndMnics); i++ {
			if opString == brnchInstrExtndMnics[i].BaseOpStr &&
				uint8(inst.Args[brnchInstrExtndMnics[i].Offset].(Mask)) == brnchInstrExtndMnics[i].Value {
				newOpStr = brnchInstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(brnchInstrExtndMnics[i].Offset))
				break
			}
		}

	// Case to handle all "Compare" instructions with one M-field operand
	case "CRB", "CGRB", "CRJ", "CGRJ", "CRT", "CGRT", "CIB", "CGIB", "CIJ", "CGIJ", "CIT", "CGIT", "CLRB", "CLGRB",
		"CLRJ", "CLGRJ", "CLRT", "CLGRT", "CLT", "CLGT", "CLIB", "CLGIB", "CLIJ", "CLGIJ", "CLFIT", "CLGIT":

		for i := 0; i < len(cmpInstrExtndMnics); i++ {
			//For CLT and CLGT instructions, M-value is the second operand.
			//Hence, set the offset to "1"
			if opString == "CLT" || opString == "CLGT" {
				cmpInstrExtndMnics[i].Offset = 1
			}

			if uint8(inst.Args[cmpInstrExtndMnics[i].Offset].(Mask)) == cmpInstrExtndMnics[i].Value {
				newOpStr = opString + cmpInstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(cmpInstrExtndMnics[i].Offset))
				break
			}
		}

	// Case to handle all "Load" and "Store" instructions with one M-field operand
	case "LOCHHI", "LOCHI", "LOCGHI", "LOCFHR", "LOCFH", "LOCR", "LOCGR", "LOC",
		"LOCG", "SELR", "SELGR", "SELFHR", "STOCFH", "STOC", "STOCG":

		for i := 0; i < len(ldSt_InstrExtndMnics); i++ {

			//For LOCFH, LOC, LOCG, SELR, SELGR, SELFHR, STOCFH, STOC, STOCG instructions,
			//M-value is the forth operand. Hence, set the offset to "3"
			if opString == "LOCFH" || opString == "LOC" || opString == "LOCG" || opString == "SELR" || opString == "SELGR" ||
				opString == "SELFHR" || opString == "STOCFH" || opString == "STOC" || opString == "STOCG" {
				ldSt_InstrExtndMnics[i].Offset = 3
			}

			if uint8(inst.Args[ldSt_InstrExtndMnics[i].Offset].(Mask)) == ldSt_InstrExtndMnics[i].Value {
				newOpStr = opString + ldSt_InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(ldSt_InstrExtndMnics[i].Offset))
				break
			}
		}

	// Case to handle all "Vector" instructions with one M-field operand
	case "VAVG", "VAVGL", "VERLLV", "VESLV", "VESRAV", "VESRLV", "VGFM", "VGM", "VMX", "VMXL", "VMRH", "VMRL", "VMN", "VMNL", "VREP",
		"VCLZ", "VCTZ", "VEC", "VECL", "VLC", "VLP", "VPOPCT", "VREPI", "VERIM", "VERLL", "VESL", "VESRA", "VESRL", "VGFMA", "VLREP",
		"VLGV", "VLVG", "VLBRREP", "VLER", "VLBR", "VSTBR", "VSTER", "VPK", "VME", "VMH", "VMLE", "VMLH", "VMLO", "VML", "VMO", "VMAE",
		"VMALE", "VMALO", "VMAL", "VMAH", "VMALH", "VMAO", "VMPH", "VMPLH", "VUPL", "VUPLL", "VSCBI", "VS", "VSUM", "VSUMG", "VSUMQ", "VA", "VACC":

		switch opString {

		case "VAVG", "VAVGL", "VERLLV", "VESLV", "VESRAV", "VESRLV", "VGFM", "VGM", "VMX", "VMXL", "VMRH", "VMRL", "VMN", "VMNL", "VREP":
			//M-field is 3rd arg for all these instructions. Hence, set the offset to "2"
			for i := 0; i < len(vecInstrExtndMnics)-2; i++ { // 0,1,2,3
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}

		case "VCLZ", "VCTZ", "VEC", "VECL", "VLC", "VLP", "VPOPCT", "VREPI":
			for i := 0; i < len(vecInstrExtndMnics)-2; i++ { //0,1,2,3
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset-1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset-1))
					break
				}
			}

		case "VERIM", "VERLL", "VESL", "VESRA", "VESRL", "VGFMA", "VLREP":
			for i := 0; i < len(vecInstrExtndMnics)-2; i++ { //0,1,2,3
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset+1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset+1))
					break
				}
			}

		case "VLGV", "VLVG":
			for i := 0; i < len(vecInstrExtndMnics)-2; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset+1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset+1))
					break
				}
			}

		case "VLBRREP", "VLER", "VSTER":
			for i := 1; i < len(vecInstrExtndMnics)-2; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset+1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset+1))
					break
				}
			}

		case "VPK":
			for i := 1; i < len(vecInstrExtndMnics)-2; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}

		case "VLBR", "VSTBR":
			for i := 1; i < len(vecInstrExtndMnics)-1; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset+1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset+1))
					break
				}
			}
		case "VME", "VMH", "VMLE", "VMLH", "VMLO", "VMO":
			for i := 0; i < len(vecInstrExtndMnics)-3; i++ { //0,1,2
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}

		case "VML":
			for i := 0; i < len(vecInstrExtndMnics)-3; i++ { //0,1,2
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == 1 {
						newOpStr = opString + string("HW")
					} else {
						newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					}
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}

		case "VMAE", "VMALE", "VMALO", "VMAL", "VMAH", "VMALH", "VMAO":
			for i := 0; i < len(vecInstrExtndMnics)-3; i++ { //0,1,2
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset+1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset+1))
					break
				}
			}

		case "VMPH", "VMPLH", "VUPL", "VUPLL": //0,1,2
			for i := 0; i < len(vecInstrExtndMnics)-3; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset-1].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset-1))
					break
				}
			}

		case "VSCBI", "VS", "VA", "VACC": // 0,1,2,3,4
			for i := 0; i < len(vecInstrExtndMnics)-1; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}
		case "VSUM", "VSUMG":
			for i := 1; i < len(vecInstrExtndMnics)-4; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}
		case "VSUMQ":
			for i := 2; i < len(vecInstrExtndMnics)-2; i++ {
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}
		}

	case "VLLEZ":
		for i := 0; i < len(vecInstrExtndMnics); i++ {
			if i == 4 {
				continue
			}
			if uint8(inst.Args[vecInstrExtndMnics[i].Offset+1].(Mask)) == vecInstrExtndMnics[i].Value {
				newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vecInstrExtndMnics[i].Offset+1))
				break
			}
		}

	case "VGBM":
		if uint16(inst.Args[1].(Imm)) == uint16(0) {
			newOpStr = "VZEO"
			removeArg(inst, int8(1))
		} else if uint16(inst.Args[1].(Imm)) == uint16(0xFFFF) {
			newOpStr = "VONE"
			removeArg(inst, int8(1))
		}
	case "VNO":
		if uint8(inst.Args[1].(VReg)) == uint8(inst.Args[2].(VReg)) { //Bitwise Not instruction(VNOT)  if V2 equal to v3
			newOpStr = opString + "T"
			removeArg(inst, int8(2))
		}

	case "VMSL":
		if uint8(inst.Args[4].(Mask)) == uint8(3) {
			newOpStr = opString + "G"
			removeArg(inst, int8(4))
		}

	case "VFLR":
		if uint8(inst.Args[2].(Mask)) == uint8(3) && ((inst.Args[3].(Mask)>>3)&0x1 == 0x1) {
			inst.Args[3] = (inst.Args[3].(Mask) ^ 0x8)
			newOpStr = "WFLRD"
			removeArg(inst, int8(2))
		} else if uint8(inst.Args[2].(Mask)) == uint8(4) && ((inst.Args[3].(Mask)>>3)&0x1 == 0x1) {
			inst.Args[3] = (inst.Args[3].(Mask) ^ 0x8)
			newOpStr = "WFLRX"
			removeArg(inst, int8(2))
		} else if uint8(inst.Args[2].(Mask)) == uint8(3) {
			newOpStr = "VFLRD"
			removeArg(inst, int8(2))
		}

	case "VLLEBRZ":
		if uint8(inst.Args[4].(Mask)) == uint8(1) {
			newOpStr = opString + "H"
			removeArg(inst, int8(4))
		} else if uint8(inst.Args[4].(Mask)) == uint8(2) {
			newOpStr = opString + "F"
			removeArg(inst, int8(4))
		} else if uint8(inst.Args[4].(Mask)) == uint8(3) {
			newOpStr = "LDRV"
			removeArg(inst, int8(4))
		} else if uint8(inst.Args[4].(Mask)) == uint8(6) {
			newOpStr = "LERV"
			removeArg(inst, int8(4))
		}

	case "VSCHP":
		if uint8(inst.Args[3].(Mask)) == uint8(2) {
			newOpStr = "VSCHSP"
			removeArg(inst, int8(3))
		} else if uint8(inst.Args[3].(Mask)) == uint8(3) {
			newOpStr = "VSCHDP"
			removeArg(inst, int8(3))
		} else if uint8(inst.Args[3].(Mask)) == uint8(4) {
			newOpStr = "VSCHXP"
			removeArg(inst, int8(3))
		}

	case "VSBCBI", "VSBI":
		if uint8(inst.Args[4].(Mask)) == uint8(4) {
			newOpStr = opString + vecInstrExtndMnics[4].ExtnOpStr
			removeArg(inst, int8(4))
		}

	case "VAC", "VACCC":
		if uint8(inst.Args[4].(Mask)) == uint8(4) {
			newOpStr = opString + vecInstrExtndMnics[3].ExtnOpStr
			removeArg(inst, int8(3))
		}

	case "VCEQ", "VCH", "VCHL":
		for i := 0; i < len(vec2InstrExtndMnics)-6; i++ {
			if uint8(inst.Args[vec2InstrExtndMnics[i].Offset1].(Mask)) == vec2InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec2InstrExtndMnics[i].Offset2].(Mask)) == vec2InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec2InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec2InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec2InstrExtndMnics[i].Offset2-1))
				break
			}
		}

	case "VPKS", "VPKLS":
		for i := 1; i < len(vec2InstrExtndMnics)-6; i++ {
			if i == 4 {
				continue
			}
			if uint8(inst.Args[vec2InstrExtndMnics[i].Offset1].(Mask)) == vec2InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec2InstrExtndMnics[i].Offset2].(Mask)) == vec2InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec2InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec2InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec2InstrExtndMnics[i].Offset2-1))
				break
			}
		}
	case "VFEE", "VFENE":
		var check bool
		for i := 0; i < len(vec21InstrExtndMnics); i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2].(Mask)) == vec21InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset2-1))
				check = true
				break
			}
		}
		if !check {
			if uint8(inst.Args[3].(Mask)) == 0 && (uint8(inst.Args[4].(Mask)) != uint8(0)) {
				newOpStr = opString + vec21InstrExtndMnics[0].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[0].Offset1))
			} else if uint8(inst.Args[3].(Mask)) == 1 && (uint8(inst.Args[4].(Mask)) != uint8(0)) {
				newOpStr = opString + vec21InstrExtndMnics[1].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[1].Offset1))
			} else if uint8(inst.Args[3].(Mask)) == 2 && (uint8(inst.Args[4].(Mask)) != uint8(0)) {
				newOpStr = opString + vec21InstrExtndMnics[2].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[2].Offset1))
			} else if uint8(inst.Args[4].(Mask)) == 0 {
				removeArg(inst, int8(vec21InstrExtndMnics[2].Offset2))
			}
		}

	case "VFAE", "VSTRC":
		off := uint8(0)
		var check bool
		if opString == "VSTRC" {
			off = uint8(1)
		}
		for i := 0; i < len(vec21InstrExtndMnics)-9; i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1+off].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2+off].(Mask)) == vec21InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1+off))
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset2+off-1))
				check = true
				break
			}
		}

		for i := 0; !(check) && (i < len(vec21InstrExtndMnics)-9); i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1+off].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2+off].(Mask)) == vec21InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1+off))
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset2+off-1))
				check = true
				break
			}
		}
		//for i := 3; !(check) && (i < len(vec21InstrExtndMnics)); i++ {
		for i := len(vec21InstrExtndMnics) - 1; !(check) && (i > 2); i-- {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1+off].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2+off].(Mask))&(vec21InstrExtndMnics[i].Value2) == vec21InstrExtndMnics[i].Value2 {
				x := uint8(inst.Args[vec21InstrExtndMnics[i].Offset2+off].(Mask)) ^ (vec21InstrExtndMnics[i].Value2)
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				if x != 0 {
					inst.Args[vec21InstrExtndMnics[i].Offset2+off] = Mask(x)
					removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1+off))
					check = true
					break
				} else {
					removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1+off))
					removeArg(inst, int8(vec21InstrExtndMnics[i].Offset2+off-1))
					check = true
					break
				}
			}
		}
		if !check && inst.Args[4+off].(Mask) == Mask(0) {
			removeArg(inst, int8(4+off))
			break
		}

	case "VSTRS":
		var check bool
		for i := 0; i < len(vec21InstrExtndMnics)-3; i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1+1].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2+1].(Mask)) == vec21InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1+1))
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset2))
				check = true
				break
			}
			if i == 2 {
				i = i + 3
			}
		}

		for i := 0; !(check) && (i < len(vec21InstrExtndMnics)-9); i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1+1].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2+1].(Mask)) != 0 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1+1))
				break
			}
		}

	case "VISTR":
		var check bool
		for i := 0; i < len(vec21InstrExtndMnics)-6; i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1-1].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2-1].(Mask)) == vec21InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1-1))
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset2-2))
				check = true
				break
			}
		}

		for i := 0; !(check) && (i < len(vec21InstrExtndMnics)-9); i++ {
			if uint8(inst.Args[vec21InstrExtndMnics[i].Offset1-1].(Mask)) == vec21InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec21InstrExtndMnics[i].Offset2-1].(Mask)) != 0 {
				newOpStr = opString + vec21InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec21InstrExtndMnics[i].Offset1-1))
				break
			}
		}

		if uint8(inst.Args[3].(Mask)) == 0 {
			removeArg(inst, int8(3))
			break
		}

	case "VCFPS":
		if inst.Args[2].(Mask) == Mask(2) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCEFB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCDGB"
			removeArg(inst, int8(2))
			break
		} else if uint8(inst.Args[2].(Mask)) == uint8(2) {
			newOpStr = "VCEFB"
			removeArg(inst, int8(2))
			break
		} else if uint8(inst.Args[2].(Mask)) == uint8(3) {
			newOpStr = "VCDGB"
			removeArg(inst, int8(2))
			break
		}

	case "VCFPL":
		if inst.Args[2].(Mask) == Mask(2) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCELFB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCDLGB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(2) {
			newOpStr = "VCELFB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) {
			newOpStr = "VCDLGB"
			removeArg(inst, int8(2))
			break
		}

	case "VCSFP":
		if inst.Args[2].(Mask) == Mask(2) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCFEB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCGDB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(2) {
			newOpStr = "VCFEB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) {
			newOpStr = "VCGDB"
			removeArg(inst, int8(2))
			break
		}

	case "VCLFP":
		if inst.Args[2].(Mask) == Mask(2) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCLFEB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			inst.Args[3] = Mask((inst.Args[3].(Mask)) ^ (0x8))
			newOpStr = "WCLGDB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(2) {
			newOpStr = "VCLFEB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) {
			newOpStr = "VCLGDB"
			removeArg(inst, int8(2))
			break
		}

	case "VFI":
		if inst.Args[2].(Mask) == Mask(2) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			newOpStr = "WFISB"
			removeArg(inst, int8(2))
			inst.Args[2] = Mask((inst.Args[2].(Mask)) ^ (0x8))
			break
		} else if inst.Args[2].(Mask) == Mask(3) && ((inst.Args[3].(Mask)>>3)&(0x3) == 1) {
			newOpStr = "WFIDB"
			removeArg(inst, int8(2))
			inst.Args[2] = Mask((inst.Args[2].(Mask)) ^ (0x8))
			break
		} else if inst.Args[2].(Mask) == Mask(4) && ((inst.Args[3].(Mask)>>3)&(0x1) == 1) {
			newOpStr = "WFIXB"
			removeArg(inst, int8(2))
			inst.Args[2] = Mask((inst.Args[2].(Mask)) ^ (0x8))
			break
		} else if inst.Args[2].(Mask) == Mask(2) {
			newOpStr = "VFISB"
			removeArg(inst, int8(2))
			break
		} else if inst.Args[2].(Mask) == Mask(3) {
			newOpStr = "VFIDB"
			removeArg(inst, int8(2))
			break
		}

	// Case to handle few vector instructions with 2 M-field operands
	case "VFA", "VFD", "VFLL", "VFMAX", "VFMIN", "VFM":
		for i := 0; i < len(vec4InstrExtndMnics); i++ {
			if opString == vec4InstrExtndMnics[i].BaseOpStr &&
				uint8(inst.Args[vec4InstrExtndMnics[i].Offset1].(Mask)) == vec4InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec4InstrExtndMnics[i].Offset2].(Mask)) == vec4InstrExtndMnics[i].Value2 {
				newOpStr = vec4InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec4InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec4InstrExtndMnics[i].Offset2-1))
				break
			}
		}

	// Case to handle few special "Vector" instructions with 2 M-field operands
	case "WFC", "WFK":
		for i := 0; i < len(vec3InstrExtndMnics); i++ {
			if uint8(inst.Args[vec3InstrExtndMnics[i].Offset1].(Mask)) == vec3InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec3InstrExtndMnics[i].Offset2].(Mask)) == vec3InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec3InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec3InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec3InstrExtndMnics[i].Offset2-1))
				break
			}
		}

	// Case to handle few vector instructions with 2 M-field operands
	case "VFMA", "VFMS", "VFNMA", "VFNMS":
		for i := 0; i < len(vec7InstrExtndMnics); i++ {
			if opString == vec7InstrExtndMnics[i].BaseOpStr &&
				uint8(inst.Args[vec7InstrExtndMnics[i].Offset1].(Mask)) == vec7InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec7InstrExtndMnics[i].Offset2].(Mask)) == vec7InstrExtndMnics[i].Value2 {
				newOpStr = vec7InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec7InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec7InstrExtndMnics[i].Offset2-1))
				break
			}
		}

	// List of instructions with 3 M-field operands.
	case "VFCE", "VFCH", "VFCHE", "VFPSO":
		for i := 0; i < len(vec6InstrExtndMnics); i++ {
			if opString == vec6InstrExtndMnics[i].BaseOpStr &&
				uint8(inst.Args[vec6InstrExtndMnics[i].Offset1].(Mask)) == vec6InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec6InstrExtndMnics[i].Offset2].(Mask)) == vec6InstrExtndMnics[i].Value2 &&
				uint8(inst.Args[vec6InstrExtndMnics[i].Offset3].(Mask)) == vec6InstrExtndMnics[i].Value3 {
				newOpStr = vec6InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec6InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec6InstrExtndMnics[i].Offset2-1))
				removeArg(inst, int8(vec6InstrExtndMnics[i].Offset3-2))
				break
			}
		}

	default:
		return opString
	}
	return newOpStr
}

func GNUSyntax(inst Inst, pc uint64) string {
	if inst.Enc == 0 {
		return ".long 0x0"
	} else if inst.Op == 0 {
		return "error: unknown instruction"
	}
	return inst.String(pc)
}

// removeArg removes the arg in inst.Args[index].
func removeArg(inst *Inst, index int8) {
	for i := int(index); i < len(inst.Args); i++ {
		if i+1 < len(inst.Args) {
			inst.Args[i] = inst.Args[i+1]
		} else {
			inst.Args[i] = nil
		}
	}
}
