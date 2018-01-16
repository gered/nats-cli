package main

import (
	"flag"
	"log"

	nc "github.com/gered/nats-cli"
	"github.com/nats-io/nats"
)

func usage() {
	log.Fatalf("nats-pub [-s server] [-tls] [-tlscert CERT_FILE] [-tlskey KEY_FILE] [-tlscacert CA_FILE] [-tlsverify] <subject> <message>")
}

func main() {
	log.SetFlags(0)

	var url = flag.String("s", nats.DefaultURL, "NATS comma-separate server URL list")
	var tls = flag.Bool("tls", false, "Enable TLS")
	var tlsCertPath = flag.String("tlscert", "", "Certificate file")
	var tlsKeyPath = flag.String("tlskey", "", "Private key file for certificate")
	var tlsCACertPath = flag.String("tlscacert", "", "Client certificate CA file")
	var tlsVerify = flag.Bool("tlsverify", true, "Enable TLS connection verification")

	flag.Usage = usage
	flag.Parse()

	if flag.NArg() < 2 {
		usage()
	}

	var subject = flag.Arg(0)
	var message = flag.Arg(1)

	conn, err := nc.Connect(*url, *tls, *tlsCertPath, *tlsKeyPath, *tlsCACertPath, *tlsVerify)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %s", err)
	}

	conn.Publish(subject, []byte(message))
	conn.Flush()

	err = conn.LastError()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Published message on subject %s\n", subject)

	conn.Close()
}
