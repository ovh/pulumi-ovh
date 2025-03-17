// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.StorageReplicationRuleDestination;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.StorageReplicationRuleFilter;
import com.pulumi.core.annotations.CustomType;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class StorageReplicationRule {
    /**
     * @return Delete marker replication
     * 
     */
    private @Nullable String deleteMarkerReplication;
    /**
     * @return Rule destination configuration
     * 
     */
    private @Nullable StorageReplicationRuleDestination destination;
    /**
     * @return Rule filters
     * 
     */
    private @Nullable StorageReplicationRuleFilter filter;
    /**
     * @return Rule ID
     * 
     */
    private @Nullable String id;
    /**
     * @return Rule priority
     * 
     */
    private @Nullable Double priority;
    /**
     * @return Rule status
     * 
     */
    private @Nullable String status;

    private StorageReplicationRule() {}
    /**
     * @return Delete marker replication
     * 
     */
    public Optional<String> deleteMarkerReplication() {
        return Optional.ofNullable(this.deleteMarkerReplication);
    }
    /**
     * @return Rule destination configuration
     * 
     */
    public Optional<StorageReplicationRuleDestination> destination() {
        return Optional.ofNullable(this.destination);
    }
    /**
     * @return Rule filters
     * 
     */
    public Optional<StorageReplicationRuleFilter> filter() {
        return Optional.ofNullable(this.filter);
    }
    /**
     * @return Rule ID
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return Rule priority
     * 
     */
    public Optional<Double> priority() {
        return Optional.ofNullable(this.priority);
    }
    /**
     * @return Rule status
     * 
     */
    public Optional<String> status() {
        return Optional.ofNullable(this.status);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(StorageReplicationRule defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String deleteMarkerReplication;
        private @Nullable StorageReplicationRuleDestination destination;
        private @Nullable StorageReplicationRuleFilter filter;
        private @Nullable String id;
        private @Nullable Double priority;
        private @Nullable String status;
        public Builder() {}
        public Builder(StorageReplicationRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.deleteMarkerReplication = defaults.deleteMarkerReplication;
    	      this.destination = defaults.destination;
    	      this.filter = defaults.filter;
    	      this.id = defaults.id;
    	      this.priority = defaults.priority;
    	      this.status = defaults.status;
        }

        @CustomType.Setter
        public Builder deleteMarkerReplication(@Nullable String deleteMarkerReplication) {

            this.deleteMarkerReplication = deleteMarkerReplication;
            return this;
        }
        @CustomType.Setter
        public Builder destination(@Nullable StorageReplicationRuleDestination destination) {

            this.destination = destination;
            return this;
        }
        @CustomType.Setter
        public Builder filter(@Nullable StorageReplicationRuleFilter filter) {

            this.filter = filter;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder priority(@Nullable Double priority) {

            this.priority = priority;
            return this;
        }
        @CustomType.Setter
        public Builder status(@Nullable String status) {

            this.status = status;
            return this;
        }
        public StorageReplicationRule build() {
            final var _resultValue = new StorageReplicationRule();
            _resultValue.deleteMarkerReplication = deleteMarkerReplication;
            _resultValue.destination = destination;
            _resultValue.filter = filter;
            _resultValue.id = id;
            _resultValue.priority = priority;
            _resultValue.status = status;
            return _resultValue;
        }
    }
}
