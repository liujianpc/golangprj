package main

import "os"
import "fmt"
import "simplemath"
import "strconv"

var Usage = func() {
    fmt.Println("Usage:Calc command [arguments]...")
    fmt.Println("\nThe command are:\n\tadd\tAddition of two values.\n\tsqrt\t Sqrt root of one integer number")
}

func main() {
    args := os.Args
    if args == nil || len(args) < 3 {
        Usage()
        return
    }
    fmt.Println(args[1])

    switch args[1] {
    case "add":
        if len(args) != 4 {
            fmt.Println("USAGE: calc add <integer1><integer2>")
            return
        }
        v1, err1 := strconv.Atoi(args[2])
        v2, err2 := strconv.Atoi(args[3])
        if err1 != nil || err2 != nil {
            fmt.Println("USAGE: calc add <integer1><integer2>")
            return
        }
        ret := simplemath.Add(v1, v2)
        fmt.Println("Result: ", ret)

    case "sqrt":
        if len(args) != 3 {
            fmt.Println("USAGE: calc sqrt <integer>")
            return
        }
        v3, err3 := strconv.Atoi(args[2])
        if err3 != nil {
            fmt.Println("USAGE:calc sart <integer>")
            return
        }
        ret := simplemath.Sqrt(v3)
        fmt.Println("Result: ", ret)

    default:
        //Usage()

    }
}
