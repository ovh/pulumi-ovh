// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain.Inputs
{

    public sealed class NamePlanArgs : global::Pulumi.ResourceArgs
    {
        [Input("configurations")]
        private InputList<Inputs.NamePlanConfigurationArgs>? _configurations;
        public InputList<Inputs.NamePlanConfigurationArgs> Configurations
        {
            get => _configurations ?? (_configurations = new InputList<Inputs.NamePlanConfigurationArgs>());
            set => _configurations = value;
        }

        /// <summary>
        /// Duration selected for the purchase of the product
        /// </summary>
        [Input("duration", required: true)]
        public Input<string> Duration { get; set; } = null!;

        /// <summary>
        /// Cart item to be linked
        /// </summary>
        [Input("itemId")]
        public Input<double>? ItemId { get; set; }

        /// <summary>
        /// Identifier of the option offer
        /// </summary>
        [Input("planCode", required: true)]
        public Input<string> PlanCode { get; set; } = null!;

        /// <summary>
        /// Pricing mode selected for the purchase of the product
        /// </summary>
        [Input("pricingMode", required: true)]
        public Input<string> PricingMode { get; set; } = null!;

        /// <summary>
        /// Quantity of product desired
        /// </summary>
        [Input("quantity")]
        public Input<double>? Quantity { get; set; }

        public NamePlanArgs()
        {
        }
        public static new NamePlanArgs Empty => new NamePlanArgs();
    }
}
