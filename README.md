# One-Year

prompt game thing

## Dev

api and client are at http://localhost:4000

### Protobuf

```sh
protoc \
  --go_out=. \
  --js_out=import_style=commonjs,binary:./client/src \
  pb/messages.proto
```
 
### Dependencies

```sh
go get
(cd /tmp; go get -u github.com/mitranim/gow; go install github.com/mitranim/gow)
(cd client; npx pnpm i)
```

### Server

```sh
go run .
```

### Client

```sh
(cd client; npm run dev)
```

