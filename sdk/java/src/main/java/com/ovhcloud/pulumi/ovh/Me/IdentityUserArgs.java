// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class IdentityUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final IdentityUserArgs Empty = new IdentityUserArgs();

    /**
     * User description.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return User description.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * User&#39;s email.
     * 
     */
    @Import(name="email", required=true)
    private Output<String> email;

    /**
     * @return User&#39;s email.
     * 
     */
    public Output<String> email() {
        return this.email;
    }

    /**
     * User&#39;s group.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return User&#39;s group.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * User&#39;s login suffix.
     * 
     */
    @Import(name="login", required=true)
    private Output<String> login;

    /**
     * @return User&#39;s login suffix.
     * 
     */
    public Output<String> login() {
        return this.login;
    }

    /**
     * User&#39;s password.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return User&#39;s password.
     * 
     */
    public Output<String> password() {
        return this.password;
    }

    private IdentityUserArgs() {}

    private IdentityUserArgs(IdentityUserArgs $) {
        this.description = $.description;
        this.email = $.email;
        this.group = $.group;
        this.login = $.login;
        this.password = $.password;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(IdentityUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private IdentityUserArgs $;

        public Builder() {
            $ = new IdentityUserArgs();
        }

        public Builder(IdentityUserArgs defaults) {
            $ = new IdentityUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param description User description.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description User description.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param email User&#39;s email.
         * 
         * @return builder
         * 
         */
        public Builder email(Output<String> email) {
            $.email = email;
            return this;
        }

        /**
         * @param email User&#39;s email.
         * 
         * @return builder
         * 
         */
        public Builder email(String email) {
            return email(Output.of(email));
        }

        /**
         * @param group User&#39;s group.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group User&#39;s group.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param login User&#39;s login suffix.
         * 
         * @return builder
         * 
         */
        public Builder login(Output<String> login) {
            $.login = login;
            return this;
        }

        /**
         * @param login User&#39;s login suffix.
         * 
         * @return builder
         * 
         */
        public Builder login(String login) {
            return login(Output.of(login));
        }

        /**
         * @param password User&#39;s password.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password User&#39;s password.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        public IdentityUserArgs build() {
            if ($.email == null) {
                throw new MissingRequiredPropertyException("IdentityUserArgs", "email");
            }
            if ($.login == null) {
                throw new MissingRequiredPropertyException("IdentityUserArgs", "login");
            }
            if ($.password == null) {
                throw new MissingRequiredPropertyException("IdentityUserArgs", "password");
            }
            return $;
        }
    }

}
