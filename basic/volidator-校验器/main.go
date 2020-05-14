package main

import (
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type User struct {
	FirstName string `validate: "required"`
	LastName  string `validate: "required"`
	Age       uint64 `validate: “gte=0, lte=130”` // uint64 是正整数
	Email     string `validate: "requires, email"`
}

func main() {
	user := &User{
		FirstName: "firstname",
		LastName:  "lastname",
		Age:       136,
		Email:     "miniok163.com",
	}
	validate := validator.New()
	//user = nil go 语言验证自定义验证错误的方法 就是将要验证的数据设置为nil
	err := validate.Struct(user)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
		}
		if errs, ok := err.(*validator.ValidationErrors); ok {
			fmt.Println(errs)
		}

	}

}
