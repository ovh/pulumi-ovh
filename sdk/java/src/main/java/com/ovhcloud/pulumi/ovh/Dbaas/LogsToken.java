// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas;

import com.ovhcloud.pulumi.ovh.Dbaas.LogsTokenArgs;
import com.ovhcloud.pulumi.ovh.Dbaas.inputs.LogsTokenState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Allows to manipulate LDP tokens.
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
 * import com.ovhcloud.pulumi.ovh.Dbaas.LogsToken;
 * import com.ovhcloud.pulumi.ovh.Dbaas.LogsTokenArgs;
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
 *         var token = new LogsToken("token", LogsTokenArgs.builder()
 *             .serviceName("ldp-xx-xxxxx")
 *             .name("ExampleToken")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:Dbaas/logsToken:LogsToken")
public class LogsToken extends com.pulumi.resources.CustomResource {
    /**
     * Cluster ID. If not provided, the default cluster_id is used
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return Cluster ID. If not provided, the default cluster_id is used
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Token creation date
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Token creation date
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Name of the token
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the token
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The LDP service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The LDP service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * ID of the token
     * 
     */
    @Export(name="tokenId", refs={String.class}, tree="[0]")
    private Output<String> tokenId;

    /**
     * @return ID of the token
     * 
     */
    public Output<String> tokenId() {
        return this.tokenId;
    }
    /**
     * Token last update date
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Token last update date
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }
    /**
     * Token value
     * 
     */
    @Export(name="value", refs={String.class}, tree="[0]")
    private Output<String> value;

    /**
     * @return Token value
     * 
     */
    public Output<String> value() {
        return this.value;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LogsToken(java.lang.String name) {
        this(name, LogsTokenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LogsToken(java.lang.String name, LogsTokenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LogsToken(java.lang.String name, LogsTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsToken:LogsToken", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LogsToken(java.lang.String name, Output<java.lang.String> id, @Nullable LogsTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsToken:LogsToken", name, state, makeResourceOptions(options, id), false);
    }

    private static LogsTokenArgs makeArgs(LogsTokenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LogsTokenArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "value"
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
    public static LogsToken get(java.lang.String name, Output<java.lang.String> id, @Nullable LogsTokenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LogsToken(name, id, state, options);
    }
}
