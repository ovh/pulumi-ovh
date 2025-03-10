// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Order
{
    public static class GetCart
    {
        public static Task<GetCartResult> InvokeAsync(GetCartArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCartResult>("ovh:Order/getCart:getCart", args ?? new GetCartArgs(), options.WithDefaults());

        public static Output<GetCartResult> Invoke(GetCartInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCartResult>("ovh:Order/getCart:getCart", args ?? new GetCartInvokeArgs(), options.WithDefaults());

        public static Output<GetCartResult> Invoke(GetCartInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCartResult>("ovh:Order/getCart:getCart", args ?? new GetCartInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCartArgs : global::Pulumi.InvokeArgs
    {
        [Input("assign")]
        public bool? Assign { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("expire")]
        public string? Expire { get; set; }

        [Input("ovhSubsidiary", required: true)]
        public string OvhSubsidiary { get; set; } = null!;

        public GetCartArgs()
        {
        }
        public static new GetCartArgs Empty => new GetCartArgs();
    }

    public sealed class GetCartInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("assign")]
        public Input<bool>? Assign { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("expire")]
        public Input<string>? Expire { get; set; }

        [Input("ovhSubsidiary", required: true)]
        public Input<string> OvhSubsidiary { get; set; } = null!;

        public GetCartInvokeArgs()
        {
        }
        public static new GetCartInvokeArgs Empty => new GetCartInvokeArgs();
    }


    [OutputType]
    public sealed class GetCartResult
    {
        public readonly bool? Assign;
        public readonly string CartId;
        public readonly string? Description;
        public readonly string Expire;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<int> Items;
        public readonly string OvhSubsidiary;
        public readonly bool ReadOnly;

        [OutputConstructor]
        private GetCartResult(
            bool? assign,

            string cartId,

            string? description,

            string expire,

            string id,

            ImmutableArray<int> items,

            string ovhSubsidiary,

            bool readOnly)
        {
            Assign = assign;
            CartId = cartId;
            Description = description;
            Expire = expire;
            Id = id;
            Items = items;
            OvhSubsidiary = ovhSubsidiary;
            ReadOnly = readOnly;
        }
    }
}
