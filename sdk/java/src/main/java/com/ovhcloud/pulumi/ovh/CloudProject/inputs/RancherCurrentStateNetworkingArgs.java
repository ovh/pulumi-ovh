// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RancherCurrentStateNetworkingArgs extends com.pulumi.resources.ResourceArgs {

    public static final RancherCurrentStateNetworkingArgs Empty = new RancherCurrentStateNetworkingArgs();

    /**
     * Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.
     * 
     */
    @Import(name="egressCidrBlocks")
    private @Nullable Output<List<String>> egressCidrBlocks;

    /**
     * @return Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.
     * 
     */
    public Optional<Output<List<String>>> egressCidrBlocks() {
        return Optional.ofNullable(this.egressCidrBlocks);
    }

    private RancherCurrentStateNetworkingArgs() {}

    private RancherCurrentStateNetworkingArgs(RancherCurrentStateNetworkingArgs $) {
        this.egressCidrBlocks = $.egressCidrBlocks;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RancherCurrentStateNetworkingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RancherCurrentStateNetworkingArgs $;

        public Builder() {
            $ = new RancherCurrentStateNetworkingArgs();
        }

        public Builder(RancherCurrentStateNetworkingArgs defaults) {
            $ = new RancherCurrentStateNetworkingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param egressCidrBlocks Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.
         * 
         * @return builder
         * 
         */
        public Builder egressCidrBlocks(@Nullable Output<List<String>> egressCidrBlocks) {
            $.egressCidrBlocks = egressCidrBlocks;
            return this;
        }

        /**
         * @param egressCidrBlocks Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.
         * 
         * @return builder
         * 
         */
        public Builder egressCidrBlocks(List<String> egressCidrBlocks) {
            return egressCidrBlocks(Output.of(egressCidrBlocks));
        }

        /**
         * @param egressCidrBlocks Specifies the CIDR ranges for egress IP addresses used by Rancher. Ensure these ranges are allowed in any IP restrictions for services that Rancher will access.
         * 
         * @return builder
         * 
         */
        public Builder egressCidrBlocks(String... egressCidrBlocks) {
            return egressCidrBlocks(List.of(egressCidrBlocks));
        }

        public RancherCurrentStateNetworkingArgs build() {
            return $;
        }
    }

}
