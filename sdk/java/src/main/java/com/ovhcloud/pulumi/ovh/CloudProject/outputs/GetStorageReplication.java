// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetStorageReplicationRule;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetStorageReplication {
    /**
     * @return Replication rules
     * 
     */
    private List<GetStorageReplicationRule> rules;

    private GetStorageReplication() {}
    /**
     * @return Replication rules
     * 
     */
    public List<GetStorageReplicationRule> rules() {
        return this.rules;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStorageReplication defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetStorageReplicationRule> rules;
        public Builder() {}
        public Builder(GetStorageReplication defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.rules = defaults.rules;
        }

        @CustomType.Setter
        public Builder rules(List<GetStorageReplicationRule> rules) {
            if (rules == null) {
              throw new MissingRequiredPropertyException("GetStorageReplication", "rules");
            }
            this.rules = rules;
            return this;
        }
        public Builder rules(GetStorageReplicationRule... rules) {
            return rules(List.of(rules));
        }
        public GetStorageReplication build() {
            final var _resultValue = new GetStorageReplication();
            _resultValue.rules = rules;
            return _resultValue;
        }
    }
}
