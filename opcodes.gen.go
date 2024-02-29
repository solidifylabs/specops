package specialops

//
// GENERATED CODE - DO NOT EDIT
//

import "github.com/ethereum/go-ethereum/core/vm"

// Aliases of all regular vm.OpCode constants that don't have "special" replacements.
const (
	STOP = opCode(vm.STOP)
	ADD = opCode(vm.ADD)
	MUL = opCode(vm.MUL)
	SUB = opCode(vm.SUB)
	DIV = opCode(vm.DIV)
	SDIV = opCode(vm.SDIV)
	MOD = opCode(vm.MOD)
	SMOD = opCode(vm.SMOD)
	ADDMOD = opCode(vm.ADDMOD)
	MULMOD = opCode(vm.MULMOD)
	EXP = opCode(vm.EXP)
	SIGNEXTEND = opCode(vm.SIGNEXTEND)
	LT = opCode(vm.LT)
	GT = opCode(vm.GT)
	SLT = opCode(vm.SLT)
	SGT = opCode(vm.SGT)
	EQ = opCode(vm.EQ)
	ISZERO = opCode(vm.ISZERO)
	AND = opCode(vm.AND)
	OR = opCode(vm.OR)
	XOR = opCode(vm.XOR)
	NOT = opCode(vm.NOT)
	BYTE = opCode(vm.BYTE)
	SHL = opCode(vm.SHL)
	SHR = opCode(vm.SHR)
	SAR = opCode(vm.SAR)
	KECCAK256 = opCode(vm.KECCAK256)
	ADDRESS = opCode(vm.ADDRESS)
	BALANCE = opCode(vm.BALANCE)
	ORIGIN = opCode(vm.ORIGIN)
	CALLER = opCode(vm.CALLER)
	CALLVALUE = opCode(vm.CALLVALUE)
	CALLDATALOAD = opCode(vm.CALLDATALOAD)
	CALLDATASIZE = opCode(vm.CALLDATASIZE)
	CALLDATACOPY = opCode(vm.CALLDATACOPY)
	CODESIZE = opCode(vm.CODESIZE)
	CODECOPY = opCode(vm.CODECOPY)
	GASPRICE = opCode(vm.GASPRICE)
	EXTCODESIZE = opCode(vm.EXTCODESIZE)
	EXTCODECOPY = opCode(vm.EXTCODECOPY)
	RETURNDATASIZE = opCode(vm.RETURNDATASIZE)
	RETURNDATACOPY = opCode(vm.RETURNDATACOPY)
	EXTCODEHASH = opCode(vm.EXTCODEHASH)
	BLOCKHASH = opCode(vm.BLOCKHASH)
	COINBASE = opCode(vm.COINBASE)
	TIMESTAMP = opCode(vm.TIMESTAMP)
	NUMBER = opCode(vm.NUMBER)
	DIFFICULTY = opCode(vm.DIFFICULTY)
	GASLIMIT = opCode(vm.GASLIMIT)
	CHAINID = opCode(vm.CHAINID)
	SELFBALANCE = opCode(vm.SELFBALANCE)
	BASEFEE = opCode(vm.BASEFEE)
	BLOBHASH = opCode(vm.BLOBHASH)
	BLOBBASEFEE = opCode(vm.BLOBBASEFEE)
	POP = opCode(vm.POP)
	MLOAD = opCode(vm.MLOAD)
	MSTORE = opCode(vm.MSTORE)
	MSTORE8 = opCode(vm.MSTORE8)
	SLOAD = opCode(vm.SLOAD)
	SSTORE = opCode(vm.SSTORE)
	JUMP = opCode(vm.JUMP)
	JUMPI = opCode(vm.JUMPI)
	PC = opCode(vm.PC)
	MSIZE = opCode(vm.MSIZE)
	GAS = opCode(vm.GAS)
	TLOAD = opCode(vm.TLOAD)
	TSTORE = opCode(vm.TSTORE)
	MCOPY = opCode(vm.MCOPY)
	PUSH0 = opCode(vm.PUSH0)
	DUP1 = opCode(vm.DUP1)
	DUP2 = opCode(vm.DUP2)
	DUP3 = opCode(vm.DUP3)
	DUP4 = opCode(vm.DUP4)
	DUP5 = opCode(vm.DUP5)
	DUP6 = opCode(vm.DUP6)
	DUP7 = opCode(vm.DUP7)
	DUP8 = opCode(vm.DUP8)
	DUP9 = opCode(vm.DUP9)
	DUP10 = opCode(vm.DUP10)
	DUP11 = opCode(vm.DUP11)
	DUP12 = opCode(vm.DUP12)
	DUP13 = opCode(vm.DUP13)
	DUP14 = opCode(vm.DUP14)
	DUP15 = opCode(vm.DUP15)
	DUP16 = opCode(vm.DUP16)
	SWAP1 = opCode(vm.SWAP1)
	SWAP2 = opCode(vm.SWAP2)
	SWAP3 = opCode(vm.SWAP3)
	SWAP4 = opCode(vm.SWAP4)
	SWAP5 = opCode(vm.SWAP5)
	SWAP6 = opCode(vm.SWAP6)
	SWAP7 = opCode(vm.SWAP7)
	SWAP8 = opCode(vm.SWAP8)
	SWAP9 = opCode(vm.SWAP9)
	SWAP10 = opCode(vm.SWAP10)
	SWAP11 = opCode(vm.SWAP11)
	SWAP12 = opCode(vm.SWAP12)
	SWAP13 = opCode(vm.SWAP13)
	SWAP14 = opCode(vm.SWAP14)
	SWAP15 = opCode(vm.SWAP15)
	SWAP16 = opCode(vm.SWAP16)
	LOG0 = opCode(vm.LOG0)
	LOG1 = opCode(vm.LOG1)
	LOG2 = opCode(vm.LOG2)
	LOG3 = opCode(vm.LOG3)
	LOG4 = opCode(vm.LOG4)
	CREATE = opCode(vm.CREATE)
	CALL = opCode(vm.CALL)
	CALLCODE = opCode(vm.CALLCODE)
	RETURN = opCode(vm.RETURN)
	DELEGATECALL = opCode(vm.DELEGATECALL)
	CREATE2 = opCode(vm.CREATE2)
	STATICCALL = opCode(vm.STATICCALL)
	REVERT = opCode(vm.REVERT)
	INVALID = opCode(vm.INVALID)
	SELFDESTRUCT = opCode(vm.SELFDESTRUCT)
)

// stackDeltas maps all valid vm.OpCode values to the number of values they
// pop and then push from/to the stack.
//
// Although DUPs technically only push a single value and SWAPs none, they are
// recorded as popping and pushing one more than they actually do as this
// implies a minimum stack depth to begin with but with the same effective
// change.
var stackDeltas = map[vm.OpCode]stackDelta{
	vm.STOP: {pop: 0, push: 0},
	vm.ADD: {pop: 2, push: 1},
	vm.MUL: {pop: 2, push: 1},
	vm.SUB: {pop: 2, push: 1},
	vm.DIV: {pop: 2, push: 1},
	vm.SDIV: {pop: 2, push: 1},
	vm.MOD: {pop: 2, push: 1},
	vm.SMOD: {pop: 2, push: 1},
	vm.ADDMOD: {pop: 3, push: 1},
	vm.MULMOD: {pop: 3, push: 1},
	vm.EXP: {pop: 2, push: 1},
	vm.SIGNEXTEND: {pop: 2, push: 1},
	vm.LT: {pop: 2, push: 1},
	vm.GT: {pop: 2, push: 1},
	vm.SLT: {pop: 2, push: 1},
	vm.SGT: {pop: 2, push: 1},
	vm.EQ: {pop: 2, push: 1},
	vm.ISZERO: {pop: 1, push: 1},
	vm.AND: {pop: 2, push: 1},
	vm.OR: {pop: 2, push: 1},
	vm.XOR: {pop: 2, push: 1},
	vm.NOT: {pop: 1, push: 1},
	vm.BYTE: {pop: 2, push: 1},
	vm.SHL: {pop: 2, push: 1},
	vm.SHR: {pop: 2, push: 1},
	vm.SAR: {pop: 2, push: 1},
	vm.KECCAK256: {pop: 2, push: 1},
	vm.ADDRESS: {pop: 0, push: 1},
	vm.BALANCE: {pop: 1, push: 1},
	vm.ORIGIN: {pop: 0, push: 1},
	vm.CALLER: {pop: 0, push: 1},
	vm.CALLVALUE: {pop: 0, push: 1},
	vm.CALLDATALOAD: {pop: 1, push: 1},
	vm.CALLDATASIZE: {pop: 0, push: 1},
	vm.CALLDATACOPY: {pop: 3, push: 0},
	vm.CODESIZE: {pop: 0, push: 1},
	vm.CODECOPY: {pop: 3, push: 0},
	vm.GASPRICE: {pop: 0, push: 1},
	vm.EXTCODESIZE: {pop: 1, push: 1},
	vm.EXTCODECOPY: {pop: 4, push: 0},
	vm.RETURNDATASIZE: {pop: 0, push: 1},
	vm.RETURNDATACOPY: {pop: 3, push: 0},
	vm.EXTCODEHASH: {pop: 1, push: 1},
	vm.BLOCKHASH: {pop: 1, push: 1},
	vm.COINBASE: {pop: 0, push: 1},
	vm.TIMESTAMP: {pop: 0, push: 1},
	vm.NUMBER: {pop: 0, push: 1},
	vm.DIFFICULTY: {pop: 0, push: 1},
	vm.GASLIMIT: {pop: 0, push: 1},
	vm.CHAINID: {pop: 0, push: 1},
	vm.SELFBALANCE: {pop: 0, push: 1},
	vm.BASEFEE: {pop: 0, push: 1},
	vm.BLOBHASH: {pop: 1, push: 1},
	vm.BLOBBASEFEE: {pop: 0, push: 1},
	vm.POP: {pop: 1, push: 0},
	vm.MLOAD: {pop: 1, push: 1},
	vm.MSTORE: {pop: 2, push: 0},
	vm.MSTORE8: {pop: 2, push: 0},
	vm.SLOAD: {pop: 1, push: 1},
	vm.SSTORE: {pop: 2, push: 0},
	vm.JUMP: {pop: 1, push: 0},
	vm.JUMPI: {pop: 2, push: 0},
	vm.PC: {pop: 0, push: 1},
	vm.MSIZE: {pop: 0, push: 1},
	vm.GAS: {pop: 0, push: 1},
	vm.JUMPDEST: {pop: 0, push: 0},
	vm.TLOAD: {pop: 1, push: 1},
	vm.TSTORE: {pop: 2, push: 0},
	vm.MCOPY: {pop: 3, push: 0},
	vm.PUSH0: {pop: 0, push: 1},
	vm.PUSH1: {pop: 0, push: 1},
	vm.PUSH2: {pop: 0, push: 1},
	vm.PUSH3: {pop: 0, push: 1},
	vm.PUSH4: {pop: 0, push: 1},
	vm.PUSH5: {pop: 0, push: 1},
	vm.PUSH6: {pop: 0, push: 1},
	vm.PUSH7: {pop: 0, push: 1},
	vm.PUSH8: {pop: 0, push: 1},
	vm.PUSH9: {pop: 0, push: 1},
	vm.PUSH10: {pop: 0, push: 1},
	vm.PUSH11: {pop: 0, push: 1},
	vm.PUSH12: {pop: 0, push: 1},
	vm.PUSH13: {pop: 0, push: 1},
	vm.PUSH14: {pop: 0, push: 1},
	vm.PUSH15: {pop: 0, push: 1},
	vm.PUSH16: {pop: 0, push: 1},
	vm.PUSH17: {pop: 0, push: 1},
	vm.PUSH18: {pop: 0, push: 1},
	vm.PUSH19: {pop: 0, push: 1},
	vm.PUSH20: {pop: 0, push: 1},
	vm.PUSH21: {pop: 0, push: 1},
	vm.PUSH22: {pop: 0, push: 1},
	vm.PUSH23: {pop: 0, push: 1},
	vm.PUSH24: {pop: 0, push: 1},
	vm.PUSH25: {pop: 0, push: 1},
	vm.PUSH26: {pop: 0, push: 1},
	vm.PUSH27: {pop: 0, push: 1},
	vm.PUSH28: {pop: 0, push: 1},
	vm.PUSH29: {pop: 0, push: 1},
	vm.PUSH30: {pop: 0, push: 1},
	vm.PUSH31: {pop: 0, push: 1},
	vm.PUSH32: {pop: 0, push: 1},
	vm.DUP1: {pop: 1, push: 2},
	vm.DUP2: {pop: 1, push: 2},
	vm.DUP3: {pop: 1, push: 2},
	vm.DUP4: {pop: 1, push: 2},
	vm.DUP5: {pop: 1, push: 2},
	vm.DUP6: {pop: 1, push: 2},
	vm.DUP7: {pop: 1, push: 2},
	vm.DUP8: {pop: 1, push: 2},
	vm.DUP9: {pop: 1, push: 2},
	vm.DUP10: {pop: 1, push: 2},
	vm.DUP11: {pop: 1, push: 2},
	vm.DUP12: {pop: 1, push: 2},
	vm.DUP13: {pop: 1, push: 2},
	vm.DUP14: {pop: 1, push: 2},
	vm.DUP15: {pop: 1, push: 2},
	vm.DUP16: {pop: 1, push: 2},
	vm.SWAP1: {pop: 1, push: 1},
	vm.SWAP2: {pop: 1, push: 1},
	vm.SWAP3: {pop: 1, push: 1},
	vm.SWAP4: {pop: 1, push: 1},
	vm.SWAP5: {pop: 1, push: 1},
	vm.SWAP6: {pop: 1, push: 1},
	vm.SWAP7: {pop: 1, push: 1},
	vm.SWAP8: {pop: 1, push: 1},
	vm.SWAP9: {pop: 1, push: 1},
	vm.SWAP10: {pop: 1, push: 1},
	vm.SWAP11: {pop: 1, push: 1},
	vm.SWAP12: {pop: 1, push: 1},
	vm.SWAP13: {pop: 1, push: 1},
	vm.SWAP14: {pop: 1, push: 1},
	vm.SWAP15: {pop: 1, push: 1},
	vm.SWAP16: {pop: 1, push: 1},
	vm.LOG0: {pop: 2, push: 0},
	vm.LOG1: {pop: 3, push: 0},
	vm.LOG2: {pop: 4, push: 0},
	vm.LOG3: {pop: 5, push: 0},
	vm.LOG4: {pop: 6, push: 0},
	vm.CREATE: {pop: 3, push: 1},
	vm.CALL: {pop: 7, push: 1},
	vm.CALLCODE: {pop: 7, push: 1},
	vm.RETURN: {pop: 2, push: 0},
	vm.DELEGATECALL: {pop: 6, push: 1},
	vm.CREATE2: {pop: 4, push: 1},
	vm.STATICCALL: {pop: 6, push: 1},
	vm.REVERT: {pop: 2, push: 0},
	vm.INVALID: {pop: 0, push: 0},
	vm.SELFDESTRUCT: {pop: 1, push: 0},
}
