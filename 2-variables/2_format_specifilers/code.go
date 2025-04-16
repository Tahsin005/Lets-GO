package main

import (
    "fmt"
)

func main() {
    // Integer types
    var age int = 25
    var smallNumber int8 = -128
    var bigNumber uint64 = 1234567890

    // Floating-point types
    var price float32 = 19.99
    var pi float64 = 3.14159265359

    // Boolean
    var isAvailable bool = true

    // String and character types
    var name string = "GoLang"
    var initial byte = 'G'      // byte is alias for uint8
    var symbol rune = '♥'       // rune is alias for int32 (Unicode)

    // Print all variables using fmt.Printf with correct format specifiers
    fmt.Println("🔸 Integer Types:")
    fmt.Printf("age (int): %d\n", age)
    fmt.Printf("smallNumber (int8): %d\n", smallNumber)
    fmt.Printf("bigNumber (uint64): %d\n", bigNumber)

    fmt.Println("\n🔸 Floating Point Types:")
    fmt.Printf("price (float32): %.2f\n", price)
    fmt.Printf("pi (float64): %.5f\n", pi)

    fmt.Println("\n🔸 Boolean:")
    fmt.Printf("isAvailable (bool): %t\n", isAvailable)

    fmt.Println("\n🔸 String and Characters:")
    fmt.Printf("name (string): %s\n", name)
    fmt.Printf("initial (byte): %c, ASCII: %d\n", initial, initial)
    fmt.Printf("symbol (rune): %c, Unicode: %U\n", symbol, symbol)

    fmt.Println("\n🔸 Type Information:")
    fmt.Printf("Type of 'age': %T\n", age)
    fmt.Printf("Type of 'price': %T\n", price)
    fmt.Printf("Type of 'name': %T\n", name)
}
