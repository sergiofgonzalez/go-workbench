# 02 &mdash; Program Structure

## Intro

In Go, as in any other programming language, one builds large programs from a small set of basic constructs:
+ Variables store values
+ Simple expressions are combined into larger ones
+ Basic types are collected into aggregates like arrays and structs.
+ Expressions are used in statements whose execution order is determined by control-flow statements like `if` and `for`.
+ Statements are grouped into functions for isolation and reuse.
+ Functions are gathered into source files and packages.
+ Packages are grouped into modules for distribution.

## Names

+ If an entity is declared within a function, it is *local* to that function.

+ If declared outside of a function, it is visible to all the files of the package to which it belongs.

+ If the name begins with an upper-case letter, it is *exported*, which means that it is visible and accessible outside of its own package, and may be referred to by other parts of the program, as with `Printf` in the `fmt` package.

+ Package names are always lower case.

Generally, the larger the scope of a name, the longer and more meaningful it should be. Conversely, local variables with small scopes should be short (e.g., `i` for loops instead of `loopIndex`).

Camel-case is preferred when combining words in Go.

The letters of acronyms and initialisms like ASCII, HTML, or URL are always rendered in the same case. Thus, a function can be called `htmlEscape` or `HTMLEscape`, but not `HtmlEscape`.

## Declarations

There are four major kinds of declarations: `var`, `const`, `type`, and `func`.

A Go program is stored in one or more files whose names end `.go`. Each file begins with a package declaration that says what package the file is part of. Then it is followed by any `import` declarations, and then a sequence of *package-level* declarations, in any order.

## Variables

A var declaration creates a variable of a particular type, attaches a name to it, and sets its initial value.

Each declaration has the general form:

```go
var name type = expression
```

If the expression is omitted, the initial value is the *zero* value for the type:
+ `0` for numbers
+ `false` for booleans
+ `""` for strings
+ `nil` for interfaces and reference types (slice, pointer, map, channel, and function).

You can declare multiple variables of different type in a single declaration, and you can intialize a set of variables by calling a function that returns multiple values.

```go
var b, g, s = true, 2.5, "catorce"
var f, err = os.Open(name)
```

## Short variable declarations

Short variable declarations (because of their brevity and flexibility), are used to declare and initialize the majority of local variables:

```go
i := 100 // short variable declaration for an int
f, err := os.Open(name)
```

A short variable declaration does not necessarily declare all the variables on its left-hand side. If some of them were already declared in the same lexical block, then the short variable declaration acts like an *assignment* to those variables.

```go
in, err := os.Open(infile) // declares both in and err
// ...
out, err := os.Create(outfile) // declares out, assigns to err
```

In those cases, a short variable declaration must declare at least one new variable, otherwise the code will not compile.

```go
f, err := os.Open(infile) // declares both f and err
// ...
f, err := os.Create(outfile) // ERR: no new variables
```

## Pointers

A pointer value is the address of a variable. That is, a pointer is the location at which a value is stored.

With a pointer, we can read or update the value of a variable indirectly, without using or even knowing the name of the variable (if it has a name).

If a variable is declared as `var x int`, the expression `&x` yields a pointer to an integer variable. That will be a value of type `*int`, which is pronounced "pointer to int". If we call that value `p` we say "`p` points to `x`", or "`p` containes the address of `x`".

The variable to which `p` points to is written `*p`, but since `*p` denotes a variable, it may appear on the left-hand side of an assignment, in which case the assignment updates the variable:

```go
x := 1
p := &x
fmt.Println(*p) // 1
*p = 2
fmt.Println(*p) // 2
```

Each component of a variable of aggregate type (field of a struct, or element of an array) is also a variable and thus, has an address too.

The zero value for a pointer of any type is `nil`.

Pointers are comparable: two pointers are equal if and only if they point to the same variable, or both are nil.

It is perfectly safe for a function to return the address of a local variable.

```go
var p = f() // p is pointer to int

func f() *int {
  v := 1
  return &v
}
```

In the code above, the local variable `v` created by this particular call to `f` will remain in existence even after the call has returned, and the pointer `p` will still refer to it.

Because a pointer contains the address of a variable, passing a pointer argument to a function makes it possible for the function to update the variable that was indirectly passed.

For example, the following code defines a function that increments the variable that its argument points to, and returns the new value of the variable so it may be used in an expression:

```go
func incr(p *int) int {
  *p++
  return *p
}

v := 1
incr(&v)  // calling for side effects, now v = 2
fmt.Println(incr(&v)) // "3", and now v = 3
```

Each time you take the address of a variable or copy a pointer, we create new aliases or ways to identify the same variable.

Pointer aliasing is useful because it allows you to access a variable without using its name, but it is a double-edged sword, because to find all the statements that access a variable you have to find all its aliases.

Aliasing also occurs when you copy values of other reference types like slices, maps, and channels, and even structs, arrays, and interfaces that contain these types.


## The `new` function

The expression `new(T)`:
+ creates a new unnamed variable of type `T`
+ initializes it to the zero value of `T`
+ returns its address, that is, a value of type `*T`


| NOTE: |
| :---- |
| The `new` function is relatively rarely used, as the struct literal syntax is far more flexible. |

`new` is a predeclared function, not a keyword. As a result, you can redefine the name for something else as shown in the following example:

```go
func delta(old, new int) int {
  return new - old
}
```

In the `delta` function, the `new` function will be unavailable.

## Lifetime of Variables

+ The lifetime of a package-level variable is the entire execution of the program.
+ Local variables live until they become unreachable, at which point its storage may be *recycled*.
+ Function parameters and results are local variables that are created each time their enclosing function is called.
+ Vars declared in a loop `for i := ...` are created only once, when the loop begins.

Because of the rules discussed above, a local variable may outlive its enclosing function:

```go
func createInt() *int {
  var dummy
  return &dummy // perfectly safe: the pointer will outlive the function
}
```

In the example above, we say `dummy` *escapes* from the function.

| NOTE: |
| :---- |
| The notion of escaping is only relevant during low-level program optimization, since each variable that escapes requires an extra memory allocation (it is heap-allocated instead of stack allocated). |

Garbage collection is a tremendous help in writing correct programs, but it doesn't relieve you from thinking about memory. For example, keeping unnecessary pointers to short-lived objects within long-lived objects (such as global variables), will prevent the garbage collector from reclaiming the short-lived objects.

## Tuple Assignment

Go allows for several variables to be assigned at once, and this capability is known as tuple assignment:

```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```

Often functions in Go return multiple results. When such call is used in an assignment, the left-hand side must have as many variables as the function has results:

```go
f, err := os.Open("foo.txt")
```

Often, functions use these additional results to indicate some kind of error, either by returning an `error`, or a `bool` usually called `ok` when the nature of the error is not important, and only signalling that it has failed is relevant.

The latter approach is used in map lookups, type assertions, and channel receive operations:

```go
v, ok = m[key]  // ok indicates that the key was found
v, ok = x.(T)   // ok indicates the assertion worked
v, ok = <-ch    // ok indicates there's data to consume
```

## Type declarations

A `type` declaration defines a new *named* type that has the same underlying type as an existing type.

The named type provides a way to separate different and perhaps incompatible uses of the underlying type, so that they can't be mixed unintentionally:
  + an `int` might represent a loop variable, a timestamp, a month...
  + a `float64` might represent a temperature, a speed, ...
  + a `string` might represent a password, a color, ...

Type declarations often appear at package level, where the named type is visible throughout the package, and if the name is *exported*, it'll be accessible from other packages as well.

For example, you can define two named types for Celsius and Fahrenheit temperatures when dealing with a temperature conversion, as that will make the program more readable:

```go
type Celsius float64

type Fahrenheit float64

func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
```

When doing so, you will not be able to compare Celsius and Fahrenheit, as they're not the same type, even though both have the same underlying type.

Additionally, you will need to make an explicit conversion using the syntax below to convert to `float64` to the corresponding named type:

```go
Celsius(t)
Fahrenheit(t)
```

For every type `T`, there is a corresponding conversion operation `T(x)` that converts the value `x` to the type `T`. This conversion will succeed is both have the same underlying type, or if both are unnamed pointer types that point to variables of the same underlying type.

In specific ocasions you will be allowed to convert between numeric types and between string and some slice types. These conversions might change the internal representation (e.g., you can convert a `float64` to an `int` and the fractional part will be discarded).

You can use comparison operators such as `==` and `<` two compare two values of the same named type, or to a value of the underlying type:

```go
tempconv.BoilingC == 100 // OK
boilingF > freezingF     // OK
```

Named types also make it possible to define new behaviors for values of the type. These behaviors are expressed as a set of functions associated with the type, and are called the type's method.

```go
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}
```

## Packages and Files

| NOTE: |
| :---- |
| This section was written before Go modules existed. |

Packages in Go serve the same purpose as libraries in other languages, supporting modularity, encapsulation, separate compilation, and reuse.

The source code for a package resides in one or more `.go` files, usually in a directory whose name ends with the import path; for instance, the files of the `gopl.io/ch1/helloworld` package are stored in `$GOPATH/src/gopli.io/ch1/helloworld`.

To refer to a function outside its package, we must qualify the identifier to make explicit which one we meant to use; for example, whether we mean `image.Decode` or `utf16.Decode`.

To illustrate the basics, let's suppose that we want to make it available to the Go community.

Let's first create a new empty repository in Github git@github.com:sergiofgonzalez/tempconv.git, and then clone it:

```bash
git clone git@github.com:sergiofgonzalez/tempconv.git
```

Then create a new module:

```bash
go mod init github.com/sergiofgonzalez/tempconv
```

Then, give the code the proper structure and content and push it to Github. Go releases work with Git tags. Therefore, you should properly tag and push the version:

```bash
git push origin main
git tag v0.1.0
git push origin v0.1.0
```

At that point you can start using it in your code by simply doing:

```bash
go get github.com/sergiofgonzalez/tempconv
```

See [Using tempconv](06_packages_using/README.md)

## Imports

Every package is identified by a unique string called its *import path* as in `"github.com/sergiofgonzalez/tempconv"` or `"fmt"`.

The language specification doesn't define where these strings come from, or what they mean.

When using `go` tool, and import path denotes a directory containing one or more Go source files that together make up the package.

In addition to its import path, each package has a *package name*, which is the typically short, and not necessarily unique name that appears in its package declaration (as in `package tempconv`).

The import declaration binds a short name to the imported package, tht must be used to refer to its content:

```go
// ...
import (
	"fmt"

	"github.com/sergiofgonzalez/tempconv"
)

//...
fmt.Println(tempconv.CtoF(17))
```

## Package initialization

Package initialization begins by initializing package-level variables in the order in which they are declared, except that dependencies are resolved first:

```go
var a = b + c   // a initialized third, a = 3
var b = f()     // b initialized second, b = 2
var c = 1       // c initialized first, c = 1

func f() int { return c + 1 }
```

If the package has multiple `.go` files, they are initialized in the order in which the files are given to the compliler.

Any file may contain any number of functions whose declaration is just:

```go
func init() { /* ... */ }
```

Such `init` functions can't be called or referenced, but otherwise are normal functions that are automatically executed when the program starts, in the order in which they are declared, and that can be used to implement initialization for complex structures.

Packages are initialized one at a time, in the order of imports in the program, dependencies first, so that any package can rely on its dependencies being fully initialized. As the initializtion proceeds from the bottom up, the `main` package is the last being initialized, and all packages are fully initializes when the `main` function begins.

## Scope

A declaration associates a name with a program entity, such as a function or a variable. The *scope* of a declaration is the part of the source code where a use of the declared name refers to that declaration.

Scope is different from lifetime. The scope is a compile-time property, while the lifetime is a run-time property.

A syntactic block is a sequence of statements enclosed in braces, such as the statements in a function or loop.

A lexical block is the grouping of declarations no explicitly surrounded by braces. The lexical block for the entire source code of an application is called the *universe block*.

+ built-in types (e.g., `int`), functions (`len`, `make`), and constants (`true`) are in the *universe block*.
+ package-level declarations (not included within any function) can be referred to from any file in the same package.
+ imported packages (e.g., when you import `fmt` in your program), can be referred to from the same file, but not from another file in the same package.
+ Many declarations are local, so they can be referred to from within the same function, or maybe just from a part of the function.

When the compiler encounters a reference to a name, it looks for a declaration, starting with the innermost enclosing lexical block, working up to the universe block. If it doesn't find a declaration, the compiler reports an error:

```go
func f() {}

var g = "g"

func main() {
  f := "f"
  fmt.Println(f) // "f": local var shadows the package level func
  fmt.Println(g) // "g": package level var is found
  fmt.Println(h) // ERROR: undefined h
}
```

Some blocks are implicitly defined. Consider:

```go
if f, err := os.Open(fname); err != nil {
  return err  // ERR: unused f
}

f.ReadByte()  // ERROR: undefined f
f.Close()     // ERROR: undefined f
```

In the snippet above, the variable f is declared in a block that does not span beyond the `if`.

