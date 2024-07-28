*reference* https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing
## NUD *(null denotation)*: prefix 
This is the atomic token like -2 or 2
## LED  *(left denotation)*: infix
Binary operator like +, -
## LBP *(left binding power value)*
Can be called also as infix binding power


## Pseudocode
```
fun parse {
	expr(tokens: $tokens, rbp: 0)
}
	
fun expr {
	let tokens: i32 $tokens
	let rbp:    i32 $rbb
	
	let left NUD(next())
	while RBP(peek()) > rbp {
		let left LED(
		    peek(), 
			[left, expr(tokens: $tokens, rbp: RBP(next())]
		)
	}
	
	ret left
}

2 + 2 * 3
  1   2

step 1:
	NUD = 2
	RPB(+) < rbp => 1 < 0

```

# How to handle parenthesis
```
let x 1 
let y (1 + 2)
let x (1 + 2 * (1 - 5 - -5) / 15)
```