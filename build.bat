rice embed-go
cd  .\_example\run\
set GOOS=linux
go build -ldflags=" -w -s" webServe.go
# start ./webServe.exe
