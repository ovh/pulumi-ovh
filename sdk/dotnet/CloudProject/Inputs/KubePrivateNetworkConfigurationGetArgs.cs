// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class KubePrivateNetworkConfigurationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// If defined, all egress traffic will be routed towards this IP address, which should belong to the private network. Empty string means disabled.
        /// </summary>
        [Input("defaultVrackGateway", required: true)]
        public Input<string> DefaultVrackGateway { get; set; } = null!;

        /// <summary>
        /// Defines whether routing should default to using the nodes' private interface, instead of their public interface. Default is false.
        /// 
        /// In order to use the gateway IP advertised by the private network subnet DHCP, the following configuration shall be used.
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        /// });
        /// ```
        /// </summary>
        [Input("privateNetworkRoutingAsDefault", required: true)]
        public Input<bool> PrivateNetworkRoutingAsDefault { get; set; } = null!;

        public KubePrivateNetworkConfigurationGetArgs()
        {
        }
        public static new KubePrivateNetworkConfigurationGetArgs Empty => new KubePrivateNetworkConfigurationGetArgs();
    }
}
