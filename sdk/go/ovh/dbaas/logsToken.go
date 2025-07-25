// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dbaas

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows to manipulate LDP tokens.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/dbaas"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := dbaas.NewLogsToken(ctx, "token", &dbaas.LogsTokenArgs{
//				ServiceName: pulumi.String("ldp-xx-xxxxx"),
//				Name:        pulumi.String("ExampleToken"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
type LogsToken struct {
	pulumi.CustomResourceState

	// Cluster ID. If not provided, the default clusterId is used
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Token creation date
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the token
	Name pulumi.StringOutput `pulumi:"name"`
	// The LDP service name
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// ID of the token
	TokenId pulumi.StringOutput `pulumi:"tokenId"`
	// Token last update date
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
	// Token value
	Value pulumi.StringOutput `pulumi:"value"`
}

// NewLogsToken registers a new resource with the given unique name, arguments, and options.
func NewLogsToken(ctx *pulumi.Context,
	name string, args *LogsTokenArgs, opts ...pulumi.ResourceOption) (*LogsToken, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"value",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LogsToken
	err := ctx.RegisterResource("ovh:Dbaas/logsToken:LogsToken", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLogsToken gets an existing LogsToken resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLogsToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LogsTokenState, opts ...pulumi.ResourceOption) (*LogsToken, error) {
	var resource LogsToken
	err := ctx.ReadResource("ovh:Dbaas/logsToken:LogsToken", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LogsToken resources.
type logsTokenState struct {
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId *string `pulumi:"clusterId"`
	// Token creation date
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the token
	Name *string `pulumi:"name"`
	// The LDP service name
	ServiceName *string `pulumi:"serviceName"`
	// ID of the token
	TokenId *string `pulumi:"tokenId"`
	// Token last update date
	UpdatedAt *string `pulumi:"updatedAt"`
	// Token value
	Value *string `pulumi:"value"`
}

type LogsTokenState struct {
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId pulumi.StringPtrInput
	// Token creation date
	CreatedAt pulumi.StringPtrInput
	// Name of the token
	Name pulumi.StringPtrInput
	// The LDP service name
	ServiceName pulumi.StringPtrInput
	// ID of the token
	TokenId pulumi.StringPtrInput
	// Token last update date
	UpdatedAt pulumi.StringPtrInput
	// Token value
	Value pulumi.StringPtrInput
}

func (LogsTokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*logsTokenState)(nil)).Elem()
}

type logsTokenArgs struct {
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId *string `pulumi:"clusterId"`
	// Name of the token
	Name *string `pulumi:"name"`
	// The LDP service name
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a LogsToken resource.
type LogsTokenArgs struct {
	// Cluster ID. If not provided, the default clusterId is used
	ClusterId pulumi.StringPtrInput
	// Name of the token
	Name pulumi.StringPtrInput
	// The LDP service name
	ServiceName pulumi.StringInput
}

func (LogsTokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*logsTokenArgs)(nil)).Elem()
}

type LogsTokenInput interface {
	pulumi.Input

	ToLogsTokenOutput() LogsTokenOutput
	ToLogsTokenOutputWithContext(ctx context.Context) LogsTokenOutput
}

func (*LogsToken) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsToken)(nil)).Elem()
}

func (i *LogsToken) ToLogsTokenOutput() LogsTokenOutput {
	return i.ToLogsTokenOutputWithContext(context.Background())
}

func (i *LogsToken) ToLogsTokenOutputWithContext(ctx context.Context) LogsTokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsTokenOutput)
}

// LogsTokenArrayInput is an input type that accepts LogsTokenArray and LogsTokenArrayOutput values.
// You can construct a concrete instance of `LogsTokenArrayInput` via:
//
//	LogsTokenArray{ LogsTokenArgs{...} }
type LogsTokenArrayInput interface {
	pulumi.Input

	ToLogsTokenArrayOutput() LogsTokenArrayOutput
	ToLogsTokenArrayOutputWithContext(context.Context) LogsTokenArrayOutput
}

type LogsTokenArray []LogsTokenInput

func (LogsTokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsToken)(nil)).Elem()
}

func (i LogsTokenArray) ToLogsTokenArrayOutput() LogsTokenArrayOutput {
	return i.ToLogsTokenArrayOutputWithContext(context.Background())
}

func (i LogsTokenArray) ToLogsTokenArrayOutputWithContext(ctx context.Context) LogsTokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsTokenArrayOutput)
}

// LogsTokenMapInput is an input type that accepts LogsTokenMap and LogsTokenMapOutput values.
// You can construct a concrete instance of `LogsTokenMapInput` via:
//
//	LogsTokenMap{ "key": LogsTokenArgs{...} }
type LogsTokenMapInput interface {
	pulumi.Input

	ToLogsTokenMapOutput() LogsTokenMapOutput
	ToLogsTokenMapOutputWithContext(context.Context) LogsTokenMapOutput
}

type LogsTokenMap map[string]LogsTokenInput

func (LogsTokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsToken)(nil)).Elem()
}

func (i LogsTokenMap) ToLogsTokenMapOutput() LogsTokenMapOutput {
	return i.ToLogsTokenMapOutputWithContext(context.Background())
}

func (i LogsTokenMap) ToLogsTokenMapOutputWithContext(ctx context.Context) LogsTokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogsTokenMapOutput)
}

type LogsTokenOutput struct{ *pulumi.OutputState }

func (LogsTokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogsToken)(nil)).Elem()
}

func (o LogsTokenOutput) ToLogsTokenOutput() LogsTokenOutput {
	return o
}

func (o LogsTokenOutput) ToLogsTokenOutputWithContext(ctx context.Context) LogsTokenOutput {
	return o
}

// Cluster ID. If not provided, the default clusterId is used
func (o LogsTokenOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Token creation date
func (o LogsTokenOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the token
func (o LogsTokenOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The LDP service name
func (o LogsTokenOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// ID of the token
func (o LogsTokenOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.TokenId }).(pulumi.StringOutput)
}

// Token last update date
func (o LogsTokenOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

// Token value
func (o LogsTokenOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v *LogsToken) pulumi.StringOutput { return v.Value }).(pulumi.StringOutput)
}

type LogsTokenArrayOutput struct{ *pulumi.OutputState }

func (LogsTokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LogsToken)(nil)).Elem()
}

func (o LogsTokenArrayOutput) ToLogsTokenArrayOutput() LogsTokenArrayOutput {
	return o
}

func (o LogsTokenArrayOutput) ToLogsTokenArrayOutputWithContext(ctx context.Context) LogsTokenArrayOutput {
	return o
}

func (o LogsTokenArrayOutput) Index(i pulumi.IntInput) LogsTokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LogsToken {
		return vs[0].([]*LogsToken)[vs[1].(int)]
	}).(LogsTokenOutput)
}

type LogsTokenMapOutput struct{ *pulumi.OutputState }

func (LogsTokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LogsToken)(nil)).Elem()
}

func (o LogsTokenMapOutput) ToLogsTokenMapOutput() LogsTokenMapOutput {
	return o
}

func (o LogsTokenMapOutput) ToLogsTokenMapOutputWithContext(ctx context.Context) LogsTokenMapOutput {
	return o
}

func (o LogsTokenMapOutput) MapIndex(k pulumi.StringInput) LogsTokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LogsToken {
		return vs[0].(map[string]*LogsToken)[vs[1].(string)]
	}).(LogsTokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LogsTokenInput)(nil)).Elem(), &LogsToken{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsTokenArrayInput)(nil)).Elem(), LogsTokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LogsTokenMapInput)(nil)).Elem(), LogsTokenMap{})
	pulumi.RegisterOutputType(LogsTokenOutput{})
	pulumi.RegisterOutputType(LogsTokenArrayOutput{})
	pulumi.RegisterOutputType(LogsTokenMapOutput{})
}
