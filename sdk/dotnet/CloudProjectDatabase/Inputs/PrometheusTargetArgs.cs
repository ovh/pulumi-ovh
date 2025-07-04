// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase.Inputs
{

    public sealed class PrometheusTargetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Host of the endpoint
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// Connection port for the endpoint
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public PrometheusTargetArgs()
        {
        }
        public static new PrometheusTargetArgs Empty => new PrometheusTargetArgs();
    }
}
