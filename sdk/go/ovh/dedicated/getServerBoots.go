// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the list of compatible netboots for a dedicated server associated with your OVHcloud Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dedicated"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dedicated.GetServerBoots(ctx, &dedicated.GetServerBootsArgs{
//				ServiceName: "myserver",
//				BootType:    pulumi.StringRef("harddisk"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServerBoots(ctx *pulumi.Context, args *GetServerBootsArgs, opts ...pulumi.InvokeOption) (*GetServerBootsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerBootsResult
	err := ctx.Invoke("ovh:Dedicated/getServerBoots:getServerBoots", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerBoots.
type GetServerBootsArgs struct {
	// Filter the value of bootType property (harddisk, rescue, internal, network)
	BootType *string `pulumi:"bootType"`
	// Filter the value of kernel property (iPXE script name)
	Kernel *string `pulumi:"kernel"`
	// The internal name of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getServerBoots.
type GetServerBootsResult struct {
	BootType *string `pulumi:"bootType"`
	// The provider-assigned unique ID for this managed resource.
	Id     string  `pulumi:"id"`
	Kernel *string `pulumi:"kernel"`
	// The list of dedicated server netboots.
	Results     []int  `pulumi:"results"`
	ServiceName string `pulumi:"serviceName"`
}

func GetServerBootsOutput(ctx *pulumi.Context, args GetServerBootsOutputArgs, opts ...pulumi.InvokeOption) GetServerBootsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServerBootsResultOutput, error) {
			args := v.(GetServerBootsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dedicated/getServerBoots:getServerBoots", args, GetServerBootsResultOutput{}, options).(GetServerBootsResultOutput), nil
		}).(GetServerBootsResultOutput)
}

// A collection of arguments for invoking getServerBoots.
type GetServerBootsOutputArgs struct {
	// Filter the value of bootType property (harddisk, rescue, internal, network)
	BootType pulumi.StringPtrInput `pulumi:"bootType"`
	// Filter the value of kernel property (iPXE script name)
	Kernel pulumi.StringPtrInput `pulumi:"kernel"`
	// The internal name of your dedicated server.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetServerBootsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerBootsArgs)(nil)).Elem()
}

// A collection of values returned by getServerBoots.
type GetServerBootsResultOutput struct{ *pulumi.OutputState }

func (GetServerBootsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerBootsResult)(nil)).Elem()
}

func (o GetServerBootsResultOutput) ToGetServerBootsResultOutput() GetServerBootsResultOutput {
	return o
}

func (o GetServerBootsResultOutput) ToGetServerBootsResultOutputWithContext(ctx context.Context) GetServerBootsResultOutput {
	return o
}

func (o GetServerBootsResultOutput) BootType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerBootsResult) *string { return v.BootType }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerBootsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerBootsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetServerBootsResultOutput) Kernel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetServerBootsResult) *string { return v.Kernel }).(pulumi.StringPtrOutput)
}

// The list of dedicated server netboots.
func (o GetServerBootsResultOutput) Results() pulumi.IntArrayOutput {
	return o.ApplyT(func(v GetServerBootsResult) []int { return v.Results }).(pulumi.IntArrayOutput)
}

func (o GetServerBootsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerBootsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerBootsResultOutput{})
}
