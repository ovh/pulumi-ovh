// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceNetworkPrivateNetwork {
    /**
     * @return Network ID
     * 
     */
    private @Nullable String id;
    /**
     * @return Existing subnet ID
     * * network_create - (Optional, Forces new resource) Information to create a new private network
     * 
     */
    private @Nullable String subnetId;

    private InstanceNetworkPrivateNetwork() {}
    /**
     * @return Network ID
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return Existing subnet ID
     * * network_create - (Optional, Forces new resource) Information to create a new private network
     * 
     */
    public Optional<String> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceNetworkPrivateNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String id;
        private @Nullable String subnetId;
        public Builder() {}
        public Builder(InstanceNetworkPrivateNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.subnetId = defaults.subnetId;
        }

        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder subnetId(@Nullable String subnetId) {

            this.subnetId = subnetId;
            return this;
        }
        public InstanceNetworkPrivateNetwork build() {
            final var _resultValue = new InstanceNetworkPrivateNetwork();
            _resultValue.id = id;
            _resultValue.subnetId = subnetId;
            return _resultValue;
        }
    }
}
