```
fun add: I32 {
	let a: I32 $a
	let b: I32 $b
	ret a + b
}

let x add {a: 1, b: 2}
let x add(a: 1, b: 2)
let x add(
	a: 1
	b: 2
)

```
