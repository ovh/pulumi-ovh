// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetDatabaseIntegrationsArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDatabaseIntegrationsArgs Empty = new GetDatabaseIntegrationsArgs();

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
     * The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
     * 
     */
    @Import(name="engine", required=true)
    private Output<String> engine;

    /**
     * @return The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
     * 
     */
    public Output<String> engine() {
        return this.engine;
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetDatabaseIntegrationsArgs() {}

    private GetDatabaseIntegrationsArgs(GetDatabaseIntegrationsArgs $) {
        this.clusterId = $.clusterId;
        this.engine = $.engine;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDatabaseIntegrationsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDatabaseIntegrationsArgs $;

        public Builder() {
            $ = new GetDatabaseIntegrationsArgs();
        }

        public Builder(GetDatabaseIntegrationsArgs defaults) {
            $ = new GetDatabaseIntegrationsArgs(Objects.requireNonNull(defaults));
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
         * @param engine The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
         * 
         * @return builder
         * 
         */
        public Builder engine(Output<String> engine) {
            $.engine = engine;
            return this;
        }

        /**
         * @param engine The engine of the database cluster you want to list integrations. To get a full list of available engine visit: [public documentation](https://docs.ovh.com/gb/en/publiccloud/databases). All engines available exept `mongodb`
         * 
         * @return builder
         * 
         */
        public Builder engine(String engine) {
            return engine(Output.of(engine));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetDatabaseIntegrationsArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationsArgs", "clusterId");
            }
            if ($.engine == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationsArgs", "engine");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetDatabaseIntegrationsArgs", "serviceName");
            }
            return $;
        }
    }

}
