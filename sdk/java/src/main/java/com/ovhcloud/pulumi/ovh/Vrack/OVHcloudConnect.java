// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Vrack;

import com.ovhcloud.pulumi.ovh.Utilities;
import com.ovhcloud.pulumi.ovh.Vrack.OVHcloudConnectArgs;
import com.ovhcloud.pulumi.ovh.Vrack.inputs.OVHcloudConnectState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Attach an OVH Cloud Connect to the vrack.
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
 * import com.ovhcloud.pulumi.ovh.Vrack.OVHcloudConnect;
 * import com.ovhcloud.pulumi.ovh.Vrack.OVHcloudConnectArgs;
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
 *         var vrackOvhcloudconnect = new OVHcloudConnect("vrackOvhcloudconnect", OVHcloudConnectArgs.builder()
 *             .ovhCloudConnect("<OVH Cloud Connect service name>")
 *             .serviceName("<vRack service name>")
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
 * Attachment of an OVH Cloud Connect and a vRack can be imported using the `service_name` (vRack identifier) and the `ovh_cloud_connect` (OVH Cloud Connect service name), separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Vrack/oVHcloudConnect:OVHcloudConnect myattach &#34;&lt;service_name&gt;/&lt;OVH Cloud Connect service name&gt;&#34;
 * ```
 * 
 */
@ResourceType(type="ovh:Vrack/oVHcloudConnect:OVHcloudConnect")
public class OVHcloudConnect extends com.pulumi.resources.CustomResource {
    /**
     * Your OVH Cloud Connect service name.
     * 
     */
    @Export(name="ovhCloudConnect", refs={String.class}, tree="[0]")
    private Output<String> ovhCloudConnect;

    /**
     * @return Your OVH Cloud Connect service name.
     * 
     */
    public Output<String> ovhCloudConnect() {
        return this.ovhCloudConnect;
    }
    /**
     * The internal name of your vrack
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The internal name of your vrack
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public OVHcloudConnect(java.lang.String name) {
        this(name, OVHcloudConnectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public OVHcloudConnect(java.lang.String name, OVHcloudConnectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public OVHcloudConnect(java.lang.String name, OVHcloudConnectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/oVHcloudConnect:OVHcloudConnect", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private OVHcloudConnect(java.lang.String name, Output<java.lang.String> id, @Nullable OVHcloudConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Vrack/oVHcloudConnect:OVHcloudConnect", name, state, makeResourceOptions(options, id), false);
    }

    private static OVHcloudConnectArgs makeArgs(OVHcloudConnectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OVHcloudConnectArgs.Empty : args;
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
    public static OVHcloudConnect get(java.lang.String name, Output<java.lang.String> id, @Nullable OVHcloudConnectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new OVHcloudConnect(name, id, state, options);
    }
}
