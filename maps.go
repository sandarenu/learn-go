package main

import (
	"fmt"
)

func main() {
	var map1 map[string]string //declare the variable
	fmt.Printf("map1 %#v %d\n", map1, len(map1))
	map1 = make(map[string]string) //actually creates the map
    fmt.Printf("map1 after make %#v %d\n", map1, len(map1))
	map1["u1"] = "User 1"
	map1["u2"] = "User 2"
	fmt.Println("map1", map1, len(map1))
	map1["u2"] = "User 222"
	fmt.Println("map1", map1, len(map1))
	fmt.Println("map1 element 2: ", map1["u2"])

	map2 := map[string]int{"apple": 100, "oranges": 222} //map inline initialization (single line)
	fmt.Println("Fruits array: ", map2)

    //map inline initialization (multi line)
	map3 := map[string]int{
        "apple": 100,
		"oranges": 222,
        "grapes": 150,
    }
    fmt.Println("map3", map3)

    emptyMap := map[string]string{}
    fmt.Println("Empty map", emptyMap)

    intMap := make(map[string]int)
    intMap["a"]++ //items are initialized to zero values, so we can safely operate on it
    fmt.Println("Int map", intMap)

    //check element is assigned or not
    //not assigned element
    e1 := emptyMap["a"]
    e2, e2s := emptyMap["a"]
    fmt.Printf("E1 %#v %#v %#v \n", e1, e2, e2s)
    //element assigned
    emptyMap["b"] = "B"
    e3 := emptyMap["b"]
    e4, e4s := emptyMap["b"]
    fmt.Printf("E3 %#v %#v %#v \n", e3, e4, e4s)

    //delete elements
    delete(emptyMap, "b")
    e5 := emptyMap["b"]
    e6, e6s := emptyMap["b"]
    fmt.Printf("E5 %#v %#v %#v \n", e5, e6, e6s)

    delete(emptyMap, "b")

    for key, value := range map3 {
        fmt.Printf("Key [%s] Value[%d]\n", key, value)
    }

}
