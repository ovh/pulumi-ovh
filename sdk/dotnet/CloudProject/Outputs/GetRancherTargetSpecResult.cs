// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class GetRancherTargetSpecResult
    {
        /// <summary>
        /// List of allowed CIDR blocks for a managed Rancher service's IP restrictions. When empty, any IP is allowed
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRancherTargetSpecIpRestrictionResult> IpRestrictions;
        /// <summary>
        /// Name of the managed Rancher service
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Plan of the managed Rancher service. Available plans for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/plan
        /// </summary>
        public readonly string Plan;
        /// <summary>
        /// Version of the managed Rancher service. Available versions for an existing managed Rancher can be retrieved using GET /rancher/rancherID/capabilities/version
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetRancherTargetSpecResult(
            ImmutableArray<Outputs.GetRancherTargetSpecIpRestrictionResult> ipRestrictions,

            string name,

            string plan,

            string version)
        {
            IpRestrictions = ipRestrictions;
            Name = name;
            Plan = plan;
            Version = version;
        }
    }
}
