// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Outputs
{

    [OutputType]
    public sealed class InstanceAttachedVolume
    {
        /// <summary>
        /// Volume id
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private InstanceAttachedVolume(string? id)
        {
            Id = id;
        }
    }
}
