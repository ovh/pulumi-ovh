// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.OVHcloud.Outputs
{

    [OutputType]
    public sealed class ConnectsOccResult
    {
        /// <summary>
        /// Service bandwidth
        /// </summary>
        public readonly string Bandwidth;
        /// <summary>
        /// Service description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// IAM resource information
        /// </summary>
        public readonly Outputs.ConnectsOccIamResult Iam;
        /// <summary>
        /// List of interfaces linked to a service
        /// </summary>
        public readonly ImmutableArray<double> InterfaceLists;
        /// <summary>
        /// Pop reference where the service is delivered
        /// </summary>
        public readonly string Pop;
        /// <summary>
        /// Port quantity
        /// </summary>
        public readonly string PortQuantity;
        /// <summary>
        /// Product name of the service
        /// </summary>
        public readonly string Product;
        /// <summary>
        /// Service provider
        /// </summary>
        public readonly string ProviderName;
        /// <summary>
        /// Service name
        /// </summary>
        public readonly string ServiceName;
        /// <summary>
        /// Service status
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// uuid of the Ovhcloud Connect service
        /// </summary>
        public readonly string Uuid;
        /// <summary>
        /// vrack linked to the service
        /// </summary>
        public readonly string Vrack;

        [OutputConstructor]
        private ConnectsOccResult(
            string bandwidth,

            string description,

            Outputs.ConnectsOccIamResult iam,

            ImmutableArray<double> interfaceLists,

            string pop,

            string portQuantity,

            string product,

            string providerName,

            string serviceName,

            string status,

            string uuid,

            string vrack)
        {
            Bandwidth = bandwidth;
            Description = description;
            Iam = iam;
            InterfaceLists = interfaceLists;
            Pop = pop;
            PortQuantity = portQuantity;
            Product = product;
            ProviderName = providerName;
            ServiceName = serviceName;
            Status = status;
            Uuid = uuid;
            Vrack = vrack;
        }
    }
}
