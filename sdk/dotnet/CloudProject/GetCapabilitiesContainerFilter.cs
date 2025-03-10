// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        public static Task<GetCapabilitiesContainerFilterResult> InvokeAsync(GetCapabilitiesContainerFilterArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCapabilitiesContainerFilterResult>("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args ?? new GetCapabilitiesContainerFilterArgs(), options.WithDefaults());

        public static Output<GetCapabilitiesContainerFilterResult> Invoke(GetCapabilitiesContainerFilterInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCapabilitiesContainerFilterResult>("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args ?? new GetCapabilitiesContainerFilterInvokeArgs(), options.WithDefaults());

        public static Output<GetCapabilitiesContainerFilterResult> Invoke(GetCapabilitiesContainerFilterInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCapabilitiesContainerFilterResult>("ovh:CloudProject/getCapabilitiesContainerFilter:getCapabilitiesContainerFilter", args ?? new GetCapabilitiesContainerFilterInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCapabilitiesContainerFilterArgs : global::Pulumi.InvokeArgs
    {
        [Input("planName", required: true)]
        public string PlanName { get; set; } = null!;

        [Input("region", required: true)]
        public string Region { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetCapabilitiesContainerFilterArgs()
        {
        }
        public static new GetCapabilitiesContainerFilterArgs Empty => new GetCapabilitiesContainerFilterArgs();
    }

    public sealed class GetCapabilitiesContainerFilterInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("planName", required: true)]
        public Input<string> PlanName { get; set; } = null!;

        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

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
        public readonly string Code;
        public readonly string CreatedAt;
        public readonly ImmutableArray<Outputs.GetCapabilitiesContainerFilterFeatureResult> Features;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string PlanName;
        public readonly string Region;
        public readonly ImmutableArray<Outputs.GetCapabilitiesContainerFilterRegistryLimitResult> RegistryLimits;
        public readonly string ServiceName;
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
