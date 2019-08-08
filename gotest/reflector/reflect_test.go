package reflector

import (
	"crypto/elliptic"
	"fmt"
	"reflect"
	"strconv"
	"testing"
	"unsafe"
)

/*
    Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
*/

func TestType(t *testing.T) {
	var booler bool = true

	// these following type has Bits()
	var inter int = -19
	var inter8 int8 = -19
	var inter16 int16 = -19
	var inter32 int32 = -19
	var inter64 int64 = -19
	var uinter uint = 19
	var uinter8 uint8 = 19
	var uinter16 uint16 = 19
	var uinter32 uint32 = 19
	var uinter64 uint64 = 19
	var byter byte = 126 // kind: uint8
	var runer rune = '哈' // kind: int32
	var floater32 float32 = 33.2
	var floater64 float64 = 33.2
	var complexer64 complex64 = complex(floater32, floater32)   // size:  8, align:  4
	var complexer128 complex128 = complex(floater64, floater64) // size: 16, align:  8
	var uintptrer uintptr = uintptr(1)

	var stringer string = "孙一峰永远是我大哥！孙一峰永远是我大哥！" // size: 16, align:  8

	// all kind is ptr
	var p_booler = &booler
	var p_inter = &inter
	var p_inter8 = &inter8
	var p_inter16 = &inter16
	var p_inter32 = &inter32
	var p_inter64 = &inter64
	var p_uinter = &uinter
	var p_uinter8 = &uinter8
	var p_uinter16 = &uinter16
	var p_uinter32 = &uinter32
	var p_uinter64 = &uinter64
	var p_byter = &byter
	var p_runer = &runer
	var p_floater32 = &floater32
	var p_floater64 = &floater64
	var p_complexer64 = &complexer64
	var p_complexer128 = &complexer128
	var p_stringer = &stringer

	// basic
	type INTER8 int8
	type UINTER64 uint64
	type PINTER8 *int8
	type PUINTER64 *uint64
	type P_PINTER8 *PINTER8
	type P_PUINTER64 *PUINTER64
	var b_inter8 = INTER8(inter8)              // kind int8
	var b_uinter64 = UINTER64(uinter64)        // kind uint64
	var pb_inter8 = PINTER8(p_inter8)          // kind ptr
	var pb_uinter64 = PUINTER64(p_uinter64)    // kind ptr
	var pp_inter8 P_PINTER8 = &pb_inter8       // kind ptr
	var pp_uinter64 P_PUINTER64 = &pb_uinter64 // kind ptr

	vars := []interface{}{inter, inter8, inter16, inter32, inter64, uinter, uinter8, uinter16, uinter32, uinter64,
		booler, byter, runer, floater32, floater64, complexer64, complexer128, stringer, uintptrer,
		p_inter, p_inter8, p_inter16, p_inter32, p_inter64, p_uinter, p_uinter8, p_uinter16, p_uinter32, p_uinter64,
		p_booler, p_byter, p_runer, p_floater32, p_floater64, p_complexer64, p_complexer128, p_stringer,
		b_inter8, b_uinter64, pb_inter8, pb_uinter64, pp_inter8, pp_uinter64}

	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		if typer.Kind() < reflect.Int || typer.Kind() > reflect.Complex128 {
			fmt.Printf("name: %12s, str: %22s, bits: %3d, size: %2d, align: %2d, kind: %8s, pkg: %s\n",
				typer.Name(), typer.String(), -1, typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
			continue
		}
		fmt.Printf("name: %12s, str: %22s, bits: %3d, size: %2d, align: %2d, kind: %8s, pkg: %s\n",
			typer.Name(), typer.String(), typer.Bits(), typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
	}
}

func TestArraySliceStruct(t *testing.T) {
	type Stct1 struct {
		inter8 int8
	}
	type Stct2 struct {
		inter8   int8
		stringer string
	}
	type Stct3 struct {
		uinter8  uint8
		stringer string
		arrayer  []int16
	}
	type Stct4 struct {
		inter16  [81]int16
		stringer string
		arrayer  []uint8
	}
	type Stct5 struct {
		inter8  int8
		inter16 int16
	}
	type Stct6 struct {
		inter8  int8
		inter16 int16
		stct4   Stct4
	}

	s1, s2 := "hahahah", "lalalal"

	// for fix array, kind is array, size = size(type)*len, align = align(type)
	var fixer [5]int16 = [5]int16{12, -234, 66, -12345, 32000}
	var fixer_str [7]string = [7]string{s1, s2}
	var fixer_ptr [7]*string = [7]*string{&s1, &s2}
	var fixer_cmp [7]complex64 = [7]complex64{complex(3.2, 3.2)}

	// for slice, kind is slice, size = 24, align = 8
	var unfixer []uint8 = []uint8{12, 234, 66, 79}
	var slicer []string = make([]string, 7)
	slicer[0] = s1
	slicer[1] = s2

	// struct
	stct1 := Stct1{18}                                                    // 1 byte
	stct2 := Stct2{18, "hahahah"}                                         // 8(1) + 16 bytes
	stct3 := Stct3{18, "hahahah", []int16{1, 2, 3, 4, 5}}                 // 8(1) + 16 + 24 bytes
	stct4 := Stct4{[81]int16{1, 2, 3, 4, 5}, "hahahah", []uint8{1, 2, 3}} // 168(162) + 16 + 24 bytes
	stct5 := Stct5{12, 12345}                                             // 2(1) + 2 = 4 bytes
	stct6 := Stct6{12, 12345, stct4}                                      // 2(1) + 6(2) + 208 bytes

	// on 64-bit machine, it tries to align the struct size to folds of 8 bytes
	// for example [79]int16 need 158 bytes to store, so align to 160 bytes,
	// while [81]int16 need 162 bytes to store, so align to 168 bytes.
	fmt.Println(unsafe.Sizeof(stct1), unsafe.Alignof(stct1.inter8))
	fmt.Println(unsafe.Sizeof(stct2), unsafe.Offsetof(stct2.inter8), unsafe.Offsetof(stct2.stringer))
	fmt.Println(unsafe.Sizeof(stct3), unsafe.Offsetof(stct3.uinter8), unsafe.Offsetof(stct3.stringer), unsafe.Offsetof(stct3.arrayer))
	fmt.Println(unsafe.Sizeof(stct4), unsafe.Offsetof(stct4.inter16), unsafe.Offsetof(stct4.stringer), unsafe.Offsetof(stct4.arrayer))
	fmt.Println(unsafe.Sizeof(stct5), unsafe.Offsetof(stct5.inter8), unsafe.Offsetof(stct5.inter16))
	fmt.Println(unsafe.Sizeof(stct6), unsafe.Offsetof(stct6.inter8), unsafe.Offsetof(stct6.inter16), unsafe.Offsetof(stct6.stct4))

	// struct array
	var stct4_array [3]Stct4 = [3]Stct4{stct4, stct4} // size = 208*3, align = 8
	var stct4_slice []Stct4 = make([]Stct4, 129)      // size = 24, align = 8

	// array, slice pointer
	var p_fixer *[5]int16 = &fixer
	var p_fixer_ptr *[7]*string = &fixer_ptr
	var p_unfixer *[]uint8 = &unfixer
	var p_stct4_array *[3]Stct4 = &stct4_array
	var p_stct4_slice *[]Stct4 = &stct4_slice

	vars := []interface{}{fixer, fixer_str, fixer_ptr, fixer_cmp, unfixer, slicer,
		stct1, stct2, stct3, stct4, stct5, stct4_array, stct4_slice,
		p_fixer, p_fixer_ptr, p_unfixer, p_stct4_array, p_stct4_slice}

	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		if typer.Kind() < reflect.Int || typer.Kind() > reflect.Complex128 {
			fmt.Printf("name: %12s, str: %22s, bits: %3d, size: %3d, align: %2d, kind: %8s, pkg: %s\n",
				typer.Name(), typer.String(), -1, typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
			continue
		}
		fmt.Printf("name: %12s, str: %22s, bits: %3d, size: %3d, align: %2d, kind: %8s, pkg: %s\n",
			typer.Name(), typer.String(), typer.Bits(), typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
	}

	// test elem, len
	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		if typer.Kind() == reflect.Array {
			fmt.Printf("str: %22s, elem: %22s, len: %2d\n", typer.String(), typer.Elem(), typer.Len())
		}
		if typer.Kind() == reflect.Slice || typer.Kind() == reflect.Ptr {
			fmt.Printf("str: %22s, elem: %22s\n", typer.String(), typer.Elem())
		}
	}
}

func TestAlign(t *testing.T) {
	type Stct struct {
		inter8   int8
		inter16  int16
		inter32  int32
		stringer string
	}
	// the data is stored in little endian
	s := &Stct{2, 11*256 + 89, 17*256*256*256 + 15*256*256 + 13*256 + 11, "aaa"}
	fmt.Println(unsafe.Sizeof(*s))
	fmt.Println(unsafe.Offsetof(s.inter8))
	fmt.Println(unsafe.Offsetof(s.inter16))
	fmt.Println(unsafe.Offsetof(s.inter32))
	fmt.Println(unsafe.Offsetof(s.stringer))
	base := unsafe.Pointer(s)
	// *int <--> unsafe.Pointer <--> uintptr
	for i := 0; i < int(unsafe.Sizeof(*s)); i++ {
		p := (*byte)(unsafe.Pointer(uintptr(base) + uintptr(i)))
		fmt.Println(i, ": ", *p)
	}
	p := (*uint64)(base)
	fmt.Println(*p)
	fmt.Println(strconv.FormatUint(*p, 2))
}

func TestPointer(t *testing.T) {
	// uint32 0x12345678 is stored in little-endian:
	//  78   56   34   12
	//   |    |    |    |
	//  p    p+1  p+2  p+3
	// *uint8(p0)=0x78,
	// *uint16(p0)=0x5678, *uint16(p1)=0x3456, *uint16(p2)=0x5678,
	// *uint32(p0)=0x12345678

	// *ArbitraryType <--> unsafe.Pointer <--> uintptr, the pointer conversion process
	inter := uint32(0x12345678)
	p32 := &inter
	fmt.Println("p32: ", p32, ", value: ", strconv.FormatUint(uint64(*p32), 16))

	// *uint32 -> unsafe.Pointer -> *uint8
	p8 := (*uint8)(unsafe.Pointer(p32))
	p16_0 := (*uint16)(unsafe.Pointer(p32))
	fmt.Println("p8_0: ", p8, ", value: ", strconv.FormatUint(uint64(*p8), 16))

	// increment the pointer
	// *uint8 -> unsafe.Pointer -> uintptr
	uintptrer := uintptr(unsafe.Pointer(p8))
	fmt.Println("uintptr: ", strconv.FormatUint(uint64(uintptrer), 16))
	uintptrer = uintptrer + uintptr(1)
	// uintptr -> unsafe.Pointer -> *uint16
	p8_1 := (*uint8)(unsafe.Pointer(uintptrer))
	p16_1 := (*uint16)(unsafe.Pointer(uintptrer))
	fmt.Println("p8_1: ", p8_1, ", value: ", strconv.FormatUint(uint64(*p8_1), 16))
	uintptrer = uintptrer + uintptr(1)
	p8_2 := (*uint8)(unsafe.Pointer(uintptrer))
	p16_2 := (*uint16)(unsafe.Pointer(uintptrer))
	fmt.Println("p8_2: ", p8_2, ", value: ", strconv.FormatUint(uint64(*p8_2), 16))
	uintptrer = uintptrer + uintptr(1)
	p8_3 := (*uint8)(unsafe.Pointer(uintptrer))
	fmt.Println("p8_3: ", p8_3, ", value: ", strconv.FormatUint(uint64(*p8_3), 16))
	fmt.Println("p16_0: ", p16_0, ", value: ", strconv.FormatUint(uint64(*p16_0), 16))
	fmt.Println("p16_1: ", p16_1, ", value: ", strconv.FormatUint(uint64(*p16_1), 16))
	fmt.Println("p16_2: ", p16_2, ", value: ", strconv.FormatUint(uint64(*p16_2), 16))
	fmt.Println()

	// when fmt.Print(*p), it prints the hex of the uintptr address
	fmt.Println("pointer: ", p32)
	uintptr_p32 := uintptr(unsafe.Pointer(p32))
	fmt.Println(strconv.FormatUint(uint64(uintptr_p32), 16))
}

func TestFuncInterfaceStruct(t *testing.T) {
	aircraft := Aircraft{1, "B787"}
	bird := Bird{"pelican", *elliptic.P224().Params()}
	p_aircraft := &aircraft
	p_bird := &bird

	vars := []interface{}{aircraft, bird, p_aircraft, p_bird}
	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		fmt.Printf("name: %12s, str: %22s, bits: %3d, align: %2d, kind: %8s, pkg: %s\n",
			typer.Name(), typer.String(), typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
	}
	fmt.Println()

	// test field func: Field, FieldByIndex, FieldByName, FieldByNameFunc, NumField
	// field.Anonymous mean that whether the field is explicitly named
	typer := reflect.TypeOf(aircraft)
	num := typer.NumField()
	for i := 0; i < num; i++ {
		field := typer.Field(i)
		fmt.Printf("name: %s, type: %s, anonymous: %t, index: %d, off: %d, pkg: %s, tag: %s\n",
			field.Name, field.Type, field.Anonymous, field.Index, field.Offset, field.PkgPath, field.Tag)
	}
	fmt.Println()

	typer = reflect.TypeOf(bird)
	num = typer.NumField()
	for i := 0; i < num; i++ {
		field := typer.Field(i)
		fmt.Printf("name: %s, type: %s, anonymous: %t, index: %d, off: %d, pkg: %s, tag: %s\n",
			field.Name, field.Type, field.Anonymous, field.Index, field.Offset, field.PkgPath, field.Tag)
	}
	num_curve, _ := typer.FieldByName("CurveParams")
	for i := 0; i < num_curve.Type.NumField(); i++ {
		field := typer.FieldByIndex([]int{1, i}) // FieldByIndex to get cascade field
		fmt.Printf("name: %s, type: %s, anonymous: %t, index: %d, off: %d, pkg: %s, tag: %s\n",
			field.Name, field.Type, field.Anonymous, field.Index, field.Offset, field.PkgPath, field.Tag)
	}
	fmt.Println()

	// it's a tricky way, make *Interface or []Interface, and get its Elem()
	// for example: var flyer = reflect.TypeOf((*Fly)(nil)).Elem()
	var flyer = reflect.TypeOf([]Fly{}).Elem()
	var facer = reflect.TypeOf([]interface{}{}).Elem()
	var stringer = reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
	fmt.Println("name: ", flyer.String(), ", kind: ", flyer.Kind())
	fmt.Println("name: ", facer.String(), ", kind: ", facer.Kind())
	fmt.Println("name: ", stringer.String(), ", kind: ", stringer.Kind())
	// if A implements an interface, *A does too;
	// if *A implements an interface, A doesn't
	fmt.Println("implements flyer: ", reflect.TypeOf(aircraft).Implements(flyer))
	fmt.Println("implements flyer: ", reflect.TypeOf(p_aircraft).Implements(flyer))
	fmt.Println("implements flyer: ", reflect.TypeOf(bird).Implements(flyer))
	fmt.Println("implements flyer: ", reflect.TypeOf(p_bird).Implements(flyer))
	fmt.Println("implements stringer: ", reflect.TypeOf(aircraft).Implements(stringer))
	fmt.Println("implements stringer: ", reflect.TypeOf(p_aircraft).Implements(stringer))
	fmt.Println("implements stringer: ", reflect.TypeOf(bird).Implements(stringer))
	fmt.Println("implements stringer: ", reflect.TypeOf(p_bird).Implements(stringer))
	// everything implements empty interface
	fmt.Println("implements facer: ", reflect.TypeOf(aircraft).Implements(facer))
	fmt.Println("implements facer: ", reflect.TypeOf(p_aircraft).Implements(facer))
	fmt.Println("implements facer: ", reflect.TypeOf(bird).Implements(facer))
	fmt.Println("implements facer: ", reflect.TypeOf(p_bird).Implements(facer))
	fmt.Println()
	// t1.AssignableTo(t2) check whether t1 can directly assign to t2, or whether t1 implements t2
	fmt.Println("assignable: ", reflect.TypeOf(aircraft).AssignableTo(flyer))
	fmt.Println("assignable: ", reflect.TypeOf(p_aircraft).AssignableTo(flyer))
	fmt.Println("assignable: ", reflect.TypeOf(bird).AssignableTo(flyer))
	fmt.Println("assignable: ", reflect.TypeOf(p_bird).AssignableTo(flyer))
	// struct that implements interface can be assigned to the interface, however the other way can not
	fmt.Println("assignable: ", flyer.AssignableTo(reflect.TypeOf(aircraft)))
	fmt.Println("assignable: ", flyer.AssignableTo(reflect.TypeOf(p_aircraft)))
	fmt.Println("assignable: ", flyer.AssignableTo(reflect.TypeOf(bird)))
	fmt.Println("assignable: ", flyer.AssignableTo(reflect.TypeOf(p_bird)))
}

func TestCompareAssignConvert(t *testing.T) {
	t_int64 := reflect.TypeOf((*int64)(nil)).Elem()
	t_int := reflect.TypeOf((*int)(nil)).Elem()
	t_int16 := reflect.TypeOf((*int16)(nil)).Elem()
	t_uint64 := reflect.TypeOf((*uint64)(nil)).Elem()
	t_uint := reflect.TypeOf((*uint)(nil)).Elem()
	t_uint16 := reflect.TypeOf((*uint16)(nil)).Elem()
	t_float32 := reflect.TypeOf((*float32)(nil)).Elem()
	t_float64 := reflect.TypeOf((*float64)(nil)).Elem()
	t_complex64 := reflect.TypeOf((*complex64)(nil)).Elem()
	types := []reflect.Type{t_int, t_int16, t_int64, t_uint, t_uint16, t_uint64, t_float32, t_float64, t_complex64}

	// integers, unsigned integers, and floats are convertible to each other, while complexes can't
	for i := 0; i < len(types); i++ {
		for j := 0; j < len(types); j++ {
			fmt.Printf("%10s %10s, assign: %5t, convert: %t\n",
				types[i], types[j], types[i].AssignableTo(types[j]), types[i].ConvertibleTo(types[j]))
		}
	}

	// test convert                            // binary representation in memory (little-endian)
	var inter64 int64 = -997654321098765432  // 88 4b 46 46 ae 9e 27 f2
	var uinter64 uint64 = 0x1234567890abcdef // ef cd ab 90 78 56 34 12
	var floater64 float64 = 1.234E24         // b2 3c 7a 62 f4 54 f0 44
	getInMem := func(base unsafe.Pointer, size uintptr) {
		for i := 0; i < int(size); i++ {
			pb := (*byte)(unsafe.Pointer(uintptr(base) + uintptr(i)))
			fmt.Print(strconv.FormatUint(uint64(*pb), 16), " ")
		}
		fmt.Println()
	}
	// when bigger int is converted to smaller int, get the lowest bytes to reprsent new number.
	// when integer and float are converted to each other, it tries to convert the real value,
	// if the value is exceeded, e.g., convert float '1.2e30' to int64, the converted value is uncertain.
	getInMem(unsafe.Pointer(&inter64), unsafe.Sizeof(inter64))
	getInMem(unsafe.Pointer(&uinter64), unsafe.Sizeof(uinter64))
	getInMem(unsafe.Pointer(&floater64), unsafe.Sizeof(floater64))
	fmt.Println()

	// int64 (88 4b 46 46 ae 9e 27 f2) to int16 (88 4b) = 19336 = 0x4b88
	fmt.Println(int16(inter64))
	// uint64 (ef cd ab 90 78 56 34 12) to int16 (ef cd) = -12817 = 0xcdef
	fmt.Println(int16(uinter64))
	fmt.Println(int16(floater64))
	fmt.Println()

	// int64 (88 4b 46 46 ae 9e 27 f2) to uint16 (88 4b) = 19336 = 0x4b88
	fmt.Println(uint16(inter64))
	// uint64 (ef cd ab 90 78 56 34 12) to uint16 (ef cd) = 52719 = 0xcdef
	fmt.Println(uint16(uinter64))
	fmt.Println(uint16(floater64))
	fmt.Println()

	// int64 (88 4b 46 46 ae 9e 27 f2) to int32 (88 4b 46 46) = 1179011976 = 0x46464b88
	fmt.Println(int32(inter64))
	// uint64 (ef cd ab 90 78 56 34 12) to int32 (ef cd ab 90) = -1867788817 = 0x90abcdef
	fmt.Println(int32(uinter64))
	fmt.Println(int32(floater64))
	fmt.Println()

	// int64 (88 4b 46 46 ae 9e 27 f2) to uint32 (88 4b 46 46) = 1179011976 = 0x46464b88
	fmt.Println(uint32(inter64))
	// uint64 (ef cd ab 90 78 56 34 12) to uint32 (ef cd ab 90) = 2427178479 = 0x90abcdef
	fmt.Println(uint32(uinter64))
	fmt.Println(uint32(floater64))
	fmt.Println()

	fmt.Println(int64(floater64))
	fmt.Println(uint64(floater64))
	fmt.Println()

	fmt.Println(float64(inter64))
	fmt.Println(float64(uinter64))
}

func TestMap(t *testing.T) {
	var a map[string]int
	var b = make(map[interface{}]interface{})
	type Namer struct {
		string
		int
	}
	var c map[Namer]*Namer
	b["dsad"] = "dsds"
	b[1] = 2
	vars := []interface{}{a, b, c}
	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		fmt.Printf("name: %3s, str: %50s, size: %3d, align: %2d, kind: %5s, pkg: %s\n",
			typer.Name(), typer.String(), typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
	}
	fmt.Println()
	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		fmt.Println("key type: ", typer.Key(), ", key kind:", typer.Key().Kind())
		fmt.Println("value type: ", typer.Elem(), ", value kind:", typer.Elem().Kind())
	}
}

func funcer(inter int, flyer Fly) string {
	return "haha" + strconv.Itoa(inter)
}

func TestFunc(t *testing.T) {
	f1 := func() {}
	f2 := func(inter ...int) {}
	f3 := func() (int, string) { return 1, "hahah" }
	vars := []interface{}{f1, f2, f3, funcer}
	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		fmt.Printf("name: %3s, str: %35s, size: %3d, align: %2d, kind: %5s, pkg: %s\n",
			typer.Name(), typer.String(), typer.Size(), typer.Align(), typer.Kind().String(), typer.PkgPath())
	}
	fmt.Println()
	// In, NumIn, Out, NumOut, IsVariadic.
	for i := range vars {
		typer := reflect.TypeOf(vars[i])
		fmt.Println("func type: ", typer, ", kind:", typer.Kind(), ", variadic: ", typer.IsVariadic())
		for i := 0; i < typer.NumIn(); i++ {
			fmt.Println("In", i, ":", typer.In(i))
		}
		for i := 0; i < typer.NumOut(); i++ {
			fmt.Println("Out", i, ":", typer.Out(i))
		}
	}
}

// method means the func belongs to struct or interface
func TestStructMethod(t *testing.T) {
	aircraft := Aircraft{1, "B787"}
	bird := Bird{"pelican", *elliptic.P224().Params()}
	p_aircraft := &aircraft
	p_bird := &bird
	flyer := reflect.TypeOf([]Fly{}).Elem()
	facer := reflect.TypeOf([]interface{}{}).Elem()
	stringer := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()

	vars := []interface{}{aircraft, bird, p_aircraft, p_bird, flyer, stringer, facer}
	for i := range vars {
		var typer reflect.Type
		var ok bool
		if typer, ok = vars[i].(reflect.Type); ok {
		} else {
			typer = reflect.TypeOf(vars[i])
		}
		fmt.Println("typer: ", typer, ", kind: ", typer.Kind())
		// Method(i) only show the public method
		for i := 0; i < typer.NumMethod(); i++ {
			method := typer.Method(i)
			fmt.Printf("id: %d, name: %10s, type: %20s, pkg: %10s\n",
				method.Index, method.Name, method.Type, method.PkgPath)
		}
		fmt.Println()
	}
	smet, _ := reflect.TypeOf(p_aircraft).MethodByName("String")
	fmt.Printf("id: %d, name: %10s, type: %20s, pkg: %10s\n",
		smet.Index, smet.Name, smet.Type, smet.PkgPath)
	sret := smet.Func.Call([]reflect.Value{reflect.ValueOf(p_aircraft)})
	fmt.Println(sret[0].Interface().(string))

	fmet, _ := reflect.TypeOf(p_aircraft).MethodByName("Fly")
	fmt.Printf("id: %d, name: %10s, type: %20s, pkg: %10s\n",
		fmet.Index, fmet.Name, fmet.Type, fmet.PkgPath)
	// use CallSlice to call variadic function, and the variadic parameter should be a slice
	fret1 := fmet.Func.CallSlice([]reflect.Value{reflect.ValueOf(p_aircraft), reflect.ValueOf([]int{1, 2, 3})})
	fmt.Println(fret1[0].Interface().(string))
	fret2 := fmet.Func.Call([]reflect.Value{reflect.ValueOf(p_aircraft)})
	fmt.Println(fret2[0].Interface().(string))
}
