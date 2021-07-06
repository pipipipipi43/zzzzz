// Code generated by protoc-gen-go-client. DO NOT EDIT.
// Sources: alert.proto

package client

import (
	fmt "fmt"
	servicehub "github.com/erda-project/erda-infra/base/servicehub"
	grpc "github.com/erda-project/erda-infra/pkg/transport/grpc"
	pb "github.com/erda-project/erda-proto-go/core/monitor/alert/pb"
	grpc1 "google.golang.org/grpc"
	reflect "reflect"
	strings "strings"
)

var dependencies = []string{
	"grpc-client@erda.core.monitor.alert",
	"grpc-client",
}

// +provider
type provider struct {
	client Client
}

func (p *provider) Init(ctx servicehub.Context) error {
	var conn grpc.ClientConnInterface
	for _, dep := range dependencies {
		c, ok := ctx.Service(dep).(grpc.ClientConnInterface)
		if ok {
			conn = c
			break
		}
	}
	if conn == nil {
		return fmt.Errorf("not found connector in (%s)", strings.Join(dependencies, ", "))
	}
	p.client = New(conn)
	return nil
}

var (
	clientsType            = reflect.TypeOf((*Client)(nil)).Elem()
	alertServiceClientType = reflect.TypeOf((*pb.AlertServiceClient)(nil)).Elem()
	alertServiceServerType = reflect.TypeOf((*pb.AlertServiceServer)(nil)).Elem()
)

func (p *provider) Provide(ctx servicehub.DependencyContext, args ...interface{}) interface{} {
	var opts []grpc1.CallOption
	for _, arg := range args {
		if opt, ok := arg.(grpc1.CallOption); ok {
			opts = append(opts, opt)
		}
	}
	switch ctx.Service() {
	case "erda.core.monitor.alert-client":
		return p.client
	case "erda.core.monitor.alert.AlertService":
		return &alertServiceWrapper{client: p.client.AlertService(), opts: opts}
	case "erda.core.monitor.alert.AlertService.client":
		return p.client.AlertService()
	}
	switch ctx.Type() {
	case clientsType:
		return p.client
	case alertServiceClientType:
		return p.client.AlertService()
	case alertServiceServerType:
		return &alertServiceWrapper{client: p.client.AlertService(), opts: opts}
	}
	return p
}

func init() {
	servicehub.Register("erda.core.monitor.alert-client", &servicehub.Spec{
		Services: []string{
			"erda.core.monitor.alert.AlertService",
			"erda.core.monitor.alert-client",
		},
		Types: []reflect.Type{
			clientsType,
			// client types
			alertServiceClientType,
			// server types
			alertServiceServerType,
		},
		OptionalDependencies: dependencies,
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}