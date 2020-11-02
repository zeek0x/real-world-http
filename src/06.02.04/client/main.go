package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

func main() {
	cert, err := tls.LoadX509KeyPair("ssl/client.crt", "ssl/client.key")
	if err != nil {
		panic(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	resp, err := client.Get("https://localhost:18443")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	log.Println(string(dump))
}

/*
うまくいかない...
証明書作成手順は ./ssl/README.md に従う

$ server.go
$ go run main.go
2020/11/02 02:46:58 http: TLS handshake error from 127.0.0.1:59754: remote error: tls: bad certificate

$ cd client
$ go run main.go
panic: Get "https://localhost:18443": x509: certificate signed by unknown authority (possibly because of "x509: invalid signature: parent certificate cannot sign this kind of certificate" while trying to verify candidate authority certificate "example.com")

goroutine 1 [running]:
main.main()
        /workspaces/real-world-http/src/06.02.04/client/main.go:26 +0x3a3
exit status 2
*/
