// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Inputs
{

    public sealed class VpsOrderDetailGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Product type of item in order
        /// </summary>
        [Input("detailType")]
        public Input<string>? DetailType { get; set; }

        [Input("domain")]
        public Input<string>? Domain { get; set; }

        [Input("orderDetailId")]
        public Input<double>? OrderDetailId { get; set; }

        [Input("quantity")]
        public Input<string>? Quantity { get; set; }

        public VpsOrderDetailGetArgs()
        {
        }
        public static new VpsOrderDetailGetArgs Empty => new VpsOrderDetailGetArgs();
    }
}