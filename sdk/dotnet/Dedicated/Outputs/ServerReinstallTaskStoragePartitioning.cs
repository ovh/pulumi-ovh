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
    public sealed class ServerReinstallTaskStoragePartitioning
    {
        /// <summary>
        /// Total number of disks in the disk group involved in the partitioning configuration (all disks of the disk group by default)
        /// </summary>
        public readonly int? Disks;
        /// <summary>
        /// Custom partitioning layout (default is the default layout of the operating system's default partitioning scheme). Accept multiple values (multiple partitions):
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerReinstallTaskStoragePartitioningLayout> Layouts;
        /// <summary>
        /// Partitioning scheme (if applicable with selected operating system)
        /// </summary>
        public readonly string? SchemeName;

        [OutputConstructor]
        private ServerReinstallTaskStoragePartitioning(
            int? disks,

            ImmutableArray<Outputs.ServerReinstallTaskStoragePartitioningLayout> layouts,

            string? schemeName)
        {
            Disks = disks;
            Layouts = layouts;
            SchemeName = schemeName;
        }
    }
}
