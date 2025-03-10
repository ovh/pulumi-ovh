// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.CloudProject
{
    [OvhResourceType("ovh:CloudProject/loadbalancer:Loadbalancer")]
    public partial class Loadbalancer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The UTC date and timestamp when the resource was created
        /// </summary>
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        /// <summary>
        /// Description of the loadbalancer
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Loadbalancer flavor id
        /// </summary>
        [Output("flavorId")]
        public Output<string> FlavorId { get; private set; } = null!;

        /// <summary>
        /// Information about floating IP
        /// </summary>
        [Output("floatingIp")]
        public Output<Outputs.LoadbalancerFloatingIp> FloatingIp { get; private set; } = null!;

        /// <summary>
        /// Listeners to create with the loadbalancer
        /// </summary>
        [Output("listeners")]
        public Output<ImmutableArray<Outputs.LoadbalancerListener>> Listeners { get; private set; } = null!;

        /// <summary>
        /// Name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network information to create the loadbalancer
        /// </summary>
        [Output("network")]
        public Output<Outputs.LoadbalancerNetwork> Network { get; private set; } = null!;

        /// <summary>
        /// Operating status of the resource
        /// </summary>
        [Output("operatingStatus")]
        public Output<string> OperatingStatus { get; private set; } = null!;

        /// <summary>
        /// Provisioning status of the resource
        /// </summary>
        [Output("provisioningStatus")]
        public Output<string> ProvisioningStatus { get; private set; } = null!;

        /// <summary>
        /// Region of the resource
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// Region name
        /// </summary>
        [Output("regionName")]
        public Output<string> RegionName { get; private set; } = null!;

        /// <summary>
        /// Service name
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// UTC date and timestamp when the resource was created
        /// </summary>
        [Output("updatedAt")]
        public Output<string> UpdatedAt { get; private set; } = null!;

        /// <summary>
        /// IP address of the Virtual IP
        /// </summary>
        [Output("vipAddress")]
        public Output<string> VipAddress { get; private set; } = null!;

        /// <summary>
        /// Openstack ID of the network for the Virtual IP
        /// </summary>
        [Output("vipNetworkId")]
        public Output<string> VipNetworkId { get; private set; } = null!;

        /// <summary>
        /// ID of the subnet for the Virtual IP
        /// </summary>
        [Output("vipSubnetId")]
        public Output<string> VipSubnetId { get; private set; } = null!;


        /// <summary>
        /// Create a Loadbalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Loadbalancer(string name, LoadbalancerArgs args, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/loadbalancer:Loadbalancer", name, args ?? new LoadbalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Loadbalancer(string name, Input<string> id, LoadbalancerState? state = null, CustomResourceOptions? options = null)
            : base("ovh:CloudProject/loadbalancer:Loadbalancer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Loadbalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Loadbalancer Get(string name, Input<string> id, LoadbalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new Loadbalancer(name, id, state, options);
        }
    }

    public sealed class LoadbalancerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the loadbalancer
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Loadbalancer flavor id
        /// </summary>
        [Input("flavorId", required: true)]
        public Input<string> FlavorId { get; set; } = null!;

        [Input("listeners")]
        private InputList<Inputs.LoadbalancerListenerArgs>? _listeners;

        /// <summary>
        /// Listeners to create with the loadbalancer
        /// </summary>
        public InputList<Inputs.LoadbalancerListenerArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.LoadbalancerListenerArgs>());
            set => _listeners = value;
        }

        /// <summary>
        /// Name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network information to create the loadbalancer
        /// </summary>
        [Input("network", required: true)]
        public Input<Inputs.LoadbalancerNetworkArgs> Network { get; set; } = null!;

        /// <summary>
        /// Region name
        /// </summary>
        [Input("regionName", required: true)]
        public Input<string> RegionName { get; set; } = null!;

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public LoadbalancerArgs()
        {
        }
        public static new LoadbalancerArgs Empty => new LoadbalancerArgs();
    }

    public sealed class LoadbalancerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The UTC date and timestamp when the resource was created
        /// </summary>
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        /// <summary>
        /// Description of the loadbalancer
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Loadbalancer flavor id
        /// </summary>
        [Input("flavorId")]
        public Input<string>? FlavorId { get; set; }

        /// <summary>
        /// Information about floating IP
        /// </summary>
        [Input("floatingIp")]
        public Input<Inputs.LoadbalancerFloatingIpGetArgs>? FloatingIp { get; set; }

        [Input("listeners")]
        private InputList<Inputs.LoadbalancerListenerGetArgs>? _listeners;

        /// <summary>
        /// Listeners to create with the loadbalancer
        /// </summary>
        public InputList<Inputs.LoadbalancerListenerGetArgs> Listeners
        {
            get => _listeners ?? (_listeners = new InputList<Inputs.LoadbalancerListenerGetArgs>());
            set => _listeners = value;
        }

        /// <summary>
        /// Name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Network information to create the loadbalancer
        /// </summary>
        [Input("network")]
        public Input<Inputs.LoadbalancerNetworkGetArgs>? Network { get; set; }

        /// <summary>
        /// Operating status of the resource
        /// </summary>
        [Input("operatingStatus")]
        public Input<string>? OperatingStatus { get; set; }

        /// <summary>
        /// Provisioning status of the resource
        /// </summary>
        [Input("provisioningStatus")]
        public Input<string>? ProvisioningStatus { get; set; }

        /// <summary>
        /// Region of the resource
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        /// <summary>
        /// Region name
        /// </summary>
        [Input("regionName")]
        public Input<string>? RegionName { get; set; }

        /// <summary>
        /// Service name
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// UTC date and timestamp when the resource was created
        /// </summary>
        [Input("updatedAt")]
        public Input<string>? UpdatedAt { get; set; }

        /// <summary>
        /// IP address of the Virtual IP
        /// </summary>
        [Input("vipAddress")]
        public Input<string>? VipAddress { get; set; }

        /// <summary>
        /// Openstack ID of the network for the Virtual IP
        /// </summary>
        [Input("vipNetworkId")]
        public Input<string>? VipNetworkId { get; set; }

        /// <summary>
        /// ID of the subnet for the Virtual IP
        /// </summary>
        [Input("vipSubnetId")]
        public Input<string>? VipSubnetId { get; set; }

        public LoadbalancerState()
        {
        }
        public static new LoadbalancerState Empty => new LoadbalancerState();
    }
}
