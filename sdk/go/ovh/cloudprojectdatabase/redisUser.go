// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// OVHcloud Managed Redis clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/redisUser:RedisUser my_user service_name/cluster_id/id
// ```
type RedisUser struct {
	pulumi.CustomResourceState

	// Categories of the user.
	Categories pulumi.StringArrayOutput `pulumi:"categories"`
	// Channels of the user.
	Channels pulumi.StringArrayOutput `pulumi:"channels"`
	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Commands of the user.
	Commands pulumi.StringArrayOutput `pulumi:"commands"`
	// Date of the creation of the user.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Keys of the user.
	Keys pulumi.StringArrayOutput `pulumi:"keys"`
	// Name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Sensitive) Password of the user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Arbitrary string to change to trigger a password update
	PasswordReset pulumi.StringPtrOutput `pulumi:"passwordReset"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the user.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewRedisUser registers a new resource with the given unique name, arguments, and options.
func NewRedisUser(ctx *pulumi.Context,
	name string, args *RedisUserArgs, opts ...pulumi.ResourceOption) (*RedisUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RedisUser
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/redisUser:RedisUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRedisUser gets an existing RedisUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRedisUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RedisUserState, opts ...pulumi.ResourceOption) (*RedisUser, error) {
	var resource RedisUser
	err := ctx.ReadResource("ovh:CloudProjectDatabase/redisUser:RedisUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RedisUser resources.
type redisUserState struct {
	// Categories of the user.
	Categories []string `pulumi:"categories"`
	// Channels of the user.
	Channels []string `pulumi:"channels"`
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Commands of the user.
	Commands []string `pulumi:"commands"`
	// Date of the creation of the user.
	CreatedAt *string `pulumi:"createdAt"`
	// Keys of the user.
	Keys []string `pulumi:"keys"`
	// Name of the user.
	Name *string `pulumi:"name"`
	// (Sensitive) Password of the user.
	Password *string `pulumi:"password"`
	// Arbitrary string to change to trigger a password update
	PasswordReset *string `pulumi:"passwordReset"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the user.
	Status *string `pulumi:"status"`
}

type RedisUserState struct {
	// Categories of the user.
	Categories pulumi.StringArrayInput
	// Channels of the user.
	Channels pulumi.StringArrayInput
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Commands of the user.
	Commands pulumi.StringArrayInput
	// Date of the creation of the user.
	CreatedAt pulumi.StringPtrInput
	// Keys of the user.
	Keys pulumi.StringArrayInput
	// Name of the user.
	Name pulumi.StringPtrInput
	// (Sensitive) Password of the user.
	Password pulumi.StringPtrInput
	// Arbitrary string to change to trigger a password update
	PasswordReset pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the user.
	Status pulumi.StringPtrInput
}

func (RedisUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*redisUserState)(nil)).Elem()
}

type redisUserArgs struct {
	// Categories of the user.
	Categories []string `pulumi:"categories"`
	// Channels of the user.
	Channels []string `pulumi:"channels"`
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Commands of the user.
	Commands []string `pulumi:"commands"`
	// Keys of the user.
	Keys []string `pulumi:"keys"`
	// Name of the user.
	Name *string `pulumi:"name"`
	// Arbitrary string to change to trigger a password update
	PasswordReset *string `pulumi:"passwordReset"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a RedisUser resource.
type RedisUserArgs struct {
	// Categories of the user.
	Categories pulumi.StringArrayInput
	// Channels of the user.
	Channels pulumi.StringArrayInput
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Commands of the user.
	Commands pulumi.StringArrayInput
	// Keys of the user.
	Keys pulumi.StringArrayInput
	// Name of the user.
	Name pulumi.StringPtrInput
	// Arbitrary string to change to trigger a password update
	PasswordReset pulumi.StringPtrInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (RedisUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*redisUserArgs)(nil)).Elem()
}

type RedisUserInput interface {
	pulumi.Input

	ToRedisUserOutput() RedisUserOutput
	ToRedisUserOutputWithContext(ctx context.Context) RedisUserOutput
}

func (*RedisUser) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisUser)(nil)).Elem()
}

func (i *RedisUser) ToRedisUserOutput() RedisUserOutput {
	return i.ToRedisUserOutputWithContext(context.Background())
}

func (i *RedisUser) ToRedisUserOutputWithContext(ctx context.Context) RedisUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisUserOutput)
}

// RedisUserArrayInput is an input type that accepts RedisUserArray and RedisUserArrayOutput values.
// You can construct a concrete instance of `RedisUserArrayInput` via:
//
//	RedisUserArray{ RedisUserArgs{...} }
type RedisUserArrayInput interface {
	pulumi.Input

	ToRedisUserArrayOutput() RedisUserArrayOutput
	ToRedisUserArrayOutputWithContext(context.Context) RedisUserArrayOutput
}

type RedisUserArray []RedisUserInput

func (RedisUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RedisUser)(nil)).Elem()
}

func (i RedisUserArray) ToRedisUserArrayOutput() RedisUserArrayOutput {
	return i.ToRedisUserArrayOutputWithContext(context.Background())
}

func (i RedisUserArray) ToRedisUserArrayOutputWithContext(ctx context.Context) RedisUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisUserArrayOutput)
}

// RedisUserMapInput is an input type that accepts RedisUserMap and RedisUserMapOutput values.
// You can construct a concrete instance of `RedisUserMapInput` via:
//
//	RedisUserMap{ "key": RedisUserArgs{...} }
type RedisUserMapInput interface {
	pulumi.Input

	ToRedisUserMapOutput() RedisUserMapOutput
	ToRedisUserMapOutputWithContext(context.Context) RedisUserMapOutput
}

type RedisUserMap map[string]RedisUserInput

func (RedisUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RedisUser)(nil)).Elem()
}

func (i RedisUserMap) ToRedisUserMapOutput() RedisUserMapOutput {
	return i.ToRedisUserMapOutputWithContext(context.Background())
}

func (i RedisUserMap) ToRedisUserMapOutputWithContext(ctx context.Context) RedisUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RedisUserMapOutput)
}

type RedisUserOutput struct{ *pulumi.OutputState }

func (RedisUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RedisUser)(nil)).Elem()
}

func (o RedisUserOutput) ToRedisUserOutput() RedisUserOutput {
	return o
}

func (o RedisUserOutput) ToRedisUserOutputWithContext(ctx context.Context) RedisUserOutput {
	return o
}

// Categories of the user.
func (o RedisUserOutput) Categories() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringArrayOutput { return v.Categories }).(pulumi.StringArrayOutput)
}

// Channels of the user.
func (o RedisUserOutput) Channels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringArrayOutput { return v.Channels }).(pulumi.StringArrayOutput)
}

// Cluster ID.
func (o RedisUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Commands of the user.
func (o RedisUserOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringArrayOutput { return v.Commands }).(pulumi.StringArrayOutput)
}

// Date of the creation of the user.
func (o RedisUserOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Keys of the user.
func (o RedisUserOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringArrayOutput { return v.Keys }).(pulumi.StringArrayOutput)
}

// Name of the user.
func (o RedisUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Sensitive) Password of the user.
func (o RedisUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Arbitrary string to change to trigger a password update
func (o RedisUserOutput) PasswordReset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringPtrOutput { return v.PasswordReset }).(pulumi.StringPtrOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o RedisUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o RedisUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *RedisUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type RedisUserArrayOutput struct{ *pulumi.OutputState }

func (RedisUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RedisUser)(nil)).Elem()
}

func (o RedisUserArrayOutput) ToRedisUserArrayOutput() RedisUserArrayOutput {
	return o
}

func (o RedisUserArrayOutput) ToRedisUserArrayOutputWithContext(ctx context.Context) RedisUserArrayOutput {
	return o
}

func (o RedisUserArrayOutput) Index(i pulumi.IntInput) RedisUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RedisUser {
		return vs[0].([]*RedisUser)[vs[1].(int)]
	}).(RedisUserOutput)
}

type RedisUserMapOutput struct{ *pulumi.OutputState }

func (RedisUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RedisUser)(nil)).Elem()
}

func (o RedisUserMapOutput) ToRedisUserMapOutput() RedisUserMapOutput {
	return o
}

func (o RedisUserMapOutput) ToRedisUserMapOutputWithContext(ctx context.Context) RedisUserMapOutput {
	return o
}

func (o RedisUserMapOutput) MapIndex(k pulumi.StringInput) RedisUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RedisUser {
		return vs[0].(map[string]*RedisUser)[vs[1].(string)]
	}).(RedisUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RedisUserInput)(nil)).Elem(), &RedisUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisUserArrayInput)(nil)).Elem(), RedisUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RedisUserMapInput)(nil)).Elem(), RedisUserMap{})
	pulumi.RegisterOutputType(RedisUserOutput{})
	pulumi.RegisterOutputType(RedisUserArrayOutput{})
	pulumi.RegisterOutputType(RedisUserMapOutput{})
}
