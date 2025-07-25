// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProjectDatabase.Inputs
{

    public sealed class OpensearchUserAclArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pattern of the ACL.
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        /// <summary>
        /// Permission of the ACL Available permission:
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        public OpensearchUserAclArgs()
        {
        }
        public static new OpensearchUserAclArgs Empty => new OpensearchUserAclArgs();
    }
}
