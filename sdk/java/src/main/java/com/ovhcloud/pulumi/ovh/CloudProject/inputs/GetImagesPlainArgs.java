// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetImagesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetImagesPlainArgs Empty = new GetImagesPlainArgs();

    /**
     * Get compatible images with flavor type
     * 
     */
    @Import(name="flavorType")
    private @Nullable String flavorType;

    /**
     * @return Get compatible images with flavor type
     * 
     */
    public Optional<String> flavorType() {
        return Optional.ofNullable(this.flavorType);
    }

    /**
     * Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     * 
     */
    @Import(name="osType")
    private @Nullable String osType;

    /**
     * @return Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     * 
     */
    public Optional<String> osType() {
        return Optional.ofNullable(this.osType);
    }

    /**
     * Image region
     * 
     */
    @Import(name="region")
    private @Nullable String region;

    /**
     * @return Image region
     * 
     */
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Public cloud project ID
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return Public cloud project ID
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetImagesPlainArgs() {}

    private GetImagesPlainArgs(GetImagesPlainArgs $) {
        this.flavorType = $.flavorType;
        this.osType = $.osType;
        this.region = $.region;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetImagesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetImagesPlainArgs $;

        public Builder() {
            $ = new GetImagesPlainArgs();
        }

        public Builder(GetImagesPlainArgs defaults) {
            $ = new GetImagesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param flavorType Get compatible images with flavor type
         * 
         * @return builder
         * 
         */
        public Builder flavorType(@Nullable String flavorType) {
            $.flavorType = flavorType;
            return this;
        }

        /**
         * @param osType Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
         * 
         * @return builder
         * 
         */
        public Builder osType(@Nullable String osType) {
            $.osType = osType;
            return this;
        }

        /**
         * @param region Image region
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable String region) {
            $.region = region;
            return this;
        }

        /**
         * @param serviceName Public cloud project ID
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetImagesPlainArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetImagesPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
