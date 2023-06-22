# Formatting Verbs Cheat Sheet

| Format Verb | Description | Example |
| :---------- | :---------- | :------ |
| `"%b"` | Prints the binary value | `fmt.Printf("%b", flag)` // 1001 |
| `"%c"` | Prints the character or rune | `fmt.Printf("%c", flag)` // 𢉩 |
| `"%d"` | Prints the decimal value | `fmt.Printf("%d", flag)` // 9 |
| `"%e"` | Prints float with exponent | `fmt.Printf("%e", num)` // 3.14e-2 |
| `"%f"` | Prints float without exponent | `fmt.Printf("%f", num)` // 3.14159 |
| `"%8.3f"` | Prints float using 8 chars in total, using 3 digits after the decimal point. | `fmt.Printf("%f", num)` // 3.14159 |
| `"%g"` | Prints a float with the most compact representation | `fmt.Printf("%g", flag)` // 3.5e2 |
| `"%o"` | Prints the octal value | `fmt.Printf("%d", flag)` // 011 |
| `"%T"` | Prints the type of the variable | `fmt.Printf("%T", noDelay)` // time.Duration |
| `"%q"` | Print quoted string or char | `fmt.Printf("%q", s)` // "hello" |
| `"%v"` | Catch-all | `fmt.Printf("%v", err)` // ERR: file not found |
| `"%x"` | Prints the hex value | `fmt.Printf("%x", flag)` // de |
| `"%X"` | Prints the hex value | `fmt.Printf("%x", flag)` // DE |
| `"% X"` | Add a space before printing | `fmt.Printf("%x % x", flag1, flag2)` // DE F1 |
| `"%#x"`<br>`"%#o"` | Emit `0x` for hex, `0` for octals | `fmt.Printf("%#x", flag)` // 0xDE |
| `"%d %[1]c %q %[2]c"` | Reuse first and second vars | `fmt.Printf("%d\t%q\t%[2]d\t0x%[2]x\n", i, r)` // 1 'H' 72 0x48 |