package main
import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  )
func main() {
  // ENCODE
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "user": "passwords",
  })
  tokenString, _ := token.SignedString([]byte("My Secret"))
  fmt.Println(tokenString)

  // DECODE
  claims := jwt.MapClaims{}
  token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
      return []byte("My Secret"), nil
  })
  for key, val := range claims {
    fmt.Println("Key: %v, value: %v\n", key, val)
  }
  fmt.Println(err)
}
