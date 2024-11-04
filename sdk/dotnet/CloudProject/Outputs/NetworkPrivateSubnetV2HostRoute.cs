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
    public sealed class NetworkPrivateSubnetV2HostRoute
    {
        public readonly string Destination;
        public readonly string Nexthop;

        [OutputConstructor]
        private NetworkPrivateSubnetV2HostRoute(
            string destination,

            string nexthop)
        {
            Destination = destination;
            Nexthop = nexthop;
        }
    }
}