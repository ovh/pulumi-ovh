// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package iploadbalancing

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get the details of Vrack network available for your IPLoadbalancer associated with your OVHcloud account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/iploadbalancing"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := iploadbalancing.GetVrackNetwork(ctx, &iploadbalancing.GetVrackNetworkArgs{
//				ServiceName:    "XXXXXX",
//				VrackNetworkId: "yyy",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVrackNetwork(ctx *pulumi.Context, args *LookupVrackNetworkArgs, opts ...pulumi.InvokeOption) (*LookupVrackNetworkResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVrackNetworkResult
	err := ctx.Invoke("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVrackNetwork.
type LookupVrackNetworkArgs struct {
	// The internal name of your IP load balancing
	ServiceName string `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network
	VrackNetworkId int `pulumi:"vrackNetworkId"`
}

// A collection of values returned by getVrackNetwork.
type LookupVrackNetworkResult struct {
	// Human readable name for your vrack network
	DisplayName string `pulumi:"displayName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
	NatIp       string `pulumi:"natIp"`
	ServiceName string `pulumi:"serviceName"`
	// IP block of the private network in the vRack
	Subnet string `pulumi:"subnet"`
	// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
	Vlan           int `pulumi:"vlan"`
	VrackNetworkId int `pulumi:"vrackNetworkId"`
}

func LookupVrackNetworkOutput(ctx *pulumi.Context, args LookupVrackNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupVrackNetworkResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVrackNetworkResultOutput, error) {
			args := v.(LookupVrackNetworkArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:IpLoadBalancing/getVrackNetwork:getVrackNetwork", args, LookupVrackNetworkResultOutput{}, options).(LookupVrackNetworkResultOutput), nil
		}).(LookupVrackNetworkResultOutput)
}

// A collection of arguments for invoking getVrackNetwork.
type LookupVrackNetworkOutputArgs struct {
	// The internal name of your IP load balancing
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Internal Load Balancer identifier of the vRack private network
	VrackNetworkId pulumi.IntInput `pulumi:"vrackNetworkId"`
}

func (LookupVrackNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVrackNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getVrackNetwork.
type LookupVrackNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupVrackNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVrackNetworkResult)(nil)).Elem()
}

func (o LookupVrackNetworkResultOutput) ToLookupVrackNetworkResultOutput() LookupVrackNetworkResultOutput {
	return o
}

func (o LookupVrackNetworkResultOutput) ToLookupVrackNetworkResultOutputWithContext(ctx context.Context) LookupVrackNetworkResultOutput {
	return o
}

// Human readable name for your vrack network
func (o LookupVrackNetworkResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVrackNetworkResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) string { return v.Id }).(pulumi.StringOutput)
}

// An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
func (o LookupVrackNetworkResultOutput) NatIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) string { return v.NatIp }).(pulumi.StringOutput)
}

func (o LookupVrackNetworkResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// IP block of the private network in the vRack
func (o LookupVrackNetworkResultOutput) Subnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) string { return v.Subnet }).(pulumi.StringOutput)
}

// VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
func (o LookupVrackNetworkResultOutput) Vlan() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) int { return v.Vlan }).(pulumi.IntOutput)
}

func (o LookupVrackNetworkResultOutput) VrackNetworkId() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVrackNetworkResult) int { return v.VrackNetworkId }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVrackNetworkResultOutput{})
}
