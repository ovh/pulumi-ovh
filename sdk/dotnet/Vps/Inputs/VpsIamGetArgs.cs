// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vps.Inputs
{

    public sealed class VpsIamGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Custom display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Unique identifier of the resource in the IAM
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags. Tags that were internally computed are prefixed with `ovh:`
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// URN of the private database, used when writing IAM policies
        /// </summary>
        [Input("urn")]
        public Input<string>? Urn { get; set; }

        public VpsIamGetArgs()
        {
        }
        public static new VpsIamGetArgs Empty => new VpsIamGetArgs();
    }
}
