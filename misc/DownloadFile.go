package misc

import (
	"io"
	"net/http"
	"os"
)

// DownloadFile ...
func DownloadFile(filepath string, url string, ch chan<- string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	ch <- filepath
	return err
}
