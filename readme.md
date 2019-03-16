Some useful information when dealing with Go

1. `go run file_name.go` - This will pass through our go code step by step. Kinda like an interpreter
2. General go project structure looks something like This
    GoWorkspace (directory)
    |------ src (this contains the executable file)
            |------ hello (executable)
            |------ hello.go (source code file)
            |------ rogerokello.com (directory)
                    |----- hello (executable)
    |------ bin (this directory has the final executable while installing the src using  `go install rogerokello.com/hello`)
            |------ hello
3. Generally the `GOPATH` must be set to point to the GoWorkspace directory for the command `go install rogerokello.com/hello` to become successful at creating a bin directory and placing the final executable there
4. Resource on how to setup a golang development environment on a mac `https://medium.com/@AkyunaAkish/setting-up-a-golang-development-environment-mac-os-x-d58e5a7ea24f`
5. Functions:
  - They are actually types in go and can return multiple values
  - They can be used like any other types and can be passed around just like they are in other languages like javascript
  - They are literals and can remember the context of where they are defined (closures)
  - They do not specify void incase they do not return a value
6. Variables:
  - In go, you cannot declare a Variables and not use them
7. Branching
  `If statement general structure(Optional else block)`
  - There are no angle brackets between the optional statement and the condition
  - The curly brackets are mandatory
  ```
  if (optional statement); condition {
    block
  }
  ```
  ```
  if (optional statement); condition {
    block
  }else{
    block
  }
  ```

  `Switch  statement`
  - There is no default fall through. So you do not have to put break statements
    like there are in other programming languages switch statements.
  - There is a keyword `fallthrough` that will make a case statement fall through
    to the next block.
  - It behaves like a set of if-else statements unless specified to behave otherwise
  - The cases can actually be expressions unlike in other programming languages
    where the cases are constants or strings that have been evaluatated at the
    beginning of the switch block.
  - It is possible to switch on types.

8. Looping (`The For statement`)
  - There is only one keyword for looping. However it can be manipulated in countless
    ways to suite the same style in other programming languages.
  - These are the several ways it could look
    1. Only specifying a condition
    NB: The condition must evaluate to true or false
    ```
    for condition {
      block
    }
    ```
    2. The For clause
    ```
    for initialiser; condition; post {
      block
    }
    ```
    3. With the range construct
    NB: A collection may be a map, slice or array, string, channel
    ```
    for i := range collection {
      block
    }
    ```
    4. For without a while loop. Internally mimicking an infinite loop where
       one may have to break out using a break statement
    ```
    for {
      block
      break // to break out of the for loop
      continue // to stop execution at that point and continue the next iteration
    }
    ```
