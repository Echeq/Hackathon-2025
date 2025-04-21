package multiprotocol

import "github.com/mitchellh/mapstructure"

func ConvertJSONToThrift(jsonData map[string]interface{}, thriftStruct interface{}) error {
    return mapstructure.Decode(jsonData, thriftStruct)
}