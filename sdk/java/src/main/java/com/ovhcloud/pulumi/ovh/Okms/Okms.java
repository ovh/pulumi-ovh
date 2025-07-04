// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms;

import com.ovhcloud.pulumi.ovh.Okms.OkmsArgs;
import com.ovhcloud.pulumi.ovh.Okms.inputs.OkmsState;
import com.ovhcloud.pulumi.ovh.Okms.outputs.OkmsIam;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
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
 * import com.ovhcloud.pulumi.ovh.Okms.Okms;
 * import com.ovhcloud.pulumi.ovh.Okms.OkmsArgs;
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
 *         var newKms = new Okms("newKms", OkmsArgs.builder()
 *             .ovhSubsidiary("FR")
 *             .region("eu-west-rbx")
 *             .displayName("terraformed KMS")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="ovh:Okms/okms:Okms")
public class Okms extends com.pulumi.resources.CustomResource {
    /**
     * (String) Resource display name
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return (String) Resource display name
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * (Attributes) IAM resource metadata (see below for nested schema)
     * 
     */
    @Export(name="iam", refs={OkmsIam.class}, tree="[0]")
    private Output<OkmsIam> iam;

    /**
     * @return (Attributes) IAM resource metadata (see below for nested schema)
     * 
     */
    public Output<OkmsIam> iam() {
        return this.iam;
    }
    /**
     * (String) KMS kmip API endpoint
     * 
     */
    @Export(name="kmipEndpoint", refs={String.class}, tree="[0]")
    private Output<String> kmipEndpoint;

    /**
     * @return (String) KMS kmip API endpoint
     * 
     */
    public Output<String> kmipEndpoint() {
        return this.kmipEndpoint;
    }
    /**
     * OVH subsidiaries
     * 
     */
    @Export(name="ovhSubsidiary", refs={String.class}, tree="[0]")
    private Output<String> ovhSubsidiary;

    /**
     * @return OVH subsidiaries
     * 
     */
    public Output<String> ovhSubsidiary() {
        return this.ovhSubsidiary;
    }
    /**
     * (String) KMS public CA (Certificate Authority)
     * 
     */
    @Export(name="publicCa", refs={String.class}, tree="[0]")
    private Output<String> publicCa;

    /**
     * @return (String) KMS public CA (Certificate Authority)
     * 
     */
    public Output<String> publicCa() {
        return this.publicCa;
    }
    /**
     * KMS region
     * 
     */
    @Export(name="region", refs={String.class}, tree="[0]")
    private Output<String> region;

    /**
     * @return KMS region
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * (String) KMS rest API endpoint
     * 
     */
    @Export(name="restEndpoint", refs={String.class}, tree="[0]")
    private Output<String> restEndpoint;

    /**
     * @return (String) KMS rest API endpoint
     * 
     */
    public Output<String> restEndpoint() {
        return this.restEndpoint;
    }
    /**
     * (String) KMS rest API swagger UI
     * 
     */
    @Export(name="swaggerEndpoint", refs={String.class}, tree="[0]")
    private Output<String> swaggerEndpoint;

    /**
     * @return (String) KMS rest API swagger UI
     * 
     */
    public Output<String> swaggerEndpoint() {
        return this.swaggerEndpoint;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Okms(java.lang.String name) {
        this(name, OkmsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Okms(java.lang.String name, OkmsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Okms(java.lang.String name, OkmsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/okms:Okms", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Okms(java.lang.String name, Output<java.lang.String> id, @Nullable OkmsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Okms/okms:Okms", name, state, makeResourceOptions(options, id), false);
    }

    private static OkmsArgs makeArgs(OkmsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? OkmsArgs.Empty : args;
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
    public static Okms get(java.lang.String name, Output<java.lang.String> id, @Nullable OkmsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Okms(name, id, state, options);
    }
}
