// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RancherCurrentStateIpRestrictionArgs extends com.pulumi.resources.ResourceArgs {

    public static final RancherCurrentStateIpRestrictionArgs Empty = new RancherCurrentStateIpRestrictionArgs();

    /**
     * Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)
     * 
     */
    @Import(name="cidrBlock")
    private @Nullable Output<String> cidrBlock;

    /**
     * @return Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)
     * 
     */
    public Optional<Output<String>> cidrBlock() {
        return Optional.ofNullable(this.cidrBlock);
    }

    /**
     * Description of the allowed CIDR block
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description of the allowed CIDR block
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    private RancherCurrentStateIpRestrictionArgs() {}

    private RancherCurrentStateIpRestrictionArgs(RancherCurrentStateIpRestrictionArgs $) {
        this.cidrBlock = $.cidrBlock;
        this.description = $.description;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RancherCurrentStateIpRestrictionArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RancherCurrentStateIpRestrictionArgs $;

        public Builder() {
            $ = new RancherCurrentStateIpRestrictionArgs();
        }

        public Builder(RancherCurrentStateIpRestrictionArgs defaults) {
            $ = new RancherCurrentStateIpRestrictionArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cidrBlock Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(@Nullable Output<String> cidrBlock) {
            $.cidrBlock = cidrBlock;
            return this;
        }

        /**
         * @param cidrBlock Allowed CIDR block (/subnet is optional, if unspecified then /32 will be used)
         * 
         * @return builder
         * 
         */
        public Builder cidrBlock(String cidrBlock) {
            return cidrBlock(Output.of(cidrBlock));
        }

        /**
         * @param description Description of the allowed CIDR block
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description of the allowed CIDR block
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public RancherCurrentStateIpRestrictionArgs build() {
            return $;
        }
    }

}
