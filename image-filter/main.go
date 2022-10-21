package main

import (
	"bmpfilter/filter"
	"fmt"
	"io"
	"os"
	"unsafe"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "bmpfilter: [image] [-gtsec]")
		fmt.Fprintln(os.Stderr, "\timage: path to input image")
		fmt.Fprintln(os.Stderr, "\tfilter: where filter is one of:")
		fmt.Fprintln(os.Stderr, "\t\t-t: threshold, default\n\t\t-g: grayscale")
		return
	}

	stat, err := os.Stat(os.Args[1])
	if err != nil {
		panic(err)
	} else if stat.IsDir() {
		fmt.Fprintln(os.Stderr, os.Args[1], " is a directory")
		return
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error in read all")
		return
	}

	// probably need better way to extract these
	off := *(*int32)(unsafe.Pointer(&bytes[10]))
	width := *(*int32)(unsafe.Pointer(&bytes[18]))
	height := *(*int32)(unsafe.Pointer(&bytes[22]))

	var (
		fil     filter.Filter
		outfile string
	)

	switch os.Args[2] {
	case "-t":
		fil, outfile = filter.Thres, "thres.bmp"
	case "-g":
		fil, outfile = filter.Gray, "gray.bmp"
	}

	apply_filter(bytes[off:], width, height, fil)

	err = os.WriteFile(outfile, bytes, 0666)
	if err != nil {
		panic(err)
	}
}

func apply_filter(image []byte, width int32, height int32, f filter.Filter) {
	var row_offset int32 = 0
	for i := int32(0); i < height; i++ {
		var pix_offset int32 = 0
		for j := int32(0); j < width; j++ {
			f(image[row_offset+pix_offset:])
			pix_offset += 3
		}
		row_offset += (width*3 + width%4)
	}
}
