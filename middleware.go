package main

import (
    "github.com/kataras/iris"
    )

func main() {
    app := iris.New()
    // register the "before" handler as the first handler which will be executed
    // on all domain's routes.
    // or use the `UseGlobal` to register a middleware which will fire across subdomains.
    app.Use(before)
    // register the "after" handler as the last handler which will be executed
    // after all domain's routes' handler(s).
    app.Done(after)

    // register our routes.
    app.Get("/", indexHandler)
    app.Get("/contact", contactHandler)

    app.Run(iris.Addr(":8080"))
}

func before(ctx iris.Context) {
    header := ctx.Request().Header["Authorization"]
    if header != nil {
        authorization_header := header[0] 
        str := "<h1>Before Action</h1><br>Authorization header : " + authorization_header
        ctx.HTML(str)
        ctx.Next()
    } else{
        ctx.HTML("Unauthorize")
    }
}

func after(ctx iris.Context) {
    ctx.HTML("<h1>After Action</h1>")
}

func indexHandler(ctx iris.Context) {
    // write something to the client as a response.
    ctx.HTML("<h1>Index</h1>")

    ctx.Next() // execute the "after" handler registered via `Done`.
}

func contactHandler(ctx iris.Context) {
    // write something to the client as a response.
    ctx.HTML("<h1>Contact</h1>")

    ctx.Next() // execute the "after" handler registered via `Done`.
}
