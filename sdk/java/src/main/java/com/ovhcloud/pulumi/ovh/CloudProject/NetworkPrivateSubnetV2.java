// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.NetworkPrivateSubnetV2Args;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.NetworkPrivateSubnetV2State;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.NetworkPrivateSubnetV2AllocationPool;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.NetworkPrivateSubnetV2HostRoute;
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
 * Creates a subnet in a private network of a public cloud region.
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
 * import com.pulumi.ovh.CloudProject.NetworkPrivateSubnetV2;
 * import com.pulumi.ovh.CloudProject.NetworkPrivateSubnetV2Args;
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
 *         var subnet = new NetworkPrivateSubnetV2("subnet", NetworkPrivateSubnetV2Args.builder()
 *             .cidr("192.168.168.0/24")
 *             .dhcp(true)
 *             .dnsNameservers("1.1.1.1")
 *             .enableGatewayIp(true)
 *             .networkId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
 *             .region("XXX1")
 *             .serviceName("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
 *             .useDefaultPublicDnsResolver(false)
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
 * Subnet in a private network of a public cloud project can be imported using the `service_name`, `region`, `network_id` and `subnet_id`, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2 mysubnet 5ceb661434891538b54a4f2c66fc4b746e/BHS5/25807101-8aaa-4ea5-b507-61f0d661b101/0f0b73a4-403b-45e4-86d0-b438f1291909
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2")
public class NetworkPrivateSubnetV2 extends com.pulumi.resources.CustomResource {
    /**
     * List of IP allocation pools Changing this value recreates the resource.
     * 
     */
    @Export(name="allocationPools", refs={List.class,NetworkPrivateSubnetV2AllocationPool.class}, tree="[0,1]")
    private Output<List<NetworkPrivateSubnetV2AllocationPool>> allocationPools;

    /**
     * @return List of IP allocation pools Changing this value recreates the resource.
     * 
     */
    public Output<List<NetworkPrivateSubnetV2AllocationPool>> allocationPools() {
        return this.allocationPools;
    }
    /**
     * IP range of the subnet Changing this value recreates the subnet.
     * 
     */
    @Export(name="cidr", refs={String.class}, tree="[0]")
    private Output<String> cidr;

    /**
     * @return IP range of the subnet Changing this value recreates the subnet.
     * 
     */
    public Output<String> cidr() {
        return this.cidr;
    }
    /**
     * Enable DHCP. Changing this forces a new resource to be created. Defaults to true.
     * 
     */
    @Export(name="dhcp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dhcp;

    /**
     * @return Enable DHCP. Changing this forces a new resource to be created. Defaults to true.
     * 
     */
    public Output<Optional<Boolean>> dhcp() {
        return Codegen.optional(this.dhcp);
    }
    /**
     * DNS nameservers used by DHCP Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
     * 
     */
    @Export(name="dnsNameservers", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> dnsNameservers;

    /**
     * @return DNS nameservers used by DHCP Changing this value recreates the resource. Defaults to OVH default DNS nameserver.
     * 
     */
    public Output<List<String>> dnsNameservers() {
        return this.dnsNameservers;
    }
    /**
     * Set to true if you want to set a default gateway IP. Changing this value recreates the resource. Defaults to true.
     * 
     */
    @Export(name="enableGatewayIp", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableGatewayIp;

    /**
     * @return Set to true if you want to set a default gateway IP. Changing this value recreates the resource. Defaults to true.
     * 
     */
    public Output<Optional<Boolean>> enableGatewayIp() {
        return Codegen.optional(this.enableGatewayIp);
    }
    /**
     * See Argument Reference above.
     * 
     */
    @Export(name="gatewayIp", refs={String.class}, tree="[0]")
    private Output<String> gatewayIp;

    /**
     * @return See Argument Reference above.
     * 
     */
    public Output<String> gatewayIp() {
        return this.gatewayIp;
    }
    /**
     * Static host routes of subnet
     * 
     */
    @Export(name="hostRoutes", refs={List.class,NetworkPrivateSubnetV2HostRoute.class}, tree="[0,1]")
    private Output</* @Nullable */ List<NetworkPrivateSubnetV2HostRoute>> hostRoutes;

    /**
     * @return Static host routes of subnet
     * 
     */
    public Output<Optional<List<NetworkPrivateSubnetV2HostRoute>>> hostRoutes() {
        return Codegen.optional(this.hostRoutes);
    }
    /**
     * Name of the subnet Changing this value recreates the subnet.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the subnet Changing this value recreates the subnet.
     * 
     */
    public Output<String> name() {
        return this.name;
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
     * Set to false if you want to use your DNS resolver. Changing this value recreates the resource.
     * 
     */
    @Export(name="useDefaultPublicDnsResolver", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useDefaultPublicDnsResolver;

    /**
     * @return Set to false if you want to use your DNS resolver. Changing this value recreates the resource.
     * 
     */
    public Output<Optional<Boolean>> useDefaultPublicDnsResolver() {
        return Codegen.optional(this.useDefaultPublicDnsResolver);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkPrivateSubnetV2(java.lang.String name) {
        this(name, NetworkPrivateSubnetV2Args.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkPrivateSubnetV2(java.lang.String name, NetworkPrivateSubnetV2Args args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkPrivateSubnetV2(java.lang.String name, NetworkPrivateSubnetV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NetworkPrivateSubnetV2(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPrivateSubnetV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/networkPrivateSubnetV2:NetworkPrivateSubnetV2", name, state, makeResourceOptions(options, id), false);
    }

    private static NetworkPrivateSubnetV2Args makeArgs(NetworkPrivateSubnetV2Args args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NetworkPrivateSubnetV2Args.Empty : args;
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
    public static NetworkPrivateSubnetV2 get(java.lang.String name, Output<java.lang.String> id, @Nullable NetworkPrivateSubnetV2State state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkPrivateSubnetV2(name, id, state, options);
    }
}
