package main

import "fmt"

// type Rect struct {
// 	width int
// 	height int
// }

// func (rect Rect) area() int {
// 	return rect.width * rect.height
// }

type authenticationInfo struct {
	username string
	password string
}

func (authInfo authenticationInfo) getBasicAuth() string {
	output := fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
	return output
}

func GETBASICAUTH(authInfo authenticationInfo) string {
	return fmt.Sprintf("Authorization: Basic %s:%s", authInfo.username, authInfo.password)
}

func test (authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
}



func main () {
	// rect1 := Rect{
	// 	width: 10,
	// 	height: 20,
	// }
	// rect2 := Rect{
	// 	width: 30,
	// 	height: 40,
	// }
	// fmt.Println("Area of rect1: ", rect1.area())
	// fmt.Println("-------------------------------")	
	// fmt.Println("Area of rect2: ", rect2.area())

	user1 := authenticationInfo{
		username: "Tahsin",
		password: "123456",
	}
	fmt.Println("-------------------------------")
	fmt.Println("Without using method for the struct")
	fmt.Println(GETBASICAUTH(user1))
	fmt.Println("-------------------------------")
	fmt.Println("Using method for the struct")
	test(user1)
	fmt.Println("-------------------------------")
}