package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/chuan-yu/go-thrift-mock/client"
)

func handleClient(c *client.HelloServiceClient) (err error) {
	resp, err := c.SayHello(context.Background(), &client.Request{
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

func runClient() error {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	addr := "127.0.0.1:8888"
	var transport thrift.TTransport
	transport, err := thrift.NewTSocket(addr)

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

	c := client.NewHelloServiceClientFactory(transport, protocolFactory)
	return handleClient(c)
}

func main() {
	err := runClient()
	if err != nil {
		panic(err.Error())
	}
	
}