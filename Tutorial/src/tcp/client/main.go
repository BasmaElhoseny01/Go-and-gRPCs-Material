package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {
    // Connect to server
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err.Error())
        return
    }
    defer conn.Close()

    // Read input from user
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text to capitalize: ")
    text, _ := reader.ReadString('\n')

    // Send input text to server
    _, err = conn.Write([]byte(strings.TrimSpace(text)))
    if err != nil {
        fmt.Println("Error sending data:", err.Error())
        return
    }

    // Read response from server
    buffer := make([]byte, 1024)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error receiving data:", err.Error())
        return
    }

    // Print capitalized text received from server
    fmt.Println("Capitalized text:", string(buffer[:n]))
}
