// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class NetworkPrivateSubnetV2AllocationPoolGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("end", required: true)]
        public Input<string> End { get; set; } = null!;

        [Input("start", required: true)]
        public Input<string> Start { get; set; } = null!;

        public NetworkPrivateSubnetV2AllocationPoolGetArgs()
        {
        }
        public static new NetworkPrivateSubnetV2AllocationPoolGetArgs Empty => new NetworkPrivateSubnetV2AllocationPoolGetArgs();
    }
}
