cd ../crun
rice embed-go
cd ..
go test -v _example\build\b_test.go -run Build$


:: 执行完成后 进行编译
go build  -o "build/windows/run.exe"  -ldflags=" -w -s" "build/run.go"

:: 编译linux版本
set GOOS=linux
go build  -o "build/linux/run"  -ldflags=" -w -s" "build/run.go"