// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Okms.Outputs
{

    [OutputType]
    public sealed class GetOkmsServiceKeyJwkIamResult
    {
        /// <summary>
        /// Resource display name
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// ID of the service key
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource tags. Tags that were internally computed are prefixed with ovh:
        /// </summary>
        public readonly ImmutableDictionary<string, string> Tags;
        /// <summary>
        /// Unique resource name used in policies
        /// </summary>
        public readonly string Urn;

        [OutputConstructor]
        private GetOkmsServiceKeyJwkIamResult(
            string displayName,

            string id,

            ImmutableDictionary<string, string> tags,

            string urn)
        {
            DisplayName = displayName;
            Id = id;
            Tags = tags;
            Urn = urn;
        }
    }
}
