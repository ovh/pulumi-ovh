// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class ServerReinstallTaskStoragePartitioningLayoutExtra
    {
        /// <summary>
        /// LVM-specific parameters (when applicable)
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerReinstallTaskStoragePartitioningLayoutExtraLv> Lvs;
        /// <summary>
        /// ZFS-specific parameters (when applicable)
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerReinstallTaskStoragePartitioningLayoutExtraZp> Zps;

        [OutputConstructor]
        private ServerReinstallTaskStoragePartitioningLayoutExtra(
            ImmutableArray<Outputs.ServerReinstallTaskStoragePartitioningLayoutExtraLv> lvs,

            ImmutableArray<Outputs.ServerReinstallTaskStoragePartitioningLayoutExtraZp> zps)
        {
            Lvs = lvs;
            Zps = zps;
        }
    }
}
