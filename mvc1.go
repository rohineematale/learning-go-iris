package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
)

func main() {
    app := iris.New()

    app.Controller("/helloworld", new(HelloWorldController))

    app.Run(iris.Addr("localhost:8080"))
}

type HelloWorldController struct {
    mvc.Controller

    // [ Your fields here ]
    // Request lifecycle data
    // Models
    // Database
    // Global properties
}

//
// GET: /helloworld

func (c *HelloWorldController) Get() string {
    return "This is my default action..."
}

//
// GET: /helloworld/{name:string}

func (c *HelloWorldController) GetBy(name string) string {
    return "Hello " + name
}

//
// GET: /helloworld/welcome

func (c *HelloWorldController) GetWelcome() (string, int) {
    return "This is the GetWelcome action func...", iris.StatusOK
}

//
// GET: /helloworld/welcome/{name:string}/{numTimes:int}

func (c *HelloWorldController) GetWelcomeBy(name string, numTimes int) {
    // Access to the low-level Context,
    // output arguments are optional of course so we don't have to use them here.
    c.Ctx.Writef("Hello %s, NumTimes is: %d", name, numTimes)
}

/*
func (c *HelloWorldController) Post() {} handles HTTP POST method requests
func (c *HelloWorldController) Put() {} handles HTTP PUT method requests
func (c *HelloWorldController) Delete() {} handles HTTP DELETE method requests
func (c *HelloWorldController) Connect() {} handles HTTP CONNECT method requests
func (c *HelloWorldController) Head() {} handles HTTP HEAD method requests
func (c *HelloWorldController) Patch() {} handles HTTP PATCH method requests
func (c *HelloWorldController) Options() {} handles HTTP OPTIONS method requests
func (c *HelloWorldController) Trace() {} handles HTTP TRACE method requests
*/

/*
func (c *HelloWorldController) All() {} handles All method requests
//        OR
func (c *HelloWorldController) Any() {} handles All method requests
*/