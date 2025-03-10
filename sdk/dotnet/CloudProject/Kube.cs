// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    [OvhResourceType("ovh:CloudProject/kube:Kube")]
    public partial class Kube : global::Pulumi.CustomResource
    {
        [Output("controlPlaneIsUpToDate")]
        public Output<bool> ControlPlaneIsUpToDate { get; private set; } = null!;

        [Output("customizationApiservers")]
        public Output<ImmutableArray<Outputs.KubeCustomizationApiserver>> CustomizationApiservers { get; private set; } = null!;

        [Output("customizationKubeProxy")]
        public Output<Outputs.KubeCustomizationKubeProxy?> CustomizationKubeProxy { get; private set; } = null!;

        [Output("customizations")]
        public Output<ImmutableArray<Outputs.KubeCustomization>> Customizations { get; private set; } = null!;

        [Output("isUpToDate")]
        public Output<bool> IsUpToDate { get; private set; } = null!;

        [Output("kubeProxyMode")]
        public Output<string> KubeProxyMode { get; private set; } = null!;

        [Output("kubeconfig")]
        public Output<string> Kubeconfig { get; private set; } = null!;

        /// <summary>
        /// The kubeconfig configuration file of the Kubernetes cluster
        /// </summary>
        [Output("kubeconfigAttributes")]
        public Output<ImmutableArray<Outputs.KubeKubeconfigAttribute>> KubeconfigAttributes { get; private set; } = null!;

        [Output("loadBalancersSubnetId")]
        public Output<string?> LoadBalancersSubnetId { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("nextUpgradeVersions")]
        public Output<ImmutableArray<string>> NextUpgradeVersions { get; private set; } = null!;

        [Output("nodesSubnetId")]
        public Output<string> NodesSubnetId { get; private set; } = null!;

        [Output("nodesUrl")]
        public Output<string> NodesUrl { get; private set; } = null!;

        [Output("privateNetworkConfiguration")]
        public Output<Outputs.KubePrivateNetworkConfiguration?> PrivateNetworkConfiguration { get; private set; } = null!;

        [Output("privateNetworkId")]
        public Output<string?> PrivateNetworkId { get; private set; } = null!;

        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("updatePolicy")]
        public Output<string> UpdatePolicy { get; private set; } = null!;

        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Kube resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Kube(string name, KubeArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/kube:Kube", name, args ?? new KubeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Kube(string name, Input<string> id, KubeState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/kube:Kube", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
                AdditionalSecretOutputs =
                {
                    "kubeconfig",
                    "kubeconfigAttributes",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Kube resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Kube Get(string name, Input<string> id, KubeState? state = null, CustomResourceOptions? options = null)
        {
            return new Kube(name, id, state, options);
        }
    }

    public sealed class KubeArgs : global::Pulumi.ResourceArgs
    {
        [Input("customizationApiservers")]
        private InputList<Inputs.KubeCustomizationApiserverArgs>? _customizationApiservers;
        public InputList<Inputs.KubeCustomizationApiserverArgs> CustomizationApiservers
        {
            get => _customizationApiservers ?? (_customizationApiservers = new InputList<Inputs.KubeCustomizationApiserverArgs>());
            set => _customizationApiservers = value;
        }

        [Input("customizationKubeProxy")]
        public Input<Inputs.KubeCustomizationKubeProxyArgs>? CustomizationKubeProxy { get; set; }

        [Input("customizations")]
        private InputList<Inputs.KubeCustomizationArgs>? _customizations;
        [Obsolete(@"Use customization_apiserver instead")]
        public InputList<Inputs.KubeCustomizationArgs> Customizations
        {
            get => _customizations ?? (_customizations = new InputList<Inputs.KubeCustomizationArgs>());
            set => _customizations = value;
        }

        [Input("kubeProxyMode")]
        public Input<string>? KubeProxyMode { get; set; }

        [Input("loadBalancersSubnetId")]
        public Input<string>? LoadBalancersSubnetId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nodesSubnetId")]
        public Input<string>? NodesSubnetId { get; set; }

        [Input("privateNetworkConfiguration")]
        public Input<Inputs.KubePrivateNetworkConfigurationArgs>? PrivateNetworkConfiguration { get; set; }

        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("updatePolicy")]
        public Input<string>? UpdatePolicy { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubeArgs()
        {
        }
        public static new KubeArgs Empty => new KubeArgs();
    }

    public sealed class KubeState : global::Pulumi.ResourceArgs
    {
        [Input("controlPlaneIsUpToDate")]
        public Input<bool>? ControlPlaneIsUpToDate { get; set; }

        [Input("customizationApiservers")]
        private InputList<Inputs.KubeCustomizationApiserverGetArgs>? _customizationApiservers;
        public InputList<Inputs.KubeCustomizationApiserverGetArgs> CustomizationApiservers
        {
            get => _customizationApiservers ?? (_customizationApiservers = new InputList<Inputs.KubeCustomizationApiserverGetArgs>());
            set => _customizationApiservers = value;
        }

        [Input("customizationKubeProxy")]
        public Input<Inputs.KubeCustomizationKubeProxyGetArgs>? CustomizationKubeProxy { get; set; }

        [Input("customizations")]
        private InputList<Inputs.KubeCustomizationGetArgs>? _customizations;
        [Obsolete(@"Use customization_apiserver instead")]
        public InputList<Inputs.KubeCustomizationGetArgs> Customizations
        {
            get => _customizations ?? (_customizations = new InputList<Inputs.KubeCustomizationGetArgs>());
            set => _customizations = value;
        }

        [Input("isUpToDate")]
        public Input<bool>? IsUpToDate { get; set; }

        [Input("kubeProxyMode")]
        public Input<string>? KubeProxyMode { get; set; }

        [Input("kubeconfig")]
        private Input<string>? _kubeconfig;
        public Input<string>? Kubeconfig
        {
            get => _kubeconfig;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _kubeconfig = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        [Input("kubeconfigAttributes")]
        private InputList<Inputs.KubeKubeconfigAttributeGetArgs>? _kubeconfigAttributes;

        /// <summary>
        /// The kubeconfig configuration file of the Kubernetes cluster
        /// </summary>
        public InputList<Inputs.KubeKubeconfigAttributeGetArgs> KubeconfigAttributes
        {
            get => _kubeconfigAttributes ?? (_kubeconfigAttributes = new InputList<Inputs.KubeKubeconfigAttributeGetArgs>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableArray.Create<Inputs.KubeKubeconfigAttributeGetArgs>());
                _kubeconfigAttributes = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("loadBalancersSubnetId")]
        public Input<string>? LoadBalancersSubnetId { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("nextUpgradeVersions")]
        private InputList<string>? _nextUpgradeVersions;
        public InputList<string> NextUpgradeVersions
        {
            get => _nextUpgradeVersions ?? (_nextUpgradeVersions = new InputList<string>());
            set => _nextUpgradeVersions = value;
        }

        [Input("nodesSubnetId")]
        public Input<string>? NodesSubnetId { get; set; }

        [Input("nodesUrl")]
        public Input<string>? NodesUrl { get; set; }

        [Input("privateNetworkConfiguration")]
        public Input<Inputs.KubePrivateNetworkConfigurationGetArgs>? PrivateNetworkConfiguration { get; set; }

        [Input("privateNetworkId")]
        public Input<string>? PrivateNetworkId { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("updatePolicy")]
        public Input<string>? UpdatePolicy { get; set; }

        [Input("url")]
        public Input<string>? Url { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public KubeState()
        {
        }
        public static new KubeState Empty => new KubeState();
    }
}
