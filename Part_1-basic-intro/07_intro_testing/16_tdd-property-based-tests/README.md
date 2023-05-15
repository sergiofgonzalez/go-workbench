# Property-based tests in Go
> building a conversor from arabic numbers to roman numerals using TDD

## v0: getting started

As we're firm believers in TDD, we'll start by writing the simplests of tests, that expects that the number `1` is correctly converted into `I`.

## v1: testing beyond I

In this version we need to enhance the tests and implementation to support more than `I`.

Start by creating a test for the conversion from `2` to `II`, and then create the necessary structures to apply table-driven tests.

After adding support for table-driven tests, add support for `III`.

## v2: refactoring

At this point we see that the current approach won't scale well as we have infinite numbers to support.

Start with the refactoring of the `ConvertToRoman` function by using a `strings.Builder` instead of a series of conditions.

With the refactoring in place, add a test for converting `4` into `IV` and `5` into `V`.

The algorithm for the conversion should look something like this:

```
  result = ""
  while n > 0:
    if n > 4:
      n -= 5
      result += "V'
    else if n > 3:
      n -= 4
      result += "IV"
    else
      n -= 1
      result += "I"
```

That algorithm should be OK to convert 6, 7 and 8 too!

## v3: going forward

Now write a test for "IX". The previous algorith won't work, but it has given you the structure to go forward. Add code to support "X" as well and validate it works up to 39 by probing individual values such as 10, 14, 18, 20, 30, and 39.

Confirm that it fails for 40. Before adding the code to make it work, it is time to do a bit of refactoring. Using a switch with so many cases is a code smell for a brute-force approach to something that should be handled in a cleaner way.

## v4: getting it cleaner

Let's start refactoring by creating a `RomanNumeral` struct that captures the relationship between an int value and its representation, and a slice with the specific values of interest: 1, 4, 5, 9, 10, 40, 50...

Then the algorithm should be changed for this one:

```go
func convertToRoman(n int) string {
	sb := strings.Builder{}

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			sb.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return sb.String()
}
```

Then add the values to support 40, and check that it works up to 50 by probing 40, 47, 49, 50.

### v5: adding the rest of symbols
After doing so, you should have a good idea around the other elements that should be added to `allRomanNumerals`.

Then test for 1984, 3999, 2014, and 790.

At this point you will have a fairly well-tested algorithm to convert Arabic numbers to their corresponding Roman Numerals.

### v6: supporting the conversion from Roman numerals to Arabic

In this version we'll add support to convert Roman numerals to arabic, starting with the test scenarios, which we won't rewrite, but instead base on the existing ones we wrote.

Then, create a naive implementation of the `ConvertToArabic` function that can work with the first test cases.

HINT: in the tests use tests[:3] to run only the first set of tests.

### v7: rewriting the conversion

Any naive implementation will fail as soon as you hit "IV".

In this step we rewrite the conversions with a pretty contrived algorithm that I just copied (nothing to do with the property-based tests).

### v8: adding property-based tests

There are a few rules we have used while working with Roman numerals:
+ There are no more than 3 consecutive symbols (e.g. 'III')
+ Only I, X, and C can be used as subtractors.
+ The result of `ConvertToArabic(ConvertToRoman(n))` should be `n`.

The test cases we've run so far are called *example-based tests*. Property based tests help you do this by throwing random data at your code and verifying certain rules in the domain, rather than the actual result.

This would look like this:

```go

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		t.Log("testing", arabic)
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("Property based test failed: ", err)
	}
}
```


