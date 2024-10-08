// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Outputs
{

    [OutputType]
    public sealed class ServerOrder
    {
        public readonly string? Date;
        /// <summary>
        /// Details object when reinstalling server (see https://eu.api.ovh.com/console/?section=%2Fdedicated%2Fserver&amp;branch=v1#post-/dedicated/server/-serviceName-/install/start)
        /// </summary>
        public readonly ImmutableArray<Outputs.ServerOrderDetail> Details;
        public readonly string? ExpirationDate;
        public readonly double? OrderId;

        [OutputConstructor]
        private ServerOrder(
            string? date,

            ImmutableArray<Outputs.ServerOrderDetail> details,

            string? expirationDate,

            double? orderId)
        {
            Date = date;
            Details = details;
            ExpirationDate = expirationDate;
            OrderId = orderId;
        }
    }
}
