// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain.Inputs
{

    public sealed class NameCurrentStateDnsConfigurationNameServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The IPv4 associated to the name server
        /// </summary>
        [Input("ipv4")]
        public Input<string>? Ipv4 { get; set; }

        /// <summary>
        /// The IPv6 associated to the name server
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        /// <summary>
        /// The host name
        /// </summary>
        [Input("nameServer")]
        public Input<string>? NameServer { get; set; }

        /// <summary>
        /// The type of name server
        /// </summary>
        [Input("nameServerType")]
        public Input<string>? NameServerType { get; set; }

        public NameCurrentStateDnsConfigurationNameServerArgs()
        {
        }
        public static new NameCurrentStateDnsConfigurationNameServerArgs Empty => new NameCurrentStateDnsConfigurationNameServerArgs();
    }
}
