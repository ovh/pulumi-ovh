// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ip

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Move struct {
	pulumi.CustomResourceState

	CanBeTerminated pulumi.BoolOutput   `pulumi:"canBeTerminated"`
	Country         pulumi.StringOutput `pulumi:"country"`
	// Custom description on your ip
	Description    pulumi.StringOutput `pulumi:"description"`
	Ip             pulumi.StringOutput `pulumi:"ip"`
	OrganisationId pulumi.StringOutput `pulumi:"organisationId"`
	// Routage information
	RoutedTo    MoveRoutedToOutput  `pulumi:"routedTo"`
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStartDate pulumi.StringOutput `pulumi:"taskStartDate"`
	// Status field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStatus pulumi.StringOutput `pulumi:"taskStatus"`
	// Possible values for ip type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMove registers a new resource with the given unique name, arguments, and options.
func NewMove(ctx *pulumi.Context,
	name string, args *MoveArgs, opts ...pulumi.ResourceOption) (*Move, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Ip == nil {
		return nil, errors.New("invalid value for required argument 'Ip'")
	}
	if args.RoutedTo == nil {
		return nil, errors.New("invalid value for required argument 'RoutedTo'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Move
	err := ctx.RegisterResource("ovh:Ip/move:Move", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMove gets an existing Move resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMove(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MoveState, opts ...pulumi.ResourceOption) (*Move, error) {
	var resource Move
	err := ctx.ReadResource("ovh:Ip/move:Move", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Move resources.
type moveState struct {
	CanBeTerminated *bool   `pulumi:"canBeTerminated"`
	Country         *string `pulumi:"country"`
	// Custom description on your ip
	Description    *string `pulumi:"description"`
	Ip             *string `pulumi:"ip"`
	OrganisationId *string `pulumi:"organisationId"`
	// Routage information
	RoutedTo    *MoveRoutedTo `pulumi:"routedTo"`
	ServiceName *string       `pulumi:"serviceName"`
	// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStartDate *string `pulumi:"taskStartDate"`
	// Status field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStatus *string `pulumi:"taskStatus"`
	// Possible values for ip type
	Type *string `pulumi:"type"`
}

type MoveState struct {
	CanBeTerminated pulumi.BoolPtrInput
	Country         pulumi.StringPtrInput
	// Custom description on your ip
	Description    pulumi.StringPtrInput
	Ip             pulumi.StringPtrInput
	OrganisationId pulumi.StringPtrInput
	// Routage information
	RoutedTo    MoveRoutedToPtrInput
	ServiceName pulumi.StringPtrInput
	// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStartDate pulumi.StringPtrInput
	// Status field of the current IP task that is in charge of changing the service the IP is attached to
	TaskStatus pulumi.StringPtrInput
	// Possible values for ip type
	Type pulumi.StringPtrInput
}

func (MoveState) ElementType() reflect.Type {
	return reflect.TypeOf((*moveState)(nil)).Elem()
}

type moveArgs struct {
	// Custom description on your ip
	Description *string `pulumi:"description"`
	Ip          string  `pulumi:"ip"`
	// Routage information
	RoutedTo MoveRoutedTo `pulumi:"routedTo"`
}

// The set of arguments for constructing a Move resource.
type MoveArgs struct {
	// Custom description on your ip
	Description pulumi.StringPtrInput
	Ip          pulumi.StringInput
	// Routage information
	RoutedTo MoveRoutedToInput
}

func (MoveArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*moveArgs)(nil)).Elem()
}

type MoveInput interface {
	pulumi.Input

	ToMoveOutput() MoveOutput
	ToMoveOutputWithContext(ctx context.Context) MoveOutput
}

func (*Move) ElementType() reflect.Type {
	return reflect.TypeOf((**Move)(nil)).Elem()
}

func (i *Move) ToMoveOutput() MoveOutput {
	return i.ToMoveOutputWithContext(context.Background())
}

func (i *Move) ToMoveOutputWithContext(ctx context.Context) MoveOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveOutput)
}

// MoveArrayInput is an input type that accepts MoveArray and MoveArrayOutput values.
// You can construct a concrete instance of `MoveArrayInput` via:
//
//	MoveArray{ MoveArgs{...} }
type MoveArrayInput interface {
	pulumi.Input

	ToMoveArrayOutput() MoveArrayOutput
	ToMoveArrayOutputWithContext(context.Context) MoveArrayOutput
}

type MoveArray []MoveInput

func (MoveArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Move)(nil)).Elem()
}

func (i MoveArray) ToMoveArrayOutput() MoveArrayOutput {
	return i.ToMoveArrayOutputWithContext(context.Background())
}

func (i MoveArray) ToMoveArrayOutputWithContext(ctx context.Context) MoveArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveArrayOutput)
}

// MoveMapInput is an input type that accepts MoveMap and MoveMapOutput values.
// You can construct a concrete instance of `MoveMapInput` via:
//
//	MoveMap{ "key": MoveArgs{...} }
type MoveMapInput interface {
	pulumi.Input

	ToMoveMapOutput() MoveMapOutput
	ToMoveMapOutputWithContext(context.Context) MoveMapOutput
}

type MoveMap map[string]MoveInput

func (MoveMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Move)(nil)).Elem()
}

func (i MoveMap) ToMoveMapOutput() MoveMapOutput {
	return i.ToMoveMapOutputWithContext(context.Background())
}

func (i MoveMap) ToMoveMapOutputWithContext(ctx context.Context) MoveMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MoveMapOutput)
}

type MoveOutput struct{ *pulumi.OutputState }

func (MoveOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Move)(nil)).Elem()
}

func (o MoveOutput) ToMoveOutput() MoveOutput {
	return o
}

func (o MoveOutput) ToMoveOutputWithContext(ctx context.Context) MoveOutput {
	return o
}

func (o MoveOutput) CanBeTerminated() pulumi.BoolOutput {
	return o.ApplyT(func(v *Move) pulumi.BoolOutput { return v.CanBeTerminated }).(pulumi.BoolOutput)
}

func (o MoveOutput) Country() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.Country }).(pulumi.StringOutput)
}

// Custom description on your ip
func (o MoveOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o MoveOutput) Ip() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.Ip }).(pulumi.StringOutput)
}

func (o MoveOutput) OrganisationId() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.OrganisationId }).(pulumi.StringOutput)
}

// Routage information
func (o MoveOutput) RoutedTo() MoveRoutedToOutput {
	return o.ApplyT(func(v *Move) MoveRoutedToOutput { return v.RoutedTo }).(MoveRoutedToOutput)
}

func (o MoveOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Starting date and time field of the current IP task that is in charge of changing the service the IP is attached to
func (o MoveOutput) TaskStartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.TaskStartDate }).(pulumi.StringOutput)
}

// Status field of the current IP task that is in charge of changing the service the IP is attached to
func (o MoveOutput) TaskStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.TaskStatus }).(pulumi.StringOutput)
}

// Possible values for ip type
func (o MoveOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Move) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type MoveArrayOutput struct{ *pulumi.OutputState }

func (MoveArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Move)(nil)).Elem()
}

func (o MoveArrayOutput) ToMoveArrayOutput() MoveArrayOutput {
	return o
}

func (o MoveArrayOutput) ToMoveArrayOutputWithContext(ctx context.Context) MoveArrayOutput {
	return o
}

func (o MoveArrayOutput) Index(i pulumi.IntInput) MoveOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Move {
		return vs[0].([]*Move)[vs[1].(int)]
	}).(MoveOutput)
}

type MoveMapOutput struct{ *pulumi.OutputState }

func (MoveMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Move)(nil)).Elem()
}

func (o MoveMapOutput) ToMoveMapOutput() MoveMapOutput {
	return o
}

func (o MoveMapOutput) ToMoveMapOutputWithContext(ctx context.Context) MoveMapOutput {
	return o
}

func (o MoveMapOutput) MapIndex(k pulumi.StringInput) MoveOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Move {
		return vs[0].(map[string]*Move)[vs[1].(string)]
	}).(MoveOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MoveInput)(nil)).Elem(), &Move{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveArrayInput)(nil)).Elem(), MoveArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MoveMapInput)(nil)).Elem(), MoveMap{})
	pulumi.RegisterOutputType(MoveOutput{})
	pulumi.RegisterOutputType(MoveArrayOutput{})
	pulumi.RegisterOutputType(MoveMapOutput{})
}
