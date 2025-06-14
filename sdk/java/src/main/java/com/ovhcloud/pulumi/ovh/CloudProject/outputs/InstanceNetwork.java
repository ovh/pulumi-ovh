// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.InstanceNetworkPrivate;
import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class InstanceNetwork {
    /**
     * @return Private network information
     * 
     */
    private @Nullable InstanceNetworkPrivate private_;
    /**
     * @return Set the new instance as public
     * 
     */
    private @Nullable Boolean public_;

    private InstanceNetwork() {}
    /**
     * @return Private network information
     * 
     */
    public Optional<InstanceNetworkPrivate> private_() {
        return Optional.ofNullable(this.private_);
    }
    /**
     * @return Set the new instance as public
     * 
     */
    public Optional<Boolean> public_() {
        return Optional.ofNullable(this.public_);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(InstanceNetwork defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable InstanceNetworkPrivate private_;
        private @Nullable Boolean public_;
        public Builder() {}
        public Builder(InstanceNetwork defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.private_ = defaults.private_;
    	      this.public_ = defaults.public_;
        }

        @CustomType.Setter("private")
        public Builder private_(@Nullable InstanceNetworkPrivate private_) {

            this.private_ = private_;
            return this;
        }
        @CustomType.Setter("public")
        public Builder public_(@Nullable Boolean public_) {

            this.public_ = public_;
            return this;
        }
        public InstanceNetwork build() {
            final var _resultValue = new InstanceNetwork();
            _resultValue.private_ = private_;
            _resultValue.public_ = public_;
            return _resultValue;
        }
    }
}
