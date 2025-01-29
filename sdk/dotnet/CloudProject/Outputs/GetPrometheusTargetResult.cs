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
    public sealed class GetPrometheusTargetResult
    {
        /// <summary>
        /// Host of the endpoint
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Connection port for the endpoint
        /// </summary>
        public readonly int Port;

        [OutputConstructor]
        private GetPrometheusTargetResult(
            string host,

            int port)
        {
            Host = host;
            Port = port;
        }
    }
}
