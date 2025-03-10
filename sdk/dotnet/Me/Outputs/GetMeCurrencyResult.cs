// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Me.Outputs
{

    [OutputType]
    public sealed class GetMeCurrencyResult
    {
        /// <summary>
        /// Currency code
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// Currency symbol
        /// </summary>
        public readonly string Symbol;

        [OutputConstructor]
        private GetMeCurrencyResult(
            string code,

            string symbol)
        {
            Code = code;
            Symbol = symbol;
        }
    }
}
