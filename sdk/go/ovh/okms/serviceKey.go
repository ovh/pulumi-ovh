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

// Creates a Service Key in an OVHcloud KMS.
type ServiceKey struct {
	pulumi.CustomResourceState

	// Context of the key
	Context pulumi.StringOutput `pulumi:"context"`
	// Creation time of the key
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Curve type for Elliptic Curve (EC) keys
	Curve pulumi.StringOutput `pulumi:"curve"`
	// Key deactivation reason
	DeactivationReason pulumi.StringOutput `pulumi:"deactivationReason"`
	// Key name
	Name pulumi.StringOutput `pulumi:"name"`
	// Okms ID
	OkmsId pulumi.StringOutput `pulumi:"okmsId"`
	// The operations for which the key is intended to be used
	Operations pulumi.StringArrayOutput `pulumi:"operations"`
	// Size of the key to be created
	Size pulumi.Float64Output `pulumi:"size"`
	// State of the key
	State pulumi.StringOutput `pulumi:"state"`
	// Type of the key to be created
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewServiceKey registers a new resource with the given unique name, arguments, and options.
func NewServiceKey(ctx *pulumi.Context,
	name string, args *ServiceKeyArgs, opts ...pulumi.ResourceOption) (*ServiceKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OkmsId == nil {
		return nil, errors.New("invalid value for required argument 'OkmsId'")
	}
	if args.Operations == nil {
		return nil, errors.New("invalid value for required argument 'Operations'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServiceKey
	err := ctx.RegisterResource("ovh:Okms/serviceKey:ServiceKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServiceKey gets an existing ServiceKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServiceKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceKeyState, opts ...pulumi.ResourceOption) (*ServiceKey, error) {
	var resource ServiceKey
	err := ctx.ReadResource("ovh:Okms/serviceKey:ServiceKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServiceKey resources.
type serviceKeyState struct {
	// Context of the key
	Context *string `pulumi:"context"`
	// Creation time of the key
	CreatedAt *string `pulumi:"createdAt"`
	// Curve type for Elliptic Curve (EC) keys
	Curve *string `pulumi:"curve"`
	// Key deactivation reason
	DeactivationReason *string `pulumi:"deactivationReason"`
	// Key name
	Name *string `pulumi:"name"`
	// Okms ID
	OkmsId *string `pulumi:"okmsId"`
	// The operations for which the key is intended to be used
	Operations []string `pulumi:"operations"`
	// Size of the key to be created
	Size *float64 `pulumi:"size"`
	// State of the key
	State *string `pulumi:"state"`
	// Type of the key to be created
	Type *string `pulumi:"type"`
}

type ServiceKeyState struct {
	// Context of the key
	Context pulumi.StringPtrInput
	// Creation time of the key
	CreatedAt pulumi.StringPtrInput
	// Curve type for Elliptic Curve (EC) keys
	Curve pulumi.StringPtrInput
	// Key deactivation reason
	DeactivationReason pulumi.StringPtrInput
	// Key name
	Name pulumi.StringPtrInput
	// Okms ID
	OkmsId pulumi.StringPtrInput
	// The operations for which the key is intended to be used
	Operations pulumi.StringArrayInput
	// Size of the key to be created
	Size pulumi.Float64PtrInput
	// State of the key
	State pulumi.StringPtrInput
	// Type of the key to be created
	Type pulumi.StringPtrInput
}

func (ServiceKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceKeyState)(nil)).Elem()
}

type serviceKeyArgs struct {
	// Context of the key
	Context *string `pulumi:"context"`
	// Curve type for Elliptic Curve (EC) keys
	Curve *string `pulumi:"curve"`
	// Key name
	Name *string `pulumi:"name"`
	// Okms ID
	OkmsId string `pulumi:"okmsId"`
	// The operations for which the key is intended to be used
	Operations []string `pulumi:"operations"`
	// Size of the key to be created
	Size *float64 `pulumi:"size"`
	// Type of the key to be created
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a ServiceKey resource.
type ServiceKeyArgs struct {
	// Context of the key
	Context pulumi.StringPtrInput
	// Curve type for Elliptic Curve (EC) keys
	Curve pulumi.StringPtrInput
	// Key name
	Name pulumi.StringPtrInput
	// Okms ID
	OkmsId pulumi.StringInput
	// The operations for which the key is intended to be used
	Operations pulumi.StringArrayInput
	// Size of the key to be created
	Size pulumi.Float64PtrInput
	// Type of the key to be created
	Type pulumi.StringInput
}

func (ServiceKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceKeyArgs)(nil)).Elem()
}

type ServiceKeyInput interface {
	pulumi.Input

	ToServiceKeyOutput() ServiceKeyOutput
	ToServiceKeyOutputWithContext(ctx context.Context) ServiceKeyOutput
}

func (*ServiceKey) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceKey)(nil)).Elem()
}

func (i *ServiceKey) ToServiceKeyOutput() ServiceKeyOutput {
	return i.ToServiceKeyOutputWithContext(context.Background())
}

func (i *ServiceKey) ToServiceKeyOutputWithContext(ctx context.Context) ServiceKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceKeyOutput)
}

// ServiceKeyArrayInput is an input type that accepts ServiceKeyArray and ServiceKeyArrayOutput values.
// You can construct a concrete instance of `ServiceKeyArrayInput` via:
//
//	ServiceKeyArray{ ServiceKeyArgs{...} }
type ServiceKeyArrayInput interface {
	pulumi.Input

	ToServiceKeyArrayOutput() ServiceKeyArrayOutput
	ToServiceKeyArrayOutputWithContext(context.Context) ServiceKeyArrayOutput
}

type ServiceKeyArray []ServiceKeyInput

func (ServiceKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceKey)(nil)).Elem()
}

func (i ServiceKeyArray) ToServiceKeyArrayOutput() ServiceKeyArrayOutput {
	return i.ToServiceKeyArrayOutputWithContext(context.Background())
}

func (i ServiceKeyArray) ToServiceKeyArrayOutputWithContext(ctx context.Context) ServiceKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceKeyArrayOutput)
}

// ServiceKeyMapInput is an input type that accepts ServiceKeyMap and ServiceKeyMapOutput values.
// You can construct a concrete instance of `ServiceKeyMapInput` via:
//
//	ServiceKeyMap{ "key": ServiceKeyArgs{...} }
type ServiceKeyMapInput interface {
	pulumi.Input

	ToServiceKeyMapOutput() ServiceKeyMapOutput
	ToServiceKeyMapOutputWithContext(context.Context) ServiceKeyMapOutput
}

type ServiceKeyMap map[string]ServiceKeyInput

func (ServiceKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceKey)(nil)).Elem()
}

func (i ServiceKeyMap) ToServiceKeyMapOutput() ServiceKeyMapOutput {
	return i.ToServiceKeyMapOutputWithContext(context.Background())
}

func (i ServiceKeyMap) ToServiceKeyMapOutputWithContext(ctx context.Context) ServiceKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceKeyMapOutput)
}

type ServiceKeyOutput struct{ *pulumi.OutputState }

func (ServiceKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceKey)(nil)).Elem()
}

func (o ServiceKeyOutput) ToServiceKeyOutput() ServiceKeyOutput {
	return o
}

func (o ServiceKeyOutput) ToServiceKeyOutputWithContext(ctx context.Context) ServiceKeyOutput {
	return o
}

// Context of the key
func (o ServiceKeyOutput) Context() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.Context }).(pulumi.StringOutput)
}

// Creation time of the key
func (o ServiceKeyOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Curve type for Elliptic Curve (EC) keys
func (o ServiceKeyOutput) Curve() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.Curve }).(pulumi.StringOutput)
}

// Key deactivation reason
func (o ServiceKeyOutput) DeactivationReason() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.DeactivationReason }).(pulumi.StringOutput)
}

// Key name
func (o ServiceKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Okms ID
func (o ServiceKeyOutput) OkmsId() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.OkmsId }).(pulumi.StringOutput)
}

// The operations for which the key is intended to be used
func (o ServiceKeyOutput) Operations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringArrayOutput { return v.Operations }).(pulumi.StringArrayOutput)
}

// Size of the key to be created
func (o ServiceKeyOutput) Size() pulumi.Float64Output {
	return o.ApplyT(func(v *ServiceKey) pulumi.Float64Output { return v.Size }).(pulumi.Float64Output)
}

// State of the key
func (o ServiceKeyOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Type of the key to be created
func (o ServiceKeyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceKey) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type ServiceKeyArrayOutput struct{ *pulumi.OutputState }

func (ServiceKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServiceKey)(nil)).Elem()
}

func (o ServiceKeyArrayOutput) ToServiceKeyArrayOutput() ServiceKeyArrayOutput {
	return o
}

func (o ServiceKeyArrayOutput) ToServiceKeyArrayOutputWithContext(ctx context.Context) ServiceKeyArrayOutput {
	return o
}

func (o ServiceKeyArrayOutput) Index(i pulumi.IntInput) ServiceKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServiceKey {
		return vs[0].([]*ServiceKey)[vs[1].(int)]
	}).(ServiceKeyOutput)
}

type ServiceKeyMapOutput struct{ *pulumi.OutputState }

func (ServiceKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServiceKey)(nil)).Elem()
}

func (o ServiceKeyMapOutput) ToServiceKeyMapOutput() ServiceKeyMapOutput {
	return o
}

func (o ServiceKeyMapOutput) ToServiceKeyMapOutputWithContext(ctx context.Context) ServiceKeyMapOutput {
	return o
}

func (o ServiceKeyMapOutput) MapIndex(k pulumi.StringInput) ServiceKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServiceKey {
		return vs[0].(map[string]*ServiceKey)[vs[1].(string)]
	}).(ServiceKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKeyInput)(nil)).Elem(), &ServiceKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKeyArrayInput)(nil)).Elem(), ServiceKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceKeyMapInput)(nil)).Elem(), ServiceKeyMap{})
	pulumi.RegisterOutputType(ServiceKeyOutput{})
	pulumi.RegisterOutputType(ServiceKeyArrayOutput{})
	pulumi.RegisterOutputType(ServiceKeyMapOutput{})
}
