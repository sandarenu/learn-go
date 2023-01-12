package main
import (
    "fmt"
)

func  main()  {

    s1 := []string {"A", "B", "C"}
    s2 := []string {"P", "Q", "R"}

    //Join two slices
    s3 := append(s1, s2...)
    fmt.Println(s3)

    //Subset of a Slice
   s4 := s1[:2]
   fmt.Println("before", s4)
    //Update origial slice
   s1[1] = "N"
   fmt.Println("after 1", s4)
   //Update subset
   s4[0] = "P"
   fmt.Println("after 2", s4)
   fmt.Println("after 3", s1)

   //Copy slice
   s5 := make([]string, 2)
   fmt.Println(len(s5), cap(s5), s5)
   copy(s5, s1[:2])
   fmt.Println("copy", s5)

   a1 := [3]string {"a", "b", "c"}
   printSlice(a1[:]) //convert array to a slice

}

func printSlice(s []string){
    fmt.Println(s)
}