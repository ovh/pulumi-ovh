// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package dedicated

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// Installation task can be imported using the `service_name` (`nsXXXX.ip...`) of the baremetal server, the `operating_system` used  and ths `task_id`, separated by "/" E.g.,
//
// bash
//
// ```sh
// $ pulumi import ovh:Dedicated/serverReinstallTask:ServerReinstallTask ovh_dedicated_server_reinstall_task nsXXXX.ipXXXX/operating_system/12345
// ```
type ServerReinstallTask struct {
	pulumi.CustomResourceState

	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy pulumi.IntPtrOutput `pulumi:"bootidOnDestroy"`
	// Details of this task. (should be `Install asked`)
	Comment pulumi.StringOutput `pulumi:"comment"`
	// Available attributes and their types are OS-dependant. Example: `hostname`.
	//
	// > __WARNING__ Some customizations may be required on some Operating Systems.  [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
	Customizations ServerReinstallTaskCustomizationsPtrOutput `pulumi:"customizations"`
	// Completion date in RFC3339 format.
	DoneDate pulumi.StringOutput `pulumi:"doneDate"`
	// Function name (should be `hardInstall`).
	Function pulumi.StringOutput `pulumi:"function"`
	// Last update
	LastUpdate pulumi.StringOutput `pulumi:"lastUpdate"`
	// Operating system to install.
	Os pulumi.StringOutput `pulumi:"os"`
	// Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Task creation date in RFC3339 format.
	StartDate pulumi.StringOutput `pulumi:"startDate"`
	// Task status (should be `done`)
	Status pulumi.StringOutput `pulumi:"status"`
	// OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
	Storages ServerReinstallTaskStorageArrayOutput `pulumi:"storages"`
}

// NewServerReinstallTask registers a new resource with the given unique name, arguments, and options.
func NewServerReinstallTask(ctx *pulumi.Context,
	name string, args *ServerReinstallTaskArgs, opts ...pulumi.ResourceOption) (*ServerReinstallTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Os == nil {
		return nil, errors.New("invalid value for required argument 'Os'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ServerReinstallTask
	err := ctx.RegisterResource("ovh:Dedicated/serverReinstallTask:ServerReinstallTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetServerReinstallTask gets an existing ServerReinstallTask resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetServerReinstallTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServerReinstallTaskState, opts ...pulumi.ResourceOption) (*ServerReinstallTask, error) {
	var resource ServerReinstallTask
	err := ctx.ReadResource("ovh:Dedicated/serverReinstallTask:ServerReinstallTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ServerReinstallTask resources.
type serverReinstallTaskState struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy *int `pulumi:"bootidOnDestroy"`
	// Details of this task. (should be `Install asked`)
	Comment *string `pulumi:"comment"`
	// Available attributes and their types are OS-dependant. Example: `hostname`.
	//
	// > __WARNING__ Some customizations may be required on some Operating Systems.  [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
	Customizations *ServerReinstallTaskCustomizations `pulumi:"customizations"`
	// Completion date in RFC3339 format.
	DoneDate *string `pulumi:"doneDate"`
	// Function name (should be `hardInstall`).
	Function *string `pulumi:"function"`
	// Last update
	LastUpdate *string `pulumi:"lastUpdate"`
	// Operating system to install.
	Os *string `pulumi:"os"`
	// Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
	Properties map[string]string `pulumi:"properties"`
	// The serviceName of your dedicated server.
	ServiceName *string `pulumi:"serviceName"`
	// Task creation date in RFC3339 format.
	StartDate *string `pulumi:"startDate"`
	// Task status (should be `done`)
	Status *string `pulumi:"status"`
	// OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
	Storages []ServerReinstallTaskStorage `pulumi:"storages"`
}

type ServerReinstallTaskState struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy pulumi.IntPtrInput
	// Details of this task. (should be `Install asked`)
	Comment pulumi.StringPtrInput
	// Available attributes and their types are OS-dependant. Example: `hostname`.
	//
	// > __WARNING__ Some customizations may be required on some Operating Systems.  [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
	Customizations ServerReinstallTaskCustomizationsPtrInput
	// Completion date in RFC3339 format.
	DoneDate pulumi.StringPtrInput
	// Function name (should be `hardInstall`).
	Function pulumi.StringPtrInput
	// Last update
	LastUpdate pulumi.StringPtrInput
	// Operating system to install.
	Os pulumi.StringPtrInput
	// Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
	Properties pulumi.StringMapInput
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringPtrInput
	// Task creation date in RFC3339 format.
	StartDate pulumi.StringPtrInput
	// Task status (should be `done`)
	Status pulumi.StringPtrInput
	// OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
	Storages ServerReinstallTaskStorageArrayInput
}

func (ServerReinstallTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*serverReinstallTaskState)(nil)).Elem()
}

type serverReinstallTaskArgs struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy *int `pulumi:"bootidOnDestroy"`
	// Available attributes and their types are OS-dependant. Example: `hostname`.
	//
	// > __WARNING__ Some customizations may be required on some Operating Systems.  [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
	Customizations *ServerReinstallTaskCustomizations `pulumi:"customizations"`
	// Operating system to install.
	Os string `pulumi:"os"`
	// Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
	Properties map[string]string `pulumi:"properties"`
	// The serviceName of your dedicated server.
	ServiceName string `pulumi:"serviceName"`
	// OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
	Storages []ServerReinstallTaskStorage `pulumi:"storages"`
}

// The set of arguments for constructing a ServerReinstallTask resource.
type ServerReinstallTaskArgs struct {
	// If set, reboot the server on the specified boot id during destroy phase.
	BootidOnDestroy pulumi.IntPtrInput
	// Available attributes and their types are OS-dependant. Example: `hostname`.
	//
	// > __WARNING__ Some customizations may be required on some Operating Systems.  [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
	Customizations ServerReinstallTaskCustomizationsPtrInput
	// Operating system to install.
	Os pulumi.StringInput
	// Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
	Properties pulumi.StringMapInput
	// The serviceName of your dedicated server.
	ServiceName pulumi.StringInput
	// OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
	Storages ServerReinstallTaskStorageArrayInput
}

func (ServerReinstallTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serverReinstallTaskArgs)(nil)).Elem()
}

type ServerReinstallTaskInput interface {
	pulumi.Input

	ToServerReinstallTaskOutput() ServerReinstallTaskOutput
	ToServerReinstallTaskOutputWithContext(ctx context.Context) ServerReinstallTaskOutput
}

func (*ServerReinstallTask) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerReinstallTask)(nil)).Elem()
}

func (i *ServerReinstallTask) ToServerReinstallTaskOutput() ServerReinstallTaskOutput {
	return i.ToServerReinstallTaskOutputWithContext(context.Background())
}

func (i *ServerReinstallTask) ToServerReinstallTaskOutputWithContext(ctx context.Context) ServerReinstallTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerReinstallTaskOutput)
}

// ServerReinstallTaskArrayInput is an input type that accepts ServerReinstallTaskArray and ServerReinstallTaskArrayOutput values.
// You can construct a concrete instance of `ServerReinstallTaskArrayInput` via:
//
//	ServerReinstallTaskArray{ ServerReinstallTaskArgs{...} }
type ServerReinstallTaskArrayInput interface {
	pulumi.Input

	ToServerReinstallTaskArrayOutput() ServerReinstallTaskArrayOutput
	ToServerReinstallTaskArrayOutputWithContext(context.Context) ServerReinstallTaskArrayOutput
}

type ServerReinstallTaskArray []ServerReinstallTaskInput

func (ServerReinstallTaskArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerReinstallTask)(nil)).Elem()
}

func (i ServerReinstallTaskArray) ToServerReinstallTaskArrayOutput() ServerReinstallTaskArrayOutput {
	return i.ToServerReinstallTaskArrayOutputWithContext(context.Background())
}

func (i ServerReinstallTaskArray) ToServerReinstallTaskArrayOutputWithContext(ctx context.Context) ServerReinstallTaskArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerReinstallTaskArrayOutput)
}

// ServerReinstallTaskMapInput is an input type that accepts ServerReinstallTaskMap and ServerReinstallTaskMapOutput values.
// You can construct a concrete instance of `ServerReinstallTaskMapInput` via:
//
//	ServerReinstallTaskMap{ "key": ServerReinstallTaskArgs{...} }
type ServerReinstallTaskMapInput interface {
	pulumi.Input

	ToServerReinstallTaskMapOutput() ServerReinstallTaskMapOutput
	ToServerReinstallTaskMapOutputWithContext(context.Context) ServerReinstallTaskMapOutput
}

type ServerReinstallTaskMap map[string]ServerReinstallTaskInput

func (ServerReinstallTaskMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerReinstallTask)(nil)).Elem()
}

func (i ServerReinstallTaskMap) ToServerReinstallTaskMapOutput() ServerReinstallTaskMapOutput {
	return i.ToServerReinstallTaskMapOutputWithContext(context.Background())
}

func (i ServerReinstallTaskMap) ToServerReinstallTaskMapOutputWithContext(ctx context.Context) ServerReinstallTaskMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerReinstallTaskMapOutput)
}

type ServerReinstallTaskOutput struct{ *pulumi.OutputState }

func (ServerReinstallTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerReinstallTask)(nil)).Elem()
}

func (o ServerReinstallTaskOutput) ToServerReinstallTaskOutput() ServerReinstallTaskOutput {
	return o
}

func (o ServerReinstallTaskOutput) ToServerReinstallTaskOutputWithContext(ctx context.Context) ServerReinstallTaskOutput {
	return o
}

// If set, reboot the server on the specified boot id during destroy phase.
func (o ServerReinstallTaskOutput) BootidOnDestroy() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.IntPtrOutput { return v.BootidOnDestroy }).(pulumi.IntPtrOutput)
}

// Details of this task. (should be `Install asked`)
func (o ServerReinstallTaskOutput) Comment() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.Comment }).(pulumi.StringOutput)
}

// Available attributes and their types are OS-dependant. Example: `hostname`.
//
// > __WARNING__ Some customizations may be required on some Operating Systems.  [Check how to list the available and required customization(s) for your operating system](https://help.ovhcloud.com/csm/en-dedicated-servers-api-os-installation?id=kb_article_view&sysparm_article=KB0061951#os-inputs) (do not forget to adapt camel case customization name to snake case parameter).
func (o ServerReinstallTaskOutput) Customizations() ServerReinstallTaskCustomizationsPtrOutput {
	return o.ApplyT(func(v *ServerReinstallTask) ServerReinstallTaskCustomizationsPtrOutput { return v.Customizations }).(ServerReinstallTaskCustomizationsPtrOutput)
}

// Completion date in RFC3339 format.
func (o ServerReinstallTaskOutput) DoneDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.DoneDate }).(pulumi.StringOutput)
}

// Function name (should be `hardInstall`).
func (o ServerReinstallTaskOutput) Function() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.Function }).(pulumi.StringOutput)
}

// Last update
func (o ServerReinstallTaskOutput) LastUpdate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.LastUpdate }).(pulumi.StringOutput)
}

// Operating system to install.
func (o ServerReinstallTaskOutput) Os() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.Os }).(pulumi.StringOutput)
}

// Arbitrary properties to pass to cloud-init's config drive datasource. It supports any key with any string value.
func (o ServerReinstallTaskOutput) Properties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringMapOutput { return v.Properties }).(pulumi.StringMapOutput)
}

// The serviceName of your dedicated server.
func (o ServerReinstallTaskOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Task creation date in RFC3339 format.
func (o ServerReinstallTaskOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

// Task status (should be `done`)
func (o ServerReinstallTaskOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *ServerReinstallTask) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// OS reinstallation storage configurations. [More details about disks, hardware/software RAID and partitioning configuration](https://help.ovhcloud.com/csm/en-dedicated-servers-api-partitioning?id=kb_article_view&sysparm_article=KB0043882) (do not forget to adapt camel case parameters to snake case parameters).
func (o ServerReinstallTaskOutput) Storages() ServerReinstallTaskStorageArrayOutput {
	return o.ApplyT(func(v *ServerReinstallTask) ServerReinstallTaskStorageArrayOutput { return v.Storages }).(ServerReinstallTaskStorageArrayOutput)
}

type ServerReinstallTaskArrayOutput struct{ *pulumi.OutputState }

func (ServerReinstallTaskArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ServerReinstallTask)(nil)).Elem()
}

func (o ServerReinstallTaskArrayOutput) ToServerReinstallTaskArrayOutput() ServerReinstallTaskArrayOutput {
	return o
}

func (o ServerReinstallTaskArrayOutput) ToServerReinstallTaskArrayOutputWithContext(ctx context.Context) ServerReinstallTaskArrayOutput {
	return o
}

func (o ServerReinstallTaskArrayOutput) Index(i pulumi.IntInput) ServerReinstallTaskOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ServerReinstallTask {
		return vs[0].([]*ServerReinstallTask)[vs[1].(int)]
	}).(ServerReinstallTaskOutput)
}

type ServerReinstallTaskMapOutput struct{ *pulumi.OutputState }

func (ServerReinstallTaskMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ServerReinstallTask)(nil)).Elem()
}

func (o ServerReinstallTaskMapOutput) ToServerReinstallTaskMapOutput() ServerReinstallTaskMapOutput {
	return o
}

func (o ServerReinstallTaskMapOutput) ToServerReinstallTaskMapOutputWithContext(ctx context.Context) ServerReinstallTaskMapOutput {
	return o
}

func (o ServerReinstallTaskMapOutput) MapIndex(k pulumi.StringInput) ServerReinstallTaskOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ServerReinstallTask {
		return vs[0].(map[string]*ServerReinstallTask)[vs[1].(string)]
	}).(ServerReinstallTaskOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServerReinstallTaskInput)(nil)).Elem(), &ServerReinstallTask{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerReinstallTaskArrayInput)(nil)).Elem(), ServerReinstallTaskArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServerReinstallTaskMapInput)(nil)).Elem(), ServerReinstallTaskMap{})
	pulumi.RegisterOutputType(ServerReinstallTaskOutput{})
	pulumi.RegisterOutputType(ServerReinstallTaskArrayOutput{})
	pulumi.RegisterOutputType(ServerReinstallTaskMapOutput{})
}
