package countdown

import (
	"io"
	"os"
)

// Countdown manages to write 3, 2, 1, Go! with 1-second pause
func Countdown(w io.Writer) {
	io.WriteString(w, "3")
}

func main() {
	Countdown(os.Stdout)
}