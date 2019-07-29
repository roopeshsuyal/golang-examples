package main

import (
	"fmt"
	"regexp"	
        "crypto/md5"
        "io"
        "encoding/hex"
)

func main() {

	url := "https://mydoamin.com/mypage.html?test=1&c=2"
	matched, err := regexp.Match(`^https?:\/\/.*$`, []byte(url))
	if err != nil {
		fmt.Println(err)
		return
	}

	if matched {
		fmt.Println("valid")
	} else {
		fmt.Println("Invalid")
	}
	
	fmt.Printf("%s",mymd5("i am a string"))
	fmt.Println(" ")
	fmt.Println(mymd5("i am a string"))

}


func mymd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
