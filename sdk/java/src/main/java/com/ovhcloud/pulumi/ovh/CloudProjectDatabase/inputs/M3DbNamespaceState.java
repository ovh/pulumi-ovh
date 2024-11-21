// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class M3DbNamespaceState extends com.pulumi.resources.ResourceArgs {

    public static final M3DbNamespaceState Empty = new M3DbNamespaceState();

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
     * Name of the namespace. A namespace named &#34;default&#34; is mapped with already created default namespace instead of creating a new namespace.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the namespace. A namespace named &#34;default&#34; is mapped with already created default namespace instead of creating a new namespace.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Import(name="resolution")
    private @Nullable Output<String> resolution;

    /**
     * @return Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Optional<Output<String>> resolution() {
        return Optional.ofNullable(this.resolution);
    }

    /**
     * Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Import(name="retentionBlockDataExpirationDuration")
    private @Nullable Output<String> retentionBlockDataExpirationDuration;

    /**
     * @return Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Optional<Output<String>> retentionBlockDataExpirationDuration() {
        return Optional.ofNullable(this.retentionBlockDataExpirationDuration);
    }

    /**
     * Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Import(name="retentionBlockSizeDuration")
    private @Nullable Output<String> retentionBlockSizeDuration;

    /**
     * @return Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Optional<Output<String>> retentionBlockSizeDuration() {
        return Optional.ofNullable(this.retentionBlockSizeDuration);
    }

    /**
     * Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Import(name="retentionBufferFutureDuration")
    private @Nullable Output<String> retentionBufferFutureDuration;

    /**
     * @return Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Optional<Output<String>> retentionBufferFutureDuration() {
        return Optional.ofNullable(this.retentionBufferFutureDuration);
    }

    /**
     * Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Import(name="retentionBufferPastDuration")
    private @Nullable Output<String> retentionBufferPastDuration;

    /**
     * @return Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Optional<Output<String>> retentionBufferPastDuration() {
        return Optional.ofNullable(this.retentionBufferPastDuration);
    }

    /**
     * Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    @Import(name="retentionPeriodDuration")
    private @Nullable Output<String> retentionPeriodDuration;

    /**
     * @return Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
     * 
     */
    public Optional<Output<String>> retentionPeriodDuration() {
        return Optional.ofNullable(this.retentionPeriodDuration);
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
     * Defines whether M3DB will create snapshot files for this namespace.
     * 
     */
    @Import(name="snapshotEnabled")
    private @Nullable Output<Boolean> snapshotEnabled;

    /**
     * @return Defines whether M3DB will create snapshot files for this namespace.
     * 
     */
    public Optional<Output<Boolean>> snapshotEnabled() {
        return Optional.ofNullable(this.snapshotEnabled);
    }

    /**
     * Type of namespace.
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of namespace.
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * Defines whether M3DB will include writes to this namespace in the commit log.
     * 
     */
    @Import(name="writesToCommitLogEnabled")
    private @Nullable Output<Boolean> writesToCommitLogEnabled;

    /**
     * @return Defines whether M3DB will include writes to this namespace in the commit log.
     * 
     */
    public Optional<Output<Boolean>> writesToCommitLogEnabled() {
        return Optional.ofNullable(this.writesToCommitLogEnabled);
    }

    private M3DbNamespaceState() {}

    private M3DbNamespaceState(M3DbNamespaceState $) {
        this.clusterId = $.clusterId;
        this.name = $.name;
        this.resolution = $.resolution;
        this.retentionBlockDataExpirationDuration = $.retentionBlockDataExpirationDuration;
        this.retentionBlockSizeDuration = $.retentionBlockSizeDuration;
        this.retentionBufferFutureDuration = $.retentionBufferFutureDuration;
        this.retentionBufferPastDuration = $.retentionBufferPastDuration;
        this.retentionPeriodDuration = $.retentionPeriodDuration;
        this.serviceName = $.serviceName;
        this.snapshotEnabled = $.snapshotEnabled;
        this.type = $.type;
        this.writesToCommitLogEnabled = $.writesToCommitLogEnabled;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(M3DbNamespaceState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private M3DbNamespaceState $;

        public Builder() {
            $ = new M3DbNamespaceState();
        }

        public Builder(M3DbNamespaceState defaults) {
            $ = new M3DbNamespaceState(Objects.requireNonNull(defaults));
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
         * @param name Name of the namespace. A namespace named &#34;default&#34; is mapped with already created default namespace instead of creating a new namespace.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the namespace. A namespace named &#34;default&#34; is mapped with already created default namespace instead of creating a new namespace.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param resolution Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder resolution(@Nullable Output<String> resolution) {
            $.resolution = resolution;
            return this;
        }

        /**
         * @param resolution Resolution for an aggregated namespace. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder resolution(String resolution) {
            return resolution(Output.of(resolution));
        }

        /**
         * @param retentionBlockDataExpirationDuration Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBlockDataExpirationDuration(@Nullable Output<String> retentionBlockDataExpirationDuration) {
            $.retentionBlockDataExpirationDuration = retentionBlockDataExpirationDuration;
            return this;
        }

        /**
         * @param retentionBlockDataExpirationDuration Controls how long we wait before expiring stale data. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBlockDataExpirationDuration(String retentionBlockDataExpirationDuration) {
            return retentionBlockDataExpirationDuration(Output.of(retentionBlockDataExpirationDuration));
        }

        /**
         * @param retentionBlockSizeDuration Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBlockSizeDuration(@Nullable Output<String> retentionBlockSizeDuration) {
            $.retentionBlockSizeDuration = retentionBlockSizeDuration;
            return this;
        }

        /**
         * @param retentionBlockSizeDuration Controls how long to keep a block in memory before flushing to a fileset on disk. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBlockSizeDuration(String retentionBlockSizeDuration) {
            return retentionBlockSizeDuration(Output.of(retentionBlockSizeDuration));
        }

        /**
         * @param retentionBufferFutureDuration Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBufferFutureDuration(@Nullable Output<String> retentionBufferFutureDuration) {
            $.retentionBufferFutureDuration = retentionBufferFutureDuration;
            return this;
        }

        /**
         * @param retentionBufferFutureDuration Controls how far into the future writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBufferFutureDuration(String retentionBufferFutureDuration) {
            return retentionBufferFutureDuration(Output.of(retentionBufferFutureDuration));
        }

        /**
         * @param retentionBufferPastDuration Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBufferPastDuration(@Nullable Output<String> retentionBufferPastDuration) {
            $.retentionBufferPastDuration = retentionBufferPastDuration;
            return this;
        }

        /**
         * @param retentionBufferPastDuration Controls how far into the past writes to the namespace will be accepted. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionBufferPastDuration(String retentionBufferPastDuration) {
            return retentionBufferPastDuration(Output.of(retentionBufferPastDuration));
        }

        /**
         * @param retentionPeriodDuration Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriodDuration(@Nullable Output<String> retentionPeriodDuration) {
            $.retentionPeriodDuration = retentionPeriodDuration;
            return this;
        }

        /**
         * @param retentionPeriodDuration Controls the duration of time that M3DB will retain data for the namespace. Should follow Rfc3339 e.g P2D, PT48H.
         * 
         * @return builder
         * 
         */
        public Builder retentionPeriodDuration(String retentionPeriodDuration) {
            return retentionPeriodDuration(Output.of(retentionPeriodDuration));
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
         * @param snapshotEnabled Defines whether M3DB will create snapshot files for this namespace.
         * 
         * @return builder
         * 
         */
        public Builder snapshotEnabled(@Nullable Output<Boolean> snapshotEnabled) {
            $.snapshotEnabled = snapshotEnabled;
            return this;
        }

        /**
         * @param snapshotEnabled Defines whether M3DB will create snapshot files for this namespace.
         * 
         * @return builder
         * 
         */
        public Builder snapshotEnabled(Boolean snapshotEnabled) {
            return snapshotEnabled(Output.of(snapshotEnabled));
        }

        /**
         * @param type Type of namespace.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of namespace.
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param writesToCommitLogEnabled Defines whether M3DB will include writes to this namespace in the commit log.
         * 
         * @return builder
         * 
         */
        public Builder writesToCommitLogEnabled(@Nullable Output<Boolean> writesToCommitLogEnabled) {
            $.writesToCommitLogEnabled = writesToCommitLogEnabled;
            return this;
        }

        /**
         * @param writesToCommitLogEnabled Defines whether M3DB will include writes to this namespace in the commit log.
         * 
         * @return builder
         * 
         */
        public Builder writesToCommitLogEnabled(Boolean writesToCommitLogEnabled) {
            return writesToCommitLogEnabled(Output.of(writesToCommitLogEnabled));
        }

        public M3DbNamespaceState build() {
            return $;
        }
    }

}
