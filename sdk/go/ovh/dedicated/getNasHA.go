// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a dedicated HA-NAS.
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
//			_, err := dedicated.GetNasHA(ctx, &dedicated.GetNasHAArgs{
//				ServiceName: "zpool-12345",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNasHA(ctx *pulumi.Context, args *GetNasHAArgs, opts ...pulumi.InvokeOption) (*GetNasHAResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNasHAResult
	err := ctx.Invoke("ovh:Dedicated/getNasHA:getNasHA", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNasHA.
type GetNasHAArgs struct {
	// The serviceName of your dedicated HA-NAS.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getNasHA.
type GetNasHAResult struct {
	// the URN of the HA-NAS instance
	NasHAURN string `pulumi:"NasHAURN"`
	// True, if partition creation is allowed on this HA-NAS
	CanCreatePartition bool `pulumi:"canCreatePartition"`
	// The name you give to the HA-NAS
	CustomName string `pulumi:"customName"`
	// area of HA-NAS
	Datacenter string `pulumi:"datacenter"`
	// the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
	DiskType string `pulumi:"diskType"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Access IP of HA-NAS
	Ip string `pulumi:"ip"`
	// Send an email to customer if any issue is detected
	Monitored bool `pulumi:"monitored"`
	// the list of the HA-NAS partitions name
	PartitionsLists []string `pulumi:"partitionsLists"`
	// The storage service name
	ServiceName string `pulumi:"serviceName"`
	// percentage of HA-NAS space used in %
	ZpoolCapacity float64 `pulumi:"zpoolCapacity"`
	// the size of the HA-NAS in GB
	ZpoolSize float64 `pulumi:"zpoolSize"`
}

func GetNasHAOutput(ctx *pulumi.Context, args GetNasHAOutputArgs, opts ...pulumi.InvokeOption) GetNasHAResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNasHAResultOutput, error) {
			args := v.(GetNasHAArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dedicated/getNasHA:getNasHA", args, GetNasHAResultOutput{}, options).(GetNasHAResultOutput), nil
		}).(GetNasHAResultOutput)
}

// A collection of arguments for invoking getNasHA.
type GetNasHAOutputArgs struct {
	// The serviceName of your dedicated HA-NAS.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetNasHAOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNasHAArgs)(nil)).Elem()
}

// A collection of values returned by getNasHA.
type GetNasHAResultOutput struct{ *pulumi.OutputState }

func (GetNasHAResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNasHAResult)(nil)).Elem()
}

func (o GetNasHAResultOutput) ToGetNasHAResultOutput() GetNasHAResultOutput {
	return o
}

func (o GetNasHAResultOutput) ToGetNasHAResultOutputWithContext(ctx context.Context) GetNasHAResultOutput {
	return o
}

// the URN of the HA-NAS instance
func (o GetNasHAResultOutput) NasHAURN() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.NasHAURN }).(pulumi.StringOutput)
}

// True, if partition creation is allowed on this HA-NAS
func (o GetNasHAResultOutput) CanCreatePartition() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNasHAResult) bool { return v.CanCreatePartition }).(pulumi.BoolOutput)
}

// The name you give to the HA-NAS
func (o GetNasHAResultOutput) CustomName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.CustomName }).(pulumi.StringOutput)
}

// area of HA-NAS
func (o GetNasHAResultOutput) Datacenter() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.Datacenter }).(pulumi.StringOutput)
}

// the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
func (o GetNasHAResultOutput) DiskType() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.DiskType }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNasHAResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.Id }).(pulumi.StringOutput)
}

// Access IP of HA-NAS
func (o GetNasHAResultOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.Ip }).(pulumi.StringOutput)
}

// Send an email to customer if any issue is detected
func (o GetNasHAResultOutput) Monitored() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNasHAResult) bool { return v.Monitored }).(pulumi.BoolOutput)
}

// the list of the HA-NAS partitions name
func (o GetNasHAResultOutput) PartitionsLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNasHAResult) []string { return v.PartitionsLists }).(pulumi.StringArrayOutput)
}

// The storage service name
func (o GetNasHAResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetNasHAResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// percentage of HA-NAS space used in %
func (o GetNasHAResultOutput) ZpoolCapacity() pulumi.Float64Output {
	return o.ApplyT(func(v GetNasHAResult) float64 { return v.ZpoolCapacity }).(pulumi.Float64Output)
}

// the size of the HA-NAS in GB
func (o GetNasHAResultOutput) ZpoolSize() pulumi.Float64Output {
	return o.ApplyT(func(v GetNasHAResult) float64 { return v.ZpoolSize }).(pulumi.Float64Output)
}

func init() {
	pulumi.RegisterOutputType(GetNasHAResultOutput{})
}
