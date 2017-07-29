# picfinder

```
go get -u github.com/golang/dep/cmd/dep # See https://github.com/golang/dep

go get github.com/dcrosby42/picfinder
cd $GOPATH/src/github.com/dcrosby42/picfinder
dep ensure
go install github.com/dcrosby42/picfinder/picfinder
```

```
tools/create=mysqldb-and-grants.sh
tools/picfinder-db.sh
```

```
picfinder db rebuild -really
```


### Generating the grpc code

1. Temporarily add imports to `main.go`:
  - `_ "github.com/golang/protobuf/proto"`
  - `_ "github.com/golang/protobuf/protoc-gen-go"`
  - `_ "google.golang.org/grpc"`
1. Run `dep ensure`
1. Remove temporary imports from `main.go`
1. Run `./grpc/gen_protobuf.sh`
