// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Ip.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetFirewallPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFirewallPlainArgs Empty = new GetFirewallPlainArgs();

    /**
     * The IP or the CIDR
     * 
     */
    @Import(name="ip", required=true)
    private String ip;

    /**
     * @return The IP or the CIDR
     * 
     */
    public String ip() {
        return this.ip;
    }

    /**
     * IPv4 address
     * 
     */
    @Import(name="ipOnFirewall", required=true)
    private String ipOnFirewall;

    /**
     * @return IPv4 address
     * 
     */
    public String ipOnFirewall() {
        return this.ipOnFirewall;
    }

    private GetFirewallPlainArgs() {}

    private GetFirewallPlainArgs(GetFirewallPlainArgs $) {
        this.ip = $.ip;
        this.ipOnFirewall = $.ipOnFirewall;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFirewallPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFirewallPlainArgs $;

        public Builder() {
            $ = new GetFirewallPlainArgs();
        }

        public Builder(GetFirewallPlainArgs defaults) {
            $ = new GetFirewallPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param ip The IP or the CIDR
         * 
         * @return builder
         * 
         */
        public Builder ip(String ip) {
            $.ip = ip;
            return this;
        }

        /**
         * @param ipOnFirewall IPv4 address
         * 
         * @return builder
         * 
         */
        public Builder ipOnFirewall(String ipOnFirewall) {
            $.ipOnFirewall = ipOnFirewall;
            return this;
        }

        public GetFirewallPlainArgs build() {
            if ($.ip == null) {
                throw new MissingRequiredPropertyException("GetFirewallPlainArgs", "ip");
            }
            if ($.ipOnFirewall == null) {
                throw new MissingRequiredPropertyException("GetFirewallPlainArgs", "ipOnFirewall");
            }
            return $;
        }
    }

}