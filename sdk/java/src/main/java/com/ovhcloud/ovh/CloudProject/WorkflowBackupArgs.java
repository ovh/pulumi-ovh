// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.CloudProject;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class WorkflowBackupArgs extends com.pulumi.resources.ResourceArgs {

    public static final WorkflowBackupArgs Empty = new WorkflowBackupArgs();

    /**
     * The name of the backup files that are created. If empty, the `name` attribute is used.
     * 
     */
    @Import(name="backupName")
    private @Nullable Output<String> backupName;

    /**
     * @return The name of the backup files that are created. If empty, the `name` attribute is used.
     * 
     */
    public Optional<Output<String>> backupName() {
        return Optional.ofNullable(this.backupName);
    }

    /**
     * The cron periodicity at which the backup workflow is scheduled
     * 
     * * `instanceId` the id of the instance to back up
     * 
     */
    @Import(name="cron", required=true)
    private Output<String> cron;

    /**
     * @return The cron periodicity at which the backup workflow is scheduled
     * 
     * * `instanceId` the id of the instance to back up
     * 
     */
    public Output<String> cron() {
        return this.cron;
    }

    @Import(name="instanceId", required=true)
    private Output<String> instanceId;

    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     * The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
     * 
     */
    @Import(name="maxExecutionCount")
    private @Nullable Output<Integer> maxExecutionCount;

    /**
     * @return The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
     * 
     */
    public Optional<Output<Integer>> maxExecutionCount() {
        return Optional.ofNullable(this.maxExecutionCount);
    }

    /**
     * The worflow name that is used in the UI
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The worflow name that is used in the UI
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The name of the openstack region.
     * 
     */
    @Import(name="regionName", required=true)
    private Output<String> regionName;

    /**
     * @return The name of the openstack region.
     * 
     */
    public Output<String> regionName() {
        return this.regionName;
    }

    /**
     * The number of backup that are retained.
     * 
     */
    @Import(name="rotation", required=true)
    private Output<Integer> rotation;

    /**
     * @return The number of backup that are retained.
     * 
     */
    public Output<Integer> rotation() {
        return this.rotation;
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

    private WorkflowBackupArgs() {}

    private WorkflowBackupArgs(WorkflowBackupArgs $) {
        this.backupName = $.backupName;
        this.cron = $.cron;
        this.instanceId = $.instanceId;
        this.maxExecutionCount = $.maxExecutionCount;
        this.name = $.name;
        this.regionName = $.regionName;
        this.rotation = $.rotation;
        this.serviceName = $.serviceName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(WorkflowBackupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private WorkflowBackupArgs $;

        public Builder() {
            $ = new WorkflowBackupArgs();
        }

        public Builder(WorkflowBackupArgs defaults) {
            $ = new WorkflowBackupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backupName The name of the backup files that are created. If empty, the `name` attribute is used.
         * 
         * @return builder
         * 
         */
        public Builder backupName(@Nullable Output<String> backupName) {
            $.backupName = backupName;
            return this;
        }

        /**
         * @param backupName The name of the backup files that are created. If empty, the `name` attribute is used.
         * 
         * @return builder
         * 
         */
        public Builder backupName(String backupName) {
            return backupName(Output.of(backupName));
        }

        /**
         * @param cron The cron periodicity at which the backup workflow is scheduled
         * 
         * * `instanceId` the id of the instance to back up
         * 
         * @return builder
         * 
         */
        public Builder cron(Output<String> cron) {
            $.cron = cron;
            return this;
        }

        /**
         * @param cron The cron periodicity at which the backup workflow is scheduled
         * 
         * * `instanceId` the id of the instance to back up
         * 
         * @return builder
         * 
         */
        public Builder cron(String cron) {
            return cron(Output.of(cron));
        }

        public Builder instanceId(Output<String> instanceId) {
            $.instanceId = instanceId;
            return this;
        }

        public Builder instanceId(String instanceId) {
            return instanceId(Output.of(instanceId));
        }

        /**
         * @param maxExecutionCount The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
         * 
         * @return builder
         * 
         */
        public Builder maxExecutionCount(@Nullable Output<Integer> maxExecutionCount) {
            $.maxExecutionCount = maxExecutionCount;
            return this;
        }

        /**
         * @param maxExecutionCount The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
         * 
         * @return builder
         * 
         */
        public Builder maxExecutionCount(Integer maxExecutionCount) {
            return maxExecutionCount(Output.of(maxExecutionCount));
        }

        /**
         * @param name The worflow name that is used in the UI
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The worflow name that is used in the UI
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param regionName The name of the openstack region.
         * 
         * @return builder
         * 
         */
        public Builder regionName(Output<String> regionName) {
            $.regionName = regionName;
            return this;
        }

        /**
         * @param regionName The name of the openstack region.
         * 
         * @return builder
         * 
         */
        public Builder regionName(String regionName) {
            return regionName(Output.of(regionName));
        }

        /**
         * @param rotation The number of backup that are retained.
         * 
         * @return builder
         * 
         */
        public Builder rotation(Output<Integer> rotation) {
            $.rotation = rotation;
            return this;
        }

        /**
         * @param rotation The number of backup that are retained.
         * 
         * @return builder
         * 
         */
        public Builder rotation(Integer rotation) {
            return rotation(Output.of(rotation));
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

        public WorkflowBackupArgs build() {
            if ($.cron == null) {
                throw new MissingRequiredPropertyException("WorkflowBackupArgs", "cron");
            }
            if ($.instanceId == null) {
                throw new MissingRequiredPropertyException("WorkflowBackupArgs", "instanceId");
            }
            if ($.regionName == null) {
                throw new MissingRequiredPropertyException("WorkflowBackupArgs", "regionName");
            }
            if ($.rotation == null) {
                throw new MissingRequiredPropertyException("WorkflowBackupArgs", "rotation");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("WorkflowBackupArgs", "serviceName");
            }
            return $;
        }
    }

}