// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class GatewayInterfaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the interface
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// IP of the interface
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Network ID of the interface
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// Subnet ID of the interface
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public GatewayInterfaceArgs()
        {
        }
        public static new GatewayInterfaceArgs Empty => new GatewayInterfaceArgs();
    }
}
