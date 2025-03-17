// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.inputs;

import com.ovhcloud.pulumi.ovh.Dedicated.inputs.ServerStoragePartitioningLayoutExtrasArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ServerStoragePartitioningLayoutArgs extends com.pulumi.resources.ResourceArgs {

    public static final ServerStoragePartitioningLayoutArgs Empty = new ServerStoragePartitioningLayoutArgs();

    /**
     * Partition extras parameters
     * 
     */
    @Import(name="extras")
    private @Nullable Output<ServerStoragePartitioningLayoutExtrasArgs> extras;

    /**
     * @return Partition extras parameters
     * 
     */
    public Optional<Output<ServerStoragePartitioningLayoutExtrasArgs>> extras() {
        return Optional.ofNullable(this.extras);
    }

    /**
     * File system type
     * 
     */
    @Import(name="fileSystem", required=true)
    private Output<String> fileSystem;

    /**
     * @return File system type
     * 
     */
    public Output<String> fileSystem() {
        return this.fileSystem;
    }

    /**
     * Mount point
     * 
     */
    @Import(name="mountPoint", required=true)
    private Output<String> mountPoint;

    /**
     * @return Mount point
     * 
     */
    public Output<String> mountPoint() {
        return this.mountPoint;
    }

    /**
     * Software raid type (default is 1)
     * 
     */
    @Import(name="raidLevel")
    private @Nullable Output<Double> raidLevel;

    /**
     * @return Software raid type (default is 1)
     * 
     */
    public Optional<Output<Double>> raidLevel() {
        return Optional.ofNullable(this.raidLevel);
    }

    /**
     * Partition size in MiB (default value is 0 which means to fill the disk with that partition)
     * 
     */
    @Import(name="size")
    private @Nullable Output<Double> size;

    /**
     * @return Partition size in MiB (default value is 0 which means to fill the disk with that partition)
     * 
     */
    public Optional<Output<Double>> size() {
        return Optional.ofNullable(this.size);
    }

    private ServerStoragePartitioningLayoutArgs() {}

    private ServerStoragePartitioningLayoutArgs(ServerStoragePartitioningLayoutArgs $) {
        this.extras = $.extras;
        this.fileSystem = $.fileSystem;
        this.mountPoint = $.mountPoint;
        this.raidLevel = $.raidLevel;
        this.size = $.size;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ServerStoragePartitioningLayoutArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ServerStoragePartitioningLayoutArgs $;

        public Builder() {
            $ = new ServerStoragePartitioningLayoutArgs();
        }

        public Builder(ServerStoragePartitioningLayoutArgs defaults) {
            $ = new ServerStoragePartitioningLayoutArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param extras Partition extras parameters
         * 
         * @return builder
         * 
         */
        public Builder extras(@Nullable Output<ServerStoragePartitioningLayoutExtrasArgs> extras) {
            $.extras = extras;
            return this;
        }

        /**
         * @param extras Partition extras parameters
         * 
         * @return builder
         * 
         */
        public Builder extras(ServerStoragePartitioningLayoutExtrasArgs extras) {
            return extras(Output.of(extras));
        }

        /**
         * @param fileSystem File system type
         * 
         * @return builder
         * 
         */
        public Builder fileSystem(Output<String> fileSystem) {
            $.fileSystem = fileSystem;
            return this;
        }

        /**
         * @param fileSystem File system type
         * 
         * @return builder
         * 
         */
        public Builder fileSystem(String fileSystem) {
            return fileSystem(Output.of(fileSystem));
        }

        /**
         * @param mountPoint Mount point
         * 
         * @return builder
         * 
         */
        public Builder mountPoint(Output<String> mountPoint) {
            $.mountPoint = mountPoint;
            return this;
        }

        /**
         * @param mountPoint Mount point
         * 
         * @return builder
         * 
         */
        public Builder mountPoint(String mountPoint) {
            return mountPoint(Output.of(mountPoint));
        }

        /**
         * @param raidLevel Software raid type (default is 1)
         * 
         * @return builder
         * 
         */
        public Builder raidLevel(@Nullable Output<Double> raidLevel) {
            $.raidLevel = raidLevel;
            return this;
        }

        /**
         * @param raidLevel Software raid type (default is 1)
         * 
         * @return builder
         * 
         */
        public Builder raidLevel(Double raidLevel) {
            return raidLevel(Output.of(raidLevel));
        }

        /**
         * @param size Partition size in MiB (default value is 0 which means to fill the disk with that partition)
         * 
         * @return builder
         * 
         */
        public Builder size(@Nullable Output<Double> size) {
            $.size = size;
            return this;
        }

        /**
         * @param size Partition size in MiB (default value is 0 which means to fill the disk with that partition)
         * 
         * @return builder
         * 
         */
        public Builder size(Double size) {
            return size(Output.of(size));
        }

        public ServerStoragePartitioningLayoutArgs build() {
            if ($.fileSystem == null) {
                throw new MissingRequiredPropertyException("ServerStoragePartitioningLayoutArgs", "fileSystem");
            }
            if ($.mountPoint == null) {
                throw new MissingRequiredPropertyException("ServerStoragePartitioningLayoutArgs", "mountPoint");
            }
            return $;
        }
    }

}
