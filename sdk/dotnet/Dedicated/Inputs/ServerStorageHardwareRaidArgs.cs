// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerStorageHardwareRaidArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of arrays (default is 1)
        /// </summary>
        [Input("arrays")]
        public Input<double>? Arrays { get; set; }

        /// <summary>
        /// Total number of disks in the disk group involved in the hardware raid configuration (all disks of the disk group by default)
        /// </summary>
        [Input("disks")]
        public Input<double>? Disks { get; set; }

        /// <summary>
        /// Hardware raid type (default is 1)
        /// </summary>
        [Input("raidLevel")]
        public Input<double>? RaidLevel { get; set; }

        /// <summary>
        /// Number of disks in the disk group involved in the spare (default is 0)
        /// </summary>
        [Input("spares")]
        public Input<double>? Spares { get; set; }

        public ServerStorageHardwareRaidArgs()
        {
        }
        public static new ServerStorageHardwareRaidArgs Empty => new ServerStorageHardwareRaidArgs();
    }
}
