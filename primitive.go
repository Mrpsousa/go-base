package main

import "fmt"

func main() {
	// Booleano
	var b bool = true

	// Inteiros
	var i8 int8 = -128
	var i16 int16 = -32768
	var i32 int32 = -2147483648
	var i64 int64 = -9223372036854775808

	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	var i int = -2147483648  // Inteiro com tamanho dependente do sistema
	var ui uint = 4294967295 // Inteiro positivo com tamanho dependente do sistema

	// Ponto flutuante
	var f32 float32 = 3.1415
	var f64 float64 = 3.141592653589793

	// Complexos
	var c64 complex64 = complex(3, 4)
	var c128 complex128 = complex(3, 4)

	// String
	var s string = "Hello, World!"

	// Byte
	var by byte = 'A'

	// Rune
	var r rune = '東'

	fmt.Println(b)
	fmt.Println(i8)
	fmt.Println(i16)
	fmt.Println(i32)
	fmt.Println(i64)
	fmt.Println(ui8)
	fmt.Println(ui16)
	fmt.Println(ui32)
	fmt.Println(ui64)
	fmt.Println(i)
	fmt.Println(ui)
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(c64)
	fmt.Println(c128)
	fmt.Println(s)
	fmt.Println(by)
	fmt.Println(r)
}
