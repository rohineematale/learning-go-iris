package main

import (
  "github.com/kataras/iris"
  "github.com/dgrijalva/jwt-go"
  jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func main() {
  app := iris.New()

  jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
      return []byte("My Secret"), nil
    },
    SigningMethod: jwt.SigningMethodHS256,
  })

  // Header => Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmb28iOiJiYXIifQ.4aQan-XvJBjDUDbCVnuh2P_xy54b2aRKKsgKcHUa8uw

  app.Use(jwtHandler.Serve)

  // middelware
  // register the "before" handler as the first handler which will be executed on all domain's routes.
  // or use the `UseGlobal` to register a middleware which will fire across subdomains.
  app.Use(before)
  // register the "after" handler as the last handler which will be executed
  // after all domain's routes' handler(s).
  app.Done(after)
  // end middleware. Functions are written after main block

  // You got full debug messages, useful when using MVC and you want to make
  // sure that your code is aligned with the Iris' MVC Architecture.
  app.Logger().SetLevel("debug")

  //A. namespace
  app.PartyFunc("/users", func(users iris.Party) {
    // http://localhost:8080/users/42/profile
    users.Get("/{id:int}/profile", func(ctx iris.Context) {
      ctx.HTML("<b>Its from user profile</b>")
    })
    // http://localhost:8080/users/messages/1
    users.Get("/inbox/{id:int}", func(ctx iris.Context) {
      println("Inside action")
      ctx.HTML("<b>Its from user </b>")
      ctx.Next()
    })
  })
  // end of namespace

  //B. types of dynamic urls
  app.Get("/lowercase/{name:string regexp(^[a-z]+)}", func(ctx iris.Context) {
    println("name should be only lowercase, otherwise this handler will never executed: %s", ctx.Params().Get("name"))
    ctx.Next()
  })

  app.Get("/profile/{id:int max(10)}", func(ctx iris.Context) {
    // second parameter is the error but it will always nil because we use macros,
    // the validaton already happened.
    id, _ := ctx.Params().GetInt("id")
    println("Hello id: %d", id)
  })
  // end of types of dynamic urls


  //C. named routes
  h := func(ctx iris.Context) {
    ctx.HTML("named routes")
  }
  // handler registration and naming
  home := app.Get("/home", h)
  home.Name = "home"
  // or
  app.Get("/about", h).Name = "about"
  app.Get("/page/{id}", h).Name = "page"
  // end of named routes


  app.Get("/", func(ctx iris.Context) {
    ctx.HTML("<b>Yahh! Its working. Now convert it to MVC.</b>")
    ctx.Next()
  })

  app.Run(iris.Addr(":8080"), iris.WithConfiguration(iris.YAML("./configs/iris.yml")))
}

// before filter method
func before(ctx iris.Context) {
  // shareInformation := "this is a sharable information between handlers"

  requestPath := ctx.Path()
  ctx.HTML("Before the action: " + requestPath)

  // ctx.Values().Set("info", shareInformation)
  ctx.Next() // execute the next handler, in this case the main one.
}
// end before filter method

// after filter method
func after(ctx iris.Context) {
  ctx.HTML("After the action")
}
// end after filter method

/**
NOTE : 
1. Add debugger : brew install go-delve/delve/delve 
2. run : dlv debug file_name
3. add breakpoint using b main.main or b main.<function_name>
4. Hit url, you will see breakpoint in terminal.
Ref: https://github.com/derekparker/delve/issues/976#issuecomment-363181012
**/
