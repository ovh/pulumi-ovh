// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a user of a redis cluster associated with a public cloud project.
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
//			redisUser, err := cloudproject.GetRedisUser(ctx, &cloudproject.GetRedisUserArgs{
//				ServiceName: "XXX",
//				ClusterId:   "YYY",
//				Name:        "ZZZ",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("redisUserCommands", redisUser.Commands)
//			return nil
//		})
//	}
//
// ```
func GetRedisUser(ctx *pulumi.Context, args *GetRedisUserArgs, opts ...pulumi.InvokeOption) (*GetRedisUserResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetRedisUserResult
	err := ctx.Invoke("ovh:CloudProject/getRedisUser:getRedisUser", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRedisUser.
type GetRedisUserArgs struct {
	// Cluster ID
	ClusterId string `pulumi:"clusterId"`
	// Name of the user
	Name string `pulumi:"name"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// A collection of values returned by getRedisUser.
type GetRedisUserResult struct {
	// Categories of the user.
	Categories []string `pulumi:"categories"`
	// Channels of the user.
	Channels []string `pulumi:"channels"`
	// See Argument Reference above.
	ClusterId string `pulumi:"clusterId"`
	// Commands of the user.
	Commands []string `pulumi:"commands"`
	// Date of the creation of the user.
	CreatedAt string `pulumi:"createdAt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Keys of the user.
	Keys []string `pulumi:"keys"`
	// See Argument Reference above.
	Name string `pulumi:"name"`
	// Current status of the user.
	ServiceName string `pulumi:"serviceName"`
	// Current status of the user.
	Status string `pulumi:"status"`
}

func GetRedisUserOutput(ctx *pulumi.Context, args GetRedisUserOutputArgs, opts ...pulumi.InvokeOption) GetRedisUserResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetRedisUserResultOutput, error) {
			args := v.(GetRedisUserArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:CloudProject/getRedisUser:getRedisUser", args, GetRedisUserResultOutput{}, options).(GetRedisUserResultOutput), nil
		}).(GetRedisUserResultOutput)
}

// A collection of arguments for invoking getRedisUser.
type GetRedisUserOutputArgs struct {
	// Cluster ID
	ClusterId pulumi.StringInput `pulumi:"clusterId"`
	// Name of the user
	Name pulumi.StringInput `pulumi:"name"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (GetRedisUserOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRedisUserArgs)(nil)).Elem()
}

// A collection of values returned by getRedisUser.
type GetRedisUserResultOutput struct{ *pulumi.OutputState }

func (GetRedisUserResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetRedisUserResult)(nil)).Elem()
}

func (o GetRedisUserResultOutput) ToGetRedisUserResultOutput() GetRedisUserResultOutput {
	return o
}

func (o GetRedisUserResultOutput) ToGetRedisUserResultOutputWithContext(ctx context.Context) GetRedisUserResultOutput {
	return o
}

// Categories of the user.
func (o GetRedisUserResultOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRedisUserResult) []string { return v.Categories }).(pulumi.StringArrayOutput)
}

// Channels of the user.
func (o GetRedisUserResultOutput) Channels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRedisUserResult) []string { return v.Channels }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetRedisUserResultOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisUserResult) string { return v.ClusterId }).(pulumi.StringOutput)
}

// Commands of the user.
func (o GetRedisUserResultOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRedisUserResult) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

// Date of the creation of the user.
func (o GetRedisUserResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisUserResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetRedisUserResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisUserResult) string { return v.Id }).(pulumi.StringOutput)
}

// Keys of the user.
func (o GetRedisUserResultOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetRedisUserResult) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

// See Argument Reference above.
func (o GetRedisUserResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisUserResult) string { return v.Name }).(pulumi.StringOutput)
}

// Current status of the user.
func (o GetRedisUserResultOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisUserResult) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o GetRedisUserResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetRedisUserResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetRedisUserResultOutput{})
}
