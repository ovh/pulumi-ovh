// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class GatewayExternalInformationGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ips")]
        private InputList<Inputs.GatewayExternalInformationIpGetArgs>? _ips;

        /// <summary>
        /// List of external ips of the gateway
        /// </summary>
        public InputList<Inputs.GatewayExternalInformationIpGetArgs> Ips
        {
            get => _ips ?? (_ips = new InputList<Inputs.GatewayExternalInformationIpGetArgs>());
            set => _ips = value;
        }

        /// <summary>
        /// External network ID of the gateway
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        public GatewayExternalInformationGetArgs()
        {
        }
        public static new GatewayExternalInformationGetArgs Empty => new GatewayExternalInformationGetArgs();
    }
}
