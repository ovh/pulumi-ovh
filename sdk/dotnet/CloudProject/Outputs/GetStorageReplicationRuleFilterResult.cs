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
    public sealed class GetStorageReplicationRuleFilterResult
    {
        /// <summary>
        /// Prefix filter
        /// </summary>
        public readonly string Prefix;
        /// <summary>
        /// Tags filter
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;

        [OutputConstructor]
        private GetStorageReplicationRuleFilterResult(
            string prefix,

            ImmutableDictionary<string, string> tags)
        {
            Prefix = prefix;
            Tags = tags;
        }
    }
}
