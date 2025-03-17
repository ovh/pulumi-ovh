// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a user of a M3DB cluster associated with a public cloud project.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/cloudproject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			m3dbUser, err := cloudproject.GetM3dbUser(ctx, &cloudproject.GetM3dbUserArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				Name:        "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("m3dbUserGroup", m3dbUser.Group)
//			return nil
//		})
//	}
//
// ```
func GetM3dbUser(ctx *pulumi.Context, args *GetM3dbUserArgs, opts ...pulumi.InvokeOption) (*GetM3dbUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetM3dbUserResult
	err := ctx.Invoke("ovh:CloudProject/getM3dbUser:getM3dbUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getM3dbUser.
type GetM3dbUserArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Name of the user.
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getM3dbUser.
type GetM3dbUserResult struct {
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt string `pulumi:"createdAt"`
	// See Argument Reference above.
	Group string `pulumi:"group"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// Current status of the user.
	ServiceName string `pulumi:"serviceName"`
	// Current status of the user.
	Status string `pulumi:"status"`
}

func GetM3dbUserOutput(ctx *pulumi.Context, args GetM3dbUserOutputArgs, opts ...pulumi.InvokeOption) GetM3dbUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetM3dbUserResultOutput, error) {
			args := v.(GetM3dbUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getM3dbUser:getM3dbUser", args, GetM3dbUserResultOutput{}, options).(GetM3dbUserResultOutput), nil
		}).(GetM3dbUserResultOutput)
}

// A collection of arguments for invoking getM3dbUser.
type GetM3dbUserOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Name of the user.
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetM3dbUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetM3dbUserArgs)(nil)).Elem()
}

// A collection of values returned by getM3dbUser.
type GetM3dbUserResultOutput struct{ *pulumi.OutputState }

func (GetM3dbUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetM3dbUserResult)(nil)).Elem()
}

func (o GetM3dbUserResultOutput) ToGetM3dbUserResultOutput() GetM3dbUserResultOutput {
	return o
}

func (o GetM3dbUserResultOutput) ToGetM3dbUserResultOutputWithContext(ctx context.Context) GetM3dbUserResultOutput {
	return o
}

// See Argument Reference above.
func (o GetM3dbUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o GetM3dbUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetM3dbUserResultOutput) Group() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.Group }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetM3dbUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o GetM3dbUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current status of the user.
func (o GetM3dbUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o GetM3dbUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetM3dbUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetM3dbUserResultOutput{})
}
