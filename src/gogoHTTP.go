package gogoHTTP

import (
   "strings"
   "log"
    "flag"
    "net/http"
)

var dirFlag = flag.String("d",".","Specify directory to serve (default: /.)")
var portFlag = flag.String("p","8080","Specify directory to serve. (default: 8080)")

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}

func listen(){
    log.Printf("Listening on port %s",*portFlag)
    log.Fatal(http.ListenAndServe(strings.Join([]string{":",*portFlag},""), Log(http.FileServer(http.Dir(*dirFlag)))))
}

func main() {
    // Simple static webserver:
    flag.Parse()
    listen()
}
