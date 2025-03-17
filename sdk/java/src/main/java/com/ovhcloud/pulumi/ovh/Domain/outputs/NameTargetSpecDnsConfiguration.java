// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain.outputs;

import com.ovhcloud.pulumi.ovh.Domain.outputs.NameTargetSpecDnsConfigurationNameServer;
import com.pulumi.core.annotations.CustomType;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class NameTargetSpecDnsConfiguration {
    /**
     * @return The name servers to update
     * 
     */
    private @Nullable List<NameTargetSpecDnsConfigurationNameServer> nameServers;

    private NameTargetSpecDnsConfiguration() {}
    /**
     * @return The name servers to update
     * 
     */
    public List<NameTargetSpecDnsConfigurationNameServer> nameServers() {
        return this.nameServers == null ? List.of() : this.nameServers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(NameTargetSpecDnsConfiguration defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<NameTargetSpecDnsConfigurationNameServer> nameServers;
        public Builder() {}
        public Builder(NameTargetSpecDnsConfiguration defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.nameServers = defaults.nameServers;
        }

        @CustomType.Setter
        public Builder nameServers(@Nullable List<NameTargetSpecDnsConfigurationNameServer> nameServers) {

            this.nameServers = nameServers;
            return this;
        }
        public Builder nameServers(NameTargetSpecDnsConfigurationNameServer... nameServers) {
            return nameServers(List.of(nameServers));
        }
        public NameTargetSpecDnsConfiguration build() {
            final var _resultValue = new NameTargetSpecDnsConfiguration();
            _resultValue.nameServers = nameServers;
            return _resultValue;
        }
    }
}
