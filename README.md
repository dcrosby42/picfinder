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
picfinder server
```

```
# Ping the api server
picfinder ping --host localhost --port 13131

# Scan local files and send to the api:
picfinder scan update --dir ~/Pictures/Photosa --host localhost --port 13131
```


### Generating the grpc protobuf code

1. Temporarily add imports to `main.go`:
  - `_ "github.com/golang/protobuf/proto"`
  - `_ "github.com/golang/protobuf/protoc-gen-go"`
  - `_ "google.golang.org/grpc"`
1. Run `dep ensure`
1. Remove temporary imports from `main.go`
1. Run `./grpc/gen_protobuf.sh`

### Tools used

- sqlx for mysql database - http://jmoiron.github.io/sqlx/
- grpc for api - https://github.com/grpc/grpc-go
- urfave/cli for CLI commands - http://github.com/urfave/cli 
- golang/dep for dependency mgmt: https://github.com/golang/dep



dupes

~host content -> duplicate file on a different host
  insert as usual

host ~path content -> duplicate file on host at different place on disk
  insert as usual

host path ?content -> repeated send from client
  updated db if request's scan date is newer
  return file_info ID of existing record


