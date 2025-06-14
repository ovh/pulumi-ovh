// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    /// <summary>
    /// Creates a network in a public cloud project.
    /// </summary>
    [OvhResourceType("ovh:CloudProject/regionNetwork:RegionNetwork")]
    public partial class RegionNetwork : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the network
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network region returned by the API
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Network region
        /// </summary>
        [Output("regionName")]
        public Output<string> RegionName { get; private set; } = null!;

        /// <summary>
        /// The id of the public cloud project
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Parameters to create a subnet
        /// </summary>
        [Output("subnet")]
        public Output<Outputs.RegionNetworkSubnet> Subnet { get; private set; } = null!;

        /// <summary>
        /// Network visibility
        /// </summary>
        [Output("visibility")]
        public Output<string> Visibility { get; private set; } = null!;

        /// <summary>
        /// VLAN ID, between 1 and 4000
        /// </summary>
        [Output("vlanId")]
        public Output<double> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a RegionNetwork resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegionNetwork(string name, RegionNetworkArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/regionNetwork:RegionNetwork", name, args ?? new RegionNetworkArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegionNetwork(string name, Input<string> id, RegionNetworkState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/regionNetwork:RegionNetwork", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegionNetwork resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegionNetwork Get(string name, Input<string> id, RegionNetworkState? state = null, CustomResourceOptions? options = null)
        {
            return new RegionNetwork(name, id, state, options);
        }
    }

    public sealed class RegionNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network region
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        /// <summary>
        /// The id of the public cloud project
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Parameters to create a subnet
        /// </summary>
        [Input("subnet", required: true)]
        public Input<Inputs.RegionNetworkSubnetArgs> Subnet { get; set; } = null!;

        /// <summary>
        /// VLAN ID, between 1 and 4000
        /// </summary>
        [Input("vlanId")]
        public Input<double>? VlanId { get; set; }

        public RegionNetworkArgs()
        {
        }
        public static new RegionNetworkArgs Empty => new RegionNetworkArgs();
    }

    public sealed class RegionNetworkState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network region returned by the API
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Network region
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        /// <summary>
        /// The id of the public cloud project
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Parameters to create a subnet
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.RegionNetworkSubnetGetArgs>? Subnet { get; set; }

        /// <summary>
        /// Network visibility
        /// </summary>
        [Input("visibility")]
        public Input<string>? Visibility { get; set; }

        /// <summary>
        /// VLAN ID, between 1 and 4000
        /// </summary>
        [Input("vlanId")]
        public Input<double>? VlanId { get; set; }

        public RegionNetworkState()
        {
        }
        public static new RegionNetworkState Empty => new RegionNetworkState();
    }
}
