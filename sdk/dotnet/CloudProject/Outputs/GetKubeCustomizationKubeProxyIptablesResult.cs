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
    public sealed class GetKubeCustomizationKubeProxyIptablesResult
    {
        /// <summary>
        /// Minimum period that IPVS rules are refreshed in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration.
        /// </summary>
        public readonly string? MinSyncPeriod;
        /// <summary>
        /// Minimum period that IPVS rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format.
        /// </summary>
        public readonly string? SyncPeriod;

        [OutputConstructor]
        private GetKubeCustomizationKubeProxyIptablesResult(
            string? minSyncPeriod,

            string? syncPeriod)
        {
            MinSyncPeriod = minSyncPeriod;
            SyncPeriod = syncPeriod;
        }
    }
}
