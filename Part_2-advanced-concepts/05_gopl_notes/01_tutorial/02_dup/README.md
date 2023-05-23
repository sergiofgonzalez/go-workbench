# dup: looking for adjacent duplicate lines

Create a program named `dup` which looks duplicate lines.

The program should print the text of each line that appears more than once, preceded by its count.

The program should be able to find duplicates from the standard input, or from a list of files received via command-line arguments.

## v0: finding duplicates on stdin

Write a program that outputs in stdout the duplicate lines that it finds in its input. That is, the program should keep track of the line it has seen in stdin until the input is closed, and then print the lines that has been repeated more than once.

For an input like:

```
line 1
this line will be repeated
line 2
this line will be repeated
something
something
something
```

The output should look like:

```
2     this line will be repeated
3     something
```

Make the implementation using TDD and testable.

HINT: You should rely on creating a Scanner, then calling `Scan` to fetch the lines, and then calling `Text` to get the contents of the retrieved line as a string.

```go
input := bufio.NewScanner(os.Stdin)
for input.Scan() {
  line := input.Text()
  // ... do something with line ...
}
```

HINT(2): to test the CLI tool, run the main program, type a few lines, and then hit CTRL-D to end the input.

## v1: handling files

Enhance the existing program so that it can either work with input received from the standard input, or alternatively, receive a list of one or more files from the command line and scan them.

## v2: an alternative implementation

Change the existing implementation so that instead of scanning the files using a scanner (and therfore use an streaming approach) we read the entire file in memory, split it into newlines, and then perform the count.

For simplicity, this solution is implemented into `main` without requiring any supporting functions.