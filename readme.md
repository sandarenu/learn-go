# Install GO
https://go.dev/doc/install

# Create module

```
go mod init learn-go
```

## GO basics

* Package `main`
  * Special package `main`, which is required if this code is going to be run directly (usually from the terminal).
* Function `main`
  * Starting point of the code execution
* Go is **statically typed**
* Variables are initialized to **zero value** for its type.
* `:=` can use as short variable declaration.
* If variable is declared, it has to be used in the code. Unless it will be a compilation error.
* If we don’t intend to use a variable, we can use Go’s **blank identifier** underscore(`_`).

### Naming Rules

* Name should start with a letter
* If the name of a variable, function, or type **begins with a capital letter**, it is considered exported and can be accessed from packages outside the current one. 
* Names starting with a lowercase letter, it is considered unexported and can only be accessed within the current package.

### Arrays

* Arrays are initialized to zero value of its type
* Array should be defined in single line, but multiline can be used but it should have a trailing comma.
  ```
  a3 := []string{
		"This",
		"is",
		"a",
		"test",
	}
  ```
  * Array vs Slice
    * `var array [3]int`
    * `var slice []int`
  * Use `make` to create slices. Eg: `s2 := make([]int, 10)`

### Slice

https://ueokande.github.io/go-slice-tricks/
https://github.com/golang/go/wiki/SliceTricks

### Excercises

https://github.com/inancgumus/learngo