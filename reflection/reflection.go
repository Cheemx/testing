package main

import "reflect"

// While we're testing and we are in "red"(the tests failing)
// is to write the smallest amount of code possible.

func walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walk(res.Interface(), fn)
		}
	}

	// First Draft
	//
	// field := val.Field(0)
	// fn(field.String())
	// The code is very unsafe and very naive!
	// WHY ?
	// OPTIMISTIC ASSUMPTIONS!!!
	// If x is nil then val will be nil
	// If val is nil then Field will panic and code will exit
	// OR
	// If val is not a struct or doesn't contain any fields at all
	// code will again panic and execution will stop
	// OR
	// calling String() on unknown type, although this will not panic but
	// still this will be wrong!
	// Although here it might not seem as a problem but in large codebases
	// this is dangerous!
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val
}
