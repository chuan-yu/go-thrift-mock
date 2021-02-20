package client

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

type Request struct {
	Msg *string `thrift:"msg,1" db:"msg" json:"msg,omitempty"`
}

func NewRequest() *Request {
	return &Request{}
}

var Request_Msg_DEFAULT string

func (p *Request) GetMsg() string {
	if !p.IsSetMsg() {
		return Request_Msg_DEFAULT
	}
	return *p.Msg
}
func (p *Request) IsSetMsg() bool {
	return p.Msg != nil
}

func (p *Request) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

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
			if fieldTypeId == thrift.STRING {
				if err := p.ReadField1(iprot); err != nil {
					return err
				}
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
	return nil
}

func (p *Request) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return thrift.PrependError("error reading field 1: ", err)
	} else {
		p.Msg = &v
	}
	return nil
}

func (p *Request) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Request"); err != nil {
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

func (p *Request) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetMsg() {
		if err := oprot.WriteFieldBegin("msg", thrift.STRING, 1); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 1:msg: ", p), err)
		}
		if err := oprot.WriteString(string(*p.Msg)); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T.msg (1) field write error: ", p), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 1:msg: ", p), err)
		}
	}
	return err
}

func (p *Request) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Request(%+v)", *p)
}