<h1>HACKHATON 2025 BY BYTEDANCE (Third place winner)</h1>
<h2>Protocol Challenges in Cloud-Native Architecture</h2>

<h2>work partners</h2>
<p>
  <ul>Armando Qiu https://github.com/armandoqiu511</ul>
  <ul>Elbert A. Zhou https://github.com/MemerZxZ</ul>
</p>

<h2>Environment Preparation </h2>
1. go install github.com/cloudwego/thriftgo@latest
2. go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
3. go mod tidy
REMEBER USE "go mod tidy" when you will implement libs or frameworks

<h2>Thrift to Kitex </h2>
<p>// Execute under GOPATH</p>
<p>kitex -service hello ./idl/hello.thrift</p>

<p>// Execute not under GOPATH</p>
<p>kitex -service hello -module kitex-examples/kitex/thrift ./idl/hello.thrift</p>

<p>// Organize & pull dependencies</p>
<p>go mod tidy</p>

<h2>Deploy</h2>  
1. go run .
2. remember to kill the process after run it to debug

<h2>Linux usesfull command</h2>
1. Check ports: lsof -i :8888
2. Kill port: kill -9 <PID>


<h3>RESOURCES</h3>
Thrift and Kitex: https://github.com/cloudwego/kitex-examples/tree/main/kitex/thrift
