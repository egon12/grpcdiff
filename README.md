# grpcdiff
Tools to see different betweenr two different GRPC implementation with same interface

## How to install?
If you have go in your environment you can use
```
go install github.com/egon12/grpcdiff
```

## How to use it?
First you need a request in JSON format that will be converted into
GRPC Request (protobuf to be exactly). And seperate it by line

For examples:

```
{}
{"content":"hello"}
{"content":"hello_there"}
```

and you can save it as `requests.jsonrequest` or whatever name you like.
For now please make your request in one line only. (This app doesn't support pretty print
JSON format)

then run it like this.

```
cat requests.jsonrequest | grpcdiff localhost:50051 yourservice.com:50051 proto.ServiceApp Func
```

If it detect some differences it will print the differences 
And if it not detect some differences it will print the time comparison

