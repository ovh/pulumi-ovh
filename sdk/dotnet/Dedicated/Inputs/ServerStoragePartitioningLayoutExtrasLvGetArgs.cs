// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerStoragePartitioningLayoutExtrasLvGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Logical volume name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ServerStoragePartitioningLayoutExtrasLvGetArgs()
        {
        }
        public static new ServerStoragePartitioningLayoutExtrasLvGetArgs Empty => new ServerStoragePartitioningLayoutExtrasLvGetArgs();
    }
}
