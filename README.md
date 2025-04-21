Environment Preparation
1. go install github.com/cloudwego/thriftgo@latest
2. go install github.com/cloudwego/kitex/tool/cmd/kitex@latest
3. go get github.com/stretchr/testify

Commands used:
1. go get github.com/cloudwego/kitex/server
2. go get github.com/cloudwego/kitex/pkg/transport
3. go get github.com/mitchellh/mapstructure
4. go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

Generate .thrift
kitex -module kitex-multi-protocol -service UserService idl/UserService.thrift