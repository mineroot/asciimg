package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"runtime"
	"runtime/debug"
	"time"
)

var (
	charSet            = "@%#*+=-:. "
	asciiWidth         = 120
	reversed           = false
	widthToHeightRatio = 0.5
	printDebug         = false
)

func main() {
	start := time.Now()
	debug.SetGCPercent(-1)         // disable GC
	debug.SetMemoryLimit(50 << 20) // set soft limit to 50 MiB

	flag.StringVar(&charSet, "charset", charSet, "Charset to use for ASCII")
	flag.IntVar(&asciiWidth, "width", asciiWidth, "ASCII width [10-240]")
	flag.Float64Var(&widthToHeightRatio, "ratio", widthToHeightRatio, "ASCII width to height ratio")
	flag.BoolVar(&reversed, "reversed", reversed, "Reverse ASCII")
	flag.BoolVar(&printDebug, "debug", printDebug, "Print debug info")
	flag.Parse()

	if err := asciimg(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if printDebug {
		elapsed := time.Since(start)
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		_, _ = fmt.Fprintln(os.Stderr)
		_, _ = fmt.Fprintf(os.Stderr, "Took %s\n", elapsed)
		_, _ = fmt.Fprintf(os.Stderr, "Alloc = %v KiB, Sys = %v KiB, NumGC = %v\n", m.Alloc/1024, m.Sys/1024, m.NumGC)
	}
}

func asciimg() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = r.(error)
		}
	}()

	// input checks
	for _, char := range []byte(charSet) {
		if char < 32 || char > 126 {
			return fmt.Errorf("invalid ASCII character(s)")
		}
	}
	if asciiWidth < 10 || asciiWidth > 240 {
		return fmt.Errorf("width must be between 10 and 240")
	}
	if widthToHeightRatio < 0.1 || widthToHeightRatio > 10 {
		return fmt.Errorf("ratio must be between 0.1 and 10")
	}
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return fmt.Errorf("no input provided. Usage: %s < image.png", os.Args[0])
	}

	// init img and other variables
	img, _, err := image.Decode(os.Stdin)
	if err != nil {
		return err
	}
	charSetLen := len(charSet)
	grayScale := 256 / float64(charSetLen)
	imgWidth, imgHeight := img.Bounds().Max.X, img.Bounds().Max.Y
	asciiHeight := imgHeight * asciiWidth / imgWidth
	asciiHeight = int(float64(asciiHeight) * widthToHeightRatio)
	asciiWidthPx := imgWidth / asciiWidth
	asciiHeightPx := imgHeight / asciiHeight
	asciiTotalPx := asciiWidthPx * asciiHeightPx

	asciiImage := make([][]byte, asciiHeight)
	// calculate each ASCII char
	// time: O(imgWidth * imgHeight)
	for y := 0; y < asciiHeight; y++ {
		asciiImage[y] = make([]byte, asciiWidth)
		for x := 0; x < asciiWidth; x++ {
			sum := 0
			for i := 0; i < asciiWidthPx; i++ {
				for j := 0; j < asciiHeightPx; j++ {
					originalColor := img.At(x*asciiWidthPx+i, y*asciiHeightPx+j)
					grayColor := int(color.GrayModel.Convert(originalColor).(color.Gray).Y)
					sum += grayColor
				}
			}
			avgGray := float64(sum / asciiTotalPx)
			if reversed {
				avgGray = 255 - avgGray
			}
			setIndex := int(math.Round(avgGray / grayScale))
			if setIndex >= charSetLen {
				setIndex = charSetLen - 1
			}
			asciiChar := charSet[setIndex]
			asciiImage[y][x] = asciiChar
		}
	}

	// print ASCII art
	for i := 0; i < len(asciiImage); i++ {
		fmt.Println(string(asciiImage[i]))
	}

	return nil
}
