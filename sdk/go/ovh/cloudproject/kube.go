// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package cloudproject

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/scraly/pulumi-ovh/sdk/go/ovh/internal"
)

// ## Import
//
// OVHcloud Managed Kubernetes Service clusters can be imported using the `service_name` and the `id` of the cluster, separated by "/" E.g., bash <break><break>```sh<break> $ pulumi import ovh:CloudProject/kube:Kube my_kube_cluster service_name/kube_id <break>```<break><break>
type Kube struct {
	pulumi.CustomResourceState

	// True if control-plane is up-to-date.
	ControlPlaneIsUpToDate pulumi.BoolOutput `pulumi:"controlPlaneIsUpToDate"`
	// Kubernetes API server customization
	CustomizationApiservers KubeCustomizationApiserverArrayOutput `pulumi:"customizationApiservers"`
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy KubeCustomizationKubeProxyPtrOutput `pulumi:"customizationKubeProxy"`
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations KubeCustomizationArrayOutput `pulumi:"customizations"`
	// True if all nodes and control-plane are up-to-date.
	IsUpToDate pulumi.BoolOutput `pulumi:"isUpToDate"`
	// Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
	KubeProxyMode pulumi.StringOutput `pulumi:"kubeProxyMode"`
	// The kubeconfig file. Use this file to connect to your kubernetes cluster.
	Kubeconfig pulumi.StringOutput `pulumi:"kubeconfig"`
	// The kubeconfig file attributes.
	KubeconfigAttributes KubeKubeconfigAttributeArrayOutput `pulumi:"kubeconfigAttributes"`
	// The name of the kubernetes cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Kubernetes versions available for upgrade.
	NextUpgradeVersions pulumi.StringArrayOutput `pulumi:"nextUpgradeVersions"`
	// Cluster nodes URL.
	NodesUrl pulumi.StringOutput `pulumi:"nodesUrl"`
	// The private network configuration
	PrivateNetworkConfiguration KubePrivateNetworkConfigurationPtrOutput `pulumi:"privateNetworkConfiguration"`
	// OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
	//
	// > __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
	PrivateNetworkId pulumi.StringPtrOutput `pulumi:"privateNetworkId"`
	// a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions. **Changing this value recreates the resource.**
	Region pulumi.StringOutput `pulumi:"region"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringOutput `pulumi:"serviceName"`
	// Cluster status. Should be normally set to 'READY'.
	Status pulumi.StringOutput `pulumi:"status"`
	// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
	UpdatePolicy pulumi.StringOutput `pulumi:"updatePolicy"`
	// Management URL of your cluster.
	Url pulumi.StringOutput `pulumi:"url"`
	// kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewKube registers a new resource with the given unique name, arguments, and options.
func NewKube(ctx *pulumi.Context,
	name string, args *KubeArgs, opts ...pulumi.ResourceOption) (*Kube, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Region == nil {
		return nil, errors.New("invalid value for required argument 'Region'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"kubeconfig",
		"kubeconfigAttributes",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Kube
	err := ctx.RegisterResource("ovh:CloudProject/kube:Kube", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetKube gets an existing Kube resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetKube(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *KubeState, opts ...pulumi.ResourceOption) (*Kube, error) {
	var resource Kube
	err := ctx.ReadResource("ovh:CloudProject/kube:Kube", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Kube resources.
type kubeState struct {
	// True if control-plane is up-to-date.
	ControlPlaneIsUpToDate *bool `pulumi:"controlPlaneIsUpToDate"`
	// Kubernetes API server customization
	CustomizationApiservers []KubeCustomizationApiserver `pulumi:"customizationApiservers"`
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy *KubeCustomizationKubeProxy `pulumi:"customizationKubeProxy"`
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations []KubeCustomization `pulumi:"customizations"`
	// True if all nodes and control-plane are up-to-date.
	IsUpToDate *bool `pulumi:"isUpToDate"`
	// Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
	KubeProxyMode *string `pulumi:"kubeProxyMode"`
	// The kubeconfig file. Use this file to connect to your kubernetes cluster.
	Kubeconfig *string `pulumi:"kubeconfig"`
	// The kubeconfig file attributes.
	KubeconfigAttributes []KubeKubeconfigAttribute `pulumi:"kubeconfigAttributes"`
	// The name of the kubernetes cluster.
	Name *string `pulumi:"name"`
	// Kubernetes versions available for upgrade.
	NextUpgradeVersions []string `pulumi:"nextUpgradeVersions"`
	// Cluster nodes URL.
	NodesUrl *string `pulumi:"nodesUrl"`
	// The private network configuration
	PrivateNetworkConfiguration *KubePrivateNetworkConfiguration `pulumi:"privateNetworkConfiguration"`
	// OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
	//
	// > __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions. **Changing this value recreates the resource.**
	Region *string `pulumi:"region"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName *string `pulumi:"serviceName"`
	// Cluster status. Should be normally set to 'READY'.
	Status *string `pulumi:"status"`
	// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
	UpdatePolicy *string `pulumi:"updatePolicy"`
	// Management URL of your cluster.
	Url *string `pulumi:"url"`
	// kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
	Version *string `pulumi:"version"`
}

type KubeState struct {
	// True if control-plane is up-to-date.
	ControlPlaneIsUpToDate pulumi.BoolPtrInput
	// Kubernetes API server customization
	CustomizationApiservers KubeCustomizationApiserverArrayInput
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy KubeCustomizationKubeProxyPtrInput
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations KubeCustomizationArrayInput
	// True if all nodes and control-plane are up-to-date.
	IsUpToDate pulumi.BoolPtrInput
	// Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
	KubeProxyMode pulumi.StringPtrInput
	// The kubeconfig file. Use this file to connect to your kubernetes cluster.
	Kubeconfig pulumi.StringPtrInput
	// The kubeconfig file attributes.
	KubeconfigAttributes KubeKubeconfigAttributeArrayInput
	// The name of the kubernetes cluster.
	Name pulumi.StringPtrInput
	// Kubernetes versions available for upgrade.
	NextUpgradeVersions pulumi.StringArrayInput
	// Cluster nodes URL.
	NodesUrl pulumi.StringPtrInput
	// The private network configuration
	PrivateNetworkConfiguration KubePrivateNetworkConfigurationPtrInput
	// OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
	//
	// > __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
	PrivateNetworkId pulumi.StringPtrInput
	// a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions. **Changing this value recreates the resource.**
	Region pulumi.StringPtrInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringPtrInput
	// Cluster status. Should be normally set to 'READY'.
	Status pulumi.StringPtrInput
	// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
	UpdatePolicy pulumi.StringPtrInput
	// Management URL of your cluster.
	Url pulumi.StringPtrInput
	// kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
	Version pulumi.StringPtrInput
}

func (KubeState) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeState)(nil)).Elem()
}

type kubeArgs struct {
	// Kubernetes API server customization
	CustomizationApiservers []KubeCustomizationApiserver `pulumi:"customizationApiservers"`
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy *KubeCustomizationKubeProxy `pulumi:"customizationKubeProxy"`
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations []KubeCustomization `pulumi:"customizations"`
	// Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
	KubeProxyMode *string `pulumi:"kubeProxyMode"`
	// The name of the kubernetes cluster.
	Name *string `pulumi:"name"`
	// The private network configuration
	PrivateNetworkConfiguration *KubePrivateNetworkConfiguration `pulumi:"privateNetworkConfiguration"`
	// OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
	//
	// > __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
	PrivateNetworkId *string `pulumi:"privateNetworkId"`
	// a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions. **Changing this value recreates the resource.**
	Region string `pulumi:"region"`
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName string `pulumi:"serviceName"`
	// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
	UpdatePolicy *string `pulumi:"updatePolicy"`
	// kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a Kube resource.
type KubeArgs struct {
	// Kubernetes API server customization
	CustomizationApiservers KubeCustomizationApiserverArrayInput
	// Kubernetes kube-proxy customization
	CustomizationKubeProxy KubeCustomizationKubeProxyPtrInput
	// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
	//
	// Deprecated: Use customization_apiserver instead
	Customizations KubeCustomizationArrayInput
	// Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
	KubeProxyMode pulumi.StringPtrInput
	// The name of the kubernetes cluster.
	Name pulumi.StringPtrInput
	// The private network configuration
	PrivateNetworkConfiguration KubePrivateNetworkConfigurationPtrInput
	// OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
	//
	// > __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
	PrivateNetworkId pulumi.StringPtrInput
	// a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions. **Changing this value recreates the resource.**
	Region pulumi.StringInput
	// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
	ServiceName pulumi.StringInput
	// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
	UpdatePolicy pulumi.StringPtrInput
	// kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
	Version pulumi.StringPtrInput
}

func (KubeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*kubeArgs)(nil)).Elem()
}

type KubeInput interface {
	pulumi.Input

	ToKubeOutput() KubeOutput
	ToKubeOutputWithContext(ctx context.Context) KubeOutput
}

func (*Kube) ElementType() reflect.Type {
	return reflect.TypeOf((**Kube)(nil)).Elem()
}

func (i *Kube) ToKubeOutput() KubeOutput {
	return i.ToKubeOutputWithContext(context.Background())
}

func (i *Kube) ToKubeOutputWithContext(ctx context.Context) KubeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeOutput)
}

// KubeArrayInput is an input type that accepts KubeArray and KubeArrayOutput values.
// You can construct a concrete instance of `KubeArrayInput` via:
//
//	KubeArray{ KubeArgs{...} }
type KubeArrayInput interface {
	pulumi.Input

	ToKubeArrayOutput() KubeArrayOutput
	ToKubeArrayOutputWithContext(context.Context) KubeArrayOutput
}

type KubeArray []KubeInput

func (KubeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kube)(nil)).Elem()
}

func (i KubeArray) ToKubeArrayOutput() KubeArrayOutput {
	return i.ToKubeArrayOutputWithContext(context.Background())
}

func (i KubeArray) ToKubeArrayOutputWithContext(ctx context.Context) KubeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeArrayOutput)
}

// KubeMapInput is an input type that accepts KubeMap and KubeMapOutput values.
// You can construct a concrete instance of `KubeMapInput` via:
//
//	KubeMap{ "key": KubeArgs{...} }
type KubeMapInput interface {
	pulumi.Input

	ToKubeMapOutput() KubeMapOutput
	ToKubeMapOutputWithContext(context.Context) KubeMapOutput
}

type KubeMap map[string]KubeInput

func (KubeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kube)(nil)).Elem()
}

func (i KubeMap) ToKubeMapOutput() KubeMapOutput {
	return i.ToKubeMapOutputWithContext(context.Background())
}

func (i KubeMap) ToKubeMapOutputWithContext(ctx context.Context) KubeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeMapOutput)
}

type KubeOutput struct{ *pulumi.OutputState }

func (KubeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Kube)(nil)).Elem()
}

func (o KubeOutput) ToKubeOutput() KubeOutput {
	return o
}

func (o KubeOutput) ToKubeOutputWithContext(ctx context.Context) KubeOutput {
	return o
}

// True if control-plane is up-to-date.
func (o KubeOutput) ControlPlaneIsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v *Kube) pulumi.BoolOutput { return v.ControlPlaneIsUpToDate }).(pulumi.BoolOutput)
}

// Kubernetes API server customization
func (o KubeOutput) CustomizationApiservers() KubeCustomizationApiserverArrayOutput {
	return o.ApplyT(func(v *Kube) KubeCustomizationApiserverArrayOutput { return v.CustomizationApiservers }).(KubeCustomizationApiserverArrayOutput)
}

// Kubernetes kube-proxy customization
func (o KubeOutput) CustomizationKubeProxy() KubeCustomizationKubeProxyPtrOutput {
	return o.ApplyT(func(v *Kube) KubeCustomizationKubeProxyPtrOutput { return v.CustomizationKubeProxy }).(KubeCustomizationKubeProxyPtrOutput)
}

// **Deprecated** (Optional) Use `customizationApiserver` and `customizationKubeProxy` instead. Kubernetes cluster customization
//
// Deprecated: Use customization_apiserver instead
func (o KubeOutput) Customizations() KubeCustomizationArrayOutput {
	return o.ApplyT(func(v *Kube) KubeCustomizationArrayOutput { return v.Customizations }).(KubeCustomizationArrayOutput)
}

// True if all nodes and control-plane are up-to-date.
func (o KubeOutput) IsUpToDate() pulumi.BoolOutput {
	return o.ApplyT(func(v *Kube) pulumi.BoolOutput { return v.IsUpToDate }).(pulumi.BoolOutput)
}

// Selected mode for kube-proxy. **Changing this value recreates the resource, including ETCD user data.** Defaults to `iptables`.
func (o KubeOutput) KubeProxyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.KubeProxyMode }).(pulumi.StringOutput)
}

// The kubeconfig file. Use this file to connect to your kubernetes cluster.
func (o KubeOutput) Kubeconfig() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.Kubeconfig }).(pulumi.StringOutput)
}

// The kubeconfig file attributes.
func (o KubeOutput) KubeconfigAttributes() KubeKubeconfigAttributeArrayOutput {
	return o.ApplyT(func(v *Kube) KubeKubeconfigAttributeArrayOutput { return v.KubeconfigAttributes }).(KubeKubeconfigAttributeArrayOutput)
}

// The name of the kubernetes cluster.
func (o KubeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Kubernetes versions available for upgrade.
func (o KubeOutput) NextUpgradeVersions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringArrayOutput { return v.NextUpgradeVersions }).(pulumi.StringArrayOutput)
}

// Cluster nodes URL.
func (o KubeOutput) NodesUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.NodesUrl }).(pulumi.StringOutput)
}

// The private network configuration
func (o KubeOutput) PrivateNetworkConfiguration() KubePrivateNetworkConfigurationPtrOutput {
	return o.ApplyT(func(v *Kube) KubePrivateNetworkConfigurationPtrOutput { return v.PrivateNetworkConfiguration }).(KubePrivateNetworkConfigurationPtrOutput)
}

// OpenStack private network (or vRack) ID to use. **Changing this value recreates the resource, including ETCD user data.** Defaults - not use private network.
//
// > __WARNING__ Updating the private network ID resets the cluster so that all user data is deleted.
func (o KubeOutput) PrivateNetworkId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringPtrOutput { return v.PrivateNetworkId }).(pulumi.StringPtrOutput)
}

// a valid OVHcloud public cloud region ID in which the kubernetes cluster will be available. Ex.: "GRA1". Defaults to all public cloud regions. **Changing this value recreates the resource.**
func (o KubeOutput) Region() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.Region }).(pulumi.StringOutput)
}

// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
func (o KubeOutput) ServiceName() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.ServiceName }).(pulumi.StringOutput)
}

// Cluster status. Should be normally set to 'READY'.
func (o KubeOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Cluster update policy. Choose between [ALWAYS_UPDATE, MINIMAL_DOWNTIME, NEVER_UPDATE].
func (o KubeOutput) UpdatePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.UpdatePolicy }).(pulumi.StringOutput)
}

// Management URL of your cluster.
func (o KubeOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// kubernetes version to use. Changing this value updates the resource. Defaults to the latest available.
func (o KubeOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Kube) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type KubeArrayOutput struct{ *pulumi.OutputState }

func (KubeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Kube)(nil)).Elem()
}

func (o KubeArrayOutput) ToKubeArrayOutput() KubeArrayOutput {
	return o
}

func (o KubeArrayOutput) ToKubeArrayOutputWithContext(ctx context.Context) KubeArrayOutput {
	return o
}

func (o KubeArrayOutput) Index(i pulumi.IntInput) KubeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Kube {
		return vs[0].([]*Kube)[vs[1].(int)]
	}).(KubeOutput)
}

type KubeMapOutput struct{ *pulumi.OutputState }

func (KubeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Kube)(nil)).Elem()
}

func (o KubeMapOutput) ToKubeMapOutput() KubeMapOutput {
	return o
}

func (o KubeMapOutput) ToKubeMapOutputWithContext(ctx context.Context) KubeMapOutput {
	return o
}

func (o KubeMapOutput) MapIndex(k pulumi.StringInput) KubeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Kube {
		return vs[0].(map[string]*Kube)[vs[1].(string)]
	}).(KubeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KubeInput)(nil)).Elem(), &Kube{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeArrayInput)(nil)).Elem(), KubeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KubeMapInput)(nil)).Elem(), KubeMap{})
	pulumi.RegisterOutputType(KubeOutput{})
	pulumi.RegisterOutputType(KubeArrayOutput{})
	pulumi.RegisterOutputType(KubeMapOutput{})
}
