# contact-api

A very simple API that receives POST requests containing messages and sends them using a gmail account.

To use it install the required libraries:

```
go get gopkg.in/gcfg.v1
go get github.com/SlyMarbo/gmail
```

and rename `config.gcfg.dist` in `config.gcfg` after having updated your configuration. Start the server with `go run contact-api.go`.
