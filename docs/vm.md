# VM - Context
## Context.Consts
- Numbers
- Literals
## Context.Dynamic
- Objects

# OpCodes
- OpLoadConst  usize <- Context.Consts
- OpLoadName   usize <- Context.Dynamic
- OpBinaryAdd
- OpBinarySub
- OpBinaryMul
- OpBinaryDiv
- OpStoreName  usize -> Context.Dynamic