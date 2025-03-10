// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetInstallationTemplate(ctx *pulumi.Context, args *GetInstallationTemplateArgs, opts ...pulumi.InvokeOption) (*GetInstallationTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetInstallationTemplateResult
	err := ctx.Invoke("ovh:index/getInstallationTemplate:getInstallationTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getInstallationTemplate.
type GetInstallationTemplateArgs struct {
	TemplateName string `pulumi:"templateName"`
}

// A collection of values returned by getInstallationTemplate.
type GetInstallationTemplateResult struct {
	BitFormat             int      `pulumi:"bitFormat"`
	Category              string   `pulumi:"category"`
	Description           string   `pulumi:"description"`
	Distribution          string   `pulumi:"distribution"`
	EndOfInstall          string   `pulumi:"endOfInstall"`
	Family                string   `pulumi:"family"`
	Filesystems           []string `pulumi:"filesystems"`
	HardRaidConfiguration bool     `pulumi:"hardRaidConfiguration"`
	// The provider-assigned unique ID for this managed resource.
	Id                    string                           `pulumi:"id"`
	Inputs                []GetInstallationTemplateInput   `pulumi:"inputs"`
	Licenses              []GetInstallationTemplateLicense `pulumi:"licenses"`
	LvmReady              bool                             `pulumi:"lvmReady"`
	NoPartitioning        bool                             `pulumi:"noPartitioning"`
	Projects              []GetInstallationTemplateProject `pulumi:"projects"`
	SoftRaidOnlyMirroring bool                             `pulumi:"softRaidOnlyMirroring"`
	Subfamily             string                           `pulumi:"subfamily"`
	TemplateName          string                           `pulumi:"templateName"`
}

func GetInstallationTemplateOutput(ctx *pulumi.Context, args GetInstallationTemplateOutputArgs, opts ...pulumi.InvokeOption) GetInstallationTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetInstallationTemplateResultOutput, error) {
			args := v.(GetInstallationTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("ovh:index/getInstallationTemplate:getInstallationTemplate", args, GetInstallationTemplateResultOutput{}, options).(GetInstallationTemplateResultOutput), nil
		}).(GetInstallationTemplateResultOutput)
}

// A collection of arguments for invoking getInstallationTemplate.
type GetInstallationTemplateOutputArgs struct {
	TemplateName pulumi.StringInput `pulumi:"templateName"`
}

func (GetInstallationTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstallationTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getInstallationTemplate.
type GetInstallationTemplateResultOutput struct{ *pulumi.OutputState }

func (GetInstallationTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInstallationTemplateResult)(nil)).Elem()
}

func (o GetInstallationTemplateResultOutput) ToGetInstallationTemplateResultOutput() GetInstallationTemplateResultOutput {
	return o
}

func (o GetInstallationTemplateResultOutput) ToGetInstallationTemplateResultOutputWithContext(ctx context.Context) GetInstallationTemplateResultOutput {
	return o
}

func (o GetInstallationTemplateResultOutput) BitFormat() pulumi.IntOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) int { return v.BitFormat }).(pulumi.IntOutput)
}

func (o GetInstallationTemplateResultOutput) Category() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.Category }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) Distribution() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.Distribution }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) EndOfInstall() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.EndOfInstall }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) Family() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.Family }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) Filesystems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) []string { return v.Filesystems }).(pulumi.StringArrayOutput)
}

func (o GetInstallationTemplateResultOutput) HardRaidConfiguration() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) bool { return v.HardRaidConfiguration }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetInstallationTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) Inputs() GetInstallationTemplateInputArrayOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) []GetInstallationTemplateInput { return v.Inputs }).(GetInstallationTemplateInputArrayOutput)
}

func (o GetInstallationTemplateResultOutput) Licenses() GetInstallationTemplateLicenseArrayOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) []GetInstallationTemplateLicense { return v.Licenses }).(GetInstallationTemplateLicenseArrayOutput)
}

func (o GetInstallationTemplateResultOutput) LvmReady() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) bool { return v.LvmReady }).(pulumi.BoolOutput)
}

func (o GetInstallationTemplateResultOutput) NoPartitioning() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) bool { return v.NoPartitioning }).(pulumi.BoolOutput)
}

func (o GetInstallationTemplateResultOutput) Projects() GetInstallationTemplateProjectArrayOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) []GetInstallationTemplateProject { return v.Projects }).(GetInstallationTemplateProjectArrayOutput)
}

func (o GetInstallationTemplateResultOutput) SoftRaidOnlyMirroring() pulumi.BoolOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) bool { return v.SoftRaidOnlyMirroring }).(pulumi.BoolOutput)
}

func (o GetInstallationTemplateResultOutput) Subfamily() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.Subfamily }).(pulumi.StringOutput)
}

func (o GetInstallationTemplateResultOutput) TemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v GetInstallationTemplateResult) string { return v.TemplateName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetInstallationTemplateResultOutput{})
}
