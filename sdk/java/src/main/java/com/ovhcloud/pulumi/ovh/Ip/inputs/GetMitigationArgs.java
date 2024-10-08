// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Ip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetMitigationArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMitigationArgs Empty = new GetMitigationArgs();

    /**
     * The IP or the CIDR
     * 
     */
    @Import(name="ip", required=true)
    private Output<String> ip;

    /**
     * @return The IP or the CIDR
     * 
     */
    public Output<String> ip() {
        return this.ip;
    }

    /**
     * IPv4 address
     * 
     */
    @Import(name="ipOnMitigation", required=true)
    private Output<String> ipOnMitigation;

    /**
     * @return IPv4 address
     * 
     */
    public Output<String> ipOnMitigation() {
        return this.ipOnMitigation;
    }

    private GetMitigationArgs() {}

    private GetMitigationArgs(GetMitigationArgs $) {
        this.ip = $.ip;
        this.ipOnMitigation = $.ipOnMitigation;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMitigationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMitigationArgs $;

        public Builder() {
            $ = new GetMitigationArgs();
        }

        public Builder(GetMitigationArgs defaults) {
            $ = new GetMitigationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip The IP or the CIDR
         * 
         * @return builder
         * 
         */
        public Builder ip(Output<String> ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ip The IP or the CIDR
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            return ip(Output.of(ip));
        }

        /**
         * @param ipOnMitigation IPv4 address
         * 
         * @return builder
         * 
         */
        public Builder ipOnMitigation(Output<String> ipOnMitigation) {
            $.ipOnMitigation = ipOnMitigation;
            return this;
        }

        /**
         * @param ipOnMitigation IPv4 address
         * 
         * @return builder
         * 
         */
        public Builder ipOnMitigation(String ipOnMitigation) {
            return ipOnMitigation(Output.of(ipOnMitigation));
        }

        public GetMitigationArgs build() {
            if ($.ip == null) {
                throw new MissingRequiredPropertyException("GetMitigationArgs", "ip");
            }
            if ($.ipOnMitigation == null) {
                throw new MissingRequiredPropertyException("GetMitigationArgs", "ipOnMitigation");
            }
            return $;
        }
    }

}
