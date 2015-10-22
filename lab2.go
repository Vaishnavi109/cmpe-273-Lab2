package main
import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
    
)
type (  
    // User represents the structure of our resource
    Response struct {
        Greetings string `json:"greeting"`
        }
)
type (  
    // User represents the structure of our resource
    Request struct {
        Name string `json:"name"`
        }
)
func hello(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
    
    
    var resp Response
    var request Request
    json.NewDecoder(req.Body).Decode(&request)
    resp.Greetings = "Hello,"+request.Name+"!"
     w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
       
    uj, _ := json.Marshal(resp)
    fmt.Fprintf(w, "%s",uj)
}
func main() {
    mux := httprouter.New()
    mux.POST("/hello", hello)
    server := http.Server{
            Addr:        "0.0.0.0:8080",
            Handler: mux,
    }
    server.ListenAndServe()
}