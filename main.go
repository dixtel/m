package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strconv"
)

const (
	VEC_TYPE = 0x7B
)

type SectionId = byte

const (
	SECTION_CUSTOM_ID     SectionId = byte(0)
	SECTION_TYPE_ID       SectionId = byte(1)
	SECTION_IMPORT_ID     SectionId = byte(2)
	SECTION_FUNCTION_ID   SectionId = byte(3)
	SECTION_TABLE_ID      SectionId = byte(4)
	SECTION_MEMORY_ID     SectionId = byte(5)
	SECTION_GLOBAL_ID     SectionId = byte(6)
	SECTION_EXPORT_ID     SectionId = byte(7)
	SECTION_START_ID      SectionId = byte(8)
	SECTION_ELEMENT_ID    SectionId = byte(9)
	SECTION_CODE_ID       SectionId = byte(10)
	SECTION_DATA_ID       SectionId = byte(11)
	SECTION_DATA_COUNT_ID SectionId = byte(12)
)

const (
	TYPE_FUNCTION = 0x60
)

// https://stackoverflow.com/questions/68137460/how-to-write-leb128-in-go
func uleb128(x uint64) []byte {
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint64(bs, x)
	return bs
}

func leb128_u32(x int) []byte {
	println("leb128 of: ", x)

	if x < 0 {
		panic("want unsigned u32")
	}

	bin := strconv.FormatInt(int64(x), 2)
	pad := 7 - len(bin)%7
	for pad > 0 {
		bin = "0" + bin
		pad -= 1
	}

	bytes := []string{}
	for i := len(bin); i > 0; i -= 7 {
		b := bin[i-7 : i]
		bytes = append(bytes, b)
	}

	res := ""
	for i := 0; i < len(bytes)-1; i++ {
		res += "1" + bytes[i]
	}

	bin = res + "0" + bytes[len(bytes)-1]

	if math.Ceil(float64(len(bin))/float64(7)) > 5 {
		panic("The total number of bytes encoding a value of type uN must not exceed ceil(N/7) bytes")
	}

	print("leb128 bin: ")

	for i := 0; i < len(bin); i += 8 {
		print(bin[i:i+8], " ")
	}
	println()

	ret := []byte{}

	for i := 0; i < len(bin); i += 8 {
		b := bin[i : i+8]
		n := byte(0)

		for i := 0; i < 8; i++ {
			n <<= 1
			n |= byte(b[i] - '0')
		}

		ret = append(ret, n)
	}

	fmt.Printf("leb128 hex: % X \n", ret)

	return ret
}

// Source: https://en.wikipedia.org/wiki/LEB128
func LEB128_U32(val uint32) (ret []byte) {
	for {
		var b = byte(val) & 0x7F // Get 7 low-order bits of `val`
		val >>= 7

		if val != 0 {
			b |= 0x80 // Set the most significant bit if the next byte is non-zero
		}

		ret = append(ret, b)

		if val == 0 {
			return
		}
	}
}

// Source: https://en.wikipedia.org/wiki/LEB128
func LEB128_I32(val int32) (ret []byte) {
	for {
		var b = byte(val) & 0x7F // Get 7 low-order bits of `val`
		val >>= 7

		if (val == 0 && (b&0x40) == 0) || (val == -1 && (b&0x40) != 0) {
			ret = append(ret, b)
			return
		}

		ret = append(ret, b|0x80)
	}
}

func section(sectionId SectionId, contents []byte) (ret []byte) {
	ret = append(ret, sectionId)
	ret = append(ret, leb128_u32(len(contents))...)
	ret = append(ret, contents...)
	return
}

func spread(bs ...[]byte) (ret []byte) {
	for _, b := range bs {
		ret = append(ret, b...)
	}
	return
}

func vec(length int, c []byte) (ret []byte) {
	ret = append(ret, LEB128_U32(uint32(length))...)
	ret = append(ret, c...)
	return
}

type ImportFunc struct {
	mod     string
	nm      string
	typeidx int
}

func importSection(imports []ImportFunc) []byte {
	importBinary := []byte{}
	for _, i := range imports {
		importBinary = append(importBinary, vec(len(i.mod), []byte(i.mod))...)
		importBinary = append(importBinary, vec(len(i.nm), []byte(i.nm))...)
		importBinary = append(importBinary, 0x00)
		importBinary = append(importBinary, LEB128_U32(uint32(i.typeidx))...)
	}

	return section(SECTION_IMPORT_ID, vec(len(imports), importBinary))
}

func main() {
	wasm := []byte{
		0x00, 0x61, 0x73, 0x6D,
		0x01, 0x00, 0x00, 0x00,
	}

	wasm = append(wasm, section(
		SECTION_TYPE_ID,
		vec(
			2,
			spread(
				functionType(
					[]ValueType{},
					[]ValueType{VALUE_TYPE_I32},
				),
				functionType( // fd_write
					[]ValueType{
						VALUE_TYPE_I32,
						VALUE_TYPE_I32,
						VALUE_TYPE_I32,
						VALUE_TYPE_I32,
					},
					[]ValueType{
						VALUE_TYPE_I32,
					},
				),
			),
		),
	)...)

	wasm = append(wasm, importSection(
		[]ImportFunc{
			{
				mod:     "wasi_snapshot_preview1",
				nm:      "fd_write",
				typeidx: 1,
			},
		},
	)...)

	wasm = append(wasm, section(
		SECTION_FUNCTION_ID,
		spread(
			leb128_u32(1),
			INDICE_TYPE_IDX(0),
		),
	)...)

	wasm = append(wasm, section(
		SECTION_TABLE_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	wasm = append(wasm, section(
		SECTION_MEMORY_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	wasm = append(wasm, section(
		SECTION_GLOBAL_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	wasm = append(wasm, section(
		SECTION_EXPORT_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	// wasm = append(wasm, section(
	// 	SECTION_START_ID,
	// 	spread(
	// 	leb128_u32(0),
	// 	),
	// )...)

	wasm = append(wasm, section(
		SECTION_ELEMENT_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	wasm = append(wasm, section(
		SECTION_DATA_COUNT_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	/*
		codesec ::= code*:section10(vec(code)) ⇒ code*
		code    ::= size:u32 code:func         ⇒ code (if size = ||func||)
		func    ::= (𝑡*)*:vec(locals) 𝑒:expr    ⇒ concat((𝑡*)*), 𝑒 (if |concat((𝑡*)*)| < 232)
		locals  ::= 𝑛:u32 𝑡:valtype             ⇒ 𝑡𝑛


		// https://github.com/WebAssembly/design/issues/1037
		local variables count include function arguments!
		total locals = function arguments + local variables.
		first local variable index = num function arguments + 0
		second local variable index = num function arguments + 1

		2.5.3 Functions
	*/
	wasm = append(wasm, section(
		SECTION_CODE_ID,
		spread(
			leb128_u32(1), // Number of code entries
			function(
				[]FunctionLocal{ // Length of locals + locals
					{
						count:   2,
						valtype: VALUE_TYPE_I32,
					},
				},
				spread( // function body
					[]byte{INSTR_I32_CONST},
					LEB128_I32(1),
					[]byte{INSTR_I32_CONST},
					LEB128_I32(2),
					[]byte{INSTR_I32_ADD},
					[]byte{INSTR_END_MARKER},
				),
			),
		),
	)...)

	wasm = append(wasm, section(
		SECTION_DATA_ID,
		spread(
			leb128_u32(0),
		),
	)...)

	save(wasm)
}

func save(w []byte) {
	f, e := os.OpenFile("out.wasm", os.O_CREATE|os.O_TRUNC|os.O_APPEND|os.O_WRONLY, 0o644)
	if e != nil {
		panic(e)
	}
	defer f.Close()
	_, e = f.Write(w)
	if e != nil {
		panic(e)
	}
}

const (
	INSTR_END_MARKER = 0x0B

	// Variable instructions
	INSTR_LOCAL_GET = 0x20
	INSTR_LOCAL_SET = 0x21

	// Numeric Instructions
	INSTR_I32_CONST = 0x41

	INSTR_I32_ADD = 0x6A
)

type FunctionLocal struct {
	count   int // u32
	valtype ValueType
}

func function(locals []FunctionLocal, expr []byte) []byte {
	ret := leb128_u32(len(locals))

	for _, local := range locals {
		ret = append(ret, leb128_u32(local.count)...)
		ret = append(ret, byte(local.valtype))
	}

	ret = append(ret, expr...)

	ret = append(leb128_u32(len(ret)), ret...) // push front function size

	return ret
}

type ValueType byte

const (
	VALUE_TYPE_I32 ValueType = 0x7F
)

func functionType(r1 []ValueType, r2 []ValueType) []byte {
	ret := []byte{TYPE_FUNCTION}

	ret = append(ret, leb128_u32(len(r1))...)
	for _, vt := range r1 {
		ret = append(ret, byte(vt))
	}

	ret = append(ret, leb128_u32(len(r2))...)
	for _, vt := range r2 {
		ret = append(ret, byte(vt))
	}

	return ret
}

var INDICE_TYPE_IDX = leb128_u32
