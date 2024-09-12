// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Me.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstallationTemplatePartitionSchemePartitionState extends com.pulumi.resources.ResourceArgs {

    public static final InstallationTemplatePartitionSchemePartitionState Empty = new InstallationTemplatePartitionSchemePartitionState();

    /**
     * Partition filesystem. Enum with possibles values:
     * - btrfs
     * - ext3
     * - ext4
     * - ntfs
     * - reiserfs
     * - swap
     * - ufs
     * - xfs
     * - zfs
     * 
     */
    @Import(name="filesystem")
    private @Nullable Output<String> filesystem;

    /**
     * @return Partition filesystem. Enum with possibles values:
     * - btrfs
     * - ext3
     * - ext4
     * - ntfs
     * - reiserfs
     * - swap
     * - ufs
     * - xfs
     * - zfs
     * 
     */
    public Optional<Output<String>> filesystem() {
        return Optional.ofNullable(this.filesystem);
    }

    /**
     * partition mount point.
     * 
     */
    @Import(name="mountpoint")
    private @Nullable Output<String> mountpoint;

    /**
     * @return partition mount point.
     * 
     */
    public Optional<Output<String>> mountpoint() {
        return Optional.ofNullable(this.mountpoint);
    }

    /**
     * step or order. specifies the creation order of the partition on the disk
     * 
     */
    @Import(name="order")
    private @Nullable Output<Integer> order;

    /**
     * @return step or order. specifies the creation order of the partition on the disk
     * 
     */
    public Optional<Output<Integer>> order() {
        return Optional.ofNullable(this.order);
    }

    /**
     * raid partition type. Enum with possible values:
     * - raid0
     * - raid1
     * - raid10
     * - raid5
     * - raid6
     * 
     */
    @Import(name="raid")
    private @Nullable Output<String> raid;

    /**
     * @return raid partition type. Enum with possible values:
     * - raid0
     * - raid1
     * - raid10
     * - raid5
     * - raid6
     * 
     */
    public Optional<Output<String>> raid() {
        return Optional.ofNullable(this.raid);
    }

    /**
     * The partition scheme name.
     * 
     */
    @Import(name="schemeName")
    private @Nullable Output<String> schemeName;

    /**
     * @return The partition scheme name.
     * 
     */
    public Optional<Output<String>> schemeName() {
        return Optional.ofNullable(this.schemeName);
    }

    /**
     * size of partition in MB, 0 =&gt; rest of the space.
     * 
     */
    @Import(name="size")
    private @Nullable Output<Integer> size;

    /**
     * @return size of partition in MB, 0 =&gt; rest of the space.
     * 
     */
    public Optional<Output<Integer>> size() {
        return Optional.ofNullable(this.size);
    }

    /**
     * The template name of the partition scheme.
     * 
     */
    @Import(name="templateName")
    private @Nullable Output<String> templateName;

    /**
     * @return The template name of the partition scheme.
     * 
     */
    public Optional<Output<String>> templateName() {
        return Optional.ofNullable(this.templateName);
    }

    /**
     * partition type. Enum with possible values:
     * - lv
     * - primary
     * - logical
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return partition type. Enum with possible values:
     * - lv
     * - primary
     * - logical
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    /**
     * The volume name needed for proxmox distribution
     * 
     */
    @Import(name="volumeName")
    private @Nullable Output<String> volumeName;

    /**
     * @return The volume name needed for proxmox distribution
     * 
     */
    public Optional<Output<String>> volumeName() {
        return Optional.ofNullable(this.volumeName);
    }

    private InstallationTemplatePartitionSchemePartitionState() {}

    private InstallationTemplatePartitionSchemePartitionState(InstallationTemplatePartitionSchemePartitionState $) {
        this.filesystem = $.filesystem;
        this.mountpoint = $.mountpoint;
        this.order = $.order;
        this.raid = $.raid;
        this.schemeName = $.schemeName;
        this.size = $.size;
        this.templateName = $.templateName;
        this.type = $.type;
        this.volumeName = $.volumeName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstallationTemplatePartitionSchemePartitionState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstallationTemplatePartitionSchemePartitionState $;

        public Builder() {
            $ = new InstallationTemplatePartitionSchemePartitionState();
        }

        public Builder(InstallationTemplatePartitionSchemePartitionState defaults) {
            $ = new InstallationTemplatePartitionSchemePartitionState(Objects.requireNonNull(defaults));
        }

        /**
         * @param filesystem Partition filesystem. Enum with possibles values:
         * - btrfs
         * - ext3
         * - ext4
         * - ntfs
         * - reiserfs
         * - swap
         * - ufs
         * - xfs
         * - zfs
         * 
         * @return builder
         * 
         */
        public Builder filesystem(@Nullable Output<String> filesystem) {
            $.filesystem = filesystem;
            return this;
        }

        /**
         * @param filesystem Partition filesystem. Enum with possibles values:
         * - btrfs
         * - ext3
         * - ext4
         * - ntfs
         * - reiserfs
         * - swap
         * - ufs
         * - xfs
         * - zfs
         * 
         * @return builder
         * 
         */
        public Builder filesystem(String filesystem) {
            return filesystem(Output.of(filesystem));
        }

        /**
         * @param mountpoint partition mount point.
         * 
         * @return builder
         * 
         */
        public Builder mountpoint(@Nullable Output<String> mountpoint) {
            $.mountpoint = mountpoint;
            return this;
        }

        /**
         * @param mountpoint partition mount point.
         * 
         * @return builder
         * 
         */
        public Builder mountpoint(String mountpoint) {
            return mountpoint(Output.of(mountpoint));
        }

        /**
         * @param order step or order. specifies the creation order of the partition on the disk
         * 
         * @return builder
         * 
         */
        public Builder order(@Nullable Output<Integer> order) {
            $.order = order;
            return this;
        }

        /**
         * @param order step or order. specifies the creation order of the partition on the disk
         * 
         * @return builder
         * 
         */
        public Builder order(Integer order) {
            return order(Output.of(order));
        }

        /**
         * @param raid raid partition type. Enum with possible values:
         * - raid0
         * - raid1
         * - raid10
         * - raid5
         * - raid6
         * 
         * @return builder
         * 
         */
        public Builder raid(@Nullable Output<String> raid) {
            $.raid = raid;
            return this;
        }

        /**
         * @param raid raid partition type. Enum with possible values:
         * - raid0
         * - raid1
         * - raid10
         * - raid5
         * - raid6
         * 
         * @return builder
         * 
         */
        public Builder raid(String raid) {
            return raid(Output.of(raid));
        }

        /**
         * @param schemeName The partition scheme name.
         * 
         * @return builder
         * 
         */
        public Builder schemeName(@Nullable Output<String> schemeName) {
            $.schemeName = schemeName;
            return this;
        }

        /**
         * @param schemeName The partition scheme name.
         * 
         * @return builder
         * 
         */
        public Builder schemeName(String schemeName) {
            return schemeName(Output.of(schemeName));
        }

        /**
         * @param size size of partition in MB, 0 =&gt; rest of the space.
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Integer> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size size of partition in MB, 0 =&gt; rest of the space.
         * 
         * @return builder
         * 
         */
        public Builder size(Integer size) {
            return size(Output.of(size));
        }

        /**
         * @param templateName The template name of the partition scheme.
         * 
         * @return builder
         * 
         */
        public Builder templateName(@Nullable Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName The template name of the partition scheme.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        /**
         * @param type partition type. Enum with possible values:
         * - lv
         * - primary
         * - logical
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type partition type. Enum with possible values:
         * - lv
         * - primary
         * - logical
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        /**
         * @param volumeName The volume name needed for proxmox distribution
         * 
         * @return builder
         * 
         */
        public Builder volumeName(@Nullable Output<String> volumeName) {
            $.volumeName = volumeName;
            return this;
        }

        /**
         * @param volumeName The volume name needed for proxmox distribution
         * 
         * @return builder
         * 
         */
        public Builder volumeName(String volumeName) {
            return volumeName(Output.of(volumeName));
        }

        public InstallationTemplatePartitionSchemePartitionState build() {
            return $;
        }
    }

}