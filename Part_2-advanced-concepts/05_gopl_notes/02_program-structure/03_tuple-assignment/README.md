# gcd
> calculating the greatest common divisor using tuple assignment

## v0

Write a function that computes the greatest common divisor of two integers using the following algorithm:

Given two integers x, y:
  + get the remainder of x / y
    + if the remainder of x / y is zero, the gcd is x
    + otherwise readjust:
      + x = y
      + y = x / y
  