// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class GetServerSpecificationsNetworkOlaResult
    {
        /// <summary>
        /// Is the OLA feature available
        /// </summary>
        public readonly bool Available;
        /// <summary>
        /// What modes are supported
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerSpecificationsNetworkOlaAvailableModeResult> AvailableModes;
        /// <summary>
        /// (DEPRECATED) What modes are supported
        /// </summary>
        public readonly ImmutableArray<string> SupportedModes;

        [OutputConstructor]
        private GetServerSpecificationsNetworkOlaResult(
            bool available,

            ImmutableArray<Outputs.GetServerSpecificationsNetworkOlaAvailableModeResult> availableModes,

            ImmutableArray<string> supportedModes)
        {
            Available = available;
            AvailableModes = availableModes;
            SupportedModes = supportedModes;
        }
    }
}
