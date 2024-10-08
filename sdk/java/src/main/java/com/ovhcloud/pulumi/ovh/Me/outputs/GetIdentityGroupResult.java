// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetIdentityGroupResult {
    /**
     * @return Identity URN of the group.
     * 
     */
    private String GroupURN;
    /**
     * @return Creation date of this group.
     * 
     */
    private String creation;
    /**
     * @return Is the group a default and immutable one.
     * 
     */
    private Boolean defaultGroup;
    /**
     * @return Group description.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Date of the last update of this group.
     * 
     */
    private String lastUpdate;
    private String name;
    /**
     * @return Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
     * 
     */
    private String role;

    private GetIdentityGroupResult() {}
    /**
     * @return Identity URN of the group.
     * 
     */
    public String GroupURN() {
        return this.GroupURN;
    }
    /**
     * @return Creation date of this group.
     * 
     */
    public String creation() {
        return this.creation;
    }
    /**
     * @return Is the group a default and immutable one.
     * 
     */
    public Boolean defaultGroup() {
        return this.defaultGroup;
    }
    /**
     * @return Group description.
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Date of the last update of this group.
     * 
     */
    public String lastUpdate() {
        return this.lastUpdate;
    }
    public String name() {
        return this.name;
    }
    /**
     * @return Role associated with the group. Valid roles are ADMIN, REGULAR, UNPRIVILEGED, and NONE.
     * 
     */
    public String role() {
        return this.role;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetIdentityGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String GroupURN;
        private String creation;
        private Boolean defaultGroup;
        private String description;
        private String id;
        private String lastUpdate;
        private String name;
        private String role;
        public Builder() {}
        public Builder(GetIdentityGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.GroupURN = defaults.GroupURN;
    	      this.creation = defaults.creation;
    	      this.defaultGroup = defaults.defaultGroup;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.lastUpdate = defaults.lastUpdate;
    	      this.name = defaults.name;
    	      this.role = defaults.role;
        }

        @CustomType.Setter
        public Builder GroupURN(String GroupURN) {
            if (GroupURN == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "GroupURN");
            }
            this.GroupURN = GroupURN;
            return this;
        }
        @CustomType.Setter
        public Builder creation(String creation) {
            if (creation == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "creation");
            }
            this.creation = creation;
            return this;
        }
        @CustomType.Setter
        public Builder defaultGroup(Boolean defaultGroup) {
            if (defaultGroup == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "defaultGroup");
            }
            this.defaultGroup = defaultGroup;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lastUpdate(String lastUpdate) {
            if (lastUpdate == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "lastUpdate");
            }
            this.lastUpdate = lastUpdate;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            if (role == null) {
              throw new MissingRequiredPropertyException("GetIdentityGroupResult", "role");
            }
            this.role = role;
            return this;
        }
        public GetIdentityGroupResult build() {
            final var _resultValue = new GetIdentityGroupResult();
            _resultValue.GroupURN = GroupURN;
            _resultValue.creation = creation;
            _resultValue.defaultGroup = defaultGroup;
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.lastUpdate = lastUpdate;
            _resultValue.name = name;
            _resultValue.role = role;
            return _resultValue;
        }
    }
}
