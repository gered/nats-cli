package nats_cli

import (
	"errors"
	ctls "crypto/tls"
	"github.com/nats-io/nats"
)

func Connect(url string, tls bool, certPath string, keyPath string, caCertPath string, verify bool) (*nats.Conn, error) {
	if tls {
		if len(certPath) == 0 {
			return nil, errors.New("tlscert not set")
		}
		if len(keyPath) == 0 {
			return nil, errors.New("tlskey not set")
		}

		cert := nats.ClientCert(certPath, keyPath)
		var conn *nats.Conn
		var err error

		tlsverify := func(o *nats.Options) error {
			if o.TLSConfig != nil {
				o.TLSConfig.InsecureSkipVerify = !verify
			} else {
				o.TLSConfig = &ctls.Config{InsecureSkipVerify: !verify}
			}
			return nil
		}

		if len(caCertPath) > 0 {
			conn, err = nats.Connect(url, nats.RootCAs(caCertPath), cert, tlsverify)
		} else {
			conn, err = nats.Connect(url, cert, tlsverify)
		}

		if err != nil {
			return nil, err
		}

		return conn, nil

	} else {
		conn, err := nats.Connect(url)
		if err != nil {
			return nil, err
		}

		return conn, nil
	}
}
