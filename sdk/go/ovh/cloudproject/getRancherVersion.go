// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetRancherVersion(ctx *pulumi.Context, args *GetRancherVersionArgs, opts ...pulumi.InvokeOption) (*GetRancherVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRancherVersionResult
	err := ctx.Invoke("ovh:CloudProject/getRancherVersion:getRancherVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRancherVersion.
type GetRancherVersionArgs struct {
	ProjectId string `pulumi:"projectId"`
}

// A collection of values returned by getRancherVersion.
type GetRancherVersionResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id        string                     `pulumi:"id"`
	ProjectId string                     `pulumi:"projectId"`
	Versions  []GetRancherVersionVersion `pulumi:"versions"`
}

func GetRancherVersionOutput(ctx *pulumi.Context, args GetRancherVersionOutputArgs, opts ...pulumi.InvokeOption) GetRancherVersionResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRancherVersionResultOutput, error) {
			args := v.(GetRancherVersionArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getRancherVersion:getRancherVersion", args, GetRancherVersionResultOutput{}, options).(GetRancherVersionResultOutput), nil
		}).(GetRancherVersionResultOutput)
}

// A collection of arguments for invoking getRancherVersion.
type GetRancherVersionOutputArgs struct {
	ProjectId pulumi.StringInput `pulumi:"projectId"`
}

func (GetRancherVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRancherVersionArgs)(nil)).Elem()
}

// A collection of values returned by getRancherVersion.
type GetRancherVersionResultOutput struct{ *pulumi.OutputState }

func (GetRancherVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRancherVersionResult)(nil)).Elem()
}

func (o GetRancherVersionResultOutput) ToGetRancherVersionResultOutput() GetRancherVersionResultOutput {
	return o
}

func (o GetRancherVersionResultOutput) ToGetRancherVersionResultOutputWithContext(ctx context.Context) GetRancherVersionResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetRancherVersionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRancherVersionResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetRancherVersionResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRancherVersionResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

func (o GetRancherVersionResultOutput) Versions() GetRancherVersionVersionArrayOutput {
	return o.ApplyT(func(v GetRancherVersionResult) []GetRancherVersionVersion { return v.Versions }).(GetRancherVersionVersionArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRancherVersionResultOutput{})
}
