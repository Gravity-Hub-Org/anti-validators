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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x608060405261083b806100136000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a0823114610149578063a457c2d71461016f578063a9059cbb1461019b578063dd62ed3e146101c757610088565b8063095ea7b31461008d57806318160ddd146100cd57806323b872dd146100e7578063395093511461011d575b600080fd5b6100b9600480360360408110156100a357600080fd5b506001600160a01b0381351690602001356101f5565b604080519115158252519081900360200190f35b6100d5610212565b60408051918252519081900360200190f35b6100b9600480360360608110156100fd57600080fd5b506001600160a01b03813581169160208101359091169060400135610218565b6100b96004803603604081101561013357600080fd5b506001600160a01b0381351690602001356102a5565b6100d56004803603602081101561015f57600080fd5b50356001600160a01b03166102f9565b6100b96004803603604081101561018557600080fd5b506001600160a01b038135169060200135610314565b6100b9600480360360408110156101b157600080fd5b506001600160a01b038135169060200135610382565b6100d5600480360360408110156101dd57600080fd5b506001600160a01b0381358116916020013516610396565b60006102096102026103c1565b84846103c5565b50600192915050565b60025490565b60006102258484846104b1565b61029b846102316103c1565b61029685604051806060016040528060288152602001610771602891396001600160a01b038a1660009081526001602052604081209061026f6103c1565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61060d16565b6103c5565b5060019392505050565b60006102096102b26103c1565b8461029685600160006102c36103c1565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff6106a416565b6001600160a01b031660009081526020819052604090205490565b60006102096103216103c1565b84610296856040518060600160405280602581526020016107e2602591396001600061034b6103c1565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61060d16565b600061020961038f6103c1565b84846104b1565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b03831661040a5760405162461bcd60e51b81526004018080602001828103825260248152602001806107be6024913960400191505060405180910390fd5b6001600160a01b03821661044f5760405162461bcd60e51b81526004018080602001828103825260228152602001806107296022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b0383166104f65760405162461bcd60e51b81526004018080602001828103825260258152602001806107996025913960400191505060405180910390fd5b6001600160a01b03821661053b5760405162461bcd60e51b81526004018080602001828103825260238152602001806107066023913960400191505060405180910390fd5b61057e8160405180606001604052806026815260200161074b602691396001600160a01b038616600090815260208190526040902054919063ffffffff61060d16565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546105b3908263ffffffff6106a416565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b6000818484111561069c5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610661578181015183820152602001610649565b50505050905090810190601f16801561068e5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b6000828201838110156106fe576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820d6c6cf767c97aed997eb8d371f7d80730ae3b517edfd1fee423f2eccf8c878ff64736f6c63430005100032"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleABI is the input ABI used to generate the binding from.
const MinterRoleABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MinterRoleFuncSigs maps the 4-byte function signature to its string representation.
var MinterRoleFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"aa271e1a": "isMinter(address)",
	"98650275": "renounceMinter()",
}

// MinterRole is an auto generated Go binding around an Ethereum contract.
type MinterRole struct {
	MinterRoleCaller     // Read-only binding to the contract
	MinterRoleTransactor // Write-only binding to the contract
	MinterRoleFilterer   // Log filterer for contract events
}

// MinterRoleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MinterRoleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MinterRoleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MinterRoleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MinterRoleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MinterRoleSession struct {
	Contract     *MinterRole       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MinterRoleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MinterRoleCallerSession struct {
	Contract *MinterRoleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MinterRoleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MinterRoleTransactorSession struct {
	Contract     *MinterRoleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MinterRoleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MinterRoleRaw struct {
	Contract *MinterRole // Generic contract binding to access the raw methods on
}

// MinterRoleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MinterRoleCallerRaw struct {
	Contract *MinterRoleCaller // Generic read-only contract binding to access the raw methods on
}

// MinterRoleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MinterRoleTransactorRaw struct {
	Contract *MinterRoleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMinterRole creates a new instance of MinterRole, bound to a specific deployed contract.
func NewMinterRole(address common.Address, backend bind.ContractBackend) (*MinterRole, error) {
	contract, err := bindMinterRole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MinterRole{MinterRoleCaller: MinterRoleCaller{contract: contract}, MinterRoleTransactor: MinterRoleTransactor{contract: contract}, MinterRoleFilterer: MinterRoleFilterer{contract: contract}}, nil
}

// NewMinterRoleCaller creates a new read-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleCaller(address common.Address, caller bind.ContractCaller) (*MinterRoleCaller, error) {
	contract, err := bindMinterRole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleCaller{contract: contract}, nil
}

// NewMinterRoleTransactor creates a new write-only instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleTransactor(address common.Address, transactor bind.ContractTransactor) (*MinterRoleTransactor, error) {
	contract, err := bindMinterRole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MinterRoleTransactor{contract: contract}, nil
}

// NewMinterRoleFilterer creates a new log filterer instance of MinterRole, bound to a specific deployed contract.
func NewMinterRoleFilterer(address common.Address, filterer bind.ContractFilterer) (*MinterRoleFilterer, error) {
	contract, err := bindMinterRole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MinterRoleFilterer{contract: contract}, nil
}

// bindMinterRole binds a generic wrapper to an already deployed contract.
func bindMinterRole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MinterRoleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.MinterRoleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.MinterRoleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MinterRole *MinterRoleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MinterRole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MinterRole *MinterRoleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MinterRole *MinterRoleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MinterRole.Contract.contract.Transact(opts, method, params...)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MinterRole.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_MinterRole *MinterRoleCallerSession) IsMinter(account common.Address) (bool, error) {
	return _MinterRole.Contract.IsMinter(&_MinterRole.CallOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_MinterRole *MinterRoleTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _MinterRole.Contract.AddMinter(&_MinterRole.TransactOpts, account)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MinterRole.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_MinterRole *MinterRoleTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _MinterRole.Contract.RenounceMinter(&_MinterRole.TransactOpts)
}

// MinterRoleMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the MinterRole contract.
type MinterRoleMinterAddedIterator struct {
	Event *MinterRoleMinterAdded // Event containing the contract specifics and raw log

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
func (it *MinterRoleMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterAdded)
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
		it.Event = new(MinterRoleMinterAdded)
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
func (it *MinterRoleMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterAdded represents a MinterAdded event raised by the MinterRole contract.
type MinterRoleMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterAddedIterator{contract: _MinterRole.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterAdded)
				if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterAdded(log types.Log) (*MinterRoleMinterAdded, error) {
	event := new(MinterRoleMinterAdded)
	if err := _MinterRole.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// MinterRoleMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the MinterRole contract.
type MinterRoleMinterRemovedIterator struct {
	Event *MinterRoleMinterRemoved // Event containing the contract specifics and raw log

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
func (it *MinterRoleMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MinterRoleMinterRemoved)
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
		it.Event = new(MinterRoleMinterRemoved)
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
func (it *MinterRoleMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MinterRoleMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MinterRoleMinterRemoved represents a MinterRemoved event raised by the MinterRole contract.
type MinterRoleMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*MinterRoleMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &MinterRoleMinterRemovedIterator{contract: _MinterRole.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *MinterRoleMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _MinterRole.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MinterRoleMinterRemoved)
				if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_MinterRole *MinterRoleFilterer) ParseMinterRemoved(log types.Log) (*MinterRoleMinterRemoved, error) {
	event := new(MinterRoleMinterRemoved)
	if err := _MinterRole.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ModelsABI is the input ABI used to generate the binding from.
const ModelsABI = "[]"

// ModelsBin is the compiled bytecode used for deploying new contracts.
var ModelsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820a7a99bc55ac316de0ddeca140346ecf43f03f1b28dc687866837e6a85514c39664736f6c63430005100032"

// DeployModels deploys a new Ethereum contract, binding an instance of Models to it.
func DeployModels(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Models, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ModelsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// Models is an auto generated Go binding around an Ethereum contract.
type Models struct {
	ModelsCaller     // Read-only binding to the contract
	ModelsTransactor // Write-only binding to the contract
	ModelsFilterer   // Log filterer for contract events
}

// ModelsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ModelsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ModelsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ModelsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ModelsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ModelsSession struct {
	Contract     *Models           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ModelsCallerSession struct {
	Contract *ModelsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ModelsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ModelsTransactorSession struct {
	Contract     *ModelsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ModelsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ModelsRaw struct {
	Contract *Models // Generic contract binding to access the raw methods on
}

// ModelsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ModelsCallerRaw struct {
	Contract *ModelsCaller // Generic read-only contract binding to access the raw methods on
}

// ModelsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ModelsTransactorRaw struct {
	Contract *ModelsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewModels creates a new instance of Models, bound to a specific deployed contract.
func NewModels(address common.Address, backend bind.ContractBackend) (*Models, error) {
	contract, err := bindModels(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Models{ModelsCaller: ModelsCaller{contract: contract}, ModelsTransactor: ModelsTransactor{contract: contract}, ModelsFilterer: ModelsFilterer{contract: contract}}, nil
}

// NewModelsCaller creates a new read-only instance of Models, bound to a specific deployed contract.
func NewModelsCaller(address common.Address, caller bind.ContractCaller) (*ModelsCaller, error) {
	contract, err := bindModels(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsCaller{contract: contract}, nil
}

// NewModelsTransactor creates a new write-only instance of Models, bound to a specific deployed contract.
func NewModelsTransactor(address common.Address, transactor bind.ContractTransactor) (*ModelsTransactor, error) {
	contract, err := bindModels(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ModelsTransactor{contract: contract}, nil
}

// NewModelsFilterer creates a new log filterer instance of Models, bound to a specific deployed contract.
func NewModelsFilterer(address common.Address, filterer bind.ContractFilterer) (*ModelsFilterer, error) {
	contract, err := bindModels(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ModelsFilterer{contract: contract}, nil
}

// bindModels binds a generic wrapper to an already deployed contract.
func bindModels(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ModelsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.ModelsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.ModelsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Models *ModelsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Models.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Models *ModelsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Models.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Models *ModelsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Models.Contract.contract.Transact(opts, method, params...)
}

// RolesABI is the input ABI used to generate the binding from.
const RolesABI = "[]"

// RolesBin is the compiled bytecode used for deploying new contracts.
var RolesBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72315820b0f39eb352d7f2a878a4627faa4c44696ba384a0fcb5f7667add53552038060164736f6c63430005100032"

// DeployRoles deploys a new Ethereum contract, binding an instance of Roles to it.
func DeployRoles(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Roles, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RolesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// Roles is an auto generated Go binding around an Ethereum contract.
type Roles struct {
	RolesCaller     // Read-only binding to the contract
	RolesTransactor // Write-only binding to the contract
	RolesFilterer   // Log filterer for contract events
}

// RolesCaller is an auto generated read-only Go binding around an Ethereum contract.
type RolesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RolesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RolesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RolesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RolesSession struct {
	Contract     *Roles            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RolesCallerSession struct {
	Contract *RolesCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RolesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RolesTransactorSession struct {
	Contract     *RolesTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RolesRaw is an auto generated low-level Go binding around an Ethereum contract.
type RolesRaw struct {
	Contract *Roles // Generic contract binding to access the raw methods on
}

// RolesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RolesCallerRaw struct {
	Contract *RolesCaller // Generic read-only contract binding to access the raw methods on
}

// RolesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RolesTransactorRaw struct {
	Contract *RolesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoles creates a new instance of Roles, bound to a specific deployed contract.
func NewRoles(address common.Address, backend bind.ContractBackend) (*Roles, error) {
	contract, err := bindRoles(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Roles{RolesCaller: RolesCaller{contract: contract}, RolesTransactor: RolesTransactor{contract: contract}, RolesFilterer: RolesFilterer{contract: contract}}, nil
}

// NewRolesCaller creates a new read-only instance of Roles, bound to a specific deployed contract.
func NewRolesCaller(address common.Address, caller bind.ContractCaller) (*RolesCaller, error) {
	contract, err := bindRoles(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RolesCaller{contract: contract}, nil
}

// NewRolesTransactor creates a new write-only instance of Roles, bound to a specific deployed contract.
func NewRolesTransactor(address common.Address, transactor bind.ContractTransactor) (*RolesTransactor, error) {
	contract, err := bindRoles(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RolesTransactor{contract: contract}, nil
}

// NewRolesFilterer creates a new log filterer instance of Roles, bound to a specific deployed contract.
func NewRolesFilterer(address common.Address, filterer bind.ContractFilterer) (*RolesFilterer, error) {
	contract, err := bindRoles(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RolesFilterer{contract: contract}, nil
}

// bindRoles binds a generic wrapper to an already deployed contract.
func bindRoles(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RolesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.RolesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.RolesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Roles *RolesCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Roles.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Roles *RolesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Roles *RolesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Roles.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158203126bbdfc42af697f463be4aeee881e9677bcd4a7964ec59308b60659550308964736f6c63430005100032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// SupersymmetryABI is the input ABI used to generate the binding from.
const SupersymmetryABI = "[{\"inputs\":[{\"internalType\":\"address[5]\",\"name\":\"newAdmins\",\"type\":\"address[5]\"},{\"internalType\":\"uint8\",\"name\":\"newBftCoefficent\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"ethAssetId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"newTimeout\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"target\",\"type\":\"string\"}],\"name\":\"NewRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Return\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumModels.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"StatusChanged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"admins\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bftCoefficent\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8[5]\",\"name\":\"v\",\"type\":\"uint8[5]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"r\",\"type\":\"bytes32[5]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"s\",\"type\":\"bytes32[5]\"},{\"internalType\":\"enumModels.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"changeStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8[5]\",\"name\":\"v\",\"type\":\"uint8[5]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"r\",\"type\":\"bytes32[5]\"},{\"internalType\":\"bytes32[5]\",\"name\":\"s\",\"type\":\"bytes32[5]\"},{\"internalType\":\"enumModels.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"changeTokenStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"target\",\"type\":\"string\"},{\"internalType\":\"enumModels.RqType\",\"name\":\"rqType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"unlockAmount\",\"type\":\"uint256\"}],\"name\":\"createEthRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"target\",\"type\":\"string\"},{\"internalType\":\"enumModels.RqType\",\"name\":\"rqType\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"createTokenRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"registerInputToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"}],\"name\":\"registerNativeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"enumModels.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumModels.RqType\",\"name\":\"rType\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"height\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"target\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"assetId\",\"type\":\"string\"},{\"internalType\":\"enumModels.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"enumModels.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestHash\",\"type\":\"bytes32\"}],\"name\":\"withdrawPastDueReqest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SupersymmetryFuncSigs maps the 4-byte function signature to its string representation.
var SupersymmetryFuncSigs = map[string]string{
	"14bfd6d0": "admins(uint256)",
	"dce5e03b": "bftCoefficent()",
	"a8413510": "changeStatus(bytes32,uint8[5],bytes32[5],bytes32[5],uint8)",
	"f488f802": "changeTokenStatus(address,uint8[5],bytes32[5],bytes32[5],uint8)",
	"e4a56b2f": "createEthRequest(string,uint8,uint256)",
	"889347ea": "createTokenRequest(string,uint8,uint256,address)",
	"e122be20": "registerInputToken(string,string,string,uint8)",
	"6944afe7": "registerNativeToken(address,string)",
	"9d866985": "requests(bytes32)",
	"e4860339": "tokens(address)",
	"329881c6": "withdrawPastDueReqest(bytes32)",
}

// SupersymmetryBin is the compiled bytecode used for deploying new contracts.
var SupersymmetryBin = "0x60806040523480156200001157600080fd5b5060405162003f5838038062003f5883398181016040526101008110156200003857600080fd5b60a082015160c08301805160405192938501918591846401000000008211156200006157600080fd5b9083019060208201858111156200007757600080fd5b82516401000000008111828201881017156200009257600080fd5b82525081516020918201929091019080838360005b83811015620000c1578181015183820152602001620000a7565b50505050905090810190601f168015620000ef5780820380516001836020036101000a031916815260200191505b506040526020015185519092506001600160a01b0316151590506200014b576040805162461bcd60e51b815260206004820152600d60248201526c656d707479206164647265737360981b604482015290519081900360640190fd5b60208401516001600160a01b03166200019b576040805162461bcd60e51b815260206004820152600d60248201526c656d707479206164647265737360981b604482015290519081900360640190fd5b60408401516001600160a01b0316620001eb576040805162461bcd60e51b815260206004820152600d60248201526c656d707479206164647265737360981b604482015290519081900360640190fd5b60608401516001600160a01b03166200023b576040805162461bcd60e51b815260206004820152600d60248201526c656d707479206164647265737360981b604482015290519081900360640190fd5b60808401516001600160a01b03166200028b576040805162461bcd60e51b815260206004820152600d60248201526c656d707479206164647265737360981b604482015290519081900360640190fd5b604080516080810190915282815260208101600081526020016003815260126020918201526000805260088152815180517f5eff886ea0ce6ca488a3d6e336d6c0f75f46d19b42c06ce5ee98e42c96d256c792620002ee9284929101906200038a565b50602082015160018083018054909160ff199091169083818111156200031057fe5b0217905550604082015160018201805461ff0019166101008360048111156200033557fe5b0217905550606091909101516001918201805462ff000019166201000060ff9384160217905560008390556006805460ff19169186169190911790556200037f908560056200040f565b5050505050620004af565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003cd57805160ff1916838001178555620003fd565b82800160010185558215620003fd579182015b82811115620003fd578251825591602001919060010190620003e0565b506200040b92915062000468565b5090565b82600581019282156200045a579160200282015b828111156200045a57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000423565b506200040b92915062000488565b6200048591905b808211156200040b57600081556001016200046f565b90565b6200048591905b808211156200040b5780546001600160a01b03191681556001016200048f565b613a9980620004bf6000396000f3fe608060405260043610620000aa5760003560e01c8063a8413510116200006d578063a841351014620003e7578063dce5e03b14620004aa578063e122be2014620004d8578063e486033914620006a5578063e4a56b2f1462000790578063f488f802146200084257620000aa565b806314bfd6d014620000af578063329881c614620000f95780636944afe71462000129578063889347ea14620001f15780639d86698514620002d2575b600080fd5b348015620000bc57600080fd5b50620000dd60048036036020811015620000d557600080fd5b50356200090e565b604080516001600160a01b039092168252519081900360200190f35b3480156200010657600080fd5b5062000127600480360360208110156200011f57600080fd5b50356200092c565b005b3480156200013657600080fd5b5062000127600480360360408110156200014f57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156200017a57600080fd5b8201836020820111156200018d57600080fd5b803590602001918460018302840111600160201b83111715620001af57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955062000b6e945050505050565b348015620001fe57600080fd5b50620002c0600480360360808110156200021757600080fd5b810190602081018135600160201b8111156200023257600080fd5b8201836020820111156200024557600080fd5b803590602001918460018302840111600160201b831117156200026757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050813560ff16925050602081013590604001356001600160a01b031662000d1e565b60408051918252519081900360200190f35b348015620002df57600080fd5b506200030060048036036020811015620002f857600080fd5b50356200140d565b604051808860048111156200031157fe5b60ff1681526020018760038111156200032657fe5b60ff168152602001866001600160a01b03166001600160a01b0316815260200185815260200180602001848152602001836001600160a01b03166001600160a01b03168152602001828103825285818151815260200191508051906020019080838360005b83811015620003a55781810151838201526020016200038b565b50505050905090810190601f168015620003d35780820380516001836020036101000a031916815260200191505b509850505050505050505060405180910390f35b348015620003f457600080fd5b506200012760048036036102208110156200040e57600080fd5b6040805160a0818101909252833593928301929160c08301919060208401906005908390839080828437600092019190915250506040805160a081810190925292959493818101939250906005908390839080828437600092019190915250506040805160a08181019092529295949381810193925090600590839083908082843760009201919091525091945050503560ff169050620014f1565b348015620004b757600080fd5b50620004c262001d9a565b6040805160ff9092168252519081900360200190f35b348015620004e557600080fd5b506200012760048036036080811015620004fe57600080fd5b810190602081018135600160201b8111156200051957600080fd5b8201836020820111156200052c57600080fd5b803590602001918460018302840111600160201b831117156200054e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115620005a157600080fd5b820183602082011115620005b457600080fd5b803590602001918460018302840111600160201b83111715620005d657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156200062957600080fd5b8201836020820111156200063c57600080fd5b803590602001918460018302840111600160201b831117156200065e57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505050903560ff16915062001da39050565b348015620006b257600080fd5b50620006dc60048036036020811015620006cb57600080fd5b50356001600160a01b031662001f84565b6040518080602001856001811115620006f157fe5b60ff1681526020018460048111156200070657fe5b60ff1681526020018360ff1660ff168152602001828103825286818151815260200191508051906020019080838360005b838110156200075157818101518382015260200162000737565b50505050905090810190601f1680156200077f5780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b620002c060048036036060811015620007a857600080fd5b810190602081018135600160201b811115620007c357600080fd5b820183602082011115620007d657600080fd5b803590602001918460018302840111600160201b83111715620007f857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505060ff833516935050506020013562002045565b3480156200084f57600080fd5b506200012760048036036102208110156200086957600080fd5b6040805160a08181019092526001600160a01b0384351693928301929160c08301919060208401906005908390839080828437600092019190915250506040805160a081810190925292959493818101939250906005908390839080828437600092019190915250506040805160a08181019092529295949381810193925090600590839083908082843760009201919091525091945050503560ff169050620022ed565b600181600581106200091c57fe5b01546001600160a01b0316905081565b600160008281526007602052604090205460ff1660048111156200094c57fe5b1462000993576040805162461bcd60e51b81526020600482015260116024820152701c995c5d595cdd081b9bdd08195e1a5cdd607a1b604482015290519081900360640190fd5b600080548282526007602052604090912060010154014311620009ed576040805162461bcd60e51b815260206004820152600d60248201526c7271206e6f742065787069726560981b604482015290519081900360640190fd5b6000818152600760205260409020600401546001600160a01b031662000a655760008181526007602052604080822080546003909101549151620100009091046001600160a01b0316926108fc831502929190818181858888f1935050505015801562000a5e573d6000803e3d6000fd5b5062000b53565b6000818152600760209081526040808320600480820154825460039093015484516323b872dd60e01b815230938101939093526001600160a01b0362010000909404841660248401526044830152925191909216936323b872dd93606480850194919392918390030190829087803b15801562000ae157600080fd5b505af115801562000af6573d6000803e3d6000fd5b505050506040513d602081101562000b0d57600080fd5b505162000b53576040805162461bcd60e51b815260206004820152600f60248201526e696e76616c69642062616c616e636560881b604482015290519081900360640190fd5b6000908152600760205260409020805460ff19166004179055565b60006001600160a01b038316600090815260086020526040902060010154610100900460ff16600481111562000ba057fe5b1462000be9576040805162461bcd60e51b81526020600482015260136024820152726976616c696420746f6b656e4164647265737360681b604482015290519081900360640190fd5b6040805160808101909152818152602081016000815260200160018152602001836001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801562000c4357600080fd5b505afa15801562000c58573d6000803e3d6000fd5b505050506040513d602081101562000c6f57600080fd5b505160ff1690526001600160a01b038316600090815260086020908152604090912082518051919262000ca8928492909101906200266a565b50602082015160018083018054909160ff1990911690838181111562000cca57fe5b0217905550604082015160018201805461ff00191661010083600481111562000cef57fe5b0217905550606091909101516001909101805460ff909216620100000262ff0000199092169190911790555050565b60006001600160a01b03821662000d73576040805162461bcd60e51b8152602060048201526014602482015273696e76616c696420746f6b656e4164647265737360601b604482015290519081900360640190fd5b60036001600160a01b038316600090815260086020526040902060010154610100900460ff16600481111562000da557fe5b1462000dee576040805162461bcd60e51b81526020600482015260136024820152726976616c696420746f6b656e4164647265737360681b604482015290519081900360640190fd5b6000831162000e3d576040805162461bcd60e51b8152602060048201526016602482015275076616c7565206c657373206f7220657175616c7320360541b604482015290519081900360640190fd5b6001600160a01b038216600090815260086020526040812060019081015460ff169081111562000e6957fe5b14801562000e9a5750600084600381111562000e8157fe5b148062000e9a5750600184600381111562000e9857fe5b145b62000eec576040805162461bcd60e51b815260206004820181905260248201527f72712074797065206973206e6f7420657175616c7320746f6b656e2074797065604482015290519081900360640190fd5b6001600160a01b038216600090815260086020526040902060019081015460ff168181111562000f1857fe5b14801562000f495750600384600381111562000f3057fe5b148062000f495750600284600381111562000f4757fe5b145b62000f9b576040805162461bcd60e51b815260206004820181905260248201527f72712074797065206973206e6f7420657175616c7320746f6b656e2074797065604482015290519081900360640190fd5b600384600381111562000faa57fe5b148062000fc35750600084600381111562000fc157fe5b145b8015620010505750604080516323b872dd60e01b81523360048201523060248201526044810185905290516001600160a01b038416916323b872dd9160648083019260209291908290030181600087803b1580156200102157600080fd5b505af115801562001036573d6000803e3d6000fd5b505050506040513d60208110156200104d57600080fd5b50515b62001094576040805162461bcd60e51b815260206004820152600f60248201526e696e76616c69642062616c616e636560881b604482015290519081900360640190fd5b60006002303388874360405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b815260140184805190602001908083835b60208310620011075780518252601f199092019160209182019101620010e6565b51815160209384036101000a60001901801990921691161790529201948552508381019290925250604080518084038301815292810190819052825192955093508392508401908083835b60208310620011735780518252601f19909201916020918201910162001152565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa158015620011b3573d6000803e3d6000fd5b5050506040513d6020811015620011c957600080fd5b505190506000808281526007602052604090205460ff166004811115620011ec57fe5b146200122f576040805162461bcd60e51b815260206004820152600d60248201526c1c995c5d595cdd08195e1a5cdd609a1b604482015290519081900360640190fd5b6040805160e0810190915280600181526020018660038111156200124f57fe5b81523360208083019190915243604080840191909152606083018a9052608083018890526001600160a01b03871660a090930192909252600084815260079091522081518154829060ff19166001836004811115620012aa57fe5b021790555060208201518154829061ff001916610100836003811115620012cd57fe5b0217905550604082015181546001600160a01b03909116620100000262010000600160b01b03199091161781556060820151600182015560808201518051620013219160028401916020909101906200266a565b5060a0820151600382015560c090910151600490910180546001600160a01b0319166001600160a01b039092169190911790556040805182815233602080830182905260609383018481528a519484019490945289517fe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c1158969486948c939092608084019185019080838360005b83811015620013c7578181015183820152602001620013ad565b50505050905090810190601f168015620013f55780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a195945050505050565b60076020908152600091825260409182902080546001808301546002808501805488516101009582161586026000190190911692909204601f810188900488028301880190985287825260ff808616989486041696620100009095046001600160a01b031695929492939192909190830182828015620014d15780601f10620014a557610100808354040283529160200191620014d1565b820191906000526020600020905b815481529060010190602001808311620014b357829003601f168201915b5050505060038301546004909301549192916001600160a01b0316905087565b600160008681526007602052604090205460ff1660048111156200151157fe5b1462001558576040805162461bcd60e51b8152602060048201526011602482015270737461747573206973206e6f77206e657760781b604482015290519081900360640190fd5b60038160048111156200156757fe5b148062001580575060028160048111156200157e57fe5b145b620015c3576040805162461bcd60e51b815260206004820152600e60248201526d696e76616c69642073746174757360901b604482015290519081900360640190fd5b600080548682526007602052604090912060010154014311156200161a576040805162461bcd60e51b815260206004820152600960248201526872712065787069726560b81b604482015290519081900360640190fd5b6000858152600760205260408120600401546001600160a01b031690805b6005811015620017f457600181600581106200165057fe5b01546001600160a01b03166001898660048111156200166b57fe5b604051602001808381526020018260ff1660ff1660f81b81526001019250505060405160208183030381529060405260405160200180807f19457468657265756d205369676e6564204d6573736167653a0a333300000000815250601c0182805190602001908083835b60208310620016f65780518252601f199092019160209182019101620016d5565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001208984600581106200173e57fe5b60200201518985600581106200175057fe5b60200201518986600581106200176257fe5b602002015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015620017bf573d6000803e3d6000fd5b505050602060405103516001600160a01b031614620017e0576000620017e3565b60015b60ff16919091019060010162001638565b5060065460ff168112156200183b5760405162461bcd60e51b815260040180806020018281038252602781526020018062003a3e6027913960400191505060405180910390fd5b60038360048111156200184a57fe5b141562001ae6576002600088815260076020526040902054610100900460ff1660038111156200187657fe5b141562001965576000878152600760209081526040808320805460039091015482516340c10f1960e01b81526001600160a01b03620100009093048316600482015260248101919091529151908616936340c10f1993604480850194919392918390030190829087803b158015620018ed57600080fd5b505af115801562001902573d6000803e3d6000fd5b505050506040513d60208110156200191957600080fd5b50516200195f576040805162461bcd60e51b815260206004820152600f60248201526e696e76616c69642062616c616e636560881b604482015290519081900360640190fd5b62001ae0565b6001600088815260076020526040902054610100900460ff1660038111156200198a57fe5b141562001ae0576001600160a01b038216620019f85760008781526007602052604080822080546003909101549151620100009091046001600160a01b0316926108fc831502929190818181858888f19350505050158015620019f1573d6000803e3d6000fd5b5062001ae0565b6000878152600760209081526040808320805460039091015482516323b872dd60e01b81523060048201526001600160a01b03620100009093048316602482015260448101919091529151908616936323b872dd93606480850194919392918390030190829087803b15801562001a6e57600080fd5b505af115801562001a83573d6000803e3d6000fd5b505050506040513d602081101562001a9a57600080fd5b505162001ae0576040805162461bcd60e51b815260206004820152600f60248201526e696e76616c69642062616c616e636560881b604482015290519081900360640190fd5b62001d18565b600283600481111562001af557fe5b141562001d18576003600088815260076020526040902054610100900460ff16600381111562001b2157fe5b141562001b9e576000878152600760209081526040808320805460039091015482516323b872dd60e01b81523060048201526001600160a01b03620100009093048316602482015260448101919091529151908616936323b872dd93606480850194919392918390030190829087803b15801562001a6e57600080fd5b60008088815260076020526040902054610100900460ff16600381111562001bc257fe5b141562001d18576001600160a01b03821662001c305760008781526007602052604080822080546003909101549151620100009091046001600160a01b0316926108fc831502929190818181858888f1935050505015801562001c29573d6000803e3d6000fd5b5062001d18565b6000878152600760209081526040808320805460039091015482516323b872dd60e01b81523060048201526001600160a01b03620100009093048316602482015260448101919091529151908616936323b872dd93606480850194919392918390030190829087803b15801562001ca657600080fd5b505af115801562001cbb573d6000803e3d6000fd5b505050506040513d602081101562001cd257600080fd5b505162001d18576040805162461bcd60e51b815260206004820152600f60248201526e696e76616c69642062616c616e636560881b604482015290519081900360640190fd5b6000878152600760205260409020805484919060ff1916600183600481111562001d3e57fe5b02179055507f22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c87846040518083815260200182600481111562001d7d57fe5b60ff1681526020019250505060405180910390a150505050505050565b60065460ff1681565b600083838360405162001db690620026ef565b60ff82166040820152606080825284519082015283518190602080830191608084019188019080838360005b8381101562001dfc57818101518382015260200162001de2565b50505050905090810190601f16801562001e2a5780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b8381101562001e5f57818101518382015260200162001e45565b50505050905090810190601f16801562001e8d5780820380516001836020036101000a031916815260200191505b5095505050505050604051809103906000f08015801562001eb2573d6000803e3d6000fd5b50604080516080810190915286815290915060208101600181526020016001815260ff84166020918201526001600160a01b0383166000908152600882526040902082518051919262001f0b928492909101906200266a565b50602082015160018083018054909160ff1990911690838181111562001f2d57fe5b0217905550604082015160018201805461ff00191661010083600481111562001f5257fe5b0217905550606091909101516001909101805460ff909216620100000262ff0000199092169190911790555050505050565b60086020908152600091825260409182902080548351601f600260001961010060018616150201909316929092049182018490048402810184019094528084529092918391908301828280156200201f5780601f1062001ff3576101008083540402835291602001916200201f565b820191906000526020600020905b8154815290600101906020018083116200200157829003601f168201915b5050506001909301549192505060ff808216916101008104821691620100009091041684565b6000808360038111156200205557fe5b14806200206e575060018360038111156200206c57fe5b145b620020c0576040805162461bcd60e51b815260206004820181905260248201527f72712074797065206973206e6f7420657175616c7320746f6b656e2074797065604482015290519081900360640190fd5b600080846003811115620020d057fe5b14620020dd5782620020df565b345b905060006002303388854360405160200180866001600160a01b03166001600160a01b031660601b8152601401856001600160a01b03166001600160a01b031660601b815260140184805190602001908083835b60208310620021545780518252601f19909201916020918201910162002133565b51815160209384036101000a60001901801990921691161790529201948552508381019290925250604080518084038301815292810190819052825192955093508392508401908083835b60208310620021c05780518252601f1990920191602091820191016200219f565b51815160209384036101000a60001901801990921691161790526040519190930194509192505080830381855afa15801562002200573d6000803e3d6000fd5b5050506040513d60208110156200221657600080fd5b505190506000808281526007602052604090205460ff1660048111156200223957fe5b146200227c576040805162461bcd60e51b815260206004820152600d60248201526c1c995c5d595cdd08195e1a5cdd609a1b604482015290519081900360640190fd5b6040805160e0810190915280600181526020018660038111156200229c57fe5b81523360208083019190915243604080840191909152606083018a905260808301869052600060a09093018390528483526007909152902081518154829060ff19166001836004811115620012aa57fe5b60016001600160a01b038616600090815260086020526040902060010154610100900460ff1660048111156200231f57fe5b1462002368576040805162461bcd60e51b81526020600482015260136024820152726976616c696420746f6b656e4164647265737360681b604482015290519081900360640190fd5b60038160048111156200237757fe5b14620023ba576040805162461bcd60e51b815260206004820152600d60248201526c6976616c69642073746174757360981b604482015290519081900360640190fd5b6002816004811115620023c957fe5b146200240c576040805162461bcd60e51b815260206004820152600d60248201526c6976616c69642073746174757360981b604482015290519081900360640190fd5b6000805b6005811015620025e157600181600581106200242857fe5b01546001600160a01b03166001888560048111156200244357fe5b60405160200180836001600160a01b03166001600160a01b031660601b81526014018260ff1660ff1660f81b81526001019250505060405160208183030381529060405260405160200180807f19457468657265756d205369676e6564204d6573736167653a0a323200000000815250601c0182805190602001908083835b60208310620024e35780518252601f199092019160209182019101620024c2565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051602081830303815290604052805190602001208884600581106200252b57fe5b60200201518885600581106200253d57fe5b60200201518886600581106200254f57fe5b602002015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa158015620025ac573d6000803e3d6000fd5b505050602060405103516001600160a01b031614620025cd576000620025d0565b60015b60ff16919091019060010162002410565b5060065460ff16811215620026285760405162461bcd60e51b815260040180806020018281038252602781526020018062003a3e6027913960400191505060405180910390fd5b6001600160a01b0386166000908152600860205260409020600101805483919061ff0019166101008360048111156200265d57fe5b0217905550505050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620026ad57805160ff1916838001178555620026dd565b82800160010185558215620026dd579182015b82811115620026dd578251825591602001919060010190620026c0565b50620026eb929150620026fd565b5090565b611320806200271e83390190565b6200271a91905b80821115620026eb576000815560010162002704565b9056fe60806040523480156200001157600080fd5b506040516200132038038062001320833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b50604052602001519150620001ca9050620001bb6001600160e01b036200021c16565b6001600160e01b036200022116565b8251620001df90600490602086019062000369565b508151620001f590600590602085019062000369565b506006805460ff191660ff831617905562000213620001bb6200021c565b5050506200040b565b335b90565b6200023c8160036200027360201b62000c521790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6200028882826001600160e01b036200030016565b15620002db576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620003495760405162461bcd60e51b8152600401808060200182810382526022815260200180620012fe6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003ac57805160ff1916838001178555620003dc565b82800160010185558215620003dc579182015b82811115620003dc578251825591602001919060010190620003bf565b50620003ea929150620003ee565b5090565b6200021e91905b80821115620003ea5760008155600101620003f5565b610ee3806200041b6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d7146102db578063a9059cbb14610307578063aa271e1a14610333578063dd62ed3e14610359576100f5565b806370a082311461027d57806395d89b41146102a3578063983b2d56146102ab57806398650275146102d3576100f5565b806323b872dd116100d357806323b872dd146101d1578063313ce56714610207578063395093511461022557806340c10f1914610251576100f5565b806306fdde03146100fa578063095ea7b31461017757806318160ddd146101b7575b600080fd5b610102610387565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013c578181015183820152602001610124565b50505050905090810190601f1680156101695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a36004803603604081101561018d57600080fd5b506001600160a01b03813516906020013561041d565b604080519115158252519081900360200190f35b6101bf61043a565b60408051918252519081900360200190f35b6101a3600480360360608110156101e757600080fd5b506001600160a01b03813581169160208101359091169060400135610440565b61020f6104cd565b6040805160ff9092168252519081900360200190f35b6101a36004803603604081101561023b57600080fd5b506001600160a01b0381351690602001356104d6565b6101a36004803603604081101561026757600080fd5b506001600160a01b03813516906020013561052a565b6101bf6004803603602081101561029357600080fd5b50356001600160a01b0316610581565b61010261059c565b6102d1600480360360208110156102c157600080fd5b50356001600160a01b03166105fd565b005b6102d161064f565b6101a3600480360360408110156102f157600080fd5b506001600160a01b038135169060200135610661565b6101a36004803603604081101561031d57600080fd5b506001600160a01b0381351690602001356106cf565b6101a36004803603602081101561034957600080fd5b50356001600160a01b03166106e3565b6101bf6004803603604081101561036f57600080fd5b506001600160a01b03813581169160200135166106fc565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104135780601f106103e857610100808354040283529160200191610413565b820191906000526020600020905b8154815290600101906020018083116103f657829003601f168201915b5050505050905090565b600061043161042a610727565b848461072b565b50600192915050565b60025490565b600061044d848484610817565b6104c384610459610727565b6104be85604051806060016040528060288152602001610df7602891396001600160a01b038a16600090815260016020526040812090610497610727565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61097316565b61072b565b5060019392505050565b60065460ff1690565b60006104316104e3610727565b846104be85600160006104f4610727565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff610a0a16565b600061053c610537610727565b6106e3565b6105775760405162461bcd60e51b8152600401808060200182810382526030815260200180610da66030913960400191505060405180910390fd5b6104318383610a6b565b6001600160a01b031660009081526020819052604090205490565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104135780601f106103e857610100808354040283529160200191610413565b610608610537610727565b6106435760405162461bcd60e51b8152600401808060200182810382526030815260200180610da66030913960400191505060405180910390fd5b61064c81610b5b565b50565b61065f61065a610727565b610ba3565b565b600061043161066e610727565b846104be85604051806060016040528060258152602001610e8a6025913960016000610698610727565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61097316565b60006104316106dc610727565b8484610817565b60006106f660038363ffffffff610beb16565b92915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166107705760405162461bcd60e51b8152600401808060200182810382526024815260200180610e666024913960400191505060405180910390fd5b6001600160a01b0382166107b55760405162461bcd60e51b8152600401808060200182810382526022815260200180610d5e6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661085c5760405162461bcd60e51b8152600401808060200182810382526025815260200180610e416025913960400191505060405180910390fd5b6001600160a01b0382166108a15760405162461bcd60e51b8152600401808060200182810382526023815260200180610d3b6023913960400191505060405180910390fd5b6108e481604051806060016040528060268152602001610d80602691396001600160a01b038616600090815260208190526040902054919063ffffffff61097316565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610919908263ffffffff610a0a16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610a025760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109c75781810151838201526020016109af565b50505050905090810190601f1680156109f45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610a64576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b038216610ac6576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254610ad9908263ffffffff610a0a16565b6002556001600160a01b038216600090815260208190526040902054610b05908263ffffffff610a0a16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b610b6c60038263ffffffff610c5216565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610bb460038263ffffffff610cd316565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610c325760405162461bcd60e51b8152600401808060200182810382526022815260200180610e1f6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610c5c8282610beb565b15610cae576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b610cdd8282610beb565b610d185760405162461bcd60e51b8152600401808060200182810382526021815260200180610dd66021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff1916905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820a0fd174d9fb4a843cfda8d411e74be8e3d289af97e1c49347b4785936176637264736f6c63430005100032526f6c65733a206163636f756e7420697320746865207a65726f206164647265737361646d696e7320766f746520636f756e74206973206c65737320626674436f6566666963656e74a265627a7a723158205dd37db9062c20769cc1d37cee1644cb7e07fe61f536e2eadf6fd77300c5da0c64736f6c63430005100032"

// DeploySupersymmetry deploys a new Ethereum contract, binding an instance of Supersymmetry to it.
func DeploySupersymmetry(auth *bind.TransactOpts, backend bind.ContractBackend, newAdmins [5]common.Address, newBftCoefficent uint8, ethAssetId string, newTimeout *big.Int) (common.Address, *types.Transaction, *Supersymmetry, error) {
	parsed, err := abi.JSON(strings.NewReader(SupersymmetryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SupersymmetryBin), backend, newAdmins, newBftCoefficent, ethAssetId, newTimeout)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Supersymmetry{SupersymmetryCaller: SupersymmetryCaller{contract: contract}, SupersymmetryTransactor: SupersymmetryTransactor{contract: contract}, SupersymmetryFilterer: SupersymmetryFilterer{contract: contract}}, nil
}

// Supersymmetry is an auto generated Go binding around an Ethereum contract.
type Supersymmetry struct {
	SupersymmetryCaller     // Read-only binding to the contract
	SupersymmetryTransactor // Write-only binding to the contract
	SupersymmetryFilterer   // Log filterer for contract events
}

// SupersymmetryCaller is an auto generated read-only Go binding around an Ethereum contract.
type SupersymmetryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupersymmetryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SupersymmetryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupersymmetryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SupersymmetryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SupersymmetrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SupersymmetrySession struct {
	Contract     *Supersymmetry    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SupersymmetryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SupersymmetryCallerSession struct {
	Contract *SupersymmetryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SupersymmetryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SupersymmetryTransactorSession struct {
	Contract     *SupersymmetryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SupersymmetryRaw is an auto generated low-level Go binding around an Ethereum contract.
type SupersymmetryRaw struct {
	Contract *Supersymmetry // Generic contract binding to access the raw methods on
}

// SupersymmetryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SupersymmetryCallerRaw struct {
	Contract *SupersymmetryCaller // Generic read-only contract binding to access the raw methods on
}

// SupersymmetryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SupersymmetryTransactorRaw struct {
	Contract *SupersymmetryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSupersymmetry creates a new instance of Supersymmetry, bound to a specific deployed contract.
func NewSupersymmetry(address common.Address, backend bind.ContractBackend) (*Supersymmetry, error) {
	contract, err := bindSupersymmetry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Supersymmetry{SupersymmetryCaller: SupersymmetryCaller{contract: contract}, SupersymmetryTransactor: SupersymmetryTransactor{contract: contract}, SupersymmetryFilterer: SupersymmetryFilterer{contract: contract}}, nil
}

// NewSupersymmetryCaller creates a new read-only instance of Supersymmetry, bound to a specific deployed contract.
func NewSupersymmetryCaller(address common.Address, caller bind.ContractCaller) (*SupersymmetryCaller, error) {
	contract, err := bindSupersymmetry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SupersymmetryCaller{contract: contract}, nil
}

// NewSupersymmetryTransactor creates a new write-only instance of Supersymmetry, bound to a specific deployed contract.
func NewSupersymmetryTransactor(address common.Address, transactor bind.ContractTransactor) (*SupersymmetryTransactor, error) {
	contract, err := bindSupersymmetry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SupersymmetryTransactor{contract: contract}, nil
}

// NewSupersymmetryFilterer creates a new log filterer instance of Supersymmetry, bound to a specific deployed contract.
func NewSupersymmetryFilterer(address common.Address, filterer bind.ContractFilterer) (*SupersymmetryFilterer, error) {
	contract, err := bindSupersymmetry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SupersymmetryFilterer{contract: contract}, nil
}

// bindSupersymmetry binds a generic wrapper to an already deployed contract.
func bindSupersymmetry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SupersymmetryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supersymmetry *SupersymmetryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Supersymmetry.Contract.SupersymmetryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supersymmetry *SupersymmetryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supersymmetry.Contract.SupersymmetryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supersymmetry *SupersymmetryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supersymmetry.Contract.SupersymmetryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Supersymmetry *SupersymmetryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Supersymmetry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Supersymmetry *SupersymmetryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Supersymmetry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Supersymmetry *SupersymmetryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Supersymmetry.Contract.contract.Transact(opts, method, params...)
}

// Admins is a free data retrieval call binding the contract method 0x14bfd6d0.
//
// Solidity: function admins(uint256 ) constant returns(address)
func (_Supersymmetry *SupersymmetryCaller) Admins(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Supersymmetry.contract.Call(opts, out, "admins", arg0)
	return *ret0, err
}

// Admins is a free data retrieval call binding the contract method 0x14bfd6d0.
//
// Solidity: function admins(uint256 ) constant returns(address)
func (_Supersymmetry *SupersymmetrySession) Admins(arg0 *big.Int) (common.Address, error) {
	return _Supersymmetry.Contract.Admins(&_Supersymmetry.CallOpts, arg0)
}

// Admins is a free data retrieval call binding the contract method 0x14bfd6d0.
//
// Solidity: function admins(uint256 ) constant returns(address)
func (_Supersymmetry *SupersymmetryCallerSession) Admins(arg0 *big.Int) (common.Address, error) {
	return _Supersymmetry.Contract.Admins(&_Supersymmetry.CallOpts, arg0)
}

// BftCoefficent is a free data retrieval call binding the contract method 0xdce5e03b.
//
// Solidity: function bftCoefficent() constant returns(uint8)
func (_Supersymmetry *SupersymmetryCaller) BftCoefficent(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Supersymmetry.contract.Call(opts, out, "bftCoefficent")
	return *ret0, err
}

// BftCoefficent is a free data retrieval call binding the contract method 0xdce5e03b.
//
// Solidity: function bftCoefficent() constant returns(uint8)
func (_Supersymmetry *SupersymmetrySession) BftCoefficent() (uint8, error) {
	return _Supersymmetry.Contract.BftCoefficent(&_Supersymmetry.CallOpts)
}

// BftCoefficent is a free data retrieval call binding the contract method 0xdce5e03b.
//
// Solidity: function bftCoefficent() constant returns(uint8)
func (_Supersymmetry *SupersymmetryCallerSession) BftCoefficent() (uint8, error) {
	return _Supersymmetry.Contract.BftCoefficent(&_Supersymmetry.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(uint8 status, uint8 rType, address owner, uint256 height, string target, uint256 tokenAmount, address tokenAddress)
func (_Supersymmetry *SupersymmetryCaller) Requests(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Status       uint8
	RType        uint8
	Owner        common.Address
	Height       *big.Int
	Target       string
	TokenAmount  *big.Int
	TokenAddress common.Address
}, error) {
	ret := new(struct {
		Status       uint8
		RType        uint8
		Owner        common.Address
		Height       *big.Int
		Target       string
		TokenAmount  *big.Int
		TokenAddress common.Address
	})
	out := ret
	err := _Supersymmetry.contract.Call(opts, out, "requests", arg0)
	return *ret, err
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(uint8 status, uint8 rType, address owner, uint256 height, string target, uint256 tokenAmount, address tokenAddress)
func (_Supersymmetry *SupersymmetrySession) Requests(arg0 [32]byte) (struct {
	Status       uint8
	RType        uint8
	Owner        common.Address
	Height       *big.Int
	Target       string
	TokenAmount  *big.Int
	TokenAddress common.Address
}, error) {
	return _Supersymmetry.Contract.Requests(&_Supersymmetry.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 ) constant returns(uint8 status, uint8 rType, address owner, uint256 height, string target, uint256 tokenAmount, address tokenAddress)
func (_Supersymmetry *SupersymmetryCallerSession) Requests(arg0 [32]byte) (struct {
	Status       uint8
	RType        uint8
	Owner        common.Address
	Height       *big.Int
	Target       string
	TokenAmount  *big.Int
	TokenAddress common.Address
}, error) {
	return _Supersymmetry.Contract.Requests(&_Supersymmetry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) constant returns(string assetId, uint8 tokenType, uint8 status, uint8 decimals)
func (_Supersymmetry *SupersymmetryCaller) Tokens(opts *bind.CallOpts, arg0 common.Address) (struct {
	AssetId   string
	TokenType uint8
	Status    uint8
	Decimals  uint8
}, error) {
	ret := new(struct {
		AssetId   string
		TokenType uint8
		Status    uint8
		Decimals  uint8
	})
	out := ret
	err := _Supersymmetry.contract.Call(opts, out, "tokens", arg0)
	return *ret, err
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) constant returns(string assetId, uint8 tokenType, uint8 status, uint8 decimals)
func (_Supersymmetry *SupersymmetrySession) Tokens(arg0 common.Address) (struct {
	AssetId   string
	TokenType uint8
	Status    uint8
	Decimals  uint8
}, error) {
	return _Supersymmetry.Contract.Tokens(&_Supersymmetry.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0xe4860339.
//
// Solidity: function tokens(address ) constant returns(string assetId, uint8 tokenType, uint8 status, uint8 decimals)
func (_Supersymmetry *SupersymmetryCallerSession) Tokens(arg0 common.Address) (struct {
	AssetId   string
	TokenType uint8
	Status    uint8
	Decimals  uint8
}, error) {
	return _Supersymmetry.Contract.Tokens(&_Supersymmetry.CallOpts, arg0)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xa8413510.
//
// Solidity: function changeStatus(bytes32 requestHash, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 status) returns()
func (_Supersymmetry *SupersymmetryTransactor) ChangeStatus(opts *bind.TransactOpts, requestHash [32]byte, v [5]uint8, r [5][32]byte, s [5][32]byte, status uint8) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "changeStatus", requestHash, v, r, s, status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xa8413510.
//
// Solidity: function changeStatus(bytes32 requestHash, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 status) returns()
func (_Supersymmetry *SupersymmetrySession) ChangeStatus(requestHash [32]byte, v [5]uint8, r [5][32]byte, s [5][32]byte, status uint8) (*types.Transaction, error) {
	return _Supersymmetry.Contract.ChangeStatus(&_Supersymmetry.TransactOpts, requestHash, v, r, s, status)
}

// ChangeStatus is a paid mutator transaction binding the contract method 0xa8413510.
//
// Solidity: function changeStatus(bytes32 requestHash, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 status) returns()
func (_Supersymmetry *SupersymmetryTransactorSession) ChangeStatus(requestHash [32]byte, v [5]uint8, r [5][32]byte, s [5][32]byte, status uint8) (*types.Transaction, error) {
	return _Supersymmetry.Contract.ChangeStatus(&_Supersymmetry.TransactOpts, requestHash, v, r, s, status)
}

// ChangeTokenStatus is a paid mutator transaction binding the contract method 0xf488f802.
//
// Solidity: function changeTokenStatus(address tokenAddress, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 status) returns()
func (_Supersymmetry *SupersymmetryTransactor) ChangeTokenStatus(opts *bind.TransactOpts, tokenAddress common.Address, v [5]uint8, r [5][32]byte, s [5][32]byte, status uint8) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "changeTokenStatus", tokenAddress, v, r, s, status)
}

// ChangeTokenStatus is a paid mutator transaction binding the contract method 0xf488f802.
//
// Solidity: function changeTokenStatus(address tokenAddress, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 status) returns()
func (_Supersymmetry *SupersymmetrySession) ChangeTokenStatus(tokenAddress common.Address, v [5]uint8, r [5][32]byte, s [5][32]byte, status uint8) (*types.Transaction, error) {
	return _Supersymmetry.Contract.ChangeTokenStatus(&_Supersymmetry.TransactOpts, tokenAddress, v, r, s, status)
}

// ChangeTokenStatus is a paid mutator transaction binding the contract method 0xf488f802.
//
// Solidity: function changeTokenStatus(address tokenAddress, uint8[5] v, bytes32[5] r, bytes32[5] s, uint8 status) returns()
func (_Supersymmetry *SupersymmetryTransactorSession) ChangeTokenStatus(tokenAddress common.Address, v [5]uint8, r [5][32]byte, s [5][32]byte, status uint8) (*types.Transaction, error) {
	return _Supersymmetry.Contract.ChangeTokenStatus(&_Supersymmetry.TransactOpts, tokenAddress, v, r, s, status)
}

// CreateEthRequest is a paid mutator transaction binding the contract method 0xe4a56b2f.
//
// Solidity: function createEthRequest(string target, uint8 rqType, uint256 unlockAmount) returns(bytes32)
func (_Supersymmetry *SupersymmetryTransactor) CreateEthRequest(opts *bind.TransactOpts, target string, rqType uint8, unlockAmount *big.Int) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "createEthRequest", target, rqType, unlockAmount)
}

// CreateEthRequest is a paid mutator transaction binding the contract method 0xe4a56b2f.
//
// Solidity: function createEthRequest(string target, uint8 rqType, uint256 unlockAmount) returns(bytes32)
func (_Supersymmetry *SupersymmetrySession) CreateEthRequest(target string, rqType uint8, unlockAmount *big.Int) (*types.Transaction, error) {
	return _Supersymmetry.Contract.CreateEthRequest(&_Supersymmetry.TransactOpts, target, rqType, unlockAmount)
}

// CreateEthRequest is a paid mutator transaction binding the contract method 0xe4a56b2f.
//
// Solidity: function createEthRequest(string target, uint8 rqType, uint256 unlockAmount) returns(bytes32)
func (_Supersymmetry *SupersymmetryTransactorSession) CreateEthRequest(target string, rqType uint8, unlockAmount *big.Int) (*types.Transaction, error) {
	return _Supersymmetry.Contract.CreateEthRequest(&_Supersymmetry.TransactOpts, target, rqType, unlockAmount)
}

// CreateTokenRequest is a paid mutator transaction binding the contract method 0x889347ea.
//
// Solidity: function createTokenRequest(string target, uint8 rqType, uint256 amount, address tokenAddress) returns(bytes32)
func (_Supersymmetry *SupersymmetryTransactor) CreateTokenRequest(opts *bind.TransactOpts, target string, rqType uint8, amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "createTokenRequest", target, rqType, amount, tokenAddress)
}

// CreateTokenRequest is a paid mutator transaction binding the contract method 0x889347ea.
//
// Solidity: function createTokenRequest(string target, uint8 rqType, uint256 amount, address tokenAddress) returns(bytes32)
func (_Supersymmetry *SupersymmetrySession) CreateTokenRequest(target string, rqType uint8, amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Supersymmetry.Contract.CreateTokenRequest(&_Supersymmetry.TransactOpts, target, rqType, amount, tokenAddress)
}

// CreateTokenRequest is a paid mutator transaction binding the contract method 0x889347ea.
//
// Solidity: function createTokenRequest(string target, uint8 rqType, uint256 amount, address tokenAddress) returns(bytes32)
func (_Supersymmetry *SupersymmetryTransactorSession) CreateTokenRequest(target string, rqType uint8, amount *big.Int, tokenAddress common.Address) (*types.Transaction, error) {
	return _Supersymmetry.Contract.CreateTokenRequest(&_Supersymmetry.TransactOpts, target, rqType, amount, tokenAddress)
}

// RegisterInputToken is a paid mutator transaction binding the contract method 0xe122be20.
//
// Solidity: function registerInputToken(string assetId, string name, string symbol, uint8 decimals) returns()
func (_Supersymmetry *SupersymmetryTransactor) RegisterInputToken(opts *bind.TransactOpts, assetId string, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "registerInputToken", assetId, name, symbol, decimals)
}

// RegisterInputToken is a paid mutator transaction binding the contract method 0xe122be20.
//
// Solidity: function registerInputToken(string assetId, string name, string symbol, uint8 decimals) returns()
func (_Supersymmetry *SupersymmetrySession) RegisterInputToken(assetId string, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Supersymmetry.Contract.RegisterInputToken(&_Supersymmetry.TransactOpts, assetId, name, symbol, decimals)
}

// RegisterInputToken is a paid mutator transaction binding the contract method 0xe122be20.
//
// Solidity: function registerInputToken(string assetId, string name, string symbol, uint8 decimals) returns()
func (_Supersymmetry *SupersymmetryTransactorSession) RegisterInputToken(assetId string, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _Supersymmetry.Contract.RegisterInputToken(&_Supersymmetry.TransactOpts, assetId, name, symbol, decimals)
}

// RegisterNativeToken is a paid mutator transaction binding the contract method 0x6944afe7.
//
// Solidity: function registerNativeToken(address tokenAddress, string assetId) returns()
func (_Supersymmetry *SupersymmetryTransactor) RegisterNativeToken(opts *bind.TransactOpts, tokenAddress common.Address, assetId string) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "registerNativeToken", tokenAddress, assetId)
}

// RegisterNativeToken is a paid mutator transaction binding the contract method 0x6944afe7.
//
// Solidity: function registerNativeToken(address tokenAddress, string assetId) returns()
func (_Supersymmetry *SupersymmetrySession) RegisterNativeToken(tokenAddress common.Address, assetId string) (*types.Transaction, error) {
	return _Supersymmetry.Contract.RegisterNativeToken(&_Supersymmetry.TransactOpts, tokenAddress, assetId)
}

// RegisterNativeToken is a paid mutator transaction binding the contract method 0x6944afe7.
//
// Solidity: function registerNativeToken(address tokenAddress, string assetId) returns()
func (_Supersymmetry *SupersymmetryTransactorSession) RegisterNativeToken(tokenAddress common.Address, assetId string) (*types.Transaction, error) {
	return _Supersymmetry.Contract.RegisterNativeToken(&_Supersymmetry.TransactOpts, tokenAddress, assetId)
}

// WithdrawPastDueReqest is a paid mutator transaction binding the contract method 0x329881c6.
//
// Solidity: function withdrawPastDueReqest(bytes32 requestHash) returns()
func (_Supersymmetry *SupersymmetryTransactor) WithdrawPastDueReqest(opts *bind.TransactOpts, requestHash [32]byte) (*types.Transaction, error) {
	return _Supersymmetry.contract.Transact(opts, "withdrawPastDueReqest", requestHash)
}

// WithdrawPastDueReqest is a paid mutator transaction binding the contract method 0x329881c6.
//
// Solidity: function withdrawPastDueReqest(bytes32 requestHash) returns()
func (_Supersymmetry *SupersymmetrySession) WithdrawPastDueReqest(requestHash [32]byte) (*types.Transaction, error) {
	return _Supersymmetry.Contract.WithdrawPastDueReqest(&_Supersymmetry.TransactOpts, requestHash)
}

// WithdrawPastDueReqest is a paid mutator transaction binding the contract method 0x329881c6.
//
// Solidity: function withdrawPastDueReqest(bytes32 requestHash) returns()
func (_Supersymmetry *SupersymmetryTransactorSession) WithdrawPastDueReqest(requestHash [32]byte) (*types.Transaction, error) {
	return _Supersymmetry.Contract.WithdrawPastDueReqest(&_Supersymmetry.TransactOpts, requestHash)
}

// SupersymmetryMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Supersymmetry contract.
type SupersymmetryMintIterator struct {
	Event *SupersymmetryMint // Event containing the contract specifics and raw log

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
func (it *SupersymmetryMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupersymmetryMint)
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
		it.Event = new(SupersymmetryMint)
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
func (it *SupersymmetryMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupersymmetryMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupersymmetryMint represents a Mint event raised by the Supersymmetry contract.
type SupersymmetryMint struct {
	RequestHash [32]byte
	Owner       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xf0bcd414edc501835c6c541058514fe45a43fa62a41671c9f62317c9889fcf7f.
//
// Solidity: event Mint(bytes32 requestHash, address owner, uint256 amount)
func (_Supersymmetry *SupersymmetryFilterer) FilterMint(opts *bind.FilterOpts) (*SupersymmetryMintIterator, error) {

	logs, sub, err := _Supersymmetry.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &SupersymmetryMintIterator{contract: _Supersymmetry.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xf0bcd414edc501835c6c541058514fe45a43fa62a41671c9f62317c9889fcf7f.
//
// Solidity: event Mint(bytes32 requestHash, address owner, uint256 amount)
func (_Supersymmetry *SupersymmetryFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SupersymmetryMint) (event.Subscription, error) {

	logs, sub, err := _Supersymmetry.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupersymmetryMint)
				if err := _Supersymmetry.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xf0bcd414edc501835c6c541058514fe45a43fa62a41671c9f62317c9889fcf7f.
//
// Solidity: event Mint(bytes32 requestHash, address owner, uint256 amount)
func (_Supersymmetry *SupersymmetryFilterer) ParseMint(log types.Log) (*SupersymmetryMint, error) {
	event := new(SupersymmetryMint)
	if err := _Supersymmetry.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SupersymmetryNewRequestIterator is returned from FilterNewRequest and is used to iterate over the raw logs and unpacked data for NewRequest events raised by the Supersymmetry contract.
type SupersymmetryNewRequestIterator struct {
	Event *SupersymmetryNewRequest // Event containing the contract specifics and raw log

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
func (it *SupersymmetryNewRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupersymmetryNewRequest)
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
		it.Event = new(SupersymmetryNewRequest)
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
func (it *SupersymmetryNewRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupersymmetryNewRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupersymmetryNewRequest represents a NewRequest event raised by the Supersymmetry contract.
type SupersymmetryNewRequest struct {
	RequestHash [32]byte
	Owner       common.Address
	Target      string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterNewRequest is a free log retrieval operation binding the contract event 0xe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c115896.
//
// Solidity: event NewRequest(bytes32 requestHash, address owner, string target)
func (_Supersymmetry *SupersymmetryFilterer) FilterNewRequest(opts *bind.FilterOpts) (*SupersymmetryNewRequestIterator, error) {

	logs, sub, err := _Supersymmetry.contract.FilterLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return &SupersymmetryNewRequestIterator{contract: _Supersymmetry.contract, event: "NewRequest", logs: logs, sub: sub}, nil
}

// WatchNewRequest is a free log subscription operation binding the contract event 0xe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c115896.
//
// Solidity: event NewRequest(bytes32 requestHash, address owner, string target)
func (_Supersymmetry *SupersymmetryFilterer) WatchNewRequest(opts *bind.WatchOpts, sink chan<- *SupersymmetryNewRequest) (event.Subscription, error) {

	logs, sub, err := _Supersymmetry.contract.WatchLogs(opts, "NewRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupersymmetryNewRequest)
				if err := _Supersymmetry.contract.UnpackLog(event, "NewRequest", log); err != nil {
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

// ParseNewRequest is a log parse operation binding the contract event 0xe5048cb72529de5404e374d7fdbda8558caeb479015b507903292cf51c115896.
//
// Solidity: event NewRequest(bytes32 requestHash, address owner, string target)
func (_Supersymmetry *SupersymmetryFilterer) ParseNewRequest(log types.Log) (*SupersymmetryNewRequest, error) {
	event := new(SupersymmetryNewRequest)
	if err := _Supersymmetry.contract.UnpackLog(event, "NewRequest", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SupersymmetryReturnIterator is returned from FilterReturn and is used to iterate over the raw logs and unpacked data for Return events raised by the Supersymmetry contract.
type SupersymmetryReturnIterator struct {
	Event *SupersymmetryReturn // Event containing the contract specifics and raw log

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
func (it *SupersymmetryReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupersymmetryReturn)
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
		it.Event = new(SupersymmetryReturn)
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
func (it *SupersymmetryReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupersymmetryReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupersymmetryReturn represents a Return event raised by the Supersymmetry contract.
type SupersymmetryReturn struct {
	RequestHash [32]byte
	Owner       common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReturn is a free log retrieval operation binding the contract event 0x86bd878938275851d348ab8cd7e09276a8abb4b70910d2e9e913a38815997313.
//
// Solidity: event Return(bytes32 requestHash, address owner, uint256 amount)
func (_Supersymmetry *SupersymmetryFilterer) FilterReturn(opts *bind.FilterOpts) (*SupersymmetryReturnIterator, error) {

	logs, sub, err := _Supersymmetry.contract.FilterLogs(opts, "Return")
	if err != nil {
		return nil, err
	}
	return &SupersymmetryReturnIterator{contract: _Supersymmetry.contract, event: "Return", logs: logs, sub: sub}, nil
}

// WatchReturn is a free log subscription operation binding the contract event 0x86bd878938275851d348ab8cd7e09276a8abb4b70910d2e9e913a38815997313.
//
// Solidity: event Return(bytes32 requestHash, address owner, uint256 amount)
func (_Supersymmetry *SupersymmetryFilterer) WatchReturn(opts *bind.WatchOpts, sink chan<- *SupersymmetryReturn) (event.Subscription, error) {

	logs, sub, err := _Supersymmetry.contract.WatchLogs(opts, "Return")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupersymmetryReturn)
				if err := _Supersymmetry.contract.UnpackLog(event, "Return", log); err != nil {
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

// ParseReturn is a log parse operation binding the contract event 0x86bd878938275851d348ab8cd7e09276a8abb4b70910d2e9e913a38815997313.
//
// Solidity: event Return(bytes32 requestHash, address owner, uint256 amount)
func (_Supersymmetry *SupersymmetryFilterer) ParseReturn(log types.Log) (*SupersymmetryReturn, error) {
	event := new(SupersymmetryReturn)
	if err := _Supersymmetry.contract.UnpackLog(event, "Return", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SupersymmetryStatusChangedIterator is returned from FilterStatusChanged and is used to iterate over the raw logs and unpacked data for StatusChanged events raised by the Supersymmetry contract.
type SupersymmetryStatusChangedIterator struct {
	Event *SupersymmetryStatusChanged // Event containing the contract specifics and raw log

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
func (it *SupersymmetryStatusChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SupersymmetryStatusChanged)
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
		it.Event = new(SupersymmetryStatusChanged)
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
func (it *SupersymmetryStatusChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SupersymmetryStatusChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SupersymmetryStatusChanged represents a StatusChanged event raised by the Supersymmetry contract.
type SupersymmetryStatusChanged struct {
	RequestHash [32]byte
	Status      uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterStatusChanged is a free log retrieval operation binding the contract event 0x22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c.
//
// Solidity: event StatusChanged(bytes32 requestHash, uint8 status)
func (_Supersymmetry *SupersymmetryFilterer) FilterStatusChanged(opts *bind.FilterOpts) (*SupersymmetryStatusChangedIterator, error) {

	logs, sub, err := _Supersymmetry.contract.FilterLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return &SupersymmetryStatusChangedIterator{contract: _Supersymmetry.contract, event: "StatusChanged", logs: logs, sub: sub}, nil
}

// WatchStatusChanged is a free log subscription operation binding the contract event 0x22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c.
//
// Solidity: event StatusChanged(bytes32 requestHash, uint8 status)
func (_Supersymmetry *SupersymmetryFilterer) WatchStatusChanged(opts *bind.WatchOpts, sink chan<- *SupersymmetryStatusChanged) (event.Subscription, error) {

	logs, sub, err := _Supersymmetry.contract.WatchLogs(opts, "StatusChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SupersymmetryStatusChanged)
				if err := _Supersymmetry.contract.UnpackLog(event, "StatusChanged", log); err != nil {
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

// ParseStatusChanged is a log parse operation binding the contract event 0x22ce671cd3f328f71ef109d846af764df2d16bb8eb664c4f41065656c6e9962c.
//
// Solidity: event StatusChanged(bytes32 requestHash, uint8 status)
func (_Supersymmetry *SupersymmetryFilterer) ParseStatusChanged(log types.Log) (*SupersymmetryStatusChanged, error) {
	event := new(SupersymmetryStatusChanged)
	if err := _Supersymmetry.contract.UnpackLog(event, "StatusChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MinterRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isMinter\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceMinter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenFuncSigs maps the 4-byte function signature to its string representation.
var TokenFuncSigs = map[string]string{
	"983b2d56": "addMinter(address)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"aa271e1a": "isMinter(address)",
	"40c10f19": "mint(address,uint256)",
	"06fdde03": "name()",
	"98650275": "renounceMinter()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// TokenBin is the compiled bytecode used for deploying new contracts.
var TokenBin = "0x60806040523480156200001157600080fd5b506040516200132038038062001320833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82516401000000008111828201881017156200008957600080fd5b82525081516020918201929091019080838360005b83811015620000b85781810151838201526020016200009e565b50505050905090810190601f168015620000e65780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200010a57600080fd5b9083019060208201858111156200012057600080fd5b82516401000000008111828201881017156200013b57600080fd5b82525081516020918201929091019080838360005b838110156200016a57818101518382015260200162000150565b50505050905090810190601f168015620001985780820380516001836020036101000a031916815260200191505b50604052602001519150620001ca9050620001bb6001600160e01b036200021c16565b6001600160e01b036200022116565b8251620001df90600490602086019062000369565b508151620001f590600590602085019062000369565b506006805460ff191660ff831617905562000213620001bb6200021c565b5050506200040b565b335b90565b6200023c8160036200027360201b62000c521790919060201c565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b6200028882826001600160e01b036200030016565b15620002db576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b60006001600160a01b038216620003495760405162461bcd60e51b8152600401808060200182810382526022815260200180620012fe6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003ac57805160ff1916838001178555620003dc565b82800160010185558215620003dc579182015b82811115620003dc578251825591602001919060010190620003bf565b50620003ea929150620003ee565b5090565b6200021e91905b80821115620003ea5760008155600101620003f5565b610ee3806200041b6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d7146102db578063a9059cbb14610307578063aa271e1a14610333578063dd62ed3e14610359576100f5565b806370a082311461027d57806395d89b41146102a3578063983b2d56146102ab57806398650275146102d3576100f5565b806323b872dd116100d357806323b872dd146101d1578063313ce56714610207578063395093511461022557806340c10f1914610251576100f5565b806306fdde03146100fa578063095ea7b31461017757806318160ddd146101b7575b600080fd5b610102610387565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013c578181015183820152602001610124565b50505050905090810190601f1680156101695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a36004803603604081101561018d57600080fd5b506001600160a01b03813516906020013561041d565b604080519115158252519081900360200190f35b6101bf61043a565b60408051918252519081900360200190f35b6101a3600480360360608110156101e757600080fd5b506001600160a01b03813581169160208101359091169060400135610440565b61020f6104cd565b6040805160ff9092168252519081900360200190f35b6101a36004803603604081101561023b57600080fd5b506001600160a01b0381351690602001356104d6565b6101a36004803603604081101561026757600080fd5b506001600160a01b03813516906020013561052a565b6101bf6004803603602081101561029357600080fd5b50356001600160a01b0316610581565b61010261059c565b6102d1600480360360208110156102c157600080fd5b50356001600160a01b03166105fd565b005b6102d161064f565b6101a3600480360360408110156102f157600080fd5b506001600160a01b038135169060200135610661565b6101a36004803603604081101561031d57600080fd5b506001600160a01b0381351690602001356106cf565b6101a36004803603602081101561034957600080fd5b50356001600160a01b03166106e3565b6101bf6004803603604081101561036f57600080fd5b506001600160a01b03813581169160200135166106fc565b60048054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104135780601f106103e857610100808354040283529160200191610413565b820191906000526020600020905b8154815290600101906020018083116103f657829003601f168201915b5050505050905090565b600061043161042a610727565b848461072b565b50600192915050565b60025490565b600061044d848484610817565b6104c384610459610727565b6104be85604051806060016040528060288152602001610df7602891396001600160a01b038a16600090815260016020526040812090610497610727565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61097316565b61072b565b5060019392505050565b60065460ff1690565b60006104316104e3610727565b846104be85600160006104f4610727565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff610a0a16565b600061053c610537610727565b6106e3565b6105775760405162461bcd60e51b8152600401808060200182810382526030815260200180610da66030913960400191505060405180910390fd5b6104318383610a6b565b6001600160a01b031660009081526020819052604090205490565b60058054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104135780601f106103e857610100808354040283529160200191610413565b610608610537610727565b6106435760405162461bcd60e51b8152600401808060200182810382526030815260200180610da66030913960400191505060405180910390fd5b61064c81610b5b565b50565b61065f61065a610727565b610ba3565b565b600061043161066e610727565b846104be85604051806060016040528060258152602001610e8a6025913960016000610698610727565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61097316565b60006104316106dc610727565b8484610817565b60006106f660038363ffffffff610beb16565b92915050565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166107705760405162461bcd60e51b8152600401808060200182810382526024815260200180610e666024913960400191505060405180910390fd5b6001600160a01b0382166107b55760405162461bcd60e51b8152600401808060200182810382526022815260200180610d5e6022913960400191505060405180910390fd5b6001600160a01b03808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b6001600160a01b03831661085c5760405162461bcd60e51b8152600401808060200182810382526025815260200180610e416025913960400191505060405180910390fd5b6001600160a01b0382166108a15760405162461bcd60e51b8152600401808060200182810382526023815260200180610d3b6023913960400191505060405180910390fd5b6108e481604051806060016040528060268152602001610d80602691396001600160a01b038616600090815260208190526040902054919063ffffffff61097316565b6001600160a01b038085166000908152602081905260408082209390935590841681522054610919908263ffffffff610a0a16565b6001600160a01b038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b60008184841115610a025760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156109c75781810151838201526020016109af565b50505050905090810190601f1680156109f45780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b600082820183811015610a64576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b6001600160a01b038216610ac6576040805162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015290519081900360640190fd5b600254610ad9908263ffffffff610a0a16565b6002556001600160a01b038216600090815260208190526040902054610b05908263ffffffff610a0a16565b6001600160a01b0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b610b6c60038263ffffffff610c5216565b6040516001600160a01b038216907f6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f690600090a250565b610bb460038263ffffffff610cd316565b6040516001600160a01b038216907fe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb6669290600090a250565b60006001600160a01b038216610c325760405162461bcd60e51b8152600401808060200182810382526022815260200180610e1f6022913960400191505060405180910390fd5b506001600160a01b03166000908152602091909152604090205460ff1690565b610c5c8282610beb565b15610cae576040805162461bcd60e51b815260206004820152601f60248201527f526f6c65733a206163636f756e7420616c72656164792068617320726f6c6500604482015290519081900360640190fd5b6001600160a01b0316600090815260209190915260409020805460ff19166001179055565b610cdd8282610beb565b610d185760405162461bcd60e51b8152600401808060200182810382526021815260200180610dd66021913960400191505060405180910390fd5b6001600160a01b0316600090815260209190915260409020805460ff1916905556fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e63654d696e746572526f6c653a2063616c6c657220646f6573206e6f74206861766520746865204d696e74657220726f6c65526f6c65733a206163636f756e7420646f6573206e6f74206861766520726f6c6545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e6365526f6c65733a206163636f756e7420697320746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa265627a7a72315820a0fd174d9fb4a843cfda8d411e74be8e3d289af97e1c49347b4785936176637264736f6c63430005100032526f6c65733a206163636f756e7420697320746865207a65726f2061646472657373"

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Token *TokenCallerSession) Decimals() (uint8, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_Token *TokenCaller) IsMinter(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "isMinter", account)
	return *ret0, err
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_Token *TokenSession) IsMinter(account common.Address) (bool, error) {
	return _Token.Contract.IsMinter(&_Token.CallOpts, account)
}

// IsMinter is a free data retrieval call binding the contract method 0xaa271e1a.
//
// Solidity: function isMinter(address account) constant returns(bool)
func (_Token *TokenCallerSession) IsMinter(account common.Address) (bool, error) {
	return _Token.Contract.IsMinter(&_Token.CallOpts, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Token *TokenCallerSession) Symbol() (string, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_Token *TokenTransactor) AddMinter(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "addMinter", account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_Token *TokenSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddMinter(&_Token.TransactOpts, account)
}

// AddMinter is a paid mutator transaction binding the contract method 0x983b2d56.
//
// Solidity: function addMinter(address account) returns()
func (_Token *TokenTransactorSession) AddMinter(account common.Address) (*types.Transaction, error) {
	return _Token.Contract.AddMinter(&_Token.TransactOpts, account)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Token *TokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.DecreaseAllowance(&_Token.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Token *TokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Token.Contract.IncreaseAllowance(&_Token.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Token *TokenSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, account, amount)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_Token *TokenTransactor) RenounceMinter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "renounceMinter")
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_Token *TokenSession) RenounceMinter() (*types.Transaction, error) {
	return _Token.Contract.RenounceMinter(&_Token.TransactOpts)
}

// RenounceMinter is a paid mutator transaction binding the contract method 0x98650275.
//
// Solidity: function renounceMinter() returns()
func (_Token *TokenTransactorSession) RenounceMinter() (*types.Transaction, error) {
	return _Token.Contract.RenounceMinter(&_Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, sender, recipient, amount)
}

// TokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Token contract.
type TokenApprovalIterator struct {
	Event *TokenApproval // Event containing the contract specifics and raw log

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
func (it *TokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenApproval)
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
		it.Event = new(TokenApproval)
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
func (it *TokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenApproval represents a Approval event raised by the Token contract.
type TokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TokenApprovalIterator{contract: _Token.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenApproval)
				if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Token *TokenFilterer) ParseApproval(log types.Log) (*TokenApproval, error) {
	event := new(TokenApproval)
	if err := _Token.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMinterAddedIterator is returned from FilterMinterAdded and is used to iterate over the raw logs and unpacked data for MinterAdded events raised by the Token contract.
type TokenMinterAddedIterator struct {
	Event *TokenMinterAdded // Event containing the contract specifics and raw log

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
func (it *TokenMinterAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinterAdded)
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
		it.Event = new(TokenMinterAdded)
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
func (it *TokenMinterAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinterAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinterAdded represents a MinterAdded event raised by the Token contract.
type TokenMinterAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterAdded is a free log retrieval operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_Token *TokenFilterer) FilterMinterAdded(opts *bind.FilterOpts, account []common.Address) (*TokenMinterAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenMinterAddedIterator{contract: _Token.contract, event: "MinterAdded", logs: logs, sub: sub}, nil
}

// WatchMinterAdded is a free log subscription operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_Token *TokenFilterer) WatchMinterAdded(opts *bind.WatchOpts, sink chan<- *TokenMinterAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinterAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinterAdded)
				if err := _Token.contract.UnpackLog(event, "MinterAdded", log); err != nil {
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

// ParseMinterAdded is a log parse operation binding the contract event 0x6ae172837ea30b801fbfcdd4108aa1d5bf8ff775444fd70256b44e6bf3dfc3f6.
//
// Solidity: event MinterAdded(address indexed account)
func (_Token *TokenFilterer) ParseMinterAdded(log types.Log) (*TokenMinterAdded, error) {
	event := new(TokenMinterAdded)
	if err := _Token.contract.UnpackLog(event, "MinterAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenMinterRemovedIterator is returned from FilterMinterRemoved and is used to iterate over the raw logs and unpacked data for MinterRemoved events raised by the Token contract.
type TokenMinterRemovedIterator struct {
	Event *TokenMinterRemoved // Event containing the contract specifics and raw log

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
func (it *TokenMinterRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenMinterRemoved)
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
		it.Event = new(TokenMinterRemoved)
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
func (it *TokenMinterRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenMinterRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenMinterRemoved represents a MinterRemoved event raised by the Token contract.
type TokenMinterRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinterRemoved is a free log retrieval operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_Token *TokenFilterer) FilterMinterRemoved(opts *bind.FilterOpts, account []common.Address) (*TokenMinterRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &TokenMinterRemovedIterator{contract: _Token.contract, event: "MinterRemoved", logs: logs, sub: sub}, nil
}

// WatchMinterRemoved is a free log subscription operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_Token *TokenFilterer) WatchMinterRemoved(opts *bind.WatchOpts, sink chan<- *TokenMinterRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "MinterRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenMinterRemoved)
				if err := _Token.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
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

// ParseMinterRemoved is a log parse operation binding the contract event 0xe94479a9f7e1952cc78f2d6baab678adc1b772d936c6583def489e524cb66692.
//
// Solidity: event MinterRemoved(address indexed account)
func (_Token *TokenFilterer) ParseMinterRemoved(log types.Log) (*TokenMinterRemoved, error) {
	event := new(TokenMinterRemoved)
	if err := _Token.contract.UnpackLog(event, "MinterRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Token *TokenFilterer) ParseTransfer(log types.Log) (*TokenTransfer, error) {
	event := new(TokenTransfer)
	if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
