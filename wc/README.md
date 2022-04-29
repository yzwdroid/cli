# Go Modules

go mod init github.com/yzwdroid/go-modules-example

```bash
// Create by go init
go.mod
// Created by go build or go test
go.sum
```

# Compling your tool for different platforms

```bash
GOOS=windows go build -o mytool.exe
GOOS=linux go build -o mytool
```
