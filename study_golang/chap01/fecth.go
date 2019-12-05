package main
import (
	"fmt"
	"net/http"
	"os"
	"io"
	"strings"
)


func main_1_5() {
	for _, url := range os.Args[1:] {
		t := strings.HasPrefix(url, "http://") 
		if !t {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v", err)
			os.Exit(1)
		}
		// b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		fmt.Println(resp.Status)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}