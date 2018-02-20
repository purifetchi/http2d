package main

import (
        "strconv"
        "fmt"
        "net/http"
        "flag"
        "os/exec"
        "golang.org/x/net/http2"
)

func main() {
        var srv http.Server

        keyPtr := flag.String("key", "server.key", "Path to the SSL key")
        crtPtr := flag.String("cert", "server.crt", "Path to the SSL certificate")
        insecurePtr := flag.Bool("http1", false, "Run in HTTP/1.1 mode (default is false)")
        portPtr := flag.Int("port", 8080, "Port to run on")
        flag.Parse()

        if *portPtr != 80 {
                srv.Addr = ":" + strconv.Itoa(*portPtr)
        }

        if *insecurePtr == false {
                fmt.Println("[INFO] Server starting.")
                http2.ConfigureServer(&srv, nil)
                http.HandleFunc("/", cmdexec)

                srv.ListenAndServeTLS(*crtPtr, *keyPtr)
        } else {
                fmt.Println("[INFO] Server starting in HTTP/1.1 mode.")
                http.HandleFunc("/", cmdexec)
                srv.ListenAndServe()
        }

}

func cmdexec(w http.ResponseWriter, r *http.Request) {
        out, err := exec.Command(flag.Args()[0], r.URL.Path).Output()
        var status = 200
        if err != nil {
                status = 500
                w.WriteHeader(status)
                w.Write([]byte("<h1>HTTP 500 - Internal Server Error</h1><pre>" + err.Error() + "</pre>"))
        } else {
                w.Write([]byte(out))
        }
        fmt.Println("[INFO] " + r.RemoteAddr + " - " + r.URL.Path + "(" + strconv.Itoa(status) + ")")
}