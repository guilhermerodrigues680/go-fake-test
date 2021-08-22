package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fgBlackBgGreen := color.New(color.FgBlack, color.BgGreen).SprintFunc()
	fgGreen := color.New(color.FgGreen).SprintFunc()
	fgBlackBgRed := color.New(color.FgBlack, color.BgRed).SprintFunc()
	fgRed := color.New(color.FgRed).SprintFunc()

	for i := 0; i < 10; i++ {
		d := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("  Running Test %d", i)
		pass := randBool()
		time.Sleep(d)
		fmt.Print("\r")
		const strFormat = "%s %s Test %d (%d ms)\n"
		if pass {
			fmt.Printf(strFormat, fgGreen("✓"), fgBlackBgGreen(" PASS "), i, d.Milliseconds())
		} else {
			fmt.Printf(strFormat, fgRed("✖"), fgBlackBgRed(" FAIL "), i, d.Milliseconds())
		}
	}

}

func randBool() bool {
	return rand.Float32() < 0.5
}
