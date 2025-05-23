// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseIntegrationPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseIntegrationPlainArgs Empty = new GetDatabaseIntegrationPlainArgs();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
     * 
     */
    @Import(name="engine", required=true)
    private String engine;

    /**
     * @return The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
     * 
     */
    public String engine() {
        return this.engine;
    }

    /**
     * Integration ID
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return Integration ID
     * 
     */
    public String id() {
        return this.id;
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetDatabaseIntegrationPlainArgs() {}

    private GetDatabaseIntegrationPlainArgs(GetDatabaseIntegrationPlainArgs $) {
        this.clusterId = $.clusterId;
        this.engine = $.engine;
        this.id = $.id;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseIntegrationPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseIntegrationPlainArgs $;

        public Builder() {
            $ = new GetDatabaseIntegrationPlainArgs();
        }

        public Builder(GetDatabaseIntegrationPlainArgs defaults) {
            $ = new GetDatabaseIntegrationPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param engine The engine of the database cluster you want to add. You can find the complete list of available engine in the [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param id Integration ID
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetDatabaseIntegrationPlainArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationPlainArgs", "clusterId");
            }
            if ($.engine == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationPlainArgs", "engine");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationPlainArgs", "id");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
