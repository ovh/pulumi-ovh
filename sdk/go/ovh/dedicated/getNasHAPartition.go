// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to retrieve information about a dedicated HA-NAS partition.
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
//			_, err := dedicated.GetNasHAPartition(ctx, &dedicated.GetNasHAPartitionArgs{
//				Name:        "my-zpool-partition",
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
func LookupNasHAPartition(ctx *pulumi.Context, args *LookupNasHAPartitionArgs, opts ...pulumi.InvokeOption) (*LookupNasHAPartitionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNasHAPartitionResult
	err := ctx.Invoke("ovh:Dedicated/getNasHAPartition:getNasHAPartition", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNasHAPartition.
type LookupNasHAPartitionArgs struct {
	// The name of your dedicated HA-NAS partition.
	Name string `pulumi:"name"`
	// The serviceName of your dedicated HA-NAS.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getNasHAPartition.
type LookupNasHAPartitionResult struct {
	// Percentage of partition space used in %
	Capacity int `pulumi:"capacity"`
	// A brief description of the partition
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// one of "NFS", "CIFS" or "NFS_CIFS"
	Protocol    string `pulumi:"protocol"`
	ServiceName string `pulumi:"serviceName"`
	// size of the partition in GB
	Size int `pulumi:"size"`
	// Percentage of partition space used by snapshots in %
	UsedBySnapshots int `pulumi:"usedBySnapshots"`
}

func LookupNasHAPartitionOutput(ctx *pulumi.Context, args LookupNasHAPartitionOutputArgs, opts ...pulumi.InvokeOption) LookupNasHAPartitionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNasHAPartitionResultOutput, error) {
			args := v.(LookupNasHAPartitionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Dedicated/getNasHAPartition:getNasHAPartition", args, LookupNasHAPartitionResultOutput{}, options).(LookupNasHAPartitionResultOutput), nil
		}).(LookupNasHAPartitionResultOutput)
}

// A collection of arguments for invoking getNasHAPartition.
type LookupNasHAPartitionOutputArgs struct {
	// The name of your dedicated HA-NAS partition.
	Name pulumi.StringInput `pulumi:"name"`
	// The serviceName of your dedicated HA-NAS.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupNasHAPartitionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNasHAPartitionArgs)(nil)).Elem()
}

// A collection of values returned by getNasHAPartition.
type LookupNasHAPartitionResultOutput struct{ *pulumi.OutputState }

func (LookupNasHAPartitionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNasHAPartitionResult)(nil)).Elem()
}

func (o LookupNasHAPartitionResultOutput) ToLookupNasHAPartitionResultOutput() LookupNasHAPartitionResultOutput {
	return o
}

func (o LookupNasHAPartitionResultOutput) ToLookupNasHAPartitionResultOutputWithContext(ctx context.Context) LookupNasHAPartitionResultOutput {
	return o
}

// Percentage of partition space used in %
func (o LookupNasHAPartitionResultOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) int { return v.Capacity }).(pulumi.IntOutput)
}

// A brief description of the partition
func (o LookupNasHAPartitionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNasHAPartitionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNasHAPartitionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) string { return v.Name }).(pulumi.StringOutput)
}

// one of "NFS", "CIFS" or "NFS_CIFS"
func (o LookupNasHAPartitionResultOutput) Protocol() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) string { return v.Protocol }).(pulumi.StringOutput)
}

func (o LookupNasHAPartitionResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// size of the partition in GB
func (o LookupNasHAPartitionResultOutput) Size() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) int { return v.Size }).(pulumi.IntOutput)
}

// Percentage of partition space used by snapshots in %
func (o LookupNasHAPartitionResultOutput) UsedBySnapshots() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNasHAPartitionResult) int { return v.UsedBySnapshots }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNasHAPartitionResultOutput{})
}
