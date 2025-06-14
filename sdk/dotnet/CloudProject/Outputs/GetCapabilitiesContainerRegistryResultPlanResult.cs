// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class GetCapabilitiesContainerRegistryResultPlanResult
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
        public readonly ImmutableArray<Outputs.GetCapabilitiesContainerRegistryResultPlanFeatureResult> Features;
        /// <summary>
        /// Plan ID
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Plan name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Container registry limits
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCapabilitiesContainerRegistryResultPlanRegistryLimitResult> RegistryLimits;
        /// <summary>
        /// Plan last update date
        /// </summary>
        public readonly string UpdatedAt;

        [OutputConstructor]
        private GetCapabilitiesContainerRegistryResultPlanResult(
            string code,

            string createdAt,

            ImmutableArray<Outputs.GetCapabilitiesContainerRegistryResultPlanFeatureResult> features,

            string id,

            string name,

            ImmutableArray<Outputs.GetCapabilitiesContainerRegistryResultPlanRegistryLimitResult> registryLimits,

            string updatedAt)
        {
            Code = code;
            CreatedAt = createdAt;
            Features = features;
            Id = id;
            Name = name;
            RegistryLimits = registryLimits;
            UpdatedAt = updatedAt;
        }
    }
}
