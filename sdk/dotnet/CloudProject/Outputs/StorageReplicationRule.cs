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
    public sealed class StorageReplicationRule
    {
        /// <summary>
        /// Delete marker replication
        /// </summary>
        public readonly string? DeleteMarkerReplication;
        /// <summary>
        /// Rule destination configuration
        /// </summary>
        public readonly Outputs.StorageReplicationRuleDestination? Destination;
        /// <summary>
        /// Rule filters
        /// </summary>
        public readonly Outputs.StorageReplicationRuleFilter? Filter;
        /// <summary>
        /// Rule ID
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Rule priority
        /// </summary>
        public readonly double? Priority;
        /// <summary>
        /// Rule status
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private StorageReplicationRule(
            string? deleteMarkerReplication,

            Outputs.StorageReplicationRuleDestination? destination,

            Outputs.StorageReplicationRuleFilter? filter,

            string? id,

            double? priority,

            string? status)
        {
            DeleteMarkerReplication = deleteMarkerReplication;
            Destination = destination;
            Filter = filter;
            Id = id;
            Priority = priority;
            Status = status;
        }
    }
}
