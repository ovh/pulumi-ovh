// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseLogSubscriptionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseLogSubscriptionPlainArgs Empty = new GetDatabaseLogSubscriptionPlainArgs();

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
     * The database engine for which you want to retrieve a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Import(name="engine", required=true)
    private String engine;

    /**
     * @return The database engine for which you want to retrieve a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    public String engine() {
        return this.engine;
    }

    /**
     * Id of the log subscription.
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return Id of the log subscription.
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

    private GetDatabaseLogSubscriptionPlainArgs() {}

    private GetDatabaseLogSubscriptionPlainArgs(GetDatabaseLogSubscriptionPlainArgs $) {
        this.clusterId = $.clusterId;
        this.engine = $.engine;
        this.id = $.id;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseLogSubscriptionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseLogSubscriptionPlainArgs $;

        public Builder() {
            $ = new GetDatabaseLogSubscriptionPlainArgs();
        }

        public Builder(GetDatabaseLogSubscriptionPlainArgs defaults) {
            $ = new GetDatabaseLogSubscriptionPlainArgs(Objects.requireNonNull(defaults));
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
         * @param engine The database engine for which you want to retrieve a subscription. To get a full list of available engine visit. [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param id Id of the log subscription.
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

        public GetDatabaseLogSubscriptionPlainArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionPlainArgs", "clusterId");
            }
            if ($.engine == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionPlainArgs", "engine");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionPlainArgs", "id");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
