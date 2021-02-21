# GO-THRIFT-MOCK
![Test](https://github.com/chuan-yu/go-thrift-mock/actions/workflows/ci.yml/badge.svg)

`go-thrift-mock` is a thrift mock server written in `Go`. It can return any mock thrift response without a predefined IDL.

This project is inspired by [thrift-mock](https://github.com/didi/thrift-mock) project.

## Intallation
```
go get https://github.com/chuan-yu/go-thrift-mock
```

## Quick Start
### Start a mock server instance
```go
serverAddr := ":8888"
s := server.MustNewMockServer(serverAddr)

go func() {
    if err := s.Start(); err != nil {
        fmt.Printf("failed to start server: %s", err.Error())
        return
    }
}()

// wait for mock server to start properly
time.Sleep(1 * time.Second)
```
### Set expected return

Assume you have an IDL below.
```thrift
struct Request{
    1:optional string msg,
}

struct Response{
    1:required i32 code,
    2:required string responseMsg,
}

service HelloService {
    Response sayHello(1:required Request request);
}
```

To set the `sayHello` method to return a mock `Response` instance:
```go
result := Response{
    Code: 200,
    ResponseMsg: "mock message",
}
expectedReturn := server.ExpectedReturn{
    Response: &result,
}

s.SetExpectedReturn("sayHello", expectedReturn)
```
Now when a client calls the server, the above mock response is returned.

You can also mock an error response:
```go
expectedReturn = server.ExpectedReturn{
    Err: errors.New("mock error"),
}
s.SetExpectedReturn("sayHello", expectedReturn)
```
### Stop the mock server
```go
s.Stop()
```

## Example Code

See `example/main.go` for a working example
