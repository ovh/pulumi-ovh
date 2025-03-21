// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetNasHAPartitionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetNasHAPartitionPlainArgs Empty = new GetNasHAPartitionPlainArgs();

    /**
     * The name of your dedicated HA-NAS partition.
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return The name of your dedicated HA-NAS partition.
     * 
     */
    public String name() {
        return this.name;
    }

    /**
     * The service_name of your dedicated HA-NAS.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The service_name of your dedicated HA-NAS.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    private GetNasHAPartitionPlainArgs() {}

    private GetNasHAPartitionPlainArgs(GetNasHAPartitionPlainArgs $) {
        this.name = $.name;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetNasHAPartitionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetNasHAPartitionPlainArgs $;

        public Builder() {
            $ = new GetNasHAPartitionPlainArgs();
        }

        public Builder(GetNasHAPartitionPlainArgs defaults) {
            $ = new GetNasHAPartitionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The name of your dedicated HA-NAS partition.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
        }

        /**
         * @param serviceName The service_name of your dedicated HA-NAS.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        public GetNasHAPartitionPlainArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetNasHAPartitionPlainArgs", "name");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetNasHAPartitionPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
