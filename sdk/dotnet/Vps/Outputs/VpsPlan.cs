// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Outputs
{

    [OutputType]
    public sealed class VpsPlan
    {
        /// <summary>
        /// Representation of a configuration item for personalizing product
        /// </summary>
        public readonly ImmutableArray<Outputs.VpsPlanConfiguration> Configurations;
        /// <summary>
        /// duration
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Cart item to be linked
        /// </summary>
        public readonly double? ItemId;
        /// <summary>
        /// Plan code
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Pricing model identifier
        /// </summary>
        public readonly string PricingMode;
        /// <summary>
        /// Quantity of product desired
        /// </summary>
        public readonly double? Quantity;

        [OutputConstructor]
        private VpsPlan(
            ImmutableArray<Outputs.VpsPlanConfiguration> configurations,

            string duration,

            double? itemId,

            string planCode,

            string pricingMode,

            double? quantity)
        {
            Configurations = configurations;
            Duration = duration;
            ItemId = itemId;
            PlanCode = planCode;
            PricingMode = pricingMode;
            Quantity = quantity;
        }
    }
}
