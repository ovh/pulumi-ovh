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
    public sealed class DatabaseNode
    {
        /// <summary>
        /// Private network id in which the node should be deployed. It's the regional openstackId of the private network
        /// </summary>
        public readonly string? NetworkId;
        /// <summary>
        /// Public cloud region in which the node should be deployed. Ex: "GRA'.
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Private subnet ID in which the node is.
        /// </summary>
        public readonly string? SubnetId;

        [OutputConstructor]
        private DatabaseNode(
            string? networkId,

            string region,

            string? subnetId)
        {
            NetworkId = networkId;
            Region = region;
            SubnetId = subnetId;
        }
    }
}
