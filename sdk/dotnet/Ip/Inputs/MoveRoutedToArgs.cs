// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Ip.Inputs
{

    public sealed class MoveRoutedToArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Service where ip is routed to
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public MoveRoutedToArgs()
        {
        }
        public static new MoveRoutedToArgs Empty => new MoveRoutedToArgs();
    }
}
