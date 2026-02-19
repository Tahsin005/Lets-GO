package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type User struct {
	Name  string `validate:"min=2,max=10"`
	Email string `validate:"required,email"`
}

func validate(val interface{}) error {
	t := reflect.TypeOf(val)
	v := reflect.ValueOf(val)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return errors.New("validate: expected a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		tag := fieldType.Tag.Get("validate")
		if tag == "" {
			continue
		}

		rules := strings.Split(tag, ",")

		for _, rule := range rules {

			switch {

			case rule == "required":
				if fieldValue.Kind() == reflect.String && fieldValue.Len() == 0 {
					return fmt.Errorf("%s is required", fieldType.Name)
				}

			case rule == "email":
				if fieldValue.Kind() == reflect.String {
					val := fieldValue.String()
					if !strings.Contains(val, "@") {
						return fmt.Errorf("%s must be a valid email", fieldType.Name)
					}
				}

			case strings.HasPrefix(rule, "min="):
				min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
				if fieldValue.Kind() == reflect.String && fieldValue.Len() < min {
					return fmt.Errorf("%s must be at least %d characters", fieldType.Name, min)
				}

			case strings.HasPrefix(rule, "max="):
				max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
				if fieldValue.Kind() == reflect.String && fieldValue.Len() > max {
					return fmt.Errorf("%s must be at most %d characters", fieldType.Name, max)
				}
			}
		}
	}

	return nil
}

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