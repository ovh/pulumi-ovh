// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Okms.Inputs
{

    public sealed class OkmsIamArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource display name
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Unique identifier of the resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags. Tags that were internally computed are prefixed with ovh:
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// Unique resource name used in policies
        /// </summary>
        [Input("urn")]
        public Input<string>? Urn { get; set; }

        public OkmsIamArgs()
        {
        }
        public static new OkmsIamArgs Empty => new OkmsIamArgs();
    }
}
