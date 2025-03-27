// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.RegionLoadBalancerLogSubscriptionArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RegionLoadBalancerLogSubscriptionState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Subscribe to a Managed Loadbalance Logs Service in a public cloud project.
 * 
 * ## Example Usage
 * 
 * Create a subscription
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.ovh.CloudProject.RegionLoadBalancerLogSubscription;
 * import com.pulumi.ovh.CloudProject.RegionLoadBalancerLogSubscriptionArgs;
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
 *         var subscription = new RegionLoadBalancerLogSubscription("subscription", RegionLoadBalancerLogSubscriptionArgs.builder()
 *             .kind("haproxy")
 *             .loadbalancerId("aaaaaaaa-bbbb-cccc-dddd-eeeeeeeeeeee")
 *             .regionName("yyyy")
 *             .serviceName("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
 *             .streamId("ffffffff-gggg-hhhh-iiii-jjjjjjjjjjjj")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:CloudProject/regionLoadBalancerLogSubscription:RegionLoadBalancerLogSubscription")
public class RegionLoadBalancerLogSubscription extends com.pulumi.resources.CustomResource {
    /**
     * The date of the subscription creation
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return The date of the subscription creation
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * haproxy **Changing this value recreates the resource.**
     * 
     */
    @Export(name="kind", refs={String.class}, tree="[0]")
    private Output<String> kind;

    /**
     * @return haproxy **Changing this value recreates the resource.**
     * 
     */
    public Output<String> kind() {
        return this.kind;
    }
    /**
     * LDP service name
     * 
     */
    @Export(name="ldpServiceName", refs={String.class}, tree="[0]")
    private Output<String> ldpServiceName;

    /**
     * @return LDP service name
     * 
     */
    public Output<String> ldpServiceName() {
        return this.ldpServiceName;
    }
    /**
     * Loadbalancer id to get the logs **Changing this value recreates the resource.**
     * 
     */
    @Export(name="loadbalancerId", refs={String.class}, tree="[0]")
    private Output<String> loadbalancerId;

    /**
     * @return Loadbalancer id to get the logs **Changing this value recreates the resource.**
     * 
     */
    public Output<String> loadbalancerId() {
        return this.loadbalancerId;
    }
    /**
     * The operation ID
     * 
     */
    @Export(name="operationId", refs={String.class}, tree="[0]")
    private Output<String> operationId;

    /**
     * @return The operation ID
     * 
     */
    public Output<String> operationId() {
        return this.operationId;
    }
    /**
     * A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: &#34;GRA11&#34;. **Changing this value recreates the resource.**
     * 
     */
    @Export(name="regionName", refs={String.class}, tree="[0]")
    private Output<String> regionName;

    /**
     * @return A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: &#34;GRA11&#34;. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }
    /**
     * The resource name
     * 
     */
    @Export(name="resourceName", refs={String.class}, tree="[0]")
    private Output<String> resourceName;

    /**
     * @return The resource name
     * 
     */
    public Output<String> resourceName() {
        return this.resourceName;
    }
    /**
     * The resource type
     * 
     */
    @Export(name="resourceType", refs={String.class}, tree="[0]")
    private Output<String> resourceType;

    /**
     * @return The resource type
     * 
     */
    public Output<String> resourceType() {
        return this.resourceType;
    }
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used. **Changing this value recreates the resource.**
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Data stream id to use for the subscription **Changing this value recreates the resource.**
     * 
     */
    @Export(name="streamId", refs={String.class}, tree="[0]")
    private Output<String> streamId;

    /**
     * @return Data stream id to use for the subscription **Changing this value recreates the resource.**
     * 
     */
    public Output<String> streamId() {
        return this.streamId;
    }
    /**
     * The subscription id
     * 
     */
    @Export(name="subscriptionId", refs={String.class}, tree="[0]")
    private Output<String> subscriptionId;

    /**
     * @return The subscription id
     * 
     */
    public Output<String> subscriptionId() {
        return this.subscriptionId;
    }
    /**
     * The last update of the subscription
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return The last update of the subscription
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RegionLoadBalancerLogSubscription(java.lang.String name) {
        this(name, RegionLoadBalancerLogSubscriptionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RegionLoadBalancerLogSubscription(java.lang.String name, RegionLoadBalancerLogSubscriptionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RegionLoadBalancerLogSubscription(java.lang.String name, RegionLoadBalancerLogSubscriptionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/regionLoadBalancerLogSubscription:RegionLoadBalancerLogSubscription", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RegionLoadBalancerLogSubscription(java.lang.String name, Output<java.lang.String> id, @Nullable RegionLoadBalancerLogSubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/regionLoadBalancerLogSubscription:RegionLoadBalancerLogSubscription", name, state, makeResourceOptions(options, id), false);
    }

    private static RegionLoadBalancerLogSubscriptionArgs makeArgs(RegionLoadBalancerLogSubscriptionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RegionLoadBalancerLogSubscriptionArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
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
    public static RegionLoadBalancerLogSubscription get(java.lang.String name, Output<java.lang.String> id, @Nullable RegionLoadBalancerLogSubscriptionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RegionLoadBalancerLogSubscription(name, id, state, options);
    }
}
