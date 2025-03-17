// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class ProjectOrderGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// date
        /// </summary>
        [Input("date")]
        public Input<string>? Date { get; set; }

        [Input("details")]
        private InputList<Inputs.ProjectOrderDetailGetArgs>? _details;

        /// <summary>
        /// Information about a Bill entry
        /// </summary>
        public InputList<Inputs.ProjectOrderDetailGetArgs> Details
        {
            get => _details ?? (_details = new InputList<Inputs.ProjectOrderDetailGetArgs>());
            set => _details = value;
        }

        /// <summary>
        /// expiration date
        /// </summary>
        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        /// <summary>
        /// order id, the same as the `id`
        /// </summary>
        [Input("orderId")]
        public Input<int>? OrderId { get; set; }

        public ProjectOrderGetArgs()
        {
        }
        public static new ProjectOrderGetArgs Empty => new ProjectOrderGetArgs();
    }
}
