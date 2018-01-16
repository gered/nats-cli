# NATS CLI Client

Just a really, really simple NATS CLI client. This is primarily intended for my own personal use. I needed a client 
that I could use for simple tasks here and there. I didn't see an existing one (with a minimum of dependencies) that 
supported TLS connections to the NATS server, which I need.

Includes:

```
nats-pub [-s server] [-tls] [-tlscert CERT_FILE] [-tlskey KEY_FILE] [-tlscacert CA_FILE] [-tlsverify] <subject> <message>

nats-sub [-s server] [-ts] [-tls] [-tlscert CERT_FILE] [-tlskey KEY_FILE] [-tlscacert CA_FILE] [-tlsverify] <subject>

```

I will add more support for extra features, etc. later if/when I need them.
