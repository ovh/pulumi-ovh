// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadBalancerNetworkPrivateGatewayGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the gateway
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public LoadBalancerNetworkPrivateGatewayGetArgs()
        {
        }
        public static new LoadBalancerNetworkPrivateGatewayGetArgs Empty => new LoadBalancerNetworkPrivateGatewayGetArgs();
    }
}
