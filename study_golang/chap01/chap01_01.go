package main

import (
    "fmt"
    "os"
    "bufio"
    "io/ioutil"
    "strings"
)

func main1_1() {
    fmt.Println("Hello, 世界")
}


func main1_2()  {
    var s, sep string
    //声明并初始化（下面四种全都等价）
    // s := ""  (只能用在函数内部，不能用于包变量)
    // var s string (直接初始化为0)
    // var s = ""
    // var s string = ""    ()
    // 两种循环
    for i := 0; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }

    // for _, arg := range os.Args[1:] {
    //     s += sep + arg
    //     sep = " "
    // }
    fmt.Println(s)
}


func main_dup1() {
    counts := make(map[string]int)
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        counts[input.Text()]++
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}


func main_dup2() {
    counts := make(map[string]int)
    files := os.Args[1:]
    if len(files) == 0 {
        countLines(os.Stdin, counts)
     }else {
        for _, arg := range files {
            f, err := os.Open(arg)
            if err != nil {
                fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
                continue
            }
            countLines(f, counts)
            f.Close()
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
            fmt.Printf("当前文件 %s 出现重复行", string(files[0]))
        }
    }
}

func countLines(f *os.File, counts map[string]int) {
    input := bufio.NewScanner(f)
    for input.Scan() {
        counts[input.Text()]++
    }
}


func main_dup3() {
    counts := make(map[string]int)
    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil{
            fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
        }
    }
    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\n", n, line)
        }
    }
}