// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudprojectdatabase

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// OVHcloud Managed M3DB clusters users can be imported using the `service_name`, `cluster_id` and `id` of the user, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:CloudProjectDatabase/m3DbUser:M3DbUser my_user service_name/cluster_id/id
// ```
type M3DbUser struct {
	pulumi.CustomResourceState

	// Cluster ID.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Group of the user.
	Group pulumi.StringPtrOutput `pulumi:"group"`
	// Name of the user. A user named "avnadmin" is mapped with already created admin user instead of creating a new user.
	Name pulumi.StringOutput `pulumi:"name"`
	// (Sensitive) Password of the user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Arbitrary string to change to trigger a password update
	PasswordReset pulumi.StringPtrOutput `pulumi:"passwordReset"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Current status of the user.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewM3DbUser registers a new resource with the given unique name, arguments, and options.
func NewM3DbUser(ctx *pulumi.Context,
	name string, args *M3DbUserArgs, opts ...pulumi.ResourceOption) (*M3DbUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource M3DbUser
	err := ctx.RegisterResource("ovh:CloudProjectDatabase/m3DbUser:M3DbUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetM3DbUser gets an existing M3DbUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetM3DbUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *M3DbUserState, opts ...pulumi.ResourceOption) (*M3DbUser, error) {
	var resource M3DbUser
	err := ctx.ReadResource("ovh:CloudProjectDatabase/m3DbUser:M3DbUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering M3DbUser resources.
type m3dbUserState struct {
	// Cluster ID.
	ClusterId *string `pulumi:"clusterId"`
	// Date of the creation of the user.
	CreatedAt *string `pulumi:"createdAt"`
	// Group of the user.
	Group *string `pulumi:"group"`
	// Name of the user. A user named "avnadmin" is mapped with already created admin user instead of creating a new user.
	Name *string `pulumi:"name"`
	// (Sensitive) Password of the user.
	Password *string `pulumi:"password"`
	// Arbitrary string to change to trigger a password update
	PasswordReset *string `pulumi:"passwordReset"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
	// Current status of the user.
	Status *string `pulumi:"status"`
}

type M3DbUserState struct {
	// Cluster ID.
	ClusterId pulumi.StringPtrInput
	// Date of the creation of the user.
	CreatedAt pulumi.StringPtrInput
	// Group of the user.
	Group pulumi.StringPtrInput
	// Name of the user. A user named "avnadmin" is mapped with already created admin user instead of creating a new user.
	Name pulumi.StringPtrInput
	// (Sensitive) Password of the user.
	Password pulumi.StringPtrInput
	// Arbitrary string to change to trigger a password update
	PasswordReset pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
	// Current status of the user.
	Status pulumi.StringPtrInput
}

func (M3DbUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*m3dbUserState)(nil)).Elem()
}

type m3dbUserArgs struct {
	// Cluster ID.
	ClusterId string `pulumi:"clusterId"`
	// Group of the user.
	Group *string `pulumi:"group"`
	// Name of the user. A user named "avnadmin" is mapped with already created admin user instead of creating a new user.
	Name *string `pulumi:"name"`
	// Arbitrary string to change to trigger a password update
	PasswordReset *string `pulumi:"passwordReset"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName *string `pulumi:"serviceName"`
}

// The set of arguments for constructing a M3DbUser resource.
type M3DbUserArgs struct {
	// Cluster ID.
	ClusterId pulumi.StringInput
	// Group of the user.
	Group pulumi.StringPtrInput
	// Name of the user. A user named "avnadmin" is mapped with already created admin user instead of creating a new user.
	Name pulumi.StringPtrInput
	// Arbitrary string to change to trigger a password update
	PasswordReset pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
	ServiceName pulumi.StringPtrInput
}

func (M3DbUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*m3dbUserArgs)(nil)).Elem()
}

type M3DbUserInput interface {
	pulumi.Input

	ToM3DbUserOutput() M3DbUserOutput
	ToM3DbUserOutputWithContext(ctx context.Context) M3DbUserOutput
}

func (*M3DbUser) ElementType() reflect.Type {
	return reflect.TypeOf((**M3DbUser)(nil)).Elem()
}

func (i *M3DbUser) ToM3DbUserOutput() M3DbUserOutput {
	return i.ToM3DbUserOutputWithContext(context.Background())
}

func (i *M3DbUser) ToM3DbUserOutputWithContext(ctx context.Context) M3DbUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(M3DbUserOutput)
}

// M3DbUserArrayInput is an input type that accepts M3DbUserArray and M3DbUserArrayOutput values.
// You can construct a concrete instance of `M3DbUserArrayInput` via:
//
//	M3DbUserArray{ M3DbUserArgs{...} }
type M3DbUserArrayInput interface {
	pulumi.Input

	ToM3DbUserArrayOutput() M3DbUserArrayOutput
	ToM3DbUserArrayOutputWithContext(context.Context) M3DbUserArrayOutput
}

type M3DbUserArray []M3DbUserInput

func (M3DbUserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*M3DbUser)(nil)).Elem()
}

func (i M3DbUserArray) ToM3DbUserArrayOutput() M3DbUserArrayOutput {
	return i.ToM3DbUserArrayOutputWithContext(context.Background())
}

func (i M3DbUserArray) ToM3DbUserArrayOutputWithContext(ctx context.Context) M3DbUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(M3DbUserArrayOutput)
}

// M3DbUserMapInput is an input type that accepts M3DbUserMap and M3DbUserMapOutput values.
// You can construct a concrete instance of `M3DbUserMapInput` via:
//
//	M3DbUserMap{ "key": M3DbUserArgs{...} }
type M3DbUserMapInput interface {
	pulumi.Input

	ToM3DbUserMapOutput() M3DbUserMapOutput
	ToM3DbUserMapOutputWithContext(context.Context) M3DbUserMapOutput
}

type M3DbUserMap map[string]M3DbUserInput

func (M3DbUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*M3DbUser)(nil)).Elem()
}

func (i M3DbUserMap) ToM3DbUserMapOutput() M3DbUserMapOutput {
	return i.ToM3DbUserMapOutputWithContext(context.Background())
}

func (i M3DbUserMap) ToM3DbUserMapOutputWithContext(ctx context.Context) M3DbUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(M3DbUserMapOutput)
}

type M3DbUserOutput struct{ *pulumi.OutputState }

func (M3DbUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**M3DbUser)(nil)).Elem()
}

func (o M3DbUserOutput) ToM3DbUserOutput() M3DbUserOutput {
	return o
}

func (o M3DbUserOutput) ToM3DbUserOutputWithContext(ctx context.Context) M3DbUserOutput {
	return o
}

// Cluster ID.
func (o M3DbUserOutput) ClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringOutput { return v.ClusterId }).(pulumi.StringOutput)
}

// Date of the creation of the user.
func (o M3DbUserOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Group of the user.
func (o M3DbUserOutput) Group() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringPtrOutput { return v.Group }).(pulumi.StringPtrOutput)
}

// Name of the user. A user named "avnadmin" is mapped with already created admin user instead of creating a new user.
func (o M3DbUserOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// (Sensitive) Password of the user.
func (o M3DbUserOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// Arbitrary string to change to trigger a password update
func (o M3DbUserOutput) PasswordReset() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringPtrOutput { return v.PasswordReset }).(pulumi.StringPtrOutput)
}

// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
func (o M3DbUserOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status of the user.
func (o M3DbUserOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *M3DbUser) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type M3DbUserArrayOutput struct{ *pulumi.OutputState }

func (M3DbUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*M3DbUser)(nil)).Elem()
}

func (o M3DbUserArrayOutput) ToM3DbUserArrayOutput() M3DbUserArrayOutput {
	return o
}

func (o M3DbUserArrayOutput) ToM3DbUserArrayOutputWithContext(ctx context.Context) M3DbUserArrayOutput {
	return o
}

func (o M3DbUserArrayOutput) Index(i pulumi.IntInput) M3DbUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *M3DbUser {
		return vs[0].([]*M3DbUser)[vs[1].(int)]
	}).(M3DbUserOutput)
}

type M3DbUserMapOutput struct{ *pulumi.OutputState }

func (M3DbUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*M3DbUser)(nil)).Elem()
}

func (o M3DbUserMapOutput) ToM3DbUserMapOutput() M3DbUserMapOutput {
	return o
}

func (o M3DbUserMapOutput) ToM3DbUserMapOutputWithContext(ctx context.Context) M3DbUserMapOutput {
	return o
}

func (o M3DbUserMapOutput) MapIndex(k pulumi.StringInput) M3DbUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *M3DbUser {
		return vs[0].(map[string]*M3DbUser)[vs[1].(string)]
	}).(M3DbUserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*M3DbUserInput)(nil)).Elem(), &M3DbUser{})
	pulumi.RegisterInputType(reflect.TypeOf((*M3DbUserArrayInput)(nil)).Elem(), M3DbUserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*M3DbUserMapInput)(nil)).Elem(), M3DbUserMap{})
	pulumi.RegisterOutputType(M3DbUserOutput{})
	pulumi.RegisterOutputType(M3DbUserArrayOutput{})
	pulumi.RegisterOutputType(M3DbUserMapOutput{})
}
