package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Test_JWT() {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("recite-words"))

	fmt.Println(tokenString, err)
}

func Test_JWT2() {
	fmt.Println("this is JWT2.")
	// nowTime := time.Now()
	// expireTime := nowTime.Add(7100)

	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := newJwt.SignedString([]byte("recite-words"))

	fmt.Println(tokenString, err)
}
