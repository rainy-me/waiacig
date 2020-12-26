package vm

import (
	"waiacig/code"
	"waiacig/object"
)

type Frame struct {
	fn          *object.CompiledFunction
	ip          int
	basePointer int
}

func NewFrame(fn *object.CompiledFunction, basePointer int) *Frame {
	f := &Frame{
		fn:          fn,
		ip:          -1,
		basePointer: basePointer,
	}
	return f
}

func (f *Frame) Instructions() code.Instructions {
	return f.fn.Instructions
}

func (vm *VM) currentFrame() *Frame {
	return vm.frames[vm.framesIndex-1]
}
func (vm *VM) pushFrame(f *Frame) {
	vm.frames[vm.framesIndex] = f
	vm.framesIndex++
}
func (vm *VM) popFrame() *Frame {
	vm.framesIndex--
	return vm.frames[vm.framesIndex]
}
