package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/chuan-yu/go-thrift-mock/resouces/gen-go/resources"
	"github.com/chuan-yu/go-thrift-mock/server"
)

func handleClient(c *resources.HelloServiceClient) (err error) {
	resp, err := c.SayHello(context.Background(), &resources.Request{
		Msg: &[]string{"hello"}[0],
	})
	if err == nil {
		fmt.Println(resp.Code)
		fmt.Println(resp.ResponseMsg)
	} else {
		fmt.Println(err)
	}
	return nil
} 

func runClient(serverAddr string) error {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	var transport thrift.TTransport
	transport, err := thrift.NewTSocket(serverAddr)

	if err != nil {
		return fmt.Errorf("Error opening socket:", err) 
	}

	if transport == nil {
		return fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return err
	}

	err = transport.Open()
	if err != nil {
		return fmt.Errorf("error running client: %s", err)
	}

	defer transport.Close()

	c := resources.NewHelloServiceClientFactory(transport, protocolFactory)
	return handleClient(c)
}

func main() {
	serverAddr := ":8888"
	s := server.MustNewMockServer(serverAddr)
	result := resources.Response{
		Code: 200,
		ResponseMsg: "mock message",
	}
	s.SetExpectedReturn("sayHello", &result)

	go func() {
		s.Start()
	}()

	fmt.Println("I am here")

	runClient(serverAddr)
	s.Stop()
}