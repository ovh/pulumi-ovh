// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain.Inputs
{

    public sealed class NamePlanOptionConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Label for your configuration item
        /// </summary>
        [Input("label", required: true)]
        public Input<string> Label { get; set; } = null!;

        /// <summary>
        /// Value or resource URL on API.OVH.COM of your configuration item
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NamePlanOptionConfigurationArgs()
        {
        }
        public static new NamePlanOptionConfigurationArgs Empty => new NamePlanOptionConfigurationArgs();
    }
}
