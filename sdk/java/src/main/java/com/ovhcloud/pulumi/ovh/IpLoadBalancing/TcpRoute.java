// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing;

import com.ovhcloud.pulumi.ovh.IpLoadBalancing.TcpRouteArgs;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.inputs.TcpRouteState;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.outputs.TcpRouteAction;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.outputs.TcpRouteRule;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manage TCP route for a loadbalancer service
 * 
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
 * import com.pulumi.ovh.IpLoadBalancing.TcpRoute;
 * import com.pulumi.ovh.IpLoadBalancing.TcpRouteArgs;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.TcpRouteActionArgs;
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
 *         var tcpReject = new TcpRoute("tcpReject", TcpRouteArgs.builder()
 *             .action(TcpRouteActionArgs.builder()
 *                 .type("reject")
 *                 .build())
 *             .serviceName("loadbalancer-xxxxxxxxxxxxxxxxxx")
 *             .weight(1)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * TCP route can be imported using the following format `service_name` and the `id` of the route separated by &#34;/&#34; e.g.
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:IpLoadBalancing/tcpRoute:TcpRoute tcpreject service_name/route_id
 * ```
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/tcpRoute:TcpRoute")
public class TcpRoute extends com.pulumi.resources.CustomResource {
    /**
     * Action triggered when all rules match
     * 
     */
    @Export(name="action", refs={TcpRouteAction.class}, tree="[0]")
    private Output<TcpRouteAction> action;

    /**
     * @return Action triggered when all rules match
     * 
     */
    public Output<TcpRouteAction> action() {
        return this.action;
    }
    /**
     * Human readable name for your route, this field is for you
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Human readable name for your route, this field is for you
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Route traffic for this frontend
     * 
     */
    @Export(name="frontendId", refs={Integer.class}, tree="[0]")
    private Output<Integer> frontendId;

    /**
     * @return Route traffic for this frontend
     * 
     */
    public Output<Integer> frontendId() {
        return this.frontendId;
    }
    /**
     * List of rules to match to trigger action
     * 
     */
    @Export(name="rules", refs={List.class,TcpRouteRule.class}, tree="[0,1]")
    private Output<List<TcpRouteRule>> rules;

    /**
     * @return List of rules to match to trigger action
     * 
     */
    public Output<List<TcpRouteRule>> rules() {
        return this.rules;
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
     * Route status. Routes in &#34;ok&#34; state are ready to operate
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Route status. Routes in &#34;ok&#34; state are ready to operate
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
     * 
     */
    @Export(name="weight", refs={Integer.class}, tree="[0]")
    private Output<Integer> weight;

    /**
     * @return Route priority ([0..255]). 0 if null. Highest priority routes are evaluated first. Only the first matching route will trigger an action
     * 
     */
    public Output<Integer> weight() {
        return this.weight;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TcpRoute(java.lang.String name) {
        this(name, TcpRouteArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TcpRoute(java.lang.String name, TcpRouteArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TcpRoute(java.lang.String name, TcpRouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/tcpRoute:TcpRoute", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TcpRoute(java.lang.String name, Output<java.lang.String> id, @Nullable TcpRouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/tcpRoute:TcpRoute", name, state, makeResourceOptions(options, id), false);
    }

    private static TcpRouteArgs makeArgs(TcpRouteArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TcpRouteArgs.Empty : args;
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
    public static TcpRoute get(java.lang.String name, Output<java.lang.String> id, @Nullable TcpRouteState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TcpRoute(name, id, state, options);
    }
}
