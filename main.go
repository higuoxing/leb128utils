package main

import (
	"fmt"
	"flag"
	"strconv"
	"os"
)

var Decode *bool
var Input *string
var IsBinary *bool
var IsSigned *bool

func init() {
	Decode = flag.Bool("d", false, "decode a hex or a binary string")
	Input = flag.String("i", "0", "input")
	IsBinary = flag.Bool("b", false, "the input string is in binary format")
	IsSigned = flag.Bool("s", false, "decode/encode a signed LEB128 integer")
	flag.Parse()
}

func printBuffer(buf []byte) {
	fmt.Printf("[ ")
	for _, b := range buf {
		fmt.Printf("0x%02x ", b)
	}
	fmt.Printf("]\n")
}

func check(e error) {
	if (e == nil) {
		return
	}
	fmt.Fprintf(os.Stderr, "%v\n", e)
	os.Exit(1)
}

func main() {
	if Decode != nil && *Decode {
		return
	}

	var buf []byte
	if IsSigned != nil && *IsSigned {
		signed, err := strconv.ParseInt(*Input, 0, 64)
		check(err)
		buf = AppendSleb128(buf, signed)
	} else {
		unsigned, err := strconv.ParseUint(*Input, 0, 64)
		check(err)
		buf = AppendUleb128(buf, unsigned)
	}

	printBuffer(buf)
}
