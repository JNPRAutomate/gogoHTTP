package main

import (
   "strings"
   "log"
   "flag"
   "net/http"
   "os/signal"
   "os"
)

var dirFlag = flag.String("d",".","Specify directory to serve (default: /.)")
var portFlag = flag.String("p","8080","Specify directory to serve. (default: 8080)")

var servecount int

func Log(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        servecount = servecount + 1
        log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
        handler.ServeHTTP(w, r)
    })
}

func listen(){
    //listen for ctrl-c
    go func(){
        sigchan := make(chan os.Signal, 1)
        signal.Notify(sigchan, os.Interrupt)
        <-sigchan
        log.Printf("Served %d requests", servecount)
        log.Println("Terminating gogoHTTP, Goodbye!")
        os.Exit(0)
    }()
    log.Printf("Listening on port %s",*portFlag)
    log.Fatal(http.ListenAndServe(strings.Join([]string{":",*portFlag},""), Log(http.FileServer(http.Dir(*dirFlag)))))
}

func main() {
    // Simple static webserver:
    servecount = 0
    flag.Parse()
    listen()
}
