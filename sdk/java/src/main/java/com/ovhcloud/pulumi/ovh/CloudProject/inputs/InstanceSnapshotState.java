// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstanceSnapshotState extends com.pulumi.resources.ResourceArgs {

    public static final InstanceSnapshotState Empty = new InstanceSnapshotState();

    /**
     * Image creation date
     * 
     */
    @Import(name="creationDate")
    private @Nullable Output<String> creationDate;

    /**
     * @return Image creation date
     * 
     */
    public Optional<Output<String>> creationDate() {
        return Optional.ofNullable(this.creationDate);
    }

    /**
     * Image usable only for this type of flavor if not null
     * 
     */
    @Import(name="flavorType")
    private @Nullable Output<String> flavorType;

    /**
     * @return Image usable only for this type of flavor if not null
     * 
     */
    public Optional<Output<String>> flavorType() {
        return Optional.ofNullable(this.flavorType);
    }

    /**
     * Instance ID
     * 
     */
    @Import(name="instanceId")
    private @Nullable Output<String> instanceId;

    /**
     * @return Instance ID
     * 
     */
    public Optional<Output<String>> instanceId() {
        return Optional.ofNullable(this.instanceId);
    }

    /**
     * Minimum disks required to use image
     * 
     */
    @Import(name="minDisk")
    private @Nullable Output<Double> minDisk;

    /**
     * @return Minimum disks required to use image
     * 
     */
    public Optional<Output<Double>> minDisk() {
        return Optional.ofNullable(this.minDisk);
    }

    /**
     * Minimum RAM required to use image
     * 
     */
    @Import(name="minRam")
    private @Nullable Output<Double> minRam;

    /**
     * @return Minimum RAM required to use image
     * 
     */
    public Optional<Output<Double>> minRam() {
        return Optional.ofNullable(this.minRam);
    }

    /**
     * Snapshot name
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Snapshot name
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Order plan code
     * 
     */
    @Import(name="planCode")
    private @Nullable Output<String> planCode;

    /**
     * @return Order plan code
     * 
     */
    public Optional<Output<String>> planCode() {
        return Optional.ofNullable(this.planCode);
    }

    /**
     * Image region
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return Image region
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * Service name
     * 
     */
    @Import(name="serviceName")
    private @Nullable Output<String> serviceName;

    /**
     * @return Service name
     * 
     */
    public Optional<Output<String>> serviceName() {
        return Optional.ofNullable(this.serviceName);
    }

    /**
     * Image size (in GiB)
     * 
     */
    @Import(name="size")
    private @Nullable Output<Double> size;

    /**
     * @return Image size (in GiB)
     * 
     */
    public Optional<Output<Double>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * Image status
     * 
     */
    @Import(name="status")
    private @Nullable Output<String> status;

    /**
     * @return Image status
     * 
     */
    public Optional<Output<String>> status() {
        return Optional.ofNullable(this.status);
    }

    /**
     * Tags about the image
     * 
     */
    @Import(name="tags")
    private @Nullable Output<List<String>> tags;

    /**
     * @return Tags about the image
     * 
     */
    public Optional<Output<List<String>>> tags() {
        return Optional.ofNullable(this.tags);
    }

    /**
     * Image type
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Image type
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * User to connect with
     * 
     */
    @Import(name="user")
    private @Nullable Output<String> user;

    /**
     * @return User to connect with
     * 
     */
    public Optional<Output<String>> user() {
        return Optional.ofNullable(this.user);
    }

    /**
     * Image visibility
     * 
     */
    @Import(name="visibility")
    private @Nullable Output<String> visibility;

    /**
     * @return Image visibility
     * 
     */
    public Optional<Output<String>> visibility() {
        return Optional.ofNullable(this.visibility);
    }

    private InstanceSnapshotState() {}

    private InstanceSnapshotState(InstanceSnapshotState $) {
        this.creationDate = $.creationDate;
        this.flavorType = $.flavorType;
        this.instanceId = $.instanceId;
        this.minDisk = $.minDisk;
        this.minRam = $.minRam;
        this.name = $.name;
        this.planCode = $.planCode;
        this.region = $.region;
        this.serviceName = $.serviceName;
        this.size = $.size;
        this.status = $.status;
        this.tags = $.tags;
        this.type = $.type;
        this.user = $.user;
        this.visibility = $.visibility;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstanceSnapshotState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstanceSnapshotState $;

        public Builder() {
            $ = new InstanceSnapshotState();
        }

        public Builder(InstanceSnapshotState defaults) {
            $ = new InstanceSnapshotState(Objects.requireNonNull(defaults));
        }

        /**
         * @param creationDate Image creation date
         * 
         * @return builder
         * 
         */
        public Builder creationDate(@Nullable Output<String> creationDate) {
            $.creationDate = creationDate;
            return this;
        }

        /**
         * @param creationDate Image creation date
         * 
         * @return builder
         * 
         */
        public Builder creationDate(String creationDate) {
            return creationDate(Output.of(creationDate));
        }

        /**
         * @param flavorType Image usable only for this type of flavor if not null
         * 
         * @return builder
         * 
         */
        public Builder flavorType(@Nullable Output<String> flavorType) {
            $.flavorType = flavorType;
            return this;
        }

        /**
         * @param flavorType Image usable only for this type of flavor if not null
         * 
         * @return builder
         * 
         */
        public Builder flavorType(String flavorType) {
            return flavorType(Output.of(flavorType));
        }

        /**
         * @param instanceId Instance ID
         * 
         * @return builder
         * 
         */
        public Builder instanceId(@Nullable Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        /**
         * @param instanceId Instance ID
         * 
         * @return builder
         * 
         */
        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param minDisk Minimum disks required to use image
         * 
         * @return builder
         * 
         */
        public Builder minDisk(@Nullable Output<Double> minDisk) {
            $.minDisk = minDisk;
            return this;
        }

        /**
         * @param minDisk Minimum disks required to use image
         * 
         * @return builder
         * 
         */
        public Builder minDisk(Double minDisk) {
            return minDisk(Output.of(minDisk));
        }

        /**
         * @param minRam Minimum RAM required to use image
         * 
         * @return builder
         * 
         */
        public Builder minRam(@Nullable Output<Double> minRam) {
            $.minRam = minRam;
            return this;
        }

        /**
         * @param minRam Minimum RAM required to use image
         * 
         * @return builder
         * 
         */
        public Builder minRam(Double minRam) {
            return minRam(Output.of(minRam));
        }

        /**
         * @param name Snapshot name
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Snapshot name
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param planCode Order plan code
         * 
         * @return builder
         * 
         */
        public Builder planCode(@Nullable Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode Order plan code
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        /**
         * @param region Image region
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region Image region
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(@Nullable Output<String> serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param serviceName Service name
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            return serviceName(Output.of(serviceName));
        }

        /**
         * @param size Image size (in GiB)
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Double> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Image size (in GiB)
         * 
         * @return builder
         * 
         */
        public Builder size(Double size) {
            return size(Output.of(size));
        }

        /**
         * @param status Image status
         * 
         * @return builder
         * 
         */
        public Builder status(@Nullable Output<String> status) {
            $.status = status;
            return this;
        }

        /**
         * @param status Image status
         * 
         * @return builder
         * 
         */
        public Builder status(String status) {
            return status(Output.of(status));
        }

        /**
         * @param tags Tags about the image
         * 
         * @return builder
         * 
         */
        public Builder tags(@Nullable Output<List<String>> tags) {
            $.tags = tags;
            return this;
        }

        /**
         * @param tags Tags about the image
         * 
         * @return builder
         * 
         */
        public Builder tags(List<String> tags) {
            return tags(Output.of(tags));
        }

        /**
         * @param tags Tags about the image
         * 
         * @return builder
         * 
         */
        public Builder tags(String... tags) {
            return tags(List.of(tags));
        }

        /**
         * @param type Image type
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Image type
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param user User to connect with
         * 
         * @return builder
         * 
         */
        public Builder user(@Nullable Output<String> user) {
            $.user = user;
            return this;
        }

        /**
         * @param user User to connect with
         * 
         * @return builder
         * 
         */
        public Builder user(String user) {
            return user(Output.of(user));
        }

        /**
         * @param visibility Image visibility
         * 
         * @return builder
         * 
         */
        public Builder visibility(@Nullable Output<String> visibility) {
            $.visibility = visibility;
            return this;
        }

        /**
         * @param visibility Image visibility
         * 
         * @return builder
         * 
         */
        public Builder visibility(String visibility) {
            return visibility(Output.of(visibility));
        }

        public InstanceSnapshotState build() {
            return $;
        }
    }

}
