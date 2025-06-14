// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package me

import (
	"context"
	"reflect"

	"errors"
	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an OAuth2 service account.
//
// ## Example Usage
//
// An OAuth2 client for an app hosted at `my-app.com`, that uses the authorization code flow to authenticate.
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := me.NewAPIOAuth2Client(ctx, "my_oauth2_client_auth_code", &me.APIOAuth2ClientArgs{
//				Name:        pulumi.String("OAuth2 authorization code service account"),
//				Flow:        pulumi.String("AUTHORIZATION_CODE"),
//				Description: pulumi.String("An OAuth2 client using the authorization code flow for my-app.com"),
//				CallbackUrls: pulumi.StringArray{
//					pulumi.String("https://my-app.com/callback"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// An OAuth2 client for an app hosted at `my-app.com`, that uses the client credentials flow to authenticate.
//
// ```go
// package main
//
// import (
//
//	"github.com/ovh/pulumi-ovh/sdk/v2/go/ovh/me"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := me.NewAPIOAuth2Client(ctx, "my_oauth2_client_client_creds", &me.APIOAuth2ClientArgs{
//				Name:        pulumi.String("client credentials service account"),
//				Description: pulumi.String("An OAuth2 client using the client credentials flow for my app"),
//				Flow:        pulumi.String("CLIENT_CREDENTIALS"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// OAuth2 clients can be imported using their `client_id`:
//
// bash
//
// ```sh
// $ pulumi import ovh:Me/aPIOAuth2Client:APIOAuth2Client my_oauth2_client client_id
// ```
//
// Because the client_secret is only available for resources created using terraform, OAuth2 clients can also be imported using a `client_id` and a `client_secret` with a pipe separator:
//
// bash
//
// ```sh
// $ pulumi import ovh:Me/aPIOAuth2Client:APIOAuth2Client my_oauth2_client 'client_id|client_secret'
// ```
type APIOAuth2Client struct {
	pulumi.CustomResourceState

	// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
	CallbackUrls pulumi.StringArrayOutput `pulumi:"callbackUrls"`
	// Client ID of the created service account.
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// Client secret of the created service account.
	ClientSecret pulumi.StringOutput `pulumi:"clientSecret"`
	// OAuth2 client description.
	Description pulumi.StringOutput `pulumi:"description"`
	// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
	Flow pulumi.StringOutput `pulumi:"flow"`
	// URN that will allow you to associate this oauth2 client with an access policy
	Identity pulumi.StringOutput `pulumi:"identity"`
	// OAuth2 client name.
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewAPIOAuth2Client registers a new resource with the given unique name, arguments, and options.
func NewAPIOAuth2Client(ctx *pulumi.Context,
	name string, args *APIOAuth2ClientArgs, opts ...pulumi.ResourceOption) (*APIOAuth2Client, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Flow == nil {
		return nil, errors.New("invalid value for required argument 'Flow'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientSecret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource APIOAuth2Client
	err := ctx.RegisterResource("ovh:Me/aPIOAuth2Client:APIOAuth2Client", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAPIOAuth2Client gets an existing APIOAuth2Client resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAPIOAuth2Client(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *APIOAuth2ClientState, opts ...pulumi.ResourceOption) (*APIOAuth2Client, error) {
	var resource APIOAuth2Client
	err := ctx.ReadResource("ovh:Me/aPIOAuth2Client:APIOAuth2Client", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering APIOAuth2Client resources.
type apioauth2ClientState struct {
	// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
	CallbackUrls []string `pulumi:"callbackUrls"`
	// Client ID of the created service account.
	ClientId *string `pulumi:"clientId"`
	// Client secret of the created service account.
	ClientSecret *string `pulumi:"clientSecret"`
	// OAuth2 client description.
	Description *string `pulumi:"description"`
	// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
	Flow *string `pulumi:"flow"`
	// URN that will allow you to associate this oauth2 client with an access policy
	Identity *string `pulumi:"identity"`
	// OAuth2 client name.
	Name *string `pulumi:"name"`
}

type APIOAuth2ClientState struct {
	// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
	CallbackUrls pulumi.StringArrayInput
	// Client ID of the created service account.
	ClientId pulumi.StringPtrInput
	// Client secret of the created service account.
	ClientSecret pulumi.StringPtrInput
	// OAuth2 client description.
	Description pulumi.StringPtrInput
	// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
	Flow pulumi.StringPtrInput
	// URN that will allow you to associate this oauth2 client with an access policy
	Identity pulumi.StringPtrInput
	// OAuth2 client name.
	Name pulumi.StringPtrInput
}

func (APIOAuth2ClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*apioauth2ClientState)(nil)).Elem()
}

type apioauth2ClientArgs struct {
	// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
	CallbackUrls []string `pulumi:"callbackUrls"`
	// OAuth2 client description.
	Description string `pulumi:"description"`
	// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
	Flow string `pulumi:"flow"`
	// OAuth2 client name.
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a APIOAuth2Client resource.
type APIOAuth2ClientArgs struct {
	// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
	CallbackUrls pulumi.StringArrayInput
	// OAuth2 client description.
	Description pulumi.StringInput
	// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
	Flow pulumi.StringInput
	// OAuth2 client name.
	Name pulumi.StringPtrInput
}

func (APIOAuth2ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*apioauth2ClientArgs)(nil)).Elem()
}

type APIOAuth2ClientInput interface {
	pulumi.Input

	ToAPIOAuth2ClientOutput() APIOAuth2ClientOutput
	ToAPIOAuth2ClientOutputWithContext(ctx context.Context) APIOAuth2ClientOutput
}

func (*APIOAuth2Client) ElementType() reflect.Type {
	return reflect.TypeOf((**APIOAuth2Client)(nil)).Elem()
}

func (i *APIOAuth2Client) ToAPIOAuth2ClientOutput() APIOAuth2ClientOutput {
	return i.ToAPIOAuth2ClientOutputWithContext(context.Background())
}

func (i *APIOAuth2Client) ToAPIOAuth2ClientOutputWithContext(ctx context.Context) APIOAuth2ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIOAuth2ClientOutput)
}

// APIOAuth2ClientArrayInput is an input type that accepts APIOAuth2ClientArray and APIOAuth2ClientArrayOutput values.
// You can construct a concrete instance of `APIOAuth2ClientArrayInput` via:
//
//	APIOAuth2ClientArray{ APIOAuth2ClientArgs{...} }
type APIOAuth2ClientArrayInput interface {
	pulumi.Input

	ToAPIOAuth2ClientArrayOutput() APIOAuth2ClientArrayOutput
	ToAPIOAuth2ClientArrayOutputWithContext(context.Context) APIOAuth2ClientArrayOutput
}

type APIOAuth2ClientArray []APIOAuth2ClientInput

func (APIOAuth2ClientArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIOAuth2Client)(nil)).Elem()
}

func (i APIOAuth2ClientArray) ToAPIOAuth2ClientArrayOutput() APIOAuth2ClientArrayOutput {
	return i.ToAPIOAuth2ClientArrayOutputWithContext(context.Background())
}

func (i APIOAuth2ClientArray) ToAPIOAuth2ClientArrayOutputWithContext(ctx context.Context) APIOAuth2ClientArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIOAuth2ClientArrayOutput)
}

// APIOAuth2ClientMapInput is an input type that accepts APIOAuth2ClientMap and APIOAuth2ClientMapOutput values.
// You can construct a concrete instance of `APIOAuth2ClientMapInput` via:
//
//	APIOAuth2ClientMap{ "key": APIOAuth2ClientArgs{...} }
type APIOAuth2ClientMapInput interface {
	pulumi.Input

	ToAPIOAuth2ClientMapOutput() APIOAuth2ClientMapOutput
	ToAPIOAuth2ClientMapOutputWithContext(context.Context) APIOAuth2ClientMapOutput
}

type APIOAuth2ClientMap map[string]APIOAuth2ClientInput

func (APIOAuth2ClientMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIOAuth2Client)(nil)).Elem()
}

func (i APIOAuth2ClientMap) ToAPIOAuth2ClientMapOutput() APIOAuth2ClientMapOutput {
	return i.ToAPIOAuth2ClientMapOutputWithContext(context.Background())
}

func (i APIOAuth2ClientMap) ToAPIOAuth2ClientMapOutputWithContext(ctx context.Context) APIOAuth2ClientMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(APIOAuth2ClientMapOutput)
}

type APIOAuth2ClientOutput struct{ *pulumi.OutputState }

func (APIOAuth2ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**APIOAuth2Client)(nil)).Elem()
}

func (o APIOAuth2ClientOutput) ToAPIOAuth2ClientOutput() APIOAuth2ClientOutput {
	return o
}

func (o APIOAuth2ClientOutput) ToAPIOAuth2ClientOutputWithContext(ctx context.Context) APIOAuth2ClientOutput {
	return o
}

// List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
func (o APIOAuth2ClientOutput) CallbackUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringArrayOutput { return v.CallbackUrls }).(pulumi.StringArrayOutput)
}

// Client ID of the created service account.
func (o APIOAuth2ClientOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringOutput { return v.ClientId }).(pulumi.StringOutput)
}

// Client secret of the created service account.
func (o APIOAuth2ClientOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringOutput { return v.ClientSecret }).(pulumi.StringOutput)
}

// OAuth2 client description.
func (o APIOAuth2ClientOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
func (o APIOAuth2ClientOutput) Flow() pulumi.StringOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringOutput { return v.Flow }).(pulumi.StringOutput)
}

// URN that will allow you to associate this oauth2 client with an access policy
func (o APIOAuth2ClientOutput) Identity() pulumi.StringOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringOutput { return v.Identity }).(pulumi.StringOutput)
}

// OAuth2 client name.
func (o APIOAuth2ClientOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *APIOAuth2Client) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type APIOAuth2ClientArrayOutput struct{ *pulumi.OutputState }

func (APIOAuth2ClientArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*APIOAuth2Client)(nil)).Elem()
}

func (o APIOAuth2ClientArrayOutput) ToAPIOAuth2ClientArrayOutput() APIOAuth2ClientArrayOutput {
	return o
}

func (o APIOAuth2ClientArrayOutput) ToAPIOAuth2ClientArrayOutputWithContext(ctx context.Context) APIOAuth2ClientArrayOutput {
	return o
}

func (o APIOAuth2ClientArrayOutput) Index(i pulumi.IntInput) APIOAuth2ClientOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *APIOAuth2Client {
		return vs[0].([]*APIOAuth2Client)[vs[1].(int)]
	}).(APIOAuth2ClientOutput)
}

type APIOAuth2ClientMapOutput struct{ *pulumi.OutputState }

func (APIOAuth2ClientMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*APIOAuth2Client)(nil)).Elem()
}

func (o APIOAuth2ClientMapOutput) ToAPIOAuth2ClientMapOutput() APIOAuth2ClientMapOutput {
	return o
}

func (o APIOAuth2ClientMapOutput) ToAPIOAuth2ClientMapOutputWithContext(ctx context.Context) APIOAuth2ClientMapOutput {
	return o
}

func (o APIOAuth2ClientMapOutput) MapIndex(k pulumi.StringInput) APIOAuth2ClientOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *APIOAuth2Client {
		return vs[0].(map[string]*APIOAuth2Client)[vs[1].(string)]
	}).(APIOAuth2ClientOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*APIOAuth2ClientInput)(nil)).Elem(), &APIOAuth2Client{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIOAuth2ClientArrayInput)(nil)).Elem(), APIOAuth2ClientArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*APIOAuth2ClientMapInput)(nil)).Elem(), APIOAuth2ClientMap{})
	pulumi.RegisterOutputType(APIOAuth2ClientOutput{})
	pulumi.RegisterOutputType(APIOAuth2ClientArrayOutput{})
	pulumi.RegisterOutputType(APIOAuth2ClientMapOutput{})
}
