// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetInstanceAddress;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetInstanceAttachedVolume;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInstanceResult {
    /**
     * @return Instance IP addresses
     * 
     */
    private List<GetInstanceAddress> addresses;
    /**
     * @return Volumes attached to the instance
     * 
     */
    private List<GetInstanceAttachedVolume> attachedVolumes;
    /**
     * @return Availability zone of the instance
     * 
     */
    private String availabilityZone;
    /**
     * @return Flavor id
     * 
     */
    private String flavorId;
    /**
     * @return Flavor name
     * 
     */
    private String flavorName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Image id
     * 
     */
    private String imageId;
    private String instanceId;
    /**
     * @return Instance name
     * 
     */
    private String name;
    private String region;
    private String serviceName;
    /**
     * @return SSH Keypair
     * 
     */
    private String sshKey;
    /**
     * @return Instance status
     * 
     */
    private String status;
    /**
     * @return Instance task state
     * 
     */
    private String taskState;

    private GetInstanceResult() {}
    /**
     * @return Instance IP addresses
     * 
     */
    public List<GetInstanceAddress> addresses() {
        return this.addresses;
    }
    /**
     * @return Volumes attached to the instance
     * 
     */
    public List<GetInstanceAttachedVolume> attachedVolumes() {
        return this.attachedVolumes;
    }
    /**
     * @return Availability zone of the instance
     * 
     */
    public String availabilityZone() {
        return this.availabilityZone;
    }
    /**
     * @return Flavor id
     * 
     */
    public String flavorId() {
        return this.flavorId;
    }
    /**
     * @return Flavor name
     * 
     */
    public String flavorName() {
        return this.flavorName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Image id
     * 
     */
    public String imageId() {
        return this.imageId;
    }
    public String instanceId() {
        return this.instanceId;
    }
    /**
     * @return Instance name
     * 
     */
    public String name() {
        return this.name;
    }
    public String region() {
        return this.region;
    }
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return SSH Keypair
     * 
     */
    public String sshKey() {
        return this.sshKey;
    }
    /**
     * @return Instance status
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return Instance task state
     * 
     */
    public String taskState() {
        return this.taskState;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstanceResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetInstanceAddress> addresses;
        private List<GetInstanceAttachedVolume> attachedVolumes;
        private String availabilityZone;
        private String flavorId;
        private String flavorName;
        private String id;
        private String imageId;
        private String instanceId;
        private String name;
        private String region;
        private String serviceName;
        private String sshKey;
        private String status;
        private String taskState;
        public Builder() {}
        public Builder(GetInstanceResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.addresses = defaults.addresses;
    	      this.attachedVolumes = defaults.attachedVolumes;
    	      this.availabilityZone = defaults.availabilityZone;
    	      this.flavorId = defaults.flavorId;
    	      this.flavorName = defaults.flavorName;
    	      this.id = defaults.id;
    	      this.imageId = defaults.imageId;
    	      this.instanceId = defaults.instanceId;
    	      this.name = defaults.name;
    	      this.region = defaults.region;
    	      this.serviceName = defaults.serviceName;
    	      this.sshKey = defaults.sshKey;
    	      this.status = defaults.status;
    	      this.taskState = defaults.taskState;
        }

        @CustomType.Setter
        public Builder addresses(List<GetInstanceAddress> addresses) {
            if (addresses == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "addresses");
            }
            this.addresses = addresses;
            return this;
        }
        public Builder addresses(GetInstanceAddress... addresses) {
            return addresses(List.of(addresses));
        }
        @CustomType.Setter
        public Builder attachedVolumes(List<GetInstanceAttachedVolume> attachedVolumes) {
            if (attachedVolumes == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "attachedVolumes");
            }
            this.attachedVolumes = attachedVolumes;
            return this;
        }
        public Builder attachedVolumes(GetInstanceAttachedVolume... attachedVolumes) {
            return attachedVolumes(List.of(attachedVolumes));
        }
        @CustomType.Setter
        public Builder availabilityZone(String availabilityZone) {
            if (availabilityZone == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "availabilityZone");
            }
            this.availabilityZone = availabilityZone;
            return this;
        }
        @CustomType.Setter
        public Builder flavorId(String flavorId) {
            if (flavorId == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "flavorId");
            }
            this.flavorId = flavorId;
            return this;
        }
        @CustomType.Setter
        public Builder flavorName(String flavorName) {
            if (flavorName == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "flavorName");
            }
            this.flavorName = flavorName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder imageId(String imageId) {
            if (imageId == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "imageId");
            }
            this.imageId = imageId;
            return this;
        }
        @CustomType.Setter
        public Builder instanceId(String instanceId) {
            if (instanceId == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "instanceId");
            }
            this.instanceId = instanceId;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder sshKey(String sshKey) {
            if (sshKey == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "sshKey");
            }
            this.sshKey = sshKey;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder taskState(String taskState) {
            if (taskState == null) {
              throw new MissingRequiredPropertyException("GetInstanceResult", "taskState");
            }
            this.taskState = taskState;
            return this;
        }
        public GetInstanceResult build() {
            final var _resultValue = new GetInstanceResult();
            _resultValue.addresses = addresses;
            _resultValue.attachedVolumes = attachedVolumes;
            _resultValue.availabilityZone = availabilityZone;
            _resultValue.flavorId = flavorId;
            _resultValue.flavorName = flavorName;
            _resultValue.id = id;
            _resultValue.imageId = imageId;
            _resultValue.instanceId = instanceId;
            _resultValue.name = name;
            _resultValue.region = region;
            _resultValue.serviceName = serviceName;
            _resultValue.sshKey = sshKey;
            _resultValue.status = status;
            _resultValue.taskState = taskState;
            return _resultValue;
        }
    }
}
