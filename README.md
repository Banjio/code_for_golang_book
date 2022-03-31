# code_for_golang_book
Code for learning golang

## Getting started

1. Create a go.mod file that tracks the dependencies (similiar to requirements.txt in python) of your project (package) inside the project root. When creating this file you need to give it a path where the package is living, e.g. `github.com/mymodule`. For the sack of learning we can do something like

    $ go mid init example/hello

2. Go is a compiled language, thus for running a programm we can use `go run`

When the file you want to run is inside the same directory as your console we can run the hello.go programm by using

    $ go run . 

The command 

    $ go help 
    
will give you more information about go


### Call code from external packages

1. Under pkg.go.dev you can browse published packages. 
2. If you found a package you want to include you can see under index what to import
3. If you import the package for the first time it will throw an error. This is because the go mod file need to be updated so that it includes the imported module
4. Call `go mod tidy` so that the requirements are downloaded

### Create a module callable from another module

https://go.dev/doc/tutorial/create-module

1. Create module in the same way you created hello_world
2. Create a function that takes input parameter(s) and returns somethin
   
```go
func TestFunc(param string, param1 int) nil {
	return nil
}
```

3. Call code from other module
   1. For that we edit the `hello.go` file so that it uses the `greetings.Hello`  function and import example.com/greetings
   2. In order for the hello package to find greetings we have to edit the go mod file `$ go mod edit -replace example.com/greetings=../greetings`
   3. Once again run `go mod tidy` so that all dependencies are synced

### Error handling

https://go.dev/doc/tutorial/handle-errors

1. See modified `greetings.Hello` function
2. See modified `hello.main` function


### Slice type

https://go.dev/doc/tutorial/random-greeting

* Go slices are like list in slices. Unlike gos arrays slices change dynamically as items are added or removed
* `[]string{}`initializes a slice type `[]` omits the size of the array making it a slice
* init functions are automatically called on statup of th eprogramm

### For loop and maps

https://go.dev/doc/tutorial/greetings-multiple-people

* We can also add slices to functions and let it return maps like this func A(list []string) (map[string], error){}
* In go maps are initialized with `make(map[key-type]value-type)`
* `for` loop in `range`returns two value the index of the current item and a copy of the item's value. Similiar like in python you can use the balank identifier `_`if you want to omit one of both. &Rightarrow; Syntax of loop: `for _, name := range names { ... }`

### Build in test module

https://go.dev/doc/tutorial/add-a-test

* Files that end with `_test` are automatically recogniced by `$ go test`. 
* Test funcion have the form `TestName`
* Giving a pointer to the T type  inside a test function `func MyTest(t * testing.T)`, thus it methods can be used (similiar method like assert and) for testing logging
*  Run tests with `$ go test` or `$ go test -v` for verbose output

### Build the programm hello_world with

* `$ go build` inside a go directory will build the package to an executable, which has the correct format depending on the plattform architecture
* Currently you can only execute the programm inside the directory 
Thus to install a programm and make it executable from anywhere on the system we have to do the following:

* To check gos install path run `$ go list -f '{{.Target}}'`, e.g., the command's output might say /home/gopher/bin/hello, meaning that binaries are installed to /home/gopher/bin. You'll need this install directory in the next step.
* To add the go install path to your shell run `$  export PATH=$PATH:/home/maxbeier/go/bin/`
* Or alternatively if you already have a bin path on your path where yo want to install programm you can do `go env -w GOBIN=/path/to/your/bin`
* Finally run `$ go install`to install the program

## Multi Module workspaces

If you develop code in multiple packages in the same folder you need to set up a multi module workspaces, because otherwise the go language server will throws errors and tell "you are not on a go path" (That filthy language server tho)

# General remarks

* Only functions with upper case are exported for use in other modules (thus functions with lower case are private)


# PeP 8 of GO 

https://go.dev/doc/effective_go
