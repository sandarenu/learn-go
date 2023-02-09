package main

import (
	"fmt"
)

type Product struct{
    Name string
    Amount float32
}

func (p *Product) giveOffer(discount int){
    if p == nil {
        return
    }
    p.Amount = p.Amount - (p.Amount * float32(discount)/100.0)
    fmt.Println("Amount after offer", p.Amount)
}

func main() {

    a := 100
    fmt.Printf("A %d\n", a)
    pa := &a
    fmt.Printf("Pointer to A %d\n", pa)
    fmt.Printf("Value of pointer pa %d\n", *pa)
    *pa += 200
    fmt.Printf("Value of pointer pa %d\n", *pa)
    fmt.Printf("A %d\n", a)

    double(a)
    fmt.Printf("A %d\n", a)

    doubleP(&a)
    fmt.Printf("A %d\n", a)
    fmt.Printf("Value of pointer pa %d\n", *pa)

    var apple Product = Product {
        "Apple",
        2.50,
    }

    fmt.Printf("Apple %#v \n", apple)
//    giveOffer(&apple, 10)
//    fmt.Printf("Apple %#v \n", apple)

    apple.giveOffer(10)
    fmt.Printf("Apple %#v \n", apple)

    banana := &Product{
        "Banana",
        1.50,
    }
    fmt.Printf("Banana %#v \n", banana)

    var orange *Product

    fmt.Printf("Orange %#v \n", orange)
    orange.giveOffer(10)


}


func giveOffer(p *Product, discount int){
    p.Amount = p.Amount - (p.Amount * float32(discount)/100.0)
    fmt.Println("Amount after offer", p.Amount)
}

func double(a int) int {
	a = a * 2
	return a
}

func doubleP(a *int){
    *a = *a * 2
}

//func doubleP(a *int) int {
//	*a = *a * 2
//	return *a
//}
