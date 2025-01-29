// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject.Inputs
{

    public sealed class InstanceFlavorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flavor ID. Flavors can be retrieved using `GET /cloud/project/{serviceName}/flavor`
        /// </summary>
        [Input("flavorId", required: true)]
        public Input<string> FlavorId { get; set; } = null!;

        public InstanceFlavorArgs()
        {
        }
        public static new InstanceFlavorArgs Empty => new InstanceFlavorArgs();
    }
}
