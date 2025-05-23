// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain;

import com.ovhcloud.pulumi.ovh.Domain.DSRecordsArgs;
import com.ovhcloud.pulumi.ovh.Domain.inputs.DSRecordsState;
import com.ovhcloud.pulumi.ovh.Domain.outputs.DSRecordsDsRecord;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Use this resource to manage a domain&#39;s DS records.
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
 * import com.ovhcloud.pulumi.ovh.Domain.DSRecords;
 * import com.ovhcloud.pulumi.ovh.Domain.DSRecordsArgs;
 * import com.pulumi.ovh.Domain.inputs.DSRecordsDsRecordArgs;
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
 *         var dsRecords = new DSRecords("dsRecords", DSRecordsArgs.builder()
 *             .domain("mydomain.ovh")
 *             .dsRecords(DSRecordsDsRecordArgs.builder()
 *                 .algorithm("RSASHA1_NSEC3_SHA1")
 *                 .flags("KEY_SIGNING_KEY")
 *                 .publicKey("my_base64_encoded_public_key")
 *                 .tag(12345)
 *                 .build())
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
 * DS records can be imported using their `domain`.
 * 
 * Using the following configuration:
 * 
 * terraform
 * 
 * import {
 * 
 *   to = ovh_domain_ds_records.ds_records
 * 
 *   id = &#34;&lt;domain name&gt;&#34;
 * 
 * }
 * 
 * You can then run:
 * 
 * bash
 * 
 * $ pulumi preview -generate-config-out=ds_records.tf
 * 
 * $ pulumi up
 * 
 * The file `ds_records.tf` will then contain the imported resource&#39;s configuration, that can be copied next to the `import` block above. See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 * 
 */
@ResourceType(type="ovh:Domain/dSRecords:DSRecords")
public class DSRecords extends com.pulumi.resources.CustomResource {
    /**
     * Domain name for which to manage DS records
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return Domain name for which to manage DS records
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * Details about a DS record
     * 
     */
    @Export(name="dsRecords", refs={List.class,DSRecordsDsRecord.class}, tree="[0,1]")
    private Output<List<DSRecordsDsRecord>> dsRecords;

    /**
     * @return Details about a DS record
     * 
     */
    public Output<List<DSRecordsDsRecord>> dsRecords() {
        return this.dsRecords;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DSRecords(java.lang.String name) {
        this(name, DSRecordsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DSRecords(java.lang.String name, DSRecordsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DSRecords(java.lang.String name, DSRecordsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/dSRecords:DSRecords", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private DSRecords(java.lang.String name, Output<java.lang.String> id, @Nullable DSRecordsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/dSRecords:DSRecords", name, state, makeResourceOptions(options, id), false);
    }

    private static DSRecordsArgs makeArgs(DSRecordsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? DSRecordsArgs.Empty : args;
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
    public static DSRecords get(java.lang.String name, Output<java.lang.String> id, @Nullable DSRecordsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DSRecords(name, id, state, options);
    }
}
