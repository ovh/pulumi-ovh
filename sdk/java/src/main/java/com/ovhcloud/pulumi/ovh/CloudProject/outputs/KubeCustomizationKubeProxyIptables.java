// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class KubeCustomizationKubeProxyIptables {
    /**
     * @return Period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. `PT60S`). Must be greater than 0.
     * 
     */
    private @Nullable String minSyncPeriod;
    /**
     * @return Minimum period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. `PT60S`).
     * 
     */
    private @Nullable String syncPeriod;

    private KubeCustomizationKubeProxyIptables() {}
    /**
     * @return Period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. `PT60S`). Must be greater than 0.
     * 
     */
    public Optional<String> minSyncPeriod() {
        return Optional.ofNullable(this.minSyncPeriod);
    }
    /**
     * @return Minimum period that iptables rules are refreshed, in [RFC3339](https://www.rfc-editor.org/rfc/rfc3339) duration format (e.g. `PT60S`).
     * 
     */
    public Optional<String> syncPeriod() {
        return Optional.ofNullable(this.syncPeriod);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(KubeCustomizationKubeProxyIptables defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String minSyncPeriod;
        private @Nullable String syncPeriod;
        public Builder() {}
        public Builder(KubeCustomizationKubeProxyIptables defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.minSyncPeriod = defaults.minSyncPeriod;
    	      this.syncPeriod = defaults.syncPeriod;
        }

        @CustomType.Setter
        public Builder minSyncPeriod(@Nullable String minSyncPeriod) {

            this.minSyncPeriod = minSyncPeriod;
            return this;
        }
        @CustomType.Setter
        public Builder syncPeriod(@Nullable String syncPeriod) {

            this.syncPeriod = syncPeriod;
            return this;
        }
        public KubeCustomizationKubeProxyIptables build() {
            final var _resultValue = new KubeCustomizationKubeProxyIptables();
            _resultValue.minSyncPeriod = minSyncPeriod;
            _resultValue.syncPeriod = syncPeriod;
            return _resultValue;
        }
    }
}