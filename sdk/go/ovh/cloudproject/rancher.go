// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage a Rancher service in a public cloud project.
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
//			rancher, err := cloudproject.NewRancher(ctx, "rancher", &cloudproject.RancherArgs{
//				ProjectId: pulumi.String("<public cloud project ID>"),
//				TargetSpec: &cloudproject.RancherTargetSpecArgs{
//					Name: pulumi.String("MyRancher"),
//					Plan: pulumi.String("STANDARD"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("rancherUrl", rancher.CurrentState.ApplyT(func(currentState cloudproject.RancherCurrentState) (*string, error) {
//				return &currentState.Url, nil
//			}).(pulumi.StringPtrOutput))
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// A share in a public cloud project can be imported using the `project_id` and `id` attributes. Using the following configuration:
//
// terraform
//
// import {
//
//	id = "<project_id>/<id>"
//
//	to = ovh_cloud_project_rancher.rancher
//
// }
//
// You can then run:
//
// bash
//
// $ pulumi preview -generate-config-out=rancher.tf
//
// $ pulumi up
//
// The file `rancher.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
type Rancher struct {
	pulumi.CustomResourceState

	// Date of the managed Rancher service creation
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Current configuration applied to the managed Rancher service
	CurrentState RancherCurrentStateOutput `pulumi:"currentState"`
	// Asynchronous operations ongoing on the managed Rancher service
	CurrentTasks RancherCurrentTaskArrayOutput `pulumi:"currentTasks"`
	// Project ID
	ProjectId pulumi.StringOutput `pulumi:"projectId"`
	// Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
	ResourceStatus pulumi.StringOutput `pulumi:"resourceStatus"`
	// Target specification for the managed Rancher service
	TargetSpec RancherTargetSpecOutput `pulumi:"targetSpec"`
	// Date of the last managed Rancher service update
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewRancher registers a new resource with the given unique name, arguments, and options.
func NewRancher(ctx *pulumi.Context,
	name string, args *RancherArgs, opts ...pulumi.ResourceOption) (*Rancher, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProjectId == nil {
		return nil, errors.New("invalid value for required argument 'ProjectId'")
	}
	if args.TargetSpec == nil {
		return nil, errors.New("invalid value for required argument 'TargetSpec'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Rancher
	err := ctx.RegisterResource("ovh:CloudProject/rancher:Rancher", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRancher gets an existing Rancher resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRancher(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RancherState, opts ...pulumi.ResourceOption) (*Rancher, error) {
	var resource Rancher
	err := ctx.ReadResource("ovh:CloudProject/rancher:Rancher", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Rancher resources.
type rancherState struct {
	// Date of the managed Rancher service creation
	CreatedAt *string `pulumi:"createdAt"`
	// Current configuration applied to the managed Rancher service
	CurrentState *RancherCurrentState `pulumi:"currentState"`
	// Asynchronous operations ongoing on the managed Rancher service
	CurrentTasks []RancherCurrentTask `pulumi:"currentTasks"`
	// Project ID
	ProjectId *string `pulumi:"projectId"`
	// Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
	ResourceStatus *string `pulumi:"resourceStatus"`
	// Target specification for the managed Rancher service
	TargetSpec *RancherTargetSpec `pulumi:"targetSpec"`
	// Date of the last managed Rancher service update
	UpdatedAt *string `pulumi:"updatedAt"`
}

type RancherState struct {
	// Date of the managed Rancher service creation
	CreatedAt pulumi.StringPtrInput
	// Current configuration applied to the managed Rancher service
	CurrentState RancherCurrentStatePtrInput
	// Asynchronous operations ongoing on the managed Rancher service
	CurrentTasks RancherCurrentTaskArrayInput
	// Project ID
	ProjectId pulumi.StringPtrInput
	// Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
	ResourceStatus pulumi.StringPtrInput
	// Target specification for the managed Rancher service
	TargetSpec RancherTargetSpecPtrInput
	// Date of the last managed Rancher service update
	UpdatedAt pulumi.StringPtrInput
}

func (RancherState) ElementType() reflect.Type {
	return reflect.TypeOf((*rancherState)(nil)).Elem()
}

type rancherArgs struct {
	// Project ID
	ProjectId string `pulumi:"projectId"`
	// Target specification for the managed Rancher service
	TargetSpec RancherTargetSpec `pulumi:"targetSpec"`
}

// The set of arguments for constructing a Rancher resource.
type RancherArgs struct {
	// Project ID
	ProjectId pulumi.StringInput
	// Target specification for the managed Rancher service
	TargetSpec RancherTargetSpecInput
}

func (RancherArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rancherArgs)(nil)).Elem()
}

type RancherInput interface {
	pulumi.Input

	ToRancherOutput() RancherOutput
	ToRancherOutputWithContext(ctx context.Context) RancherOutput
}

func (*Rancher) ElementType() reflect.Type {
	return reflect.TypeOf((**Rancher)(nil)).Elem()
}

func (i *Rancher) ToRancherOutput() RancherOutput {
	return i.ToRancherOutputWithContext(context.Background())
}

func (i *Rancher) ToRancherOutputWithContext(ctx context.Context) RancherOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RancherOutput)
}

// RancherArrayInput is an input type that accepts RancherArray and RancherArrayOutput values.
// You can construct a concrete instance of `RancherArrayInput` via:
//
//	RancherArray{ RancherArgs{...} }
type RancherArrayInput interface {
	pulumi.Input

	ToRancherArrayOutput() RancherArrayOutput
	ToRancherArrayOutputWithContext(context.Context) RancherArrayOutput
}

type RancherArray []RancherInput

func (RancherArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rancher)(nil)).Elem()
}

func (i RancherArray) ToRancherArrayOutput() RancherArrayOutput {
	return i.ToRancherArrayOutputWithContext(context.Background())
}

func (i RancherArray) ToRancherArrayOutputWithContext(ctx context.Context) RancherArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RancherArrayOutput)
}

// RancherMapInput is an input type that accepts RancherMap and RancherMapOutput values.
// You can construct a concrete instance of `RancherMapInput` via:
//
//	RancherMap{ "key": RancherArgs{...} }
type RancherMapInput interface {
	pulumi.Input

	ToRancherMapOutput() RancherMapOutput
	ToRancherMapOutputWithContext(context.Context) RancherMapOutput
}

type RancherMap map[string]RancherInput

func (RancherMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rancher)(nil)).Elem()
}

func (i RancherMap) ToRancherMapOutput() RancherMapOutput {
	return i.ToRancherMapOutputWithContext(context.Background())
}

func (i RancherMap) ToRancherMapOutputWithContext(ctx context.Context) RancherMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RancherMapOutput)
}

type RancherOutput struct{ *pulumi.OutputState }

func (RancherOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Rancher)(nil)).Elem()
}

func (o RancherOutput) ToRancherOutput() RancherOutput {
	return o
}

func (o RancherOutput) ToRancherOutputWithContext(ctx context.Context) RancherOutput {
	return o
}

// Date of the managed Rancher service creation
func (o RancherOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Rancher) pulumi.StringOutput { return v.CreatedAt }).(pulumi.StringOutput)
}

// Current configuration applied to the managed Rancher service
func (o RancherOutput) CurrentState() RancherCurrentStateOutput {
	return o.ApplyT(func(v *Rancher) RancherCurrentStateOutput { return v.CurrentState }).(RancherCurrentStateOutput)
}

// Asynchronous operations ongoing on the managed Rancher service
func (o RancherOutput) CurrentTasks() RancherCurrentTaskArrayOutput {
	return o.ApplyT(func(v *Rancher) RancherCurrentTaskArrayOutput { return v.CurrentTasks }).(RancherCurrentTaskArrayOutput)
}

// Project ID
func (o RancherOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Rancher) pulumi.StringOutput { return v.ProjectId }).(pulumi.StringOutput)
}

// Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
func (o RancherOutput) ResourceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *Rancher) pulumi.StringOutput { return v.ResourceStatus }).(pulumi.StringOutput)
}

// Target specification for the managed Rancher service
func (o RancherOutput) TargetSpec() RancherTargetSpecOutput {
	return o.ApplyT(func(v *Rancher) RancherTargetSpecOutput { return v.TargetSpec }).(RancherTargetSpecOutput)
}

// Date of the last managed Rancher service update
func (o RancherOutput) UpdatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Rancher) pulumi.StringOutput { return v.UpdatedAt }).(pulumi.StringOutput)
}

type RancherArrayOutput struct{ *pulumi.OutputState }

func (RancherArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Rancher)(nil)).Elem()
}

func (o RancherArrayOutput) ToRancherArrayOutput() RancherArrayOutput {
	return o
}

func (o RancherArrayOutput) ToRancherArrayOutputWithContext(ctx context.Context) RancherArrayOutput {
	return o
}

func (o RancherArrayOutput) Index(i pulumi.IntInput) RancherOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Rancher {
		return vs[0].([]*Rancher)[vs[1].(int)]
	}).(RancherOutput)
}

type RancherMapOutput struct{ *pulumi.OutputState }

func (RancherMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Rancher)(nil)).Elem()
}

func (o RancherMapOutput) ToRancherMapOutput() RancherMapOutput {
	return o
}

func (o RancherMapOutput) ToRancherMapOutputWithContext(ctx context.Context) RancherMapOutput {
	return o
}

func (o RancherMapOutput) MapIndex(k pulumi.StringInput) RancherOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Rancher {
		return vs[0].(map[string]*Rancher)[vs[1].(string)]
	}).(RancherOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RancherInput)(nil)).Elem(), &Rancher{})
	pulumi.RegisterInputType(reflect.TypeOf((*RancherArrayInput)(nil)).Elem(), RancherArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RancherMapInput)(nil)).Elem(), RancherMap{})
	pulumi.RegisterOutputType(RancherOutput{})
	pulumi.RegisterOutputType(RancherArrayOutput{})
	pulumi.RegisterOutputType(RancherMapOutput{})
}
