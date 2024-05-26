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
    public sealed class GetServerSpecificationsNetworkOlaAvailableModeResult
    {
        /// <summary>
        /// Whether it is the default configuration of the server
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// Interface layout
        /// </summary>
        public readonly ImmutableArray<Outputs.GetServerSpecificationsNetworkOlaAvailableModeInterfaceResult> Interfaces;
        /// <summary>
        /// Switch name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetServerSpecificationsNetworkOlaAvailableModeResult(
            bool @default,

            ImmutableArray<Outputs.GetServerSpecificationsNetworkOlaAvailableModeInterfaceResult> interfaces,

            string name)
        {
            Default = @default;
            Interfaces = interfaces;
            Name = name;
        }
    }
}
