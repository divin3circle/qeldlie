package books

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func SayBookName() {
	fmt.Println("The Go Programming language")
}

type File struct {
	count    int
	fileName string
}

func Duplicate() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		text := strings.TrimSpace(input.Text())
		counts[text]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func Duplicate2() {
	counts := make(map[string]File)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "duplicate: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n.count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.count, line, n.fileName)
		}
	}
}

func countLines(f *os.File, counts map[string]File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := strings.TrimSpace(input.Text())
		file, ok := counts[text]
		if ok {
			file.count++
		} else {
			file = File{1, f.Name()}
		}
		counts[text] = file
	}
}
