package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
)

func main() {
	port := flag.Int("port", 9100, "port number")
	bind := flag.String("bind", "", "bind address")
	flag.Parse()

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		for _, domain := range flag.Args() {
			ips, _ := net.LookupIP(domain)
			fmt.Fprintf(w, "dns_a_replies{domain=\"%s\"} %d\n", domain, len(ips))
		}
	})

	http.ListenAndServe(fmt.Sprintf("%s:%d", *bind, *port), nil)
}
