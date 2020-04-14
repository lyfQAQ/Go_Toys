package dependencyInjection

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func TestGreet(t *testing.T) {

	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
