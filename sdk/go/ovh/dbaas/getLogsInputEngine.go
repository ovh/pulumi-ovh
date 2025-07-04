// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a DBaas logs input engine.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.GetLogsInputEngine(ctx, &dbaas.GetLogsInputEngineArgs{
//				ServiceName:  "ldp-xx-xxxxx",
//				Name:         pulumi.StringRef("logstash"),
//				Version:      pulumi.StringRef("6.8"),
//				IsDeprecated: pulumi.BoolRef(true),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLogsInputEngine(ctx *pulumi.Context, args *GetLogsInputEngineArgs, opts ...pulumi.InvokeOption) (*GetLogsInputEngineResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLogsInputEngineResult
	err := ctx.Invoke("ovh:Dbaas/getLogsInputEngine:getLogsInputEngine", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLogsInputEngine.
type GetLogsInputEngineArgs struct {
	// Indicates if engine will soon not be supported.
	IsDeprecated *bool `pulumi:"isDeprecated"`
	// The name of the logs input engine.
	Name *string `pulumi:"name"`
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName string `pulumi:"serviceName"`
	// Software version
	Version *string `pulumi:"version"`
}

// A collection of values returned by getLogsInputEngine.
type GetLogsInputEngineResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id           string `pulumi:"id"`
	IsDeprecated bool   `pulumi:"isDeprecated"`
	Name         string `pulumi:"name"`
	ServiceName  string `pulumi:"serviceName"`
	Version      string `pulumi:"version"`
}

func GetLogsInputEngineOutput(ctx *pulumi.Context, args GetLogsInputEngineOutputArgs, opts ...pulumi.InvokeOption) GetLogsInputEngineResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetLogsInputEngineResultOutput, error) {
			args := v.(GetLogsInputEngineArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dbaas/getLogsInputEngine:getLogsInputEngine", args, GetLogsInputEngineResultOutput{}, options).(GetLogsInputEngineResultOutput), nil
		}).(GetLogsInputEngineResultOutput)
}

// A collection of arguments for invoking getLogsInputEngine.
type GetLogsInputEngineOutputArgs struct {
	// Indicates if engine will soon not be supported.
	IsDeprecated pulumi.BoolPtrInput `pulumi:"isDeprecated"`
	// The name of the logs input engine.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The service name. It's the ID of your Logs Data Platform instance.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Software version
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (GetLogsInputEngineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsInputEngineArgs)(nil)).Elem()
}

// A collection of values returned by getLogsInputEngine.
type GetLogsInputEngineResultOutput struct{ *pulumi.OutputState }

func (GetLogsInputEngineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLogsInputEngineResult)(nil)).Elem()
}

func (o GetLogsInputEngineResultOutput) ToGetLogsInputEngineResultOutput() GetLogsInputEngineResultOutput {
	return o
}

func (o GetLogsInputEngineResultOutput) ToGetLogsInputEngineResultOutputWithContext(ctx context.Context) GetLogsInputEngineResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLogsInputEngineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsInputEngineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLogsInputEngineResultOutput) IsDeprecated() pulumi.BoolOutput {
	return o.ApplyT(func(v GetLogsInputEngineResult) bool { return v.IsDeprecated }).(pulumi.BoolOutput)
}

func (o GetLogsInputEngineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsInputEngineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetLogsInputEngineResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsInputEngineResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func (o GetLogsInputEngineResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v GetLogsInputEngineResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLogsInputEngineResultOutput{})
}
