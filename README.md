# SpecOps [![Go](https://github.com/solidifylabs/specops/actions/workflows/go.yml/badge.svg)](https://github.com/solidifylabs/specops/actions/workflows/go.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/solidifylabs/specops.svg)](https://pkg.go.dev/github.com/solidifylabs/specops)

> `specops` is a low-level, domain-specific language (and compiler) for crafting [Ethereum VM](https://ethereum.org/en/developers/docs/evm) bytecode in Go.

This is a _very_ early release. In fact, it's just a weekend project gone rogue
so is less than a week old.

## _special_ opcodes

Writing bytecode is hard. There's always that temptation to give up and use a
higher-level language with all of its conveniences, but that defeats the point.
What if we could maintain full control of the opcode placement, but with
syntactic sugar to help the medicine go down?

*Special* opcodes provide just that. Some are interpreted by the compiler,
converting them to
[regular](https://pkg.go.dev/github.com/ethereum/go-ethereum/core/vm#OpCode)
equivalents, while others are simply compiler hints that leave the resulting
bytecode unchanged.

## Features

- [x] `JUMPDEST` labels (absolute)
- [ ] `JUMPDEST` labels (relative to `PC`)
- [x] Function-like syntax (optional)
- [x] Inverted `DUP`/`SWAP` special opcodes from "bottom" of stack (a.k.a. pseudo-variables)
- [x] `PUSH<T>` for native Go types
- [X] `PUSH(v)` length detection
- [x] Macros
- [x] Compiler-state assertions (e.g. expected stack depth)
- [ ] Automatic stack permutation
- [ ] Standalone compiler
- [x] In-process EVM execution (geth)
- [x] Debugger
  * [x] Stepping
  * [ ] Breakpoints
  * [x] Programmatic inspection (e.g. native Go tests at opcode resolution)
    * [x] Memory
    * [x] Stack
  * [ ] User interface
- [ ] Source mapping
- [ ] Coverage analysis
- [ ] Fork testing with RPC URL

### Documentation

The [`specops` Go
documentation](https://pkg.go.dev/github.com/solidifylabs/specops) covers all
functionality.

## Examples

### Hello world

TODO: link to Go playground; for now, here's the [real implementation](https://github.com/solidifylabs/specops/blob/41efe932c9a85e45ce705b231577447e6c944487/examples_test.go#L12).

The `specops` Go package has a minimal footprint to allow for dot-importing,
making all exported symbols available. TODO: expand on the implications,
rationale, and recommendations as this goes against the style guide.

```go
import . github.com/solidifylabs/specops

…

hello := []byte("Hello world")
code := Code{
    // The compiler determines the shortest-possible PUSH<n> opcode.
    // Fn() simply reverses its arguments (a surprisingly powerful construct)!
    Fn(MSTORE, PUSH0, PUSH(hello)),
    Fn(RETURN, PUSH(32-len(hello)), PUSH(len(hello))),
}

// ----- COMPILE -----
bytecode, err := code.Compile()
// ...

// ----- EXECUTE -----

result, err := code.Run(nil /*callData*/ /*, [runopts.Options]...*/)
// ...

// ----- DEBUG (Programmatic) -----

dbg, results := code.StartDebugging(nil /*callData*/ /*, Options...*/)
defer dbg.FastForward() // best practice to avoid resource leaks

state := dbg.State() // is updated on calls to Step() / FastForward()

for !dbg.Done() {
  dbg.Step()
  fmt.Println("Peek-a-boo", state.ScopeContext.Stack().Back(0))
}

result, err := results()
//...
```

### Other examples

- [Verbatim reimplementation of well-known contracts](https://github.com/solidifylabs/specops/blob/41efe932c9a85e45ce705b231577447e6c944487/examples_test.go#L34) 
  * EIP-1167 Minimal Proxy ([original](https://eips.ethereum.org/EIPS/eip-1167#specification))
  * 0age/metamorphic ([original](https://github.com/0age/metamorphic/blob/55adac1d2487046002fc33a5dff7d669b5419a3a/contracts/MetamorphicContractFactory.sol#L55))
- [Monte Carlo approximation of pi](https://github.com/solidifylabs/specops/blob/41efe932c9a85e45ce705b231577447e6c944487/examples_test.go#L158)
- [`sqrt()`](https://github.com/solidifylabs/specops/blob/41efe932c9a85e45ce705b231577447e6c944487/examples_test.go#L246) as seen ~~on TV~~ in `prb-math` ([original](https://github.com/PaulRBerg/prb-math/blob/5b6279a0cf7c1b1b6a5cc96082811f7ef620cf60/src/Common.sol#L595))

## Acknowledgements

Some of SpecOps was, of course, inspired by
[Huff](https://github.com/huff-language). I hope to provide something different,
of value, and to inspire them too.
