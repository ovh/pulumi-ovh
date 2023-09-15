// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Import
//
// OVHcloud Managed MongoDB clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g., bash
//
// ```sh
//
//	$ pulumi import ovh:CloudProjectDatabase/mongoDbUser:MongoDbUser my_user service_name/cluster_id/id
//
// ```
type MongoDbUser struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Name of the user.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Sensitive) Password of the user.
	Password pulumi.StringOutput `pulumi:"password"`
	// See Argument Reference above.
	PasswordReset pulumi.StringPtrOutput `pulumi:"passwordReset"`
	// Roles the user belongs to.
	// Available roles:
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the user.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewMongoDbUser registers a new resource with the given unique name, arguments, and options.
func NewMongoDbUser(ctx *pulumi.Context,
	name string, args *MongoDbUserArgs, opts ...pulumi.ResourceOption) (*MongoDbUser, error) {
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
	var resource MongoDbUser
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/mongoDbUser:MongoDbUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMongoDbUser gets an existing MongoDbUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMongoDbUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MongoDbUserState, opts ...pulumi.ResourceOption) (*MongoDbUser, error) {
	var resource MongoDbUser
	err := ctx.ReadResource("ovh:CloudProjectDatabase/mongoDbUser:MongoDbUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MongoDbUser resources.
type mongoDbUserState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt *string `pulumi:"createdAt"`
	// Name of the user.
	Name *string `pulumi:"name"`
	// (Sensitive) Password of the user.
	Password *string `pulumi:"password"`
	// See Argument Reference above.
	PasswordReset *string `pulumi:"passwordReset"`
	// Roles the user belongs to.
	// Available roles:
	Roles []string `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the user.
	Status *string `pulumi:"status"`
}

type MongoDbUserState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Date of the creation of the user.
	CreatedAt pulumi.StringPtrInput
	// Name of the user.
	Name pulumi.StringPtrInput
	// (Sensitive) Password of the user.
	Password pulumi.StringPtrInput
	// See Argument Reference above.
	PasswordReset pulumi.StringPtrInput
	// Roles the user belongs to.
	// Available roles:
	Roles pulumi.StringArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the user.
	Status pulumi.StringPtrInput
}

func (MongoDbUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDbUserState)(nil)).Elem()
}

type mongoDbUserArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Name of the user.
	Name *string `pulumi:"name"`
	// See Argument Reference above.
	PasswordReset *string `pulumi:"passwordReset"`
	// Roles the user belongs to.
	// Available roles:
	Roles []string `pulumi:"roles"`
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a MongoDbUser resource.
type MongoDbUserArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Name of the user.
	Name pulumi.StringPtrInput
	// See Argument Reference above.
	PasswordReset pulumi.StringPtrInput
	// Roles the user belongs to.
	// Available roles:
	Roles pulumi.StringArrayInput
	// The id of the public cloud project. If omitted,
	// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringInput
}

func (MongoDbUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mongoDbUserArgs)(nil)).Elem()
}

type MongoDbUserInput interface {
	pulumi.Input

	ToMongoDbUserOutput() MongoDbUserOutput
	ToMongoDbUserOutputWithContext(ctx context.Context) MongoDbUserOutput
}

func (*MongoDbUser) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDbUser)(nil)).Elem()
}

func (i *MongoDbUser) ToMongoDbUserOutput() MongoDbUserOutput {
	return i.ToMongoDbUserOutputWithContext(context.Background())
}

func (i *MongoDbUser) ToMongoDbUserOutputWithContext(ctx context.Context) MongoDbUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDbUserOutput)
}

func (i *MongoDbUser) ToOutput(ctx context.Context) pulumix.Output[*MongoDbUser] {
	return pulumix.Output[*MongoDbUser]{
		OutputState: i.ToMongoDbUserOutputWithContext(ctx).OutputState,
	}
}

// MongoDbUserArrayInput is an input type that accepts MongoDbUserArray and MongoDbUserArrayOutput values.
// You can construct a concrete instance of `MongoDbUserArrayInput` via:
//
//	MongoDbUserArray{ MongoDbUserArgs{...} }
type MongoDbUserArrayInput interface {
	pulumi.Input

	ToMongoDbUserArrayOutput() MongoDbUserArrayOutput
	ToMongoDbUserArrayOutputWithContext(context.Context) MongoDbUserArrayOutput
}

type MongoDbUserArray []MongoDbUserInput

func (MongoDbUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoDbUser)(nil)).Elem()
}

func (i MongoDbUserArray) ToMongoDbUserArrayOutput() MongoDbUserArrayOutput {
	return i.ToMongoDbUserArrayOutputWithContext(context.Background())
}

func (i MongoDbUserArray) ToMongoDbUserArrayOutputWithContext(ctx context.Context) MongoDbUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDbUserArrayOutput)
}

func (i MongoDbUserArray) ToOutput(ctx context.Context) pulumix.Output[[]*MongoDbUser] {
	return pulumix.Output[[]*MongoDbUser]{
		OutputState: i.ToMongoDbUserArrayOutputWithContext(ctx).OutputState,
	}
}

// MongoDbUserMapInput is an input type that accepts MongoDbUserMap and MongoDbUserMapOutput values.
// You can construct a concrete instance of `MongoDbUserMapInput` via:
//
//	MongoDbUserMap{ "key": MongoDbUserArgs{...} }
type MongoDbUserMapInput interface {
	pulumi.Input

	ToMongoDbUserMapOutput() MongoDbUserMapOutput
	ToMongoDbUserMapOutputWithContext(context.Context) MongoDbUserMapOutput
}

type MongoDbUserMap map[string]MongoDbUserInput

func (MongoDbUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoDbUser)(nil)).Elem()
}

func (i MongoDbUserMap) ToMongoDbUserMapOutput() MongoDbUserMapOutput {
	return i.ToMongoDbUserMapOutputWithContext(context.Background())
}

func (i MongoDbUserMap) ToMongoDbUserMapOutputWithContext(ctx context.Context) MongoDbUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDbUserMapOutput)
}

func (i MongoDbUserMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MongoDbUser] {
	return pulumix.Output[map[string]*MongoDbUser]{
		OutputState: i.ToMongoDbUserMapOutputWithContext(ctx).OutputState,
	}
}

type MongoDbUserOutput struct{ *pulumi.OutputState }

func (MongoDbUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDbUser)(nil)).Elem()
}

func (o MongoDbUserOutput) ToMongoDbUserOutput() MongoDbUserOutput {
	return o
}

func (o MongoDbUserOutput) ToMongoDbUserOutputWithContext(ctx context.Context) MongoDbUserOutput {
	return o
}

func (o MongoDbUserOutput) ToOutput(ctx context.Context) pulumix.Output[*MongoDbUser] {
	return pulumix.Output[*MongoDbUser]{
		OutputState: o.OutputState,
	}
}

// Cluster ID.
func (o MongoDbUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o MongoDbUserOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Name of the user.
func (o MongoDbUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Sensitive) Password of the user.
func (o MongoDbUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// See Argument Reference above.
func (o MongoDbUserOutput) PasswordReset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringPtrOutput { return v.PasswordReset }).(pulumi.StringPtrOutput)
}

// Roles the user belongs to.
// Available roles:
func (o MongoDbUserOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringArrayOutput { return v.Roles }).(pulumi.StringArrayOutput)
}

// The id of the public cloud project. If omitted,
// the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o MongoDbUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o MongoDbUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MongoDbUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type MongoDbUserArrayOutput struct{ *pulumi.OutputState }

func (MongoDbUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MongoDbUser)(nil)).Elem()
}

func (o MongoDbUserArrayOutput) ToMongoDbUserArrayOutput() MongoDbUserArrayOutput {
	return o
}

func (o MongoDbUserArrayOutput) ToMongoDbUserArrayOutputWithContext(ctx context.Context) MongoDbUserArrayOutput {
	return o
}

func (o MongoDbUserArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MongoDbUser] {
	return pulumix.Output[[]*MongoDbUser]{
		OutputState: o.OutputState,
	}
}

func (o MongoDbUserArrayOutput) Index(i pulumi.IntInput) MongoDbUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MongoDbUser {
		return vs[0].([]*MongoDbUser)[vs[1].(int)]
	}).(MongoDbUserOutput)
}

type MongoDbUserMapOutput struct{ *pulumi.OutputState }

func (MongoDbUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MongoDbUser)(nil)).Elem()
}

func (o MongoDbUserMapOutput) ToMongoDbUserMapOutput() MongoDbUserMapOutput {
	return o
}

func (o MongoDbUserMapOutput) ToMongoDbUserMapOutputWithContext(ctx context.Context) MongoDbUserMapOutput {
	return o
}

func (o MongoDbUserMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MongoDbUser] {
	return pulumix.Output[map[string]*MongoDbUser]{
		OutputState: o.OutputState,
	}
}

func (o MongoDbUserMapOutput) MapIndex(k pulumi.StringInput) MongoDbUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MongoDbUser {
		return vs[0].(map[string]*MongoDbUser)[vs[1].(string)]
	}).(MongoDbUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MongoDbUserInput)(nil)).Elem(), &MongoDbUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoDbUserArrayInput)(nil)).Elem(), MongoDbUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MongoDbUserMapInput)(nil)).Elem(), MongoDbUserMap{})
	pulumi.RegisterOutputType(MongoDbUserOutput{})
	pulumi.RegisterOutputType(MongoDbUserArrayOutput{})
	pulumi.RegisterOutputType(MongoDbUserMapOutput{})
}
