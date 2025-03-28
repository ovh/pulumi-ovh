// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Ovh.Vrack
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
    ///     var myAccount = Ovh.Me.GetMe.Invoke();
    /// 
    ///     var myCart = Ovh.Order.GetCart.Invoke(new()
    ///     {
    ///         OvhSubsidiary = myAccount.Apply(getMeResult =&gt; getMeResult.OvhSubsidiary),
    ///     });
    /// 
    ///     var vrackCartProductPlan = Ovh.Order.GetCartProductPlan.Invoke(new()
    ///     {
    ///         CartId = myCart.Apply(getCartResult =&gt; getCartResult.Id),
    ///         PriceCapacity = "renew",
    ///         Product = "vrack",
    ///         PlanCode = "vrack",
    ///     });
    /// 
    ///     var vrackVrack = new Ovh.Vrack.Vrack("vrackVrack", new()
    ///     {
    ///         OvhSubsidiary = myCart.Apply(getCartResult =&gt; getCartResult.OvhSubsidiary),
    ///         Description = "my vrack",
    ///         Plan = new Ovh.Vrack.Inputs.VrackPlanArgs
    ///         {
    ///             Duration = vrackCartProductPlan.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.SelectedPrices[0]?.Duration),
    ///             PlanCode = vrackCartProductPlan.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.PlanCode),
    ///             PricingMode = vrackCartProductPlan.Apply(getCartProductPlanResult =&gt; getCartProductPlanResult.SelectedPrices[0]?.PricingMode),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A vRack can be imported using the `service_name`. Using the following configuration:
    /// 
    /// terraform
    /// 
    /// import {
    /// 
    ///   to = ovh_vrack.vrack
    /// 
    ///   id = "&lt;service name&gt;"
    /// 
    /// }
    /// 
    /// You can then run:
    /// 
    /// bash
    /// 
    /// $ pulumi preview -generate-config-out=vrack.tf
    /// 
    /// $ pulumi up
    /// 
    /// The file `vrack.tf` will then contain the imported resource's configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
    /// </summary>
    [OvhResourceType("ovh:Vrack/vrack:Vrack")]
    public partial class Vrack : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The URN of the vrack, used with IAM permissions
        /// </summary>
        [Output("VrackURN")]
        public Output<string> VrackURN { get; private set; } = null!;

        /// <summary>
        /// yourvrackdescription
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// yourvrackname
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Details about an Order
        /// </summary>
        [Output("orders")]
        public Output<ImmutableArray<Outputs.VrackOrder>> Orders { get; private set; } = null!;

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
        public Output<Outputs.VrackPlan> Plan { get; private set; } = null!;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        [Output("planOptions")]
        public Output<ImmutableArray<Outputs.VrackPlanOption>> PlanOptions { get; private set; } = null!;

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Output("serviceName")]
        public Output<string> ServiceName { get; private set; } = null!;


        /// <summary>
        /// Create a Vrack resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Vrack(string name, VrackArgs? args = null, CustomResourceOptions? options = null)
            : base("ovh:Vrack/vrack:Vrack", name, args ?? new VrackArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Vrack(string name, Input<string> id, VrackState? state = null, CustomResourceOptions? options = null)
            : base("ovh:Vrack/vrack:Vrack", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Vrack resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Vrack Get(string name, Input<string> id, VrackState? state = null, CustomResourceOptions? options = null)
        {
            return new Vrack(name, id, state, options);
        }
    }

    public sealed class VrackArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// yourvrackdescription
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// yourvrackname
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("orders")]
        private InputList<Inputs.VrackOrderArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.VrackOrderArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.VrackOrderArgs>());
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
        public Input<Inputs.VrackPlanArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.VrackPlanOptionArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.VrackPlanOptionArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.VrackPlanOptionArgs>());
            set => _planOptions = value;
        }

        public VrackArgs()
        {
        }
        public static new VrackArgs Empty => new VrackArgs();
    }

    public sealed class VrackState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URN of the vrack, used with IAM permissions
        /// </summary>
        [Input("VrackURN")]
        public Input<string>? VrackURN { get; set; }

        /// <summary>
        /// yourvrackdescription
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// yourvrackname
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("orders")]
        private InputList<Inputs.VrackOrderGetArgs>? _orders;

        /// <summary>
        /// Details about an Order
        /// </summary>
        public InputList<Inputs.VrackOrderGetArgs> Orders
        {
            get => _orders ?? (_orders = new InputList<Inputs.VrackOrderGetArgs>());
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
        public Input<Inputs.VrackPlanGetArgs>? Plan { get; set; }

        [Input("planOptions")]
        private InputList<Inputs.VrackPlanOptionGetArgs>? _planOptions;

        /// <summary>
        /// Product Plan to order
        /// </summary>
        public InputList<Inputs.VrackPlanOptionGetArgs> PlanOptions
        {
            get => _planOptions ?? (_planOptions = new InputList<Inputs.VrackPlanOptionGetArgs>());
            set => _planOptions = value;
        }

        /// <summary>
        /// The internal name of your vrack
        /// </summary>
        [Input("serviceName")]
        public Input<string>? ServiceName { get; set; }

        public VrackState()
        {
        }
        public static new VrackState Empty => new VrackState();
    }
}
