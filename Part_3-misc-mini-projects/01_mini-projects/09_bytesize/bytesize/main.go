package main

import "fmt"

// ByteSize represents a storage size in bytes
type ByteSize float64

const (
	_ = iota													// 0
	KB ByteSize = 1 << (10 * iota)		// 2^10 ~ 1024
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)


func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%.2fYB", b / YB)
	case b >= ZB:
		return fmt.Sprintf("%.2fZB", b / ZB)
	case b >= EB:
		return fmt.Sprintf("%.2fEB", b / EB)
	case b >= PB:
		return fmt.Sprintf("%.2fPB", b / PB)
	case b >= TB:
		return fmt.Sprintf("%.2fTB", b / TB)
	case b >= GB:
		return fmt.Sprintf("%.2fGB", b / GB)
	case b >= MB:
		return fmt.Sprintf("%.2fMB", b / MB)
	case b >= KB:
		return fmt.Sprintf("%.2fKB", b / KB)
	}

	return fmt.Sprintf("%.2fB", b)
}


func main() {

	fmt.Println("Table of values")
	fmt.Printf("1 KB = %v Bytes\n", float64(KB))
	fmt.Printf("1 MB = %v Bytes\n", float64(MB))
	fmt.Printf("1 GB = %v Bytes\n", float64(GB))
	fmt.Printf("1 TB = %v Bytes\n", float64(TB))
	fmt.Printf("1 PB = %v Bytes\n", float64(PB))
	fmt.Printf("1 ZB = %v Bytes\n", float64(ZB))
	fmt.Printf("1 YB = %v Bytes\n", float64(YB))
	fmt.Println("==")

	b := ByteSize(1024)
	fmt.Println(b)

}