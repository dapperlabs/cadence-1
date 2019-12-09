package interpreter

import (
	"github.com/dapperlabs/flow-go/language/runtime/ast"
	"github.com/dapperlabs/flow-go/language/runtime/errors"
	"github.com/dapperlabs/flow-go/language/runtime/sema"
	"github.com/raviqqe/hamt"
	// revive:disable:dot-imports
	. "github.com/dapperlabs/flow-go/language/runtime/trampoline"
	// revive:enable
)

// Invocation

type Invocation struct {
	Arguments     []Value
	ArgumentTypes []sema.Type
	Location      LocationPosition
	Interpreter   *Interpreter
}

// FunctionValue

type FunctionValue interface {
	Value
	isFunctionValue()
	invoke(Invocation) Trampoline
}

// InterpretedFunctionValue

type InterpretedFunctionValue struct {
	Interpreter *Interpreter
	Expression  *ast.FunctionExpression
	Type        *sema.FunctionType
	Activation  hamt.Map
}

func (InterpretedFunctionValue) isValue() {}

func (f InterpretedFunctionValue) Copy() Value {
	return f
}

func (InterpretedFunctionValue) GetOwner() string {
	// value is never owned
	return ""
}

func (InterpretedFunctionValue) SetOwner(owner string) {
	// NO-OP: value cannot be owned
}

func (InterpretedFunctionValue) isFunctionValue() {}

func newInterpretedFunction(
	interpreter *Interpreter,
	expression *ast.FunctionExpression,
	functionType *sema.FunctionType,
	activation hamt.Map,
) InterpretedFunctionValue {
	return InterpretedFunctionValue{
		Interpreter: interpreter,
		Expression:  expression,
		Type:        functionType,
		Activation:  activation,
	}
}

func (f InterpretedFunctionValue) invoke(invocation Invocation) Trampoline {
	return f.Interpreter.invokeInterpretedFunction(f, invocation.Arguments)
}

// HostFunctionValue

type HostFunction func(invocation Invocation) Trampoline

type HostFunctionValue struct {
	Function HostFunction
	Members  map[string]Value
}

func NewHostFunctionValue(
	function HostFunction,
) HostFunctionValue {
	return HostFunctionValue{
		Function: function,
	}
}

func (HostFunctionValue) isValue() {}

func (f HostFunctionValue) Copy() Value {
	return f
}

func (HostFunctionValue) GetOwner() string {
	// value is never owned
	return ""
}

func (HostFunctionValue) SetOwner(owner string) {
	// NO-OP: value cannot be owned
}

func (HostFunctionValue) isFunctionValue() {}

func (f HostFunctionValue) invoke(invocation Invocation) Trampoline {
	return f.Function(invocation)
}

func (f HostFunctionValue) GetMember(interpreter *Interpreter, _ LocationRange, name string) Value {
	return f.Members[name]
}

func (f HostFunctionValue) SetMember(_ *Interpreter, _ LocationRange, _ string, _ Value) {
	panic(errors.NewUnreachableError())
}
