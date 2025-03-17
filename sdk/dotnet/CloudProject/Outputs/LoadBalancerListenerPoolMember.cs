// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class LoadBalancerListenerPoolMember
    {
        /// <summary>
        /// IP address of the resource
        /// </summary>
        public readonly string? Address;
        /// <summary>
        /// Name of the member
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Protocol port number for the resource
        /// </summary>
        public readonly double? ProtocolPort;
        /// <summary>
        /// Weight of a member determines the portion of requests or connections it services compared to the other members of the pool. Between 1 and 256.
        /// </summary>
        public readonly double? Weight;

        [OutputConstructor]
        private LoadBalancerListenerPoolMember(
            string? address,

            string? name,

            double? protocolPort,

            double? weight)
        {
            Address = address;
            Name = name;
            ProtocolPort = protocolPort;
            Weight = weight;
        }
    }
}
