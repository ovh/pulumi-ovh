// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package okms

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Import an existing key in the JWK format in an OVHcloud KMS.
type ServiceKeyJWK struct {
	pulumi.CustomResourceState

	// Context of the key
	Context pulumi.StringPtrOutput `pulumi:"context"`
	// Creation time of the key
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Key deactivation reason
	DeactivationReason pulumi.StringOutput `pulumi:"deactivationReason"`
	// IAM resource metadata
	Iam ServiceKeyJWKIamOutput `pulumi:"iam"`
	// Set of JSON Web Keys to import
	Keys ServiceKeyJWKKeyArrayOutput `pulumi:"keys"`
	// Key name
	Name pulumi.StringOutput `pulumi:"name"`
	// Okms ID
	OkmsId pulumi.StringOutput `pulumi:"okmsId"`
	// Size of the key to be created
	Size pulumi.Float64Output `pulumi:"size"`
	// State of the key
	State pulumi.StringOutput `pulumi:"state"`
	// Type of the key to be created
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceKeyJWK registers a new resource with the given unique name, arguments, and options.
func NewServiceKeyJWK(ctx *pulumi.Context,
	name string, args *ServiceKeyJWKArgs, opts ...pulumi.ResourceOption) (*ServiceKeyJWK, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Keys == nil {
		return nil, errors.New("invalid value for required argument 'Keys'")
	}
	if args.OkmsId == nil {
		return nil, errors.New("invalid value for required argument 'OkmsId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceKeyJWK
	err := ctx.RegisterResource("ovh:Okms/serviceKeyJWK:ServiceKeyJWK", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceKeyJWK gets an existing ServiceKeyJWK resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceKeyJWK(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceKeyJWKState, opts ...pulumi.ResourceOption) (*ServiceKeyJWK, error) {
	var resource ServiceKeyJWK
	err := ctx.ReadResource("ovh:Okms/serviceKeyJWK:ServiceKeyJWK", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceKeyJWK resources.
type serviceKeyJWKState struct {
	// Context of the key
	Context *string `pulumi:"context"`
	// Creation time of the key
	CreatedAt *string `pulumi:"createdAt"`
	// Key deactivation reason
	DeactivationReason *string `pulumi:"deactivationReason"`
	// IAM resource metadata
	Iam *ServiceKeyJWKIam `pulumi:"iam"`
	// Set of JSON Web Keys to import
	Keys []ServiceKeyJWKKey `pulumi:"keys"`
	// Key name
	Name *string `pulumi:"name"`
	// Okms ID
	OkmsId *string `pulumi:"okmsId"`
	// Size of the key to be created
	Size *float64 `pulumi:"size"`
	// State of the key
	State *string `pulumi:"state"`
	// Type of the key to be created
	Type *string `pulumi:"type"`
}

type ServiceKeyJWKState struct {
	// Context of the key
	Context pulumi.StringPtrInput
	// Creation time of the key
	CreatedAt pulumi.StringPtrInput
	// Key deactivation reason
	DeactivationReason pulumi.StringPtrInput
	// IAM resource metadata
	Iam ServiceKeyJWKIamPtrInput
	// Set of JSON Web Keys to import
	Keys ServiceKeyJWKKeyArrayInput
	// Key name
	Name pulumi.StringPtrInput
	// Okms ID
	OkmsId pulumi.StringPtrInput
	// Size of the key to be created
	Size pulumi.Float64PtrInput
	// State of the key
	State pulumi.StringPtrInput
	// Type of the key to be created
	Type pulumi.StringPtrInput
}

func (ServiceKeyJWKState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceKeyJWKState)(nil)).Elem()
}

type serviceKeyJWKArgs struct {
	// Context of the key
	Context *string `pulumi:"context"`
	// Set of JSON Web Keys to import
	Keys []ServiceKeyJWKKey `pulumi:"keys"`
	// Key name
	Name *string `pulumi:"name"`
	// Okms ID
	OkmsId string `pulumi:"okmsId"`
}

// The set of arguments for constructing a ServiceKeyJWK resource.
type ServiceKeyJWKArgs struct {
	// Context of the key
	Context pulumi.StringPtrInput
	// Set of JSON Web Keys to import
	Keys ServiceKeyJWKKeyArrayInput
	// Key name
	Name pulumi.StringPtrInput
	// Okms ID
	OkmsId pulumi.StringInput
}

func (ServiceKeyJWKArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceKeyJWKArgs)(nil)).Elem()
}

type ServiceKeyJWKInput interface {
	pulumi.Input

	ToServiceKeyJWKOutput() ServiceKeyJWKOutput
	ToServiceKeyJWKOutputWithContext(ctx context.Context) ServiceKeyJWKOutput
}

func (*ServiceKeyJWK) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceKeyJWK)(nil)).Elem()
}

func (i *ServiceKeyJWK) ToServiceKeyJWKOutput() ServiceKeyJWKOutput {
	return i.ToServiceKeyJWKOutputWithContext(context.Background())
}

func (i *ServiceKeyJWK) ToServiceKeyJWKOutputWithContext(ctx context.Context) ServiceKeyJWKOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceKeyJWKOutput)
}

// ServiceKeyJWKArrayInput is an input type that accepts ServiceKeyJWKArray and ServiceKeyJWKArrayOutput values.
// You can construct a concrete instance of `ServiceKeyJWKArrayInput` via:
//
//	ServiceKeyJWKArray{ ServiceKeyJWKArgs{...} }
type ServiceKeyJWKArrayInput interface {
	pulumi.Input

	ToServiceKeyJWKArrayOutput() ServiceKeyJWKArrayOutput
	ToServiceKeyJWKArrayOutputWithContext(context.Context) ServiceKeyJWKArrayOutput
}

type ServiceKeyJWKArray []ServiceKeyJWKInput

func (ServiceKeyJWKArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceKeyJWK)(nil)).Elem()
}

func (i ServiceKeyJWKArray) ToServiceKeyJWKArrayOutput() ServiceKeyJWKArrayOutput {
	return i.ToServiceKeyJWKArrayOutputWithContext(context.Background())
}

func (i ServiceKeyJWKArray) ToServiceKeyJWKArrayOutputWithContext(ctx context.Context) ServiceKeyJWKArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceKeyJWKArrayOutput)
}

// ServiceKeyJWKMapInput is an input type that accepts ServiceKeyJWKMap and ServiceKeyJWKMapOutput values.
// You can construct a concrete instance of `ServiceKeyJWKMapInput` via:
//
//	ServiceKeyJWKMap{ "key": ServiceKeyJWKArgs{...} }
type ServiceKeyJWKMapInput interface {
	pulumi.Input

	ToServiceKeyJWKMapOutput() ServiceKeyJWKMapOutput
	ToServiceKeyJWKMapOutputWithContext(context.Context) ServiceKeyJWKMapOutput
}

type ServiceKeyJWKMap map[string]ServiceKeyJWKInput

func (ServiceKeyJWKMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceKeyJWK)(nil)).Elem()
}

func (i ServiceKeyJWKMap) ToServiceKeyJWKMapOutput() ServiceKeyJWKMapOutput {
	return i.ToServiceKeyJWKMapOutputWithContext(context.Background())
}

func (i ServiceKeyJWKMap) ToServiceKeyJWKMapOutputWithContext(ctx context.Context) ServiceKeyJWKMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceKeyJWKMapOutput)
}

type ServiceKeyJWKOutput struct{ *pulumi.OutputState }

func (ServiceKeyJWKOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceKeyJWK)(nil)).Elem()
}

func (o ServiceKeyJWKOutput) ToServiceKeyJWKOutput() ServiceKeyJWKOutput {
	return o
}

func (o ServiceKeyJWKOutput) ToServiceKeyJWKOutputWithContext(ctx context.Context) ServiceKeyJWKOutput {
	return o
}

// Context of the key
func (o ServiceKeyJWKOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringPtrOutput { return v.Context }).(pulumi.StringPtrOutput)
}

// Creation time of the key
func (o ServiceKeyJWKOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Key deactivation reason
func (o ServiceKeyJWKOutput) DeactivationReason() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringOutput { return v.DeactivationReason }).(pulumi.StringOutput)
}

// IAM resource metadata
func (o ServiceKeyJWKOutput) Iam() ServiceKeyJWKIamOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) ServiceKeyJWKIamOutput { return v.Iam }).(ServiceKeyJWKIamOutput)
}

// Set of JSON Web Keys to import
func (o ServiceKeyJWKOutput) Keys() ServiceKeyJWKKeyArrayOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) ServiceKeyJWKKeyArrayOutput { return v.Keys }).(ServiceKeyJWKKeyArrayOutput)
}

// Key name
func (o ServiceKeyJWKOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Okms ID
func (o ServiceKeyJWKOutput) OkmsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringOutput { return v.OkmsId }).(pulumi.StringOutput)
}

// Size of the key to be created
func (o ServiceKeyJWKOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// State of the key
func (o ServiceKeyJWKOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Type of the key to be created
func (o ServiceKeyJWKOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKeyJWK) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ServiceKeyJWKArrayOutput struct{ *pulumi.OutputState }

func (ServiceKeyJWKArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceKeyJWK)(nil)).Elem()
}

func (o ServiceKeyJWKArrayOutput) ToServiceKeyJWKArrayOutput() ServiceKeyJWKArrayOutput {
	return o
}

func (o ServiceKeyJWKArrayOutput) ToServiceKeyJWKArrayOutputWithContext(ctx context.Context) ServiceKeyJWKArrayOutput {
	return o
}

func (o ServiceKeyJWKArrayOutput) Index(i pulumi.IntInput) ServiceKeyJWKOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceKeyJWK {
		return vs[0].([]*ServiceKeyJWK)[vs[1].(int)]
	}).(ServiceKeyJWKOutput)
}

type ServiceKeyJWKMapOutput struct{ *pulumi.OutputState }

func (ServiceKeyJWKMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceKeyJWK)(nil)).Elem()
}

func (o ServiceKeyJWKMapOutput) ToServiceKeyJWKMapOutput() ServiceKeyJWKMapOutput {
	return o
}

func (o ServiceKeyJWKMapOutput) ToServiceKeyJWKMapOutputWithContext(ctx context.Context) ServiceKeyJWKMapOutput {
	return o
}

func (o ServiceKeyJWKMapOutput) MapIndex(k pulumi.StringInput) ServiceKeyJWKOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceKeyJWK {
		return vs[0].(map[string]*ServiceKeyJWK)[vs[1].(string)]
	}).(ServiceKeyJWKOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKeyJWKInput)(nil)).Elem(), &ServiceKeyJWK{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKeyJWKArrayInput)(nil)).Elem(), ServiceKeyJWKArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKeyJWKMapInput)(nil)).Elem(), ServiceKeyJWKMap{})
	pulumi.RegisterOutputType(ServiceKeyJWKOutput{})
	pulumi.RegisterOutputType(ServiceKeyJWKArrayOutput{})
	pulumi.RegisterOutputType(ServiceKeyJWKMapOutput{})
}
