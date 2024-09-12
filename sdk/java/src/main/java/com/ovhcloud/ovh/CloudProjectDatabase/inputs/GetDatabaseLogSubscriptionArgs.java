// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseLogSubscriptionArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseLogSubscriptionArgs Empty = new GetDatabaseLogSubscriptionArgs();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    @Import(name="engine", required=true)
    private Output<String> engine;

    /**
     * @return The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
     * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }

    /**
     * Id of the log subscription.
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return Id of the log subscription.
     * 
     */
    public Output<String> id() {
        return this.id;
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetDatabaseLogSubscriptionArgs() {}

    private GetDatabaseLogSubscriptionArgs(GetDatabaseLogSubscriptionArgs $) {
        this.clusterId = $.clusterId;
        this.engine = $.engine;
        this.id = $.id;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseLogSubscriptionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseLogSubscriptionArgs $;

        public Builder() {
            $ = new GetDatabaseLogSubscriptionArgs();
        }

        public Builder(GetDatabaseLogSubscriptionArgs defaults) {
            $ = new GetDatabaseLogSubscriptionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param engine The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The database engine for which you want to retrieve a subscription. To get a full list of available engine visit.
         * [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases).
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param id Id of the log subscription.
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Id of the log subscription.
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetDatabaseLogSubscriptionArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionArgs", "clusterId");
            }
            if ($.engine == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionArgs", "engine");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionArgs", "id");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetDatabaseLogSubscriptionArgs", "serviceName");
            }
            return $;
        }
    }

}