package Response

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"log"
	"golang.org/x/net/context"
	"github.com/lampard1014/aphro/CommonBiz/Response/PB"
	"github.com/golang/protobuf/ptypes/any"
)



func NewCommonBizResponse(code int64, message string,resultMsg proto.Message )(*Aphro_CommonBiz.Response,error) {
	     any, err := MarshalAny(resultMsg)
	     r := &Aphro_CommonBiz.Response{Code:code,Message:message,Result:any}
	     return r,err
}

func MarshalAny(protoMsg proto.Message)(*any.Any, error) {
	any, err := ptypes.MarshalAny(protoMsg)
	return any,err
}

func UnmarshalAny(any *any.Any, pb proto.Message)(error) {
	return ptypes.UnmarshalAny(any,pb)
}


func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("before handling. Info: %+v", info)
	resp, err := handler(ctx, req)

	v,_ := resp.(proto.Message)

	var code int64 = 0
	if err != nil{
		code = 1
	}
	var message string
	if err != nil {
		message = err.Error()
	}

	x,err := NewCommonBizResponse(code, message, v)

	log.Printf("after handling. resp: %+v", x)
	return x, err
}
// StreamServerInterceptor is a gRPC server-side interceptor that provides Prometheus monitoring for Streaming RPCs.
func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("before handling. Info: %+v", info)
	err := handler(srv, ss)
	log.Printf("after handling. err: %v", err)
	return err
}