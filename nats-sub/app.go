package main

import (
	"flag"
	"log"

	"runtime"

	nc "github.com/gered/nats-cli"
	"github.com/nats-io/nats"
)

func usage() {
	log.Fatalf("nats-sub [-s server] [-ts] [-tls] [-tlscert CERT_FILE] [-tlskey KEY_FILE] [-tlscacert CA_FILE] [-tlsverify] <subject>")
}

func main() {
	log.SetFlags(0)

	var url = flag.String("s", nats.DefaultURL, "NATS comma-separate server URL list")
	var ts = flag.Bool("ts", false, "Display timestamp on logging output")
	var tls = flag.Bool("tls", false, "Enable TLS")
	var tlsCertPath = flag.String("tlscert", "", "Certificate file")
	var tlsKeyPath = flag.String("tlskey", "", "Private key file for certificate")
	var tlsCACertPath = flag.String("tlscacert", "", "Client certificate CA file")
	var tlsVerify = flag.Bool("tlsverify", false, "Enable TLS connection verification")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 1 {
		usage()
		return
	}

	var subject = flag.Arg(0)

	if *ts {
		log.SetFlags(log.LstdFlags)
	}

	conn, err := nc.Connect(*url, *tls, *tlsCertPath, *tlsKeyPath, *tlsCACertPath, *tlsVerify)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %s", err)
	}

	conn.Subscribe(subject, func(msg *nats.Msg) {
		log.Printf("[%s]: %s\n", msg.Subject, string(msg.Data))
	})
	conn.Flush()

	err = conn.LastError()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on subject: %s\n", subject)

	runtime.Goexit()
}
