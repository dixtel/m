(module
(;@b     ;)  (type (;0;) (func))
(;@e     ;)  (type (;1;) (func (param i32 i32 i32 i32) (result i32)))
(;@19    ;)  (import "wasi_snapshot_preview1" "fd_write" (func (;0;) (type 1)))
(;@64    ;)  (func (;1;) (type 0)
(;@65    ;)    i32.const 0
(;@67    ;)    i32.const 100
(;@6a    ;)    i32.store
(;@6d    ;)    i32.const 4
(;@6f    ;)    i32.const 11
(;@71    ;)    i32.store
(;@74    ;)    i32.const 1
(;@76    ;)    i32.const 0
(;@78    ;)    i32.const 1
(;@7a    ;)    i32.const 16
(;@7c    ;)    call 0
(;@7e    ;)    drop
             )
(;@45    ;)  (memory (;0;) 1)
(;@4d    ;)  (export "main" (func 1))
(;@54    ;)  (export "memory" (memory 0))
(;@83    ;)  (data (;0;) (i32.const 100) "hello world")
           )
(;@94    ;)
