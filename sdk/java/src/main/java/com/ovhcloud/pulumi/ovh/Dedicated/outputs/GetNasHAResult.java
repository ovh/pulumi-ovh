// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Dedicated.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetNasHAResult {
    /**
     * @return the URN of the HA-NAS instance
     * 
     */
    private String NasHAURN;
    /**
     * @return True, if partition creation is allowed on this HA-NAS
     * 
     */
    private Boolean canCreatePartition;
    /**
     * @return The name you give to the HA-NAS
     * 
     */
    private String customName;
    /**
     * @return area of HA-NAS
     * 
     */
    private String datacenter;
    /**
     * @return the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
     * 
     */
    private String diskType;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Access IP of HA-NAS
     * 
     */
    private String ip;
    /**
     * @return Send an email to customer if any issue is detected
     * 
     */
    private Boolean monitored;
    /**
     * @return the list of the HA-NAS partitions name
     * 
     */
    private List<String> partitionsLists;
    /**
     * @return The storage service name
     * 
     */
    private String serviceName;
    /**
     * @return percentage of HA-NAS space used in %
     * 
     */
    private Double zpoolCapacity;
    /**
     * @return the size of the HA-NAS in GB
     * 
     */
    private Double zpoolSize;

    private GetNasHAResult() {}
    /**
     * @return the URN of the HA-NAS instance
     * 
     */
    public String NasHAURN() {
        return this.NasHAURN;
    }
    /**
     * @return True, if partition creation is allowed on this HA-NAS
     * 
     */
    public Boolean canCreatePartition() {
        return this.canCreatePartition;
    }
    /**
     * @return The name you give to the HA-NAS
     * 
     */
    public String customName() {
        return this.customName;
    }
    /**
     * @return area of HA-NAS
     * 
     */
    public String datacenter() {
        return this.datacenter;
    }
    /**
     * @return the disk type of the HA-NAS. Possible values are: `hdd`, `ssd`, `nvme`
     * 
     */
    public String diskType() {
        return this.diskType;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Access IP of HA-NAS
     * 
     */
    public String ip() {
        return this.ip;
    }
    /**
     * @return Send an email to customer if any issue is detected
     * 
     */
    public Boolean monitored() {
        return this.monitored;
    }
    /**
     * @return the list of the HA-NAS partitions name
     * 
     */
    public List<String> partitionsLists() {
        return this.partitionsLists;
    }
    /**
     * @return The storage service name
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return percentage of HA-NAS space used in %
     * 
     */
    public Double zpoolCapacity() {
        return this.zpoolCapacity;
    }
    /**
     * @return the size of the HA-NAS in GB
     * 
     */
    public Double zpoolSize() {
        return this.zpoolSize;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNasHAResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String NasHAURN;
        private Boolean canCreatePartition;
        private String customName;
        private String datacenter;
        private String diskType;
        private String id;
        private String ip;
        private Boolean monitored;
        private List<String> partitionsLists;
        private String serviceName;
        private Double zpoolCapacity;
        private Double zpoolSize;
        public Builder() {}
        public Builder(GetNasHAResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.NasHAURN = defaults.NasHAURN;
    	      this.canCreatePartition = defaults.canCreatePartition;
    	      this.customName = defaults.customName;
    	      this.datacenter = defaults.datacenter;
    	      this.diskType = defaults.diskType;
    	      this.id = defaults.id;
    	      this.ip = defaults.ip;
    	      this.monitored = defaults.monitored;
    	      this.partitionsLists = defaults.partitionsLists;
    	      this.serviceName = defaults.serviceName;
    	      this.zpoolCapacity = defaults.zpoolCapacity;
    	      this.zpoolSize = defaults.zpoolSize;
        }

        @CustomType.Setter
        public Builder NasHAURN(String NasHAURN) {
            if (NasHAURN == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "NasHAURN");
            }
            this.NasHAURN = NasHAURN;
            return this;
        }
        @CustomType.Setter
        public Builder canCreatePartition(Boolean canCreatePartition) {
            if (canCreatePartition == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "canCreatePartition");
            }
            this.canCreatePartition = canCreatePartition;
            return this;
        }
        @CustomType.Setter
        public Builder customName(String customName) {
            if (customName == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "customName");
            }
            this.customName = customName;
            return this;
        }
        @CustomType.Setter
        public Builder datacenter(String datacenter) {
            if (datacenter == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "datacenter");
            }
            this.datacenter = datacenter;
            return this;
        }
        @CustomType.Setter
        public Builder diskType(String diskType) {
            if (diskType == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "diskType");
            }
            this.diskType = diskType;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ip(String ip) {
            if (ip == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "ip");
            }
            this.ip = ip;
            return this;
        }
        @CustomType.Setter
        public Builder monitored(Boolean monitored) {
            if (monitored == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "monitored");
            }
            this.monitored = monitored;
            return this;
        }
        @CustomType.Setter
        public Builder partitionsLists(List<String> partitionsLists) {
            if (partitionsLists == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "partitionsLists");
            }
            this.partitionsLists = partitionsLists;
            return this;
        }
        public Builder partitionsLists(String... partitionsLists) {
            return partitionsLists(List.of(partitionsLists));
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder zpoolCapacity(Double zpoolCapacity) {
            if (zpoolCapacity == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "zpoolCapacity");
            }
            this.zpoolCapacity = zpoolCapacity;
            return this;
        }
        @CustomType.Setter
        public Builder zpoolSize(Double zpoolSize) {
            if (zpoolSize == null) {
              throw new MissingRequiredPropertyException("GetNasHAResult", "zpoolSize");
            }
            this.zpoolSize = zpoolSize;
            return this;
        }
        public GetNasHAResult build() {
            final var _resultValue = new GetNasHAResult();
            _resultValue.NasHAURN = NasHAURN;
            _resultValue.canCreatePartition = canCreatePartition;
            _resultValue.customName = customName;
            _resultValue.datacenter = datacenter;
            _resultValue.diskType = diskType;
            _resultValue.id = id;
            _resultValue.ip = ip;
            _resultValue.monitored = monitored;
            _resultValue.partitionsLists = partitionsLists;
            _resultValue.serviceName = serviceName;
            _resultValue.zpoolCapacity = zpoolCapacity;
            _resultValue.zpoolSize = zpoolSize;
            return _resultValue;
        }
    }
}
