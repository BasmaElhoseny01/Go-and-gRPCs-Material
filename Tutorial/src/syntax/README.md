# Go Basic Syntax

This file is taken from https://go.dev/tour, visit for more info.

### Data Types

```Go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
```

### Zero Values

- 0 for numeric types,
- false for the boolean type, and
- "" (the empty string) for strings.

### Type Conversion

```go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

Or, put more simply:

```go
i := 42
f := float64(i)
u := uint(f)
```

### Type Inference

When declaring a variable without specifying an explicit type (either by using the `:=` syntax or `var = expression` syntax), the variable's type is inferred from the value on the right hand side.

```go
var i int
j := i // j is an int
```

### Constants

- Constants are declared like variables, but with the const keyword.
- Constants can be character, string, boolean, or numeric values.
- Constants cannot be declared using the := syntax.

```go
const Pi = 3.14
```

---

## Functions

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

```go
func add(x, y int) int {
	return x + y
}
```

### Multiple results

A function can return any number of results.

```go
func swap(x, y string) (string, string) {
    return y, x
}
func main() {
    a, b := swap("hello", "world")
    fmt.Println(a, b)
}
```

---

## For Loop

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

- the init statement: executed before the first iteration
- the condition expression: evaluated before every iteration
- the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in the scope of the for statement.

The loop will stop iterating once the boolean condition evaluates to false.

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components of the for statement and the braces { } are always required.

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

The init and post statements are optional.

```go
sum := 1
for ; sum < 1000; {
    sum += sum
}
```

At that point you can drop the semicolons: C's `while` is spelled `for` in Go.

```go
sum := 1
for sum < 1000 {
    sum += sum
}
```

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.

```go
for {
}
```

---

## If

Go's `if` statements are like its `for` loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

```go
func sqrt(x float64) string {
    if x < 0 {
        return sqrt(-x) + "i"
    }
    return fmt.Sprint(math.Sqrt(x))
}
```

Like `for`, the `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.

```go
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}
```

---

## Defer

A defer statement defers the execution of a function until the surrounding function returns.

The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

```go
func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}
// output:
// hello
// world
```

### Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

```go
func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}
// output:
// counting
// done
// 9
// 8
// 7
// 6
// 5
// 4
// 3
// 2
// 1
// 0
```

## Pointers

Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

```go
var p *int
```

The & operator generates a pointer to its operand.

```go
i := 42
p = &i
```

The \* operator denotes the pointer's underlying value.

```go
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

This is known as "dereferencing" or "indirecting".

Unlike C, Go has no pointer arithmetic.

## Structs

A struct is a collection of fields.

```go
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```
## Arrays
The type `[n]T` is an array of `n` values of type `T`.

The expression
```go
var a [10]int
```
declares a variable a as an array of ten integers.

An array's length is part of its type, so arrays cannot be resized. This seems limiting, but don't worry; Go provides a convenient way of working with arrays.
## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the func keyword and the method name.

In this example, the `Abs` method has a receiver of type `Vertex` named `v`.

```go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```
