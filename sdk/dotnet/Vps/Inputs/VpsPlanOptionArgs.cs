// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Inputs
{

    public sealed class VpsPlanOptionArgs : global::Pulumi.ResourceArgs
    {
        [Input("configurations")]
        private InputList<Inputs.VpsPlanOptionConfigurationArgs>? _configurations;

        /// <summary>
        /// Representation of a configuration item for personalizing product
        /// </summary>
        public InputList<Inputs.VpsPlanOptionConfigurationArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.VpsPlanOptionConfigurationArgs>());
            set => _configurations = value;
        }

        /// <summary>
        /// duration
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        /// <summary>
        /// Plan code
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Pricing model identifier
        /// </summary>
        [Input("pricingMode", required: true)]
        public Input<string> PricingMode { get; set; } = null!;

        /// <summary>
        /// Quantity of product desired
        /// </summary>
        [Input("quantity", required: true)]
        public Input<double> Quantity { get; set; } = null!;

        public VpsPlanOptionArgs()
        {
        }
        public static new VpsPlanOptionArgs Empty => new VpsPlanOptionArgs();
    }
}
