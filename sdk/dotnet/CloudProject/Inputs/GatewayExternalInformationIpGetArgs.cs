// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class GatewayExternalInformationIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// External IP of the gateway
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Subnet ID of the ip
        /// </summary>
        [Input("subnetId")]
        public Input<string>? SubnetId { get; set; }

        public GatewayExternalInformationIpGetArgs()
        {
        }
        public static new GatewayExternalInformationIpGetArgs Empty => new GatewayExternalInformationIpGetArgs();
    }
}
