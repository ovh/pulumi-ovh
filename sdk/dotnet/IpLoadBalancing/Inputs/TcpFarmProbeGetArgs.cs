// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.IpLoadBalancing.Inputs
{

    public sealed class TcpFarmProbeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("forceSsl")]
        public Input<bool>? ForceSsl { get; set; }

        [Input("interval")]
        public Input<int>? Interval { get; set; }

        [Input("match")]
        public Input<string>? Match { get; set; }

        [Input("method")]
        public Input<string>? Method { get; set; }

        [Input("negate")]
        public Input<bool>? Negate { get; set; }

        [Input("pattern")]
        public Input<string>? Pattern { get; set; }

        [Input("port")]
        public Input<int>? Port { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("url")]
        public Input<string>? Url { get; set; }

        public TcpFarmProbeGetArgs()
        {
        }
        public static new TcpFarmProbeGetArgs Empty => new TcpFarmProbeGetArgs();
    }
}
