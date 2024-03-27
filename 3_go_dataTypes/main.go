package main

import (
	"fmt"
	"math/bits"
	"unsafe"
)

func main() {
	var num int
	var num16 int16
	var num32 int32
	var num64 int64

	fmt.Printf("num  %d Bytes \n", unsafe.Sizeof(num))
	fmt.Printf("num16  %d Bytes \n", unsafe.Sizeof(num16))
	fmt.Printf("num32  %d Bytes \n", unsafe.Sizeof(num32))
	fmt.Printf("num64  %d Bytes \n", unsafe.Sizeof(num64))
	// --- float ---
	var b float32
	fmt.Printf("number %f \n", b)
	// --- def sizing ---
	var a = bits.UintSize // size default sys
	fmt.Println(a)	
}	
	
