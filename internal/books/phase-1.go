package books

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
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

func Fetch() {
	for _, url := range os.Args[1:] {
		if strings.HasPrefix(url, "http://") {
			url = url[:]
		} else {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("%s\t%s", b, resp.Status)
	}
}
