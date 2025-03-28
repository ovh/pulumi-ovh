// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.LoadBalancerListenerPool;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class LoadBalancerListener {
    /**
     * @return The allowed CIDRs
     * 
     */
    private @Nullable List<String> allowedCidrs;
    /**
     * @return The description of the listener
     * 
     */
    private @Nullable String description;
    /**
     * @return Name of the listener
     * 
     */
    private @Nullable String name;
    /**
     * @return Listener pool
     * 
     */
    private @Nullable LoadBalancerListenerPool pool;
    /**
     * @return Listener port
     * 
     */
    private Double port;
    /**
     * @return Protocol for the listener
     * 
     */
    private String protocol;
    /**
     * @return Secret ID to get certificate for SSL listener creation
     * 
     */
    private @Nullable String secretId;
    /**
     * @return Timeout client data of the listener
     * 
     */
    private @Nullable Double timeoutClientData;
    /**
     * @return Timeout member data of the listener
     * 
     */
    private @Nullable Double timeoutMemberData;
    /**
     * @return TLS versions of the listener
     * 
     */
    private @Nullable List<String> tlsVersions;

    private LoadBalancerListener() {}
    /**
     * @return The allowed CIDRs
     * 
     */
    public List<String> allowedCidrs() {
        return this.allowedCidrs == null ? List.of() : this.allowedCidrs;
    }
    /**
     * @return The description of the listener
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return Name of the listener
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return Listener pool
     * 
     */
    public Optional<LoadBalancerListenerPool> pool() {
        return Optional.ofNullable(this.pool);
    }
    /**
     * @return Listener port
     * 
     */
    public Double port() {
        return this.port;
    }
    /**
     * @return Protocol for the listener
     * 
     */
    public String protocol() {
        return this.protocol;
    }
    /**
     * @return Secret ID to get certificate for SSL listener creation
     * 
     */
    public Optional<String> secretId() {
        return Optional.ofNullable(this.secretId);
    }
    /**
     * @return Timeout client data of the listener
     * 
     */
    public Optional<Double> timeoutClientData() {
        return Optional.ofNullable(this.timeoutClientData);
    }
    /**
     * @return Timeout member data of the listener
     * 
     */
    public Optional<Double> timeoutMemberData() {
        return Optional.ofNullable(this.timeoutMemberData);
    }
    /**
     * @return TLS versions of the listener
     * 
     */
    public List<String> tlsVersions() {
        return this.tlsVersions == null ? List.of() : this.tlsVersions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(LoadBalancerListener defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> allowedCidrs;
        private @Nullable String description;
        private @Nullable String name;
        private @Nullable LoadBalancerListenerPool pool;
        private Double port;
        private String protocol;
        private @Nullable String secretId;
        private @Nullable Double timeoutClientData;
        private @Nullable Double timeoutMemberData;
        private @Nullable List<String> tlsVersions;
        public Builder() {}
        public Builder(LoadBalancerListener defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedCidrs = defaults.allowedCidrs;
    	      this.description = defaults.description;
    	      this.name = defaults.name;
    	      this.pool = defaults.pool;
    	      this.port = defaults.port;
    	      this.protocol = defaults.protocol;
    	      this.secretId = defaults.secretId;
    	      this.timeoutClientData = defaults.timeoutClientData;
    	      this.timeoutMemberData = defaults.timeoutMemberData;
    	      this.tlsVersions = defaults.tlsVersions;
        }

        @CustomType.Setter
        public Builder allowedCidrs(@Nullable List<String> allowedCidrs) {

            this.allowedCidrs = allowedCidrs;
            return this;
        }
        public Builder allowedCidrs(String... allowedCidrs) {
            return allowedCidrs(List.of(allowedCidrs));
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder pool(@Nullable LoadBalancerListenerPool pool) {

            this.pool = pool;
            return this;
        }
        @CustomType.Setter
        public Builder port(Double port) {
            if (port == null) {
              throw new MissingRequiredPropertyException("LoadBalancerListener", "port");
            }
            this.port = port;
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            if (protocol == null) {
              throw new MissingRequiredPropertyException("LoadBalancerListener", "protocol");
            }
            this.protocol = protocol;
            return this;
        }
        @CustomType.Setter
        public Builder secretId(@Nullable String secretId) {

            this.secretId = secretId;
            return this;
        }
        @CustomType.Setter
        public Builder timeoutClientData(@Nullable Double timeoutClientData) {

            this.timeoutClientData = timeoutClientData;
            return this;
        }
        @CustomType.Setter
        public Builder timeoutMemberData(@Nullable Double timeoutMemberData) {

            this.timeoutMemberData = timeoutMemberData;
            return this;
        }
        @CustomType.Setter
        public Builder tlsVersions(@Nullable List<String> tlsVersions) {

            this.tlsVersions = tlsVersions;
            return this;
        }
        public Builder tlsVersions(String... tlsVersions) {
            return tlsVersions(List.of(tlsVersions));
        }
        public LoadBalancerListener build() {
            final var _resultValue = new LoadBalancerListener();
            _resultValue.allowedCidrs = allowedCidrs;
            _resultValue.description = description;
            _resultValue.name = name;
            _resultValue.pool = pool;
            _resultValue.port = port;
            _resultValue.protocol = protocol;
            _resultValue.secretId = secretId;
            _resultValue.timeoutClientData = timeoutClientData;
            _resultValue.timeoutMemberData = timeoutMemberData;
            _resultValue.tlsVersions = tlsVersions;
            return _resultValue;
        }
    }
}
