# Mini-project: Grokking pointers vs. values

Go has pointers, but in many instances, you don't feel you're working with pointers.

Because of that, it is important to understand the differences and when to use them to achieve the desired effects.

## Functions

### Using primitives: numbers, booleans, strings, and runes

Numbers, booleans, strings, and runes are passed by value which means that you can change them within the invoked function but those changes won't be visible to the caller once the call has completed.

As a result, if you want to change the value of an argument, make sure that the function receives a pointer, and not a value:

```go
func doubleMe(num *int) {
  *num *= 2
}

myNum := 5
doubleMe(&num)
fmt.Println(num) // 10
```

| NOTE: |
| :---- |
| In most of the cases it is preferable to make the function return the changed value. |

### Using arrays

Arrays in Go are not very common, and you will end up using slices more as they're more flexible.

However, it is important to know that when passing arrays to functions they are passed by value as if they were primitives. This means that an actual copy of the array is done. As a result, you cannot change the elements of the array, and you cannot reassign it to a different array.

If you need to change the elements of an array, you need to pass it by reference. That way, you'll get a copy of the address of the array, and you will be able to change its elements.

| NOTE: |
| :---- |
| You will not be able to reassign the passed array to a new one even if passing an address. This happens because what you get is a copy of the pointer and therefore, the caller will get the original address when the invoked function has complete. |

```go
func doubleElems(arr *[3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] *= 2
	}
}

arr := [3]int{0, 1, 2}
doubleElems(&arr)
fmt.Println(arr) // 0, 2, 4
```

| NOTE: |
| :---- |
| Note that within the function, you don't need to dereference the argument and can use the pointer directly.<br>Note also that if an array is big, you might want to pass it by reference to prevent the copy. |

### Using structs

The same that applies to arrays applies to structs: they are copied when passed as arguments to functions. As a result, if you need to change its contents you must pass a reference to the struct.

```go
func doublePersonAge(p *person) {
	p.age *= 2
}

p := person{"Sergio", 49}
doublePersonAge(&p)
fmt.Println(p.age)
```

The same ideas that apply regarding syntax and performance (copy vs. reference) apply here too.

### Maps and slices

Maps and slices are a bit different, because they are complex data structures that hide some underlying storage.

As a result, when passing maps and slices by either value or reference you will get write access to the underlying storage.

Note that when passing the reference of a map or slice the received argument cannot be reassigned though, because when doing so, you'll get a copy of the address that will be reverted back to its original value once the function has completed its execution.

```go
func capitalizeKeys(m map[string]int) {
	for k, v := range m {
		m[strings.ToUpper(k)] = v
	}
}

m := map[string]int{"Spain": 0, "France": 1}
capitalizeKeys(m)
fmt.Println(m)  // map[FRANCE:1 France:1 SPAIN:0 Spain:0]
```

```go
func doubleSliceElems(s []int) {
	for i := 0; i < len(s); i++ {
		s[i] *= 2
	}
}

s := []int{0, 1, 2, 3, 4}
doubleSliceElems(s)
fmt.Println(s)  // 0, 2, 4, 6, 8
```

## Methods

In methods, the compiler will automatically take the address, or dereference a pointer as required, so you can have:

```go
func (s someStruct) valueReceiver() {
  ...
}

func (s *someStruct) ptrReceiver() {
  ...
}

myStruct := someStruct{}
myStructPtr := &myStruct

myStruct.valueReceiver()      // OK
myStructPtr.valueReceiver()   // Compiler automatically dereferences the pointer

myStruct.ptrReceiver()        // Compiler automatically takes the address of the value
myStructPtr.ptrReceiver()     // OK
```

With respect to changing the values, the same that applies to the functions apply to the methods with respect to the receiver.

A common technique you'll find in methods with pointer receivers is the following:

```go
func (p *person) changePerson(other person) {
	*p = other
}

p1 := person{
	name: "Jason Isaacs",
	age: 57,
}
p1.changePerson(person{"Idris Elba", 49})
```

Note how in the function implementation, we just make `*p`, that is, the contents of the memory where `p` is pointing to have different values by simply copying the struct.

The same technique can be used for slices.