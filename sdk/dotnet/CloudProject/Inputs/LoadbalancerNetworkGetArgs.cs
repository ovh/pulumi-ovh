// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadbalancerNetworkGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information to private network
        /// </summary>
        [Input("private", required: true)]
        public Input<Inputs.LoadbalancerNetworkPrivateGetArgs> Private { get; set; } = null!;

        public LoadbalancerNetworkGetArgs()
        {
        }
        public static new LoadbalancerNetworkGetArgs Empty => new LoadbalancerNetworkGetArgs();
    }
}
