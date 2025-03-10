// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerStoragePartitioningArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Total number of disks in the disk group involved in the partitioning configuration (all disks of the disk group by default)
        /// </summary>
        [Input("disks")]
        public Input<double>? Disks { get; set; }

        [Input("layouts")]
        private InputList<Inputs.ServerStoragePartitioningLayoutArgs>? _layouts;

        /// <summary>
        /// Custom partitioning layout (default is the default layout of the operating system's default partitioning scheme)
        /// </summary>
        public InputList<Inputs.ServerStoragePartitioningLayoutArgs> Layouts
        {
            get => _layouts ?? (_layouts = new InputList<Inputs.ServerStoragePartitioningLayoutArgs>());
            set => _layouts = value;
        }

        /// <summary>
        /// Partitioning scheme (if applicable with selected operating system)
        /// </summary>
        [Input("schemeName")]
        public Input<string>? SchemeName { get; set; }

        public ServerStoragePartitioningArgs()
        {
        }
        public static new ServerStoragePartitioningArgs Empty => new ServerStoragePartitioningArgs();
    }
}
