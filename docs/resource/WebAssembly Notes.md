https://webassembly.github.io/wabt/demo/wasm2wat/
https://charlycst.github.io/posts/wasm-encoding/
https://developer.mozilla.org/en-US/docs/WebAssembly/


‎__wasi_ciovec_t https://github.com/WebAssembly/wasi-libc/blob/ebac9aee23f28aaddf17507712696f56727513d4/libc-bottom-half/headers/public/wasi/api.h#L697 

# Modules
[[WebAssembly.pdf#page=159|WebAssembly, page 159]]
	- [[WebAssembly.pdf#page=157&selection=6,7,6,19&color=yellow|Code Section]]
section encoding: [[WebAssembly.pdf#search=5.5.2 Sections|5.5.2 Sections]]

# Function
## Locals

> ([[WebAssembly.pdf#page=24&selection=133,0,138,111&color=note|p.20]])
> The locals declare a vector of mutable local variables and their types. These variables are referenced through local indices in the function’s body. The index of the first local is the smallest index not referencing a parameter.

## Encoding
[[WebAssembly.pdf#page=24&selection=92,0,100,29|The funcs component of a module defines a vector of functions with the following structure:]]


# LEB128
https://en.wikipedia.org/wiki/LEB128

[[WebAssembly.pdf#page=139&selection=32,0,33,2&color=note|All integers are encoded using the LEB128]]

##  encode unsigned integer
```
val: uint32 = ...

do {
	byte = val & 0x7F
	val >>= 7

	if val != 0 {
		byte |= 0x80 // set most siginificant bit to 1
	}

	emit byte
} while (val != 0)



```

The sign does not matter for addition
https://stackoverflow.com/a/63454934

## ones' complement and two's ceomplement
1. positive number, most significant bit must be  0, rest is just a binary number
2. if (one's complement) if negative
	1. represent number in binary (-3) => 3
	2. flip the bits
	3. we now have number in ones' complements when we have -0 and 0
3. if (two's complement) if negative
	1. 1. represent number in binary (-3) => 3
	2. flip the bits
	3. add 1
	4. we have number in two's complements with only one 0 represenation
