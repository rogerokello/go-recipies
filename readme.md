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
  - References
    1. Function usage Advanced: http://golang.org/doc/codewalk/functions/
    2. Function types: http://golang.org/ref/spec#Function_types
6. Variables:
  - In go, you cannot declare a Variables and not use them
  - References
    1. Function types: http://golang.org/ref/spec
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
    - References
      - http://golang.org/ref/spec#For_statements
      - http://golang.org/doc/effective_go.html#for
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
9. Maps

  - Analogous to Dictionaries in python. They have unique keys that map to values
  - Keys in maps need to have equality operators defined for them. This means if
    you want to use a type for a key in your map, it has to be one that allows
    for equal comparisons. Most built in types (strings, ints, floats) support the
    equality operator. This is important because certain types like a slice do not
    have an equality operator defined for them.
  - Maps are reference types. This means that if you have a map variable it behaves
    very much like a pointer. When you pass a map to function, any changes made to
    the map within the function will propagate to the original value because the
    actual reference is passed in an not a copy.
  - Maps are not thread safe. Do not use maps in concurrent code unless you
    absolutely have to.
  `Map operations`
    1. Insert
    2. Update
    3. Delete
    4. Check for Existence
  - References
    1. https://golang.org/doc/effective_go.html#maps

10. Slices
  - These are similar to datatypes such as lists in other languages and lend
    themselves the benefits of such datatypes.
  - `What is an array?`
    This is basically a collection that has a fixed size of a particular type.
    The array is stored in a sequential memory location.
    In go, the syntax for creating an array may look like this: `var n[3]int`
    `Properties`
    1. They are fixed sized
    2. Array type is size and underlying type. An array of 3 ints is different from
       an array of 5 ints.
    3. No initialisation required (0 valued). All values are zeroed out.
    4. It is not a pointer. It is a value type. This means if you pass an Array
       to a function, that function will get a copy of the array and will actually
       be modifying the copy and not the actual type as opposed to a reference type
       like a map.
  - `What is a slice?`
    It is an abstraction over an array. The problem with an array is that it is a
    fixed size and the different arrays of different sizes are different types.
    This becomes more of a problem because if you wanted to use an array as a
    parameter to function then size would limit the number of arrays that you would
    be able to pass in there.
    Slices of different sizes are still the same type. This is were the slices will
    beat the Arrays.
    `Basic syntax for a slice: var aSlice []type = make([]type, length, capacity)`
    - A slice is actually a pointer to an array.
    - When working with a slice you will be taking a portion of that array and working
      with it.
    `Properties of a slice`
    1. A slice is of a fixed size but we can use append to make it grow and add more
       to it.
    2. Type is slice of the underlying type and it does not have a length. In short, if
       you pass it as a parameter to a function you only need to specify the type and
       not the length as you would an array.
    3. A slice must be initialised using `make` otherwise it will be nil. If one does not
       use make to to initialise, then the slice will be nil.
    4. A slice points to an array. A slice is actually a reference type unlike an array which
       is a value type. If you pass a slice to function and you modify that slice from with in
       the function, it is going to modify the original array pointed to by the slice.
       Modifying the slice of a slice will affect the original as well.
    5. Because slices are arrays underneath, we use the `append` to add items to the slice.
       In case the slice was made with lower capacity a new underlying array is created and
       all the previous items copied to it together with the new item to be added. This is
       normally an expensive operation and setting a slice using make with a larger capacity
       is preferred. Basic syntax of append is `newSlice = append(slice, itemToAddToSlice)`.
       `append` takes a variable number of arguments of the type of slice.

10. Go Methods
    Define methods by defining a function which specifies which types it operates on. It will behave just
    like a normal function with parameters and return values. Go methods can operate on any named types.
    It will operate on any named type in the same package as it was declared.
    `NB`: A classical example could be you could rename an integer type to myInt and then define a method
    on it. What this just means is that we can create a method for any type so long as we name that type first

11. Pointer methods
    If we create a method on a named type we cannot modify that named type, but if we create a method on a Pointer
    to a named type we can modify the named type. Think of a method in go as taking one additional parameter, which
    is a type.
    `Basic rule`: If you declare a method that is going to work on a pointer to a named type, then it is more specific,
    it is only going to work if you pass in a pointer to that named type.
    ```
      func(*myType)myMethod(i int)
    ```
    But if you declare a method that is going to work on a named type, then you can utilise than method from that named type
    itself or a pointer to that named type.
    ```
      func(myType)myMethod(i int)
    ```

12. Interfaces
    It just simple states that what does something do. If something obeys this contract/interface then it belongs to that
    interface type. In go we do not have to explicitly implement interfaces.
    `NB` Any type that has the same methods as an interface implements that interface.
    In order to implement an interface in Go all that has to happen is a named type has to implement the methods defined
    by the interface. This means one can use that type any where the interface is specified.
    It is possible for a function to take an interface as a parameter and then use that interface and know that those Methods
    that exist on that interface exist and then utilise it through that interface. This gives one the flexibility of passing
    in several types into that function and using them via that interface.

    `Empty Interface`
    If we define an interface that has no requirements then every single type in the system will implement that interface.
    ```
    interface {}
    ```
    - References
      1. https://golang.org/doc/effective_go.html#methods
      2. https://golang.org/doc/effective_go.html#interfaces

13. Concurrency
    `Note:` Do not communicate by sharing data but instead share data by communicating.
    1. Goroutines
      - It is basically a light weight thread and it is managed by the go runtime.
      - It uses the keyword `go` placed in front of a function. When placed in front of a function, it is basically saying
        run this and keep on going. I do not want to block waiting for this to finish. We do not have to worry that this is
        another thread or how this thread is being created or the life cycle of this thread. Instead it is the go runtime
        that will manage all that.
      - You if you create a go routine in main, you have to be careful enough to make sure it gets completed before
        the main exists. ie. main may have to wait for it to complete.

    2. Channels
      - Because one has to synchronise how goroutines get run and ensure main does not complete before the go routines,
        go uses channels to synchronise this communication between goroutines. Actually main is also a go routine.
      - Essentially what happens is one go routine (goroutine 1) feeds a channel with something and another go routine
        (goroutine 2) waits to extract something out of that channel.
      `Channel types`
      1. Un Buffered
         This means before something new is put onto the channel, we have to wait for whatever was present to be removed
         fro the channel. In short the channel can only be used once at a time. If we have to read from the channel,
         we have to wait until data is actually in the channel.
      2. Buffered
         This is going to allow for how many items have to be in the channel before we actually block on the channel.
      NB: The buffer specifies how many times you have to write to a channel before that data starts being read from that
          channel. In some cases data is read from a buffered channel as long as something has now come through. For example
          If a channel has a buffer size of 2 and one item has already been sent to it but other lines of code continue to
          execute before other items are sent to channel, then it means an item can be read off from the channel before
          another item is place to it.
          It appears when one adds more items to a buffer than its size, it blocks until all items are read.

      NB: A channel is actually used to communicate between threads.
