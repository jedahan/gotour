package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) (counts map[string]int) {
    counts = make(map[string]int)
    
    for _, v := range strings.Fields(s) {
        counts[v]++
    }

    return
}

func main() {
    wc.Test(WordCount)
}