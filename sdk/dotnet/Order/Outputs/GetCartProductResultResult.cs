// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Order.Outputs
{

    [OutputType]
    public sealed class GetCartProductResultResult
    {
        /// <summary>
        /// Product offer identifier
        /// </summary>
        public readonly string PlanCode;
        /// <summary>
        /// Prices of the product offer
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCartProductResultPriceResult> Prices;
        /// <summary>
        /// Name of the product
        /// </summary>
        public readonly string ProductName;
        /// <summary>
        /// Product type
        /// </summary>
        public readonly string ProductType;

        [OutputConstructor]
        private GetCartProductResultResult(
            string planCode,

            ImmutableArray<Outputs.GetCartProductResultPriceResult> prices,

            string productName,

            string productType)
        {
            PlanCode = planCode;
            Prices = prices;
            ProductName = productName;
            ProductType = productType;
        }
    }
}
