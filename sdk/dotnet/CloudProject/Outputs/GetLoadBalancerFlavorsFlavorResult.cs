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
    public sealed class GetLoadBalancerFlavorsFlavorResult
    {
        /// <summary>
        /// Flavor id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Flavor name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Region name
        /// </summary>
        public readonly string Region;

        [OutputConstructor]
        private GetLoadBalancerFlavorsFlavorResult(
            string id,

            string name,

            string region)
        {
            Id = id;
            Name = name;
            Region = region;
        }
    }
}
