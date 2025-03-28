// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetStorageObjectPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetStorageObjectPlainArgs Empty = new GetStorageObjectPlainArgs();

    /**
     * Key
     * 
     */
    @Import(name="key", required=true)
    private String key;

    /**
     * @return Key
     * 
     */
    public String key() {
        return this.key;
    }

    /**
     * Name
     * 
     */
    @Import(name="name", required=true)
    private String name;

    /**
     * @return Name
     * 
     */
    public String name() {
        return this.name;
    }

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

    private GetStorageObjectPlainArgs() {}

    private GetStorageObjectPlainArgs(GetStorageObjectPlainArgs $) {
        this.key = $.key;
        this.name = $.name;
        this.regionName = $.regionName;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetStorageObjectPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetStorageObjectPlainArgs $;

        public Builder() {
            $ = new GetStorageObjectPlainArgs();
        }

        public Builder(GetStorageObjectPlainArgs defaults) {
            $ = new GetStorageObjectPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param key Key
         * 
         * @return builder
         * 
         */
        public Builder key(String key) {
            $.key = key;
            return this;
        }

        /**
         * @param name Name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            $.name = name;
            return this;
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

        public GetStorageObjectPlainArgs build() {
            if ($.key == null) {
                throw new MissingRequiredPropertyException("GetStorageObjectPlainArgs", "key");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetStorageObjectPlainArgs", "name");
            }
            if ($.regionName == null) {
                throw new MissingRequiredPropertyException("GetStorageObjectPlainArgs", "regionName");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetStorageObjectPlainArgs", "serviceName");
            }
            return $;
        }
    }

}
