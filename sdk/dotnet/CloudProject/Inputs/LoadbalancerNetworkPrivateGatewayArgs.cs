// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadbalancerNetworkPrivateGatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the gateway
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public LoadbalancerNetworkPrivateGatewayArgs()
        {
        }
        public static new LoadbalancerNetworkPrivateGatewayArgs Empty => new LoadbalancerNetworkPrivateGatewayArgs();
    }
}
