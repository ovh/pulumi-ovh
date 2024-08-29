// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovh.ovh.Ip;

import com.ovh.ovh.Ip.MitigationArgs;
import com.ovh.ovh.Ip.inputs.MitigationState;
import com.ovh.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Use this resource to manage an IP permanent mitigation.
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
 * import com.pulumi.ovh.Ip.Mitigation;
 * import com.pulumi.ovh.Ip.MitigationArgs;
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
 *         var mitigation = new Mitigation("mitigation", MitigationArgs.builder()
 *             .ip("XXXXXX")
 *             .ipOnMitigation("XXXXXX")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:Ip/mitigation:Mitigation")
public class Mitigation extends com.pulumi.resources.CustomResource {
    /**
     * Set on true if the IP is on auto-mitigation
     * 
     */
    @Export(name="auto", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> auto;

    /**
     * @return Set on true if the IP is on auto-mitigation
     * 
     */
    public Output<Boolean> auto() {
        return this.auto;
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
     * * ` permanent  ` - Set on true if the IP is on permanent mitigation
     * 
     */
    @Export(name="ipOnMitigation", refs={String.class}, tree="[0]")
    private Output<String> ipOnMitigation;

    /**
     * @return IPv4 address
     * * ` permanent  ` - Set on true if the IP is on permanent mitigation
     * 
     */
    public Output<String> ipOnMitigation() {
        return this.ipOnMitigation;
    }
    /**
     * Set on true if your ip is on permanent mitigation
     * 
     */
    @Export(name="permanent", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> permanent;

    /**
     * @return Set on true if your ip is on permanent mitigation
     * 
     */
    public Output<Boolean> permanent() {
        return this.permanent;
    }
    /**
     * Current state of the IP on mitigation
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return Current state of the IP on mitigation
     * 
     */
    public Output<String> state() {
        return this.state;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Mitigation(java.lang.String name) {
        this(name, MitigationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Mitigation(java.lang.String name, MitigationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Mitigation(java.lang.String name, MitigationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Ip/mitigation:Mitigation", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Mitigation(java.lang.String name, Output<java.lang.String> id, @Nullable MitigationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Ip/mitigation:Mitigation", name, state, makeResourceOptions(options, id), false);
    }

    private static MitigationArgs makeArgs(MitigationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? MitigationArgs.Empty : args;
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
    public static Mitigation get(java.lang.String name, Output<java.lang.String> id, @Nullable MitigationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Mitigation(name, id, state, options);
    }
}
