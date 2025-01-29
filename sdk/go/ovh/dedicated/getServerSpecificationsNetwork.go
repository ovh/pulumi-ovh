// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the network information about a dedicated server associated with your OVHcloud Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/dedicated"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dedicated.GetServerSpecificationsNetwork(ctx, &dedicated.GetServerSpecificationsNetworkArgs{
//				ServiceName: "myserver",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetServerSpecificationsNetwork(ctx *pulumi.Context, args *GetServerSpecificationsNetworkArgs, opts ...pulumi.InvokeOption) (*GetServerSpecificationsNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetServerSpecificationsNetworkResult
	err := ctx.Invoke("ovh:Dedicated/getServerSpecificationsNetwork:getServerSpecificationsNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getServerSpecificationsNetwork.
type GetServerSpecificationsNetworkArgs struct {
	// The internal name of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getServerSpecificationsNetwork.
type GetServerSpecificationsNetworkResult struct {
	// vrack bandwidth limitation
	Bandwidth GetServerSpecificationsNetworkBandwidth `pulumi:"bandwidth"`
	// Network connection flow rate
	ConnectionVal GetServerSpecificationsNetworkConnectionVal `pulumi:"connectionVal"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// OLA details
	Ola GetServerSpecificationsNetworkOla `pulumi:"ola"`
	// Routing details
	Routing     GetServerSpecificationsNetworkRouting `pulumi:"routing"`
	ServiceName string                                `pulumi:"serviceName"`
	// Switching details
	Switching GetServerSpecificationsNetworkSwitching `pulumi:"switching"`
	// Traffic details
	Traffic GetServerSpecificationsNetworkTraffic `pulumi:"traffic"`
	// VMAC information for this dedicated server
	Vmac GetServerSpecificationsNetworkVmac `pulumi:"vmac"`
	// vRack details
	Vrack GetServerSpecificationsNetworkVrack `pulumi:"vrack"`
}

func GetServerSpecificationsNetworkOutput(ctx *pulumi.Context, args GetServerSpecificationsNetworkOutputArgs, opts ...pulumi.InvokeOption) GetServerSpecificationsNetworkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetServerSpecificationsNetworkResultOutput, error) {
			args := v.(GetServerSpecificationsNetworkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dedicated/getServerSpecificationsNetwork:getServerSpecificationsNetwork", args, GetServerSpecificationsNetworkResultOutput{}, options).(GetServerSpecificationsNetworkResultOutput), nil
		}).(GetServerSpecificationsNetworkResultOutput)
}

// A collection of arguments for invoking getServerSpecificationsNetwork.
type GetServerSpecificationsNetworkOutputArgs struct {
	// The internal name of your dedicated server.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetServerSpecificationsNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerSpecificationsNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getServerSpecificationsNetwork.
type GetServerSpecificationsNetworkResultOutput struct{ *pulumi.OutputState }

func (GetServerSpecificationsNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetServerSpecificationsNetworkResult)(nil)).Elem()
}

func (o GetServerSpecificationsNetworkResultOutput) ToGetServerSpecificationsNetworkResultOutput() GetServerSpecificationsNetworkResultOutput {
	return o
}

func (o GetServerSpecificationsNetworkResultOutput) ToGetServerSpecificationsNetworkResultOutputWithContext(ctx context.Context) GetServerSpecificationsNetworkResultOutput {
	return o
}

// vrack bandwidth limitation
func (o GetServerSpecificationsNetworkResultOutput) Bandwidth() GetServerSpecificationsNetworkBandwidthOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkBandwidth {
		return v.Bandwidth
	}).(GetServerSpecificationsNetworkBandwidthOutput)
}

// Network connection flow rate
func (o GetServerSpecificationsNetworkResultOutput) ConnectionVal() GetServerSpecificationsNetworkConnectionValOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkConnectionVal {
		return v.ConnectionVal
	}).(GetServerSpecificationsNetworkConnectionValOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetServerSpecificationsNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// OLA details
func (o GetServerSpecificationsNetworkResultOutput) Ola() GetServerSpecificationsNetworkOlaOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkOla { return v.Ola }).(GetServerSpecificationsNetworkOlaOutput)
}

// Routing details
func (o GetServerSpecificationsNetworkResultOutput) Routing() GetServerSpecificationsNetworkRoutingOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkRouting { return v.Routing }).(GetServerSpecificationsNetworkRoutingOutput)
}

func (o GetServerSpecificationsNetworkResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Switching details
func (o GetServerSpecificationsNetworkResultOutput) Switching() GetServerSpecificationsNetworkSwitchingOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkSwitching {
		return v.Switching
	}).(GetServerSpecificationsNetworkSwitchingOutput)
}

// Traffic details
func (o GetServerSpecificationsNetworkResultOutput) Traffic() GetServerSpecificationsNetworkTrafficOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkTraffic { return v.Traffic }).(GetServerSpecificationsNetworkTrafficOutput)
}

// VMAC information for this dedicated server
func (o GetServerSpecificationsNetworkResultOutput) Vmac() GetServerSpecificationsNetworkVmacOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkVmac { return v.Vmac }).(GetServerSpecificationsNetworkVmacOutput)
}

// vRack details
func (o GetServerSpecificationsNetworkResultOutput) Vrack() GetServerSpecificationsNetworkVrackOutput {
	return o.ApplyT(func(v GetServerSpecificationsNetworkResult) GetServerSpecificationsNetworkVrack { return v.Vrack }).(GetServerSpecificationsNetworkVrackOutput)
}

func init() {
	pulumi.RegisterOutputType(GetServerSpecificationsNetworkResultOutput{})
}
