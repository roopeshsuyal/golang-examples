package main

import (
	"fmt"
	"github.com/asaskevich/govalidator"
)

func main() {

	url := "http://mydoamin.com/mypage.html?test=1&c=2"
	err := validateURL(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(url + " is valid URL")

}

func validateURL(u string) error {
	valid := govalidator.IsRequestURL(u)
	if valid == false {
		return fmt.Errorf("%v is a invalid url", u)
	}
	return nil
}
