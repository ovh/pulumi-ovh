// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.StorageReplicationRuleFilterTag;
import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class StorageReplicationRuleFilter {
    /**
     * @return Prefix filter
     * 
     */
    private @Nullable String prefix;
    /**
     * @return Tags filter
     * 
     */
    private @Nullable List<StorageReplicationRuleFilterTag> tags;

    private StorageReplicationRuleFilter() {}
    /**
     * @return Prefix filter
     * 
     */
    public Optional<String> prefix() {
        return Optional.ofNullable(this.prefix);
    }
    /**
     * @return Tags filter
     * 
     */
    public List<StorageReplicationRuleFilterTag> tags() {
        return this.tags == null ? List.of() : this.tags;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(StorageReplicationRuleFilter defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String prefix;
        private @Nullable List<StorageReplicationRuleFilterTag> tags;
        public Builder() {}
        public Builder(StorageReplicationRuleFilter defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.prefix = defaults.prefix;
    	      this.tags = defaults.tags;
        }

        @CustomType.Setter
        public Builder prefix(@Nullable String prefix) {

            this.prefix = prefix;
            return this;
        }
        @CustomType.Setter
        public Builder tags(@Nullable List<StorageReplicationRuleFilterTag> tags) {

            this.tags = tags;
            return this;
        }
        public Builder tags(StorageReplicationRuleFilterTag... tags) {
            return tags(List.of(tags));
        }
        public StorageReplicationRuleFilter build() {
            final var _resultValue = new StorageReplicationRuleFilter();
            _resultValue.prefix = prefix;
            _resultValue.tags = tags;
            return _resultValue;
        }
    }
}
