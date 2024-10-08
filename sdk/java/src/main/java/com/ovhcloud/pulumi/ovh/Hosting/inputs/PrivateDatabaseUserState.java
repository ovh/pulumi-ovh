// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Hosting.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class PrivateDatabaseUserState extends com.pulumi.resources.ResourceArgs {

    public static final PrivateDatabaseUserState Empty = new PrivateDatabaseUserState();

    /**
     * Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The internal name of your private database.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The internal name of your private database.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * User name used to connect on your databases
     * 
     */
    @Import(name="userName")
    private @Nullable Output<String> userName;

    /**
     * @return User name used to connect on your databases
     * 
     */
    public Optional<Output<String>> userName() {
        return Optional.ofNullable(this.userName);
    }

    private PrivateDatabaseUserState() {}

    private PrivateDatabaseUserState(PrivateDatabaseUserState $) {
        this.password = $.password;
        this.serviceName = $.serviceName;
        this.userName = $.userName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(PrivateDatabaseUserState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private PrivateDatabaseUserState $;

        public Builder() {
            $ = new PrivateDatabaseUserState();
        }

        public Builder(PrivateDatabaseUserState defaults) {
            $ = new PrivateDatabaseUserState(Objects.requireNonNull(defaults));
        }

        /**
         * @param password Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password Password for the new user (alphanumeric, minimum one number and 8 characters minimum)
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param serviceName The internal name of your private database.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The internal name of your private database.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param userName User name used to connect on your databases
         * 
         * @return builder
         * 
         */
        public Builder userName(@Nullable Output<String> userName) {
            $.userName = userName;
            return this;
        }

        /**
         * @param userName User name used to connect on your databases
         * 
         * @return builder
         * 
         */
        public Builder userName(String userName) {
            return userName(Output.of(userName));
        }

        public PrivateDatabaseUserState build() {
            return $;
        }
    }

}
