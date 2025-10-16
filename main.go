package main

import (
	"ecommere.com/cmd"
)

// "crypto/sha256"
// "fmt"
// "fmt"

// "encoding/base64"

// import

func main() {
	// fmt.Println("server start")

	cmd.Serve()

	// var s string

	// s = "hello world"

	// byteArr := []byte(s)

	// fmt.Println(s)

	// fmt.Println(byteArr)

	// enc := base64.URLEncoding
	// enc = enc.WithPadding(base64.NoPadding)
	// b64Str := enc.EncodeToString(byteArr)

	// fmt.Println(b64Str)
	// decodedStr, err := enc.DecodeString(b64Str)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(decodedStr)

	// base64.URLEncoding.WithPadding(base64.NoPadding)

	// data := []byte("Hello Tanjil")
	// hash := sha256.Sum256(data)
	// fmt.Println("Hash after SHA-256: ", hash)

	// secret := []byte("my-secret")
	// message := []byte("Hello World")

	// h := hmac.New(sha256.New, secret)
	// h.Write(message)

	// text := h.Sum(nil)

	// fmt.Println(text)
	// jwt, err := utility.CreateJwt("my-secret", utility.Payload{
	// 	Sub:         20,
	// 	FirstName:   "Shahriar",
	// 	LastName:    "Hasan",
	// 	Email:       "shahriart29@gmail.com",
	// 	IsShopOwner: false,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// fmt.Println(jwt)

}
