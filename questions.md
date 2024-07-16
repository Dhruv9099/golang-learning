1. Why should one learn Golang? What are the advantages of Golang over other languages?

        Go language follows the principle of maximum effect with minimum efforts. Every feature and syntax of Go was developed to ease the life of programmers. Following are the advantages of Go Language:

        Simple and Understandable: Go is very simple to learn and understand. There are no unnecessary features included. Every single line of the Go code is very easily readable and thereby easily understandable irrespective of the size of the codebase. Go was developed by keeping simplicity, maintainability and readability in mind.

        Standard Powerful Library: Go supports all standard libraries and packages that help in writing code easily and efficiently.

        Support for concurrency: Go provides very good support for concurrency using Go Routines or channels. They take advantage of efficient memory management strategies and multi-core processor architecture for implementing concurrency.

        Static Type Checking:
        Go is a very strong and statically typed programming language. Statically typed means every variable has types assigned to it. The data type cannot be changed once created and strongly typed means that there are rules and restrictions while performing type conversion. This ensures that the code is type-safe and all type conversions are handled efficiently. This is done for reducing the chances of errors at runtime.

        Easy to install Binaries:
        Go provides support for generating binaries for the applications with all required dependencies. These binaries help to install tools or applications written in Go very easily without the need for a Go compiler or package managers or runtime.

        Good Testing Support:
        Go has good support for writing unit test cases along with our code. There are libraries that support checking code coverage and generating code documentation.


2. What are Golang packages?

        Go Packages (in short pkg) are nothing but directories in the Go workspace that contains Go source files or other Go packages themselves. Every single piece of code starting from variables to functions are written in the source files are in turn stored in a linked package. Every source file should belong to a package.

        From the image below, we can see that a Go Package is represented as a box where we can store multiple Go source files of the .go extension. We can also store Go packages as well within a package