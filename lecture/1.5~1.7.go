package main

//for and range
func superAdd(numbers ...int) int {
    total := 0
    for _, number := range numbers{
        total += number
    }
    return total
}

//if else switch, variable expressions
func canIDrink(age int) bool {
    if koreanAge := age + 2; koreanAge < 20 {
        return false
    }
    return true
    // switch koreanAge := age + 2; koreanAge{
    // case koreanAge < 20:
    //     return false
    // case koreanAge >= 20:
    //     return true
    // }
}

