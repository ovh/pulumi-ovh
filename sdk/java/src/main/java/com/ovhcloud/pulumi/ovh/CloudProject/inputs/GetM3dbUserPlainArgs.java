// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetM3dbUserPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetM3dbUserPlainArgs Empty = new GetM3dbUserPlainArgs();

    /**
     * Cluster ID
     * 
     */
    @Import(name="clusterId", required=true)
    private String clusterId;

    /**
     * @return Cluster ID
     * 
     */
    public String clusterId() {
        return this.clusterId;
    }

    /**
     * Name of the user.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name of the user.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetM3dbUserPlainArgs() {}

    private GetM3dbUserPlainArgs(GetM3dbUserPlainArgs $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetM3dbUserPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetM3dbUserPlainArgs $;

        public Builder() {
            $ = new GetM3dbUserPlainArgs();
        }

        public Builder(GetM3dbUserPlainArgs defaults) {
            $ = new GetM3dbUserPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param name Name of the user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
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
            $.serviceName = serviceName;
            return this;
        }

        public GetM3dbUserPlainArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("GetM3dbUserPlainArgs", "clusterId");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetM3dbUserPlainArgs", "name");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetM3dbUserPlainArgs", "serviceName");
            }
            return $;
        }
    }

}