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
    public sealed class ProjectOrderDetail
    {
        /// <summary>
        /// A description associated with the user.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// expiration date
        /// </summary>
        public readonly string? Domain;
        /// <summary>
        /// order detail id
        /// </summary>
        public readonly int? OrderDetailId;
        /// <summary>
        /// quantity
        /// </summary>
        public readonly string? Quantity;

        [OutputConstructor]
        private ProjectOrderDetail(
            string? description,

            string? domain,

            int? orderDetailId,

            string? quantity)
        {
            Description = description;
            Domain = domain;
            OrderDetailId = orderDetailId;
            Quantity = quantity;
        }
    }
}
