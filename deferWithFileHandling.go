package main

import (
    "fmt"
    "os"
)

func main() {
    // Example with multiple defers
    fmt.Println("Opening files and creating resources")

    file1, err := os.Create("file1.txt")
    if err != nil {
        fmt.Println("Error creating file1:", err)
        return
    }
    defer func() {
        fmt.Println("Closing file1")
        file1.Close()
    }()

    file2, err := os.Create("file2.txt")
    if err != nil {
        fmt.Println("Error creating file2:", err)
        return
    }
    defer func() {
        fmt.Println("Closing file2")
        file2.Close()
    }()

    fmt.Println("Writing to files")
    file1.WriteString("Hello, this is file 1")
    file2.WriteString("Hello, this is file 2")

    fmt.Println("All done")
}
