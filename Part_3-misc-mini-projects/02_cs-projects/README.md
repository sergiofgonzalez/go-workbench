# Computer Science-y related projects
> A collection of CS related concepts and projects implemented in Go

## Algorithms & Abstract Data Types

> An abstract data type (ADT) is a collection of related data along with a set of operations you can apply over the data.

It is important to note that an ADT is independent of its actual implementation which can be done using a specific programming language, and within the language using certain features such as data structures, interfaces, native data types, etc.

Execution time is an important characteristic to take into account when assessing a particular algorithm.

When doing analysis, we typically rely in the number of times you execute a particular program instruction, taken into account the data set size you're feeding the program.

For example, you might want to estimate the execution how good an algorithm is then reading N numbers and sorting them.

Many times the execution time might not rely only in the execution time, and instead depend on the arrangement of the data.

That's why in many cases you'd want to consider:
+ Worst case: the scenario in which the data arrangement renders the longest execution time (e.g., N numbers fed to the sorting algorithm in descending order).
+ Best case: the scenario rendering the shortest execution time (e.g., N numbers already sorted)
+ Mean case/Most probably case: the scenarios that fall in the middle (e.g., N random numbers).

To estimate the complexity of an algorithm we use the $ O(N) $ notation. This notation gives us an order of magnitude so that:

$$
O(1) \lt O(log N) \lt O(N) \lt O(N \cdot log N) \lt O(N^2) \lt O(N^3) \lt ... \lt O(2^N) \lt O(N!)
$$

Where:
+ $ O(1) $ is constant complexity
+ $ O(log N) $ is logarithmic complexity
+ $ O(N) $ is linear complexity
+ $ O(N^2) $ is quadratic complexity
+ $ O(2^N) $ is exponential complexity
+ $ O(N!) $ is factorial complexity

Note that $ O(N) $ gives us an order of magnitude, rather than a exact estimate.

A simple statement has a constant complexity $ O(1) $:

```go
x := x + 5
```

But a set of simple statements also have constant complexity $ O(1) $:

```go
x := x + 5
y := x / 2
x := y * 3
```

A loop iterating n times over a single statement has $ O(n) $ complexity:

```go
for i := 0; i < n; i++ {
  x := x + i
}
```

And a loop iterating n times over multiple statements also has $ O(n) $:

```go
for i := 0; i < n; i++ {
  x := x + i
  y := x / 2
  x := y * 3
}
```

A program with an outer and an inner loop has $ O(n^2) $ complexity:

```go
for i := 0; i < n; i++ {
  for j := 0; j < n; j++ {
    x := x + i + j
  }
}
```

### Example: Compute the complexity of the following block

```go
const n = ...

type Vector [n]int

func BubbleSort(v *Vector) {
  for i:= 0; i < n; i++ {
    for j := n - 1; j > i; j-- {
      if v[j - 1] > v[j] {
        v[j - 1], v[j] = v[j], v[j - 1]
      }
    }
  }
}
```

The size of the input is `n`. The inner loop has complexity $ O(1) $ and it is executed `n - i` times. The outer loop is executed `n` times.

That is, the inner statements are executed:

$$
(n - 1) + (n - 2) + ... + 1 = \frac{n \cdot (n - 1)}{2} = O(n^2)
$$

### Complexity calculation cheatsheet

1. Simple statement cost is $ O(1) $

    Therefore, assignments, arithmetic operations, read and write operations... on simple data.

    Slice/Array operations shouldn't be considered.

2. In a sequence of statements, the complexity of the sequence would be the maximun among the complexities of all the statements.

3. In a condition operation, the associated complexity is the complexity of the most complicated path between the different options.

4. In loops, compute the complexity as the sum of the complexity of the body + complexity of the termination condition:

$$
\sum_{num\_iterations}(complexity(body) + complexity(loop\_termination\_condition))
$$

5. When invoking functions or methods, you need to trace the sequence of calls, computing the complexity of each of its statements (recursively if required)

6. Recall the formulas:

    $
    \displaystyle \sum_{i = 0}^{n - 1} 1 = n
    $

    $
    \displaystyle \sum_{i = 0}^{n - 1} i = \frac{(n - 1) \cdot n}{2}
    $

    $
    \displaystyle \sum_{i = 0}^{n - 1} i^2 = \frac{(n - 2) \cdot n \cdot  (2n - 1)}{6}
    $

    $
    \displaystyle \sum_{i = 0}^{n - 1} i^k = O(n^{k+1})
    $

6. If a part of program given an input of size $ n $, is executed for $ n/2 $, $ n/4 $, $ n/8 $... then the complexity is $ O(log n) $.


## Lists: sequential representation

## Lists

A list is a collection of elements of the same type and arranged in a certain order.

| NOTE: |
| :---- |
| Order in a list means that for each element in the list there is a previous one (except for the first one), and there is always a next element (except for the last one).<br>That is, each element has a position within the list, which might not necessarily be based on the element's value. |

## Stacks

A stack is a collection of elements of the same type in which the order of the elements is decided based on when they have been inserted into it.

The top of the stack is the element that has been inserted in the last place.

This structure is known as last-in, first-out or LIFO, because the last inserted element is the first one being retrieved.