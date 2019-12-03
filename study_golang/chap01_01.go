package main

import "fmt"

func main1() {
    fmt.Println("Hello, 世界")
}


func main()  {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}