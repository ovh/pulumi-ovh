// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Domain.Inputs
{

    public sealed class NameServersServerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The server hostname
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The server IP
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        public NameServersServerArgs()
        {
        }
        public static new NameServersServerArgs Empty => new NameServersServerArgs();
    }
}
