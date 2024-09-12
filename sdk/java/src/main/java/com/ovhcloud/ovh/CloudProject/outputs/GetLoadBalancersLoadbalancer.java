// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject.outputs;

import com.ovhcloud.ovh.CloudProject.outputs.GetLoadBalancersLoadbalancerFloatingIp;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetLoadBalancersLoadbalancer {
    /**
     * @return Date of creation of the loadbalancer
     * 
     */
    private String createdAt;
    /**
     * @return ID of the flavor
     * 
     */
    private String flavorId;
    /**
     * @return Information about the floating IP
     * 
     */
    private GetLoadBalancersLoadbalancerFloatingIp floatingIp;
    /**
     * @return ID of the floating IP
     * 
     */
    private String id;
    /**
     * @return Name of the loadbalancer
     * 
     */
    private String name;
    /**
     * @return Operating status of the loadbalancer
     * 
     */
    private String operatingStatus;
    /**
     * @return Provisioning status of the loadbalancer
     * 
     */
    private String provisioningStatus;
    /**
     * @return Region of the loadbalancer
     * 
     */
    private String region;
    /**
     * @return Last update date of the loadbalancer
     * 
     */
    private String updatedAt;
    /**
     * @return IP address of the Virtual IP
     * 
     */
    private String vipAddress;
    /**
     * @return Openstack ID of the network for the Virtual IP
     * 
     */
    private String vipNetworkId;
    /**
     * @return ID of the subnet for the Virtual IP
     * 
     */
    private String vipSubnetId;

    private GetLoadBalancersLoadbalancer() {}
    /**
     * @return Date of creation of the loadbalancer
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return ID of the flavor
     * 
     */
    public String flavorId() {
        return this.flavorId;
    }
    /**
     * @return Information about the floating IP
     * 
     */
    public GetLoadBalancersLoadbalancerFloatingIp floatingIp() {
        return this.floatingIp;
    }
    /**
     * @return ID of the floating IP
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Name of the loadbalancer
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Operating status of the loadbalancer
     * 
     */
    public String operatingStatus() {
        return this.operatingStatus;
    }
    /**
     * @return Provisioning status of the loadbalancer
     * 
     */
    public String provisioningStatus() {
        return this.provisioningStatus;
    }
    /**
     * @return Region of the loadbalancer
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return Last update date of the loadbalancer
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }
    /**
     * @return IP address of the Virtual IP
     * 
     */
    public String vipAddress() {
        return this.vipAddress;
    }
    /**
     * @return Openstack ID of the network for the Virtual IP
     * 
     */
    public String vipNetworkId() {
        return this.vipNetworkId;
    }
    /**
     * @return ID of the subnet for the Virtual IP
     * 
     */
    public String vipSubnetId() {
        return this.vipSubnetId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetLoadBalancersLoadbalancer defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String flavorId;
        private GetLoadBalancersLoadbalancerFloatingIp floatingIp;
        private String id;
        private String name;
        private String operatingStatus;
        private String provisioningStatus;
        private String region;
        private String updatedAt;
        private String vipAddress;
        private String vipNetworkId;
        private String vipSubnetId;
        public Builder() {}
        public Builder(GetLoadBalancersLoadbalancer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.flavorId = defaults.flavorId;
    	      this.floatingIp = defaults.floatingIp;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.operatingStatus = defaults.operatingStatus;
    	      this.provisioningStatus = defaults.provisioningStatus;
    	      this.region = defaults.region;
    	      this.updatedAt = defaults.updatedAt;
    	      this.vipAddress = defaults.vipAddress;
    	      this.vipNetworkId = defaults.vipNetworkId;
    	      this.vipSubnetId = defaults.vipSubnetId;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder flavorId(String flavorId) {
            if (flavorId == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "flavorId");
            }
            this.flavorId = flavorId;
            return this;
        }
        @CustomType.Setter
        public Builder floatingIp(GetLoadBalancersLoadbalancerFloatingIp floatingIp) {
            if (floatingIp == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "floatingIp");
            }
            this.floatingIp = floatingIp;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder operatingStatus(String operatingStatus) {
            if (operatingStatus == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "operatingStatus");
            }
            this.operatingStatus = operatingStatus;
            return this;
        }
        @CustomType.Setter
        public Builder provisioningStatus(String provisioningStatus) {
            if (provisioningStatus == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "provisioningStatus");
            }
            this.provisioningStatus = provisioningStatus;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        @CustomType.Setter
        public Builder vipAddress(String vipAddress) {
            if (vipAddress == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "vipAddress");
            }
            this.vipAddress = vipAddress;
            return this;
        }
        @CustomType.Setter
        public Builder vipNetworkId(String vipNetworkId) {
            if (vipNetworkId == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "vipNetworkId");
            }
            this.vipNetworkId = vipNetworkId;
            return this;
        }
        @CustomType.Setter
        public Builder vipSubnetId(String vipSubnetId) {
            if (vipSubnetId == null) {
              throw new MissingRequiredPropertyException("GetLoadBalancersLoadbalancer", "vipSubnetId");
            }
            this.vipSubnetId = vipSubnetId;
            return this;
        }
        public GetLoadBalancersLoadbalancer build() {
            final var _resultValue = new GetLoadBalancersLoadbalancer();
            _resultValue.createdAt = createdAt;
            _resultValue.flavorId = flavorId;
            _resultValue.floatingIp = floatingIp;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.operatingStatus = operatingStatus;
            _resultValue.provisioningStatus = provisioningStatus;
            _resultValue.region = region;
            _resultValue.updatedAt = updatedAt;
            _resultValue.vipAddress = vipAddress;
            _resultValue.vipNetworkId = vipNetworkId;
            _resultValue.vipSubnetId = vipSubnetId;
            return _resultValue;
        }
    }
}