package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/url"
	"regexp"
)

func main() {

	URL := "https://mydoamin.com/mysubpage/mypage.html?test=1&c=2"

	urlPath, err := getURLPath(URL)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(urlPath)

	fmt.Println("===============MD5============")

	fmt.Printf("%s", mymd5("i am a string"))
	fmt.Println(" ")
	fmt.Println(mymd5("i am a string"))

}

func getURLPath(URL string) (string, error) {

	matched, err := validateHTTPURL(URL)
	if err != nil {
		return "", err
	}

	if !matched {
		return "", errors.New(URL + " is not valid url")
	}

	u, err := url.ParseRequestURI(URL)
	if err != nil {
		return "", err
	}

	return u.Path, nil
}

func validateHTTPURL(URL string) (bool, error) {

	matched, err := regexp.Match(`^https?:\/\/.*$`, []byte(URL))
	if err != nil {
		return false, err
	}
	return matched, nil
}

func mymd5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return hex.EncodeToString(h.Sum(nil))
}
