package main

import "fmt"

func check(e error, message string) {
    if e != nil {
        fmt.Printf("%s - %e", message, e)
        // panic(e)
    }
}