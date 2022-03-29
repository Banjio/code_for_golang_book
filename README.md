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