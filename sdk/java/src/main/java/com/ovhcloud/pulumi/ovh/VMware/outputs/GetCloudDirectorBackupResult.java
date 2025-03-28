// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.VMware.outputs;

import com.ovhcloud.pulumi.ovh.VMware.outputs.GetCloudDirectorBackupCurrentState;
import com.ovhcloud.pulumi.ovh.VMware.outputs.GetCloudDirectorBackupCurrentTask;
import com.ovhcloud.pulumi.ovh.VMware.outputs.GetCloudDirectorBackupIam;
import com.ovhcloud.pulumi.ovh.VMware.outputs.GetCloudDirectorBackupTargetSpec;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetCloudDirectorBackupResult {
    /**
     * @return Backup ID
     * 
     */
    private String backupId;
    /**
     * @return Datetime when backup was enabled
     * 
     */
    private String createdAt;
    /**
     * @return VMware Cloud Director Backup service current state
     * 
     */
    private GetCloudDirectorBackupCurrentState currentState;
    /**
     * @return Asynchronous operations ongoing on the VMware Cloud Director organization backup service
     * 
     */
    private List<GetCloudDirectorBackupCurrentTask> currentTasks;
    /**
     * @return IAM resource metadata
     * 
     */
    private GetCloudDirectorBackupIam iam;
    /**
     * @return Unique identifier of the VMware Cloud Director backup
     * 
     */
    private String id;
    /**
     * @return Reflects the readiness of the VMware Cloud Director organization backup service
     * 
     */
    private String resourceStatus;
    /**
     * @return VMware Cloud Director Backup target spec
     * 
     */
    private GetCloudDirectorBackupTargetSpec targetSpec;
    /**
     * @return Datetime when backup is modified
     * 
     */
    private String updatedAt;

    private GetCloudDirectorBackupResult() {}
    /**
     * @return Backup ID
     * 
     */
    public String backupId() {
        return this.backupId;
    }
    /**
     * @return Datetime when backup was enabled
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return VMware Cloud Director Backup service current state
     * 
     */
    public GetCloudDirectorBackupCurrentState currentState() {
        return this.currentState;
    }
    /**
     * @return Asynchronous operations ongoing on the VMware Cloud Director organization backup service
     * 
     */
    public List<GetCloudDirectorBackupCurrentTask> currentTasks() {
        return this.currentTasks;
    }
    /**
     * @return IAM resource metadata
     * 
     */
    public GetCloudDirectorBackupIam iam() {
        return this.iam;
    }
    /**
     * @return Unique identifier of the VMware Cloud Director backup
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Reflects the readiness of the VMware Cloud Director organization backup service
     * 
     */
    public String resourceStatus() {
        return this.resourceStatus;
    }
    /**
     * @return VMware Cloud Director Backup target spec
     * 
     */
    public GetCloudDirectorBackupTargetSpec targetSpec() {
        return this.targetSpec;
    }
    /**
     * @return Datetime when backup is modified
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudDirectorBackupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String backupId;
        private String createdAt;
        private GetCloudDirectorBackupCurrentState currentState;
        private List<GetCloudDirectorBackupCurrentTask> currentTasks;
        private GetCloudDirectorBackupIam iam;
        private String id;
        private String resourceStatus;
        private GetCloudDirectorBackupTargetSpec targetSpec;
        private String updatedAt;
        public Builder() {}
        public Builder(GetCloudDirectorBackupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backupId = defaults.backupId;
    	      this.createdAt = defaults.createdAt;
    	      this.currentState = defaults.currentState;
    	      this.currentTasks = defaults.currentTasks;
    	      this.iam = defaults.iam;
    	      this.id = defaults.id;
    	      this.resourceStatus = defaults.resourceStatus;
    	      this.targetSpec = defaults.targetSpec;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder backupId(String backupId) {
            if (backupId == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "backupId");
            }
            this.backupId = backupId;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder currentState(GetCloudDirectorBackupCurrentState currentState) {
            if (currentState == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "currentState");
            }
            this.currentState = currentState;
            return this;
        }
        @CustomType.Setter
        public Builder currentTasks(List<GetCloudDirectorBackupCurrentTask> currentTasks) {
            if (currentTasks == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "currentTasks");
            }
            this.currentTasks = currentTasks;
            return this;
        }
        public Builder currentTasks(GetCloudDirectorBackupCurrentTask... currentTasks) {
            return currentTasks(List.of(currentTasks));
        }
        @CustomType.Setter
        public Builder iam(GetCloudDirectorBackupIam iam) {
            if (iam == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "iam");
            }
            this.iam = iam;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder resourceStatus(String resourceStatus) {
            if (resourceStatus == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "resourceStatus");
            }
            this.resourceStatus = resourceStatus;
            return this;
        }
        @CustomType.Setter
        public Builder targetSpec(GetCloudDirectorBackupTargetSpec targetSpec) {
            if (targetSpec == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "targetSpec");
            }
            this.targetSpec = targetSpec;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorBackupResult", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        public GetCloudDirectorBackupResult build() {
            final var _resultValue = new GetCloudDirectorBackupResult();
            _resultValue.backupId = backupId;
            _resultValue.createdAt = createdAt;
            _resultValue.currentState = currentState;
            _resultValue.currentTasks = currentTasks;
            _resultValue.iam = iam;
            _resultValue.id = id;
            _resultValue.resourceStatus = resourceStatus;
            _resultValue.targetSpec = targetSpec;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
