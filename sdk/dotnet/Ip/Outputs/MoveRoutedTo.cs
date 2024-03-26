// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Ip.Outputs
{

    [OutputType]
    public sealed class MoveRoutedTo
    {
        /// <summary>
        /// Name of the service to route the IP to. IP will be parked if this value is an empty string
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private MoveRoutedTo(string serviceName)
        {
            ServiceName = serviceName;
        }
    }
}
