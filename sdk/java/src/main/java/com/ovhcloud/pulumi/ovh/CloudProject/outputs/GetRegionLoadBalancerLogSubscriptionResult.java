// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRegionLoadBalancerLogSubscriptionResult {
    /**
     * @return The date of the subscription creation
     * 
     */
    private String createdAt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Router used for forwarding log
     * 
     */
    private String kind;
    /**
     * @return LDP service name
     * 
     */
    private String ldpServiceName;
    /**
     * @return Loadbalancer id to get the logs
     * 
     */
    private String loadbalancerId;
    /**
     * @return A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: &#34;GRA11&#34;.
     * 
     */
    private String regionName;
    /**
     * @return The resource name
     * 
     */
    private String resourceName;
    /**
     * @return The resource type
     * 
     */
    private String resourceType;
    /**
     * @return The id of the public cloud project.
     * 
     */
    private String serviceName;
    /**
     * @return Data stream id to use for the subscription
     * 
     */
    private String streamId;
    /**
     * @return The subscription id
     * 
     */
    private String subscriptionId;
    /**
     * @return The last update of the subscription
     * 
     */
    private String updatedAt;

    private GetRegionLoadBalancerLogSubscriptionResult() {}
    /**
     * @return The date of the subscription creation
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Router used for forwarding log
     * 
     */
    public String kind() {
        return this.kind;
    }
    /**
     * @return LDP service name
     * 
     */
    public String ldpServiceName() {
        return this.ldpServiceName;
    }
    /**
     * @return Loadbalancer id to get the logs
     * 
     */
    public String loadbalancerId() {
        return this.loadbalancerId;
    }
    /**
     * @return A valid OVHcloud public cloud region name in which the loadbalancer will be available. Ex.: &#34;GRA11&#34;.
     * 
     */
    public String regionName() {
        return this.regionName;
    }
    /**
     * @return The resource name
     * 
     */
    public String resourceName() {
        return this.resourceName;
    }
    /**
     * @return The resource type
     * 
     */
    public String resourceType() {
        return this.resourceType;
    }
    /**
     * @return The id of the public cloud project.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Data stream id to use for the subscription
     * 
     */
    public String streamId() {
        return this.streamId;
    }
    /**
     * @return The subscription id
     * 
     */
    public String subscriptionId() {
        return this.subscriptionId;
    }
    /**
     * @return The last update of the subscription
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRegionLoadBalancerLogSubscriptionResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private String kind;
        private String ldpServiceName;
        private String loadbalancerId;
        private String regionName;
        private String resourceName;
        private String resourceType;
        private String serviceName;
        private String streamId;
        private String subscriptionId;
        private String updatedAt;
        public Builder() {}
        public Builder(GetRegionLoadBalancerLogSubscriptionResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.kind = defaults.kind;
    	      this.ldpServiceName = defaults.ldpServiceName;
    	      this.loadbalancerId = defaults.loadbalancerId;
    	      this.regionName = defaults.regionName;
    	      this.resourceName = defaults.resourceName;
    	      this.resourceType = defaults.resourceType;
    	      this.serviceName = defaults.serviceName;
    	      this.streamId = defaults.streamId;
    	      this.subscriptionId = defaults.subscriptionId;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder kind(String kind) {
            if (kind == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "kind");
            }
            this.kind = kind;
            return this;
        }
        @CustomType.Setter
        public Builder ldpServiceName(String ldpServiceName) {
            if (ldpServiceName == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "ldpServiceName");
            }
            this.ldpServiceName = ldpServiceName;
            return this;
        }
        @CustomType.Setter
        public Builder loadbalancerId(String loadbalancerId) {
            if (loadbalancerId == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "loadbalancerId");
            }
            this.loadbalancerId = loadbalancerId;
            return this;
        }
        @CustomType.Setter
        public Builder regionName(String regionName) {
            if (regionName == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "regionName");
            }
            this.regionName = regionName;
            return this;
        }
        @CustomType.Setter
        public Builder resourceName(String resourceName) {
            if (resourceName == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "resourceName");
            }
            this.resourceName = resourceName;
            return this;
        }
        @CustomType.Setter
        public Builder resourceType(String resourceType) {
            if (resourceType == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "resourceType");
            }
            this.resourceType = resourceType;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder streamId(String streamId) {
            if (streamId == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "streamId");
            }
            this.streamId = streamId;
            return this;
        }
        @CustomType.Setter
        public Builder subscriptionId(String subscriptionId) {
            if (subscriptionId == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "subscriptionId");
            }
            this.subscriptionId = subscriptionId;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetRegionLoadBalancerLogSubscriptionResult", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        public GetRegionLoadBalancerLogSubscriptionResult build() {
            final var _resultValue = new GetRegionLoadBalancerLogSubscriptionResult();
            _resultValue.createdAt = createdAt;
            _resultValue.id = id;
            _resultValue.kind = kind;
            _resultValue.ldpServiceName = ldpServiceName;
            _resultValue.loadbalancerId = loadbalancerId;
            _resultValue.regionName = regionName;
            _resultValue.resourceName = resourceName;
            _resultValue.resourceType = resourceType;
            _resultValue.serviceName = serviceName;
            _resultValue.streamId = streamId;
            _resultValue.subscriptionId = subscriptionId;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
