Environment Preparation
1. go install github.com/cloudwego/thriftgo@latest
2. go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
3. go get github.com/stretchr/testify

Commands used:
kitex -module kitex-multi-protocol -service UserService idl/UserService.thrift
go get github.com/cloudwego/kitex/server
go get github.com/cloudwego/kitex/pkg/transport
go get github.com/mitchellh/mapstructure