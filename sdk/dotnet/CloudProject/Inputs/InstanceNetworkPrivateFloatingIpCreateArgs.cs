// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class InstanceNetworkPrivateFloatingIpCreateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Floating IP description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        public InstanceNetworkPrivateFloatingIpCreateArgs()
        {
        }
        public static new InstanceNetworkPrivateFloatingIpCreateArgs Empty => new InstanceNetworkPrivateFloatingIpCreateArgs();
    }
}
