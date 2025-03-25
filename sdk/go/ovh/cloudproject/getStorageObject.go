// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get information about an object in a S3™* compatible storage container. \* S3 is a trademark filed by Amazon Technologies,Inc. OVHcloud's service is not sponsored by, endorsed by, or otherwise affiliated with Amazon Technologies,Inc.
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
//			_, err := cloudproject.GetStorageObject(ctx, &cloudproject.GetStorageObjectArgs{
//				Key:         "<object name>",
//				Name:        "<bucket name>",
//				RegionName:  "GRA",
//				ServiceName: "<public cloud project ID>",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupStorageObject(ctx *pulumi.Context, args *LookupStorageObjectArgs, opts ...pulumi.InvokeOption) (*LookupStorageObjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupStorageObjectResult
	err := ctx.Invoke("ovh:CloudProject/getStorageObject:getStorageObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStorageObject.
type LookupStorageObjectArgs struct {
	// Key
	Key string `pulumi:"key"`
	// Name
	Name string `pulumi:"name"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getStorageObject.
type LookupStorageObjectResult struct {
	// ETag
	Etag string `pulumi:"etag"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Whether this object is a delete marker
	IsDeleteMarker bool `pulumi:"isDeleteMarker"`
	// Whether this is the latest version of the object
	IsLatest bool `pulumi:"isLatest"`
	// Key
	Key string `pulumi:"key"`
	// Last modification date
	LastModified string `pulumi:"lastModified"`
	// Name
	Name string `pulumi:"name"`
	// Region name
	RegionName string `pulumi:"regionName"`
	// Service name
	ServiceName string `pulumi:"serviceName"`
	// Size (bytes)
	Size float64 `pulumi:"size"`
	// Storage class
	StorageClass string `pulumi:"storageClass"`
	// Version ID of the object
	VersionId string `pulumi:"versionId"`
}

func LookupStorageObjectOutput(ctx *pulumi.Context, args LookupStorageObjectOutputArgs, opts ...pulumi.InvokeOption) LookupStorageObjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupStorageObjectResultOutput, error) {
			args := v.(LookupStorageObjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getStorageObject:getStorageObject", args, LookupStorageObjectResultOutput{}, options).(LookupStorageObjectResultOutput), nil
		}).(LookupStorageObjectResultOutput)
}

// A collection of arguments for invoking getStorageObject.
type LookupStorageObjectOutputArgs struct {
	// Key
	Key pulumi.StringInput `pulumi:"key"`
	// Name
	Name pulumi.StringInput `pulumi:"name"`
	// Region name
	RegionName pulumi.StringInput `pulumi:"regionName"`
	// Service name
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupStorageObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageObjectArgs)(nil)).Elem()
}

// A collection of values returned by getStorageObject.
type LookupStorageObjectResultOutput struct{ *pulumi.OutputState }

func (LookupStorageObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStorageObjectResult)(nil)).Elem()
}

func (o LookupStorageObjectResultOutput) ToLookupStorageObjectResultOutput() LookupStorageObjectResultOutput {
	return o
}

func (o LookupStorageObjectResultOutput) ToLookupStorageObjectResultOutputWithContext(ctx context.Context) LookupStorageObjectResultOutput {
	return o
}

// ETag
func (o LookupStorageObjectResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStorageObjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Whether this object is a delete marker
func (o LookupStorageObjectResultOutput) IsDeleteMarker() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) bool { return v.IsDeleteMarker }).(pulumi.BoolOutput)
}

// Whether this is the latest version of the object
func (o LookupStorageObjectResultOutput) IsLatest() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) bool { return v.IsLatest }).(pulumi.BoolOutput)
}

// Key
func (o LookupStorageObjectResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.Key }).(pulumi.StringOutput)
}

// Last modification date
func (o LookupStorageObjectResultOutput) LastModified() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.LastModified }).(pulumi.StringOutput)
}

// Name
func (o LookupStorageObjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// Region name
func (o LookupStorageObjectResultOutput) RegionName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.RegionName }).(pulumi.StringOutput)
}

// Service name
func (o LookupStorageObjectResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Size (bytes)
func (o LookupStorageObjectResultOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v LookupStorageObjectResult) float64 { return v.Size }).(pulumi.Float64Output)
}

// Storage class
func (o LookupStorageObjectResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

// Version ID of the object
func (o LookupStorageObjectResultOutput) VersionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStorageObjectResult) string { return v.VersionId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStorageObjectResultOutput{})
}
