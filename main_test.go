package main

import (
	"fmt"
	"testing"
)

func TestBinaryRep(t *testing.T) {
	// fmt.Println(strconv.FormatInt(
	// 	('1' - '0') << 7, 2,
	// ))
	// fmt.Println(strconv.FormatInt(200, 2))

	// fmt.Println(leb128_u32(624485))

}

func TestLEB128_32(t *testing.T) {

	fmt.Printf("%#v\n", LEB128_U32(624485))


	fmt.Printf("%#v\n", LEB128_I32(-123456))


}
