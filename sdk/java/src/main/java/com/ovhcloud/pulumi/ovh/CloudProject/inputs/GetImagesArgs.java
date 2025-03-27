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


public final class GetImagesArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetImagesArgs Empty = new GetImagesArgs();

    /**
     * Get compatible images with flavor type
     * 
     */
    @Import(name="flavorType")
    private @Nullable Output<String> flavorType;

    /**
     * @return Get compatible images with flavor type
     * 
     */
    public Optional<Output<String>> flavorType() {
        return Optional.ofNullable(this.flavorType);
    }

    /**
     * Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     * 
     */
    @Import(name="osType")
    private @Nullable Output<String> osType;

    /**
     * @return Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
     * 
     */
    public Optional<Output<String>> osType() {
        return Optional.ofNullable(this.osType);
    }

    /**
     * Image region
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Image region
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Public cloud project ID
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return Public cloud project ID
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private GetImagesArgs() {}

    private GetImagesArgs(GetImagesArgs $) {
        this.flavorType = $.flavorType;
        this.osType = $.osType;
        this.region = $.region;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetImagesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetImagesArgs $;

        public Builder() {
            $ = new GetImagesArgs();
        }

        public Builder(GetImagesArgs defaults) {
            $ = new GetImagesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param flavorType Get compatible images with flavor type
         * 
         * @return builder
         * 
         */
        public Builder flavorType(@Nullable Output<String> flavorType) {
            $.flavorType = flavorType;
            return this;
        }

        /**
         * @param flavorType Get compatible images with flavor type
         * 
         * @return builder
         * 
         */
        public Builder flavorType(String flavorType) {
            return flavorType(Output.of(flavorType));
        }

        /**
         * @param osType Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
         * 
         * @return builder
         * 
         */
        public Builder osType(@Nullable Output<String> osType) {
            $.osType = osType;
            return this;
        }

        /**
         * @param osType Image OS (Allowed values: baremetal-linux ┃ bsd ┃ linux ┃ windows)
         * 
         * @return builder
         * 
         */
        public Builder osType(String osType) {
            return osType(Output.of(osType));
        }

        /**
         * @param region Image region
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Image region
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceName Public cloud project ID
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Public cloud project ID
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public GetImagesArgs build() {
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetImagesArgs", "serviceName");
            }
            return $;
        }
    }

}
