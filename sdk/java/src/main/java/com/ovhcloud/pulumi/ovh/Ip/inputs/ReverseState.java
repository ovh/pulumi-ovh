// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Ip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ReverseState extends com.pulumi.resources.ResourceArgs {

    public static final ReverseState Empty = new ReverseState();

    /**
     * The IP to set the reverse of
     * 
     */
    @Import(name="ReverseIp")
    private @Nullable Output<String> ReverseIp;

    /**
     * @return The IP to set the reverse of
     * 
     */
    public Optional<Output<String>> ReverseIp() {
        return Optional.ofNullable(this.ReverseIp);
    }

    /**
     * The value of the reverse
     * 
     */
    @Import(name="ReverseValue")
    private @Nullable Output<String> ReverseValue;

    /**
     * @return The value of the reverse
     * 
     */
    public Optional<Output<String>> ReverseValue() {
        return Optional.ofNullable(this.ReverseValue);
    }

    /**
     * The IP block to which the IP belongs
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return The IP block to which the IP belongs
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    private ReverseState() {}

    private ReverseState(ReverseState $) {
        this.ReverseIp = $.ReverseIp;
        this.ReverseValue = $.ReverseValue;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ReverseState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ReverseState $;

        public Builder() {
            $ = new ReverseState();
        }

        public Builder(ReverseState defaults) {
            $ = new ReverseState(Objects.requireNonNull(defaults));
        }

        /**
         * @param ReverseIp The IP to set the reverse of
         * 
         * @return builder
         * 
         */
        public Builder ReverseIp(@Nullable Output<String> ReverseIp) {
            $.ReverseIp = ReverseIp;
            return this;
        }

        /**
         * @param ReverseIp The IP to set the reverse of
         * 
         * @return builder
         * 
         */
        public Builder ReverseIp(String ReverseIp) {
            return ReverseIp(Output.of(ReverseIp));
        }

        /**
         * @param ReverseValue The value of the reverse
         * 
         * @return builder
         * 
         */
        public Builder ReverseValue(@Nullable Output<String> ReverseValue) {
            $.ReverseValue = ReverseValue;
            return this;
        }

        /**
         * @param ReverseValue The value of the reverse
         * 
         * @return builder
         * 
         */
        public Builder ReverseValue(String ReverseValue) {
            return ReverseValue(Output.of(ReverseValue));
        }

        /**
         * @param ip The IP block to which the IP belongs
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The IP block to which the IP belongs
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        public ReverseState build() {
            return $;
        }
    }

}
