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


public final class StorageReplicationRuleDestinationArgs extends com.pulumi.resources.ResourceArgs {

    public static final StorageReplicationRuleDestinationArgs Empty = new StorageReplicationRuleDestinationArgs();

    /**
     * Destination bucket name
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Destination bucket name
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * Destination region
     * 
     */
    @Import(name="region", required=true)
    private Output<String> region;

    /**
     * @return Destination region
     * 
     */
    public Output<String> region() {
        return this.region;
    }

    /**
     * Destination storage class
     * 
     */
    @Import(name="storageClass")
    private @Nullable Output<String> storageClass;

    /**
     * @return Destination storage class
     * 
     */
    public Optional<Output<String>> storageClass() {
        return Optional.ofNullable(this.storageClass);
    }

    private StorageReplicationRuleDestinationArgs() {}

    private StorageReplicationRuleDestinationArgs(StorageReplicationRuleDestinationArgs $) {
        this.name = $.name;
        this.region = $.region;
        this.storageClass = $.storageClass;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StorageReplicationRuleDestinationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StorageReplicationRuleDestinationArgs $;

        public Builder() {
            $ = new StorageReplicationRuleDestinationArgs();
        }

        public Builder(StorageReplicationRuleDestinationArgs defaults) {
            $ = new StorageReplicationRuleDestinationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Destination bucket name
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Destination bucket name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param region Destination region
         * 
         * @return builder
         * 
         */
        public Builder region(Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Destination region
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param storageClass Destination storage class
         * 
         * @return builder
         * 
         */
        public Builder storageClass(@Nullable Output<String> storageClass) {
            $.storageClass = storageClass;
            return this;
        }

        /**
         * @param storageClass Destination storage class
         * 
         * @return builder
         * 
         */
        public Builder storageClass(String storageClass) {
            return storageClass(Output.of(storageClass));
        }

        public StorageReplicationRuleDestinationArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("StorageReplicationRuleDestinationArgs", "name");
            }
            if ($.region == null) {
                throw new MissingRequiredPropertyException("StorageReplicationRuleDestinationArgs", "region");
            }
            return $;
        }
    }

}
