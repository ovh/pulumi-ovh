// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Scraly.PulumiPackage.Ovh.CloudProjectDatabase.Inputs
{

    public sealed class OpensearchUserAclGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Pattern of the ACL.
        /// </summary>
        [Input("pattern", required: true)]
        public Input<string> Pattern { get; set; } = null!;

        /// <summary>
        /// Permission of the ACL
        /// Available permission:
        /// </summary>
        [Input("permission", required: true)]
        public Input<string> Permission { get; set; } = null!;

        public OpensearchUserAclGetArgs()
        {
        }
        public static new OpensearchUserAclGetArgs Empty => new OpensearchUserAclGetArgs();
    }
}