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
    public sealed class ServerStoragePartitioningLayout
    {
        /// <summary>
        /// Partition extras parameters
        /// </summary>
        public readonly Outputs.ServerStoragePartitioningLayoutExtras? Extras;
        /// <summary>
        /// File system type
        /// </summary>
        public readonly string FileSystem;
        /// <summary>
        /// Mount point
        /// </summary>
        public readonly string MountPoint;
        /// <summary>
        /// Software raid type (default is 1)
        /// </summary>
        public readonly double? RaidLevel;
        /// <summary>
        /// Partition size in MiB (default value is 0 which means to fill the disk with that partition)
        /// </summary>
        public readonly double? Size;

        [OutputConstructor]
        private ServerStoragePartitioningLayout(
            Outputs.ServerStoragePartitioningLayoutExtras? extras,

            string fileSystem,

            string mountPoint,

            double? raidLevel,

            double? size)
        {
            Extras = extras;
            FileSystem = fileSystem;
            MountPoint = mountPoint;
            RaidLevel = raidLevel;
            Size = size;
        }
    }
}
