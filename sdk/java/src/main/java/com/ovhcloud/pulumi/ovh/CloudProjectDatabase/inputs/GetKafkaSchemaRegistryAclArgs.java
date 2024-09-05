// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetKafkaSchemaRegistryAclArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetKafkaSchemaRegistryAclArgs Empty = new GetKafkaSchemaRegistryAclArgs();

    /**
     * Cluster ID
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * Schema registry ACL ID
     * 
     */
    @Import(name="id", required=true)
    private Output<String> id;

    /**
     * @return Schema registry ACL ID
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

    private GetKafkaSchemaRegistryAclArgs() {}

    private GetKafkaSchemaRegistryAclArgs(GetKafkaSchemaRegistryAclArgs $) {
        this.clusterId = $.clusterId;
        this.id = $.id;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKafkaSchemaRegistryAclArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKafkaSchemaRegistryAclArgs $;

        public Builder() {
            $ = new GetKafkaSchemaRegistryAclArgs();
        }

        public Builder(GetKafkaSchemaRegistryAclArgs defaults) {
            $ = new GetKafkaSchemaRegistryAclArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param id Schema registry ACL ID
         * 
         * @return builder
         * 
         */
        public Builder id(Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Schema registry ACL ID
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

        public GetKafkaSchemaRegistryAclArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetKafkaSchemaRegistryAclArgs", "clusterId");
            }
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetKafkaSchemaRegistryAclArgs", "id");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetKafkaSchemaRegistryAclArgs", "serviceName");
            }
            return $;
        }
    }

}