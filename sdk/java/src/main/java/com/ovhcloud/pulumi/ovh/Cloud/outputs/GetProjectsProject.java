// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Cloud.outputs;

import com.ovhcloud.pulumi.ovh.Cloud.outputs.GetProjectsProjectIam;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetProjectsProject {
    /**
     * @return Project access
     * 
     */
    private String access;
    /**
     * @return Project creation date
     * 
     */
    private String creationDate;
    /**
     * @return Description of your project
     * 
     */
    private String description;
    /**
     * @return Expiration date of your project. After this date, your project will be deleted
     * 
     */
    private String expiration;
    /**
     * @return IAM resource information
     * 
     */
    private GetProjectsProjectIam iam;
    /**
     * @return Manual quota prevent automatic quota upgrade
     * 
     */
    private Boolean manualQuota;
    /**
     * @return Project order ID
     * 
     */
    private Double orderId;
    /**
     * @return Order plan code
     * 
     */
    private String planCode;
    /**
     * @return Project ID
     * 
     */
    private String projectId;
    /**
     * @return Project name
     * 
     */
    private String projectName;
    /**
     * @return ID of the public cloud project
     * 
     */
    private String serviceName;
    /**
     * @return Current status
     * 
     */
    private String status;
    /**
     * @return Project unleashed
     * 
     */
    private Boolean unleash;

    private GetProjectsProject() {}
    /**
     * @return Project access
     * 
     */
    public String access() {
        return this.access;
    }
    /**
     * @return Project creation date
     * 
     */
    public String creationDate() {
        return this.creationDate;
    }
    /**
     * @return Description of your project
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return Expiration date of your project. After this date, your project will be deleted
     * 
     */
    public String expiration() {
        return this.expiration;
    }
    /**
     * @return IAM resource information
     * 
     */
    public GetProjectsProjectIam iam() {
        return this.iam;
    }
    /**
     * @return Manual quota prevent automatic quota upgrade
     * 
     */
    public Boolean manualQuota() {
        return this.manualQuota;
    }
    /**
     * @return Project order ID
     * 
     */
    public Double orderId() {
        return this.orderId;
    }
    /**
     * @return Order plan code
     * 
     */
    public String planCode() {
        return this.planCode;
    }
    /**
     * @return Project ID
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return Project name
     * 
     */
    public String projectName() {
        return this.projectName;
    }
    /**
     * @return ID of the public cloud project
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Current status
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Project unleashed
     * 
     */
    public Boolean unleash() {
        return this.unleash;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetProjectsProject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String access;
        private String creationDate;
        private String description;
        private String expiration;
        private GetProjectsProjectIam iam;
        private Boolean manualQuota;
        private Double orderId;
        private String planCode;
        private String projectId;
        private String projectName;
        private String serviceName;
        private String status;
        private Boolean unleash;
        public Builder() {}
        public Builder(GetProjectsProject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.access = defaults.access;
    	      this.creationDate = defaults.creationDate;
    	      this.description = defaults.description;
    	      this.expiration = defaults.expiration;
    	      this.iam = defaults.iam;
    	      this.manualQuota = defaults.manualQuota;
    	      this.orderId = defaults.orderId;
    	      this.planCode = defaults.planCode;
    	      this.projectId = defaults.projectId;
    	      this.projectName = defaults.projectName;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
    	      this.unleash = defaults.unleash;
        }

        @CustomType.Setter
        public Builder access(String access) {
            if (access == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "access");
            }
            this.access = access;
            return this;
        }
        @CustomType.Setter
        public Builder creationDate(String creationDate) {
            if (creationDate == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "creationDate");
            }
            this.creationDate = creationDate;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder expiration(String expiration) {
            if (expiration == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "expiration");
            }
            this.expiration = expiration;
            return this;
        }
        @CustomType.Setter
        public Builder iam(GetProjectsProjectIam iam) {
            if (iam == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "iam");
            }
            this.iam = iam;
            return this;
        }
        @CustomType.Setter
        public Builder manualQuota(Boolean manualQuota) {
            if (manualQuota == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "manualQuota");
            }
            this.manualQuota = manualQuota;
            return this;
        }
        @CustomType.Setter
        public Builder orderId(Double orderId) {
            if (orderId == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "orderId");
            }
            this.orderId = orderId;
            return this;
        }
        @CustomType.Setter
        public Builder planCode(String planCode) {
            if (planCode == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "planCode");
            }
            this.planCode = planCode;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder projectName(String projectName) {
            if (projectName == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "projectName");
            }
            this.projectName = projectName;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder unleash(Boolean unleash) {
            if (unleash == null) {
              throw new MissingRequiredPropertyException("GetProjectsProject", "unleash");
            }
            this.unleash = unleash;
            return this;
        }
        public GetProjectsProject build() {
            final var _resultValue = new GetProjectsProject();
            _resultValue.access = access;
            _resultValue.creationDate = creationDate;
            _resultValue.description = description;
            _resultValue.expiration = expiration;
            _resultValue.iam = iam;
            _resultValue.manualQuota = manualQuota;
            _resultValue.orderId = orderId;
            _resultValue.planCode = planCode;
            _resultValue.projectId = projectId;
            _resultValue.projectName = projectName;
            _resultValue.serviceName = serviceName;
            _resultValue.status = status;
            _resultValue.unleash = unleash;
            return _resultValue;
        }
    }
}
