// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerReinstallTaskStorageArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk group id to install the OS to (default is 0, meaning automatic).
        /// </summary>
        [Input("diskGroupId")]
        public Input<int>? DiskGroupId { get; set; }

        [Input("hardwareRaids")]
        private InputList<Inputs.ServerReinstallTaskStorageHardwareRaidArgs>? _hardwareRaids;

        /// <summary>
        /// Hardware Raid configurations (if not specified, all disks of the chosen disk group id will be configured in JBOD mode).
        /// </summary>
        public InputList<Inputs.ServerReinstallTaskStorageHardwareRaidArgs> HardwareRaids
        {
            get => _hardwareRaids ?? (_hardwareRaids = new InputList<Inputs.ServerReinstallTaskStorageHardwareRaidArgs>());
            set => _hardwareRaids = value;
        }

        [Input("partitionings")]
        private InputList<Inputs.ServerReinstallTaskStoragePartitioningArgs>? _partitionings;

        /// <summary>
        /// Partitioning configuration
        /// </summary>
        public InputList<Inputs.ServerReinstallTaskStoragePartitioningArgs> Partitionings
        {
            get => _partitionings ?? (_partitionings = new InputList<Inputs.ServerReinstallTaskStoragePartitioningArgs>());
            set => _partitionings = value;
        }

        public ServerReinstallTaskStorageArgs()
        {
        }
        public static new ServerReinstallTaskStorageArgs Empty => new ServerReinstallTaskStorageArgs();
    }
}
