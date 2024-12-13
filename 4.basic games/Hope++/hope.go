package main

import "fmt"

// Simple Hope game Task from Quera
func main() {
    var p, q int
    fmt.Scanf("%d %d", &p, &q)
    
    for i := 1 ; i <= q ; i++ {
        quotient := i / p
        if i % p == 0{
            for j := 0 ; j < quotient ; j++{
                fmt.Print("Hope ")
            }
            fmt.Println("")
        } else {
            fmt.Println(i)
        }
    } 
}