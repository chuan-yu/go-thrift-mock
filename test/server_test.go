package test

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/chuan-yu/go-thrift-mock/resources/gen-go/resources"
	"github.com/chuan-yu/go-thrift-mock/server"
	"github.com/stretchr/testify/assert"
)

func runClient(serverAddr string) (*resources.Response, error) {

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	var transport thrift.TTransport
	transport, err := thrift.NewTSocket(serverAddr)

	if err != nil {
		return nil, fmt.Errorf("Error opening socket: %s", err) 
	}

	if transport == nil {
		return nil, fmt.Errorf("Error opening socket, got nil transport. Is server available?")
	}

	transport, err = transportFactory.GetTransport(transport)
	if err != nil {
		return nil, err
	}

	err = transport.Open()
	if err != nil {
		return nil, fmt.Errorf("error running client: %s", err)
	}

	defer transport.Close()

	c := resources.NewHelloServiceClientFactory(transport, protocolFactory)

	resp, err := c.SayHello(context.Background(), &resources.Request{
		Msg: &[]string{"hello"}[0],
	})

	if err != nil {
		return nil, err
	}
	return resp, nil 
}

func TestServer(t *testing.T) {

	// start mock server
	serverAddr := ":8888"
	s := server.MustNewMockServer(serverAddr)

	go func() {
		if err := s.Start(); err != nil {
			fmt.Printf("failed to start server: %s", err.Error())
			return
		}
	}()
	time.Sleep(1 * time.Second)

	// test success response
	expectedReturn := server.ExpectedReturn{
		Response: &resources.Response{
			Code: 200,
			ResponseMsg: "mock message",
		},
	}
	s.SetExpectedReturn("sayHello", expectedReturn)

	resp, err := runClient(serverAddr)
	if assert.NoError(t, err) {
		assert.Equal(t, expectedReturn.Response, resp)
	}

	// test error response
	expectedErrMsg := "mock server"
	expectedReturn = server.ExpectedReturn{
		Err: errors.New(expectedErrMsg),
	}
	s.SetExpectedReturn("sayHello", expectedReturn)
	resp, err = runClient(serverAddr)
	assert.Contains(t, err.Error(), expectedErrMsg)

	// stop the mock server
	s.Stop()
}