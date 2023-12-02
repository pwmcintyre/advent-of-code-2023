package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// parse params
	var (
		cookie    string
		day, year int
		today     = time.Now()
	)
	flag.IntVar(&day, "day", today.Day(), "day number to fetch, 1-25")
	flag.IntVar(&year, "year", today.Year(), "AOC year")
	flag.StringVar(&cookie, "cookie", os.Getenv("AOC_SESSION_COOKIE"), "AOC session cookie")
	flag.Parse()

	// get input
	res, err := Fetch(ctx, day, year, cookie)
	if err != nil {
		log.Fatalf("failed to get input data: %s", err)
	}

	// print
	fmt.Fprint(os.Stdout, string(res))
}

func Fetch(ctx context.Context, year, day int, cookie string) ([]byte, error) {

	// create request
	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	// add cookie
	req.AddCookie(&http.Cookie{Name: "session", Value: cookie})

	// send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("sending request: %w", err)
	}

	// read body
	return io.ReadAll(res.Body)
}
