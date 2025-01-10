# JWT-decoder
JSON Web Token decoder tool written in golang.

## Usage

```console
go run main.go -t eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTY5NDQ5NDUzNSwiZXhwIjoxNjk0NDk4MTM1fQ.UkhULVq2XdcHf64uPyw2UiHTIjyeuvWuhqgdXc3PhLo
```

## Output

```console
-------------------------------------
Header:  {"typ":"JWT","alg":"HS256"}
Payload:  {
  "sub": "1234567890",
  "name": "John Doe",
  "admin": true,
  "iat": 1694494535,
  "exp": 1694498135
}
Signature:  UkhULVq2XdcHf64uPyw2UiHTIjyeuvWuhqgdXc3PhLo
-------------------------------------
```

## Build Binary

```sh
go build main.go
```
