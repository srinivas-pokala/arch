package s390xasm

//import "fmt"

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

	//Vector instructions
	vecInstrExtndMnics := []typ2_ExtndMnics{
		typ2_ExtndMnics{Value: 0, Offset: 3, ExtnOpStr: "B"},
		typ2_ExtndMnics{Value: 1, Offset: 3, ExtnOpStr: "H"},
		typ2_ExtndMnics{Value: 2, Offset: 3, ExtnOpStr: "F"},
		typ2_ExtndMnics{Value: 3, Offset: 3, ExtnOpStr: "G"},
		typ2_ExtndMnics{Value: 4, Offset: 3, ExtnOpStr: "Q"},
		typ2_ExtndMnics{Value: 6, Offset: 3, ExtnOpStr: "E"},
	}

	vec2InstrExtndMnics := []typ3_ExtndMnics{
		typ3_ExtndMnics{Value1: 0, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "B"},
		typ3_ExtndMnics{Value1: 1, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "H"},
		typ3_ExtndMnics{Value1: 2, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "F"},
		typ3_ExtndMnics{Value1: 3, Value2: 0, Offset1: 3, Offset2: 4, ExtnOpStr: "G"},
		typ3_ExtndMnics{Value1: 0, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "BS"},
		typ3_ExtndMnics{Value1: 1, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "HS"},
		typ3_ExtndMnics{Value1: 2, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "FS"},
		typ3_ExtndMnics{Value1: 3, Value2: 1, Offset1: 3, Offset2: 4, ExtnOpStr: "GS"},
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

	vec5InstrExtndMnics := []typ1_ExtndMnics{
		// VSCHP - DECIMAL SCALE AND CONVERT TO HFP
		typ1_ExtndMnics{BaseOpStr: "VSCHP", Value: 2, Offset: 3, ExtnOpStr: "VSCHSP"},
		typ1_ExtndMnics{BaseOpStr: "VSCHP", Value: 3, Offset: 3, ExtnOpStr: "VSCHDP"},
		typ1_ExtndMnics{BaseOpStr: "VSCHP", Value: 4, Offset: 3, ExtnOpStr: "VSCHXP"},

		// VAC - VECTOR ADD WITH CARRY instruction
		typ1_ExtndMnics{BaseOpStr: "VAC", Value: 4, Offset: 4, ExtnOpStr: "VACQ"},

		// VACCC - VECTOR ADD WITH CARRY COMPUTE CARRY instruction
		typ1_ExtndMnics{BaseOpStr: "VACCC", Value: 4, Offset: 4, ExtnOpStr: "VACCCQ"},

		// VMSL - VECTOR MULTIPLY SUM LOGICAL instruction
		typ1_ExtndMnics{BaseOpStr: "VMSL", Value: 3, Offset: 4, ExtnOpStr: "VMSLG"},

		// VSBI - VECTOR SUBTRACT WITH BORROW INDICATION instruction
		typ1_ExtndMnics{BaseOpStr: "VSBI", Value: 4, Offset: 4, ExtnOpStr: "VSBIQ"},

		// VSBCBI - VECTOR SUBTRACT WITH BORROW COMPUTE BORROW INDICATION instruction
		typ1_ExtndMnics{BaseOpStr: "VSBCBI", Value: 4, Offset: 4, ExtnOpStr: "VSBCBIQ"},

		// VSTEBRF - VECTOR STORE BYTE REVERSED ELEMENT instruction
		typ1_ExtndMnics{BaseOpStr: "VSTEBRF", Value: 0, Offset: 4, ExtnOpStr: "STERV"},

		// VSTEBRG - VECTOR STORE BYTE REVERSED ELEMENT instruction
		typ1_ExtndMnics{BaseOpStr: "VSTEBRG", Value: 0, Offset: 4, ExtnOpStr: "STDRV"},
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

	vec8InstrExtndMnics := []typ6_ExtndMnics{
		// VGBM - VECTOR GENERATE BYTE MASK
		typ6_ExtndMnics{Value: 0, Offset: 1, ExtnOpStr: "VZERO"},
		typ6_ExtndMnics{Value: 0xFFFF, Offset: 1, ExtnOpStr: "VONE"},
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
				//skipArgList[0] = int8(brnchInstrExtndMnics[i].Offset)
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
				//skipArgList[0] = int8(cmpInstrExtndMnics[i].Offset)
				break
			}
		}

	// Case to handle all "Load" and "Store" instructions with one M-field operand
	case "LOCHHI", "LOCHI", "LOCGHI", "LOCFHR", "LOCFH", "LOCR", "LOCGR", "LOC",
		"LOCG", "SELR", "SELGR", "SELFHR", "STOCFH", "STOC", "STOCG":

		for i := 0; i < len(ldSt_InstrExtndMnics); i++ {

			//For LOCFH, LOC, LOCG, SELR, SELGR, SELFHR, STOCFH, STOC, STOCG instructions,
			//M-value is the forth operand. Hence, set the offset to "3"
			if opString == "LOCFH" || opString == "LOC" || opString == "SELR" || opString == "SELGR" ||
				opString == "SELFHR" || opString == "STOCFH" || opString == "STOC" || opString == "STOCG" {
				ldSt_InstrExtndMnics[i].Offset = 3
			}

			if uint8(inst.Args[ldSt_InstrExtndMnics[i].Offset].(Mask)) == ldSt_InstrExtndMnics[i].Value {
				newOpStr = opString + ldSt_InstrExtndMnics[i].ExtnOpStr
				//skipArgList[0] = int8(ldSt_InstrExtndMnics[i].Offset)
				removeArg(inst, int8(ldSt_InstrExtndMnics[i].Offset))
				break
			}
		}

	// Case to handle all "Vector" instructions with one M-field operand
	case "VA", "VACC", "VAVG", "VAVGL", "VCLZ", "VCTZ", "VEC", "VECL", "VERIM", "VERLLV", "VERLL",
		"VESLV", "VESL", "VESRAV", "VESRA", "VESRLV", "VESRL", "VGFM", "VGFMA", "VGM", "VLREP", "VLLEBRZ",
		"VLBR", "VLC", "VLER", "VLGV", "VLLEZ", "VLP", "VLVG", "VMX", "VMXL", "VMRH", "VMRL", "VMN", "VMNL",
		"VMAE", "VMALE", "VMALO", "VMAL", "VMAH", "VMALH", "VMAO", "VME", "VMH", "VMLE", "VMLH", "VMLO",
		"VML", "VMO", "VPK", "VPOPCT", "VREP", "VREPI", "VSEG", "VSTBR", "VSTER", "VS", "VSCBI", "VSUMG",
		"VSUMQ", "VSUM", "VUPH", "VUPL", "VUPLH", "VUPLL", "VLBRREP":

		switch opString {

		case "VCLZ", "VCTZ", "VEC", "VECL", "VLC", "VLP", "VPOPCT",
			"VREPI", "VSEG", "VUPH", "VUPLH", "VUPL", "VUPLL":

			//M-field is 3rd arg for all these instructions. Hence, set the offset to "2"
			for i := 0; i < len(vecInstrExtndMnics); i++ {
				vecInstrExtndMnics[i].Offset = 2
				if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					//skipArgList[0] = int8(vecInstrExtndMnics[i].Offset)
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}

		case "VERLL", "VESL", "VESRA", "VESRL", "VGFMA", "VLREP", "VLBRREP", "VLLEBRZ", "VLBR", "VLER",
			"VLLEZ", "VMAE", "VMALE", "VMALO", "VMAL", "VMAH", "VMALH", "VMAO", "VSTBR", "VSTER", "VERIM", "VLGV", "VLVG":
			//M-field is 5th arg for all these instructions. Hence, set the offset to "4"
			for i := 0; i < len(vecInstrExtndMnics); i++ {
				vecInstrExtndMnics[i].Offset = 4
				//fmt.Printf("Srinivas:Mnemonic:%s ExtnOpStr:%s Arg:%s,Type:%T off:%v\n", inst.Op.String(), vecInstrExtndMnics[i].ExtnOpStr, inst.Args[vecInstrExtndMnics[i].Offset].String(0xFFFF), inst.Args[vecInstrExtndMnics[i].Offset], vecInstrExtndMnics[i].Offset)
				if opString == "VLLEZ" && uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == 6 {
					newOpStr = opString + string("LF")
					//skipArgList[0] = int8(vecInstrExtndMnics[i].Offset)
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				} else if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					//skipArgList[0] = int8(vecInstrExtndMnics[i].Offset)
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}

		default:

			//"VA", "VACC", "VAVG", "VAVGL", "VERLLV", "VESLV", "VESRAV", "VESRLV", "VGFM", "VGM", "VMX", "VMXL",
			//"VMRH", "VMRL", "VMN", "VMNL", "VME", "VMH", "VMLE", "VMLH", "VMLO", "VML", "VMO", "VPK", "VREP",
			//"VS", "VSCBI", "VSUMG", "VSUMQ", "VSUM"
			//M-field is 4th arg for all these instructions.
			for i := 0; i < len(vecInstrExtndMnics); i++ {
				//fmt.Printf("Srinivas:Mnemonic:%s ExtnOpStr:%s Arg:%s,Type:%T off:%v\n", inst.Op.String(), vecInstrExtndMnics[i].ExtnOpStr, inst.Args[vecInstrExtndMnics[i].Offset].String(0xFFFF), inst.Args[vecInstrExtndMnics[i].Offset], vecInstrExtndMnics[i].Offset)
				if opString == "VML" && uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == 1 {
					newOpStr = opString + string("HW")
					//skipArgList[0] = int8(vecInstrExtndMnics[i].Offset)
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				} else if uint8(inst.Args[vecInstrExtndMnics[i].Offset].(Mask)) == vecInstrExtndMnics[i].Value {
					newOpStr = opString + vecInstrExtndMnics[i].ExtnOpStr
					//skipArgList[0] = int8(vecInstrExtndMnics[i].Offset)
					removeArg(inst, int8(vecInstrExtndMnics[i].Offset))
					break
				}
			}
		}

	// Case to handle all "Vector" instructions with 2 M-field operands
	case "VCEQ", "VCH", "VCHL", "VPKS", "VPKLS", "VFEE", "VFENE", "VISTR", "VFAE", "VSTRC", "VSTRS":
		for i := 0; i < len(vec2InstrExtndMnics); i++ {
			if uint8(inst.Args[vec2InstrExtndMnics[i].Offset1].(Mask)) == vec2InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec2InstrExtndMnics[i].Offset2].(Mask)) == vec2InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec2InstrExtndMnics[i].ExtnOpStr
				//skipArgList[0] = int8(vec2InstrExtndMnics[i].Offset1)
				//skipArgList[1] = int8(vec2InstrExtndMnics[i].Offset2)
				removeArg(inst, int8(vec2InstrExtndMnics[i].Offset1))
				if (opString != "VFAE" && opString != "VSTRC") &&
					!((opString == "VFEE" || opString == "VFENE" || opString == "VISTR" || opString == "VSTRS") &&
						(vec2InstrExtndMnics[i].Value2 == 0)) {
					removeArg(inst, int8(vec2InstrExtndMnics[i].Offset2-1))
				}
				break
			}
		}

	// Case to handle few special "Vector" instructions with 2 M-field operands
	case "WFC", "WFK":

		for i := 0; i < len(vec3InstrExtndMnics); i++ {
			if uint8(inst.Args[vec3InstrExtndMnics[i].Offset1].(Mask)) == vec3InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec3InstrExtndMnics[i].Offset2].(Mask)) == vec3InstrExtndMnics[i].Value2 {
				newOpStr = opString + vec3InstrExtndMnics[i].ExtnOpStr
				//skipArgList[0] = int8(vec3InstrExtndMnics[i].Offset1)
				//skipArgList[1] = int8(vec3InstrExtndMnics[i].Offset2)
				removeArg(inst, int8(vec3InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec3InstrExtndMnics[i].Offset2-1))
				break
			}
		}

	// Case to handle few vector instructions with 2 M-field operands
	case "VFA", "VFD", "VFLL", "VFMAX", "VFMIN", "VFM":

		for i := 0; i < len(vec4InstrExtndMnics); i++ {
			if opString == vec4InstrExtndMnics[i].BaseOpStr &&
				uint8(inst.Args[vec4InstrExtndMnics[i].Offset1].(Mask)) == vec4InstrExtndMnics[i].Value1 &&
				uint8(inst.Args[vec4InstrExtndMnics[i].Offset2].(Mask)) == vec4InstrExtndMnics[i].Value2 {
				newOpStr = vec4InstrExtndMnics[i].ExtnOpStr
				//skipArgList[0] = int8(vec4InstrExtndMnics[i].Offset1)
				//skipArgList[1] = int8(vec4InstrExtndMnics[i].Offset2)
				removeArg(inst, int8(vec4InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec4InstrExtndMnics[i].Offset2-1))
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
				//skipArgList[0] = int8(vec7InstrExtndMnics[i].Offse)
				//skipArgList[1] = int8(vec7InstrExtndMnics[i].Offset2)
				removeArg(inst, int8(vec7InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec7InstrExtndMnics[i].Offset2-1))
				break
			}
		}

		// Case to handle few vector instructions with one M-field operand
	case "VSCHP", "VAC", "VACCC", "VMSL", "VSBI", "VSBCBI", "VSTEBRF", "VSTEBRG":
		//case "VSCHP", "VAC":
		for i := 0; i < len(vec5InstrExtndMnics); i++ {
			if (opString == vec5InstrExtndMnics[i].BaseOpStr) && (uint8(inst.Args[vec5InstrExtndMnics[i].Offset].(Mask)) == vec5InstrExtndMnics[i].Value) {
				newOpStr = vec5InstrExtndMnics[i].ExtnOpStr
				//skipArgList[0] = int8(vec5InstrExtndMnics[i].Offset)
				removeArg(inst, int8(vec5InstrExtndMnics[i].Offset))
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
				//skipArgList[0] = int8(vec6InstrExtndMnics[i].Offset1)
				// skipArgList[1] = int8(vec6InstrExtndMnics[i].Offset2)
				// skipArgList[2] = int8(vec6InstrExtndMnics[i].Offset3)
				removeArg(inst, int8(vec6InstrExtndMnics[i].Offset1))
				removeArg(inst, int8(vec6InstrExtndMnics[i].Offset2-1))
				removeArg(inst, int8(vec6InstrExtndMnics[i].Offset3-2))
				break
			}
		}

	case "VGBM":
		for i := 0; i < len(vec8InstrExtndMnics); i++ {
			if uint16(inst.Args[vec8InstrExtndMnics[i].Offset].(Imm)) == vec8InstrExtndMnics[i].Value {
				newOpStr = vec8InstrExtndMnics[i].ExtnOpStr
				removeArg(inst, int8(vec8InstrExtndMnics[i].Offset))
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
