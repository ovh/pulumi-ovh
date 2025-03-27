// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class M3DbUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final M3DbUserArgs Empty = new M3DbUserArgs();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId", required=true)
    private Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Output<String> clusterId() {
        return this.clusterId;
    }

    /**
     * Group of the user.
     * 
     */
    @Import(name="group")
    private @Nullable Output<String> group;

    /**
     * @return Group of the user.
     * 
     */
    public Optional<Output<String>> group() {
        return Optional.ofNullable(this.group);
    }

    /**
     * Name of the user. A user named &#34;avnadmin&#34; is mapped with already created admin user instead of creating a new user.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the user. A user named &#34;avnadmin&#34; is mapped with already created admin user instead of creating a new user.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Arbitrary string to change to trigger a password update
     * 
     */
    @Import(name="passwordReset")
    private @Nullable Output<String> passwordReset;

    /**
     * @return Arbitrary string to change to trigger a password update
     * 
     */
    public Optional<Output<String>> passwordReset() {
        return Optional.ofNullable(this.passwordReset);
    }

    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Output<String> serviceName() {
        return this.serviceName;
    }

    private M3DbUserArgs() {}

    private M3DbUserArgs(M3DbUserArgs $) {
        this.clusterId = $.clusterId;
        this.group = $.group;
        this.name = $.name;
        this.passwordReset = $.passwordReset;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(M3DbUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private M3DbUserArgs $;

        public Builder() {
            $ = new M3DbUserArgs();
        }

        public Builder(M3DbUserArgs defaults) {
            $ = new M3DbUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(Output<String> clusterId) {
            $.clusterId = clusterId;
            return this;
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(String clusterId) {
            return clusterId(Output.of(clusterId));
        }

        /**
         * @param group Group of the user.
         * 
         * @return builder
         * 
         */
        public Builder group(@Nullable Output<String> group) {
            $.group = group;
            return this;
        }

        /**
         * @param group Group of the user.
         * 
         * @return builder
         * 
         */
        public Builder group(String group) {
            return group(Output.of(group));
        }

        /**
         * @param name Name of the user. A user named &#34;avnadmin&#34; is mapped with already created admin user instead of creating a new user.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the user. A user named &#34;avnadmin&#34; is mapped with already created admin user instead of creating a new user.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param passwordReset Arbitrary string to change to trigger a password update
         * 
         * @return builder
         * 
         */
        public Builder passwordReset(@Nullable Output<String> passwordReset) {
            $.passwordReset = passwordReset;
            return this;
        }

        /**
         * @param passwordReset Arbitrary string to change to trigger a password update
         * 
         * @return builder
         * 
         */
        public Builder passwordReset(String passwordReset) {
            return passwordReset(Output.of(passwordReset));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        public M3DbUserArgs build() {
            if ($.clusterId == null) {
                throw new MissingRequiredPropertyException("M3DbUserArgs", "clusterId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("M3DbUserArgs", "serviceName");
            }
            return $;
        }
    }

}
