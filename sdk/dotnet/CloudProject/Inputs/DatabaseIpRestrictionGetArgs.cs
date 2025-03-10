// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class DatabaseIpRestrictionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the IP restriction
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Authorized IP
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// Current status of the IP restriction
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public DatabaseIpRestrictionGetArgs()
        {
        }
        public static new DatabaseIpRestrictionGetArgs Empty => new DatabaseIpRestrictionGetArgs();
    }
}
