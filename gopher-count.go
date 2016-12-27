package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		if i > 0 {
			doCount(v)
		}
	}
}

func doCount(filename string) {
	fd, err := os.Open(filename)
	if err != nil {
		println("Cannot open file", filename, err.Error())
		return
	}

	stats, err := fd.Stat()
	if err != nil {
		println("Cannot stat file", filename, err.Error())
		return
	}

	fmt.Printf("Compiled JS: %s = %d bytes\n", filename, stats.Size())

	data := make([]byte, stats.Size())
	_, err = fd.Read(data)
	if err != nil {
		println("Error reading file ", filename, err.Error())
		return
	}

	// loop through a line at a time
	startLine := 0
	startOfLine := true
	packageName := ""
	for i, v := range data {
		if startOfLine {
			if string(data[i:i+11]) == `$packages["` {
				// scan out past 22 to get the terminating "
				for ii := i + 12; ; ii++ {
					if data[ii] == 34 {
						// println("ThePackage", data[i+11:ii])
						packageName = string(data[i+11 : ii])
						break
					}
					if data[ii] == 10 {
						packageName = ""
						break
					}
				}
				startOfLine = false
			}
		}
		if v == '\n' {
			if packageName != "" {
				// println("Package", packageName, "goes from", startLine, "to", i, "which is", i-startLine, "bytes")
				fmt.Printf("%d\t%s\n", i-startLine, packageName)
				if packageName == "main" {
					// all done
					return
				}
			}
			// println("Newline at", i, startLine, i-startLine)
			startLine = i + 1
			startOfLine = true
			packageName = ""
		}
	}

}
