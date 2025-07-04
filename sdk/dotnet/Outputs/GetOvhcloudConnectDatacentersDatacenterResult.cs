// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Outputs
{

    [OutputType]
    public sealed class GetOvhcloudConnectDatacentersDatacenterResult
    {
        /// <summary>
        /// Get availability to add new configuration on it
        /// </summary>
        public readonly bool Available;
        /// <summary>
        /// Id
        /// </summary>
        public readonly double Id;
        /// <summary>
        /// name of the datacenter
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// region of the datacenter
        /// </summary>
        public readonly string Region;
        /// <summary>
        /// region type of the datacenter
        /// </summary>
        public readonly string RegionType;

        [OutputConstructor]
        private GetOvhcloudConnectDatacentersDatacenterResult(
            bool available,

            double id,

            string name,

            string region,

            string regionType)
        {
            Available = available;
            Id = id;
            Name = name;
            Region = region;
            RegionType = regionType;
        }
    }
}
