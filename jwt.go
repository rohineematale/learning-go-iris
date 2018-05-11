// iris provides some basic middleware, most for your learning courve.
// You can use any net/http compatible middleware with iris.FromStd wrapper.
//
// JWT net/http video tutorial for golang newcomers: https://www.youtube.com/watch?v=dgJFeqeXVKw
//
// This middleware is the only one cloned from external source: https://github.com/auth0/go-jwt-middleware
// (because it used "context" to define the user but we don't need that so a simple iris.FromStd wouldn't work as expected.)

// $ go get -u github.com/dgrijalva/jwt-go
// go get -u github.com/iris-contrib/middleware/jwt If you have not set proper GOPATH
// $ go run main.go

package main

import (
  "github.com/kataras/iris"
  "github.com/dgrijalva/jwt-go"
  jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

func myHandler(ctx iris.Context) {
  user := ctx.Values().Get("jwt").(*jwt.Token)

  ctx.Writef("This is an authenticated request\n")
  ctx.Writef("Claim content:\n")

  ctx.Writef("%s", user.Signature)
}

func main() {
  app := iris.New()

  jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
    ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
      return []byte("My Secret"), nil
    },
    SigningMethod: jwt.SigningMethodHS256,
  })

  app.Use(jwtHandler.Serve)

  app.Get("/ping", myHandler)
  app.Get("/", func(ctx iris.Context) {
    ctx.HTML("<b>Yahh! Its working. Now convert it to MVC.</b>")
    ctx.Next()
  })
  app.Run(iris.Addr("localhost:8080"))
} // don't forget to look ../jwt_test.go to seee how to set your own custom claims