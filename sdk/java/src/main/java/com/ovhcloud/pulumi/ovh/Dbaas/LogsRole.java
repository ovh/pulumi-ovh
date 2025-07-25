// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dbaas;

import com.ovhcloud.pulumi.ovh.Dbaas.LogsRoleArgs;
import com.ovhcloud.pulumi.ovh.Dbaas.inputs.LogsRoleState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Reference a DBaaS logs role.
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
 * import com.ovhcloud.pulumi.ovh.Dbaas.LogsRole;
 * import com.ovhcloud.pulumi.ovh.Dbaas.LogsRoleArgs;
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
 *         var ro = new LogsRole("ro", LogsRoleArgs.builder()
 *             .serviceName("ldp-xx-xxxxx")
 *             .name("Devops - RO")
 *             .description("Devops - RO")
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
 * OVHcloud DBaaS Log Role can be imported using the `service_name` and `role_id` of the role, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Dbaas/logsRole:LogsRole  ovh_dbaas_logs_role.ro ldp-ra-XX/dc145bc2-eb01-4efe-a802-XXXXXX
 * ```
 * 
 */
@ResourceType(type="ovh:Dbaas/logsRole:LogsRole")
public class LogsRole extends com.pulumi.resources.CustomResource {
    /**
     * Role creation date
     * 
     */
    @Export(name="createdAt", refs={String.class}, tree="[0]")
    private Output<String> createdAt;

    /**
     * @return Role creation date
     * 
     */
    public Output<String> createdAt() {
        return this.createdAt;
    }
    /**
     * The role description
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return The role description
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The role name
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The role name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * number of member for the role
     * 
     */
    @Export(name="nbMember", refs={Integer.class}, tree="[0]")
    private Output<Integer> nbMember;

    /**
     * @return number of member for the role
     * 
     */
    public Output<Integer> nbMember() {
        return this.nbMember;
    }
    /**
     * number of configured permission for the role
     * 
     */
    @Export(name="nbPermission", refs={Integer.class}, tree="[0]")
    private Output<Integer> nbPermission;

    /**
     * @return number of configured permission for the role
     * 
     */
    public Output<Integer> nbPermission() {
        return this.nbPermission;
    }
    /**
     * Role identifier
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return Role identifier
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * The service name
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Role last update date
     * 
     */
    @Export(name="updatedAt", refs={String.class}, tree="[0]")
    private Output<String> updatedAt;

    /**
     * @return Role last update date
     * 
     */
    public Output<String> updatedAt() {
        return this.updatedAt;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public LogsRole(java.lang.String name) {
        this(name, LogsRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public LogsRole(java.lang.String name, LogsRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public LogsRole(java.lang.String name, LogsRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsRole:LogsRole", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private LogsRole(java.lang.String name, Output<java.lang.String> id, @Nullable LogsRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Dbaas/logsRole:LogsRole", name, state, makeResourceOptions(options, id), false);
    }

    private static LogsRoleArgs makeArgs(LogsRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? LogsRoleArgs.Empty : args;
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
    public static LogsRole get(java.lang.String name, Output<java.lang.String> id, @Nullable LogsRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new LogsRole(name, id, state, options);
    }
}
