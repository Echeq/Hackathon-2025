namespace go user

struct User {
    1: required string ID,
    2: required string Name,
}

service UserService {
    User GetUser(1: string ID),
}