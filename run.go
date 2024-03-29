package specops

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
	"github.com/solidifylabs/specops/evmdebug"
	"github.com/solidifylabs/specops/runopts"
)

// Run calls c.Compile() and runs the compiled bytecode on a freshly
// instantiated vm.EVMInterpreter. The default EVM parameters MUST NOT be
// considered stable: they are currently such that code runs on the Cancun fork
// with no state DB.
func (c Code) Run(callData []byte, opts ...runopts.Option) ([]byte, error) {
	compiled, err := c.Compile()
	if err != nil {
		return nil, fmt.Errorf("%T.Compile(): %v", c, err)
	}
	return runBytecode(compiled, callData, opts...)
}

// StartDebugging appends a runopts.Debugger (`dbg`) to the Options, calls
// c.Run() in a new goroutine, and returns `dbg` along with a function to
// retrieve ther esults of Run(). The function will block until Run() returns,
// i.e. when dbg.Done() returns true. There is no need to call dbg.Wait().
//
// If execution never completes, such that dbg.Done() always returns false, then
// the goroutine will be leaked.
//
// Any compilation error will be returned by StartDebugging() while execution
// errors are returned by a call to the returned function. Said execution errors
// can be errors.Unwrap()d to access the same error available in
// `dbg.State().Err`.
func (c Code) StartDebugging(callData []byte, opts ...runopts.Option) (*evmdebug.Debugger, func() ([]byte, error), error) {
	compiled, err := c.Compile()
	if err != nil {
		return nil, nil, fmt.Errorf("%T.Compile(): %v", c, err)
	}

	dbg, opt := runopts.WithNewDebugger()
	opts = append(opts, opt)

	var (
		result []byte
		resErr error
	)
	done := make(chan struct{})
	go func() {
		result, resErr = runBytecode(compiled, callData, opts...)
		close(done)
	}()

	dbg.Wait()

	return dbg, func() ([]byte, error) {
		<-done
		return result, resErr
	}, nil
}

// RunTerminalDebugger is equivalent to StartDebugging(), but instead of
// returning the Debugger and results function, it calls
// Debugger.RunTerminalUI().
func (c Code) RunTerminalDebugger(callData []byte, opts ...runopts.Option) error {
	var contract *vm.Contract
	opts = append(opts, runopts.Func(func(c *runopts.Configuration) error {
		contract = c.Contract
		return nil
	}))
	dbg, results, err := c.StartDebugging(callData, opts...)
	if err != nil {
		return err
	}
	defer dbg.FastForward()
	return dbg.RunTerminalUI(callData, results, contract)
}

func runBytecode(compiled, callData []byte, opts ...runopts.Option) ([]byte, error) {
	cfg, err := newRunConfig(compiled, opts...)
	if err != nil {
		return nil, err
	}
	interp := vm.NewEVM(
		cfg.BlockCtx,
		cfg.TxCtx,
		cfg.StateDB,
		cfg.ChainConfig,
		cfg.VMConfig,
	).Interpreter()

	out, err := interp.Run(cfg.Contract, callData, cfg.ReadOnly)
	if err != nil {
		return nil, fmt.Errorf("%T.Run([%T.Compile()], [callData], readOnly=%t): %w", interp, Code{}, cfg.ReadOnly, err)
	}
	return out, nil
}

func newRunConfig(compiled []byte, opts ...runopts.Option) (*runopts.Configuration, error) {
	cfg := &runopts.Configuration{
		Contract: &vm.Contract{
			Code: compiled,
			Gas:  30e6,
		},
		BlockCtx: vm.BlockContext{
			BlockNumber: big.NewInt(0),
			Random:      &common.Hash{}, // post merge
		},
		ChainConfig: &params.ChainConfig{
			LondonBlock: big.NewInt(0),
			CancunTime:  new(uint64),
		},
	}
	for _, o := range opts {
		if err := o.Apply(cfg); err != nil {
			return nil, fmt.Errorf("runopts.Option[%T].Apply(): %v", o, err)
		}
	}
	return cfg, nil
}
