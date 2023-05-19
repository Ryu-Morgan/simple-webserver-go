main package

import(
	"fmt"
	"log"
	"net/http"
)

func root_handler(w http.ResponseWriter , r *http.Request){
	fmt.Fprintln(w, "Hello agian")
}

func main()  {
	http.HandleFunc("/", root_handler)
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}