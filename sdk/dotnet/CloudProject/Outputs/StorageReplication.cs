// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class StorageReplication
    {
        /// <summary>
        /// Replication rules
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageReplicationRule> Rules;

        [OutputConstructor]
        private StorageReplication(ImmutableArray<Outputs.StorageReplicationRule> rules)
        {
            Rules = rules;
        }
    }
}
