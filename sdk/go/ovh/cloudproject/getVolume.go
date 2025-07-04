// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about a volume in a public cloud project
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloudproject.GetVolume(ctx, &cloudproject.GetVolumeArgs{
//				RegionName:  "xxx",
//				ServiceName: "yyy",
//				VolumeId:    "aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVolumeResult
	err := ctx.Invoke("ovh:CloudProject/getVolume:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVolume.
type LookupVolumeArgs struct {
	// A valid OVHcloud public cloud region name in which the volume is available. Ex.: "GRA11".
	RegionName string `pulumi:"regionName"`
	// The id of the public cloud project.
	ServiceName string `pulumi:"serviceName"`
	// Volume id to get the informations
	VolumeId string `pulumi:"volumeId"`
}

// A collection of values returned by getVolume.
type LookupVolumeResult struct {
	Id string `pulumi:"id"`
	// The name of the volume (E.g.: "GRA", meaning Gravelines, for region "GRA1")
	Name string `pulumi:"name"`
	// The region name where volume is available
	RegionName string `pulumi:"regionName"`
	// The id of the public cloud project.
	ServiceName string `pulumi:"serviceName"`
	// The size of the volume
	Size float64 `pulumi:"size"`
	// The id of the volume
	VolumeId string `pulumi:"volumeId"`
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVolumeResultOutput, error) {
			args := v.(LookupVolumeArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getVolume:getVolume", args, LookupVolumeResultOutput{}, options).(LookupVolumeResultOutput), nil
		}).(LookupVolumeResultOutput)
}

// A collection of arguments for invoking getVolume.
type LookupVolumeOutputArgs struct {
	// A valid OVHcloud public cloud region name in which the volume is available. Ex.: "GRA11".
	RegionName pulumi.StringInput `pulumi:"regionName"`
	// The id of the public cloud project.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Volume id to get the informations
	VolumeId pulumi.StringInput `pulumi:"volumeId"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}

// A collection of values returned by getVolume.
type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the volume (E.g.: "GRA", meaning Gravelines, for region "GRA1")
func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

// The region name where volume is available
func (o LookupVolumeResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.RegionName }).(pulumi.StringOutput)
}

// The id of the public cloud project.
func (o LookupVolumeResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// The size of the volume
func (o LookupVolumeResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// The id of the volume
func (o LookupVolumeResultOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
