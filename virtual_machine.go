package js

import (
	"errors"
	"github.com/robertkrimen/otto"
	_ "github.com/robertkrimen/otto/underscore"
)

type virtualMachine struct {
	vm     *otto.Otto
	script *otto.Script
}

func NewVirtualMachine() *virtualMachine {
	exec := &virtualMachine{
		vm: otto.New(),
	}
	exec.init()
	return exec
}

func (e *virtualMachine) init() {
}

// Destroy 销毁虚拟机，为了性能考虑，现在只是将之还给虚拟机池。
func (e *virtualMachine) Destroy() {
}

// LoadString 加载字符串，并编译成字节码
func (e *virtualMachine) LoadString(source string) error {
	script, err := e.vm.Compile("", source)
	if err != nil {
		return err
	}
	e.script = script
	return nil
}

// LoadFile 加载文件，并编译成字节码
func (e *virtualMachine) LoadFile(filePath string) error {
	script, err := e.vm.Compile(filePath, nil)
	if err != nil {
		return err
	}
	e.script = script
	return nil
}

// Execute 执行已编译的lua代码
func (e *virtualMachine) Execute() error {
	if e.script == nil {
		return errors.New("no js")
	}
	_, err := e.vm.Run(e.script)
	return err
}

// ExecuteString 直接执行字符串
func (e *virtualMachine) ExecuteString(source string) error {
	_, err := e.vm.Run(source)
	return err
}

// ExecuteFile 直接执行lua文件
func (e *virtualMachine) ExecuteFile(filePath string) error {
	return e.compileAndRun(filePath)
}

func (e *virtualMachine) compileAndRun(filePath string) error {
	script, err := e.vm.Compile(filePath, nil)
	if err != nil {
		return err
	}
	_, err = e.vm.Run(script)
	if err != nil {
		return err
	}
	return nil
}

// Register 注册方法或者变量到js
func (e *virtualMachine) Register(name string, value interface{}) error {
	err := e.vm.Set(name, value)
	return err
}

// CallFunction 调用js中的方法
func (e *virtualMachine) CallFunction(name string, args ...interface{}) error {
	_, err := e.vm.Call(name, nil, args...)
	return err
}
