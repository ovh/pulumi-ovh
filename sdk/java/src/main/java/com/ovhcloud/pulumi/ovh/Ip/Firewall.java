// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Ip;

import com.ovhcloud.pulumi.ovh.Ip.FirewallArgs;
import com.ovhcloud.pulumi.ovh.Ip.inputs.FirewallState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Use this resource to manage an IP firewall.
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
 * import com.ovhcloud.pulumi.ovh.Ip.Firewall;
 * import com.ovhcloud.pulumi.ovh.Ip.FirewallArgs;
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
 *         var myFirewall = new Firewall("myFirewall", FirewallArgs.builder()
 *             .ip("XXXXXX")
 *             .ipOnFirewall("XXXXXX")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:Ip/firewall:Firewall")
public class Firewall extends com.pulumi.resources.CustomResource {
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }
    /**
     * The IP or the CIDR
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return The IP or the CIDR
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * IPv4 address
     * * ` enabled  ` - Whether firewall should be enabled
     * 
     */
    @Export(name="ipOnFirewall", refs={String.class}, tree="[0]")
    private Output<String> ipOnFirewall;

    /**
     * @return IPv4 address
     * * ` enabled  ` - Whether firewall should be enabled
     * 
     */
    public Output<String> ipOnFirewall() {
        return this.ipOnFirewall;
    }
    /**
     * Current state of your ip on firewall
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Current state of your ip on firewall
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Firewall(java.lang.String name) {
        this(name, FirewallArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Firewall(java.lang.String name, FirewallArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Firewall(java.lang.String name, FirewallArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Ip/firewall:Firewall", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Firewall(java.lang.String name, Output<java.lang.String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Ip/firewall:Firewall", name, state, makeResourceOptions(options, id), false);
    }

    private static FirewallArgs makeArgs(FirewallArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FirewallArgs.Empty : args;
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
    public static Firewall get(java.lang.String name, Output<java.lang.String> id, @Nullable FirewallState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Firewall(name, id, state, options);
    }
}
