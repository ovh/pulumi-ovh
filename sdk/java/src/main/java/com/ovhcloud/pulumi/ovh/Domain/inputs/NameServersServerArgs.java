// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Domain.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NameServersServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final NameServersServerArgs Empty = new NameServersServerArgs();

    /**
     * The server hostname
     * 
     */
    @Import(name="host", required=true)
    private Output<String> host;

    /**
     * @return The server hostname
     * 
     */
    public Output<String> host() {
        return this.host;
    }

    /**
     * The server IP
     * 
     */
    @Import(name="ip")
    private @Nullable Output<String> ip;

    /**
     * @return The server IP
     * 
     */
    public Optional<Output<String>> ip() {
        return Optional.ofNullable(this.ip);
    }

    private NameServersServerArgs() {}

    private NameServersServerArgs(NameServersServerArgs $) {
        this.host = $.host;
        this.ip = $.ip;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NameServersServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NameServersServerArgs $;

        public Builder() {
            $ = new NameServersServerArgs();
        }

        public Builder(NameServersServerArgs defaults) {
            $ = new NameServersServerArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param host The server hostname
         * 
         * @return builder
         * 
         */
        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The server hostname
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param ip The server IP
         * 
         * @return builder
         * 
         */
        public Builder ip(@Nullable Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The server IP
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        public NameServersServerArgs build() {
            if ($.host == null) {
                throw new MissingRequiredPropertyException("NameServersServerArgs", "host");
            }
            return $;
        }
    }

}
