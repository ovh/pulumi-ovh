// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetIdentityGroupsResult {
    /**
     * @return The list of the group names of all the identity groups.
     * 
     */
    private List<String> groups;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;

    private GetIdentityGroupsResult() {}
    /**
     * @return The list of the group names of all the identity groups.
     * 
     */
    public List<String> groups() {
        return this.groups;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityGroupsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<String> groups;
        private String id;
        public Builder() {}
        public Builder(GetIdentityGroupsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.id = defaults.id;
        }

        @CustomType.Setter
        public Builder groups(List<String> groups) {
            if (groups == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupsResult", "groups");
            }
            this.groups = groups;
            return this;
        }
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupsResult", "id");
            }
            this.id = id;
            return this;
        }
        public GetIdentityGroupsResult build() {
            final var _resultValue = new GetIdentityGroupsResult();
            _resultValue.groups = groups;
            _resultValue.id = id;
            return _resultValue;
        }
    }
}
