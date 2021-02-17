package server

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/chuan-yu/go-thrift-mock/processor"
)

type ExpectedReturn struct {
	Err error
	Result thrift.TStruct
}

type MockServer struct {
	host string
	Server *thrift.TSimpleServer
	processor *processor.MockProcessor
	protocolFactory *thrift.TBinaryProtocolFactory
	transportFactory *thrift.TTransportFactory
}

func MustNewMockServer(host string) *MockServer {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
	transport, err := thrift.NewTServerSocket(host)
	if err != nil {
		panic("failed to create a MockServer instance: " + err.Error())
	}
	p := processor.NewMockProcessor()
	server := thrift.NewTSimpleServer4(p, transport, transportFactory, protocolFactory)
	return &MockServer{
		host: host,
		processor: p,
		protocolFactory: protocolFactory,
		transportFactory: &transportFactory,
		Server: server,
	}
}

func (s *MockServer) Start() (err error) {
	fmt.Printf("starting the simple server... on %s \n", s.host)
	return s.Server.Serve()
}

func (s *MockServer) Stop() {
	s.Server.Stop()
}

func (s *MockServer) SetExpectedReturn(methodName string, expected ExpectedReturn) {
	processFunc := processor.MockProcessorFunction{
		MethodName: methodName,
		Result: expected.Result,
		Err: expected.Err,
	}
	s.processor.AddToProcessorMap(methodName, &processFunc)
}

