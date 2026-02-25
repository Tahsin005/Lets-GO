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
			var err error

			switch {
			case rule == "required":
				err = validateRequired(fieldType, fieldValue)

			case rule == "email":
				err = validateEmail(fieldType, fieldValue)

			case strings.HasPrefix(rule, "min="):
				err = validateMinLength(fieldType, fieldValue, rule)

			case strings.HasPrefix(rule, "max="):
				err = validateMaxLength(fieldType, fieldValue, rule)
			}

			if err != nil {
				return err
			}
		}
	}

	return nil
}

func validateRequired(fieldType reflect.StructField, fieldValue reflect.Value) error {
	if fieldValue.Kind() == reflect.String && fieldValue.Len() == 0 {
		return fmt.Errorf("%s is required", fieldType.Name)
	}
	return nil
}

func validateEmail(fieldType reflect.StructField, fieldValue reflect.Value) error {
	if fieldValue.Kind() == reflect.String {
		val := fieldValue.String()
		if !strings.Contains(val, "@") {
			return fmt.Errorf("%s must be a valid email", fieldType.Name)
		}
	}
	return nil
}

func validateMinLength(fieldType reflect.StructField, fieldValue reflect.Value, rule string) error {
	min, _ := strconv.Atoi(strings.TrimPrefix(rule, "min="))
	if fieldValue.Kind() == reflect.String && fieldValue.Len() < min {
		return fmt.Errorf("%s must be at least %d characters", fieldType.Name, min)
	}
	return nil
}

func validateMaxLength(fieldType reflect.StructField, fieldValue reflect.Value, rule string) error {
	max, _ := strconv.Atoi(strings.TrimPrefix(rule, "max="))
	if fieldValue.Kind() == reflect.String && fieldValue.Len() > max {
		return fmt.Errorf("%s must be at most %d characters", fieldType.Name, max)
	}
	return nil
}
