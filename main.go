package main

import (
    "github.com/AaronO/gogo-proxy"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    upstream_server := os.Getenv("UPSTREAM_SERVER")
    p, _ := proxy.New(proxy.ProxyOptions{
        Balancer: func(req *http.Request) (string, error) {
            return upstream_server, nil
        },
    })

    http.ListenAndServe(":"+port, p)
}
