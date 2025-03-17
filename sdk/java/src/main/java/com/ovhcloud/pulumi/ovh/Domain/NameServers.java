// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain;

import com.ovhcloud.pulumi.ovh.Domain.NameServersArgs;
import com.ovhcloud.pulumi.ovh.Domain.inputs.NameServersState;
import com.ovhcloud.pulumi.ovh.Domain.outputs.NameServersServer;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Use this resource to manage a domain&#39;s name servers.
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
 * import com.pulumi.ovh.Domain.NameServers;
 * import com.pulumi.ovh.Domain.NameServersArgs;
 * import com.pulumi.ovh.Domain.inputs.NameServersServerArgs;
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
 *         var nameServers = new NameServers("nameServers", NameServersArgs.builder()
 *             .domain("mydomain.ovh")
 *             .servers(            
 *                 NameServersServerArgs.builder()
 *                     .host("dns105.ovh.net")
 *                     .ip("213.251.188.144")
 *                     .build(),
 *                 NameServersServerArgs.builder()
 *                     .host("ns105.ovh.net")
 *                     .build())
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
 * Name servers can be imported using their `domain`.
 * 
 * Using the following configuration:
 * 
 * hcl
 * 
 * import {
 * 
 *   to = ovh_domain_name_servers.name_servers
 * 
 *   id = &#34;&lt;domain name&gt;&#34;
 * 
 * }
 * 
 * You can then run:
 * 
 * bash
 * 
 * $ pulumi preview -generate-config-out=name_servers.tf
 * 
 * $ pulumi up
 * 
 * The file `name_servers.tf` will then contain the imported resource&#39;s configuration, that can be copied next to the `import` block above.
 * 
 * See https://developer.hashicorp.com/terraform/language/import/generating-configuration for more details.
 * 
 */
@ResourceType(type="ovh:Domain/nameServers:NameServers")
public class NameServers extends com.pulumi.resources.CustomResource {
    /**
     * Domain name for which to manage name servers
     * 
     */
    @Export(name="domain", refs={String.class}, tree="[0]")
    private Output<String> domain;

    /**
     * @return Domain name for which to manage name servers
     * 
     */
    public Output<String> domain() {
        return this.domain;
    }
    /**
     * Details about a name server
     * 
     */
    @Export(name="servers", refs={List.class,NameServersServer.class}, tree="[0,1]")
    private Output<List<NameServersServer>> servers;

    /**
     * @return Details about a name server
     * 
     */
    public Output<List<NameServersServer>> servers() {
        return this.servers;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NameServers(java.lang.String name) {
        this(name, NameServersArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NameServers(java.lang.String name, NameServersArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NameServers(java.lang.String name, NameServersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/nameServers:NameServers", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private NameServers(java.lang.String name, Output<java.lang.String> id, @Nullable NameServersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Domain/nameServers:NameServers", name, state, makeResourceOptions(options, id), false);
    }

    private static NameServersArgs makeArgs(NameServersArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? NameServersArgs.Empty : args;
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
    public static NameServers get(java.lang.String name, Output<java.lang.String> id, @Nullable NameServersState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NameServers(name, id, state, options);
    }
}
