# Mart Language
Mart is a simple programming language, similar to Go.


- Max 80 words per line

# Reserved keywords
- let
- if
- elif
- else
- for
- fun
- enum
- data
- pub

# builtin types
- Num (64 double precision, IEEE 754-2008)
- - converted to integer 32 when performing bitwise operations
- Str (utf-8)

# Basic types

Builtin basic types:
- num (can be i32, i64, f32 or f64)
- str

```
let x     1
let price 12.99
let hello "world"    

```

# Functions

```
let x      do_db_call() // propagate error to the parent in case of a any error
let x, err do_db_call() // error will be handled manually


fun do_db_call() {
    validate_sql($sql) // error will be propgeted to the caller

    res, err call_db($sql)

    err ? ret fmt("cannot make db call: {}", err)
    // sugar syntax to:
    if err {
        ret fmt("cannot make db call: {}", err)        
    }

    ret res
}


```

# Comples types

Builtin comples types:
- data
- dyn

```

data Book {
    name   str
    author str
    year   num
}

fun Book.get_name() {
    ret $self.name
}

let my_fav_book Book (
    name   "1984"
    author "George Orwell"
    year   1949
)

dyn GetName {
    get_name() str
}

let name_getter:     GetName my_fav_book 
let name_of_the_book         name_getter.get_name()

```


# Declaring variables

```
// global functions must be upper case
let TOKEN_ID "abcd"
// upper case variables are automatically set as connstant

let y 1
let x 2

let x "234" // variable can be shadowed

// Math on variables

let number1 4
let num2    8
let res1    (number1 + num2) // parenthsis are required when assigned value consists with more than 1 element
let res2    (number1 - num2)
let res3    (number1 * num2)
let res4    (number1 / num2)
```

# Control Flow
```
if res1 == 4 {

} elif res1 < 42 {

} elif res1 > 42 {

} else {

}
```

# Loops
```
for bytes {
    let idx $idx
    let obj $elm
}
```

# Functions
- annotation for return types are required


```

fun sum_two_numbers {
    let num1 $num1
    let num2 $num2
    ret num1 + num2
}

fun sum_two_numbers -> num {
    let num1 $num1
    let num2 $num2
    ret num1 + num2
}

fun sum_numbers -> num {
    let sum 0
    for $ {
        let idx $idx
        let obj $elm
        sum sum + obj
    }
    ret sum
}

let x sum_two_numbers(num1 2, num2 4)
let x sum_two_numbers(
    num1 2
    num2 4
)

let y sum_numbers(1, 2, 3, 4, 5)
let y sum_numbers(
    1
    2
    3
    4
    5
)

fun print_owner_dogs {
    let owner_name $owner_name 
    
    // $ = print_owner_dogs_args
    // print_owner_dogs_kwargs = { owner_name str }
    // print_owner_dogs_args = [{idx, elm}, ...]
    
    for $ {
        print(format "{} has dog with name {}", owner_name, $elm)
    }
    
    // for loop above is a sugar syntax for:

    for $ {
        print(format "{} has dog with name {}", owner_name, $elm)
    }
}

fun iter {
    let arr $arr
    let cb $cb



}


let owner_name "Marian"

print_owner_dogs(owner_name owner_name, "Alex", "Fredo", "Melisa")
// the same as
print_owner_dogs(owner_name, "Alex", "Fredo", "Melisa")




```
# Enums - no booleans, only enums
```
// old approach:
// is_object_exists = true

// new approach:
let Object enum {
	Exists
	NotExists
}

// defining an enum
enum ConnStatus { 
    NotConnected // first variant is always a default one
    Connected
    Disconnected
} 

// enum can by anonymous
let Status enum { 
    NotStarted
    Started
}

// enum can hold any data
let Player enum { 
    Player { 
	    nickname: str
	    exp:      num
	}
    PlayerNotFound
}

// in stdlib
@allow_snake_case(Type) a 
@allow_snake_case(Variant)
let bool enum {
	true
	false
}


let player enum { 
    Player {
        nickname: str
        exp:      num
        active:   bool
    }
    PlayerNotFound
}
```

# Data
Data are like structs in other programming languages

## Defining
```
// anonymous,
// i don’t think it is a good idea
// there should be only one way of doing things
let player {
    name       "banana"
    exp        0
    created_at "12:12"
    updated_at "12:13"
}

data Player {
    name              str
    exp               num
    created_at        str
    updated_at        str
    // this field will not be exposed
    // to external modules
    _some_private_var [num]
}

// this struct will not be exposed 
// to external modules
data _PrivateStruct {}

let player1 Player

// we can override default values 
let player2 Player(
    name       "player2"
    created_at time.now()
    updated_at time.now()
)

```
## Anonymous data
```
let player_data {
    name       "banana"
    exp        0
    created_at
    updated_at
}

// some code...

// updating player details
player_data.exp        2
player_data.updated_at time.now()
```
## Data type
```
// error: data types must be specified
data PlayerData {
    name       
    exp
    created_at
    updated_at
}

let player_data {
    name       "banana"
    exp        0
    created_at
    updated_at
}

// some code...

// updating player details
player_data.exp        2
player_data.updated_at time.now()
```

## Adding methods to data
```
data PlayerData {
    name       
    exp
    created_at
    updated_at
}

fun PlayerData.new {
    ret PlayerData(name $name, exp 0, created_at time.now())
}

fun PlayerData.change_exp {
    if $new_exp > 100 {
        panic("exp limit exceeded")
    }

    $self.exp $new_exp
}

```

# What if we exceed 80 line size limit
```
some_struct_with_very_long_long_long_variable.and_this_is_very_long_long_long_field_of_this_struct time.now()

// The above code will not compile

// To decrease line size we can break up the line like following:
some_struct_with_very_long_long_long_variable.
    and_this_is_very_long_long_long_field_of_this_struct time.now()

// or like this
some_struct_with_very_long_long_long_variable.
    and_this_is_very_long_long_long_field_of_this_struct
    time.now()

// other examples

player_data.created_at
    time.now()

```


