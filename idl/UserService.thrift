namespace go UserService

struct GetUserRequest {
    1: required i64 id;
}

struct GetUserResponse {
    1: required string name;
}

service UserService {
    GetUserResponse GetUser(1: GetUserRequest request);
}