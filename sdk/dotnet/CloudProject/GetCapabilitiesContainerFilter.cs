// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    public static class GetCapabilitiesContainerFilter
    {
        /// <summary>
        /// Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.
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
        ///     var capability = Ovh.CloudProject.GetCapabilitiesContainerFilter.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Region = "GRA",
        ///         PlanName = "SMALL",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetCapabilitiesContainerFilterResult> InvokeAsync(GetCapabilitiesContainerFilterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCapabilitiesContainerFilterResult>("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args ?? new GetCapabilitiesContainerFilterArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.
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
        ///     var capability = Ovh.CloudProject.GetCapabilitiesContainerFilter.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Region = "GRA",
        ///         PlanName = "SMALL",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCapabilitiesContainerFilterResult> Invoke(GetCapabilitiesContainerFilterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCapabilitiesContainerFilterResult>("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args ?? new GetCapabilitiesContainerFilterInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to filter the list of container registry capabilities associated with a public cloud project to match one and only one capability.
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
        ///     var capability = Ovh.CloudProject.GetCapabilitiesContainerFilter.Invoke(new()
        ///     {
        ///         ServiceName = "XXXXXX",
        ///         Region = "GRA",
        ///         PlanName = "SMALL",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetCapabilitiesContainerFilterResult> Invoke(GetCapabilitiesContainerFilterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCapabilitiesContainerFilterResult>("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args ?? new GetCapabilitiesContainerFilterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCapabilitiesContainerFilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
        /// </summary>
        [Input("planName", required: true)]
        public string PlanName { get; set; } = null!;

        /// <summary>
        /// The region name
        /// </summary>
        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCapabilitiesContainerFilterArgs()
        {
        }
        public static new GetCapabilitiesContainerFilterArgs Empty => new GetCapabilitiesContainerFilterArgs();
    }

    public sealed class GetCapabilitiesContainerFilterInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The plan name. It can be 'SMALL', 'MEDIUM' or 'LARGE'.
        /// </summary>
        [Input("planName", required: true)]
        public Input<string> PlanName { get; set; } = null!;

        /// <summary>
        /// The region name
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public GetCapabilitiesContainerFilterInvokeArgs()
        {
        }
        public static new GetCapabilitiesContainerFilterInvokeArgs Empty => new GetCapabilitiesContainerFilterInvokeArgs();
    }


    [OutputType]
    public sealed class GetCapabilitiesContainerFilterResult
    {
        /// <summary>
        /// Plan code from the catalog
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Plan creation date
        /// </summary>
        public readonly string CreatedAt;
        /// <summary>
        /// Features of the plan
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCapabilitiesContainerFilterFeatureResult> Features;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Plan name
        /// </summary>
        public readonly string Name;
        public readonly string PlanName;
        public readonly string Region;
        /// <summary>
        /// Container registry limits
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCapabilitiesContainerFilterRegistryLimitResult> RegistryLimits;
        public readonly string ServiceName;
        /// <summary>
        /// Plan last update date
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetCapabilitiesContainerFilterResult(
            string code,

            string createdAt,

            ImmutableArray<Outputs.GetCapabilitiesContainerFilterFeatureResult> features,

            string id,

            string name,

            string planName,

            string region,

            ImmutableArray<Outputs.GetCapabilitiesContainerFilterRegistryLimitResult> registryLimits,

            string serviceName,

            string updatedAt)
        {
            Code = code;
            CreatedAt = createdAt;
            Features = features;
            Id = id;
            Name = name;
            PlanName = planName;
            Region = region;
            RegistryLimits = registryLimits;
            ServiceName = serviceName;
            UpdatedAt = updatedAt;
        }
    }
}
