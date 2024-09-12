// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UserRole {
    /**
     * @return A description associated with the user.
     * 
     */
    private @Nullable String description;
    /**
     * @return id of the role
     * 
     */
    private @Nullable String id;
    /**
     * @return name of the role
     * 
     */
    private @Nullable String name;
    /**
     * @return list of permissions associated with the role
     * 
     */
    private @Nullable List<String> permissions;

    private UserRole() {}
    /**
     * @return A description associated with the user.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return id of the role
     * 
     */
    public Optional<String> id() {
        return Optional.ofNullable(this.id);
    }
    /**
     * @return name of the role
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }
    /**
     * @return list of permissions associated with the role
     * 
     */
    public List<String> permissions() {
        return this.permissions == null ? List.of() : this.permissions;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UserRole defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String description;
        private @Nullable String id;
        private @Nullable String name;
        private @Nullable List<String> permissions;
        public Builder() {}
        public Builder(UserRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.permissions = defaults.permissions;
        }

        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder id(@Nullable String id) {

            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(@Nullable String name) {

            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder permissions(@Nullable List<String> permissions) {

            this.permissions = permissions;
            return this;
        }
        public Builder permissions(String... permissions) {
            return permissions(List.of(permissions));
        }
        public UserRole build() {
            final var _resultValue = new UserRole();
            _resultValue.description = description;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.permissions = permissions;
            return _resultValue;
        }
    }
}