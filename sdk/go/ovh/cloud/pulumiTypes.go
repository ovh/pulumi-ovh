// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloud

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = internal.GetEnvOrDefault

type GetProjectIam struct {
	// Resource display name
	DisplayName string `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id string `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags map[string]string `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn string `pulumi:"urn"`
}

// GetProjectIamInput is an input type that accepts GetProjectIamArgs and GetProjectIamOutput values.
// You can construct a concrete instance of `GetProjectIamInput` via:
//
//	GetProjectIamArgs{...}
type GetProjectIamInput interface {
	pulumi.Input

	ToGetProjectIamOutput() GetProjectIamOutput
	ToGetProjectIamOutputWithContext(context.Context) GetProjectIamOutput
}

type GetProjectIamArgs struct {
	// Resource display name
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id pulumi.StringInput `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn pulumi.StringInput `pulumi:"urn"`
}

func (GetProjectIamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectIam)(nil)).Elem()
}

func (i GetProjectIamArgs) ToGetProjectIamOutput() GetProjectIamOutput {
	return i.ToGetProjectIamOutputWithContext(context.Background())
}

func (i GetProjectIamArgs) ToGetProjectIamOutputWithContext(ctx context.Context) GetProjectIamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectIamOutput)
}

type GetProjectIamOutput struct{ *pulumi.OutputState }

func (GetProjectIamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectIam)(nil)).Elem()
}

func (o GetProjectIamOutput) ToGetProjectIamOutput() GetProjectIamOutput {
	return o
}

func (o GetProjectIamOutput) ToGetProjectIamOutputWithContext(ctx context.Context) GetProjectIamOutput {
	return o
}

// Resource display name
func (o GetProjectIamOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectIam) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Unique identifier of the resource in the IAM
func (o GetProjectIamOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectIam) string { return v.Id }).(pulumi.StringOutput)
}

// Resource tags. Tags that were internally computed are prefixed with `ovh:`
func (o GetProjectIamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetProjectIam) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// URN of the private database, used when writing IAM policies
func (o GetProjectIamOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectIam) string { return v.Urn }).(pulumi.StringOutput)
}

type GetProjectsProject struct {
	// Project access
	Access string `pulumi:"access"`
	// Project creation date
	CreationDate string `pulumi:"creationDate"`
	// Description of your project
	Description string `pulumi:"description"`
	// Expiration date of your project. After this date, your project will be deleted
	Expiration string `pulumi:"expiration"`
	// IAM resource information
	Iam GetProjectsProjectIam `pulumi:"iam"`
	// Manual quota prevent automatic quota upgrade
	ManualQuota bool `pulumi:"manualQuota"`
	// Project order ID
	OrderId float64 `pulumi:"orderId"`
	// Order plan code
	PlanCode string `pulumi:"planCode"`
	// Project ID
	ProjectId string `pulumi:"projectId"`
	// Project name
	ProjectName string `pulumi:"projectName"`
	// ID of the public cloud project
	ServiceName string `pulumi:"serviceName"`
	// Current status
	Status string `pulumi:"status"`
	// Project unleashed
	Unleash bool `pulumi:"unleash"`
}

// GetProjectsProjectInput is an input type that accepts GetProjectsProjectArgs and GetProjectsProjectOutput values.
// You can construct a concrete instance of `GetProjectsProjectInput` via:
//
//	GetProjectsProjectArgs{...}
type GetProjectsProjectInput interface {
	pulumi.Input

	ToGetProjectsProjectOutput() GetProjectsProjectOutput
	ToGetProjectsProjectOutputWithContext(context.Context) GetProjectsProjectOutput
}

type GetProjectsProjectArgs struct {
	// Project access
	Access pulumi.StringInput `pulumi:"access"`
	// Project creation date
	CreationDate pulumi.StringInput `pulumi:"creationDate"`
	// Description of your project
	Description pulumi.StringInput `pulumi:"description"`
	// Expiration date of your project. After this date, your project will be deleted
	Expiration pulumi.StringInput `pulumi:"expiration"`
	// IAM resource information
	Iam GetProjectsProjectIamInput `pulumi:"iam"`
	// Manual quota prevent automatic quota upgrade
	ManualQuota pulumi.BoolInput `pulumi:"manualQuota"`
	// Project order ID
	OrderId pulumi.Float64Input `pulumi:"orderId"`
	// Order plan code
	PlanCode pulumi.StringInput `pulumi:"planCode"`
	// Project ID
	ProjectId pulumi.StringInput `pulumi:"projectId"`
	// Project name
	ProjectName pulumi.StringInput `pulumi:"projectName"`
	// ID of the public cloud project
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Current status
	Status pulumi.StringInput `pulumi:"status"`
	// Project unleashed
	Unleash pulumi.BoolInput `pulumi:"unleash"`
}

func (GetProjectsProjectArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProject)(nil)).Elem()
}

func (i GetProjectsProjectArgs) ToGetProjectsProjectOutput() GetProjectsProjectOutput {
	return i.ToGetProjectsProjectOutputWithContext(context.Background())
}

func (i GetProjectsProjectArgs) ToGetProjectsProjectOutputWithContext(ctx context.Context) GetProjectsProjectOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectOutput)
}

// GetProjectsProjectArrayInput is an input type that accepts GetProjectsProjectArray and GetProjectsProjectArrayOutput values.
// You can construct a concrete instance of `GetProjectsProjectArrayInput` via:
//
//	GetProjectsProjectArray{ GetProjectsProjectArgs{...} }
type GetProjectsProjectArrayInput interface {
	pulumi.Input

	ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput
	ToGetProjectsProjectArrayOutputWithContext(context.Context) GetProjectsProjectArrayOutput
}

type GetProjectsProjectArray []GetProjectsProjectInput

func (GetProjectsProjectArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectsProject)(nil)).Elem()
}

func (i GetProjectsProjectArray) ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput {
	return i.ToGetProjectsProjectArrayOutputWithContext(context.Background())
}

func (i GetProjectsProjectArray) ToGetProjectsProjectArrayOutputWithContext(ctx context.Context) GetProjectsProjectArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectArrayOutput)
}

type GetProjectsProjectOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProject)(nil)).Elem()
}

func (o GetProjectsProjectOutput) ToGetProjectsProjectOutput() GetProjectsProjectOutput {
	return o
}

func (o GetProjectsProjectOutput) ToGetProjectsProjectOutputWithContext(ctx context.Context) GetProjectsProjectOutput {
	return o
}

// Project access
func (o GetProjectsProjectOutput) Access() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Access }).(pulumi.StringOutput)
}

// Project creation date
func (o GetProjectsProjectOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.CreationDate }).(pulumi.StringOutput)
}

// Description of your project
func (o GetProjectsProjectOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Description }).(pulumi.StringOutput)
}

// Expiration date of your project. After this date, your project will be deleted
func (o GetProjectsProjectOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Expiration }).(pulumi.StringOutput)
}

// IAM resource information
func (o GetProjectsProjectOutput) Iam() GetProjectsProjectIamOutput {
	return o.ApplyT(func(v GetProjectsProject) GetProjectsProjectIam { return v.Iam }).(GetProjectsProjectIamOutput)
}

// Manual quota prevent automatic quota upgrade
func (o GetProjectsProjectOutput) ManualQuota() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectsProject) bool { return v.ManualQuota }).(pulumi.BoolOutput)
}

// Project order ID
func (o GetProjectsProjectOutput) OrderId() pulumi.Float64Output {
	return o.ApplyT(func(v GetProjectsProject) float64 { return v.OrderId }).(pulumi.Float64Output)
}

// Order plan code
func (o GetProjectsProjectOutput) PlanCode() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.PlanCode }).(pulumi.StringOutput)
}

// Project ID
func (o GetProjectsProjectOutput) ProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.ProjectId }).(pulumi.StringOutput)
}

// Project name
func (o GetProjectsProjectOutput) ProjectName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.ProjectName }).(pulumi.StringOutput)
}

// ID of the public cloud project
func (o GetProjectsProjectOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.ServiceName }).(pulumi.StringOutput)
}

// Current status
func (o GetProjectsProjectOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProject) string { return v.Status }).(pulumi.StringOutput)
}

// Project unleashed
func (o GetProjectsProjectOutput) Unleash() pulumi.BoolOutput {
	return o.ApplyT(func(v GetProjectsProject) bool { return v.Unleash }).(pulumi.BoolOutput)
}

type GetProjectsProjectArrayOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetProjectsProject)(nil)).Elem()
}

func (o GetProjectsProjectArrayOutput) ToGetProjectsProjectArrayOutput() GetProjectsProjectArrayOutput {
	return o
}

func (o GetProjectsProjectArrayOutput) ToGetProjectsProjectArrayOutputWithContext(ctx context.Context) GetProjectsProjectArrayOutput {
	return o
}

func (o GetProjectsProjectArrayOutput) Index(i pulumi.IntInput) GetProjectsProjectOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetProjectsProject {
		return vs[0].([]GetProjectsProject)[vs[1].(int)]
	}).(GetProjectsProjectOutput)
}

type GetProjectsProjectIam struct {
	// Resource display name
	DisplayName string `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id string `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags map[string]string `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn string `pulumi:"urn"`
}

// GetProjectsProjectIamInput is an input type that accepts GetProjectsProjectIamArgs and GetProjectsProjectIamOutput values.
// You can construct a concrete instance of `GetProjectsProjectIamInput` via:
//
//	GetProjectsProjectIamArgs{...}
type GetProjectsProjectIamInput interface {
	pulumi.Input

	ToGetProjectsProjectIamOutput() GetProjectsProjectIamOutput
	ToGetProjectsProjectIamOutputWithContext(context.Context) GetProjectsProjectIamOutput
}

type GetProjectsProjectIamArgs struct {
	// Resource display name
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	// Unique identifier of the resource in the IAM
	Id pulumi.StringInput `pulumi:"id"`
	// Resource tags. Tags that were internally computed are prefixed with `ovh:`
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// URN of the private database, used when writing IAM policies
	Urn pulumi.StringInput `pulumi:"urn"`
}

func (GetProjectsProjectIamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProjectIam)(nil)).Elem()
}

func (i GetProjectsProjectIamArgs) ToGetProjectsProjectIamOutput() GetProjectsProjectIamOutput {
	return i.ToGetProjectsProjectIamOutputWithContext(context.Background())
}

func (i GetProjectsProjectIamArgs) ToGetProjectsProjectIamOutputWithContext(ctx context.Context) GetProjectsProjectIamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetProjectsProjectIamOutput)
}

type GetProjectsProjectIamOutput struct{ *pulumi.OutputState }

func (GetProjectsProjectIamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetProjectsProjectIam)(nil)).Elem()
}

func (o GetProjectsProjectIamOutput) ToGetProjectsProjectIamOutput() GetProjectsProjectIamOutput {
	return o
}

func (o GetProjectsProjectIamOutput) ToGetProjectsProjectIamOutputWithContext(ctx context.Context) GetProjectsProjectIamOutput {
	return o
}

// Resource display name
func (o GetProjectsProjectIamOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProjectIam) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Unique identifier of the resource in the IAM
func (o GetProjectsProjectIamOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProjectIam) string { return v.Id }).(pulumi.StringOutput)
}

// Resource tags. Tags that were internally computed are prefixed with `ovh:`
func (o GetProjectsProjectIamOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetProjectsProjectIam) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// URN of the private database, used when writing IAM policies
func (o GetProjectsProjectIamOutput) Urn() pulumi.StringOutput {
	return o.ApplyT(func(v GetProjectsProjectIam) string { return v.Urn }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectIamInput)(nil)).Elem(), GetProjectIamArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectInput)(nil)).Elem(), GetProjectsProjectArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectArrayInput)(nil)).Elem(), GetProjectsProjectArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetProjectsProjectIamInput)(nil)).Elem(), GetProjectsProjectIamArgs{})
	pulumi.RegisterOutputType(GetProjectIamOutput{})
	pulumi.RegisterOutputType(GetProjectsProjectOutput{})
	pulumi.RegisterOutputType(GetProjectsProjectArrayOutput{})
	pulumi.RegisterOutputType(GetProjectsProjectIamOutput{})
}
