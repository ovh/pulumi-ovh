// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms;

import com.ovhcloud.pulumi.ovh.Okms.ServiceKeyArgs;
import com.ovhcloud.pulumi.ovh.Okms.inputs.ServiceKeyState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Creates a Service Key in an OVHcloud KMS.
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
 * import com.ovhcloud.pulumi.ovh.Okms.ServiceKey;
 * import com.ovhcloud.pulumi.ovh.Okms.ServiceKeyArgs;
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
 *         var keySymetric = new ServiceKey("keySymetric", ServiceKeyArgs.builder()
 *             .okmsId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
 *             .operations(            
 *                 "encrypt",
 *                 "decrypt")
 *             .size(256.0)
 *             .type("oct")
 *             .build());
 * 
 *         var keyRsa = new ServiceKey("keyRsa", ServiceKeyArgs.builder()
 *             .okmsId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
 *             .operations(            
 *                 "sign",
 *                 "verify")
 *             .size(2048.0)
 *             .type("RSA")
 *             .build());
 * 
 *         var keyEcdsa = new ServiceKey("keyEcdsa", ServiceKeyArgs.builder()
 *             .curve("P-256")
 *             .okmsId("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx")
 *             .operations(            
 *                 "sign",
 *                 "verify")
 *             .type("EC")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:Okms/serviceKey:ServiceKey")
public class ServiceKey extends com.pulumi.resources.CustomResource {
    /**
     * Context of the key
     * 
     */
    @Export(name="context", refs={String.class}, tree="[0]")
    private Output<String> context;

    /**
     * @return Context of the key
     * 
     */
    public Output<String> context() {
        return this.context;
    }
    /**
     * Creation time of the key
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Creation time of the key
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * Curve type for Elliptic Curve (EC) keys
     * 
     */
    @Export(name="curve", refs={String.class}, tree="[0]")
    private Output<String> curve;

    /**
     * @return Curve type for Elliptic Curve (EC) keys
     * 
     */
    public Output<String> curve() {
        return this.curve;
    }
    /**
     * Key deactivation reason
     * 
     */
    @Export(name="deactivationReason", refs={String.class}, tree="[0]")
    private Output<String> deactivationReason;

    /**
     * @return Key deactivation reason
     * 
     */
    public Output<String> deactivationReason() {
        return this.deactivationReason;
    }
    /**
     * Key name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Key name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Okms ID
     * 
     */
    @Export(name="okmsId", refs={String.class}, tree="[0]")
    private Output<String> okmsId;

    /**
     * @return Okms ID
     * 
     */
    public Output<String> okmsId() {
        return this.okmsId;
    }
    /**
     * The operations for which the key is intended to be used
     * 
     */
    @Export(name="operations", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> operations;

    /**
     * @return The operations for which the key is intended to be used
     * 
     */
    public Output<List<String>> operations() {
        return this.operations;
    }
    /**
     * Size of the key to be created
     * 
     */
    @Export(name="size", refs={Double.class}, tree="[0]")
    private Output<Double> size;

    /**
     * @return Size of the key to be created
     * 
     */
    public Output<Double> size() {
        return this.size;
    }
    /**
     * State of the key
     * 
     */
    @Export(name="state", refs={String.class}, tree="[0]")
    private Output<String> state;

    /**
     * @return State of the key
     * 
     */
    public Output<String> state() {
        return this.state;
    }
    /**
     * Type of the key to be created
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return Type of the key to be created
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ServiceKey(java.lang.String name) {
        this(name, ServiceKeyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ServiceKey(java.lang.String name, ServiceKeyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ServiceKey(java.lang.String name, ServiceKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/serviceKey:ServiceKey", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ServiceKey(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/serviceKey:ServiceKey", name, state, makeResourceOptions(options, id), false);
    }

    private static ServiceKeyArgs makeArgs(ServiceKeyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ServiceKeyArgs.Empty : args;
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
    public static ServiceKey get(java.lang.String name, Output<java.lang.String> id, @Nullable ServiceKeyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ServiceKey(name, id, state, options);
    }
}
