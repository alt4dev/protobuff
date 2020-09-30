<a href="https://alt4.dev"><img src="https://alt4.dev/banner.svg" alt="" height="120"></a>

## Protocol Buffers

Alt4 uses <a href="" target="_blank">grpc</a> to ingest logs.
This repo holds protocol buffer definitions used by Alt4dev 
client libraries and internally by the server.


*Unless making a custom client, you don't need to use these definitions
Instead we'd recommend you use the ready-made client libraries*

#### Client Libraries
- [go](https://github.com/alt4dev/go)

#### Code Generation
```shell script
# Go Lang
./build.sh go
```

### Usage
This section refers to code that is generated and ready to be used publicly

#### Golang
In your project install:
```shell script
go get github.com/alt4dev/protobuff
```
