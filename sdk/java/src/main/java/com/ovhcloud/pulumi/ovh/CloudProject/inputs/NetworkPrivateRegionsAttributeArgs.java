// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class NetworkPrivateRegionsAttributeArgs extends com.pulumi.resources.ResourceArgs {

    public static final NetworkPrivateRegionsAttributeArgs Empty = new NetworkPrivateRegionsAttributeArgs();

    @Import(name="openstackid")
    private @Nullable Output<String> openstackid;

    public Optional<Output<String>> openstackid() {
        return Optional.ofNullable(this.openstackid);
    }

    @Import(name="region")
    private @Nullable Output<String> region;

    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * the status of the network. should be normally set to &#39;ACTIVE&#39;.
     * 
     */
    @Import(name="status", required=true)
    private Output<String> status;

    /**
     * @return the status of the network. should be normally set to &#39;ACTIVE&#39;.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    private NetworkPrivateRegionsAttributeArgs() {}

    private NetworkPrivateRegionsAttributeArgs(NetworkPrivateRegionsAttributeArgs $) {
        this.openstackid = $.openstackid;
        this.region = $.region;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(NetworkPrivateRegionsAttributeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private NetworkPrivateRegionsAttributeArgs $;

        public Builder() {
            $ = new NetworkPrivateRegionsAttributeArgs();
        }

        public Builder(NetworkPrivateRegionsAttributeArgs defaults) {
            $ = new NetworkPrivateRegionsAttributeArgs(Objects.requireNonNull(defaults));
        }

        public Builder openstackid(@Nullable Output<String> openstackid) {
            $.openstackid = openstackid;
            return this;
        }

        public Builder openstackid(String openstackid) {
            return openstackid(Output.of(openstackid));
        }

        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param status the status of the network. should be normally set to &#39;ACTIVE&#39;.
         * 
         * @return builder
         * 
         */
        public Builder status(Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status the status of the network. should be normally set to &#39;ACTIVE&#39;.
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public NetworkPrivateRegionsAttributeArgs build() {
            if ($.status == null) {
                throw new MissingRequiredPropertyException("NetworkPrivateRegionsAttributeArgs", "status");
            }
            return $;
        }
    }

}
