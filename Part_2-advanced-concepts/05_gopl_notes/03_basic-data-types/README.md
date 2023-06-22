# 03 &mdash; Basic Data Types

## Intro

Go's types fall into four categories:
+ basic types
  + numbers
  + strings
  + booleans
+ aggregate types
  + arrays
  + structs
+ reference types
  + pointers
  + slices
  + maps
  + functions
  + channels
+ interface types

Basic types is the topic of this chapter.

## Notes on Integers

+ `int` and `uint` are the natural and most efficient size for signed and unsigned integers on a particular platform.
+ `rune` is a synonym for `int32` and conventionally indicates that a value is a Unicode code point.
+ `byte` is a synonym for `uint8`, and it is typically used to indicate a piece of raw data (rather than a small numeric quantity).
+ `uintptr` is an unsigned integer type defined to hold all the bits of a pointer value.

Binary operators for arithmetic logic and comparison, in order of decreasing precedence are:

```
*    /    %    <<    >>     &    &^
+    -    |    ^
==   !=   <    <=    >      >=
&&
||
```

It the result of an arithmetic operation has more bits than can be represented in the result type, it is said to overflow. The high-order bits that do not fit are silently discarded.

```go
var u uint8 = 255
fmt.Println(u, u+1, u*u) // 255 0 1

var i int8 = 127
fmt.Println(i, i+1, i*i) // 127 -128 1
```

Two integers of the saame type may be compared using the well-knwon binary comparison operators. The type of such comparison is a boolean.

As a reminder, Go supports the following bitwise binary operators, with the first four treat their operands as bit patterns with no concept of arithmetic carry or sign:

```
&   bitwise AND
|   bitwise OR
^   bitwise XOR
&^  bit clear (AND NOT)
<<  left shift
>>  right shift
```

The `^` operator is the bitwise negation when used as a unary operator (and XOR when used as a binary operator).

The left shift and right shift are understood as follows:

```go
byte << numberOfBitsPositionsToShiftLeft
byte >> numberOfBitsPositionsToShiftRight
```

Left shifts fill the vacated bits with zeros, as do right shifts of unsigned numbers.

| EXAMPLE: |
| :------- |
| See [01_int-bitwise](01_int-bitwise/README.md) for a runnable example. |


However, right shifts of signed numbers fill the vacated bits with copies of the sign bit.

It's common to use signed `int` form for quantities that can't be negative, such as the length of an array, though `uint` should be more appropriate. For example, the `len` built-in function returns an int. This is needed for loop conditions in which you compare lengths with a number that is less than zero.

| NOTE: |
| :---- |
| Unsigned numbers tend to be used only when using bitwise operators, or when implementing bit sets, parsing binary data, or using cryptography. |

Type mismatches with ints are typically fixed by converting everything to a common type:

```go
var apples int32 = 1
var oranges int16 = 2
// var mix = apples + oranges // Compile-time error
var mix = int(apples) + int(oranges)
```

Float to integer conversion discards any fractional part, truncating towards zero.

```go
f := 3.141
i := int(f) // 3

f = 1.99
i = int(f) // 1
```

Integer literals can be written as ordinary decimals, octal numbers (if they begin with 0, as in 0666), or as hexadecimal (if they begin with 0x or 0X, as ix 0xde0b).

We can control the prinint of these numbers with the formats:

+ `%d` &mdash; decimal
+ `%o` &mdash; octal
+ `%x` &mdash; hex
+ `%b` &mdash; binary

Consider the following program that uses a few interesting `fmt` tricks:

```go
o := 0666 // octal, for POSIX file-permissions
fmt.Printf("%d %[1]o %#[1]o\n", o)

x := int64(0xdeadbeef)
fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
```

The `[1]` adverbs let us reuse a given argument multiple times in the format specifier. The `#` adverb tells `fmt` to emit `0` for octals, and `0x` or `0X` for hexadecimals.

| EXAMPLE: |
| :------- |
| See [02_int-fmtprnt](02_int-fmtprnt/README.md) for a runnable example. |

Rune literals are written as a character within single quotes. They are used to hold a single Unicode code point.

They are printed using `%c` or `%q` if quoting is desired.

```go
	ascii := 'a'
	unicode := '𢉩'
	newline := '\n'

	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]q\n", newline) // no %c to prevent newline
```

| EXAMPLE: |
| :------- |
| See [03_int-runprnt](03_int-runeprnt/README.md) for a runnable example. |

## Notes on Floating-Point Numbers

The limits for both `float32` and `float64` can be found in the `math` package.

A `float32` provides approximately 6 decimal digits of precision, while `float64` provides about 15 digits.

Very small or very large numbers should be written in scientific notation with the letter `e` or `E` preceding the decimal exponent:

```go
const Planck = 6.62606957e-34
```

Floating point numbers are conveniently printed with `Printf`'s `%g` verb which chooses the most compact representation. For tables of data you can also use `%e` to always use exponent, or `%f` to never use exponent.
All three verbs allow for for field width and numeric precision representation as is `%8.3f` (take up 8 characters, use 3 digits after the decimal point).

The `math` package also defines `NaN`, `+Inf`, and `-Inf`.

```go
var z float64
fmt.Println(z, -z, 1/z, -1/z, z/z) // 0 -0 +Inf -Inf NaN
```

You can use `math.IsNan()` to check whether an argument is "not-a-number". Any comparison with the result of the function `math.NaN` renders false.

If a function returning a floating point is likely to report an error it is recommended to use the following approach:

```go
func compute() (value float64, ok bool) {
  // ... computations that might fail ...
  if failed {
    return 0, false
  }
  return result, true
}
```

| EXAMPLE: |
| :------- |
| See [04_floats-surface-plot](04_floats-surface-plot/README.md) for a runnable example on floats computations. |

## Notes on Complex Numbers

Go provides `complex64` and `complex128` for complex numbers, whose components are `float32` and `float64` respectively.

```go
var x complex128 = complex(1, 2) // 1+2i
fmt.Println(real(x)) // 1
fmt.Println(imag(y)) // 2
```

If a floating-point literal, or decimal integer literal is immediately followed by `i` such as `2i` it becomes and imaginary literal.

You can write complex numbers and operations naturally in Go:

```go
x := 1 + 2i
y := 3 + 4i
fmt.Println(x == y) // false
fmt.Println(x * y)  // -5 + 10i
```

The `math/cmplx` package provides library functions for working with complex numbers, such as square root and exponentiation functions.

```go
fmt.Println(cmplx.Sqrt(-1)) // i
```

| EXAMPLE: |
| :------- |
| See [05_complex-mandelbrot](05_complex-mandelbrot/README.md) for an example involving complex numbers. |

## Notes on Booleans

Boolean values can be combined with `&&` and `||`, which have short-circuit behavior: if the answer is already determined by the value of the left operand, the right operand is not evaluated, which makes safe to right expressions such as:

```go
if s != "" && s[0] == 'x' { ... } // doesn't panic if s is empty
```

`&&` has higher precedence than `||` which lets you use no parentheses in expressions such as:

```go
if 'a' <= c && c <= 'z' ||
   'A' <= c && c <= 'Z' ||
   '0' <= c && c <= '9' {
    // ... ASCII alphanumeric char ...
   }
```

There is no implicit conversion from a boolean to a numeric value like `0` or `1`.

## Notes on Strings

A string is an immutable sequence of bytes, which usually contains human-readable text, but may contain arbitrary dta.

Text strings are interpreted as UTF-8 encoded sequences of Unicode code points, known as *runes* in Go.

Getting the length in bytes (not runes) and accessing the individual bytes looks like this:

```go
s := "Hello, world!"
fmt.Println(len(s))       // 13
fmt.Println(s[0], s[7])
```

Attempting to access a byte outside this range results in a panic:

```go
c := s[len(s)] // panic: index out of range!
```

The substring operation `s[i:j]` yields a new string consisting of the bytes of the original string starting at index `i` and continuing up to, but not including the byte at index `j`, so that the result contains `j-i` bytes.

```go
s := "Hello, world!"
fmt.Println(s[0:5]) // "Hello"
```

Either or both of the `i` and `j` operands may be omitted:
+ `i` defaults to `0`, the beginning of the string
+ `j` defaults to `len(s)`, the end of the string

```go
s := "Hello, world!"
fmt.Println(s[:5]) // "Hello"
fmt.Println(s[7:]) // "world!"
fmt.Println(s[:]) // "Hello, world!"
```

The `+` operator makes a new string by concatenating two strings:

```go
s := "Hello, world!"
fmt.Println("goodbye" + s[:5]) // "goodbye world!"
```

Strings may be compared with comparison operators like `==` and `<`. This comparison is done byte-by-byte.

Strings are immutable, which means that the byte sequence contained in a string value can never be changed. You can though assign a new value to a string variable:

```go
s := "left foot"
s += ", right foot"
s[0] = "L" // compile-time error: cannot assign to s[0] (immutable!)
```

### Notes on String literals

A string value can be written as a string literal, a sequence of bytes encloses in double quotes. Unicode code points can be used in string literals:

```go
s := "Hello, 世界"
```

Within a double-quoted string literal, escape sequences that begin with a backslash can be used to insert arbitrary byte values into the string, such as `\n`, `\t`, `\'`, `\\`...

Arbitrary bytes can also be included using hexadecimal or octal escapes as in `\xhh` or `\od`.

A raw string literal is written with backquotes as in `...`. Within a raw string literal, no escape sequences are processed; the contents are taken literally including backslashes and newlines.

These are especially useful to write regular expressions, and HTML and JSON literals.

```go
const GoUsage = `Go is a tool for managing Go source code.
Usage:
    go command [args]
    `
```

### Notes on Unicode

The natural data type to hold a single Unicode code point, which in Go is called a *rune* is an `int32`.

The `int32` type has a synonym `rune` to represent Unicode code points.

While this is simple and uniform, is a huge waste of space as many Unicode code points only utilize 8 bits.

UTF-8 is a variable length encoding of Unicode code points as bytes. It uses between 1 and 4 bytes to represent each *rune*.

```
0xxxxxxx runes              0-127     (7-bit ASCII)
11xxxxxx 10xxxxxx           128-2047  (values <128 unused)
110xxxxx 10xxxxxx 10xxxxxx  2048-65535 (values < 2048 unused)
1110xxxx 10xxxxxx 10xxxxxx 10xxxxxx 65536-0x10FFFF (other values unused)
```

Go source files are always encoded in UTF-8, and UTF-8 is the preferred encoding for text strings manipulated by Go programs.

The `unicode` package provides functions for working with individual runes (such as distinguishing letters from numbers, converting to uppercase/lowercase), and the `unicode/utf8` package provides functions for encoding and decoding runes as bytes using UTF-8.

Oftentimes you will need to use Unicode escapes to in Go string literals. These let you specify Unicode code points by their numeric code point value.

There are two forms (each `h` is a hex digit):
+ `\uhhhh` for a 16 bit value.
+ `\Uhhhhhhhh` for a 32 bit value.


The following string literals represent the same six-byte string:

```go
"世界"
"\xe4\xb8\x96\xe7\x95\x8c"  // using byte encoding
"\u4e16\u754c"              // using unicode (16-bit values)
"\U00004e16\U0000754c"
```

Unicode escapes may also be used in rune literals:

```go
'界'
'\u4e16'
'\U00004e16'
```

A rune whose value is less than 256 maybe written with a single hexadecimal escape, such as `\x41` for `'A'`, but for higher values, `\u` or `\U` must be used. Thus, `'\xe4\xb8\x96'` is not considered a valid rune literal.

Many string operations that check individual runes do not require decoding:

```go

// HasPrefix returns true if s starts with prefix, false otherwise
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// HasSuffix returns true if s ends with prefix, false otherwise
func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

// Contains returns true if s contains substr in some part of it, false otherwise
func Contains(s, substr string) bool {
	for i := 0; i < len(s) + 1; i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
```

| EXAMPLE: |
| :------- |
| See [06_strings-stringfns](06_strings-stringfns/README.md) for a runnable example. |


However, if we really care abouth the individual Unicode characters, we will need to rely on the `unicode` package.

```go
s := "世界"
fmt.Println(s)                         // 世界
fmt.Println(len(s))                    // 6 bytes
fmt.Println(utf8.RuneCountInString(s)) // 2 chars

// To process the characters individually we need to decode them
// 0       世
// 3       界
for i := 0; i < len(s); {
  r, size := utf8.DecodeRuneInString(s[i:])
  fmt.Printf("%d\t%c\n", i, r)
  i += size
}

// We can also use a simple range s to get the runes
// 世
// 界
for _, r := range s {
  fmt.Printf("%c\n", r)
}
```

We can get the code points of a given rune with a simple loop:

```go
// 0       'H'     72      0x48
// 1       'e'     101     0x65
// 2       'l'     108     0x6c
// 3       'l'     108     0x6c
// 4       'o'     111     0x6f
// 5       ','     44      0x2c
// 6       ' '     32      0x20
// 7       '世'    19990   0x4e16
// 10      '界'    30028   0x754c
for i, r := range "Hello, 世界" {
  fmt.Printf("%d\t%q\t%[2]d\t0x%[2]x\n", i, r)
}
```

And with another loop we can easily get the number of runes in a string:

```go
n := 0
s := "Hello, 世界"
for range s {
  n++
}
fmt.Printf("runes: %d, bytes: %d\n", n, len(s)) // runes: 9, bytes: 13
```

| EXAMPLE: |
| :------- |
| See [06_strings-stringfns](06_strings-stringfns/README.md) for a runnable example. |

UTF-8 is very convenient as an interchange format, but oftentimes, `rune` (alias of `int32`) is more convenient within a program they are uniform in size and therefore easily indexed in arrays and slices.

A `[]rune` conversion applied to a UTF-8 encoded string returns the sequence of Unicode code points that the string encodes:

```go
s := "プログラム"
fmt.Printf("% x\n", s)  // spc is intentional (another formatting trick!)

// Transform the utf-8 string in an array of runes
r := []rune(s)
fmt.Printf("%x\n", r)

// Transform an array/slice of runes into an utf-8 encoded string
s = string(r)
fmt.Println(s)

// Converting an integer value to a string
fmt.Println(string(65)) // "A", not "65"

// A rune is already an uint32, so no conversion needed
c := 'A'
fmt.Println(c)  // 65

// A string with one char do needs conversion, but it is implicit
s = "A"
fmt.Println(s[0])  // 65

// When an invalid rune is used the replacement character is used
fmt.Println(string(1234567))  // ERR: �
fmt.Println(string(0x4eac))		// OK:  京
```

An unexpected input byte generates a special Unicode replacement character with value `\uFFFD` which is usually printed as �.

| EXAMPLE: |
| :------- |
| See [08_strings-runes](08_strings-runes/README.md) for a runnable example of the concepts of this section. |

### Strings and Byte Slices

The following four packages are relevant for manipulating strings:
+ `strings` &mdash; functions for searching, replacing, comparing, trimming, splitting and joining strings. Also functions like `ToUpper` and `ToLower` that return a new string with the specified transformation applied.
+ `bytes` &mdash; similar functions for manipulating slices of bytes (`[]byte`). Because strings are immutable, many times you'll use `bytes.Buffer` when you need to build up strings.
+ `strconv` &mdash; functions for converting boolean, integer, and floating-point values to and form their string representations, and functions for quoting and unquoting strings.
+ `unicode` &mdash; functions like `IsDigit`, `IsLetter`, `IsUpper`, and `IsLower` for classifying runes. Also, functions like `ToUpper` and `ToLower` to convert a rune into the given case.

For example, consider this implementation of the `basename` utility that relies on no libraries to perform its duties:

+ `basename("a/b/c.go")` -> `"c"`
+ `basename("c.d.go")` -> `"c.d"`
+ `basename("abc")` -> `"abc"`

```go
func basename(s string) string {
	// Trim everything until the last '/' if any
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// Trim everything after the last '.' if any
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
```

This implementation can be greatly simplified using `strings.LastIndex`:

```go
func basename(s string) string {
	// Trim everything until the last '/' if any
	s = s[strings.LastIndex(s, "/")+1:]

	// Trim everything after the last '.' if any
	if posDot := strings.LastIndex(s, "."); posDot >= 0 {
		s = s[:posDot]
	}

	return s
}
```

| EXAMPLE: |
| :------- |
| See [09_strings_and_bytes-basename](09_strings_and_bytes-basename/README.md) for a runnable example.<br>See also [10_strings_and_bytes-int-comma-separator](10_strings_and_bytes-int-comma-separator/README.md) for another example. |

While a string contains an array of bytes that once created is immutable, the elements of a byte slice can be freely modified and those can be converted back and forth very easily:

```go
s := "abc"
b := []byte(s)	// allocates a new byte array holding a copy of the bytes of s
s2 := string(b)	// makes a copy to ensure immutability of s2
```

To avoid some of the unnecessary copying and conversions, many of the utility functions in the `bytes` package and the `strings` package are the same in nature:

| strings | bytes |
| :------ | :---- |
| `func Contains(s, substr string) bool` | `func Contains(b, subslice []byte) bool` |
| `func Count(s, sep string) int` | `func Count(s, sep byte[]) int` |
| `func Fields(s string) []string` | `func Fields(s []byte) [][]byte` |
| `func HasPrefix(s, prefix string) bool` | `func HasPrefix(s, prefix []byte) bool` |
| `func Index(s, sep string) int` | `func Index(s, sep []byte) int` |
| `func Join(a []string, sep string) string` | `func Join(s [][]byte, sep []byte) []byte` |

The `bytes` package provides the `Buffer` type for efficient manipulation of byte slices.
A `Buffer` starts out empty but grows as data of types like `string`, `byte`, and `[]byte` are written to it.

Its zero value is usable, so it require no initialization.

The following snippet illustrates an example:

```go
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
```

When appending the UTF-8 encoding of an arbitrary rune, `buf.WriteRune` will work best, but when appending ASCII characters such as `[` you can safely use `buf.WriteByte` as seen in the snippet above.

Note also its versatility, in the `fmt.Fprintf` statement, in which the buffer is used as if it were a file.

| EXAMPLE: |
| :------- |
| See [11_strings_and_bytes-buffer-printints](11_strings_and_bytes-buffer-printints/README.md) for a runnable example. |

### Conversions between Strings and Numbers

Conversions between numeric values and their string representations are done with functions from the `strconv` package.

Alternatively, you can do conversions from numbers to strings using `fmt.Sprintf`.

To convert an integer to a string you can use `strconve.Itoa` (integer to ASCII):

```go
x := 123
y := strconv.Itoa(x)  		// "123"
z := fmt.Sprintf("%d", x) // "123"
```

`FormatInt` and `FormatUint` can be used to format numbers in a different base:

```go
fmt.Println(strconv.FormatInt(Int64(x), 2)) // 1111011
```

But in general, using Sprintf is more convenient:

```go
s := fmt.Sprintf("x=%b", x)) // x=1111011
```

To parse a string representing an integer, you can use `Atoi`,`ParseInt`, or `ParseUint`:

```go
x, err := strconv.Atoi("123")
y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
```

When using `ParseInt`, passing a 64 assumes the result is an `int64`, a 16 an `int16`, and a 0 an `int`.

Conversely to what happens with `fmt.Sprintf` (which is more flexible than `Itoa`), `fmt.Scanf` is not recommended to convert strings to numbers.

## Constants

Constants are expressions whose value is known to the compiler, so that their evaluation is guaranteed to occur at compile time.

The underlying type of every constant is:
+ boolean
+ string
+ number

Since their values are known at compile time, constants can be used in type definitions:

```go
const IPv4Len = 4

var p [IPv4Len]byte
```

A constant declaration may specify a type, otherwise, the type is inferred from the expression on the right-hand side:

```go
const noDelay time.Duration = 0 // underlying type is int64
```

When a sequence of constants is declared as a group, the right-hand side expression may be omitted for all but the first of the group, implying that the previous expression and its type should be used again:

```go
const (
	a = 1			// 1
	b					// 1
	c = 2			// 2
	d					// 2
)
```

A `const` declaration may use the *constant generator* `iota`, which is used to create a sequence of related values without spelling out each one explicitly.

The value of `iota` begins at zero and increments by one for each item in the sequence:

```go
type Weekday int

const (
	Sunday Weekday = iota		// 0
	Monday									// 1
	Tuesday									// 2
	Wednesday								// 3
	Thursday								// 4
	Friday									// 5
	Saturday								// 6
)
```

This can be used in more complex expressions too:

```go
type Flags uint

const (
	One Flags = 1 << iota	// 0001 1
	Two										// 0010 2
	Four									// 0100 4
	Eight									// 1000 8
)
```

### Untyped Constants

Although a constant can have any of the basic data types like `int` or `float64`, or types with an underlying basic type like `time.Duration`, many constants are not committed to a particular type.

In these case the compiler represents them with much greater precision than the values of basic types, and arithmetic on them is more precise than machine arithmetic; you may assume at least 256 bits of precision.

Consider the following definitions:

```go
const (
	_ = 1 << (10 * iota)		// 1
	KiB									    // 2^10 = 1024
	MiB
	GiB
	TiB											// exceeds 1 << 32
	PiB
	EiB
	ZiB											// exceeds 1 << 64
	YiB
)
```

Untyped constants not only retain their higher precision until later, but they can also participate in expressions such as:

```go
fmt.Println(YiB/ZiB)	// 1024
```

Additionally, an uncommitted constant can participate in expressions without requiring explicit conversions.

Only constants can be untyped. Also, once assigned to a variable, or appears on the right-hand side of a variable declaration with an explicit type, the constant is implicitly converted to the type of that variable if possible.

```go
const Pi64 float64 = math.Pi	// no longer an untype constant
```