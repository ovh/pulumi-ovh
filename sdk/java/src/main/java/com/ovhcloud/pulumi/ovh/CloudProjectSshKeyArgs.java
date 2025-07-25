// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class CloudProjectSshKeyArgs extends com.pulumi.resources.ResourceArgs {

    public static final CloudProjectSshKeyArgs Empty = new CloudProjectSshKeyArgs();

    /**
     * SSH key name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return SSH key name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * SSH public key
     * 
     */
    @Import(name="publicKey", required=true)
    private Output<String> publicKey;

    /**
     * @return SSH public key
     * 
     */
    public Output<String> publicKey() {
        return this.publicKey;
    }

    /**
     * Region to create SSH key
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Region to create SSH key
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Service name
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return Service name
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private CloudProjectSshKeyArgs() {}

    private CloudProjectSshKeyArgs(CloudProjectSshKeyArgs $) {
        this.name = $.name;
        this.publicKey = $.publicKey;
        this.region = $.region;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(CloudProjectSshKeyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private CloudProjectSshKeyArgs $;

        public Builder() {
            $ = new CloudProjectSshKeyArgs();
        }

        public Builder(CloudProjectSshKeyArgs defaults) {
            $ = new CloudProjectSshKeyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name SSH key name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name SSH key name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param publicKey SSH public key
         * 
         * @return builder
         * 
         */
        public Builder publicKey(Output<String> publicKey) {
            $.publicKey = publicKey;
            return this;
        }

        /**
         * @param publicKey SSH public key
         * 
         * @return builder
         * 
         */
        public Builder publicKey(String publicKey) {
            return publicKey(Output.of(publicKey));
        }

        /**
         * @param region Region to create SSH key
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Region to create SSH key
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public CloudProjectSshKeyArgs build() {
            if ($.publicKey == null) {
                throw new MissingRequiredPropertyException("CloudProjectSshKeyArgs", "publicKey");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("CloudProjectSshKeyArgs", "serviceName");
            }
            return $;
        }
    }

}
