// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.FailoverIpAttachArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.FailoverIpAttachState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Attaches a failover IP address to a compute instance
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
 * import com.pulumi.ovh.CloudProject.FailoverIpAttach;
 * import com.pulumi.ovh.CloudProject.FailoverIpAttachArgs;
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
 *         var myFailoverIp = new FailoverIpAttach("myFailoverIp", FailoverIpAttachArgs.builder()
 *             .ip("XXXXXX")
 *             .routedTo("XXXXXX")
 *             .serviceName("XXXXXX")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:CloudProject/failoverIpAttach:FailoverIpAttach")
public class FailoverIpAttach extends com.pulumi.resources.CustomResource {
    /**
     * The IP block
     * * `continentCode` - The Ip continent
     * 
     */
    @Export(name="block", refs={String.class}, tree="[0]")
    private Output<String> block;

    /**
     * @return The IP block
     * * `continentCode` - The Ip continent
     * 
     */
    public Output<String> block() {
        return this.block;
    }
    /**
     * Ip continent
     * 
     */
    @Export(name="continentCode", refs={String.class}, tree="[0]")
    private Output<String> continentCode;

    /**
     * @return Ip continent
     * 
     */
    public Output<String> continentCode() {
        return this.continentCode;
    }
    /**
     * Ip location
     * 
     */
    @Export(name="geoLoc", refs={String.class}, tree="[0]")
    private Output<String> geoLoc;

    /**
     * @return Ip location
     * 
     */
    public Output<String> geoLoc() {
        return this.geoLoc;
    }
    /**
     * The failover ip address to attach
     * 
     */
    @Export(name="ip", refs={String.class}, tree="[0]")
    private Output<String> ip;

    /**
     * @return The failover ip address to attach
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }
    /**
     * Current operation progress in percent
     * * `routedTo` - Instance where ip is routed to
     * 
     */
    @Export(name="progress", refs={Integer.class}, tree="[0]")
    private Output<Integer> progress;

    /**
     * @return Current operation progress in percent
     * * `routedTo` - Instance where ip is routed to
     * 
     */
    public Output<Integer> progress() {
        return this.progress;
    }
    /**
     * The GUID of an instance to which the failover IP address is be attached
     * 
     */
    @Export(name="routedTo", refs={String.class}, tree="[0]")
    private Output<String> routedTo;

    /**
     * @return The GUID of an instance to which the failover IP address is be attached
     * 
     */
    public Output<String> routedTo() {
        return this.routedTo;
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
     * Ip status, can be `ok` or `operationPending`
     * * `subType` - IP sub type, can be `cloud` or `ovh`
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return Ip status, can be `ok` or `operationPending`
     * * `subType` - IP sub type, can be `cloud` or `ovh`
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * IP sub type
     * 
     */
    @Export(name="subType", refs={String.class}, tree="[0]")
    private Output<String> subType;

    /**
     * @return IP sub type
     * 
     */
    public Output<String> subType() {
        return this.subType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FailoverIpAttach(java.lang.String name) {
        this(name, FailoverIpAttachArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FailoverIpAttach(java.lang.String name, FailoverIpAttachArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FailoverIpAttach(java.lang.String name, FailoverIpAttachArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/failoverIpAttach:FailoverIpAttach", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private FailoverIpAttach(java.lang.String name, Output<java.lang.String> id, @Nullable FailoverIpAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProject/failoverIpAttach:FailoverIpAttach", name, state, makeResourceOptions(options, id), false);
    }

    private static FailoverIpAttachArgs makeArgs(FailoverIpAttachArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? FailoverIpAttachArgs.Empty : args;
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
    public static FailoverIpAttach get(java.lang.String name, Output<java.lang.String> id, @Nullable FailoverIpAttachState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FailoverIpAttach(name, id, state, options);
    }
}
