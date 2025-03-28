// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetKubeNodePoolTemplate;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetKubeNodePoolResult {
    /**
     * @return (Optional) should the pool use the anti-affinity feature. Default to `false`.
     * 
     */
    private Boolean antiAffinity;
    /**
     * @return (Optional) Enable auto-scaling for the pool. Default to `false`.
     * 
     */
    private Boolean autoscale;
    /**
     * @return (Optional) scaleDownUnneededTimeSeconds autoscaling parameter How long a node should be unneeded before it is eligible for scale down
     * 
     */
    private Integer autoscalingScaleDownUnneededTimeSeconds;
    /**
     * @return (Optional) scaleDownUnreadyTimeSeconds autoscaling parameter How long an unready node should be unneeded before it is eligible for scale down
     * 
     */
    private Integer autoscalingScaleDownUnreadyTimeSeconds;
    /**
     * @return (Optional) scaleDownUtilizationThreshold autoscaling parameter Node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
     * 
     */
    private Double autoscalingScaleDownUtilizationThreshold;
    /**
     * @return Number of nodes which are actually ready in the pool
     * 
     */
    private Integer availableNodes;
    /**
     * @return Creation date
     * 
     */
    private String createdAt;
    /**
     * @return Number of nodes present in the pool
     * 
     */
    private Integer currentNodes;
    /**
     * @return Number of nodes you desire in the pool
     * 
     */
    private Integer desiredNodes;
    /**
     * @return Flavor name
     * 
     */
    private String flavor;
    /**
     * @return a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: &#34;b2-7&#34;. Changing this value recreates the resource. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
     * 
     */
    private String flavorName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String kubeId;
    /**
     * @return maximum number of nodes allowed in the pool. Setting `desired_nodes` over this value will raise an error.
     * 
     */
    private Integer maxNodes;
    /**
     * @return minimum number of nodes allowed in the pool. Setting `desired_nodes` under this value will raise an error.
     * 
     */
    private Integer minNodes;
    /**
     * @return (Optional) should the nodes be billed on a monthly basis. Default to `false`.
     * 
     */
    private Boolean monthlyBilled;
    /**
     * @return (Optional) The name of the nodepool. Changing this value recreates the resource. Warning: &#34;_&#34; char is not allowed!
     * 
     */
    private String name;
    /**
     * @return Project id
     * 
     */
    private String projectId;
    /**
     * @return See Argument Reference above.
     * 
     */
    private String serviceName;
    /**
     * @return Status describing the state between number of nodes wanted and available ones
     * 
     */
    private String sizeStatus;
    /**
     * @return Current status
     * 
     */
    private String status;
    private @Nullable GetKubeNodePoolTemplate template;
    /**
     * @return Number of nodes with the latest version installed in the pool
     * 
     */
    private Integer upToDateNodes;
    /**
     * @return Last update date
     * 
     */
    private String updatedAt;

    private GetKubeNodePoolResult() {}
    /**
     * @return (Optional) should the pool use the anti-affinity feature. Default to `false`.
     * 
     */
    public Boolean antiAffinity() {
        return this.antiAffinity;
    }
    /**
     * @return (Optional) Enable auto-scaling for the pool. Default to `false`.
     * 
     */
    public Boolean autoscale() {
        return this.autoscale;
    }
    /**
     * @return (Optional) scaleDownUnneededTimeSeconds autoscaling parameter How long a node should be unneeded before it is eligible for scale down
     * 
     */
    public Integer autoscalingScaleDownUnneededTimeSeconds() {
        return this.autoscalingScaleDownUnneededTimeSeconds;
    }
    /**
     * @return (Optional) scaleDownUnreadyTimeSeconds autoscaling parameter How long an unready node should be unneeded before it is eligible for scale down
     * 
     */
    public Integer autoscalingScaleDownUnreadyTimeSeconds() {
        return this.autoscalingScaleDownUnreadyTimeSeconds;
    }
    /**
     * @return (Optional) scaleDownUtilizationThreshold autoscaling parameter Node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
     * 
     */
    public Double autoscalingScaleDownUtilizationThreshold() {
        return this.autoscalingScaleDownUtilizationThreshold;
    }
    /**
     * @return Number of nodes which are actually ready in the pool
     * 
     */
    public Integer availableNodes() {
        return this.availableNodes;
    }
    /**
     * @return Creation date
     * 
     */
    public String createdAt() {
        return this.createdAt;
    }
    /**
     * @return Number of nodes present in the pool
     * 
     */
    public Integer currentNodes() {
        return this.currentNodes;
    }
    /**
     * @return Number of nodes you desire in the pool
     * 
     */
    public Integer desiredNodes() {
        return this.desiredNodes;
    }
    /**
     * @return Flavor name
     * 
     */
    public String flavor() {
        return this.flavor;
    }
    /**
     * @return a valid OVHcloud public cloud flavor ID in which the nodes will be started. Ex: &#34;b2-7&#34;. Changing this value recreates the resource. You can find the list of flavor IDs: https://www.ovhcloud.com/fr/public-cloud/prices/
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
     * @return See Argument Reference above.
     * 
     */
    public String kubeId() {
        return this.kubeId;
    }
    /**
     * @return maximum number of nodes allowed in the pool. Setting `desired_nodes` over this value will raise an error.
     * 
     */
    public Integer maxNodes() {
        return this.maxNodes;
    }
    /**
     * @return minimum number of nodes allowed in the pool. Setting `desired_nodes` under this value will raise an error.
     * 
     */
    public Integer minNodes() {
        return this.minNodes;
    }
    /**
     * @return (Optional) should the nodes be billed on a monthly basis. Default to `false`.
     * 
     */
    public Boolean monthlyBilled() {
        return this.monthlyBilled;
    }
    /**
     * @return (Optional) The name of the nodepool. Changing this value recreates the resource. Warning: &#34;_&#34; char is not allowed!
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Project id
     * 
     */
    public String projectId() {
        return this.projectId;
    }
    /**
     * @return See Argument Reference above.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Status describing the state between number of nodes wanted and available ones
     * 
     */
    public String sizeStatus() {
        return this.sizeStatus;
    }
    /**
     * @return Current status
     * 
     */
    public String status() {
        return this.status;
    }
    public Optional<GetKubeNodePoolTemplate> template() {
        return Optional.ofNullable(this.template);
    }
    /**
     * @return Number of nodes with the latest version installed in the pool
     * 
     */
    public Integer upToDateNodes() {
        return this.upToDateNodes;
    }
    /**
     * @return Last update date
     * 
     */
    public String updatedAt() {
        return this.updatedAt;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetKubeNodePoolResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Boolean antiAffinity;
        private Boolean autoscale;
        private Integer autoscalingScaleDownUnneededTimeSeconds;
        private Integer autoscalingScaleDownUnreadyTimeSeconds;
        private Double autoscalingScaleDownUtilizationThreshold;
        private Integer availableNodes;
        private String createdAt;
        private Integer currentNodes;
        private Integer desiredNodes;
        private String flavor;
        private String flavorName;
        private String id;
        private String kubeId;
        private Integer maxNodes;
        private Integer minNodes;
        private Boolean monthlyBilled;
        private String name;
        private String projectId;
        private String serviceName;
        private String sizeStatus;
        private String status;
        private @Nullable GetKubeNodePoolTemplate template;
        private Integer upToDateNodes;
        private String updatedAt;
        public Builder() {}
        public Builder(GetKubeNodePoolResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.antiAffinity = defaults.antiAffinity;
    	      this.autoscale = defaults.autoscale;
    	      this.autoscalingScaleDownUnneededTimeSeconds = defaults.autoscalingScaleDownUnneededTimeSeconds;
    	      this.autoscalingScaleDownUnreadyTimeSeconds = defaults.autoscalingScaleDownUnreadyTimeSeconds;
    	      this.autoscalingScaleDownUtilizationThreshold = defaults.autoscalingScaleDownUtilizationThreshold;
    	      this.availableNodes = defaults.availableNodes;
    	      this.createdAt = defaults.createdAt;
    	      this.currentNodes = defaults.currentNodes;
    	      this.desiredNodes = defaults.desiredNodes;
    	      this.flavor = defaults.flavor;
    	      this.flavorName = defaults.flavorName;
    	      this.id = defaults.id;
    	      this.kubeId = defaults.kubeId;
    	      this.maxNodes = defaults.maxNodes;
    	      this.minNodes = defaults.minNodes;
    	      this.monthlyBilled = defaults.monthlyBilled;
    	      this.name = defaults.name;
    	      this.projectId = defaults.projectId;
    	      this.serviceName = defaults.serviceName;
    	      this.sizeStatus = defaults.sizeStatus;
    	      this.status = defaults.status;
    	      this.template = defaults.template;
    	      this.upToDateNodes = defaults.upToDateNodes;
    	      this.updatedAt = defaults.updatedAt;
        }

        @CustomType.Setter
        public Builder antiAffinity(Boolean antiAffinity) {
            if (antiAffinity == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "antiAffinity");
            }
            this.antiAffinity = antiAffinity;
            return this;
        }
        @CustomType.Setter
        public Builder autoscale(Boolean autoscale) {
            if (autoscale == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "autoscale");
            }
            this.autoscale = autoscale;
            return this;
        }
        @CustomType.Setter
        public Builder autoscalingScaleDownUnneededTimeSeconds(Integer autoscalingScaleDownUnneededTimeSeconds) {
            if (autoscalingScaleDownUnneededTimeSeconds == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "autoscalingScaleDownUnneededTimeSeconds");
            }
            this.autoscalingScaleDownUnneededTimeSeconds = autoscalingScaleDownUnneededTimeSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder autoscalingScaleDownUnreadyTimeSeconds(Integer autoscalingScaleDownUnreadyTimeSeconds) {
            if (autoscalingScaleDownUnreadyTimeSeconds == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "autoscalingScaleDownUnreadyTimeSeconds");
            }
            this.autoscalingScaleDownUnreadyTimeSeconds = autoscalingScaleDownUnreadyTimeSeconds;
            return this;
        }
        @CustomType.Setter
        public Builder autoscalingScaleDownUtilizationThreshold(Double autoscalingScaleDownUtilizationThreshold) {
            if (autoscalingScaleDownUtilizationThreshold == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "autoscalingScaleDownUtilizationThreshold");
            }
            this.autoscalingScaleDownUtilizationThreshold = autoscalingScaleDownUtilizationThreshold;
            return this;
        }
        @CustomType.Setter
        public Builder availableNodes(Integer availableNodes) {
            if (availableNodes == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "availableNodes");
            }
            this.availableNodes = availableNodes;
            return this;
        }
        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder currentNodes(Integer currentNodes) {
            if (currentNodes == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "currentNodes");
            }
            this.currentNodes = currentNodes;
            return this;
        }
        @CustomType.Setter
        public Builder desiredNodes(Integer desiredNodes) {
            if (desiredNodes == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "desiredNodes");
            }
            this.desiredNodes = desiredNodes;
            return this;
        }
        @CustomType.Setter
        public Builder flavor(String flavor) {
            if (flavor == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "flavor");
            }
            this.flavor = flavor;
            return this;
        }
        @CustomType.Setter
        public Builder flavorName(String flavorName) {
            if (flavorName == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "flavorName");
            }
            this.flavorName = flavorName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder kubeId(String kubeId) {
            if (kubeId == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "kubeId");
            }
            this.kubeId = kubeId;
            return this;
        }
        @CustomType.Setter
        public Builder maxNodes(Integer maxNodes) {
            if (maxNodes == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "maxNodes");
            }
            this.maxNodes = maxNodes;
            return this;
        }
        @CustomType.Setter
        public Builder minNodes(Integer minNodes) {
            if (minNodes == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "minNodes");
            }
            this.minNodes = minNodes;
            return this;
        }
        @CustomType.Setter
        public Builder monthlyBilled(Boolean monthlyBilled) {
            if (monthlyBilled == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "monthlyBilled");
            }
            this.monthlyBilled = monthlyBilled;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder projectId(String projectId) {
            if (projectId == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "projectId");
            }
            this.projectId = projectId;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder sizeStatus(String sizeStatus) {
            if (sizeStatus == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "sizeStatus");
            }
            this.sizeStatus = sizeStatus;
            return this;
        }
        @CustomType.Setter
        public Builder status(String status) {
            if (status == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "status");
            }
            this.status = status;
            return this;
        }
        @CustomType.Setter
        public Builder template(@Nullable GetKubeNodePoolTemplate template) {

            this.template = template;
            return this;
        }
        @CustomType.Setter
        public Builder upToDateNodes(Integer upToDateNodes) {
            if (upToDateNodes == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "upToDateNodes");
            }
            this.upToDateNodes = upToDateNodes;
            return this;
        }
        @CustomType.Setter
        public Builder updatedAt(String updatedAt) {
            if (updatedAt == null) {
              throw new MissingRequiredPropertyException("GetKubeNodePoolResult", "updatedAt");
            }
            this.updatedAt = updatedAt;
            return this;
        }
        public GetKubeNodePoolResult build() {
            final var _resultValue = new GetKubeNodePoolResult();
            _resultValue.antiAffinity = antiAffinity;
            _resultValue.autoscale = autoscale;
            _resultValue.autoscalingScaleDownUnneededTimeSeconds = autoscalingScaleDownUnneededTimeSeconds;
            _resultValue.autoscalingScaleDownUnreadyTimeSeconds = autoscalingScaleDownUnreadyTimeSeconds;
            _resultValue.autoscalingScaleDownUtilizationThreshold = autoscalingScaleDownUtilizationThreshold;
            _resultValue.availableNodes = availableNodes;
            _resultValue.createdAt = createdAt;
            _resultValue.currentNodes = currentNodes;
            _resultValue.desiredNodes = desiredNodes;
            _resultValue.flavor = flavor;
            _resultValue.flavorName = flavorName;
            _resultValue.id = id;
            _resultValue.kubeId = kubeId;
            _resultValue.maxNodes = maxNodes;
            _resultValue.minNodes = minNodes;
            _resultValue.monthlyBilled = monthlyBilled;
            _resultValue.name = name;
            _resultValue.projectId = projectId;
            _resultValue.serviceName = serviceName;
            _resultValue.sizeStatus = sizeStatus;
            _resultValue.status = status;
            _resultValue.template = template;
            _resultValue.upToDateNodes = upToDateNodes;
            _resultValue.updatedAt = updatedAt;
            return _resultValue;
        }
    }
}
