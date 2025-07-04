// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivateFloatingIp;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivateFloatingIpCreate;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivateGateway;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivateGatewayCreate;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivateNetwork;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivateNetworkCreate;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceNetworkPrivate {
    /**
     * @return Existing floating IP
     * 
     */
    private @Nullable InstanceNetworkPrivateFloatingIp floatingIp;
    /**
     * @return Information to create a new floating IP
     * 
     */
    private @Nullable InstanceNetworkPrivateFloatingIpCreate floatingIpCreate;
    /**
     * @return Existing gateway
     * 
     */
    private @Nullable InstanceNetworkPrivateGateway gateway;
    /**
     * @return Information to create a new gateway
     * 
     */
    private @Nullable InstanceNetworkPrivateGatewayCreate gatewayCreate;
    /**
     * @return Instance IP in the private network
     * 
     */
    private @Nullable String ip;
    /**
     * @return Existing private network
     * 
     */
    private @Nullable InstanceNetworkPrivateNetwork network;
    /**
     * @return Information to create a new private network
     * 
     */
    private @Nullable InstanceNetworkPrivateNetworkCreate networkCreate;

    private InstanceNetworkPrivate() {}
    /**
     * @return Existing floating IP
     * 
     */
    public Optional<InstanceNetworkPrivateFloatingIp> floatingIp() {
        return Optional.ofNullable(this.floatingIp);
    }
    /**
     * @return Information to create a new floating IP
     * 
     */
    public Optional<InstanceNetworkPrivateFloatingIpCreate> floatingIpCreate() {
        return Optional.ofNullable(this.floatingIpCreate);
    }
    /**
     * @return Existing gateway
     * 
     */
    public Optional<InstanceNetworkPrivateGateway> gateway() {
        return Optional.ofNullable(this.gateway);
    }
    /**
     * @return Information to create a new gateway
     * 
     */
    public Optional<InstanceNetworkPrivateGatewayCreate> gatewayCreate() {
        return Optional.ofNullable(this.gatewayCreate);
    }
    /**
     * @return Instance IP in the private network
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }
    /**
     * @return Existing private network
     * 
     */
    public Optional<InstanceNetworkPrivateNetwork> network() {
        return Optional.ofNullable(this.network);
    }
    /**
     * @return Information to create a new private network
     * 
     */
    public Optional<InstanceNetworkPrivateNetworkCreate> networkCreate() {
        return Optional.ofNullable(this.networkCreate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceNetworkPrivate defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable InstanceNetworkPrivateFloatingIp floatingIp;
        private @Nullable InstanceNetworkPrivateFloatingIpCreate floatingIpCreate;
        private @Nullable InstanceNetworkPrivateGateway gateway;
        private @Nullable InstanceNetworkPrivateGatewayCreate gatewayCreate;
        private @Nullable String ip;
        private @Nullable InstanceNetworkPrivateNetwork network;
        private @Nullable InstanceNetworkPrivateNetworkCreate networkCreate;
        public Builder() {}
        public Builder(InstanceNetworkPrivate defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.floatingIp = defaults.floatingIp;
    	      this.floatingIpCreate = defaults.floatingIpCreate;
    	      this.gateway = defaults.gateway;
    	      this.gatewayCreate = defaults.gatewayCreate;
    	      this.ip = defaults.ip;
    	      this.network = defaults.network;
    	      this.networkCreate = defaults.networkCreate;
        }

        @CustomType.Setter
        public Builder floatingIp(@Nullable InstanceNetworkPrivateFloatingIp floatingIp) {

            this.floatingIp = floatingIp;
            return this;
        }
        @CustomType.Setter
        public Builder floatingIpCreate(@Nullable InstanceNetworkPrivateFloatingIpCreate floatingIpCreate) {

            this.floatingIpCreate = floatingIpCreate;
            return this;
        }
        @CustomType.Setter
        public Builder gateway(@Nullable InstanceNetworkPrivateGateway gateway) {

            this.gateway = gateway;
            return this;
        }
        @CustomType.Setter
        public Builder gatewayCreate(@Nullable InstanceNetworkPrivateGatewayCreate gatewayCreate) {

            this.gatewayCreate = gatewayCreate;
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {

            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder network(@Nullable InstanceNetworkPrivateNetwork network) {

            this.network = network;
            return this;
        }
        @CustomType.Setter
        public Builder networkCreate(@Nullable InstanceNetworkPrivateNetworkCreate networkCreate) {

            this.networkCreate = networkCreate;
            return this;
        }
        public InstanceNetworkPrivate build() {
            final var _resultValue = new InstanceNetworkPrivate();
            _resultValue.floatingIp = floatingIp;
            _resultValue.floatingIpCreate = floatingIpCreate;
            _resultValue.gateway = gateway;
            _resultValue.gatewayCreate = gatewayCreate;
            _resultValue.ip = ip;
            _resultValue.network = network;
            _resultValue.networkCreate = networkCreate;
            return _resultValue;
        }
    }
}
