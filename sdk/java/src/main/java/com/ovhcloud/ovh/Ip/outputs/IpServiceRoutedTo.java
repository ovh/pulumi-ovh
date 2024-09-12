// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Ip.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class IpServiceRoutedTo {
    /**
     * @return service name
     * 
     */
    private @Nullable String serviceName;

    private IpServiceRoutedTo() {}
    /**
     * @return service name
     * 
     */
    public Optional<String> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(IpServiceRoutedTo defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String serviceName;
        public Builder() {}
        public Builder(IpServiceRoutedTo defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.serviceName = defaults.serviceName;
        }

        @CustomType.Setter
        public Builder serviceName(@Nullable String serviceName) {

            this.serviceName = serviceName;
            return this;
        }
        public IpServiceRoutedTo build() {
            final var _resultValue = new IpServiceRoutedTo();
            _resultValue.serviceName = serviceName;
            return _resultValue;
        }
    }
}