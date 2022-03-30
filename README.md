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