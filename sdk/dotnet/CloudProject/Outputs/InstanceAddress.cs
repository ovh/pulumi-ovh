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
    public sealed class InstanceAddress
    {
        /// <summary>
        /// IP address
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// IP version
        /// </summary>
        public readonly int? Version;

        [OutputConstructor]
        private InstanceAddress(
            string? ip,

            int? version)
        {
            Ip = ip;
            Version = version;
        }
    }
}
