// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about an IP service.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/ip"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := ip.GetService(ctx, &ip.GetServiceArgs{
//				ServiceName: "XXXXXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetService(ctx *pulumi.Context, args *GetServiceArgs, opts ...pulumi.InvokeOption) (*GetServiceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServiceResult
	err := ctx.Invoke("ovh:Ip/getService:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getService.
type GetServiceArgs struct {
	// The service name
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getService.
type GetServiceResult struct {
	// can be terminated
	CanBeTerminated bool `pulumi:"canBeTerminated"`
	// country
	Country string `pulumi:"country"`
	// Custom description on your ip
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// ip block
	Ip string `pulumi:"ip"`
	// IP block organisation Id
	OrganisationId string `pulumi:"organisationId"`
	// Routage information
	RoutedTos []GetServiceRoutedTo `pulumi:"routedTos"`
	// Service where ip is routed to
	ServiceName string `pulumi:"serviceName"`
	// Possible values for ip type ( "cdn", "cloud", "dedicated", "failover", "hostedSsl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")
	Type string `pulumi:"type"`
}

func GetServiceOutput(ctx *pulumi.Context, args GetServiceOutputArgs, opts ...pulumi.InvokeOption) GetServiceResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServiceResultOutput, error) {
			args := v.(GetServiceArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Ip/getService:getService", args, GetServiceResultOutput{}, options).(GetServiceResultOutput), nil
		}).(GetServiceResultOutput)
}

// A collection of arguments for invoking getService.
type GetServiceOutputArgs struct {
	// The service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceArgs)(nil)).Elem()
}

// A collection of values returned by getService.
type GetServiceResultOutput struct{ *pulumi.OutputState }

func (GetServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServiceResult)(nil)).Elem()
}

func (o GetServiceResultOutput) ToGetServiceResultOutput() GetServiceResultOutput {
	return o
}

func (o GetServiceResultOutput) ToGetServiceResultOutputWithContext(ctx context.Context) GetServiceResultOutput {
	return o
}

// can be terminated
func (o GetServiceResultOutput) CanBeTerminated() pulumi.BoolOutput {
	return o.ApplyT(func(v GetServiceResult) bool { return v.CanBeTerminated }).(pulumi.BoolOutput)
}

// country
func (o GetServiceResultOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Country }).(pulumi.StringOutput)
}

// Custom description on your ip
func (o GetServiceResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServiceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Id }).(pulumi.StringOutput)
}

// ip block
func (o GetServiceResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Ip }).(pulumi.StringOutput)
}

// IP block organisation Id
func (o GetServiceResultOutput) OrganisationId() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.OrganisationId }).(pulumi.StringOutput)
}

// Routage information
func (o GetServiceResultOutput) RoutedTos() GetServiceRoutedToArrayOutput {
	return o.ApplyT(func(v GetServiceResult) []GetServiceRoutedTo { return v.RoutedTos }).(GetServiceRoutedToArrayOutput)
}

// Service where ip is routed to
func (o GetServiceResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Possible values for ip type ( "cdn", "cloud", "dedicated", "failover", "hostedSsl", "housing", "loadBalancing", "mail", "overthebox", "pcc", "pci", "private", "vpn", "vps", "vrack", "xdsl")
func (o GetServiceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetServiceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServiceResultOutput{})
}
