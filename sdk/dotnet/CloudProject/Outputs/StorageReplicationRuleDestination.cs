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
    public sealed class StorageReplicationRuleDestination
    {
        /// <summary>
        /// Destination bucket name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Destination region
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// Destination storage class
        /// </summary>
        public readonly string? StorageClass;

        [OutputConstructor]
        private StorageReplicationRuleDestination(
            string name,

            string region,

            string? storageClass)
        {
            Name = name;
            Region = region;
            StorageClass = storageClass;
        }
    }
}
