rice embed-go
cd  .\_example\run\
go build -ldflags=" -w -s" webServe.go
start ./webServe.exe
