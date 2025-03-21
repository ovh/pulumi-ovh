// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.StorageReplicationRuleDestinationArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.StorageReplicationRuleFilterArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class StorageReplicationRuleArgs extends com.pulumi.resources.ResourceArgs {

    public static final StorageReplicationRuleArgs Empty = new StorageReplicationRuleArgs();

    /**
     * Delete marker replication
     * 
     */
    @Import(name="deleteMarkerReplication")
    private @Nullable Output<String> deleteMarkerReplication;

    /**
     * @return Delete marker replication
     * 
     */
    public Optional<Output<String>> deleteMarkerReplication() {
        return Optional.ofNullable(this.deleteMarkerReplication);
    }

    /**
     * Rule destination configuration
     * 
     */
    @Import(name="destination")
    private @Nullable Output<StorageReplicationRuleDestinationArgs> destination;

    /**
     * @return Rule destination configuration
     * 
     */
    public Optional<Output<StorageReplicationRuleDestinationArgs>> destination() {
        return Optional.ofNullable(this.destination);
    }

    /**
     * Rule filters
     * 
     */
    @Import(name="filter")
    private @Nullable Output<StorageReplicationRuleFilterArgs> filter;

    /**
     * @return Rule filters
     * 
     */
    public Optional<Output<StorageReplicationRuleFilterArgs>> filter() {
        return Optional.ofNullable(this.filter);
    }

    /**
     * Rule ID
     * 
     */
    @Import(name="id")
    private @Nullable Output<String> id;

    /**
     * @return Rule ID
     * 
     */
    public Optional<Output<String>> id() {
        return Optional.ofNullable(this.id);
    }

    /**
     * Rule priority
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Double> priority;

    /**
     * @return Rule priority
     * 
     */
    public Optional<Output<Double>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * Rule status
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Rule status
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    private StorageReplicationRuleArgs() {}

    private StorageReplicationRuleArgs(StorageReplicationRuleArgs $) {
        this.deleteMarkerReplication = $.deleteMarkerReplication;
        this.destination = $.destination;
        this.filter = $.filter;
        this.id = $.id;
        this.priority = $.priority;
        this.status = $.status;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(StorageReplicationRuleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private StorageReplicationRuleArgs $;

        public Builder() {
            $ = new StorageReplicationRuleArgs();
        }

        public Builder(StorageReplicationRuleArgs defaults) {
            $ = new StorageReplicationRuleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param deleteMarkerReplication Delete marker replication
         * 
         * @return builder
         * 
         */
        public Builder deleteMarkerReplication(@Nullable Output<String> deleteMarkerReplication) {
            $.deleteMarkerReplication = deleteMarkerReplication;
            return this;
        }

        /**
         * @param deleteMarkerReplication Delete marker replication
         * 
         * @return builder
         * 
         */
        public Builder deleteMarkerReplication(String deleteMarkerReplication) {
            return deleteMarkerReplication(Output.of(deleteMarkerReplication));
        }

        /**
         * @param destination Rule destination configuration
         * 
         * @return builder
         * 
         */
        public Builder destination(@Nullable Output<StorageReplicationRuleDestinationArgs> destination) {
            $.destination = destination;
            return this;
        }

        /**
         * @param destination Rule destination configuration
         * 
         * @return builder
         * 
         */
        public Builder destination(StorageReplicationRuleDestinationArgs destination) {
            return destination(Output.of(destination));
        }

        /**
         * @param filter Rule filters
         * 
         * @return builder
         * 
         */
        public Builder filter(@Nullable Output<StorageReplicationRuleFilterArgs> filter) {
            $.filter = filter;
            return this;
        }

        /**
         * @param filter Rule filters
         * 
         * @return builder
         * 
         */
        public Builder filter(StorageReplicationRuleFilterArgs filter) {
            return filter(Output.of(filter));
        }

        /**
         * @param id Rule ID
         * 
         * @return builder
         * 
         */
        public Builder id(@Nullable Output<String> id) {
            $.id = id;
            return this;
        }

        /**
         * @param id Rule ID
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            return id(Output.of(id));
        }

        /**
         * @param priority Rule priority
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Double> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Rule priority
         * 
         * @return builder
         * 
         */
        public Builder priority(Double priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param status Rule status
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Rule status
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        public StorageReplicationRuleArgs build() {
            return $;
        }
    }

}
