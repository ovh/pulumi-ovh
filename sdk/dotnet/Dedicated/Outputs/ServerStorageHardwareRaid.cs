// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class ServerStorageHardwareRaid
    {
        /// <summary>
        /// Number of arrays (default is 1)
        /// </summary>
        public readonly double? Arrays;
        /// <summary>
        /// Total number of disks in the disk group involved in the hardware raid configuration (all disks of the disk group by default)
        /// </summary>
        public readonly double? Disks;
        /// <summary>
        /// Hardware raid type (default is 1)
        /// </summary>
        public readonly double? RaidLevel;
        /// <summary>
        /// Number of disks in the disk group involved in the spare (default is 0)
        /// </summary>
        public readonly double? Spares;

        [OutputConstructor]
        private ServerStorageHardwareRaid(
            double? arrays,

            double? disks,

            double? raidLevel,

            double? spares)
        {
            Arrays = arrays;
            Disks = disks;
            RaidLevel = raidLevel;
            Spares = spares;
        }
    }
}
