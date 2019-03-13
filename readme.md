Some useful information when dealing with Go

1. `go run file_name.go` - This will pass through our go code step by step. Kinda like an interpreter
2. General go project structure looks something like This
    GoWorkspace (directory)
    |------ src (this contains the executable file)
            |------ hello (executable)
            |------ hello.go (source code file)
            |------ rogerokello.com (directory)
                    |----- hello (executable)
    |------ bin (this directory has the final executable while building the src using  `go build src/rogerokello.com/hello`)
            |------ hello
3. Generally the `GOPATH` must be set to point to the GoWorkspace directory for the command `go build src/rogerokello.com/hello` to become successful at creating a bin directory and placing the final executable there
