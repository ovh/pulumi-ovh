// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.NetworkPrivateSubnetArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.NetworkPrivateSubnetState;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.NetworkPrivateSubnetIpPool;
import com.ovhcloud.pulumi.ovh.Utilities;
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
 * Creates a subnet in a private network of a public cloud project.
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
 * import com.ovhcloud.pulumi.ovh.CloudProject.NetworkPrivateSubnet;
 * import com.ovhcloud.pulumi.ovh.CloudProject.NetworkPrivateSubnetArgs;
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
 *         var subnet = new NetworkPrivateSubnet("subnet", NetworkPrivateSubnetArgs.builder()
 *             .serviceName("xxxxx")
 *             .networkId("0234543")
 *             .region("GRA1")
 *             .start("192.168.168.100")
 *             .end("192.168.168.200")
 *             .network("192.168.168.0/24")
 *             .dhcp(true)
 *             .noGateway(false)
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
 * Subnet in a private network of a public cloud project can be imported using the `service_name` , the `network_id` as `pn-xxxx` format and the `subnet_id`, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet mysubnet service_name/network_id/subnet_id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet")
public class NetworkPrivateSubnet extends com.pulumi.resources.CustomResource {
    /**
     * Ip Block representing the subnet cidr.
     * 
     */
    @Export(name="cidr", refs={String.class}, tree="[0]")
    private Output<String> cidr;

    /**
     * @return Ip Block representing the subnet cidr.
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }
    /**
     * Enable DHCP. Changing this forces a new resource to be created. Defaults to false.
     * 
     */
    @Export(name="dhcp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dhcp;

    /**
     * @return Enable DHCP. Changing this forces a new resource to be created. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> dhcp() {
        return Codegen.optional(this.dhcp);
    }
    /**
     * Last ip for this region. Changing this value recreates the subnet.
     * 
     */
    @Export(name="end", refs={String.class}, tree="[0]")
    private Output<String> end;

    /**
     * @return Last ip for this region. Changing this value recreates the subnet.
     * 
     */
    public Output<String> end() {
        return this.end;
    }
    /**
     * The IP of the gateway
     * 
     */
    @Export(name="gatewayIp", refs={String.class}, tree="[0]")
    private Output<String> gatewayIp;

    /**
     * @return The IP of the gateway
     * 
     */
    public Output<String> gatewayIp() {
        return this.gatewayIp;
    }
    /**
     * List of ip pools allocated in the subnet.
     * * `ip_pools/network` - Global network with cidr.
     * * `ip_pools/region` - Region where this subnet is created.
     * * `ip_pools/dhcp` - DHCP enabled.
     * * `ip_pools/end` - Last ip for this region.
     * * `ip_pools/start` - First ip for this region.
     * 
     */
    @Export(name="ipPools", refs={List.class,NetworkPrivateSubnetIpPool.class}, tree="[0,1]")
    private Output<List<NetworkPrivateSubnetIpPool>> ipPools;

    /**
     * @return List of ip pools allocated in the subnet.
     * * `ip_pools/network` - Global network with cidr.
     * * `ip_pools/region` - Region where this subnet is created.
     * * `ip_pools/dhcp` - DHCP enabled.
     * * `ip_pools/end` - Last ip for this region.
     * * `ip_pools/start` - First ip for this region.
     * 
     */
    public Output<List<NetworkPrivateSubnetIpPool>> ipPools() {
        return this.ipPools;
    }
    /**
     * Global network in CIDR format. Changing this value recreates the subnet
     * 
     */
    @Export(name="network", refs={String.class}, tree="[0]")
    private Output<String> network;

    /**
     * @return Global network in CIDR format. Changing this value recreates the subnet
     * 
     */
    public Output<String> network() {
        return this.network;
    }
    /**
     * The id of the network. Changing this forces a new resource to be created.
     * 
     */
    @Export(name="networkId", refs={String.class}, tree="[0]")
    private Output<String> networkId;

    /**
     * @return The id of the network. Changing this forces a new resource to be created.
     * 
     */
    public Output<String> networkId() {
        return this.networkId;
    }
    /**
     * Set to true if you don&#39;t want to set a default gateway IP. Changing this value recreates the resource. Defaults to false.
     * 
     */
    @Export(name="noGateway", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> noGateway;

    /**
     * @return Set to true if you don&#39;t want to set a default gateway IP. Changing this value recreates the resource. Defaults to false.
     * 
     */
    public Output<Optional<Boolean>> noGateway() {
        return Codegen.optional(this.noGateway);
    }
    /**
     * The region in which the network subnet will be created. Ex.: &#34;GRA1&#34;. Changing this value recreates the resource.
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return The region in which the network subnet will be created. Ex.: &#34;GRA1&#34;. Changing this value recreates the resource.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * First ip for this region. Changing this value recreates the subnet.
     * 
     */
    @Export(name="start", refs={String.class}, tree="[0]")
    private Output<String> start;

    /**
     * @return First ip for this region. Changing this value recreates the subnet.
     * 
     */
    public Output<String> start() {
        return this.start;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkPrivateSubnet(java.lang.String name) {
        this(name, NetworkPrivateSubnetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkPrivateSubnet(java.lang.String name, NetworkPrivateSubnetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkPrivateSubnet(java.lang.String name, NetworkPrivateSubnetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NetworkPrivateSubnet(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPrivateSubnetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/networkPrivateSubnet:NetworkPrivateSubnet", name, state, makeResourceOptions(options, id), false);
    }

    private static NetworkPrivateSubnetArgs makeArgs(NetworkPrivateSubnetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NetworkPrivateSubnetArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .pluginDownloadURL("github://api.github.com/ovh/pulumi-ovh")
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
    public static NetworkPrivateSubnet get(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPrivateSubnetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkPrivateSubnet(name, id, state, options);
    }
}
