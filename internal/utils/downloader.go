package utils

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

// DownloadURLs downloads all provided URLs concurrently. It limits
// the number of concurrent downloads to the provided concurrencyLevel.
func DownloadURLs(urls []url.URL, concurrencyLevel int, outputDirectory string, out io.Writer, quiet bool) error {
	sem := NewSemaphore(concurrencyLevel)
	errors := make(chan error)
	var wg sync.WaitGroup

	for _, u := range urls {
		wg.Add(1)

		go func(u url.URL) {
			defer wg.Done()

			sem.Down()
			if u.Scheme == "http" || u.Scheme == "https" {
				filePath := path.Join(outputDirectory, path.Base(u.Path))
				err := httpFetch(filePath, u.String())
				errors <- err
			}
			sem.Up()
		}(u)
	}

	go func() {
		wg.Wait()
		close(errors)
	}()

	for err := range errors {
		if err != nil {
			fmt.Fprintf(out, "error downloading file: %v\n", err)
		}
	}

	return nil
}

func httpFetch(filePath string, url string) error {
	out, err := os.Create(filePath + ".part")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	err = os.Rename(filePath + ".part", filePath)
	if err != nil {
		return err
	}

	return nil
}
