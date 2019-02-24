// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// LighthouseABI is the input ABI used to generate the binding from.
const LighthouseABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newAuth\",\"type\":\"address\"}],\"name\":\"changeAuth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peekData\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint128\"},{\"name\":\"b\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"read\",\"outputs\":[{\"name\":\"x\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peek\",\"outputs\":[{\"name\":\"v\",\"type\":\"bytes32\"},{\"name\":\"ok\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newMaxAge\",\"type\":\"uint256\"}],\"name\":\"setMaxAge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"DataValue\",\"type\":\"uint256\"},{\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"write\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peekUpdated\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint32\"},{\"name\":\"b\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"peekLastNonce\",\"outputs\":[{\"name\":\"v\",\"type\":\"uint32\"},{\"name\":\"b\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSeeker\",\"type\":\"address\"}],\"name\":\"changeSearcher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"notTooLongSinceUpdated\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LighthouseBin is the compiled bytecode used for deploying new contracts.
const LighthouseBin = `0x60806040526000805433600160a060020a0319918216811790925560018054909116909117905534801561003257600080fd5b50610912806100426000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166322969eac81146100c9578063420b81f6146100ec57806357de26a41461012d57806359e02dd7146101545780635ae28fc9146101825780638da5cb5b1461019a5780639c0e3f7a146101cb578063a6f9dae1146101e6578063bdf384a814610207578063becfbf691461023c578063d6e848ac14610251578063de9375f214610272578063e2f9063214610287575b600080fd5b3480156100d557600080fd5b506100ea600160a060020a03600435166102b0565b005b3480156100f857600080fd5b5061010161032f565b604080516fffffffffffffffffffffffffffffffff909316835290151560208301528051918290030190f35b34801561013957600080fd5b5061014261034f565b60408051918252519081900360200190f35b34801561016057600080fd5b506101696103d5565b6040805192835290151560208301528051918290030190f35b34801561018e57600080fd5b506100ea6004356103f4565b3480156101a657600080fd5b506101af610449565b60408051600160a060020a039092168252519081900360200190f35b3480156101d757600080fd5b506100ea600435602435610458565b3480156101f257600080fd5b506100ea600160a060020a036004351661063e565b34801561021357600080fd5b5061021c6106bd565b6040805163ffffffff909316835290151560208301528051918290030190f35b34801561024857600080fd5b5061021c6106f3565b34801561025d57600080fd5b506100ea600160a060020a036004351661071d565b34801561027e57600080fd5b506101af61087a565b34801561029357600080fd5b5061029c610889565b604080519115158252519081900360200190f35b600054600160a060020a03163314610300576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600061033c610889565b8015610349575060035415155b90509091565b6000610359610889565b8015610366575060035415155b15156103bc576040805160e560020a62461bcd02815260206004820152601360248201527f496e76616c696420646174612073746f72656400000000000000000000000000604482015290519081900360640190fd5b506003546fffffffffffffffffffffffffffffffff1690565b6003546fffffffffffffffffffffffffffffffff16600061033c610889565b600054600160a060020a03163314610444576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b600455565b600054600160a060020a031681565b600154600160a060020a031633146104a8576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b700100000000000000000000000000000000820415610511576040805160e560020a62461bcd02815260206004820152600f60248201527f56616c756520746f6f206c617267650000000000000000000000000000000000604482015290519081900360640190fd5b64010000000081041561056e576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f6e636520746f6f206c617267650000000000000000000000000000000000604482015290519081900360640190fd5b780100000000000000000000000000000000000000000000000081028201700100000000000000000000000000000000420201600355600254600160a060020a03161561063a57600260009054906101000a9004600160a060020a0316600160a060020a031663181783586040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561062157600080fd5b505af1158015610635573d6000803e3d6000fd5b505050505b5050565b600054600160a060020a0316331461068e576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60035470010000000000000000000000000000000090046000816106df610889565b80156106ec575060035415155b9150509091565b600354780100000000000000000000000000000000000000000000000090046000816106df610889565b600054600160a060020a0316331461076d576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517feeb728660000000000000000000000000000000000000000000000000000000081529051929091169163eeb72866916004808201926020929091908290030181600087803b1580156107f157600080fd5b505af1158015610805573d6000803e3d6000fd5b505050506040513d602081101561081b57600080fd5b5051630da4b05514610877576040805160e560020a62461bcd02815260206004820152601060248201527f696e76616c696420736561726368657200000000000000000000000000000000604482015290519081900360640190fd5b50565b600154600160a060020a031681565b600354600454600091700100000000000000000000000000000000900467ffffffffffffffff164203908110806108c05750600454155b915050905600556e617574686f72697365642061636365737300000000000000000000000000a165627a7a723058204ed92e16931f43e968015a1bcf62afea7450a01f58ccf02b72b79a8f536410270029`

// DeployLighthouse deploys a new Ethereum contract, binding an instance of Lighthouse to it.
func DeployLighthouse(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Lighthouse, error) {
	parsed, err := abi.JSON(strings.NewReader(LighthouseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LighthouseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lighthouse{LighthouseCaller: LighthouseCaller{contract: contract}, LighthouseTransactor: LighthouseTransactor{contract: contract}, LighthouseFilterer: LighthouseFilterer{contract: contract}}, nil
}

// Lighthouse is an auto generated Go binding around an Ethereum contract.
type Lighthouse struct {
	LighthouseCaller     // Read-only binding to the contract
	LighthouseTransactor // Write-only binding to the contract
	LighthouseFilterer   // Log filterer for contract events
}

// LighthouseCaller is an auto generated read-only Go binding around an Ethereum contract.
type LighthouseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LighthouseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LighthouseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LighthouseSession struct {
	Contract     *Lighthouse       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LighthouseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LighthouseCallerSession struct {
	Contract *LighthouseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LighthouseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LighthouseTransactorSession struct {
	Contract     *LighthouseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LighthouseRaw is an auto generated low-level Go binding around an Ethereum contract.
type LighthouseRaw struct {
	Contract *Lighthouse // Generic contract binding to access the raw methods on
}

// LighthouseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LighthouseCallerRaw struct {
	Contract *LighthouseCaller // Generic read-only contract binding to access the raw methods on
}

// LighthouseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LighthouseTransactorRaw struct {
	Contract *LighthouseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLighthouse creates a new instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouse(address common.Address, backend bind.ContractBackend) (*Lighthouse, error) {
	contract, err := bindLighthouse(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lighthouse{LighthouseCaller: LighthouseCaller{contract: contract}, LighthouseTransactor: LighthouseTransactor{contract: contract}, LighthouseFilterer: LighthouseFilterer{contract: contract}}, nil
}

// NewLighthouseCaller creates a new read-only instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouseCaller(address common.Address, caller bind.ContractCaller) (*LighthouseCaller, error) {
	contract, err := bindLighthouse(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LighthouseCaller{contract: contract}, nil
}

// NewLighthouseTransactor creates a new write-only instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouseTransactor(address common.Address, transactor bind.ContractTransactor) (*LighthouseTransactor, error) {
	contract, err := bindLighthouse(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LighthouseTransactor{contract: contract}, nil
}

// NewLighthouseFilterer creates a new log filterer instance of Lighthouse, bound to a specific deployed contract.
func NewLighthouseFilterer(address common.Address, filterer bind.ContractFilterer) (*LighthouseFilterer, error) {
	contract, err := bindLighthouse(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LighthouseFilterer{contract: contract}, nil
}

// bindLighthouse binds a generic wrapper to an already deployed contract.
func bindLighthouse(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LighthouseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lighthouse *LighthouseRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lighthouse.Contract.LighthouseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lighthouse *LighthouseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lighthouse.Contract.LighthouseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lighthouse *LighthouseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lighthouse.Contract.LighthouseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lighthouse *LighthouseCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Lighthouse.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lighthouse *LighthouseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lighthouse.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lighthouse *LighthouseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lighthouse.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() constant returns(address)
func (_Lighthouse *LighthouseCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lighthouse.contract.Call(opts, out, "auth")
	return *ret0, err
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() constant returns(address)
func (_Lighthouse *LighthouseSession) Auth() (common.Address, error) {
	return _Lighthouse.Contract.Auth(&_Lighthouse.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() constant returns(address)
func (_Lighthouse *LighthouseCallerSession) Auth() (common.Address, error) {
	return _Lighthouse.Contract.Auth(&_Lighthouse.CallOpts)
}

// NotTooLongSinceUpdated is a free data retrieval call binding the contract method 0xe2f90632.
//
// Solidity: function notTooLongSinceUpdated() constant returns(bool)
func (_Lighthouse *LighthouseCaller) NotTooLongSinceUpdated(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Lighthouse.contract.Call(opts, out, "notTooLongSinceUpdated")
	return *ret0, err
}

// NotTooLongSinceUpdated is a free data retrieval call binding the contract method 0xe2f90632.
//
// Solidity: function notTooLongSinceUpdated() constant returns(bool)
func (_Lighthouse *LighthouseSession) NotTooLongSinceUpdated() (bool, error) {
	return _Lighthouse.Contract.NotTooLongSinceUpdated(&_Lighthouse.CallOpts)
}

// NotTooLongSinceUpdated is a free data retrieval call binding the contract method 0xe2f90632.
//
// Solidity: function notTooLongSinceUpdated() constant returns(bool)
func (_Lighthouse *LighthouseCallerSession) NotTooLongSinceUpdated() (bool, error) {
	return _Lighthouse.Contract.NotTooLongSinceUpdated(&_Lighthouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lighthouse *LighthouseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Lighthouse.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lighthouse *LighthouseSession) Owner() (common.Address, error) {
	return _Lighthouse.Contract.Owner(&_Lighthouse.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Lighthouse *LighthouseCallerSession) Owner() (common.Address, error) {
	return _Lighthouse.Contract.Owner(&_Lighthouse.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(v bytes32, ok bool)
func (_Lighthouse *LighthouseCaller) Peek(opts *bind.CallOpts) (struct {
	V  [32]byte
	Ok bool
}, error) {
	ret := new(struct {
		V  [32]byte
		Ok bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peek")
	return *ret, err
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(v bytes32, ok bool)
func (_Lighthouse *LighthouseSession) Peek() (struct {
	V  [32]byte
	Ok bool
}, error) {
	return _Lighthouse.Contract.Peek(&_Lighthouse.CallOpts)
}

// Peek is a free data retrieval call binding the contract method 0x59e02dd7.
//
// Solidity: function peek() constant returns(v bytes32, ok bool)
func (_Lighthouse *LighthouseCallerSession) Peek() (struct {
	V  [32]byte
	Ok bool
}, error) {
	return _Lighthouse.Contract.Peek(&_Lighthouse.CallOpts)
}

// PeekData is a free data retrieval call binding the contract method 0x420b81f6.
//
// Solidity: function peekData() constant returns(v uint128, b bool)
func (_Lighthouse *LighthouseCaller) PeekData(opts *bind.CallOpts) (struct {
	V *big.Int
	B bool
}, error) {
	ret := new(struct {
		V *big.Int
		B bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peekData")
	return *ret, err
}

// PeekData is a free data retrieval call binding the contract method 0x420b81f6.
//
// Solidity: function peekData() constant returns(v uint128, b bool)
func (_Lighthouse *LighthouseSession) PeekData() (struct {
	V *big.Int
	B bool
}, error) {
	return _Lighthouse.Contract.PeekData(&_Lighthouse.CallOpts)
}

// PeekData is a free data retrieval call binding the contract method 0x420b81f6.
//
// Solidity: function peekData() constant returns(v uint128, b bool)
func (_Lighthouse *LighthouseCallerSession) PeekData() (struct {
	V *big.Int
	B bool
}, error) {
	return _Lighthouse.Contract.PeekData(&_Lighthouse.CallOpts)
}

// PeekLastNonce is a free data retrieval call binding the contract method 0xbecfbf69.
//
// Solidity: function peekLastNonce() constant returns(v uint32, b bool)
func (_Lighthouse *LighthouseCaller) PeekLastNonce(opts *bind.CallOpts) (struct {
	V uint32
	B bool
}, error) {
	ret := new(struct {
		V uint32
		B bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peekLastNonce")
	return *ret, err
}

// PeekLastNonce is a free data retrieval call binding the contract method 0xbecfbf69.
//
// Solidity: function peekLastNonce() constant returns(v uint32, b bool)
func (_Lighthouse *LighthouseSession) PeekLastNonce() (struct {
	V uint32
	B bool
}, error) {
	return _Lighthouse.Contract.PeekLastNonce(&_Lighthouse.CallOpts)
}

// PeekLastNonce is a free data retrieval call binding the contract method 0xbecfbf69.
//
// Solidity: function peekLastNonce() constant returns(v uint32, b bool)
func (_Lighthouse *LighthouseCallerSession) PeekLastNonce() (struct {
	V uint32
	B bool
}, error) {
	return _Lighthouse.Contract.PeekLastNonce(&_Lighthouse.CallOpts)
}

// PeekUpdated is a free data retrieval call binding the contract method 0xbdf384a8.
//
// Solidity: function peekUpdated() constant returns(v uint32, b bool)
func (_Lighthouse *LighthouseCaller) PeekUpdated(opts *bind.CallOpts) (struct {
	V uint32
	B bool
}, error) {
	ret := new(struct {
		V uint32
		B bool
	})
	out := ret
	err := _Lighthouse.contract.Call(opts, out, "peekUpdated")
	return *ret, err
}

// PeekUpdated is a free data retrieval call binding the contract method 0xbdf384a8.
//
// Solidity: function peekUpdated() constant returns(v uint32, b bool)
func (_Lighthouse *LighthouseSession) PeekUpdated() (struct {
	V uint32
	B bool
}, error) {
	return _Lighthouse.Contract.PeekUpdated(&_Lighthouse.CallOpts)
}

// PeekUpdated is a free data retrieval call binding the contract method 0xbdf384a8.
//
// Solidity: function peekUpdated() constant returns(v uint32, b bool)
func (_Lighthouse *LighthouseCallerSession) PeekUpdated() (struct {
	V uint32
	B bool
}, error) {
	return _Lighthouse.Contract.PeekUpdated(&_Lighthouse.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(x bytes32)
func (_Lighthouse *LighthouseCaller) Read(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Lighthouse.contract.Call(opts, out, "read")
	return *ret0, err
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(x bytes32)
func (_Lighthouse *LighthouseSession) Read() ([32]byte, error) {
	return _Lighthouse.Contract.Read(&_Lighthouse.CallOpts)
}

// Read is a free data retrieval call binding the contract method 0x57de26a4.
//
// Solidity: function read() constant returns(x bytes32)
func (_Lighthouse *LighthouseCallerSession) Read() ([32]byte, error) {
	return _Lighthouse.Contract.Read(&_Lighthouse.CallOpts)
}

// ChangeAuth is a paid mutator transaction binding the contract method 0x22969eac.
//
// Solidity: function changeAuth(newAuth address) returns()
func (_Lighthouse *LighthouseTransactor) ChangeAuth(opts *bind.TransactOpts, newAuth common.Address) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "changeAuth", newAuth)
}

// ChangeAuth is a paid mutator transaction binding the contract method 0x22969eac.
//
// Solidity: function changeAuth(newAuth address) returns()
func (_Lighthouse *LighthouseSession) ChangeAuth(newAuth common.Address) (*types.Transaction, error) {
	return _Lighthouse.Contract.ChangeAuth(&_Lighthouse.TransactOpts, newAuth)
}

// ChangeAuth is a paid mutator transaction binding the contract method 0x22969eac.
//
// Solidity: function changeAuth(newAuth address) returns()
func (_Lighthouse *LighthouseTransactorSession) ChangeAuth(newAuth common.Address) (*types.Transaction, error) {
	return _Lighthouse.Contract.ChangeAuth(&_Lighthouse.TransactOpts, newAuth)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Lighthouse *LighthouseTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Lighthouse *LighthouseSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Lighthouse.Contract.ChangeOwner(&_Lighthouse.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(newOwner address) returns()
func (_Lighthouse *LighthouseTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Lighthouse.Contract.ChangeOwner(&_Lighthouse.TransactOpts, newOwner)
}

// ChangeSearcher is a paid mutator transaction binding the contract method 0xd6e848ac.
//
// Solidity: function changeSearcher(newSeeker address) returns()
func (_Lighthouse *LighthouseTransactor) ChangeSearcher(opts *bind.TransactOpts, newSeeker common.Address) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "changeSearcher", newSeeker)
}

// ChangeSearcher is a paid mutator transaction binding the contract method 0xd6e848ac.
//
// Solidity: function changeSearcher(newSeeker address) returns()
func (_Lighthouse *LighthouseSession) ChangeSearcher(newSeeker common.Address) (*types.Transaction, error) {
	return _Lighthouse.Contract.ChangeSearcher(&_Lighthouse.TransactOpts, newSeeker)
}

// ChangeSearcher is a paid mutator transaction binding the contract method 0xd6e848ac.
//
// Solidity: function changeSearcher(newSeeker address) returns()
func (_Lighthouse *LighthouseTransactorSession) ChangeSearcher(newSeeker common.Address) (*types.Transaction, error) {
	return _Lighthouse.Contract.ChangeSearcher(&_Lighthouse.TransactOpts, newSeeker)
}

// SetMaxAge is a paid mutator transaction binding the contract method 0x5ae28fc9.
//
// Solidity: function setMaxAge(newMaxAge uint256) returns()
func (_Lighthouse *LighthouseTransactor) SetMaxAge(opts *bind.TransactOpts, newMaxAge *big.Int) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "setMaxAge", newMaxAge)
}

// SetMaxAge is a paid mutator transaction binding the contract method 0x5ae28fc9.
//
// Solidity: function setMaxAge(newMaxAge uint256) returns()
func (_Lighthouse *LighthouseSession) SetMaxAge(newMaxAge *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.SetMaxAge(&_Lighthouse.TransactOpts, newMaxAge)
}

// SetMaxAge is a paid mutator transaction binding the contract method 0x5ae28fc9.
//
// Solidity: function setMaxAge(newMaxAge uint256) returns()
func (_Lighthouse *LighthouseTransactorSession) SetMaxAge(newMaxAge *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.SetMaxAge(&_Lighthouse.TransactOpts, newMaxAge)
}

// Write is a paid mutator transaction binding the contract method 0x9c0e3f7a.
//
// Solidity: function write(DataValue uint256, nonce uint256) returns()
func (_Lighthouse *LighthouseTransactor) Write(opts *bind.TransactOpts, DataValue *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Lighthouse.contract.Transact(opts, "write", DataValue, nonce)
}

// Write is a paid mutator transaction binding the contract method 0x9c0e3f7a.
//
// Solidity: function write(DataValue uint256, nonce uint256) returns()
func (_Lighthouse *LighthouseSession) Write(DataValue *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.Write(&_Lighthouse.TransactOpts, DataValue, nonce)
}

// Write is a paid mutator transaction binding the contract method 0x9c0e3f7a.
//
// Solidity: function write(DataValue uint256, nonce uint256) returns()
func (_Lighthouse *LighthouseTransactorSession) Write(DataValue *big.Int, nonce *big.Int) (*types.Transaction, error) {
	return _Lighthouse.Contract.Write(&_Lighthouse.TransactOpts, DataValue, nonce)
}

// LighthouseFactoryABI is the input ABI used to generate the binding from.
const LighthouseFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newAuth\",\"type\":\"address\"}],\"name\":\"changeAuth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberOfLighthouses\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lighthouses\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"makeLighthouse\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"auth\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newLighthouse\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"housePosition\",\"type\":\"uint256\"}],\"name\":\"newLighthouse\",\"type\":\"event\"}]"

// LighthouseFactoryBin is the compiled bytecode used for deploying new contracts.
const LighthouseFactoryBin = `0x608060405260008054600160a060020a0319163317905534801561002257600080fd5b50610d4a806100326000396000f30060806040526004361061006c5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166322969eac8114610071578063488ebb291461009457806399a06549146100bb578063dd7e88e7146100ef578063de9375f214610104575b600080fd5b34801561007d57600080fd5b50610092600160a060020a0360043516610119565b005b3480156100a057600080fd5b506100a96101c1565b60408051918252519081900360200190f35b3480156100c757600080fd5b506100d36004356101c7565b60408051600160a060020a039092168252519081900360200190f35b3480156100fb57600080fd5b506100926101ef565b34801561011057600080fd5b506100d36103ab565b600054600160a060020a0316331461019257604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f556e617574686f72697365642061636365737300000000000000000000000000604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60015490565b60018054829081106101d557fe5b600091825260209091200154600160a060020a0316905081565b600080548190600160a060020a0316331461026b57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601360248201527f556e617574686f72697365642061636365737300000000000000000000000000604482015290519081900360640190fd5b6102736103ba565b604051809103906000f08015801561028f573d6000803e3d6000fd5b5060008054604080517f22969eac000000000000000000000000000000000000000000000000000000008152600160a060020a0392831660048201529051939550908516926322969eac9260248084019391929182900301818387803b1580156102f857600080fd5b505af115801561030c573d6000803e3d6000fd5b5050600180548082018083556000929092527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf601805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038716908117909155604080519182526020820183905280519295507f035beaf06f9aaeac59c9d5cb26b559a872fc1c2f3c67c5792d13b35bbfa85856945090829003019150a15050565b600054600160a060020a031681565b604051610954806103cb83390190560060806040526000805433600160a060020a0319918216811790925560018054909116909117905534801561003257600080fd5b50610912806100426000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166322969eac81146100c9578063420b81f6146100ec57806357de26a41461012d57806359e02dd7146101545780635ae28fc9146101825780638da5cb5b1461019a5780639c0e3f7a146101cb578063a6f9dae1146101e6578063bdf384a814610207578063becfbf691461023c578063d6e848ac14610251578063de9375f214610272578063e2f9063214610287575b600080fd5b3480156100d557600080fd5b506100ea600160a060020a03600435166102b0565b005b3480156100f857600080fd5b5061010161032f565b604080516fffffffffffffffffffffffffffffffff909316835290151560208301528051918290030190f35b34801561013957600080fd5b5061014261034f565b60408051918252519081900360200190f35b34801561016057600080fd5b506101696103d5565b6040805192835290151560208301528051918290030190f35b34801561018e57600080fd5b506100ea6004356103f4565b3480156101a657600080fd5b506101af610449565b60408051600160a060020a039092168252519081900360200190f35b3480156101d757600080fd5b506100ea600435602435610458565b3480156101f257600080fd5b506100ea600160a060020a036004351661063e565b34801561021357600080fd5b5061021c6106bd565b6040805163ffffffff909316835290151560208301528051918290030190f35b34801561024857600080fd5b5061021c6106f3565b34801561025d57600080fd5b506100ea600160a060020a036004351661071d565b34801561027e57600080fd5b506101af61087a565b34801561029357600080fd5b5061029c610889565b604080519115158252519081900360200190f35b600054600160a060020a03163314610300576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600354600061033c610889565b8015610349575060035415155b90509091565b6000610359610889565b8015610366575060035415155b15156103bc576040805160e560020a62461bcd02815260206004820152601360248201527f496e76616c696420646174612073746f72656400000000000000000000000000604482015290519081900360640190fd5b506003546fffffffffffffffffffffffffffffffff1690565b6003546fffffffffffffffffffffffffffffffff16600061033c610889565b600054600160a060020a03163314610444576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b600455565b600054600160a060020a031681565b600154600160a060020a031633146104a8576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b700100000000000000000000000000000000820415610511576040805160e560020a62461bcd02815260206004820152600f60248201527f56616c756520746f6f206c617267650000000000000000000000000000000000604482015290519081900360640190fd5b64010000000081041561056e576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f6e636520746f6f206c617267650000000000000000000000000000000000604482015290519081900360640190fd5b780100000000000000000000000000000000000000000000000081028201700100000000000000000000000000000000420201600355600254600160a060020a03161561063a57600260009054906101000a9004600160a060020a0316600160a060020a031663181783586040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401600060405180830381600087803b15801561062157600080fd5b505af1158015610635573d6000803e3d6000fd5b505050505b5050565b600054600160a060020a0316331461068e576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60035470010000000000000000000000000000000090046000816106df610889565b80156106ec575060035415155b9150509091565b600354780100000000000000000000000000000000000000000000000090046000816106df610889565b600054600160a060020a0316331461076d576040805160e560020a62461bcd02815260206004820152601360248201526000805160206108c7833981519152604482015290519081900360640190fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055604080517feeb728660000000000000000000000000000000000000000000000000000000081529051929091169163eeb72866916004808201926020929091908290030181600087803b1580156107f157600080fd5b505af1158015610805573d6000803e3d6000fd5b505050506040513d602081101561081b57600080fd5b5051630da4b05514610877576040805160e560020a62461bcd02815260206004820152601060248201527f696e76616c696420736561726368657200000000000000000000000000000000604482015290519081900360640190fd5b50565b600154600160a060020a031681565b600354600454600091700100000000000000000000000000000000900467ffffffffffffffff164203908110806108c05750600454155b915050905600556e617574686f72697365642061636365737300000000000000000000000000a165627a7a723058204ed92e16931f43e968015a1bcf62afea7450a01f58ccf02b72b79a8f536410270029a165627a7a72305820ada0fa605a89559d4c619517ad379a850e6a92a45f66e75a2a7afb5a35bb194e0029`

// DeployLighthouseFactory deploys a new Ethereum contract, binding an instance of LighthouseFactory to it.
func DeployLighthouseFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LighthouseFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(LighthouseFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LighthouseFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LighthouseFactory{LighthouseFactoryCaller: LighthouseFactoryCaller{contract: contract}, LighthouseFactoryTransactor: LighthouseFactoryTransactor{contract: contract}, LighthouseFactoryFilterer: LighthouseFactoryFilterer{contract: contract}}, nil
}

// LighthouseFactory is an auto generated Go binding around an Ethereum contract.
type LighthouseFactory struct {
	LighthouseFactoryCaller     // Read-only binding to the contract
	LighthouseFactoryTransactor // Write-only binding to the contract
	LighthouseFactoryFilterer   // Log filterer for contract events
}

// LighthouseFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LighthouseFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LighthouseFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LighthouseFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LighthouseFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LighthouseFactorySession struct {
	Contract     *LighthouseFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LighthouseFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LighthouseFactoryCallerSession struct {
	Contract *LighthouseFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// LighthouseFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LighthouseFactoryTransactorSession struct {
	Contract     *LighthouseFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// LighthouseFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LighthouseFactoryRaw struct {
	Contract *LighthouseFactory // Generic contract binding to access the raw methods on
}

// LighthouseFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LighthouseFactoryCallerRaw struct {
	Contract *LighthouseFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// LighthouseFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LighthouseFactoryTransactorRaw struct {
	Contract *LighthouseFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLighthouseFactory creates a new instance of LighthouseFactory, bound to a specific deployed contract.
func NewLighthouseFactory(address common.Address, backend bind.ContractBackend) (*LighthouseFactory, error) {
	contract, err := bindLighthouseFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LighthouseFactory{LighthouseFactoryCaller: LighthouseFactoryCaller{contract: contract}, LighthouseFactoryTransactor: LighthouseFactoryTransactor{contract: contract}, LighthouseFactoryFilterer: LighthouseFactoryFilterer{contract: contract}}, nil
}

// NewLighthouseFactoryCaller creates a new read-only instance of LighthouseFactory, bound to a specific deployed contract.
func NewLighthouseFactoryCaller(address common.Address, caller bind.ContractCaller) (*LighthouseFactoryCaller, error) {
	contract, err := bindLighthouseFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LighthouseFactoryCaller{contract: contract}, nil
}

// NewLighthouseFactoryTransactor creates a new write-only instance of LighthouseFactory, bound to a specific deployed contract.
func NewLighthouseFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*LighthouseFactoryTransactor, error) {
	contract, err := bindLighthouseFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LighthouseFactoryTransactor{contract: contract}, nil
}

// NewLighthouseFactoryFilterer creates a new log filterer instance of LighthouseFactory, bound to a specific deployed contract.
func NewLighthouseFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*LighthouseFactoryFilterer, error) {
	contract, err := bindLighthouseFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LighthouseFactoryFilterer{contract: contract}, nil
}

// bindLighthouseFactory binds a generic wrapper to an already deployed contract.
func bindLighthouseFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LighthouseFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LighthouseFactory *LighthouseFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LighthouseFactory.Contract.LighthouseFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LighthouseFactory *LighthouseFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LighthouseFactory.Contract.LighthouseFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LighthouseFactory *LighthouseFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LighthouseFactory.Contract.LighthouseFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LighthouseFactory *LighthouseFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LighthouseFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LighthouseFactory *LighthouseFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LighthouseFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LighthouseFactory *LighthouseFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LighthouseFactory.Contract.contract.Transact(opts, method, params...)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() constant returns(address)
func (_LighthouseFactory *LighthouseFactoryCaller) Auth(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LighthouseFactory.contract.Call(opts, out, "auth")
	return *ret0, err
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() constant returns(address)
func (_LighthouseFactory *LighthouseFactorySession) Auth() (common.Address, error) {
	return _LighthouseFactory.Contract.Auth(&_LighthouseFactory.CallOpts)
}

// Auth is a free data retrieval call binding the contract method 0xde9375f2.
//
// Solidity: function auth() constant returns(address)
func (_LighthouseFactory *LighthouseFactoryCallerSession) Auth() (common.Address, error) {
	return _LighthouseFactory.Contract.Auth(&_LighthouseFactory.CallOpts)
}

// Lighthouses is a free data retrieval call binding the contract method 0x99a06549.
//
// Solidity: function lighthouses( uint256) constant returns(address)
func (_LighthouseFactory *LighthouseFactoryCaller) Lighthouses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LighthouseFactory.contract.Call(opts, out, "lighthouses", arg0)
	return *ret0, err
}

// Lighthouses is a free data retrieval call binding the contract method 0x99a06549.
//
// Solidity: function lighthouses( uint256) constant returns(address)
func (_LighthouseFactory *LighthouseFactorySession) Lighthouses(arg0 *big.Int) (common.Address, error) {
	return _LighthouseFactory.Contract.Lighthouses(&_LighthouseFactory.CallOpts, arg0)
}

// Lighthouses is a free data retrieval call binding the contract method 0x99a06549.
//
// Solidity: function lighthouses( uint256) constant returns(address)
func (_LighthouseFactory *LighthouseFactoryCallerSession) Lighthouses(arg0 *big.Int) (common.Address, error) {
	return _LighthouseFactory.Contract.Lighthouses(&_LighthouseFactory.CallOpts, arg0)
}

// NumberOfLighthouses is a free data retrieval call binding the contract method 0x488ebb29.
//
// Solidity: function numberOfLighthouses() constant returns(uint256)
func (_LighthouseFactory *LighthouseFactoryCaller) NumberOfLighthouses(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LighthouseFactory.contract.Call(opts, out, "numberOfLighthouses")
	return *ret0, err
}

// NumberOfLighthouses is a free data retrieval call binding the contract method 0x488ebb29.
//
// Solidity: function numberOfLighthouses() constant returns(uint256)
func (_LighthouseFactory *LighthouseFactorySession) NumberOfLighthouses() (*big.Int, error) {
	return _LighthouseFactory.Contract.NumberOfLighthouses(&_LighthouseFactory.CallOpts)
}

// NumberOfLighthouses is a free data retrieval call binding the contract method 0x488ebb29.
//
// Solidity: function numberOfLighthouses() constant returns(uint256)
func (_LighthouseFactory *LighthouseFactoryCallerSession) NumberOfLighthouses() (*big.Int, error) {
	return _LighthouseFactory.Contract.NumberOfLighthouses(&_LighthouseFactory.CallOpts)
}

// ChangeAuth is a paid mutator transaction binding the contract method 0x22969eac.
//
// Solidity: function changeAuth(newAuth address) returns()
func (_LighthouseFactory *LighthouseFactoryTransactor) ChangeAuth(opts *bind.TransactOpts, newAuth common.Address) (*types.Transaction, error) {
	return _LighthouseFactory.contract.Transact(opts, "changeAuth", newAuth)
}

// ChangeAuth is a paid mutator transaction binding the contract method 0x22969eac.
//
// Solidity: function changeAuth(newAuth address) returns()
func (_LighthouseFactory *LighthouseFactorySession) ChangeAuth(newAuth common.Address) (*types.Transaction, error) {
	return _LighthouseFactory.Contract.ChangeAuth(&_LighthouseFactory.TransactOpts, newAuth)
}

// ChangeAuth is a paid mutator transaction binding the contract method 0x22969eac.
//
// Solidity: function changeAuth(newAuth address) returns()
func (_LighthouseFactory *LighthouseFactoryTransactorSession) ChangeAuth(newAuth common.Address) (*types.Transaction, error) {
	return _LighthouseFactory.Contract.ChangeAuth(&_LighthouseFactory.TransactOpts, newAuth)
}

// MakeLighthouse is a paid mutator transaction binding the contract method 0xdd7e88e7.
//
// Solidity: function makeLighthouse() returns()
func (_LighthouseFactory *LighthouseFactoryTransactor) MakeLighthouse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LighthouseFactory.contract.Transact(opts, "makeLighthouse")
}

// MakeLighthouse is a paid mutator transaction binding the contract method 0xdd7e88e7.
//
// Solidity: function makeLighthouse() returns()
func (_LighthouseFactory *LighthouseFactorySession) MakeLighthouse() (*types.Transaction, error) {
	return _LighthouseFactory.Contract.MakeLighthouse(&_LighthouseFactory.TransactOpts)
}

// MakeLighthouse is a paid mutator transaction binding the contract method 0xdd7e88e7.
//
// Solidity: function makeLighthouse() returns()
func (_LighthouseFactory *LighthouseFactoryTransactorSession) MakeLighthouse() (*types.Transaction, error) {
	return _LighthouseFactory.Contract.MakeLighthouse(&_LighthouseFactory.TransactOpts)
}

// LighthouseFactoryNewLighthouseIterator is returned from FilterNewLighthouse and is used to iterate over the raw logs and unpacked data for NewLighthouse events raised by the LighthouseFactory contract.
type LighthouseFactoryNewLighthouseIterator struct {
	Event *LighthouseFactoryNewLighthouse // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LighthouseFactoryNewLighthouseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LighthouseFactoryNewLighthouse)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LighthouseFactoryNewLighthouse)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LighthouseFactoryNewLighthouseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LighthouseFactoryNewLighthouseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LighthouseFactoryNewLighthouse represents a NewLighthouse event raised by the LighthouseFactory contract.
type LighthouseFactoryNewLighthouse struct {
	NewLighthouse common.Address
	HousePosition *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewLighthouse is a free log retrieval operation binding the contract event 0x035beaf06f9aaeac59c9d5cb26b559a872fc1c2f3c67c5792d13b35bbfa85856.
//
// Solidity: event newLighthouse(newLighthouse address, housePosition uint256)
func (_LighthouseFactory *LighthouseFactoryFilterer) FilterNewLighthouse(opts *bind.FilterOpts) (*LighthouseFactoryNewLighthouseIterator, error) {

	logs, sub, err := _LighthouseFactory.contract.FilterLogs(opts, "newLighthouse")
	if err != nil {
		return nil, err
	}
	return &LighthouseFactoryNewLighthouseIterator{contract: _LighthouseFactory.contract, event: "newLighthouse", logs: logs, sub: sub}, nil
}

// WatchNewLighthouse is a free log subscription operation binding the contract event 0x035beaf06f9aaeac59c9d5cb26b559a872fc1c2f3c67c5792d13b35bbfa85856.
//
// Solidity: event newLighthouse(newLighthouse address, housePosition uint256)
func (_LighthouseFactory *LighthouseFactoryFilterer) WatchNewLighthouse(opts *bind.WatchOpts, sink chan<- *LighthouseFactoryNewLighthouse) (event.Subscription, error) {

	logs, sub, err := _LighthouseFactory.contract.WatchLogs(opts, "newLighthouse")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LighthouseFactoryNewLighthouse)
				if err := _LighthouseFactory.contract.UnpackLog(event, "newLighthouse", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// SearcherABI is the input ABI used to generate the binding from.
const SearcherABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"poke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"identify\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SearcherBin is the compiled bytecode used for deploying new contracts.
const SearcherBin = `0x`

// DeploySearcher deploys a new Ethereum contract, binding an instance of Searcher to it.
func DeploySearcher(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Searcher, error) {
	parsed, err := abi.JSON(strings.NewReader(SearcherABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SearcherBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Searcher{SearcherCaller: SearcherCaller{contract: contract}, SearcherTransactor: SearcherTransactor{contract: contract}, SearcherFilterer: SearcherFilterer{contract: contract}}, nil
}

// Searcher is an auto generated Go binding around an Ethereum contract.
type Searcher struct {
	SearcherCaller     // Read-only binding to the contract
	SearcherTransactor // Write-only binding to the contract
	SearcherFilterer   // Log filterer for contract events
}

// SearcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type SearcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SearcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SearcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SearcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SearcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SearcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SearcherSession struct {
	Contract     *Searcher         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SearcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SearcherCallerSession struct {
	Contract *SearcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SearcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SearcherTransactorSession struct {
	Contract     *SearcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SearcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type SearcherRaw struct {
	Contract *Searcher // Generic contract binding to access the raw methods on
}

// SearcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SearcherCallerRaw struct {
	Contract *SearcherCaller // Generic read-only contract binding to access the raw methods on
}

// SearcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SearcherTransactorRaw struct {
	Contract *SearcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSearcher creates a new instance of Searcher, bound to a specific deployed contract.
func NewSearcher(address common.Address, backend bind.ContractBackend) (*Searcher, error) {
	contract, err := bindSearcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Searcher{SearcherCaller: SearcherCaller{contract: contract}, SearcherTransactor: SearcherTransactor{contract: contract}, SearcherFilterer: SearcherFilterer{contract: contract}}, nil
}

// NewSearcherCaller creates a new read-only instance of Searcher, bound to a specific deployed contract.
func NewSearcherCaller(address common.Address, caller bind.ContractCaller) (*SearcherCaller, error) {
	contract, err := bindSearcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SearcherCaller{contract: contract}, nil
}

// NewSearcherTransactor creates a new write-only instance of Searcher, bound to a specific deployed contract.
func NewSearcherTransactor(address common.Address, transactor bind.ContractTransactor) (*SearcherTransactor, error) {
	contract, err := bindSearcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SearcherTransactor{contract: contract}, nil
}

// NewSearcherFilterer creates a new log filterer instance of Searcher, bound to a specific deployed contract.
func NewSearcherFilterer(address common.Address, filterer bind.ContractFilterer) (*SearcherFilterer, error) {
	contract, err := bindSearcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SearcherFilterer{contract: contract}, nil
}

// bindSearcher binds a generic wrapper to an already deployed contract.
func bindSearcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SearcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Searcher *SearcherRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Searcher.Contract.SearcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Searcher *SearcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Searcher.Contract.SearcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Searcher *SearcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Searcher.Contract.SearcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Searcher *SearcherCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Searcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Searcher *SearcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Searcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Searcher *SearcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Searcher.Contract.contract.Transact(opts, method, params...)
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() constant returns(uint256)
func (_Searcher *SearcherCaller) Identify(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Searcher.contract.Call(opts, out, "identify")
	return *ret0, err
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() constant returns(uint256)
func (_Searcher *SearcherSession) Identify() (*big.Int, error) {
	return _Searcher.Contract.Identify(&_Searcher.CallOpts)
}

// Identify is a free data retrieval call binding the contract method 0xeeb72866.
//
// Solidity: function identify() constant returns(uint256)
func (_Searcher *SearcherCallerSession) Identify() (*big.Int, error) {
	return _Searcher.Contract.Identify(&_Searcher.CallOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Searcher *SearcherTransactor) Poke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Searcher.contract.Transact(opts, "poke")
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Searcher *SearcherSession) Poke() (*types.Transaction, error) {
	return _Searcher.Contract.Poke(&_Searcher.TransactOpts)
}

// Poke is a paid mutator transaction binding the contract method 0x18178358.
//
// Solidity: function poke() returns()
func (_Searcher *SearcherTransactorSession) Poke() (*types.Transaction, error) {
	return _Searcher.Contract.Poke(&_Searcher.TransactOpts)
}
