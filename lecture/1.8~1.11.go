package main

import "fmt"

//struct
type person struct {
    name string
    age int
    favFood []string
}

func main() {
    //pointer
    a := 2
    b := &a
    c := &b
    fmt.Println(**c, a)

    //array
    names := [5]string{"nico", "lynn", "dal"}
    names[3] = "aaa"
    names[4] = "asdfa"
    // fmt.Println(names)

    //slice
    nameSlice := []string{"nico", "lynn", "dal"}
    nameSlice = append(nameSlice, "what")
    // fmt.Println(nameSlice)

    //maps
    nico := map[string]string{
        "name":"nico",
        "age": "12",
    }
    for keys, value := range nico {
        fmt.Println(keys, value)
    }

    //struct
    favFood := []string{"kimchi", "ramen"}
    justin := person{name:"nico", age:18, favFood: favFood}
    fmt.Println(justin)
}