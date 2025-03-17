// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ovh

import (
	"context"
	"reflect"

	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the ovh package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// The OVH API Access Token
	AccessToken pulumi.StringPtrOutput `pulumi:"accessToken"`
	// The OVH API Application Key
	ApplicationKey pulumi.StringPtrOutput `pulumi:"applicationKey"`
	// The OVH API Application Secret
	ApplicationSecret pulumi.StringPtrOutput `pulumi:"applicationSecret"`
	// OAuth 2.0 application's ID
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// OAuth 2.0 application's secret
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The OVH API Consumer Key
	ConsumerKey pulumi.StringPtrOutput `pulumi:"consumerKey"`
	// The OVH API endpoint to target (ex: "ovh-eu")
	Endpoint pulumi.StringPtrOutput `pulumi:"endpoint"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.ApplicationKey == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OVH_APPLICATION_KEY"); d != nil {
			args.ApplicationKey = pulumi.StringPtr(d.(string))
		}
	}
	if args.ApplicationSecret == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OVH_APPLICATION_SECRET"); d != nil {
			args.ApplicationSecret = pulumi.StringPtr(d.(string))
		}
	}
	if args.ConsumerKey == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OVH_CONSUMER_KEY"); d != nil {
			args.ConsumerKey = pulumi.StringPtr(d.(string))
		}
	}
	if args.Endpoint == nil {
		if d := internal.GetEnvOrDefault(nil, nil, "OVH_ENDPOINT"); d != nil {
			args.Endpoint = pulumi.StringPtr(d.(string))
		}
	}
	if args.ApplicationSecret != nil {
		args.ApplicationSecret = pulumi.ToSecret(args.ApplicationSecret).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"applicationSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:ovh", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// The OVH API Access Token
	AccessToken *string `pulumi:"accessToken"`
	// The OVH API Application Key
	ApplicationKey *string `pulumi:"applicationKey"`
	// The OVH API Application Secret
	ApplicationSecret *string `pulumi:"applicationSecret"`
	// OAuth 2.0 application's ID
	ClientId *string `pulumi:"clientId"`
	// OAuth 2.0 application's secret
	ClientSecret *string `pulumi:"clientSecret"`
	// The OVH API Consumer Key
	ConsumerKey *string `pulumi:"consumerKey"`
	// The OVH API endpoint to target (ex: "ovh-eu")
	Endpoint *string `pulumi:"endpoint"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// The OVH API Access Token
	AccessToken pulumi.StringPtrInput
	// The OVH API Application Key
	ApplicationKey pulumi.StringPtrInput
	// The OVH API Application Secret
	ApplicationSecret pulumi.StringPtrInput
	// OAuth 2.0 application's ID
	ClientId pulumi.StringPtrInput
	// OAuth 2.0 application's secret
	ClientSecret pulumi.StringPtrInput
	// The OVH API Consumer Key
	ConsumerKey pulumi.StringPtrInput
	// The OVH API endpoint to target (ex: "ovh-eu")
	Endpoint pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// The OVH API Access Token
func (o ProviderOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.AccessToken }).(pulumi.StringPtrOutput)
}

// The OVH API Application Key
func (o ProviderOutput) ApplicationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApplicationKey }).(pulumi.StringPtrOutput)
}

// The OVH API Application Secret
func (o ProviderOutput) ApplicationSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ApplicationSecret }).(pulumi.StringPtrOutput)
}

// OAuth 2.0 application's ID
func (o ProviderOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

// OAuth 2.0 application's secret
func (o ProviderOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

// The OVH API Consumer Key
func (o ProviderOutput) ConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.ConsumerKey }).(pulumi.StringPtrOutput)
}

// The OVH API endpoint to target (ex: "ovh-eu")
func (o ProviderOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.Endpoint }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
}
