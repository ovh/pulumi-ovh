// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetKubeNodePoolNodes
    {
        /// <summary>
        /// Use this data source to get a list of OVHcloud Managed Kubernetes nodes in a specific node pool.
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
        ///     var nodes = Ovh.CloudProject.GetKubeNodePoolNodes.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        ///         Name = "XXXXXX",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["nodes"] = nodes,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetKubeNodePoolNodesResult> InvokeAsync(GetKubeNodePoolNodesArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetKubeNodePoolNodesResult>("ovh:CloudProject/getKubeNodePoolNodes:getKubeNodePoolNodes", args ?? new GetKubeNodePoolNodesArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of OVHcloud Managed Kubernetes nodes in a specific node pool.
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
        ///     var nodes = Ovh.CloudProject.GetKubeNodePoolNodes.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        ///         Name = "XXXXXX",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["nodes"] = nodes,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetKubeNodePoolNodesResult> Invoke(GetKubeNodePoolNodesInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeNodePoolNodesResult>("ovh:CloudProject/getKubeNodePoolNodes:getKubeNodePoolNodes", args ?? new GetKubeNodePoolNodesInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get a list of OVHcloud Managed Kubernetes nodes in a specific node pool.
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
        ///     var nodes = Ovh.CloudProject.GetKubeNodePoolNodes.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         KubeId = "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxx",
        ///         Name = "XXXXXX",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["nodes"] = nodes,
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetKubeNodePoolNodesResult> Invoke(GetKubeNodePoolNodesInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetKubeNodePoolNodesResult>("ovh:CloudProject/getKubeNodePoolNodes:getKubeNodePoolNodes", args ?? new GetKubeNodePoolNodesInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetKubeNodePoolNodesArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public string KubeId { get; set; } = null!;

        /// <summary>
        /// Name of the node pool from which we want the nodes.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetKubeNodePoolNodesArgs()
        {
        }
        public static new GetKubeNodePoolNodesArgs Empty => new GetKubeNodePoolNodesArgs();
    }

    public sealed class GetKubeNodePoolNodesInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the managed kubernetes cluster.
        /// </summary>
        [Input("kubeId", required: true)]
        public Input<string> KubeId { get; set; } = null!;

        /// <summary>
        /// Name of the node pool from which we want the nodes.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetKubeNodePoolNodesInvokeArgs()
        {
        }
        public static new GetKubeNodePoolNodesInvokeArgs Empty => new GetKubeNodePoolNodesInvokeArgs();
    }


    [OutputType]
    public sealed class GetKubeNodePoolNodesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string KubeId;
        /// <summary>
        /// Name of the node.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of all nodes composing the kubernetes cluster.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetKubeNodePoolNodesNodeResult> Nodes;
        /// <summary>
        /// See Argument Reference above.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private GetKubeNodePoolNodesResult(
            string id,

            string kubeId,

            string name,

            ImmutableArray<Outputs.GetKubeNodePoolNodesNodeResult> nodes,

            string serviceName)
        {
            Id = id;
            KubeId = kubeId;
            Name = name;
            Nodes = nodes;
            ServiceName = serviceName;
        }
    }
}
