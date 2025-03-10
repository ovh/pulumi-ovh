// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class NetworkPrivateSubnetIpPoolArgs : global::Pulumi.ResourceArgs
    {
        [Input("dhcp")]
        public Input<bool>? Dhcp { get; set; }

        [Input("end")]
        public Input<string>? End { get; set; }

        [Input("network")]
        public Input<string>? Network { get; set; }

        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("start")]
        public Input<string>? Start { get; set; }

        public NetworkPrivateSubnetIpPoolArgs()
        {
        }
        public static new NetworkPrivateSubnetIpPoolArgs Empty => new NetworkPrivateSubnetIpPoolArgs();
    }
}
