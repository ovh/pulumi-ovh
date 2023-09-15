// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Creates an OIDC configuration in an OVHcloud Managed Private Registry.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/go/ovh/CloudProject"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := CloudProject.NewContainerRegistryOIDC(ctx, "my-oidc", &CloudProject.ContainerRegistryOIDCArgs{
//				ServiceName:      pulumi.String("XXXXXX"),
//				RegistryId:       pulumi.String("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx"),
//				OidcName:         pulumi.String("my-oidc-provider"),
//				OidcEndpoint:     pulumi.String("https://xxxx.yyy.com"),
//				OidcClientId:     pulumi.String("xxx"),
//				OidcClientSecret: pulumi.String("xxx"),
//				OidcScope:        pulumi.String("openid,profile,email,offline_access"),
//				OidcGroupsClaim:  pulumi.String("groups"),
//				OidcAdminGroup:   pulumi.String("harbor-admin"),
//				OidcVerifyCert:   pulumi.Bool(true),
//				OidcAutoOnboard:  pulumi.Bool(true),
//				OidcUserClaim:    pulumi.String("preferred_username"),
//				DeleteUsers:      pulumi.Bool(false),
//			})
//			if err != nil {
//				return err
//			}
//			ctx.Export("oidcClientSecret", my_oidc.OidcClientSecret)
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OVHcloud Managed Private Registry OIDC can be imported using the tenant `service_name` and registry id `registry_id` separated by "/" E.g., bash
//
// ```sh
//
//	$ pulumi import ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC my-oidc service_name/registry_id
//
// ```
type ContainerRegistryOIDC struct {
	pulumi.CustomResourceState

	// Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
	DeleteUsers pulumi.BoolPtrOutput `pulumi:"deleteUsers"`
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup pulumi.StringPtrOutput `pulumi:"oidcAdminGroup"`
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard pulumi.BoolPtrOutput `pulumi:"oidcAutoOnboard"`
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId pulumi.StringOutput `pulumi:"oidcClientId"`
	// The secret for the Harbor client application.
	OidcClientSecret pulumi.StringOutput `pulumi:"oidcClientSecret"`
	// The URL of an OIDC-compliant server.
	OidcEndpoint pulumi.StringOutput `pulumi:"oidcEndpoint"`
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim pulumi.StringPtrOutput `pulumi:"oidcGroupsClaim"`
	// The name of the OIDC provider.
	OidcName pulumi.StringOutput `pulumi:"oidcName"`
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope pulumi.StringOutput `pulumi:"oidcScope"`
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim pulumi.StringPtrOutput `pulumi:"oidcUserClaim"`
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert pulumi.BoolPtrOutput `pulumi:"oidcVerifyCert"`
	// The ID of the Managed Private Registry. **Changing this value recreates the resource.**
	RegistryId pulumi.StringOutput `pulumi:"registryId"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
}

// NewContainerRegistryOIDC registers a new resource with the given unique name, arguments, and options.
func NewContainerRegistryOIDC(ctx *pulumi.Context,
	name string, args *ContainerRegistryOIDCArgs, opts ...pulumi.ResourceOption) (*ContainerRegistryOIDC, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OidcClientId == nil {
		return nil, errors.New("invalid value for required argument 'OidcClientId'")
	}
	if args.OidcClientSecret == nil {
		return nil, errors.New("invalid value for required argument 'OidcClientSecret'")
	}
	if args.OidcEndpoint == nil {
		return nil, errors.New("invalid value for required argument 'OidcEndpoint'")
	}
	if args.OidcName == nil {
		return nil, errors.New("invalid value for required argument 'OidcName'")
	}
	if args.OidcScope == nil {
		return nil, errors.New("invalid value for required argument 'OidcScope'")
	}
	if args.RegistryId == nil {
		return nil, errors.New("invalid value for required argument 'RegistryId'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	if args.OidcClientSecret != nil {
		args.OidcClientSecret = pulumi.ToSecret(args.OidcClientSecret).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"oidcClientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ContainerRegistryOIDC
	err := ctx.RegisterResource("ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetContainerRegistryOIDC gets an existing ContainerRegistryOIDC resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetContainerRegistryOIDC(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ContainerRegistryOIDCState, opts ...pulumi.ResourceOption) (*ContainerRegistryOIDC, error) {
	var resource ContainerRegistryOIDC
	err := ctx.ReadResource("ovh:CloudProject/containerRegistryOIDC:ContainerRegistryOIDC", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ContainerRegistryOIDC resources.
type containerRegistryOIDCState struct {
	// Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
	DeleteUsers *bool `pulumi:"deleteUsers"`
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup *string `pulumi:"oidcAdminGroup"`
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard *bool `pulumi:"oidcAutoOnboard"`
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId *string `pulumi:"oidcClientId"`
	// The secret for the Harbor client application.
	OidcClientSecret *string `pulumi:"oidcClientSecret"`
	// The URL of an OIDC-compliant server.
	OidcEndpoint *string `pulumi:"oidcEndpoint"`
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim *string `pulumi:"oidcGroupsClaim"`
	// The name of the OIDC provider.
	OidcName *string `pulumi:"oidcName"`
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope *string `pulumi:"oidcScope"`
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim *string `pulumi:"oidcUserClaim"`
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert *bool `pulumi:"oidcVerifyCert"`
	// The ID of the Managed Private Registry. **Changing this value recreates the resource.**
	RegistryId *string `pulumi:"registryId"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName *string `pulumi:"serviceName"`
}

type ContainerRegistryOIDCState struct {
	// Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
	DeleteUsers pulumi.BoolPtrInput
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup pulumi.StringPtrInput
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard pulumi.BoolPtrInput
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId pulumi.StringPtrInput
	// The secret for the Harbor client application.
	OidcClientSecret pulumi.StringPtrInput
	// The URL of an OIDC-compliant server.
	OidcEndpoint pulumi.StringPtrInput
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim pulumi.StringPtrInput
	// The name of the OIDC provider.
	OidcName pulumi.StringPtrInput
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope pulumi.StringPtrInput
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim pulumi.StringPtrInput
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert pulumi.BoolPtrInput
	// The ID of the Managed Private Registry. **Changing this value recreates the resource.**
	RegistryId pulumi.StringPtrInput
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringPtrInput
}

func (ContainerRegistryOIDCState) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryOIDCState)(nil)).Elem()
}

type containerRegistryOIDCArgs struct {
	// Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
	DeleteUsers *bool `pulumi:"deleteUsers"`
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup *string `pulumi:"oidcAdminGroup"`
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard *bool `pulumi:"oidcAutoOnboard"`
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId string `pulumi:"oidcClientId"`
	// The secret for the Harbor client application.
	OidcClientSecret string `pulumi:"oidcClientSecret"`
	// The URL of an OIDC-compliant server.
	OidcEndpoint string `pulumi:"oidcEndpoint"`
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim *string `pulumi:"oidcGroupsClaim"`
	// The name of the OIDC provider.
	OidcName string `pulumi:"oidcName"`
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope string `pulumi:"oidcScope"`
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim *string `pulumi:"oidcUserClaim"`
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert *bool `pulumi:"oidcVerifyCert"`
	// The ID of the Managed Private Registry. **Changing this value recreates the resource.**
	RegistryId string `pulumi:"registryId"`
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName string `pulumi:"serviceName"`
}

// The set of arguments for constructing a ContainerRegistryOIDC resource.
type ContainerRegistryOIDCArgs struct {
	// Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
	DeleteUsers pulumi.BoolPtrInput
	// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
	OidcAdminGroup pulumi.StringPtrInput
	// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
	OidcAutoOnboard pulumi.BoolPtrInput
	// The client ID with which Harbor is registered as client application with the OIDC provider.
	OidcClientId pulumi.StringInput
	// The secret for the Harbor client application.
	OidcClientSecret pulumi.StringInput
	// The URL of an OIDC-compliant server.
	OidcEndpoint pulumi.StringInput
	// The name of Claim in the ID token whose value is the list of group names.
	OidcGroupsClaim pulumi.StringPtrInput
	// The name of the OIDC provider.
	OidcName pulumi.StringInput
	// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
	OidcScope pulumi.StringInput
	// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
	OidcUserClaim pulumi.StringPtrInput
	// Set it to `false` if your OIDC server is hosted via self-signed certificate.
	OidcVerifyCert pulumi.BoolPtrInput
	// The ID of the Managed Private Registry. **Changing this value recreates the resource.**
	RegistryId pulumi.StringInput
	// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringInput
}

func (ContainerRegistryOIDCArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*containerRegistryOIDCArgs)(nil)).Elem()
}

type ContainerRegistryOIDCInput interface {
	pulumi.Input

	ToContainerRegistryOIDCOutput() ContainerRegistryOIDCOutput
	ToContainerRegistryOIDCOutputWithContext(ctx context.Context) ContainerRegistryOIDCOutput
}

func (*ContainerRegistryOIDC) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryOIDC)(nil)).Elem()
}

func (i *ContainerRegistryOIDC) ToContainerRegistryOIDCOutput() ContainerRegistryOIDCOutput {
	return i.ToContainerRegistryOIDCOutputWithContext(context.Background())
}

func (i *ContainerRegistryOIDC) ToContainerRegistryOIDCOutputWithContext(ctx context.Context) ContainerRegistryOIDCOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOIDCOutput)
}

func (i *ContainerRegistryOIDC) ToOutput(ctx context.Context) pulumix.Output[*ContainerRegistryOIDC] {
	return pulumix.Output[*ContainerRegistryOIDC]{
		OutputState: i.ToContainerRegistryOIDCOutputWithContext(ctx).OutputState,
	}
}

// ContainerRegistryOIDCArrayInput is an input type that accepts ContainerRegistryOIDCArray and ContainerRegistryOIDCArrayOutput values.
// You can construct a concrete instance of `ContainerRegistryOIDCArrayInput` via:
//
//	ContainerRegistryOIDCArray{ ContainerRegistryOIDCArgs{...} }
type ContainerRegistryOIDCArrayInput interface {
	pulumi.Input

	ToContainerRegistryOIDCArrayOutput() ContainerRegistryOIDCArrayOutput
	ToContainerRegistryOIDCArrayOutputWithContext(context.Context) ContainerRegistryOIDCArrayOutput
}

type ContainerRegistryOIDCArray []ContainerRegistryOIDCInput

func (ContainerRegistryOIDCArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryOIDC)(nil)).Elem()
}

func (i ContainerRegistryOIDCArray) ToContainerRegistryOIDCArrayOutput() ContainerRegistryOIDCArrayOutput {
	return i.ToContainerRegistryOIDCArrayOutputWithContext(context.Background())
}

func (i ContainerRegistryOIDCArray) ToContainerRegistryOIDCArrayOutputWithContext(ctx context.Context) ContainerRegistryOIDCArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOIDCArrayOutput)
}

func (i ContainerRegistryOIDCArray) ToOutput(ctx context.Context) pulumix.Output[[]*ContainerRegistryOIDC] {
	return pulumix.Output[[]*ContainerRegistryOIDC]{
		OutputState: i.ToContainerRegistryOIDCArrayOutputWithContext(ctx).OutputState,
	}
}

// ContainerRegistryOIDCMapInput is an input type that accepts ContainerRegistryOIDCMap and ContainerRegistryOIDCMapOutput values.
// You can construct a concrete instance of `ContainerRegistryOIDCMapInput` via:
//
//	ContainerRegistryOIDCMap{ "key": ContainerRegistryOIDCArgs{...} }
type ContainerRegistryOIDCMapInput interface {
	pulumi.Input

	ToContainerRegistryOIDCMapOutput() ContainerRegistryOIDCMapOutput
	ToContainerRegistryOIDCMapOutputWithContext(context.Context) ContainerRegistryOIDCMapOutput
}

type ContainerRegistryOIDCMap map[string]ContainerRegistryOIDCInput

func (ContainerRegistryOIDCMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryOIDC)(nil)).Elem()
}

func (i ContainerRegistryOIDCMap) ToContainerRegistryOIDCMapOutput() ContainerRegistryOIDCMapOutput {
	return i.ToContainerRegistryOIDCMapOutputWithContext(context.Background())
}

func (i ContainerRegistryOIDCMap) ToContainerRegistryOIDCMapOutputWithContext(ctx context.Context) ContainerRegistryOIDCMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerRegistryOIDCMapOutput)
}

func (i ContainerRegistryOIDCMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*ContainerRegistryOIDC] {
	return pulumix.Output[map[string]*ContainerRegistryOIDC]{
		OutputState: i.ToContainerRegistryOIDCMapOutputWithContext(ctx).OutputState,
	}
}

type ContainerRegistryOIDCOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOIDCOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerRegistryOIDC)(nil)).Elem()
}

func (o ContainerRegistryOIDCOutput) ToContainerRegistryOIDCOutput() ContainerRegistryOIDCOutput {
	return o
}

func (o ContainerRegistryOIDCOutput) ToContainerRegistryOIDCOutputWithContext(ctx context.Context) ContainerRegistryOIDCOutput {
	return o
}

func (o ContainerRegistryOIDCOutput) ToOutput(ctx context.Context) pulumix.Output[*ContainerRegistryOIDC] {
	return pulumix.Output[*ContainerRegistryOIDC]{
		OutputState: o.OutputState,
	}
}

// Delete existing users from Harbor. OIDC can't be enabled if there is at least one user already created. This parameter is only used at OIDC configuration creation. **Changing this value recreates the resource.**
func (o ContainerRegistryOIDCOutput) DeleteUsers() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.BoolPtrOutput { return v.DeleteUsers }).(pulumi.BoolPtrOutput)
}

// Specify an OIDC admin group name. All OIDC users in this group will have harbor admin privilege. Keep it blank if you do not want to.
func (o ContainerRegistryOIDCOutput) OidcAdminGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringPtrOutput { return v.OidcAdminGroup }).(pulumi.StringPtrOutput)
}

// Skip the onboarding screen, so user cannot change its username. Username is provided from ID Token.
func (o ContainerRegistryOIDCOutput) OidcAutoOnboard() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.BoolPtrOutput { return v.OidcAutoOnboard }).(pulumi.BoolPtrOutput)
}

// The client ID with which Harbor is registered as client application with the OIDC provider.
func (o ContainerRegistryOIDCOutput) OidcClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.OidcClientId }).(pulumi.StringOutput)
}

// The secret for the Harbor client application.
func (o ContainerRegistryOIDCOutput) OidcClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.OidcClientSecret }).(pulumi.StringOutput)
}

// The URL of an OIDC-compliant server.
func (o ContainerRegistryOIDCOutput) OidcEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.OidcEndpoint }).(pulumi.StringOutput)
}

// The name of Claim in the ID token whose value is the list of group names.
func (o ContainerRegistryOIDCOutput) OidcGroupsClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringPtrOutput { return v.OidcGroupsClaim }).(pulumi.StringPtrOutput)
}

// The name of the OIDC provider.
func (o ContainerRegistryOIDCOutput) OidcName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.OidcName }).(pulumi.StringOutput)
}

// The scope sent to OIDC server during authentication. It's a comma-separated string that must contain 'openid' and usually also contains 'profile' and 'email'. To obtain refresh tokens it should also contain 'offline_access'.
func (o ContainerRegistryOIDCOutput) OidcScope() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.OidcScope }).(pulumi.StringOutput)
}

// The name of the claim in the ID Token where the username is retrieved from. If not specified, it will default to 'name' (only useful when automatic Onboarding is enabled).
func (o ContainerRegistryOIDCOutput) OidcUserClaim() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringPtrOutput { return v.OidcUserClaim }).(pulumi.StringPtrOutput)
}

// Set it to `false` if your OIDC server is hosted via self-signed certificate.
func (o ContainerRegistryOIDCOutput) OidcVerifyCert() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.BoolPtrOutput { return v.OidcVerifyCert }).(pulumi.BoolPtrOutput)
}

// The ID of the Managed Private Registry. **Changing this value recreates the resource.**
func (o ContainerRegistryOIDCOutput) RegistryId() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.RegistryId }).(pulumi.StringOutput)
}

// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
func (o ContainerRegistryOIDCOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *ContainerRegistryOIDC) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

type ContainerRegistryOIDCArrayOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOIDCArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ContainerRegistryOIDC)(nil)).Elem()
}

func (o ContainerRegistryOIDCArrayOutput) ToContainerRegistryOIDCArrayOutput() ContainerRegistryOIDCArrayOutput {
	return o
}

func (o ContainerRegistryOIDCArrayOutput) ToContainerRegistryOIDCArrayOutputWithContext(ctx context.Context) ContainerRegistryOIDCArrayOutput {
	return o
}

func (o ContainerRegistryOIDCArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*ContainerRegistryOIDC] {
	return pulumix.Output[[]*ContainerRegistryOIDC]{
		OutputState: o.OutputState,
	}
}

func (o ContainerRegistryOIDCArrayOutput) Index(i pulumi.IntInput) ContainerRegistryOIDCOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ContainerRegistryOIDC {
		return vs[0].([]*ContainerRegistryOIDC)[vs[1].(int)]
	}).(ContainerRegistryOIDCOutput)
}

type ContainerRegistryOIDCMapOutput struct{ *pulumi.OutputState }

func (ContainerRegistryOIDCMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ContainerRegistryOIDC)(nil)).Elem()
}

func (o ContainerRegistryOIDCMapOutput) ToContainerRegistryOIDCMapOutput() ContainerRegistryOIDCMapOutput {
	return o
}

func (o ContainerRegistryOIDCMapOutput) ToContainerRegistryOIDCMapOutputWithContext(ctx context.Context) ContainerRegistryOIDCMapOutput {
	return o
}

func (o ContainerRegistryOIDCMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*ContainerRegistryOIDC] {
	return pulumix.Output[map[string]*ContainerRegistryOIDC]{
		OutputState: o.OutputState,
	}
}

func (o ContainerRegistryOIDCMapOutput) MapIndex(k pulumi.StringInput) ContainerRegistryOIDCOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ContainerRegistryOIDC {
		return vs[0].(map[string]*ContainerRegistryOIDC)[vs[1].(string)]
	}).(ContainerRegistryOIDCOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryOIDCInput)(nil)).Elem(), &ContainerRegistryOIDC{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryOIDCArrayInput)(nil)).Elem(), ContainerRegistryOIDCArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContainerRegistryOIDCMapInput)(nil)).Elem(), ContainerRegistryOIDCMap{})
	pulumi.RegisterOutputType(ContainerRegistryOIDCOutput{})
	pulumi.RegisterOutputType(ContainerRegistryOIDCArrayOutput{})
	pulumi.RegisterOutputType(ContainerRegistryOIDCMapOutput{})
}
