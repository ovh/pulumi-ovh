// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class LoadbalancerFloatingIpGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// IP Address of the resource
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public LoadbalancerFloatingIpGetArgs()
        {
        }
        public static new LoadbalancerFloatingIpGetArgs Empty => new LoadbalancerFloatingIpGetArgs();
    }
}
