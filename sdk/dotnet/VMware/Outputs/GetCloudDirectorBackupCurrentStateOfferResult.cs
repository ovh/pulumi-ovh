// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.VMware.Outputs
{

    [OutputType]
    public sealed class GetCloudDirectorBackupCurrentStateOfferResult
    {
        /// <summary>
        /// Backup service offer type (BRONZE|SILVER|GOLD)
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Backup repository primary region
        /// </summary>
        public readonly string ProtectionPrimaryRegion;
        /// <summary>
        /// Backup repository replicated region
        /// </summary>
        public readonly string ProtectionReplicatedRegion;
        /// <summary>
        /// Backup repository quota in TB
        /// </summary>
        public readonly double QuotaInTb;
        /// <summary>
        /// Backup offer status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// Backup repository used space in GB
        /// </summary>
        public readonly double UsedSpaceInGb;

        [OutputConstructor]
        private GetCloudDirectorBackupCurrentStateOfferResult(
            string name,

            string protectionPrimaryRegion,

            string protectionReplicatedRegion,

            double quotaInTb,

            string status,

            double usedSpaceInGb)
        {
            Name = name;
            ProtectionPrimaryRegion = protectionPrimaryRegion;
            ProtectionReplicatedRegion = protectionReplicatedRegion;
            QuotaInTb = quotaInTb;
            Status = status;
            UsedSpaceInGb = usedSpaceInGb;
        }
    }
}
