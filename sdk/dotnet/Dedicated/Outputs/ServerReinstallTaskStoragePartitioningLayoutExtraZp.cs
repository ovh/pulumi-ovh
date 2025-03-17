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
    public sealed class ServerReinstallTaskStoragePartitioningLayoutExtraZp
    {
        /// <summary>
        /// zpool name (generated automatically if not specified, note that multiple ZFS partitions with same zpool names will be configured as multiple datasets belonging to the same zpool if compatible)
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ServerReinstallTaskStoragePartitioningLayoutExtraZp(string? name)
        {
            Name = name;
        }
    }
}
