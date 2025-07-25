// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Dedicated.Inputs
{

    public sealed class ServerOrderArgs : global::Pulumi.ResourceArgs
    {
        [Input("date")]
        public Input<string>? Date { get; set; }

        [Input("details")]
        private InputList<Inputs.ServerOrderDetailArgs>? _details;
        public InputList<Inputs.ServerOrderDetailArgs> Details
        {
            get => _details ?? (_details = new InputList<Inputs.ServerOrderDetailArgs>());
            set => _details = value;
        }

        [Input("expirationDate")]
        public Input<string>? ExpirationDate { get; set; }

        [Input("orderId")]
        public Input<double>? OrderId { get; set; }

        public ServerOrderArgs()
        {
        }
        public static new ServerOrderArgs Empty => new ServerOrderArgs();
    }
}
