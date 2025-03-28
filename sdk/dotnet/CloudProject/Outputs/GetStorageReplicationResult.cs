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
    public sealed class GetStorageReplicationResult
    {
        /// <summary>
        /// Replication rules
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStorageReplicationRuleResult> Rules;

        [OutputConstructor]
        private GetStorageReplicationResult(ImmutableArray<Outputs.GetStorageReplicationRuleResult> rules)
        {
            Rules = rules;
        }
    }
}
