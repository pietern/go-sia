# go-sia

> This implementation is very far from complete.

A SIA protocol reader/writer implementation.

Current status: it can receive and acknowledge messages from a
Honeywell Galaxy Flex 3 security system.

## Usage

There's one command in this repository at the time of writing. It
listens on port 10002 and acknowledges every SIA message that it
receives. Run it with:

``` sh
go run github.com/pietern/go-sia/cmd/ack
```


## License

MIT
