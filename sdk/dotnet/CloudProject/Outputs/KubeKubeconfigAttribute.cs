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
    public sealed class KubeKubeconfigAttribute
    {
        public readonly string? ClientCertificate;
        public readonly string? ClientKey;
        public readonly string? ClusterCaCertificate;
        public readonly string? Host;

        [OutputConstructor]
        private KubeKubeconfigAttribute(
            string? clientCertificate,

            string? clientKey,

            string? clusterCaCertificate,

            string? host)
        {
            ClientCertificate = clientCertificate;
            ClientKey = clientKey;
            ClusterCaCertificate = clusterCaCertificate;
            Host = host;
        }
    }
}
