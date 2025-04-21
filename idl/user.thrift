namespace go user

service UserService {
    string GetUser(1: i64 userID)
}