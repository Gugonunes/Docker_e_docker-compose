//Código copiado do video "Docker e Docker Compose do zero ao Deploy"
//do candal Full Cycle
//disponivel em <https://www.youtube.com/watch?v=yb2udL9GG2U>

package main
import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Olá Full Cycle Developers!")
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}