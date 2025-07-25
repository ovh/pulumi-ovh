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
    public sealed class ContainerRegistryPlan
    {
        /// <summary>
        /// Plan code from the catalog
        /// </summary>
        public readonly string? Code;
        /// <summary>
        /// Plan creation date
        /// </summary>
        public readonly string? CreatedAt;
        /// <summary>
        /// Features of the plan
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerRegistryPlanFeature> Features;
        /// <summary>
        /// Plan ID
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Registry name
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Container registry limits
        /// </summary>
        public readonly ImmutableArray<Outputs.ContainerRegistryPlanRegistryLimit> RegistryLimits;
        /// <summary>
        /// Registry last update date
        /// </summary>
        public readonly string? UpdatedAt;

        [OutputConstructor]
        private ContainerRegistryPlan(
            string? code,

            string? createdAt,

            ImmutableArray<Outputs.ContainerRegistryPlanFeature> features,

            string? id,

            string? name,

            ImmutableArray<Outputs.ContainerRegistryPlanRegistryLimit> registryLimits,

            string? updatedAt)
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
