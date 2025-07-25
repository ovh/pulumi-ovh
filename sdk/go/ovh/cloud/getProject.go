// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the details of a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/cloud"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := cloud.GetProject(ctx, &cloud.GetProjectArgs{
//				ServiceName: "XXX",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetProject(ctx *pulumi.Context, args *GetProjectArgs, opts ...pulumi.InvokeOption) (*GetProjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetProjectResult
	err := ctx.Invoke("ovh:Cloud/getProject:getProject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getProject.
type GetProjectArgs struct {
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getProject.
type GetProjectResult struct {
	// Project access
	Access string `pulumi:"access"`
	// Project creation date
	CreationDate string `pulumi:"creationDate"`
	// Description of your project
	Description string `pulumi:"description"`
	// Expiration date of your project. After this date, your project will be deleted
	Expiration string `pulumi:"expiration"`
	// IAM resource information
	Iam GetProjectIam `pulumi:"iam"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Manual quota prevent automatic quota upgrade
	ManualQuota bool `pulumi:"manualQuota"`
	// Project order ID
	OrderId float64 `pulumi:"orderId"`
	// Order plan code
	PlanCode string `pulumi:"planCode"`
	// Project ID
	ProjectId string `pulumi:"projectId"`
	// Project name
	ProjectName string `pulumi:"projectName"`
	// ID of the public cloud project
	ServiceName string `pulumi:"serviceName"`
	// Current status
	Status string `pulumi:"status"`
	// Project unleashed
	Unleash bool `pulumi:"unleash"`
}

func GetProjectOutput(ctx *pulumi.Context, args GetProjectOutputArgs, opts ...pulumi.InvokeOption) GetProjectResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetProjectResultOutput, error) {
			args := v.(GetProjectArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:Cloud/getProject:getProject", args, GetProjectResultOutput{}, options).(GetProjectResultOutput), nil
		}).(GetProjectResultOutput)
}

// A collection of arguments for invoking getProject.
type GetProjectOutputArgs struct {
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetProjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectArgs)(nil)).Elem()
}

// A collection of values returned by getProject.
type GetProjectResultOutput struct{ *pulumi.OutputState }

func (GetProjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectResult)(nil)).Elem()
}

func (o GetProjectResultOutput) ToGetProjectResultOutput() GetProjectResultOutput {
	return o
}

func (o GetProjectResultOutput) ToGetProjectResultOutputWithContext(ctx context.Context) GetProjectResultOutput {
	return o
}

// Project access
func (o GetProjectResultOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Access }).(pulumi.StringOutput)
}

// Project creation date
func (o GetProjectResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Description of your project
func (o GetProjectResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Description }).(pulumi.StringOutput)
}

// Expiration date of your project. After this date, your project will be deleted
func (o GetProjectResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Expiration }).(pulumi.StringOutput)
}

// IAM resource information
func (o GetProjectResultOutput) Iam() GetProjectIamOutput {
	return o.ApplyT(func(v GetProjectResult) GetProjectIam { return v.Iam }).(GetProjectIamOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetProjectResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Id }).(pulumi.StringOutput)
}

// Manual quota prevent automatic quota upgrade
func (o GetProjectResultOutput) ManualQuota() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectResult) bool { return v.ManualQuota }).(pulumi.BoolOutput)
}

// Project order ID
func (o GetProjectResultOutput) OrderId() pulumi.Float64Output {
	return o.ApplyT(func(v GetProjectResult) float64 { return v.OrderId }).(pulumi.Float64Output)
}

// Order plan code
func (o GetProjectResultOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.PlanCode }).(pulumi.StringOutput)
}

// Project ID
func (o GetProjectResultOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Project name
func (o GetProjectResultOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.ProjectName }).(pulumi.StringOutput)
}

// ID of the public cloud project
func (o GetProjectResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status
func (o GetProjectResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectResult) string { return v.Status }).(pulumi.StringOutput)
}

// Project unleashed
func (o GetProjectResultOutput) Unleash() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectResult) bool { return v.Unleash }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterOutputType(GetProjectResultOutput{})
}
