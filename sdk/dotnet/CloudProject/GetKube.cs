// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetKube
    {
        /// <summary>
        /// Use this data source to get a OVHcloud Managed Kubernetes Service cluster.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["version"] = myKubeCluster.Apply(getKubeResult =&gt; getKubeResult.Version),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetKubeResult> InvokeAsync(GetKubeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubeResult>("ovh:CloudProject/getKube:getKube", args ?? new GetKubeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a OVHcloud Managed Kubernetes Service cluster.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["version"] = myKubeCluster.Apply(getKubeResult =&gt; getKubeResult.Version),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetKubeResult> Invoke(GetKubeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeResult>("ovh:CloudProject/getKube:getKube", args ?? new GetKubeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a OVHcloud Managed Kubernetes Service cluster.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Ovh = Pulumi.Ovh;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var myKubeCluster = Ovh.CloudProject.GetKube.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxx",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["version"] = myKubeCluster.Apply(getKubeResult =&gt; getKubeResult.Version),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetKubeResult> Invoke(GetKubeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeResult>("ovh:CloudProject/getKube:getKube", args ?? new GetKubeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubeArgs : global::Pulumi.InvokeArgs
    {
        [Input("customizationApiservers")]
        private List<Inputs.GetKubeCustomizationApiserverArgs>? _customizationApiservers;

        /// <summary>
        /// Kubernetes API server customization
        /// </summary>
        public List<Inputs.GetKubeCustomizationApiserverArgs> CustomizationApiservers
        {
            get => _customizationApiservers ?? (_customizationApiservers = new List<Inputs.GetKubeCustomizationApiserverArgs>());
            set => _customizationApiservers = value;
        }

        /// <summary>
        /// Kubernetes kube-proxy customization
        /// </summary>
        [Input("customizationKubeProxy")]
        public Inputs.GetKubeCustomizationKubeProxyArgs? CustomizationKubeProxy { get; set; }

        [Input("customizations")]
        private List<Inputs.GetKubeCustomizationArgs>? _customizations;

        /// <summary>
        /// **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
        /// </summary>
        [Obsolete(@"Use customization_apiserver instead")]
        public List<Inputs.GetKubeCustomizationArgs> Customizations
        {
            get => _customizations ?? (_customizations = new List<Inputs.GetKubeCustomizationArgs>());
            set => _customizations = value;
        }

        /// <summary>
        /// The id of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public string KubeId { get; set; } = null!;

        /// <summary>
        /// Selected mode for kube-proxy.
        /// </summary>
        [Input("kubeProxyMode")]
        public string? KubeProxyMode { get; set; }

        /// <summary>
        /// The name of the managed kubernetes cluster.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The OVHcloud public cloud region ID of the managed kubernetes cluster.
        /// </summary>
        [Input("region")]
        public string? Region { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        /// <summary>
        /// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
        /// </summary>
        [Input("updatePolicy")]
        public string? UpdatePolicy { get; set; }

        /// <summary>
        /// Kubernetes version of the managed kubernetes cluster.
        /// </summary>
        [Input("version")]
        public string? Version { get; set; }

        public GetKubeArgs()
        {
        }
        public static new GetKubeArgs Empty => new GetKubeArgs();
    }

    public sealed class GetKubeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("customizationApiservers")]
        private InputList<Inputs.GetKubeCustomizationApiserverInputArgs>? _customizationApiservers;

        /// <summary>
        /// Kubernetes API server customization
        /// </summary>
        public InputList<Inputs.GetKubeCustomizationApiserverInputArgs> CustomizationApiservers
        {
            get => _customizationApiservers ?? (_customizationApiservers = new InputList<Inputs.GetKubeCustomizationApiserverInputArgs>());
            set => _customizationApiservers = value;
        }

        /// <summary>
        /// Kubernetes kube-proxy customization
        /// </summary>
        [Input("customizationKubeProxy")]
        public Input<Inputs.GetKubeCustomizationKubeProxyInputArgs>? CustomizationKubeProxy { get; set; }

        [Input("customizations")]
        private InputList<Inputs.GetKubeCustomizationInputArgs>? _customizations;

        /// <summary>
        /// **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
        /// </summary>
        [Obsolete(@"Use customization_apiserver instead")]
        public InputList<Inputs.GetKubeCustomizationInputArgs> Customizations
        {
            get => _customizations ?? (_customizations = new InputList<Inputs.GetKubeCustomizationInputArgs>());
            set => _customizations = value;
        }

        /// <summary>
        /// The id of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public Input<string> KubeId { get; set; } = null!;

        /// <summary>
        /// Selected mode for kube-proxy.
        /// </summary>
        [Input("kubeProxyMode")]
        public Input<string>? KubeProxyMode { get; set; }

        /// <summary>
        /// The name of the managed kubernetes cluster.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The OVHcloud public cloud region ID of the managed kubernetes cluster.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
        /// </summary>
        [Input("updatePolicy")]
        public Input<string>? UpdatePolicy { get; set; }

        /// <summary>
        /// Kubernetes version of the managed kubernetes cluster.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public GetKubeInvokeArgs()
        {
        }
        public static new GetKubeInvokeArgs Empty => new GetKubeInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubeResult
    {
        /// <summary>
        /// True if control-plane is up-to-date.
        /// </summary>
        public readonly bool ControlPlaneIsUpToDate;
        /// <summary>
        /// Kubernetes API server customization
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubeCustomizationApiserverResult> CustomizationApiservers;
        /// <summary>
        /// Kubernetes kube-proxy customization
        /// </summary>
        public readonly Outputs.GetKubeCustomizationKubeProxyResult? CustomizationKubeProxy;
        /// <summary>
        /// **Deprecated** (Optional) Use `customization_apiserver` and `customization_kube_proxy` instead. Kubernetes cluster customization
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubeCustomizationResult> Customizations;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// True if all nodes and control-plane are up-to-date.
        /// </summary>
        public readonly bool IsUpToDate;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string KubeId;
        /// <summary>
        /// Selected mode for kube-proxy.
        /// </summary>
        public readonly string? KubeProxyMode;
        /// <summary>
        /// Openstack private network (or vRack) ID to use for load balancers.
        /// </summary>
        public readonly string LoadBalancersSubnetId;
        /// <summary>
        /// The name of the managed kubernetes cluster.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Kubernetes versions available for upgrade.
        /// </summary>
        public readonly ImmutableArray<string> NextUpgradeVersions;
        /// <summary>
        /// Openstack private network (or vRack) ID to use for nodes.
        /// </summary>
        public readonly string NodesSubnetId;
        /// <summary>
        /// Cluster nodes URL.
        /// </summary>
        public readonly string NodesUrl;
        /// <summary>
        /// OpenStack private network (or vrack) ID to use.
        /// </summary>
        public readonly string PrivateNetworkId;
        /// <summary>
        /// The OVHcloud public cloud region ID of the managed kubernetes cluster.
        /// </summary>
        public readonly string? Region;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Cluster status. Should be normally set to 'READY'.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Cluster update policy. Choose between [ALWAYS_UPDATE,MINIMAL_DOWNTIME,NEVER_UPDATE]'.
        /// </summary>
        public readonly string? UpdatePolicy;
        /// <summary>
        /// Management URL of your cluster.
        /// </summary>
        public readonly string Url;
        /// <summary>
        /// Kubernetes version of the managed kubernetes cluster.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private GetKubeResult(
            bool controlPlaneIsUpToDate,

            ImmutableArray<Outputs.GetKubeCustomizationApiserverResult> customizationApiservers,

            Outputs.GetKubeCustomizationKubeProxyResult? customizationKubeProxy,

            ImmutableArray<Outputs.GetKubeCustomizationResult> customizations,

            string id,

            bool isUpToDate,

            string kubeId,

            string? kubeProxyMode,

            string loadBalancersSubnetId,

            string? name,

            ImmutableArray<string> nextUpgradeVersions,

            string nodesSubnetId,

            string nodesUrl,

            string privateNetworkId,

            string? region,

            string serviceName,

            string status,

            string? updatePolicy,

            string url,

            string? version)
        {
            ControlPlaneIsUpToDate = controlPlaneIsUpToDate;
            CustomizationApiservers = customizationApiservers;
            CustomizationKubeProxy = customizationKubeProxy;
            Customizations = customizations;
            Id = id;
            IsUpToDate = isUpToDate;
            KubeId = kubeId;
            KubeProxyMode = kubeProxyMode;
            LoadBalancersSubnetId = loadBalancersSubnetId;
            Name = name;
            NextUpgradeVersions = nextUpgradeVersions;
            NodesSubnetId = nodesSubnetId;
            NodesUrl = nodesUrl;
            PrivateNetworkId = privateNetworkId;
            Region = region;
            ServiceName = serviceName;
            Status = status;
            UpdatePolicy = updatePolicy;
            Url = url;
            Version = version;
        }
    }
}
