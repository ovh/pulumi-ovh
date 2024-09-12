// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.IpLoadBalancing;

import com.ovhcloud.ovh.IpLoadBalancing.LoadBalancerArgs;
import com.ovhcloud.ovh.IpLoadBalancing.inputs.LoadBalancerState;
import com.ovhcloud.ovh.IpLoadBalancing.outputs.LoadBalancerOrder;
import com.ovhcloud.ovh.IpLoadBalancing.outputs.LoadBalancerOrderableZone;
import com.ovhcloud.ovh.IpLoadBalancing.outputs.LoadBalancerPlan;
import com.ovhcloud.ovh.IpLoadBalancing.outputs.LoadBalancerPlanOption;
import com.ovhcloud.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.Me.MeFunctions;
 * import com.pulumi.ovh.Order.OrderFunctions;
 * import com.pulumi.ovh.Order.inputs.GetCartArgs;
 * import com.pulumi.ovh.Order.inputs.GetCartProductPlanArgs;
 * import com.pulumi.ovh.Order.inputs.GetCartProductOptionsPlanArgs;
 * import com.pulumi.ovh.IpLoadBalancing.LoadBalancer;
 * import com.pulumi.ovh.IpLoadBalancing.LoadBalancerArgs;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.LoadBalancerPlanArgs;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.LoadBalancerPlanOptionArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var myaccount = MeFunctions.getMe();
 * 
 *         final var mycart = OrderFunctions.getCart(GetCartArgs.builder()
 *             .ovhSubsidiary(myaccount.applyValue(getMeResult -> getMeResult.ovhSubsidiary()))
 *             .build());
 * 
 *         final var iplb = OrderFunctions.getCartProductPlan(GetCartProductPlanArgs.builder()
 *             .cartId(mycart.applyValue(getCartResult -> getCartResult.id()))
 *             .priceCapacity("renew")
 *             .product("ipLoadbalancing")
 *             .planCode("iplb-lb1")
 *             .build());
 * 
 *         final var bhs = OrderFunctions.getCartProductOptionsPlan(GetCartProductOptionsPlanArgs.builder()
 *             .cartId(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.cartId()))
 *             .priceCapacity(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.priceCapacity()))
 *             .product(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.product()))
 *             .planCode(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.planCode()))
 *             .optionsPlanCode("iplb-zone-lb1-rbx")
 *             .build());
 * 
 *         var iplb_lb1 = new LoadBalancer("iplb-lb1", LoadBalancerArgs.builder()
 *             .ovhSubsidiary(mycart.applyValue(getCartResult -> getCartResult.ovhSubsidiary()))
 *             .displayName("my ip loadbalancing")
 *             .plan(LoadBalancerPlanArgs.builder()
 *                 .duration(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.planCode()))
 *                 .pricingMode(iplb.applyValue(getCartProductPlanResult -> getCartProductPlanResult.selectedPrices()[0].pricingMode()))
 *                 .build())
 *             .planOptions(LoadBalancerPlanOptionArgs.builder()
 *                 .duration(bhs.applyValue(getCartProductOptionsPlanResult -> getCartProductOptionsPlanResult.selectedPrices()[0].duration()))
 *                 .planCode(bhs.applyValue(getCartProductOptionsPlanResult -> getCartProductOptionsPlanResult.planCode()))
 *                 .pricingMode(bhs.applyValue(getCartProductOptionsPlanResult -> getCartProductOptionsPlanResult.selectedPrices()[0].pricingMode()))
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/loadBalancer:LoadBalancer")
public class LoadBalancer extends com.pulumi.resources.CustomResource {
    /**
     * URN of the load balancer, used when writing IAM policies
     * 
     */
    @Export(name="LoadBalancerURN", refs={String.class}, tree="[0]")
    private Output<String> LoadBalancerURN;

    /**
     * @return URN of the load balancer, used when writing IAM policies
     * 
     */
    public Output<String> LoadBalancerURN() {
        return this.LoadBalancerURN;
    }
    /**
     * Set the name displayed in ManagerV6 for your iplb (max 50 chars)
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return Set the name displayed in ManagerV6 for your iplb (max 50 chars)
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Your IP load balancing
     * 
     */
    @Export(name="ipLoadbalancing", refs={String.class}, tree="[0]")
    private Output<String> ipLoadbalancing;

    /**
     * @return Your IP load balancing
     * 
     */
    public Output<String> ipLoadbalancing() {
        return this.ipLoadbalancing;
    }
    /**
     * The IPV4 associated to your IP load balancing
     * 
     */
    @Export(name="ipv4", refs={String.class}, tree="[0]")
    private Output<String> ipv4;

    /**
     * @return The IPV4 associated to your IP load balancing
     * 
     */
    public Output<String> ipv4() {
        return this.ipv4;
    }
    /**
     * The IPV6 associated to your IP load balancing. DEPRECATED.
     * 
     */
    @Export(name="ipv6", refs={String.class}, tree="[0]")
    private Output<String> ipv6;

    /**
     * @return The IPV6 associated to your IP load balancing. DEPRECATED.
     * 
     */
    public Output<String> ipv6() {
        return this.ipv6;
    }
    /**
     * The metrics token associated with your IP load balancing
     * 
     */
    @Export(name="metricsToken", refs={String.class}, tree="[0]")
    private Output<String> metricsToken;

    /**
     * @return The metrics token associated with your IP load balancing
     * 
     */
    public Output<String> metricsToken() {
        return this.metricsToken;
    }
    /**
     * The offer of your IP load balancing
     * 
     */
    @Export(name="offer", refs={String.class}, tree="[0]")
    private Output<String> offer;

    /**
     * @return The offer of your IP load balancing
     * 
     */
    public Output<String> offer() {
        return this.offer;
    }
    /**
     * Available additional zone for your Load Balancer
     * 
     */
    @Export(name="orderableZones", refs={List.class,LoadBalancerOrderableZone.class}, tree="[0,1]")
    private Output<List<LoadBalancerOrderableZone>> orderableZones;

    /**
     * @return Available additional zone for your Load Balancer
     * 
     */
    public Output<List<LoadBalancerOrderableZone>> orderableZones() {
        return this.orderableZones;
    }
    /**
     * Details about an Order
     * 
     */
    @Export(name="orders", refs={List.class,LoadBalancerOrder.class}, tree="[0,1]")
    private Output<List<LoadBalancerOrder>> orders;

    /**
     * @return Details about an Order
     * 
     */
    public Output<List<LoadBalancerOrder>> orders() {
        return this.orders;
    }
    /**
     * OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    @Export(name="ovhSubsidiary", refs={String.class}, tree="[0]")
    private Output<String> ovhSubsidiary;

    /**
     * @return OVHcloud Subsidiary. Country of OVHcloud legal entity you&#39;ll be billed by. List of supported subsidiaries available on API at [/1.0/me.json under `models.nichandle.OvhSubsidiaryEnum`](https://eu.api.ovh.com/1.0/me.json)
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * Ovh payment mode
     * 
     * @deprecated
     * This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used.
     * 
     */
    @Deprecated /* This field is not anymore used since the API has been deprecated in favor of /payment/mean. Now, the default payment mean is used. */
    @Export(name="paymentMean", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> paymentMean;

    /**
     * @return Ovh payment mode
     * 
     */
    public Output<Optional<String>> paymentMean() {
        return Codegen.optional(this.paymentMean);
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="plan", refs={LoadBalancerPlan.class}, tree="[0]")
    private Output<LoadBalancerPlan> plan;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<LoadBalancerPlan> plan() {
        return this.plan;
    }
    /**
     * Product Plan to order
     * 
     */
    @Export(name="planOptions", refs={List.class,LoadBalancerPlanOption.class}, tree="[0,1]")
    private Output</* @Nullable */ List<LoadBalancerPlanOption>> planOptions;

    /**
     * @return Product Plan to order
     * 
     */
    public Output<Optional<List<LoadBalancerPlanOption>>> planOptions() {
        return Codegen.optional(this.planOptions);
    }
    /**
     * The internal name of your IP load balancing
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your IP load balancing
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of &#34;intermediate&#34;, &#34;modern&#34;.
     * 
     */
    @Export(name="sslConfiguration", refs={String.class}, tree="[0]")
    private Output<String> sslConfiguration;

    /**
     * @return Modern oldest compatible clients : Firefox 27, Chrome 30, IE 11 on Windows 7, Edge, Opera 17, Safari 9, Android 5.0, and Java 8. Intermediate oldest compatible clients : Firefox 1, Chrome 1, IE 7, Opera 5, Safari 1, Windows XP IE8, Android 2.3, Java 7. Intermediate if null. one of &#34;intermediate&#34;, &#34;modern&#34;.
     * 
     */
    public Output<String> sslConfiguration() {
        return this.sslConfiguration;
    }
    /**
     * Current state of your IP
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Current state of your IP
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Vrack eligibility
     * 
     */
    @Export(name="vrackEligibility", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> vrackEligibility;

    /**
     * @return Vrack eligibility
     * 
     */
    public Output<Boolean> vrackEligibility() {
        return this.vrackEligibility;
    }
    /**
     * Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
     * 
     */
    @Export(name="vrackName", refs={String.class}, tree="[0]")
    private Output<String> vrackName;

    /**
     * @return Name of the vRack on which the current Load Balancer is attached to, as it is named on vRack product
     * 
     */
    public Output<String> vrackName() {
        return this.vrackName;
    }
    /**
     * Location where your service is
     * 
     */
    @Export(name="zones", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> zones;

    /**
     * @return Location where your service is
     * 
     */
    public Output<List<String>> zones() {
        return this.zones;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LoadBalancer(java.lang.String name) {
        this(name, LoadBalancerArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LoadBalancer(java.lang.String name, LoadBalancerArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LoadBalancer(java.lang.String name, LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/loadBalancer:LoadBalancer", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LoadBalancer(java.lang.String name, Output<java.lang.String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/loadBalancer:LoadBalancer", name, state, makeResourceOptions(options, id), false);
    }

    private static LoadBalancerArgs makeArgs(LoadBalancerArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LoadBalancerArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "metricsToken"
            ))
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static LoadBalancer get(java.lang.String name, Output<java.lang.String> id, @Nullable LoadBalancerState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LoadBalancer(name, id, state, options);
    }
}