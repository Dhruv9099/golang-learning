# Basic-Printing-Functions

Go’s fmt package offers several functions for printing, each serving different purposes:

Println:

    The Println function is straightforward. It prints each argument separated by a space and adds a newline at the end.

Printf:

    The Printf function is more versatile as it allows formatted printing. You use format specifiers to define how you want the output to be displayed.

## Format Specifiers

Format specifiers define how a value should be formatted when printed. Here are some commonly used ones:

1. %v: Default format.
2. %T: Type of the value.
3. %d: Decimal integer.
4. %s: String.
5. %f: Floating-point number.
6. %t: Boolean.
7. %p: Pointer.

%v Specifier

The %v specifier is particularly useful because it provides the default representation of a value. It’s a versatile specifier that adapts to different types.

    package main

    import "fmt"

    func main() {
        i := 42
        s := "hello"
        arr := []int{1, 2, 3}
        fmt.Printf("Integer: %v\n", i)
        fmt.Printf("String: %v\n", s)
        fmt.Printf("Array: %v\n", arr)
    }

output:

    Integer: 42
    String: hello
    Array: [1 2 3]

## Additional Formatting Options

1. %+v: Includes field names in structs.
2. %#v: Go-syntax representation of the value.
3. %q: Quoted string.

        package main

        import "fmt"

        type Person struct {
            Name string
            Age  int
        }

        func main() {
            p := Person{Name: "Alice", Age: 25}
            fmt.Printf("Default format: %v\n", p)
            fmt.Printf("Include field names: %+v\n", p)
            fmt.Printf("Go-syntax representation: %#v\n", p)
            fmt.Printf("Type: %T\n", p)
            fmt.Printf("Quoted string: %q\n", "hello")
        }

output:

        Default format: {Alice 25}
        Include field names: {Name:Alice Age:25}
        Go-syntax representation: main.Person{Name:"Alice", Age:25}
        Type: main.Person
        Quoted string: "hello"

Printing Structs:
When dealing with structs, %+v and %#v become extremely useful for debugging as they provide more detailed information.

    package main

    import "fmt"

    type Address struct {
        City    string
        Country string
    }

    type Person struct {
        Name    string
        Age     int
        Address Address
    }

    func main() {
        p := Person{
            Name: "John",
            Age:  30,
            Address: Address{
                City:    "New York",
                Country: "USA",
            },
        }
        fmt.Printf("Person: %+v\n", p)
        fmt.Printf("Person: %#v\n", p)
    }

output:

    Person: {Name:John Age:30 Address:{City:New York Country:USA}}
    Person: main.Person{Name:"John", Age:30, Address:main.Address{City:"New York", Country:"USA"}}

## Conclusion

Understanding how to use Println, Printf, and various format specifiers can greatly enhance your ability to debug and interact with your Go programs. The %v specifier is particularly useful for its versatility, and exploring other specifiers like %+v, %#v, and %q can provide even more control over your output.

## Happy coding
