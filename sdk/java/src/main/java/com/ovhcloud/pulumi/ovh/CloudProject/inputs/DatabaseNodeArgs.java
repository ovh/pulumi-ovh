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


public final class DatabaseNodeArgs extends com.pulumi.resources.ResourceArgs {

    public static final DatabaseNodeArgs Empty = new DatabaseNodeArgs();

    /**
     * Private network id in which the node should be deployed. It&#39;s the regional openstackId of the private network
     * 
     */
    @Import(name="networkId")
    private @Nullable Output<String> networkId;

    /**
     * @return Private network id in which the node should be deployed. It&#39;s the regional openstackId of the private network
     * 
     */
    public Optional<Output<String>> networkId() {
        return Optional.ofNullable(this.networkId);
    }

    /**
     * Public cloud region in which the node should be deployed. Ex: &#34;GRA&#39;.
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return Public cloud region in which the node should be deployed. Ex: &#34;GRA&#39;.
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * Private subnet ID in which the node is.
     * 
     */
    @Import(name="subnetId")
    private @Nullable Output<String> subnetId;

    /**
     * @return Private subnet ID in which the node is.
     * 
     */
    public Optional<Output<String>> subnetId() {
        return Optional.ofNullable(this.subnetId);
    }

    private DatabaseNodeArgs() {}

    private DatabaseNodeArgs(DatabaseNodeArgs $) {
        this.networkId = $.networkId;
        this.region = $.region;
        this.subnetId = $.subnetId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(DatabaseNodeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private DatabaseNodeArgs $;

        public Builder() {
            $ = new DatabaseNodeArgs();
        }

        public Builder(DatabaseNodeArgs defaults) {
            $ = new DatabaseNodeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param networkId Private network id in which the node should be deployed. It&#39;s the regional openstackId of the private network
         * 
         * @return builder
         * 
         */
        public Builder networkId(@Nullable Output<String> networkId) {
            $.networkId = networkId;
            return this;
        }

        /**
         * @param networkId Private network id in which the node should be deployed. It&#39;s the regional openstackId of the private network
         * 
         * @return builder
         * 
         */
        public Builder networkId(String networkId) {
            return networkId(Output.of(networkId));
        }

        /**
         * @param region Public cloud region in which the node should be deployed. Ex: &#34;GRA&#39;.
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Public cloud region in which the node should be deployed. Ex: &#34;GRA&#39;.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param subnetId Private subnet ID in which the node is.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(@Nullable Output<String> subnetId) {
            $.subnetId = subnetId;
            return this;
        }

        /**
         * @param subnetId Private subnet ID in which the node is.
         * 
         * @return builder
         * 
         */
        public Builder subnetId(String subnetId) {
            return subnetId(Output.of(subnetId));
        }

        public DatabaseNodeArgs build() {
            if ($.region == null) {
                throw new MissingRequiredPropertyException("DatabaseNodeArgs", "region");
            }
            return $;
        }
    }

}
