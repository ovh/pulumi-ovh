// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Scraly.PulumiPackage.Ovh.IpLoadBalancing.Inputs
{

    public sealed class LoadBalancerOrderableZoneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The zone three letter code
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Plan code
        /// </summary>
        [Input("planCode")]
        public Input<string>? PlanCode { get; set; }

        public LoadBalancerOrderableZoneArgs()
        {
        }
        public static new LoadBalancerOrderableZoneArgs Empty => new LoadBalancerOrderableZoneArgs();
    }
}