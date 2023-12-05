# Proto files
## auth.proto (Path: ProteiTestCase/api/proto/auth.proto)
<p>Authentication service with 1 rpc is Login.</p>

```
    rpc Login(LoginRequest) returns (LoginResponce) {};
```

<p>On input "LoginRequest". Fields: login(string) and password(string)</p>

```
    string login = 1;
    string password = 2;
```

<p>On output "LoginResponce". Field: token(string) is JWT token</p>

```
    string token = 1;
```

## app.proto (Path: ProteiTestCase/api/proto/app.proto)
<p>Service with 2 rpcs is GetInfoAboutUser and CheckAbsenceStatus.</p>

```
    rpc GetInfoAboutUser(GetInfoRequest) returns (GetInfoResponse) {}
    rpc CheckAbsenceStatus(AbsenceStatusRequest) returns (AbsenceStatusResponse) {}
```

<p>(1) RPC GetInfoAboutUser: On input "GetInfoRequest". Fields: array of ids(int64), name(string), work phone(string), email(string), 
dateFrom(time), dateTo(time)</p>

```
    repeated int64 id = 1;
    string name = 2;
    string workPhone = 3;
    string email = 4;
    google.protobuf.Timestamp dateFrom = 5;
    google.protobuf.Timestamp dateTo = 6;
```

<p>RPC GetInfoAboutUser: On output "GetInfoResponse". Fields: status(string), array of structs "OutputUsersData"</p>

```
    string status = 1;
    repeated OutputUsersData usersData = 2;
```
<p>OutputUsersData. Fields: id(int64), displayName(string), email(string), work phone(string)</p>

```
    int64 id = 1;
    string displayName = 2;
    string email = 3;
    string workPhone = 4;
```

<p>(2) RPC CheckAbsenceStatus:
