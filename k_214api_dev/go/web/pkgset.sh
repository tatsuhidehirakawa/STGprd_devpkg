#!/usr/bin/bash

go mod init github.com/goark/pa-api
go get github.com/kyleconroy/sqlc/cmd/sqlc
go install github.com/kyleconroy/sqlc/cmd/sqlc
sqlc version
sqlc generate --file sqlc/sqlc.yaml
cd sqlc/build_sqlc
go mod init sqlc_pkg
cd ../..
go get sqlc_pkg

# 「src/web/go.mod」の末尾に「replace sqlc_pkg => ./sqlc/build_sqlc」を加える
# go mod tidy
# go run main.go