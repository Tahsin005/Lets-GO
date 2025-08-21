package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/base-7/description/?envType=problem-list-v2&envId=n6ww7n9h

func convertToBase7(num int) string {
	return fmt.Sprintf(strconv.FormatInt(int64(num), 7))
}

func main () {

}