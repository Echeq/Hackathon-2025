Environment Preparation
1. go install github.com/cloudwego/thriftgo@latest
2. go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
3. go mod tidy
REMEBER USE "go mod tidy" when you will implement libs or frameworks

Deploy:
1. go run .
2. remember to kill the process after run it to debug

Linux usesfull command:
1. Check ports: lsof -i :8888
2. Kill port: kill -9 <PID>
