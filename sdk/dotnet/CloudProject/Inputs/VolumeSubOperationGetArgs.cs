// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class VolumeSubOperationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Affected resource of the sub-operation
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// The started date of the sub-operation
        /// </summary>
        [Input("resourceType")]
        public Input<string>? ResourceType { get; set; }

        public VolumeSubOperationGetArgs()
        {
        }
        public static new VolumeSubOperationGetArgs Empty => new VolumeSubOperationGetArgs();
    }
}
