package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/chuan-yu/go-thrift-mock/resources/gen-go/resources"
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

	// create and start mock server
	serverAddr := ":8888"
	s := server.MustNewMockServer(serverAddr)

	go func() {
		if err := s.Start(); err != nil {
			fmt.Printf("failed to start server: %s", err.Error())
			return
		}
	}()

	time.Sleep(1 * time.Second)

	// mock success response
	result := resources.Response{
		Code: 200,
		ResponseMsg: "mock message",
	}
	expectedReturn := server.ExpectedReturn{
		Response: &result,
	}
	s.SetExpectedReturn("sayHello", expectedReturn)
	if err := runClient(serverAddr); err != nil {
		panic("failed to run client: " + err.Error())
	}

	// mock error response
	expectedReturn = server.ExpectedReturn{
		Err: errors.New("mock error"),
	}
	s.SetExpectedReturn("sayHello", expectedReturn)
	if err := runClient(serverAddr); err != nil {
		panic("failed to run client: " + err.Error())
	}

	// stop the mock server
	s.Stop()
}