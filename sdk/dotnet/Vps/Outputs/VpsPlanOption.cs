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
    public sealed class VpsPlanOption
    {
        public readonly ImmutableArray<Outputs.VpsPlanOptionConfiguration> Configurations;
        /// <summary>
        /// Duration selected for the purchase of the product
        /// </summary>
        public readonly string Duration;
        /// <summary>
        /// Identifier of the option offer
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Pricing mode selected for the purchase of the product
        /// </summary>
        public readonly string PricingMode;
        /// <summary>
        /// Quantity of product desired
        /// </summary>
        public readonly double Quantity;

        [OutputConstructor]
        private VpsPlanOption(
            ImmutableArray<Outputs.VpsPlanOptionConfiguration> configurations,

            string duration,

            string planCode,

            string pricingMode,

            double quantity)
        {
            Configurations = configurations;
            Duration = duration;
            PlanCode = planCode;
            PricingMode = pricingMode;
            Quantity = quantity;
        }
    }
}
