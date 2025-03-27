// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetLoadBalancerPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetLoadBalancerPlainArgs Empty = new GetLoadBalancerPlainArgs();

    /**
     * ID of the loadbalancer
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return ID of the loadbalancer
     * 
     */
    public String id() {
        return this.id;
    }

    /**
     * Region of the loadbalancer.
     * 
     */
    @Import(name="regionName", required=true)
    private String regionName;

    /**
     * @return Region of the loadbalancer.
     * 
     */
    public String regionName() {
        return this.regionName;
    }

    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetLoadBalancerPlainArgs() {}

    private GetLoadBalancerPlainArgs(GetLoadBalancerPlainArgs $) {
        this.id = $.id;
        this.regionName = $.regionName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetLoadBalancerPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetLoadBalancerPlainArgs $;

        public Builder() {
            $ = new GetLoadBalancerPlainArgs();
        }

        public Builder(GetLoadBalancerPlainArgs defaults) {
            $ = new GetLoadBalancerPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id ID of the loadbalancer
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        /**
         * @param regionName Region of the loadbalancer.
         * 
         * @return builder
         * 
         */
        public Builder regionName(String regionName) {
            $.regionName = regionName;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetLoadBalancerPlainArgs build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetLoadBalancerPlainArgs", "id");
            }
            if ($.regionName == null) {
                throw new MissingRequiredPropertyException("GetLoadBalancerPlainArgs", "regionName");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetLoadBalancerPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
