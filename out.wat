(module
  (type (;0;) (func (result i32)))
  (type (;1;) (func (param i32 i32 i32 i32) (result i32)))
  (import "wasi_snapshot_preview1" "fd_write" (func (;0;) (type 1)))
  (func (;1;) (type 0) (result i32)
    (local i32 i32)
    i32.const 1
    i32.const 2
    i32.add
  )
)