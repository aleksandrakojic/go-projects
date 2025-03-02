package main

import (
    "fmt"
)

func Filter[T any](s []T, f func(T) bool) []T {
    var r []T
    for _, v := range s {
        if f(v) {
            r = append(r, v)
        }
    }
    return r
}

func main() {
    s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    r := Filter(s, func(v int) bool { return v % 2 == 0 })
    fmt.Println(r)

    shortStrings :=  Filter([]string{"ok", "notok", "maybe", "maybe not"}, func(s string) bool { return len(s) < 3 })     
    fmt.Println(shortStrings)
}