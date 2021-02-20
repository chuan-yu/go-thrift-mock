package processor

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

type handler struct{}

// MockProcessorFunction is a mock thrift processor function. It implements thrift's
// TProcessorFunction
type MockProcessorFunction struct {
	MethodName string
	Result     thrift.TStruct
	Err        error
}

func (f *MockProcessorFunction) Process(ctx context.Context, seqID int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	if f.Err != nil {
		x := thrift.NewTApplicationException(
			thrift.INTERNAL_ERROR,
			fmt.Sprintf("Internal error processing %s : %s", f.MethodName, f.Err),
		)
		oprot.WriteMessageBegin(f.MethodName, thrift.EXCEPTION, seqID)
		x.Write(oprot)
		oprot.Flush(ctx)
		return true, f.Err
	}

	var err2 error
	if err2 = oprot.WriteMessageBegin(f.MethodName, thrift.REPLY, seqID); err2 != nil {
		err = err2
	}
	if err2 = f.Result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, nil
}

type MockResult struct {
	funcName string
	Success  thrift.TStruct `thrift:"success,0" db:"success" json:"success,omitempty"`
}

func NewMockResult(funcName string, success thrift.TStruct) *MockResult {
	return &MockResult{funcName: funcName, Success: success}
}

var HelloServiceSayHelloResult_Success_DEFAULT thrift.TStruct

func (p *MockResult) GetSuccess() thrift.TStruct {
	if !p.IsSetSuccess() {
		return HelloServiceSayHelloResult_Success_DEFAULT
	}
	return p.Success
}

func (p *MockResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MockResult) Read(iprot thrift.TProtocol) error {

	return nil
}

func (p *MockResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin(p.funcName + "_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if p != nil {
		if err := p.writeField0(oprot); err != nil {
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

func (p *MockResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field begin error 0:success: ", p), err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T error writing struct: ", p.Success), err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return thrift.PrependError(fmt.Sprintf("%T write field end error 0:success: ", p), err)
		}
	}
	return err
}

func (p *MockResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("HelloServiceSayHelloResult(%+v)", *p)
}

// MockProcessor is a mock thrift processor. It implementes thrift's TProcessor interface.
type MockProcessor struct {
	processorFuncMap map[string]thrift.TProcessorFunction
}

func NewMockProcessor() *MockProcessor {
	p := &MockProcessor{}
	p.processorFuncMap = make(map[string]thrift.TProcessorFunction)
	return p
}

// AddToProcessorMap adds a processor function to the processor's map which maps
// a string key to a thrift processor function
func (p *MockProcessor) AddToProcessorMap(key string, processorFunc thrift.TProcessorFunction) {
	p.processorFuncMap[key] = processorFunc
}

// GetProcessorFunction gets the thrift processor function by its string key
func (p *MockProcessor) GetProcessorFunction(key string) (processorFunc thrift.TProcessorFunction, ok bool) {
	processorFunc, ok = p.processorFuncMap[key]
	return
}

// Process operates upon an input stream and writes to a output stream
func (p *MockProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqID, err := iprot.ReadMessageBegin()

	if err != nil {
		return false, err
	}

	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqID, iprot, oprot)
	}

	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x3 := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqID)
	x3.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x3
}
