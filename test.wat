(module
  (type (;0;) (func))
  (import "wasi_snapshot_preview1" "fd_write" (func $fd_write (param i32 i32 i32 i32) (result i32)))
  (memory (;0;) 1)
  (export "memory" (memory 0))
  (export "_start" (func $start))
  (func $start (type 0)
    i32.const 1 ;; stdout
    i32.const 0 ;; iovec ptr
    i32.const 1 ;; entries
    i32.const 24 ;; out bytes
    call $fd_write
    drop
  )
  (data (i32.const 0) "BCD")
)