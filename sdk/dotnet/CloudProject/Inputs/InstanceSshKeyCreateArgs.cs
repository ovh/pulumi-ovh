// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class InstanceSshKeyCreateArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SSH Key pair name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// SSH Public Key
        /// </summary>
        [Input("publicKey", required: true)]
        public Input<string> PublicKey { get; set; } = null!;

        public InstanceSshKeyCreateArgs()
        {
        }
        public static new InstanceSshKeyCreateArgs Empty => new InstanceSshKeyCreateArgs();
    }
}
