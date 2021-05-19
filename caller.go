package grpcdiff

import (
	"context"
	"time"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

type Caller struct {
	connA   *grpc.ClientConn
	connB   *grpc.ClientConn
	path    string
	reqType *desc.MessageDescriptor
	resType *desc.MessageDescriptor
}

func NewCaller(hostA, hostB string, service, method string) (*Caller, error) {
	connA, err := grpc.Dial(hostA, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	connB, err := grpc.Dial(hostB, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	rc := grpc_reflection_v1alpha.NewServerReflectionClient(connA)
	c := grpcreflect.NewClient(context.Background(), rc)

	req, res, err := getMessagesTypes(c, service, method)
	if err != nil {
		return nil, err
	}

	return &Caller{
		connA:   connA,
		connB:   connB,
		path:    service + "/" + method,
		reqType: req,
		resType: res,
	}, nil
}

func (c *Caller) Call(data string) (*Report, error) {
	resA, resB, durA, durB, err := c.callBoth(data)
	if err != nil {
		return nil, err
	}
	return NewReport(data, resA, resB, durA, durB), nil
}

func (c *Caller) callBoth(data string) (resA, resB []byte, durA, durB time.Duration, err error) {
	req := dynamic.NewMessage(c.reqType)
	err = req.UnmarshalJSON([]byte(data))
	if err != nil {
		return
	}

	var start time.Time

	start = time.Now()
	resA, err = c.call(c.connA, req)
	if err != nil {
		return
	}
	durA = time.Now().Sub(start)

	start = time.Now()
	resB, err = c.call(c.connB, req)
	if err != nil {
		return
	}
	durB = time.Now().Sub(start)

	return
}

func (c *Caller) call(conn *grpc.ClientConn, req *dynamic.Message) ([]byte, error) {
	res := dynamic.NewMessage(c.resType)

	err := conn.Invoke(context.Background(), c.path, req, res)
	if err != nil {
		return nil, err
	}

	b, err := res.MarshalJSONIndent()
	if err != nil {
		return nil, err
	}

	return b, nil
}
