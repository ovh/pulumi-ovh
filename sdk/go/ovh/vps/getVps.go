// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vps

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a vps associated with your OVHcloud Account.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/vps"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vps.GetVps(ctx, &vps.GetVpsArgs{
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
func LookupVps(ctx *pulumi.Context, args *LookupVpsArgs, opts ...pulumi.InvokeOption) (*LookupVpsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVpsResult
	err := ctx.Invoke("ovh:Vps/getVps:getVps", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVps.
type LookupVpsArgs struct {
	// The serviceName of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getVps.
type LookupVpsResult struct {
	// The URN of the vps
	VpsURN string `pulumi:"VpsURN"`
	// The OVHcloud cluster the vps is in
	Cluster string `pulumi:"cluster"`
	// The datacenter in which the vps is located
	Datacenter map[string]string `pulumi:"datacenter"`
	// The displayed name in the OVHcloud web admin
	Displayname string `pulumi:"displayname"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The list of IPs addresses attached to the vps
	Ips []string `pulumi:"ips"`
	// The keymap for the ip kvm, valid values "", "fr", "us"
	Keymap string `pulumi:"keymap"`
	// The amount of memory in MB of the vps.
	Memory int `pulumi:"memory"`
	// A dict describing the type of vps.
	Model map[string]string `pulumi:"model"`
	Name  string            `pulumi:"name"`
	// The source of the boot kernel
	Netbootmode string `pulumi:"netbootmode"`
	// The type of offer (ssd, cloud, classic)
	Offertype   string `pulumi:"offertype"`
	ServiceName string `pulumi:"serviceName"`
	// A boolean to indicate if OVHcloud SLA monitoring is active.
	Slamonitoring bool `pulumi:"slamonitoring"`
	// The state of the vps
	State string `pulumi:"state"`
	// The type of server
	Type string `pulumi:"type"`
	// The number of vcore of the vps
	Vcore int `pulumi:"vcore"`
	// The OVHcloud zone where the vps is
	Zone string `pulumi:"zone"`
}

func LookupVpsOutput(ctx *pulumi.Context, args LookupVpsOutputArgs, opts ...pulumi.InvokeOption) LookupVpsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVpsResultOutput, error) {
			args := v.(LookupVpsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Vps/getVps:getVps", args, LookupVpsResultOutput{}, options).(LookupVpsResultOutput), nil
		}).(LookupVpsResultOutput)
}

// A collection of arguments for invoking getVps.
type LookupVpsOutputArgs struct {
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupVpsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpsArgs)(nil)).Elem()
}

// A collection of values returned by getVps.
type LookupVpsResultOutput struct{ *pulumi.OutputState }

func (LookupVpsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpsResult)(nil)).Elem()
}

func (o LookupVpsResultOutput) ToLookupVpsResultOutput() LookupVpsResultOutput {
	return o
}

func (o LookupVpsResultOutput) ToLookupVpsResultOutputWithContext(ctx context.Context) LookupVpsResultOutput {
	return o
}

// The URN of the vps
func (o LookupVpsResultOutput) VpsURN() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.VpsURN }).(pulumi.StringOutput)
}

// The OVHcloud cluster the vps is in
func (o LookupVpsResultOutput) Cluster() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Cluster }).(pulumi.StringOutput)
}

// The datacenter in which the vps is located
func (o LookupVpsResultOutput) Datacenter() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpsResult) map[string]string { return v.Datacenter }).(pulumi.StringMapOutput)
}

// The displayed name in the OVHcloud web admin
func (o LookupVpsResultOutput) Displayname() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Displayname }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVpsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The list of IPs addresses attached to the vps
func (o LookupVpsResultOutput) Ips() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVpsResult) []string { return v.Ips }).(pulumi.StringArrayOutput)
}

// The keymap for the ip kvm, valid values "", "fr", "us"
func (o LookupVpsResultOutput) Keymap() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Keymap }).(pulumi.StringOutput)
}

// The amount of memory in MB of the vps.
func (o LookupVpsResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpsResult) int { return v.Memory }).(pulumi.IntOutput)
}

// A dict describing the type of vps.
func (o LookupVpsResultOutput) Model() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVpsResult) map[string]string { return v.Model }).(pulumi.StringMapOutput)
}

func (o LookupVpsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Name }).(pulumi.StringOutput)
}

// The source of the boot kernel
func (o LookupVpsResultOutput) Netbootmode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Netbootmode }).(pulumi.StringOutput)
}

// The type of offer (ssd, cloud, classic)
func (o LookupVpsResultOutput) Offertype() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Offertype }).(pulumi.StringOutput)
}

func (o LookupVpsResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// A boolean to indicate if OVHcloud SLA monitoring is active.
func (o LookupVpsResultOutput) Slamonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVpsResult) bool { return v.Slamonitoring }).(pulumi.BoolOutput)
}

// The state of the vps
func (o LookupVpsResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.State }).(pulumi.StringOutput)
}

// The type of server
func (o LookupVpsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Type }).(pulumi.StringOutput)
}

// The number of vcore of the vps
func (o LookupVpsResultOutput) Vcore() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpsResult) int { return v.Vcore }).(pulumi.IntOutput)
}

// The OVHcloud zone where the vps is
func (o LookupVpsResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpsResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpsResultOutput{})
}
