- Gradual typing


# builtin types
- i32, i64
- f32, f64
- str
- dyn


# Lang Spec


```
let_statement ::= "let" ident " " expr

ident ::= alphabetic alphanumeric* 

expr ::= ident
       | number
       | fun_call
       | "(" expr (op expr)* ")"

fun_call :== ident "(" + ...


```


# Modules
Packing system is similar to golang's packages.

```
main/                    // module mart
	main.mart
	engine/              // submodule engine
		engine.mart
		parts.mart
	helpers/             // submodule helpers
		helpers.mart
		utils/           // submodule utils
			utils.mart
```

`pub` keyword allow to access `data` and `fun` from external module. Otherwise all definitions inside module is private by default

# Macro

# Private members