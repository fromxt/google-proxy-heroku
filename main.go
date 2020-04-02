package main

import (
    "github.com/AaronO/gogo-proxy"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    p, _ := proxy.New(proxy.ProxyOptions{
        Balancer: func(req *http.Request) (string, error) {
            return <%= ENV["UPSTREAM_SERVER"] %>, nil
        },
    })

    http.ListenAndServe(":"+port, p)
}
