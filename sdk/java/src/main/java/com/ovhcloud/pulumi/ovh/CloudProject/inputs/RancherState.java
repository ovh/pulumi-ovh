// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RancherCurrentStateArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RancherCurrentTaskArgs;
import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RancherTargetSpecArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RancherState extends com.pulumi.resources.ResourceArgs {

    public static final RancherState Empty = new RancherState();

    /**
     * Date of the managed Rancher service creation
     * 
     */
    @Import(name="createdAt")
    private @Nullable Output<String> createdAt;

    /**
     * @return Date of the managed Rancher service creation
     * 
     */
    public Optional<Output<String>> createdAt() {
        return Optional.ofNullable(this.createdAt);
    }

    /**
     * Current configuration applied to the managed Rancher service
     * 
     */
    @Import(name="currentState")
    private @Nullable Output<RancherCurrentStateArgs> currentState;

    /**
     * @return Current configuration applied to the managed Rancher service
     * 
     */
    public Optional<Output<RancherCurrentStateArgs>> currentState() {
        return Optional.ofNullable(this.currentState);
    }

    /**
     * Asynchronous operations ongoing on the managed Rancher service
     * 
     */
    @Import(name="currentTasks")
    private @Nullable Output<List<RancherCurrentTaskArgs>> currentTasks;

    /**
     * @return Asynchronous operations ongoing on the managed Rancher service
     * 
     */
    public Optional<Output<List<RancherCurrentTaskArgs>>> currentTasks() {
        return Optional.ofNullable(this.currentTasks);
    }

    /**
     * Project ID
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return Project ID
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
     * 
     */
    @Import(name="resourceStatus")
    private @Nullable Output<String> resourceStatus;

    /**
     * @return Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
     * 
     */
    public Optional<Output<String>> resourceStatus() {
        return Optional.ofNullable(this.resourceStatus);
    }

    /**
     * Target specification for the managed Rancher service
     * 
     */
    @Import(name="targetSpec")
    private @Nullable Output<RancherTargetSpecArgs> targetSpec;

    /**
     * @return Target specification for the managed Rancher service
     * 
     */
    public Optional<Output<RancherTargetSpecArgs>> targetSpec() {
        return Optional.ofNullable(this.targetSpec);
    }

    /**
     * Date of the last managed Rancher service update
     * 
     */
    @Import(name="updatedAt")
    private @Nullable Output<String> updatedAt;

    /**
     * @return Date of the last managed Rancher service update
     * 
     */
    public Optional<Output<String>> updatedAt() {
        return Optional.ofNullable(this.updatedAt);
    }

    private RancherState() {}

    private RancherState(RancherState $) {
        this.createdAt = $.createdAt;
        this.currentState = $.currentState;
        this.currentTasks = $.currentTasks;
        this.projectId = $.projectId;
        this.resourceStatus = $.resourceStatus;
        this.targetSpec = $.targetSpec;
        this.updatedAt = $.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RancherState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RancherState $;

        public Builder() {
            $ = new RancherState();
        }

        public Builder(RancherState defaults) {
            $ = new RancherState(Objects.requireNonNull(defaults));
        }

        /**
         * @param createdAt Date of the managed Rancher service creation
         * 
         * @return builder
         * 
         */
        public Builder createdAt(@Nullable Output<String> createdAt) {
            $.createdAt = createdAt;
            return this;
        }

        /**
         * @param createdAt Date of the managed Rancher service creation
         * 
         * @return builder
         * 
         */
        public Builder createdAt(String createdAt) {
            return createdAt(Output.of(createdAt));
        }

        /**
         * @param currentState Current configuration applied to the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder currentState(@Nullable Output<RancherCurrentStateArgs> currentState) {
            $.currentState = currentState;
            return this;
        }

        /**
         * @param currentState Current configuration applied to the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder currentState(RancherCurrentStateArgs currentState) {
            return currentState(Output.of(currentState));
        }

        /**
         * @param currentTasks Asynchronous operations ongoing on the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder currentTasks(@Nullable Output<List<RancherCurrentTaskArgs>> currentTasks) {
            $.currentTasks = currentTasks;
            return this;
        }

        /**
         * @param currentTasks Asynchronous operations ongoing on the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder currentTasks(List<RancherCurrentTaskArgs> currentTasks) {
            return currentTasks(Output.of(currentTasks));
        }

        /**
         * @param currentTasks Asynchronous operations ongoing on the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder currentTasks(RancherCurrentTaskArgs... currentTasks) {
            return currentTasks(List.of(currentTasks));
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param resourceStatus Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
         * 
         * @return builder
         * 
         */
        public Builder resourceStatus(@Nullable Output<String> resourceStatus) {
            $.resourceStatus = resourceStatus;
            return this;
        }

        /**
         * @param resourceStatus Reflects the readiness of the managed Rancher service. A new target specification request will be accepted only in `READY` status
         * 
         * @return builder
         * 
         */
        public Builder resourceStatus(String resourceStatus) {
            return resourceStatus(Output.of(resourceStatus));
        }

        /**
         * @param targetSpec Target specification for the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder targetSpec(@Nullable Output<RancherTargetSpecArgs> targetSpec) {
            $.targetSpec = targetSpec;
            return this;
        }

        /**
         * @param targetSpec Target specification for the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder targetSpec(RancherTargetSpecArgs targetSpec) {
            return targetSpec(Output.of(targetSpec));
        }

        /**
         * @param updatedAt Date of the last managed Rancher service update
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(@Nullable Output<String> updatedAt) {
            $.updatedAt = updatedAt;
            return this;
        }

        /**
         * @param updatedAt Date of the last managed Rancher service update
         * 
         * @return builder
         * 
         */
        public Builder updatedAt(String updatedAt) {
            return updatedAt(Output.of(updatedAt));
        }

        public RancherState build() {
            return $;
        }
    }

}
