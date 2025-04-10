// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me;

import com.ovhcloud.pulumi.ovh.Me.APIOAuth2ClientArgs;
import com.ovhcloud.pulumi.ovh.Me.inputs.APIOAuth2ClientState;
import com.ovhcloud.pulumi.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Creates an OAuth2 service account.
 * 
 * ## Example Usage
 * 
 * An OAuth2 client for an app hosted at `my-app.com`, that uses the authorization code flow to authenticate.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.ovhcloud.pulumi.ovh.Me.APIOAuth2Client;
 * import com.ovhcloud.pulumi.ovh.Me.APIOAuth2ClientArgs;
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
 *         var myOauth2ClientAuthCode = new APIOAuth2Client("myOauth2ClientAuthCode", APIOAuth2ClientArgs.builder()
 *             .callbackUrls("https://my-app.com/callback")
 *             .description("An OAuth2 client using the authorization code flow for my-app.com")
 *             .flow("AUTHORIZATION_CODE")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * An OAuth2 client for an app hosted at `my-app.com`, that uses the client credentials flow to authenticate.
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.ovhcloud.pulumi.ovh.Me.APIOAuth2Client;
 * import com.ovhcloud.pulumi.ovh.Me.APIOAuth2ClientArgs;
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
 *         var myOauth2ClientClientCreds = new APIOAuth2Client("myOauth2ClientClientCreds", APIOAuth2ClientArgs.builder()
 *             .description("An OAuth2 client using the client credentials flow for my app")
 *             .flow("CLIENT_CREDENTIALS")
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
 * OAuth2 clients can be imported using their `client_id`:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Me/aPIOAuth2Client:APIOAuth2Client my_oauth2_client client_id
 * ```
 * 
 * Because the client_secret is only available for resources created using terraform, OAuth2 clients can also be imported using a `client_id` and a `client_secret` with a pipe separator:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:Me/aPIOAuth2Client:APIOAuth2Client my_oauth2_client &#39;client_id|client_secret&#39;
 * ```
 * 
 */
@ResourceType(type="ovh:Me/aPIOAuth2Client:APIOAuth2Client")
public class APIOAuth2Client extends com.pulumi.resources.CustomResource {
    /**
     * List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
     * 
     */
    @Export(name="callbackUrls", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> callbackUrls;

    /**
     * @return List of callback urls when configuring the `AUTHORIZATION_CODE` flow.
     * 
     */
    public Output<Optional<List<String>>> callbackUrls() {
        return Codegen.optional(this.callbackUrls);
    }
    /**
     * Client ID of the created service account.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output<String> clientId;

    /**
     * @return Client ID of the created service account.
     * 
     */
    public Output<String> clientId() {
        return this.clientId;
    }
    /**
     * Client secret of the created service account.
     * 
     */
    @Export(name="clientSecret", refs={String.class}, tree="[0]")
    private Output<String> clientSecret;

    /**
     * @return Client secret of the created service account.
     * 
     */
    public Output<String> clientSecret() {
        return this.clientSecret;
    }
    /**
     * OAuth2 client description.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output<String> description;

    /**
     * @return OAuth2 client description.
     * 
     */
    public Output<String> description() {
        return this.description;
    }
    /**
     * The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
     * 
     */
    @Export(name="flow", refs={String.class}, tree="[0]")
    private Output<String> flow;

    /**
     * @return The OAuth2 flow to use. `AUTHORIZATION_CODE` or `CLIENT_CREDENTIALS` are supported at the moment.
     * 
     */
    public Output<String> flow() {
        return this.flow;
    }
    /**
     * URN that will allow you to associate this oauth2 client with an access policy
     * 
     */
    @Export(name="identity", refs={String.class}, tree="[0]")
    private Output<String> identity;

    /**
     * @return URN that will allow you to associate this oauth2 client with an access policy
     * 
     */
    public Output<String> identity() {
        return this.identity;
    }
    /**
     * OAuth2 client name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return OAuth2 client name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public APIOAuth2Client(java.lang.String name) {
        this(name, APIOAuth2ClientArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public APIOAuth2Client(java.lang.String name, APIOAuth2ClientArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public APIOAuth2Client(java.lang.String name, APIOAuth2ClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/aPIOAuth2Client:APIOAuth2Client", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private APIOAuth2Client(java.lang.String name, Output<java.lang.String> id, @Nullable APIOAuth2ClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:Me/aPIOAuth2Client:APIOAuth2Client", name, state, makeResourceOptions(options, id), false);
    }

    private static APIOAuth2ClientArgs makeArgs(APIOAuth2ClientArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? APIOAuth2ClientArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "clientSecret"
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
    public static APIOAuth2Client get(java.lang.String name, Output<java.lang.String> id, @Nullable APIOAuth2ClientState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new APIOAuth2Client(name, id, state, options);
    }
}
