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
    public sealed class GetStoragesContainerReplicationResult
    {
        /// <summary>
        /// Replication rules
        /// </summary>
        public readonly ImmutableArray<Outputs.GetStoragesContainerReplicationRuleResult> Rules;

        [OutputConstructor]
        private GetStoragesContainerReplicationResult(ImmutableArray<Outputs.GetStoragesContainerReplicationRuleResult> rules)
        {
            Rules = rules;
        }
    }
}
