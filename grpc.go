package grpcdiff

import (
	"context"
	"fmt"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

func SimpleCall(host, service, method, data string) ([]byte, error) {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	rc := grpc_reflection_v1alpha.NewServerReflectionClient(conn)
	c := grpcreflect.NewClient(context.Background(), rc)

	req, res, err := getMessages(c, service, method)
	if err != nil {
		return nil, err
	}

	err = req.UnmarshalJSON([]byte(data))
	if err != nil {
		return nil, fmt.Errorf("failed unmarshal json '%s': %v", method, err)
	}

	err = conn.Invoke(context.Background(), service+"/"+method, req, res)
	if err != nil {
		return nil, fmt.Errorf("invoke failed: %v", err)
	}

	b, err := res.MarshalJSONIndent()
	if err != nil {
		return nil, fmt.Errorf("marshal response failed: %v", err)
	}

	return b, nil
}

func getMessagesTypes(c *grpcreflect.Client, service, method string) (req, res *desc.MessageDescriptor, err error) {
	s, err := c.ResolveService(service)
	if err != nil {
		return
	}

	m := s.FindMethodByName(method)
	if m == nil {
		err = fmt.Errorf("cannot find method '%s'", method)
		return
	}

	req = m.GetInputType()
	res = m.GetOutputType()
	return
}

func getMessages(c *grpcreflect.Client, service, method string) (req, res *dynamic.Message, err error) {
	dreq, dres, err := getMessagesTypes(c, service, method)
	if err != nil {
		return
	}
	req = dynamic.NewMessage(dreq)
	res = dynamic.NewMessage(dres)
	return
}
