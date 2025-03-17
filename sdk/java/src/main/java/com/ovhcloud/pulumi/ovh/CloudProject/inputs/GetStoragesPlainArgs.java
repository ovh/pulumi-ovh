// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetStoragesPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetStoragesPlainArgs Empty = new GetStoragesPlainArgs();

    /**
     * Region name
     * 
     */
    @Import(name="regionName", required=true)
    private String regionName;

    /**
     * @return Region name
     * 
     */
    public String regionName() {
        return this.regionName;
    }

    /**
     * Service name
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return Service name
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetStoragesPlainArgs() {}

    private GetStoragesPlainArgs(GetStoragesPlainArgs $) {
        this.regionName = $.regionName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetStoragesPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetStoragesPlainArgs $;

        public Builder() {
            $ = new GetStoragesPlainArgs();
        }

        public Builder(GetStoragesPlainArgs defaults) {
            $ = new GetStoragesPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param regionName Region name
         * 
         * @return builder
         * 
         */
        public Builder regionName(String regionName) {
            $.regionName = regionName;
            return this;
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetStoragesPlainArgs build() {
            if ($.regionName == null) {
                throw new MissingRequiredPropertyException("GetStoragesPlainArgs", "regionName");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetStoragesPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
