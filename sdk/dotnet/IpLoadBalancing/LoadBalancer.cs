// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.IpLoadBalancing
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Ovh = Pulumi.Ovh;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var myaccount = Ovh.Me.GetMe.Invoke();
    /// 
    ///     var mycart = Ovh.Order.GetCart.Invoke(new()
    ///     {
    ///         OvhSubsidiary = myaccount.Apply(getMeResult =&gt; getMeResult.OvhSubsidiary),
    ///     });
    /// 
    ///     var iplb = Ovh.Order.GetCartProductPlan.Invoke(new()
    ///     {
    ///         CartId = mycart.Apply(getCartResult =&gt; getCartResult.Id),
    ///         PriceCapacity = "renew",
    ///         Product = "ipLoadbalancing",
    ///         PlanCode = "iplb-lb1",
    ///     });
    /// 
    ///     var bhs = Ovh.Order.GetCartProductOptionsPlan.Invoke(new()
    ///     {
    ///         CartId = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.CartId),
    ///         PriceCapacity = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.PriceCapacity),
    ///         Product = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.Product),
    ///         PlanCode = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.PlanCode),
    ///         OptionsPlanCode = "iplb-zone-lb1-rbx",
    ///     });
    /// 
    ///     var iplb_lb1 = new Ovh.IpLoadBalancing.LoadBalancer("iplb-lb1", new()
    ///     {
    ///         OvhSubsidiary = mycart.Apply(getCartResult =&gt; getCartResult.OvhSubsidiary),
    ///         DisplayName = "my ip loadbalancing",
    ///         Plan = new Ovh.IpLoadBalancing.Inputs.LoadBalancerPlanArgs
    ///         {
    ///             Duration = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.SelectedPrices[0]?.Duration),
    ///             PlanCode = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.PlanCode),
    ///             PricingMode = iplb.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.SelectedPrices[0]?.PricingMode),
    ///         },
    ///         PlanOptions = new[]
    ///         {
    ///             new Ovh.IpLoadBalancing.Inputs.LoadBalancerPlanOptionArgs
    ///             {
    ///                 Duration = bhs.Apply(getCartProductOptionsPlanResult =&gt; getCartProductOptionsPlanResult.SelectedPrices[0]?.Duration),
    ///                 PlanCode = bhs.Apply(getCartProductOptionsPlanResult =&gt; getCartProductOptionsPlanResult.PlanCode),
    ///                 PricingMode = bhs.Apply(getCartProductOptionsPlanResult =&gt; getCartProductOptionsPlanResult.SelectedPrices[0]?.PricingMode),
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// OVHcloud IP load balancing services can be imported using its `service_name`.
    /// 
    /// Using the following configuration:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = ovh_iploadbalancing.iplb
    /// 
    ///   id = "&lt;service name&gt;"
    /// 
    /// }
    /// 
    /// You can then run:
    /// 
    /// bash
    /// 
    /// $ pulumi preview -generate-config-out=iplb.tf
    /// 
    /// $ pulumi up
    /// 
    /// The file `iplb.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
    /// </summary>
    [OvhResourceType("ovh:IpLoadBalancing/loadBalancer:LoadBalancer")]
    public partial class LoadBalancer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// URN of the load balancer, used when writing IAM policies
        /// </summary>
        [Output("LoadBalancerURN")]
        public Output<string> LoadBalancerURN { get; private set; } = null!;

        /// <summary>
        /// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// Your IP load balancing
        /// </summary>
        [Output("ipLoadbalancing")]
        public Output<string> IpLoadbalancing { get; private set; } = null!;

        /// <summary>
        /// The IPV4 associated to your IP load balancing
        /// </summary>
        [Output("ipv4")]
        public Output<string> Ipv4 { get; private set; } = null!;

        /// <summary>
        /// The IPV6 associated to your IP load balancing. DEPRECATED.
        /// </summary>
        [Output("ipv6")]
        public Output<string> Ipv6 { get; private set; } = null!;

        /// <summary>
        /// The metrics token associated with your IP load balancing
        /// </summary>
        [Output("metricsToken")]
        public Output<string> MetricsToken { get; private set; } = null!;

        /// <summary>
        /// The offer of your IP load balancing
        /// </summary>
        [Output("offer")]
        public Output<string> Offer { get; private set; } = null!;

        /// <summary>
        /// Available additional zone for your Load Balancer
        /// </summary>
        [Output("orderableZones")]
        public Output<ImmutableArray<Outputs.LoadBalancerOrderableZone>> OrderableZones { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.LoadBalancerOrder>> Orders { get; private set; } = null!;

        /// <summary>
        /// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        /// </summary>
        [Output("ovhSubsidiary")]
        public Output<string> OvhSubsidiary { get; private set; } = null!;

        /// <summary>
        /// Ovh payment mode
        /// </summary>
        [Output("paymentMean")]
        public Output<string?> PaymentMean { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("plan")]
        public Output<Outputs.LoadBalancerPlan> Plan { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.LoadBalancerPlanOption>> PlanOptions { get; private set; } = null!;

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;

        /// <summary>
        /// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
        /// </summary>
        [Output("sslConfiguration")]
        public Output<string> SslConfiguration { get; private set; } = null!;

        /// <summary>
        /// Current state of your IP
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Vrack eligibility
        /// </summary>
        [Output("vrackEligibility")]
        public Output<bool> VrackEligibility { get; private set; } = null!;

        /// <summary>
        /// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
        /// </summary>
        [Output("vrackName")]
        public Output<string> VrackName { get; private set; } = null!;

        /// <summary>
        /// Location where your service is
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a LoadBalancer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public LoadBalancer(string name, LoadBalancerArgs? args = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/loadBalancer:LoadBalancer", name, args ?? new LoadBalancerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private LoadBalancer(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
            : base("ovh:IpLoadBalancing/loadBalancer:LoadBalancer", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/ovh/pulumi-ovh",
                AdditionalSecretOutputs =
                {
                    "metricsToken",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing LoadBalancer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static LoadBalancer Get(string name, Input<string> id, LoadBalancerState? state = null, CustomResourceOptions? options = null)
        {
            return new LoadBalancer(name, id, state, options);
        }
    }

    public sealed class LoadBalancerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("orders")]
        private InputList<Inputs.LoadBalancerOrderArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.LoadBalancerOrderArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.LoadBalancerOrderArgs>());
            set => _orders = value;
        }

        /// <summary>
        /// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Ovh payment mode
        /// </summary>
        [Input("paymentMean")]
        public Input<string>? PaymentMean { get; set; }

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan")]
        public Input<Inputs.LoadBalancerPlanArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.LoadBalancerPlanOptionArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.LoadBalancerPlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.LoadBalancerPlanOptionArgs>());
            set => _planOptions = value;
        }

        /// <summary>
        /// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
        /// </summary>
        [Input("sslConfiguration")]
        public Input<string>? SslConfiguration { get; set; }

        public LoadBalancerArgs()
        {
        }
        public static new LoadBalancerArgs Empty => new LoadBalancerArgs();
    }

    public sealed class LoadBalancerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// URN of the load balancer, used when writing IAM policies
        /// </summary>
        [Input("LoadBalancerURN")]
        public Input<string>? LoadBalancerURN { get; set; }

        /// <summary>
        /// Set the name displayed in ManagerV6 for your iplb (max 50 chars)
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        /// <summary>
        /// Your IP load balancing
        /// </summary>
        [Input("ipLoadbalancing")]
        public Input<string>? IpLoadbalancing { get; set; }

        /// <summary>
        /// The IPV4 associated to your IP load balancing
        /// </summary>
        [Input("ipv4")]
        public Input<string>? Ipv4 { get; set; }

        /// <summary>
        /// The IPV6 associated to your IP load balancing. DEPRECATED.
        /// </summary>
        [Input("ipv6")]
        public Input<string>? Ipv6 { get; set; }

        [Input("metricsToken")]
        private Input<string>? _metricsToken;

        /// <summary>
        /// The metrics token associated with your IP load balancing
        /// </summary>
        public Input<string>? MetricsToken
        {
            get => _metricsToken;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _metricsToken = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The offer of your IP load balancing
        /// </summary>
        [Input("offer")]
        public Input<string>? Offer { get; set; }

        [Input("orderableZones")]
        private InputList<Inputs.LoadBalancerOrderableZoneGetArgs>? _orderableZones;

        /// <summary>
        /// Available additional zone for your Load Balancer
        /// </summary>
        public InputList<Inputs.LoadBalancerOrderableZoneGetArgs> OrderableZones
        {
            get => _orderableZones ?? (_orderableZones = new InputList<Inputs.LoadBalancerOrderableZoneGetArgs>());
            set => _orderableZones = value;
        }

        [Input("orders")]
        private InputList<Inputs.LoadBalancerOrderGetArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.LoadBalancerOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.LoadBalancerOrderGetArgs>());
            set => _orders = value;
        }

        /// <summary>
        /// OVHcloud Subsidiary. Country of OVHcloud legal entity you'll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
        /// </summary>
        [Input("ovhSubsidiary")]
        public Input<string>? OvhSubsidiary { get; set; }

        /// <summary>
        /// Ovh payment mode
        /// </summary>
        [Input("paymentMean")]
        public Input<string>? PaymentMean { get; set; }

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Input("plan")]
        public Input<Inputs.LoadBalancerPlanGetArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.LoadBalancerPlanOptionGetArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.LoadBalancerPlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.LoadBalancerPlanOptionGetArgs>());
            set => _planOptions = value;
        }

        /// <summary>
        /// The internal name of your IP load balancing
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        /// <summary>
        /// Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of "intermediate", "modern".
        /// </summary>
        [Input("sslConfiguration")]
        public Input<string>? SslConfiguration { get; set; }

        /// <summary>
        /// Current state of your IP
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Vrack eligibility
        /// </summary>
        [Input("vrackEligibility")]
        public Input<bool>? VrackEligibility { get; set; }

        /// <summary>
        /// Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
        /// </summary>
        [Input("vrackName")]
        public Input<string>? VrackName { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// Location where your service is
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public LoadBalancerState()
        {
        }
        public static new LoadBalancerState Empty => new LoadBalancerState();
    }
}
