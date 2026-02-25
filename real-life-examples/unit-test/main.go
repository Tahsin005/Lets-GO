package main

import "fmt"

func main() {
	user1 := User{
		Name: "tahsin",
		Email: "tahsin.ferdous3546@gmail.com",
	}

	if err := validate(user1); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}

	user2 := User{
		Name: "t",
		Email: "tahsin.ferdous3546gmail.com",
	}

	if err := validate(user2); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}

	user3 := User{
		Name: "tahsin",
		Email: "tahsin.ferdous3546gmail.com",
	}

	if err := validate(user3); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}

	user4 := User{
		Name: "tahsin",
		Email: "tahsin.ferdous3546@gmail.com",
	}

	if err := validate(user4); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}

	user5 := User{
		Name: "tahsin",
		Email: "",
	}

	if err := validate(user5); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}

	// t := reflect.TypeOf(user)

	// fmt.Println(t.Name())
	// fmt.Println(t.Kind())

	// for i := 0; i < t.NumField(); i++ {
	// 	field := t.Field(i)

	// 	tag := field.Tag.Get("validate")

	// 	fmt.Printf("%d. %v (%v), tag: '%v'\n", i + 1, field.Name, field.Type.Name(), tag)
	// }
}