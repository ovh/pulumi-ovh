// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LoadBalancerFloatingIp {
    /**
     * @return ID of the resource
     * 
     */
    private @Nullable String id;
    /**
     * @return IP Address of the resource
     * 
     */
    private @Nullable String ip;

    private LoadBalancerFloatingIp() {}
    /**
     * @return ID of the resource
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return IP Address of the resource
     * 
     */
    public Optional<String> ip() {
        return Optional.ofNullable(this.ip);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadBalancerFloatingIp defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String id;
        private @Nullable String ip;
        public Builder() {}
        public Builder(LoadBalancerFloatingIp defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
        }

        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ip(@Nullable String ip) {

            this.ip = ip;
            return this;
        }
        public LoadBalancerFloatingIp build() {
            final var _resultValue = new LoadBalancerFloatingIp();
            _resultValue.id = id;
            _resultValue.ip = ip;
            return _resultValue;
        }
    }
}
