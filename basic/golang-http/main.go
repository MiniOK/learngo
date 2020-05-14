package main

import (
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

func main() {

	app := iris.Default()
	app.Get("/hello", func(context iris.Context) {
		context.WriteString("Hello iris!!!")
	})

	app.Get("/user/{id:uint64}",
		func(ctx iris.Context) {
			id := ctx.Params().GetUint64Default("id", 0)
			ctx.WriteString(strconv.Itoa(int(id)))

		})
	err := app.Run(iris.Addr(":8082"))
	fmt.Println(err)
}
