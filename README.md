# Quick Start

This is a very simple guide to the request library. Just want to be easy to do what I would like to do. So I build it.

## Install

You just need to copy and paste the command to the terminal and press the Enter key.
```
go get -u github.com/tansoz/request
```

## Usage

First, You need import the library with the import statement.

```GO
import "github.com/tansoz/request"
```

Make sure the library was imported before writing a request with the library.

### GET
Use the library to send a simple GET request
```GO
if resp := request.Get("https://httpbin.org/get", map[]); resp.Error() == nil {
    var tmp bytes.Buffer
    tmp.ReadFrom(resp.Readable())
    fmt.Println(tmp.String())
} else {
    fmt.Println(resp.Error())
}
```
Result:
```JSON
{
  "args": {
    "id": "1515", 
    "method": "get"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Host": "httpbin.org", 
    "User-Agent": "request/1.1.2", 
    "X-Amzn-Trace-Id": "Root=1-61928190-67740ba50929357a16b5f917"
  }, 
  "origin": "123.154.60.166", 
  "url": "https://httpbin.org/get?id=1515&method=get"
}
```
### POST
Use the library to send a simple POST request
```GO
if resp := request.Post("https://httpbin.org/post", map[string]interface{}{
    "id":  1515,
    "foo": "bar",
}); resp.Error() == nil {
    var tmp bytes.Buffer
    tmp.ReadFrom(resp.Readable())
    fmt.Println(tmp.String())
} else {
    fmt.Println(resp.Error())
}
```
Result:
```JSON
{
  "args": {}, 
  "data": "", 
  "files": {}, 
  "form": {
    "foo": "bar", 
    "id": "1515"
  }, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "15", 
    "Content-Type": "application/x-www-form-urlencoded; charset=utf-8", 
    "Host": "httpbin.org", 
    "User-Agent": "request/1.1.2", 
    "X-Amzn-Trace-Id": "Root=1-61928192-327e74387dd38edc28ebdcb3"
  }, 
  "json": null, 
  "origin": "123.154.60.166", 
  "url": "https://httpbin.org/post"
}
```
### UPLOAD FILES WITH METHOD POST
Use the library to send a simple POST request
```GO
if resp := request.New(
    option.POST,
    option.URL("https://httpbin.org/post"),
).Go(body.MultipartBody([]body.Field{
    body.FileBodyField("files", "testfile.txt", "go.sum", "text/plain", 0, -1, nil),
}, "", nil)); resp.Error() == nil {
    var tmp bytes.Buffer
    tmp.ReadFrom(resp.Readable())
    fmt.Println(tmp.String())
} else {
    fmt.Println(resp.Error())
}
```
Result:
```JSON
{
  "args": {}, 
  "data": "", 
  "files": {
    "files": "github.com/tansoz/request v1.1.2 h1:z5So8y7pVoOwZ1WKXuqIiTzZ3GOvHuxbeGsr8gu6cro=\ngithub.com/tansoz/request v1.1.2/go.mod h1:FDMFH2cBkmqc5nykpBCF520HD7clEI/tnIUCv0w0Jac=\n"
  }, 
  "form": {}, 
  "headers": {
    "Accept-Encoding": "gzip", 
    "Content-Length": "351", 
    "Content-Type": "multipart/form-data; boundary=NTU3NzAwNjc5MzU4NDc3MTQzNQ", 
    "Host": "httpbin.org", 
    "User-Agent": "request/1.1.2", 
    "X-Amzn-Trace-Id": "Root=1-6192841a-20d704f01eada96111bb0c3d"
  }, 
  "json": null, 
  "origin": "123.154.60.166", 
  "url": "https://httpbin.org/post"
}
```

## Advance

### New a request
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request")
)
```
### Send a request
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request")
).Go(nil) // Call the Go function with a null body
```
### Set method
```GO
request.New(
    option.Method("GET"), // option.POST, option.PUT, options.HEAD ...
    option.URL("https://github.com/tansoz/request"),
)
```
### Set URL
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request"), // this option method is set request URL
)
```
### Set Headers
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request"),
    option.Headers(map[string][]string{
        "User-Agent":{"the request library by Romm Hui"}
    }),
)
```
### Set Proxy URL
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request"),
    option.Headers(map[string][]string{
        "User-Agent":{"the request library by Romm Hui"}
    }),
    option.ProxyURL("http://127.0.0.1:8080"), // supported socks, http, https scheme
)
```
### Set Timeout
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request"),
    option.Headers(map[string][]string{
        "User-Agent":{"the request library by Romm Hui"}
    }),
    option.ProxyURL("http://127.0.0.1:8080"),
    option.Timeout(1000), // 1 second here, millisecond as a unit
)
```
### Set Hostname
```GO
request.New(
    option.Method("GET"),
    option.URL("https://github.com/tansoz/request"),
    option.Headers(map[string][]string{
        "User-Agent":{"the request library by Romm Hui"}
    }),
    option.ProxyURL("http://127.0.0.1:8080"),
    option.Timeout(1000),
    option.Host("www.github.com"), // set hostname
)
```
### Send a request with a Multipart body
```GO
multipartBody := body.MultipartBody(
    []body.Field{
        body.FileBodyField("files", "testfile.txt", "go.mod", "text/plain", 0, -1, nil),
		body.ParamField("id", 15, nil),
		body.QueryBodyField("data", map[string]interface{}{
			"str":    "hello world",
			"number": 15,
			"bool":   false,
		}, nil),
    },
    "",
    nil,
)
request.New(
    option.Method("POST"),
    option.URL("https://httpbin.org/post"),
    option.Headers(map[string][]string{
        "User-Agent":{"the request library by Romm Hui"},
    }),
).Go(multipartBody)
```
Result:
```JSON
{
  "args": {},
  "data": "",
  "files": {
    "files": "module github.com/tansoz/request\n\ngo 1.13\n"
  },
  "form": {
    "data": "bool=false&number=15&str=hello+world",
    "id": "15"
  },
  "headers": {
    "Accept-Encoding": "gzip",
    "Content-Length": "520",
    "Content-Type": "multipart/form-data; boundary=NTU3NzAwNjc5MzU4NDc3NDcwOA",
    "Host": "httpbin.org",
    "User-Agent": "the request library by Romm Hui",
    "X-Amzn-Trace-Id": "Root=1-619290f8-7c81c8bd5245696b5eded914"
  },
  "json": null,
  "origin": "123.154.60.166",
  "url": "https://httpbin.org/post"
}
```