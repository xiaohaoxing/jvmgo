package instructions

import (
	"fmt"
	"jvmgo/ch09/instructions/base"
	. "jvmgo/ch09/instructions/comparisons"
	. "jvmgo/ch09/instructions/constants"
	. "jvmgo/ch09/instructions/control"
	. "jvmgo/ch09/instructions/conversions"
	. "jvmgo/ch09/instructions/extended"
	. "jvmgo/ch09/instructions/loads"
	. "jvmgo/ch09/instructions/math"
	. "jvmgo/ch09/instructions/references"
	. "jvmgo/ch09/instructions/reserved"
	. "jvmgo/ch09/instructions/stack"
	. "jvmgo/ch09/instructions/stores"
)

//单例 singleton
var (
	nop              = &NOP{}
	aconst_null      = &ACONST_NULL{}
	iconst_m1        = &ICONST_M1{}
	iconst_0         = &ICONST_0{}
	iconst_1         = &ICONST_1{}
	iconst_2         = &ICONST_2{}
	iconst_3         = &ICONST_3{}
	iconst_4         = &ICONST_4{}
	iconst_5         = &ICONST_5{}
	lconst_0         = &LCONST_0{}
	lconst_1         = &LCONST_1{}
	fconst_0         = &FCONST_0{}
	fconst_1         = &FCONST_1{}
	fconst_2         = &FCONST_2{}
	dconst_0         = &DCONST_0{}
	dconst_1         = &DCONST_1{}
	iload_0          = &ILOAD_0{}
	iload_1          = &ILOAD_1{}
	iload_2          = &ILOAD_2{}
	iload_3          = &ILOAD_3{}
	lload_0          = &LLOAD_0{}
	lload_1          = &LLOAD_1{}
	lload_2          = &LLOAD_2{}
	lload_3          = &LLOAD_3{}
	fload_0          = &FLOAD_0{}
	fload_1          = &FLOAD_1{}
	fload_2          = &FLOAD_2{}
	fload_3          = &FLOAD_3{}
	dload_0          = &DLOAD_0{}
	dload_1          = &DLOAD_1{}
	dload_2          = &DLOAD_2{}
	dload_3          = &DLOAD_3{}
	aload_0          = &ALOAD_0{}
	aload_1          = &ALOAD_1{}
	aload_2          = &ALOAD_2{}
	aload_3          = &ALOAD_3{}
	istore_0         = &ISTORE_0{}
	istore_1         = &ISTORE_1{}
	istore_2         = &ISTORE_2{}
	istore_3         = &ISTORE_3{}
	lstore_0         = &LSTORE_0{}
	lstore_1         = &LSTORE_1{}
	lstore_2         = &LSTORE_2{}
	lstore_3         = &LSTORE_3{}
	fstore_0         = &FSTORE_0{}
	fstore_1         = &FSTORE_1{}
	fstore_2         = &FSTORE_2{}
	fstore_3         = &FSTORE_3{}
	dstore_0         = &DSTORE_0{}
	dstore_1         = &DSTORE_1{}
	dstore_2         = &DSTORE_2{}
	dstore_3         = &DSTORE_3{}
	astore_0         = &ASTORE_0{}
	astore_1         = &ASTORE_1{}
	astore_2         = &ASTORE_2{}
	astore_3         = &ASTORE_3{}
	pop              = &POP{}
	pop2             = &POP2{}
	dup              = &DUP{}
	dup_x1           = &DUP_X1{}
	dup_x2           = &DUP_X2{}
	dup2             = &DUP2{}
	dup2_x1          = &DUP2_X1{}
	dup2_x2          = &DUP2_X2{}
	swap             = &SWAP{}
	iadd             = &IADD{}
	ladd             = &LADD{}
	fadd             = &FADD{}
	dadd             = &DADD{}
	isub             = &ISUB{}
	lsub             = &LSUB{}
	fsub             = &FSUB{}
	dsub             = &DSUB{}
	imul             = &IMUL{}
	lmul             = &LMUL{}
	fmul             = &FMUL{}
	dmul             = &DMUL{}
	idiv             = &IDIV{}
	ldiv             = &LDIV{}
	fdiv             = &FDIV{}
	ddiv             = &DDIV{}
	irem             = &IREM{}
	lrem             = &LREM{}
	frem             = &FREM{}
	drem             = &DREM{}
	ineg             = &INEG{}
	lneg             = &LNEG{}
	fneg             = &FNEG{}
	dneg             = &DNEG{}
	ishl             = &ISHL{}
	lshl             = &LSHL{}
	ishr             = &ISHR{}
	lshr             = &LSHR{}
	iushr            = &IUSHR{}
	lushr            = &LUSHR{}
	iand             = &IAND{}
	land             = &LAND{}
	ior              = &IOR{}
	lor              = &LOR{}
	ixor             = &IXOR{}
	lxor             = &LXOR{}
	iinc             = &IINC{}
	i2l              = &I2L{}
	i2f              = &I2F{}
	i2d              = &I2D{}
	i2b              = &I2B{}
	i2c              = &I2C{}
	i2s              = &I2S{}
	l2i              = &L2I{}
	l2f              = &L2F{}
	l2d              = &L2D{}
	f2i              = &F2I{}
	f2l              = &F2L{}
	f2d              = &F2D{}
	d2i              = &D2I{}
	d2l              = &D2L{}
	d2f              = &D2F{}
	lcmp             = &LCMP{}
	fcmpl            = &FCMPL{}
	fcmpg            = &FCMPG{}
	dcmpl            = &DCMPL{}
	dcmpg            = &DCMPG{}
	bipush           = &BIPUSH{}
	sipush           = &SIPUSH{}
	ldc              = &LDC{}
	ldc_w            = &LDC_W{}
	ldc2_w           = &LDC2_W{}
	iload            = &ILOAD{}
	lload            = &LLOAD{}
	fload            = &FLOAD{}
	dload            = &DLOAD{}
	aload            = &ALOAD{}
	aaload           = &AALOAD{}
	baload           = &BALOAD{}
	caload           = &CALOAD{}
	daload           = &DALOAD{}
	faload           = &FALOAD{}
	iaload           = &IALOAD{}
	laload           = &LALOAD{}
	saload           = &SALOAD{}
	istore           = &ISTORE{}
	lstore           = &LSTORE{}
	fstore           = &FSTORE{}
	dstore           = &DSTORE{}
	astore           = &ASTORE{}
	aastore          = &AASTORE{}
	iastore          = &IASTORE{}
	bastore          = &BASTORE{}
	castore          = &CASTORE{}
	dastore          = &DASTORE{}
	fastore          = &FASTORE{}
	lastore          = &LASTORE{}
	sastore          = &SASTORE{}
	ifeq             = &IFEQ{}
	ifne             = &IFNE{}
	iflt             = &IFLT{}
	ifge             = &IFGE{}
	ifgt             = &IFGT{}
	ifle             = &IFLE{}
	if_icmpeq        = &IF_ICMPEQ{}
	if_icmpne        = &IF_ICMPNE{}
	if_icmplt        = &IF_ICMPLT{}
	if_icmpge        = &IF_ICMPGE{}
	if_icmpgt        = &IF_ICMPGT{}
	if_icmple        = &IF_ICMPLE{}
	if_acmpeq        = &IF_ACMPEQ{}
	if_acmpne        = &IF_ACMPNE{}
	table_switch     = &TABLE_SWITCH{}
	lookup_swtich    = &LOOKUP_SWITCH{}
	ireturn          = &IRETURN{}
	lreturn          = &LRETURN{}
	freturn          = &FRETURN{}
	dreturn          = &DRETURN{}
	areturn          = &ARETURN{}
	get_static       = &GET_STATIC{}
	put_static       = &PUT_STATIC{}
	get_field        = &GET_FIELD{}
	put_field        = &PUT_FIELD{}
	invoke_virtual   = &INVOKE_VIRTUAL{}
	invoke_special   = &INVOKE_SPECIAL{}
	invoke_static    = &INVOKE_STATIC{}
	invoke_interface = &INVOKE_INTERFACE{}
	new              = &NEW{}
	new_array        = &NEW_ARRAY{}
	anew_array       = &ANEW_ARRAY{}
	arraylength      = &ARRAY_LENGTH{}
	check_cast       = &CHECK_CAST{}
	instance_of      = &INSTANCE_OF{}
	wide             = &WIDE{}
	multi_anew_array = &MULTI_ANEW_ARRAY{}
	ifnull           = &IFNULL{}
	ifnonnull        = &IFNONNULL{}
	goto_w           = &GOTO_W{}

	invoke_native = &INVOKE_NATIVE{}
)

func NewInstruction(opcode byte) base.Instruction {
	switch opcode {
	case 0x00:
		return nop
	case 0x01:
		return aconst_null
	case 0x02:
		return iconst_m1
	case 0x03:
		return iconst_0
	case 0x04:
		return iconst_1
	case 0x05:
		return iconst_2
	case 0x06:
		return iconst_3
	case 0x07:
		return iconst_4
	case 0x08:
		return iconst_5
	case 0x09:
		return lconst_0
	case 0x0a:
		return lconst_1
	case 0x0b:
		return fconst_0
	case 0x0c:
		return fconst_1
	case 0x0d:
		return fconst_2
	case 0x0e:
		return dconst_0
	case 0x0f:
		return dconst_1
	case 0x10:
		return bipush
	case 0x11:
		return sipush
	case 0x12:
		return ldc
	case 0x13:
		return ldc_w
	case 0x14:
		return ldc2_w
	case 0x15:
		return iload
	case 0x16:
		return lload
	case 0x17:
		return fload
	case 0x18:
		return dload
	case 0x19:
		return aload
	case 0x1a:
		return iload_0
	case 0x1b:
		return iload_1
	case 0x1c:
		return iload_2
	case 0x1d:
		return iload_3
	case 0x1e:
		return lload_0
	case 0x1f:
		return lload_1
	case 0x20:
		return lload_2
	case 0x21:
		return lload_3
	case 0x22:
		return fload_0
	case 0x23:
		return fload_1
	case 0x24:
		return fload_2
	case 0x25:
		return fload_3
	case 0x26:
		return dload_0
	case 0x27:
		return dload_1
	case 0x28:
		return dload_2
	case 0x29:
		return dload_3
	case 0x2a:
		return aload_0
	case 0x2b:
		return aload_1
	case 0x2c:
		return aload_2
	case 0x2d:
		return aload_3
	case 0x2e:
		return iaload
	case 0x2f:
		return laload
	case 0x30:
		return faload
	case 0x31:
		return daload
	case 0x32:
		return aaload
	case 0x33:
		return baload
	case 0x34:
		return caload
	case 0x35:
		return saload
	case 0x36:
		return istore
	case 0x37:
		return lstore
	case 0x38:
		return fstore
	case 0x39:
		return dstore
	case 0x3a:
		return astore
	case 0x3b:
		return istore_0
	case 0x3c:
		return istore_1
	case 0x3d:
		return istore_2
	case 0x3e:
		return istore_3
	case 0x3f:
		return lstore_0
	case 0x40:
		return lstore_1
	case 0x41:
		return lstore_2
	case 0x42:
		return lstore_3
	case 0x43:
		return fstore_0
	case 0x44:
		return fstore_1
	case 0x45:
		return fstore_2
	case 0x46:
		return fstore_3
	case 0x47:
		return dstore_0
	case 0x48:
		return dstore_1
	case 0x49:
		return dstore_2
	case 0x4a:
		return dstore_3
	case 0x4b:
		return astore_0
	case 0x4c:
		return astore_1
	case 0x4d:
		return astore_2
	case 0x4e:
		return astore_3
	case 0x4f:
		return iastore
	case 0x50:
		return lastore
	case 0x51:
		return fastore
	case 0x52:
		return dastore
	case 0x53:
		return aastore
	case 0x54:
		return bastore
	case 0x55:
		return castore
	case 0x56:
		return sastore
	case 0x57:
		return pop
	case 0x58:
		return pop2
	case 0x59:
		return dup
	case 0x5a:
		return dup_x1
	case 0x5b:
		return dup_x2
	case 0x5c:
		return dup2
	case 0x5d:
		return dup2_x1
	case 0x5e:
		return dup2_x2
	case 0x5f:
		return swap
	case 0x60:
		return iadd
	case 0x61:
		return ladd
	case 0x62:
		return fadd
	case 0x63:
		return dadd
	case 0x64:
		return isub
	case 0x65:
		return lsub
	case 0x66:
		return fsub
	case 0x67:
		return dsub
	case 0x68:
		return imul
	case 0x69:
		return lmul
	case 0x6a:
		return fmul
	case 0x6b:
		return dmul
	case 0x6c:
		return idiv
	case 0x6d:
		return ldiv
	case 0x6e:
		return fdiv
	case 0x6f:
		return ddiv
	case 0x70:
		return irem
	case 0x71:
		return lrem
	case 0x72:
		return frem
	case 0x73:
		return drem
	case 0x74:
		return ineg
	case 0x75:
		return lneg
	case 0x76:
		return fneg
	case 0x77:
		return dneg
	case 0x78:
		return ishl
	case 0x79:
		return lshl
	case 0x7a:
		return ishr
	case 0x7b:
		return lshr
	case 0x7c:
		return iushr
	case 0x7d:
		return lushr
	case 0x7e:
		return iand
	case 0x7f:
		return land
	case 0x80:
		return ior
	case 0x81:
		return lor
	case 0x82:
		return ixor
	case 0x83:
		return lxor
	case 0x84:
		return iinc
	case 0x85:
		return i2l
	case 0x86:
		return i2f
	case 0x87:
		return i2d
	case 0x88:
		return l2i
	case 0x89:
		return l2f
	case 0x8a:
		return l2d
	case 0x8b:
		return f2i
	case 0x8c:
		return f2l
	case 0x8d:
		return f2d
	case 0x8e:
		return d2i
	case 0x8f:
		return d2l
	case 0x90:
		return d2f
	case 0x91:
		return i2b
	case 0x92:
		return i2c
	case 0x93:
		return i2s
	case 0x94:
		return lcmp
	case 0x95:
		return fcmpl
	case 0x96:
		return fcmpg
	case 0x97:
		return dcmpl
	case 0x98:
		return dcmpg
	case 0x99:
		return ifeq
	case 0x9a:
		return ifne
	case 0x9b:
		return iflt
	case 0x9c:
		return ifge
	case 0x9d:
		return ifgt
	case 0x9e:
		return ifle
	case 0x9f:
		return if_icmpeq
	case 0xa0:
		return if_icmpne
	case 0xa1:
		return if_icmplt
	case 0xa2:
		return if_icmpge
	case 0xa3:
		return if_icmpgt
	case 0xa4:
		return if_icmple
	case 0xa5:
		return if_acmpeq
	case 0xa6:
		return if_acmpne
	case 0xa7:
		return &GOTO{}
	// case 0xa8: return &JSR{}
	// case 0xa9: return &RET{}
	case 0xaa:
		return table_switch
	case 0xab:
		return lookup_swtich
	case 0xac:
		return ireturn
	case 0xad:
		return lreturn
	case 0xae:
		return freturn
	case 0xaf:
		return dreturn
	case 0xb0:
		return areturn
	case 0xb1:
		return &RETURN{}
	case 0xb2:
		return get_static
	case 0xb3:
		return put_static
	case 0xb4:
		return get_field
	case 0xb5:
		return put_field
	case 0xb6:
		return invoke_virtual
	case 0xb7:
		return invoke_special
	case 0xb8:
		return invoke_static
	case 0xb9:
		return invoke_interface
	//case 0xba: return &INVOKE_DYNAMIC{}
	case 0xbb:
		return new
	case 0xbc:
		return new_array
	case 0xbd:
		return anew_array
	case 0xbe:
		return arraylength
	// case 0xbf: return athrow
	case 0xc0:
		return check_cast
	case 0xc1:
		return instance_of
	// case 0xc2: return monitorenter
	// case 0xc3: return monitorexit
	case 0xc4:
		return wide
	case 0xc5:
		return multi_anew_array
	case 0xc6:
		return ifnull
	case 0xc7:
		return ifnonnull
	case 0xc8:
		return goto_w
	// case 0xc9: return &JSR_W{}
	// case 0xca: breakpoint
	case 0xfe:
		return invoke_native
	// case 0xff: return impdep2
	default:
		panic(fmt.Errorf("Unsupported opcode: 0x%x!", opcode))
	}
}
