// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.IpLoadBalancing;

import com.ovhcloud.pulumi.ovh.IpLoadBalancing.TcpFarmArgs;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.inputs.TcpFarmState;
import com.ovhcloud.pulumi.ovh.IpLoadBalancing.outputs.TcpFarmProbe;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates a backend server group (farm) to be used by loadbalancing frontend(s)
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
 * import com.pulumi.ovh.IpLoadBalancing.IpLoadBalancingFunctions;
 * import com.pulumi.ovh.IpLoadBalancing.inputs.GetIpLoadBalancingArgs;
 * import com.pulumi.ovh.IpLoadBalancing.TcpFarm;
 * import com.pulumi.ovh.IpLoadBalancing.TcpFarmArgs;
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
 *         final var lb = IpLoadBalancingFunctions.getIpLoadBalancing(GetIpLoadBalancingArgs.builder()
 *             .serviceName("ip-1.2.3.4")
 *             .state("ok")
 *             .build());
 * 
 *         var farmName = new TcpFarm("farmName", TcpFarmArgs.builder()
 *             .displayName("ingress-8080-gra")
 *             .serviceName(lb.applyValue(getIpLoadBalancingResult -> getIpLoadBalancingResult.serviceName()))
 *             .zone("GRA")
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
 * TCP Farm can be imported using the following format `service_name` and the `id` of the farm, separated by &#34;/&#34; e.g.
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:IpLoadBalancing/tcpFarm:TcpFarm farmname service_name/farm_id
 * ```
 * 
 */
@ResourceType(type="ovh:IpLoadBalancing/tcpFarm:TcpFarm")
public class TcpFarm extends com.pulumi.resources.CustomResource {
    /**
     * Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     * 
     */
    @Export(name="balance", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> balance;

    /**
     * @return Load balancing algorithm. `roundrobin` if null (`first`, `leastconn`, `roundrobin`, `source`)
     * 
     */
    public Output<Optional<String>> balance() {
        return Codegen.optional(this.balance);
    }
    /**
     * Readable label for loadbalancer farm
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    /**
     * @return Readable label for loadbalancer farm
     * 
     */
    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    /**
     * Port attached to your farm ([1..49151]). Inherited from frontend if null
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> port;

    /**
     * @return Port attached to your farm ([1..49151]). Inherited from frontend if null
     * 
     */
    public Output<Optional<Integer>> port() {
        return Codegen.optional(this.port);
    }
    /**
     * define a backend healthcheck probe
     * 
     */
    @Export(name="probe", refs={TcpFarmProbe.class}, tree="[0]")
    private Output</* @Nullable */ TcpFarmProbe> probe;

    /**
     * @return define a backend healthcheck probe
     * 
     */
    public Output<Optional<TcpFarmProbe>> probe() {
        return Codegen.optional(this.probe);
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
     * Stickiness type. No stickiness if null (`sourceIp`)
     * 
     */
    @Export(name="stickiness", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> stickiness;

    /**
     * @return Stickiness type. No stickiness if null (`sourceIp`)
     * 
     */
    public Output<Optional<String>> stickiness() {
        return Codegen.optional(this.stickiness);
    }
    /**
     * Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     * 
     */
    @Export(name="vrackNetworkId", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> vrackNetworkId;

    /**
     * @return Internal Load Balancer identifier of the vRack private network to attach to your farm, mandatory when your Load Balancer is attached to a vRack
     * 
     */
    public Output<Optional<Integer>> vrackNetworkId() {
        return Codegen.optional(this.vrackNetworkId);
    }
    /**
     * Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     * 
     */
    @Export(name="zone", refs={String.class}, tree="[0]")
    private Output<String> zone;

    /**
     * @return Zone where the farm will be defined (ie. `GRA`, `BHS` also supports `ALL`)
     * 
     */
    public Output<String> zone() {
        return this.zone;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TcpFarm(java.lang.String name) {
        this(name, TcpFarmArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TcpFarm(java.lang.String name, TcpFarmArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TcpFarm(java.lang.String name, TcpFarmArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/tcpFarm:TcpFarm", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private TcpFarm(java.lang.String name, Output<java.lang.String> id, @Nullable TcpFarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:IpLoadBalancing/tcpFarm:TcpFarm", name, state, makeResourceOptions(options, id), false);
    }

    private static TcpFarmArgs makeArgs(TcpFarmArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? TcpFarmArgs.Empty : args;
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
    public static TcpFarm get(java.lang.String name, Output<java.lang.String> id, @Nullable TcpFarmState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TcpFarm(name, id, state, options);
    }
}
