./compile.sh && wasmtime out.wasm --invoke main

#  -D   debug-info[=y|n] -- Enable generation of DWARF debug information in compiled code.
#  -D  address-map[=y|n] -- Configure whether compiled code can map native addresses to wasm.
#  -D      logging[=y|n] -- Configure whether logging is enabled.
#  -D log-to-files[=y|n] -- Configure whether logs are emitted to files
#  -D       coredump=val -- Enable coredump generation to this file after a WebAssembly trap.