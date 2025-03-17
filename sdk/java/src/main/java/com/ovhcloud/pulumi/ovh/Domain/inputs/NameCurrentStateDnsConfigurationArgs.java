// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain.inputs;

import com.ovhcloud.pulumi.ovh.Domain.inputs.NameCurrentStateDnsConfigurationNameServerArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NameCurrentStateDnsConfigurationArgs extends com.pulumi.resources.ResourceArgs {

    public static final NameCurrentStateDnsConfigurationArgs Empty = new NameCurrentStateDnsConfigurationArgs();

    /**
     * The type of DNS configuration of the domain
     * 
     */
    @Import(name="configurationType")
    private @Nullable Output<String> configurationType;

    /**
     * @return The type of DNS configuration of the domain
     * 
     */
    public Optional<Output<String>> configurationType() {
        return Optional.ofNullable(this.configurationType);
    }

    /**
     * Whether the registry supports IPv6 or not
     * 
     */
    @Import(name="glueRecordIpv6supported")
    private @Nullable Output<Boolean> glueRecordIpv6supported;

    /**
     * @return Whether the registry supports IPv6 or not
     * 
     */
    public Optional<Output<Boolean>> glueRecordIpv6supported() {
        return Optional.ofNullable(this.glueRecordIpv6supported);
    }

    /**
     * Whether the registry accepts hosts or not
     * 
     */
    @Import(name="hostSupported")
    private @Nullable Output<Boolean> hostSupported;

    /**
     * @return Whether the registry accepts hosts or not
     * 
     */
    public Optional<Output<Boolean>> hostSupported() {
        return Optional.ofNullable(this.hostSupported);
    }

    /**
     * The maximum number of name servers allowed by the registry
     * 
     */
    @Import(name="maxDns")
    private @Nullable Output<Double> maxDns;

    /**
     * @return The maximum number of name servers allowed by the registry
     * 
     */
    public Optional<Output<Double>> maxDns() {
        return Optional.ofNullable(this.maxDns);
    }

    /**
     * The minimum number of name servers allowed by the registry
     * 
     */
    @Import(name="minDns")
    private @Nullable Output<Double> minDns;

    /**
     * @return The minimum number of name servers allowed by the registry
     * 
     */
    public Optional<Output<Double>> minDns() {
        return Optional.ofNullable(this.minDns);
    }

    /**
     * The name servers used by the domain name
     * 
     */
    @Import(name="nameServers")
    private @Nullable Output<List<NameCurrentStateDnsConfigurationNameServerArgs>> nameServers;

    /**
     * @return The name servers used by the domain name
     * 
     */
    public Optional<Output<List<NameCurrentStateDnsConfigurationNameServerArgs>>> nameServers() {
        return Optional.ofNullable(this.nameServers);
    }

    private NameCurrentStateDnsConfigurationArgs() {}

    private NameCurrentStateDnsConfigurationArgs(NameCurrentStateDnsConfigurationArgs $) {
        this.configurationType = $.configurationType;
        this.glueRecordIpv6supported = $.glueRecordIpv6supported;
        this.hostSupported = $.hostSupported;
        this.maxDns = $.maxDns;
        this.minDns = $.minDns;
        this.nameServers = $.nameServers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NameCurrentStateDnsConfigurationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NameCurrentStateDnsConfigurationArgs $;

        public Builder() {
            $ = new NameCurrentStateDnsConfigurationArgs();
        }

        public Builder(NameCurrentStateDnsConfigurationArgs defaults) {
            $ = new NameCurrentStateDnsConfigurationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param configurationType The type of DNS configuration of the domain
         * 
         * @return builder
         * 
         */
        public Builder configurationType(@Nullable Output<String> configurationType) {
            $.configurationType = configurationType;
            return this;
        }

        /**
         * @param configurationType The type of DNS configuration of the domain
         * 
         * @return builder
         * 
         */
        public Builder configurationType(String configurationType) {
            return configurationType(Output.of(configurationType));
        }

        /**
         * @param glueRecordIpv6supported Whether the registry supports IPv6 or not
         * 
         * @return builder
         * 
         */
        public Builder glueRecordIpv6supported(@Nullable Output<Boolean> glueRecordIpv6supported) {
            $.glueRecordIpv6supported = glueRecordIpv6supported;
            return this;
        }

        /**
         * @param glueRecordIpv6supported Whether the registry supports IPv6 or not
         * 
         * @return builder
         * 
         */
        public Builder glueRecordIpv6supported(Boolean glueRecordIpv6supported) {
            return glueRecordIpv6supported(Output.of(glueRecordIpv6supported));
        }

        /**
         * @param hostSupported Whether the registry accepts hosts or not
         * 
         * @return builder
         * 
         */
        public Builder hostSupported(@Nullable Output<Boolean> hostSupported) {
            $.hostSupported = hostSupported;
            return this;
        }

        /**
         * @param hostSupported Whether the registry accepts hosts or not
         * 
         * @return builder
         * 
         */
        public Builder hostSupported(Boolean hostSupported) {
            return hostSupported(Output.of(hostSupported));
        }

        /**
         * @param maxDns The maximum number of name servers allowed by the registry
         * 
         * @return builder
         * 
         */
        public Builder maxDns(@Nullable Output<Double> maxDns) {
            $.maxDns = maxDns;
            return this;
        }

        /**
         * @param maxDns The maximum number of name servers allowed by the registry
         * 
         * @return builder
         * 
         */
        public Builder maxDns(Double maxDns) {
            return maxDns(Output.of(maxDns));
        }

        /**
         * @param minDns The minimum number of name servers allowed by the registry
         * 
         * @return builder
         * 
         */
        public Builder minDns(@Nullable Output<Double> minDns) {
            $.minDns = minDns;
            return this;
        }

        /**
         * @param minDns The minimum number of name servers allowed by the registry
         * 
         * @return builder
         * 
         */
        public Builder minDns(Double minDns) {
            return minDns(Output.of(minDns));
        }

        /**
         * @param nameServers The name servers used by the domain name
         * 
         * @return builder
         * 
         */
        public Builder nameServers(@Nullable Output<List<NameCurrentStateDnsConfigurationNameServerArgs>> nameServers) {
            $.nameServers = nameServers;
            return this;
        }

        /**
         * @param nameServers The name servers used by the domain name
         * 
         * @return builder
         * 
         */
        public Builder nameServers(List<NameCurrentStateDnsConfigurationNameServerArgs> nameServers) {
            return nameServers(Output.of(nameServers));
        }

        /**
         * @param nameServers The name servers used by the domain name
         * 
         * @return builder
         * 
         */
        public Builder nameServers(NameCurrentStateDnsConfigurationNameServerArgs... nameServers) {
            return nameServers(List.of(nameServers));
        }

        public NameCurrentStateDnsConfigurationArgs build() {
            return $;
        }
    }

}
