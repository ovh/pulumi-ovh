// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs;

import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs.GetDatabaseEndpoint;
import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs.GetDatabaseIpRestriction;
import com.ovhcloud.pulumi.ovh.CloudProjectDatabase.outputs.GetDatabaseNode;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetDatabaseResult {
    /**
     * @return Advanced configuration key / value.
     * 
     */
    private Map<String,String> advancedConfiguration;
    /**
     * @return List of region where backups are pushed.
     * 
     */
    private List<String> backupRegions;
    /**
     * @return Time on which backups start every day.
     * 
     */
    private String backupTime;
    /**
     * @return Date of the creation of the cluster.
     * 
     */
    private String createdAt;
    /**
     * @return Description of the IP restriction
     * 
     */
    private String description;
    /**
     * @return The disk size (in GB) of the database service.
     * 
     */
    private Integer diskSize;
    /**
     * @return The disk type of the database service.
     * 
     */
    private String diskType;
    /**
     * @return List of all endpoints objects of the service.
     * 
     */
    private List<GetDatabaseEndpoint> endpoints;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String engine;
    /**
     * @return A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     * 
     */
    private String flavor;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String id;
    /**
     * @return IP Blocks authorized to access to the cluster.
     * 
     */
    private List<GetDatabaseIpRestriction> ipRestrictions;
    /**
     * @return Defines whether the REST API is enabled on a kafka cluster.
     * 
     */
    private Boolean kafkaRestApi;
    /**
     * @return Defines whether the schema registry is enabled on a Kafka cluster
     * 
     */
    private Boolean kafkaSchemaRegistry;
    /**
     * @return Time on which maintenances can start every day.
     * 
     */
    private String maintenanceTime;
    /**
     * @return Type of network of the cluster.
     * 
     */
    private String networkType;
    /**
     * @return List of nodes object.
     * 
     */
    private List<GetDatabaseNode> nodes;
    private Boolean opensearchAclsEnabled;
    /**
     * @return Plan of the cluster.
     * 
     */
    private String plan;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return Current status of the cluster.
     * 
     */
    private String status;
    /**
     * @return The version of the engine in which the service should be deployed
     * 
     */
    private String version;

    private GetDatabaseResult() {}
    /**
     * @return Advanced configuration key / value.
     * 
     */
    public Map<String,String> advancedConfiguration() {
        return this.advancedConfiguration;
    }
    /**
     * @return List of region where backups are pushed.
     * 
     */
    public List<String> backupRegions() {
        return this.backupRegions;
    }
    /**
     * @return Time on which backups start every day.
     * 
     */
    public String backupTime() {
        return this.backupTime;
    }
    /**
     * @return Date of the creation of the cluster.
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Description of the IP restriction
     * 
     */
    public String description() {
        return this.description;
    }
    /**
     * @return The disk size (in GB) of the database service.
     * 
     */
    public Integer diskSize() {
        return this.diskSize;
    }
    /**
     * @return The disk type of the database service.
     * 
     */
    public String diskType() {
        return this.diskType;
    }
    /**
     * @return List of all endpoints objects of the service.
     * 
     */
    public List<GetDatabaseEndpoint> endpoints() {
        return this.endpoints;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String engine() {
        return this.engine;
    }
    /**
     * @return A valid OVHcloud public cloud database flavor name in which the nodes will be started.
     * 
     */
    public String flavor() {
        return this.flavor;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return IP Blocks authorized to access to the cluster.
     * 
     */
    public List<GetDatabaseIpRestriction> ipRestrictions() {
        return this.ipRestrictions;
    }
    /**
     * @return Defines whether the REST API is enabled on a kafka cluster.
     * 
     */
    public Boolean kafkaRestApi() {
        return this.kafkaRestApi;
    }
    /**
     * @return Defines whether the schema registry is enabled on a Kafka cluster
     * 
     */
    public Boolean kafkaSchemaRegistry() {
        return this.kafkaSchemaRegistry;
    }
    /**
     * @return Time on which maintenances can start every day.
     * 
     */
    public String maintenanceTime() {
        return this.maintenanceTime;
    }
    /**
     * @return Type of network of the cluster.
     * 
     */
    public String networkType() {
        return this.networkType;
    }
    /**
     * @return List of nodes object.
     * 
     */
    public List<GetDatabaseNode> nodes() {
        return this.nodes;
    }
    public Boolean opensearchAclsEnabled() {
        return this.opensearchAclsEnabled;
    }
    /**
     * @return Plan of the cluster.
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Current status of the cluster.
     * 
     */
    public String status() {
        return this.status;
    }
    /**
     * @return The version of the engine in which the service should be deployed
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetDatabaseResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> advancedConfiguration;
        private List<String> backupRegions;
        private String backupTime;
        private String createdAt;
        private String description;
        private Integer diskSize;
        private String diskType;
        private List<GetDatabaseEndpoint> endpoints;
        private String engine;
        private String flavor;
        private String id;
        private List<GetDatabaseIpRestriction> ipRestrictions;
        private Boolean kafkaRestApi;
        private Boolean kafkaSchemaRegistry;
        private String maintenanceTime;
        private String networkType;
        private List<GetDatabaseNode> nodes;
        private Boolean opensearchAclsEnabled;
        private String plan;
        private String serviceName;
        private String status;
        private String version;
        public Builder() {}
        public Builder(GetDatabaseResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.advancedConfiguration = defaults.advancedConfiguration;
    	      this.backupRegions = defaults.backupRegions;
    	      this.backupTime = defaults.backupTime;
    	      this.createdAt = defaults.createdAt;
    	      this.description = defaults.description;
    	      this.diskSize = defaults.diskSize;
    	      this.diskType = defaults.diskType;
    	      this.endpoints = defaults.endpoints;
    	      this.engine = defaults.engine;
    	      this.flavor = defaults.flavor;
    	      this.id = defaults.id;
    	      this.ipRestrictions = defaults.ipRestrictions;
    	      this.kafkaRestApi = defaults.kafkaRestApi;
    	      this.kafkaSchemaRegistry = defaults.kafkaSchemaRegistry;
    	      this.maintenanceTime = defaults.maintenanceTime;
    	      this.networkType = defaults.networkType;
    	      this.nodes = defaults.nodes;
    	      this.opensearchAclsEnabled = defaults.opensearchAclsEnabled;
    	      this.plan = defaults.plan;
    	      this.serviceName = defaults.serviceName;
    	      this.status = defaults.status;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder advancedConfiguration(Map<String,String> advancedConfiguration) {
            if (advancedConfiguration == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "advancedConfiguration");
            }
            this.advancedConfiguration = advancedConfiguration;
            return this;
        }
        @CustomType.Setter
        public Builder backupRegions(List<String> backupRegions) {
            if (backupRegions == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "backupRegions");
            }
            this.backupRegions = backupRegions;
            return this;
        }
        public Builder backupRegions(String... backupRegions) {
            return backupRegions(List.of(backupRegions));
        }
        @CustomType.Setter
        public Builder backupTime(String backupTime) {
            if (backupTime == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "backupTime");
            }
            this.backupTime = backupTime;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder diskSize(Integer diskSize) {
            if (diskSize == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "diskSize");
            }
            this.diskSize = diskSize;
            return this;
        }
        @CustomType.Setter
        public Builder diskType(String diskType) {
            if (diskType == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "diskType");
            }
            this.diskType = diskType;
            return this;
        }
        @CustomType.Setter
        public Builder endpoints(List<GetDatabaseEndpoint> endpoints) {
            if (endpoints == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "endpoints");
            }
            this.endpoints = endpoints;
            return this;
        }
        public Builder endpoints(GetDatabaseEndpoint... endpoints) {
            return endpoints(List.of(endpoints));
        }
        @CustomType.Setter
        public Builder engine(String engine) {
            if (engine == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "engine");
            }
            this.engine = engine;
            return this;
        }
        @CustomType.Setter
        public Builder flavor(String flavor) {
            if (flavor == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "flavor");
            }
            this.flavor = flavor;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder ipRestrictions(List<GetDatabaseIpRestriction> ipRestrictions) {
            if (ipRestrictions == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "ipRestrictions");
            }
            this.ipRestrictions = ipRestrictions;
            return this;
        }
        public Builder ipRestrictions(GetDatabaseIpRestriction... ipRestrictions) {
            return ipRestrictions(List.of(ipRestrictions));
        }
        @CustomType.Setter
        public Builder kafkaRestApi(Boolean kafkaRestApi) {
            if (kafkaRestApi == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "kafkaRestApi");
            }
            this.kafkaRestApi = kafkaRestApi;
            return this;
        }
        @CustomType.Setter
        public Builder kafkaSchemaRegistry(Boolean kafkaSchemaRegistry) {
            if (kafkaSchemaRegistry == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "kafkaSchemaRegistry");
            }
            this.kafkaSchemaRegistry = kafkaSchemaRegistry;
            return this;
        }
        @CustomType.Setter
        public Builder maintenanceTime(String maintenanceTime) {
            if (maintenanceTime == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "maintenanceTime");
            }
            this.maintenanceTime = maintenanceTime;
            return this;
        }
        @CustomType.Setter
        public Builder networkType(String networkType) {
            if (networkType == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "networkType");
            }
            this.networkType = networkType;
            return this;
        }
        @CustomType.Setter
        public Builder nodes(List<GetDatabaseNode> nodes) {
            if (nodes == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "nodes");
            }
            this.nodes = nodes;
            return this;
        }
        public Builder nodes(GetDatabaseNode... nodes) {
            return nodes(List.of(nodes));
        }
        @CustomType.Setter
        public Builder opensearchAclsEnabled(Boolean opensearchAclsEnabled) {
            if (opensearchAclsEnabled == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "opensearchAclsEnabled");
            }
            this.opensearchAclsEnabled = opensearchAclsEnabled;
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            if (plan == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "plan");
            }
            this.plan = plan;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("GetDatabaseResult", "version");
            }
            this.version = version;
            return this;
        }
        public GetDatabaseResult build() {
            final var _resultValue = new GetDatabaseResult();
            _resultValue.advancedConfiguration = advancedConfiguration;
            _resultValue.backupRegions = backupRegions;
            _resultValue.backupTime = backupTime;
            _resultValue.createdAt = createdAt;
            _resultValue.description = description;
            _resultValue.diskSize = diskSize;
            _resultValue.diskType = diskType;
            _resultValue.endpoints = endpoints;
            _resultValue.engine = engine;
            _resultValue.flavor = flavor;
            _resultValue.id = id;
            _resultValue.ipRestrictions = ipRestrictions;
            _resultValue.kafkaRestApi = kafkaRestApi;
            _resultValue.kafkaSchemaRegistry = kafkaSchemaRegistry;
            _resultValue.maintenanceTime = maintenanceTime;
            _resultValue.networkType = networkType;
            _resultValue.nodes = nodes;
            _resultValue.opensearchAclsEnabled = opensearchAclsEnabled;
            _resultValue.plan = plan;
            _resultValue.serviceName = serviceName;
            _resultValue.status = status;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
