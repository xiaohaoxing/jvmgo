package heap

/*
各个类、方法、成员变量的修饰符的值。
*/
const (
	ACC_PUBLIC       = 0x0001 //public		|class|member|method|
	ACC_PRIVATE      = 0x0002 //private		|     |member|method|
	ACC_PROTECTED    = 0x0004 //protected		|     |member|method|
	ACC_STATIC       = 0x0008 //static		|     |member|method|
	ACC_FINAL        = 0x0010 //final			|class|member|method|
	ACC_SUPER        = 0x0020 //super			|class|      |      |
	ACC_SYNCHRONIZED = 0x0020 //synchromized	|     |      |method|
	ACC_VOLATILE     = 0x0040 //volatile		|     |member|      |
	ACC_BRIDGE       = 0x0040 //bridge		|     |      |method|
	ACC_TRANSIENT    = 0x0080 //transient		|     |member|      |
	ACC_VARARGS      = 0x0080 //varargs		|     |      |method|
	ACC_NATIVE       = 0x0100 //native		|     |      |method|
	ACC_INTERFACE    = 0x0200 //interface		|class|      |      |
	ACC_ABSTRACT     = 0x0400 //abstract		|class|      |method|
	ACC_STRICT       = 0x0800 //strict		|     |      |method|
	ACC_SYNTHETIC    = 0x1000 //synthetic		|class|member|method|
	ACC_ANNOTATION   = 0x2000 //annotation	|class|      |      |
	ACC_ENUM         = 0x4000 //enum			|class|member|      |
)
