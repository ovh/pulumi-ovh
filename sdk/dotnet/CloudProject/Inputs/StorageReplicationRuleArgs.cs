// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class StorageReplicationRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Delete marker replication
        /// </summary>
        [Input("deleteMarkerReplication")]
        public Input<string>? DeleteMarkerReplication { get; set; }

        /// <summary>
        /// Rule destination configuration
        /// </summary>
        [Input("destination")]
        public Input<Inputs.StorageReplicationRuleDestinationArgs>? Destination { get; set; }

        /// <summary>
        /// Rule filters
        /// </summary>
        [Input("filter")]
        public Input<Inputs.StorageReplicationRuleFilterArgs>? Filter { get; set; }

        /// <summary>
        /// Rule ID
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Rule priority
        /// </summary>
        [Input("priority")]
        public Input<double>? Priority { get; set; }

        /// <summary>
        /// Rule status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public StorageReplicationRuleArgs()
        {
        }
        public static new StorageReplicationRuleArgs Empty => new StorageReplicationRuleArgs();
    }
}
