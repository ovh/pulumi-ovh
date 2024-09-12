// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.IpLoadBalancing.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetVrackNetworkResult {
    /**
     * @return Human readable name for your vrack network
     * 
     */
    private String displayName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
     * 
     */
    private String natIp;
    private String serviceName;
    /**
     * @return IP block of the private network in the vRack
     * 
     */
    private String subnet;
    /**
     * @return VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
     * 
     */
    private Integer vlan;
    private Integer vrackNetworkId;

    private GetVrackNetworkResult() {}
    /**
     * @return Human readable name for your vrack network
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return An IP block used as a pool of IPs by this Load Balancer to connect to the servers in this private network. The blck must be in the private network and reserved for the Load Balancer
     * 
     */
    public String natIp() {
        return this.natIp;
    }
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return IP block of the private network in the vRack
     * 
     */
    public String subnet() {
        return this.subnet;
    }
    /**
     * @return VLAN of the private network in the vRack. 0 if the private network is not in a VLAN
     * 
     */
    public Integer vlan() {
        return this.vlan;
    }
    public Integer vrackNetworkId() {
        return this.vrackNetworkId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetVrackNetworkResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String displayName;
        private String id;
        private String natIp;
        private String serviceName;
        private String subnet;
        private Integer vlan;
        private Integer vrackNetworkId;
        public Builder() {}
        public Builder(GetVrackNetworkResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.displayName = defaults.displayName;
    	      this.id = defaults.id;
    	      this.natIp = defaults.natIp;
    	      this.serviceName = defaults.serviceName;
    	      this.subnet = defaults.subnet;
    	      this.vlan = defaults.vlan;
    	      this.vrackNetworkId = defaults.vrackNetworkId;
        }

        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder natIp(String natIp) {
            if (natIp == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "natIp");
            }
            this.natIp = natIp;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder subnet(String subnet) {
            if (subnet == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "subnet");
            }
            this.subnet = subnet;
            return this;
        }
        @CustomType.Setter
        public Builder vlan(Integer vlan) {
            if (vlan == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "vlan");
            }
            this.vlan = vlan;
            return this;
        }
        @CustomType.Setter
        public Builder vrackNetworkId(Integer vrackNetworkId) {
            if (vrackNetworkId == null) {
              throw new MissingRequiredPropertyException("GetVrackNetworkResult", "vrackNetworkId");
            }
            this.vrackNetworkId = vrackNetworkId;
            return this;
        }
        public GetVrackNetworkResult build() {
            final var _resultValue = new GetVrackNetworkResult();
            _resultValue.displayName = displayName;
            _resultValue.id = id;
            _resultValue.natIp = natIp;
            _resultValue.serviceName = serviceName;
            _resultValue.subnet = subnet;
            _resultValue.vlan = vlan;
            _resultValue.vrackNetworkId = vrackNetworkId;
            return _resultValue;
        }
    }
}