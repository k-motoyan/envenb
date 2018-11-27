# envenb

envenb is a support tool for embedding environment variables in Golang's binary.

## Install

```
go get -v github.com/k-motoyan/envenb
```

## Usage

```
cat .env | envenb | gofmt > envenb.go
go build main.go envenb.go
```

## License

MIT
