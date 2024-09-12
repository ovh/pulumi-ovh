// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProjectDatabase;

import com.ovhcloud.ovh.CloudProjectDatabase.PostgresSqlConnectionPoolArgs;
import com.ovhcloud.ovh.CloudProjectDatabase.inputs.PostgresSqlConnectionPoolState;
import com.ovhcloud.ovh.Utilities;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * OVHcloud Managed PostgreSQL clusters connection pools can be imported using the `service_name`, `cluster_id` and `id` of the connection pool, separated by &#34;/&#34; E.g.,
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import ovh:CloudProjectDatabase/postgresSqlConnectionPool:PostgresSqlConnectionPool my_connection_pool service_name/cluster_id/id
 * ```
 * 
 */
@ResourceType(type="ovh:CloudProjectDatabase/postgresSqlConnectionPool:PostgresSqlConnectionPool")
public class PostgresSqlConnectionPool extends com.pulumi.resources.CustomResource {
    /**
     * Cluster ID.
     * 
     */
    @Export(name="clusterId", refs={String.class}, tree="[0]")
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }
    /**
     * Database ID for a database that belongs to the Database cluster given above.
     * 
     */
    @Export(name="databaseId", refs={String.class}, tree="[0]")
    private Output<String> databaseId;

    /**
     * @return Database ID for a database that belongs to the Database cluster given above.
     * 
     */
    public Output<String> databaseId() {
        return this.databaseId;
    }
    /**
     * Connection mode to the connection pool
     * Available modes:
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output<String> mode;

    /**
     * @return Connection mode to the connection pool
     * Available modes:
     * 
     */
    public Output<String> mode() {
        return this.mode;
    }
    /**
     * Name of the connection pool.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the connection pool.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Port of the connection pool.
     * 
     */
    @Export(name="port", refs={Integer.class}, tree="[0]")
    private Output<Integer> port;

    /**
     * @return Port of the connection pool.
     * 
     */
    public Output<Integer> port() {
        return this.port;
    }
    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Export(name="serviceName", refs={String.class}, tree="[0]")
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }
    /**
     * Size of the connection pool.
     * 
     */
    @Export(name="size", refs={Integer.class}, tree="[0]")
    private Output<Integer> size;

    /**
     * @return Size of the connection pool.
     * 
     */
    public Output<Integer> size() {
        return this.size;
    }
    /**
     * Ssl connection mode for the pool.
     * 
     */
    @Export(name="sslMode", refs={String.class}, tree="[0]")
    private Output<String> sslMode;

    /**
     * @return Ssl connection mode for the pool.
     * 
     */
    public Output<String> sslMode() {
        return this.sslMode;
    }
    /**
     * Connection URI to the pool.
     * 
     */
    @Export(name="uri", refs={String.class}, tree="[0]")
    private Output<String> uri;

    /**
     * @return Connection URI to the pool.
     * 
     */
    public Output<String> uri() {
        return this.uri;
    }
    /**
     * Database user authorized to connect to the pool, if none all the users are allowed.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return Database user authorized to connect to the pool, if none all the users are allowed.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PostgresSqlConnectionPool(java.lang.String name) {
        this(name, PostgresSqlConnectionPoolArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PostgresSqlConnectionPool(java.lang.String name, PostgresSqlConnectionPoolArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PostgresSqlConnectionPool(java.lang.String name, PostgresSqlConnectionPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/postgresSqlConnectionPool:PostgresSqlConnectionPool", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private PostgresSqlConnectionPool(java.lang.String name, Output<java.lang.String> id, @Nullable PostgresSqlConnectionPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("ovh:CloudProjectDatabase/postgresSqlConnectionPool:PostgresSqlConnectionPool", name, state, makeResourceOptions(options, id), false);
    }

    private static PostgresSqlConnectionPoolArgs makeArgs(PostgresSqlConnectionPoolArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? PostgresSqlConnectionPoolArgs.Empty : args;
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
    public static PostgresSqlConnectionPool get(java.lang.String name, Output<java.lang.String> id, @Nullable PostgresSqlConnectionPoolState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PostgresSqlConnectionPool(name, id, state, options);
    }
}