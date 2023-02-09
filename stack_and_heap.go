package main

import "fmt"

func main() {
// b := read()
// b[0] = 0
// fmt.Println(b)

 c := make([]byte, 32)
 read2(c)
 fmt.Println(c)
}

//func read() []byte{
//    a := make([]byte, 32)
//    return a
//}

func read2(b []byte){
    b[0]=1
}