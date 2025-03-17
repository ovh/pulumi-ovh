// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class NetworkPrivateRegionsStatusGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// the status of the network. should be normally set to 'ACTIVE'.
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        public NetworkPrivateRegionsStatusGetArgs()
        {
        }
        public static new NetworkPrivateRegionsStatusGetArgs Empty => new NetworkPrivateRegionsStatusGetArgs();
    }
}
