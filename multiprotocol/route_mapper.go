package multiprotocol

import "errors"

func MapHTTPToThrift(path string) (string, error) {
    switch path {
    case "/api/UserService/GetUser":
        return "UserService::GetUser", nil
    default:
        return "", errors.New("method not found")
    }
}