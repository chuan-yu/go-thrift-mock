package client

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
	"github.com/chuan-yu/go-thrift-mock/resouces/gen-go/resources"
)

type HelloServiceClient struct {
	c thrift.TClient
}

func NewHelloServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *HelloServiceClient {
	return &HelloServiceClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewHelloServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *HelloServiceClient {
	return &HelloServiceClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewHelloServiceClient(c thrift.TClient) *HelloServiceClient {
	return &HelloServiceClient{
		c: c,
	}
}

func (p *HelloServiceClient) Client_() thrift.TClient {
	return p.c
}

// Parameters:
//  - Request
func (p *HelloServiceClient) SayHello(ctx context.Context, request *Request) (r *resources.Response, err error) {
	var _args0 HelloServiceSayHelloArgs
	_args0.Request = request
	var _result1 resources.HelloServiceSayHelloResult
	if err = p.Client_().Call(ctx, "sayHello", &_args0, &_result1); err != nil {
		return
	}
	return _result1.GetSuccess(), nil
}

type HelloServiceSayHelloArgs struct {
	Request *Request `thrift:"request,1,required" db:"request" json:"request"`
}

func NewHelloServiceSayHelloArgs() *HelloServiceSayHelloArgs {
	return &HelloServiceSayHelloArgs{}
}

var HelloServiceSayHelloArgs_Request_DEFAULT *Request

func (p *HelloServiceSayHelloArgs) GetRequest() *Request {
	if !p.IsSetRequest() {
		return HelloServiceSayHelloArgs_Request_DEFAULT
	}
	return p.Request
}
func (p *HelloServiceSayHelloArgs) IsSetRequest() bool {
	return p.Request != nil
}

func (p *HelloServiceSayHelloArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	var issetRequest bool = false

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
				issetRequest = true
			} else {
				if err := iprot.Skip(fieldTypeId); err != nil {
					return err
				}
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	if !issetRequest {
		return thrift.NewTProtocolExceptionWithType(thrift.INVALID_DATA, fmt.Errorf("Required field Request is not set"))
	}
	return nil
}

func (p *HelloServiceSayHelloArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Request = &Request{}
	if err := p.Request.Read(iprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error reading struct: ", p.Request), err)
	}
	return nil
}

func (p *HelloServiceSayHelloArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("sayHello_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField1(oprot); err != nil {
			return err
		}
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *HelloServiceSayHelloArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("request", thrift.STRUCT, 1); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:request: ", p), err)
	}
	if err := p.Request.Write(oprot); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Request), err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write field end error 1:request: ", p), err)
	}
	return err
}

func (p *HelloServiceSayHelloArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloServiceSayHelloArgs(%+v)", *p)
}

