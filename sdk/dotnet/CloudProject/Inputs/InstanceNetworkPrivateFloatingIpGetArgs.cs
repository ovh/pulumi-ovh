// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class InstanceNetworkPrivateFloatingIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Floating IP ID
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public InstanceNetworkPrivateFloatingIpGetArgs()
        {
        }
        public static new InstanceNetworkPrivateFloatingIpGetArgs Empty => new InstanceNetworkPrivateFloatingIpGetArgs();
    }
}
