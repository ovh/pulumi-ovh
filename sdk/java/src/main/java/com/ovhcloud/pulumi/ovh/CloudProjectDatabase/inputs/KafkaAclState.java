// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class KafkaAclState extends com.pulumi.resources.ResourceArgs {

    public static final KafkaAclState Empty = new KafkaAclState();

    /**
     * Cluster ID.
     * 
     */
    @Import(name="clusterId")
    private @Nullable Output<String> clusterId;

    /**
     * @return Cluster ID.
     * 
     */
    public Optional<Output<String>> clusterId() {
        return Optional.ofNullable(this.clusterId);
    }

    /**
     * Permission to give to this username on this topic.
     * Available permissions:
     * 
     */
    @Import(name="permission")
    private @Nullable Output<String> permission;

    /**
     * @return Permission to give to this username on this topic.
     * Available permissions:
     * 
     */
    public Optional<Output<String>> permission() {
        return Optional.ofNullable(this.permission);
    }

    /**
     * The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return The id of the public cloud project. If omitted,
     * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Topic affected by this ACL.
     * 
     */
    @Import(name="topic")
    private @Nullable Output<String> topic;

    /**
     * @return Topic affected by this ACL.
     * 
     */
    public Optional<Output<String>> topic() {
        return Optional.ofNullable(this.topic);
    }

    /**
     * Username affected by this ACL.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return Username affected by this ACL.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private KafkaAclState() {}

    private KafkaAclState(KafkaAclState $) {
        this.clusterId = $.clusterId;
        this.permission = $.permission;
        this.serviceName = $.serviceName;
        this.topic = $.topic;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(KafkaAclState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private KafkaAclState $;

        public Builder() {
            $ = new KafkaAclState();
        }

        public Builder(KafkaAclState defaults) {
            $ = new KafkaAclState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clusterId Cluster ID.
         * 
         * @return builder
         * 
         */
        public Builder clusterId(@Nullable Output<String> clusterId) {
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
         * @param permission Permission to give to this username on this topic.
         * Available permissions:
         * 
         * @return builder
         * 
         */
        public Builder permission(@Nullable Output<String> permission) {
            $.permission = permission;
            return this;
        }

        /**
         * @param permission Permission to give to this username on this topic.
         * Available permissions:
         * 
         * @return builder
         * 
         */
        public Builder permission(String permission) {
            return permission(Output.of(permission));
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName The id of the public cloud project. If omitted,
         * the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param topic Topic affected by this ACL.
         * 
         * @return builder
         * 
         */
        public Builder topic(@Nullable Output<String> topic) {
            $.topic = topic;
            return this;
        }

        /**
         * @param topic Topic affected by this ACL.
         * 
         * @return builder
         * 
         */
        public Builder topic(String topic) {
            return topic(Output.of(topic));
        }

        /**
         * @param username Username affected by this ACL.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Username affected by this ACL.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public KafkaAclState build() {
            return $;
        }
    }

}