package js

import (
	"fmt"
	"github.com/robertkrimen/otto"
	"github.com/stretchr/testify/assert"
	"testing"
)

type user struct {
	Name  string
	token string
}

func (u *user) SetToken(t string) {
	u.token = t
}

func (u *user) Token() string {
	return u.token
}

func TestVirtualMachine_ExecuteString(t *testing.T) {
	exe := NewVirtualMachine()
	defer exe.Destroy()

	jsString := `console.log('ddddd')`

	err := exe.LoadString(jsString)
	assert.Nil(t, err)

	type Role struct {
		Name string
	}
	type Menu struct {
		Name string
	}

	var role Role
	var menu Menu

	err = exe.Execute()
	assert.Nil(t, err)

	fmt.Println(role)
	fmt.Println(menu)
}

func TestVirtualMachine_ExecuteFile(t *testing.T) {
	exe := NewVirtualMachine()
	defer exe.Destroy()

	err := exe.LoadFile("./script/test.js")
	assert.Nil(t, err)

	u := &user{
		Name: "Tim",
	}

	err = exe.Register("u", u)
	assert.Nil(t, err)

	err = exe.Register("sayHello", func(call otto.FunctionCall) otto.Value {
		fmt.Printf("Hello, %s.\n", call.Argument(0).String())
		return otto.Value{}
	})
	assert.Nil(t, err)

	err = exe.Execute()
	assert.Nil(t, err)

	fmt.Println("Js set your token to:", u.Token())

	err = exe.CallFunction("printMessage", "hello")
	assert.Nil(t, err)
}
