package reflector

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func TestValue(t *testing.T) {
	typer := reflect.TypeOf(int32(123))
	fmt.Println(typer)
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
	var byter byte = 126            // kind: uint8
	var runer rune = '哈'           // kind: int32
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
	var b_inter8 = INTER8(inter8)                   // kind int8
	var b_uinter64 = UINTER64(uinter64)             // kind uint64
	var pb_inter8 = PINTER8(p_inter8)               // kind ptr
	var pb_uinter64 = PUINTER64(p_uinter64)         // kind ptr
	var pp_inter8 P_PINTER8 = &pb_inter8            // kind ptr
	var pp_uinter64 P_PUINTER64 = &pb_uinter64      // kind ptr

	vars := []interface{}{inter, inter8, inter16, inter32, inter64, uinter, uinter8, uinter16, uinter32, uinter64,
		booler, byter, runer, floater32, floater64, complexer64, complexer128, stringer, uintptrer,
		p_inter, p_inter8, p_inter16, p_inter32, p_inter64, p_uinter, p_uinter8, p_uinter16, p_uinter32, p_uinter64,
		p_booler, p_byter, p_runer, p_floater32, p_floater64, p_complexer64, p_complexer128, p_stringer,
		b_inter8, b_uinter64, pb_inter8, pb_uinter64, pp_inter8, pp_uinter64, }

	for i := range vars {
		value := reflect.ValueOf(vars[i])
		if value.Kind() >= reflect.Int && value.Kind() <= reflect.Int64 {
			fmt.Print("type: ", value.Type(), " value: ", value.Int())
			fmt.Println(" over: ", value.OverflowInt(2147483647), value.OverflowInt(2147483648))
		}
		if value.Kind() >= reflect.Uint && value.Kind() <= reflect.Uint64 {
			fmt.Print("type: ", value.Type(), " value: ", value.Uint())
			fmt.Println(" over: ", value.OverflowUint(4294967295), value.OverflowUint(4294967296))
		}
		if value.Kind() == reflect.Ptr {
			fmt.Print("type: ", value.Type(), ", indirect: ", reflect.Indirect(value),
				", ptr: ", strconv.FormatUint(uint64(value.Pointer()), 16), ", p: ", vars[i])
			fmt.Println()
		}
	}
}

func TestMakeMap(t *testing.T) {
	mapers := []interface{}{make(map[string]int), make(map[interface{}]interface{})}
	for i := range mapers {
		typer := reflect.TypeOf(mapers[i])
		fmt.Println("type: ", typer)
		v := reflect.MakeMap(typer)
		v.SetMapIndex(reflect.ValueOf("a"), reflect.ValueOf(1))
		v.SetMapIndex(reflect.ValueOf("b"), reflect.ValueOf(2))
		fmt.Println(v)
	}
}

func TestMakeFunc(t *testing.T) {
	// make a func with signature: func(string, ...int) string
	str_t := reflect.TypeOf("")
	int_slice_t := reflect.SliceOf(reflect.TypeOf(1))
	{
		func_t := reflect.FuncOf([]reflect.Type{str_t, int_slice_t}, []reflect.Type{str_t}, true)
		fmt.Println("func_t: ", func_t)
		funcer := reflect.MakeFunc(func_t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(args[0].String() + " " + strconv.Itoa(args[1].Len()))}
		})
		fmt.Println("funcer: ", funcer)
		{
			ret := funcer.CallSlice([]reflect.Value{reflect.ValueOf("aaa"), reflect.ValueOf([]int{2, 3, 4})})
			fmt.Println(ret[0].Interface().(string))
		}
		{
			ret := funcer.Call([]reflect.Value{reflect.ValueOf("aaa")})
			fmt.Println(ret[0].Interface().(string))
		}
		fmt.Println()
	}
	{
		var a func(string, ...int) string
		func_t := reflect.TypeOf(a)
		fmt.Println("func_t: ", func_t)
		funcer := reflect.MakeFunc(func_t, func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(args[0].String() + " " + strconv.Itoa(args[1].Len()))}
		})
		fmt.Println("funcer: ", funcer)
		a = funcer.Interface().(func(string, ...int) string)
		ret := funcer.CallSlice([]reflect.Value{reflect.ValueOf("aaa"), reflect.ValueOf([]int{2, 3, 4})})
		fmt.Println(ret[0].Interface().(string))
		fmt.Println(a("aaa", []int{1, 2, 3}...))
		fmt.Println(a("aaa", 1, 2))
		fmt.Println()
	}
	{
		var a func(string, ...int) string
		pa := &a
		fmt.Println("pa: ", pa)
		value := reflect.ValueOf(&a).Elem()
		fmt.Println("value: ", value)
		fmt.Println("func_t: ", value.Type())
		funcer := reflect.MakeFunc(value.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(args[0].String() + " " + strconv.Itoa(args[1].Len()))}
		})
		fmt.Println("funcer: ", funcer)

		// same as: a = funcer.Interface().(func(string, ...int) string)
		value.Set(funcer)
		ret := funcer.CallSlice([]reflect.Value{reflect.ValueOf("aaa"), reflect.ValueOf([]int{2, 3, 4})})
		fmt.Println(ret[0].Interface().(string))
		fmt.Println(a("aaa", []int{1, 2, 3}...))
		fmt.Println(a("aaa", 1, 2))
		fmt.Println()
	}
	{
		var a = func(string, ...int) string {return ""}  // func not nil
		pa := &a
		fmt.Println("pa: ", pa)
		value := reflect.ValueOf(pa).Elem()
		fmt.Println("value: ", value)  // example value: 0x55a790. is it the pointer in the function section?
		fmt.Println("func_t: ", value.Type())
		funcer := reflect.MakeFunc(value.Type(), func(args []reflect.Value) (results []reflect.Value) {
			return []reflect.Value{reflect.ValueOf(args[0].String() + " " + strconv.Itoa(args[1].Len()))}
		})
		fmt.Println("funcer: ", funcer) // example value: 0x4ad220. is it the pointer in the function section?

		// same as: a = funcer.Interface().(func(string, ...int) string)
		value.Set(funcer)
		ret := funcer.CallSlice([]reflect.Value{reflect.ValueOf("aaa"), reflect.ValueOf([]int{2, 3, 4})})
		fmt.Println(ret[0].Interface().(string))
		fmt.Println(a("aaa", []int{1, 2, 3}...))
		fmt.Println(a("aaa", 1, 2))
	}
}