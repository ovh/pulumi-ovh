// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetRancherCurrentStateIpRestriction;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetRancherCurrentStateNetworking;
import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetRancherCurrentStateUsage;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRancherCurrentState {
    /**
     * @return Bootstrap password of the managed Rancher service, returned only on creation
     * 
     */
    private String bootstrapPassword;
    /**
     * @return List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
     * 
     */
    private List<GetRancherCurrentStateIpRestriction> ipRestrictions;
    /**
     * @return Name of the managed Rancher service
     * 
     */
    private String name;
    /**
     * @return Networking properties of a managed Rancher service
     * 
     */
    private GetRancherCurrentStateNetworking networking;
    /**
     * @return Plan of the managed Rancher service
     * 
     */
    private String plan;
    /**
     * @return Region of the managed Rancher service
     * 
     */
    private String region;
    /**
     * @return URL of the managed Rancher service
     * 
     */
    private String url;
    /**
     * @return Latest metrics regarding the usage of the managed Rancher service
     * 
     */
    private GetRancherCurrentStateUsage usage;
    /**
     * @return Version of the managed Rancher service
     * 
     */
    private String version;

    private GetRancherCurrentState() {}
    /**
     * @return Bootstrap password of the managed Rancher service, returned only on creation
     * 
     */
    public String bootstrapPassword() {
        return this.bootstrapPassword;
    }
    /**
     * @return List of allowed CIDR blocks for a managed Rancher service&#39;s IP restrictions. When empty, any IP is allowed
     * 
     */
    public List<GetRancherCurrentStateIpRestriction> ipRestrictions() {
        return this.ipRestrictions;
    }
    /**
     * @return Name of the managed Rancher service
     * 
     */
    public String name() {
        return this.name;
    }
    /**
     * @return Networking properties of a managed Rancher service
     * 
     */
    public GetRancherCurrentStateNetworking networking() {
        return this.networking;
    }
    /**
     * @return Plan of the managed Rancher service
     * 
     */
    public String plan() {
        return this.plan;
    }
    /**
     * @return Region of the managed Rancher service
     * 
     */
    public String region() {
        return this.region;
    }
    /**
     * @return URL of the managed Rancher service
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return Latest metrics regarding the usage of the managed Rancher service
     * 
     */
    public GetRancherCurrentStateUsage usage() {
        return this.usage;
    }
    /**
     * @return Version of the managed Rancher service
     * 
     */
    public String version() {
        return this.version;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRancherCurrentState defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String bootstrapPassword;
        private List<GetRancherCurrentStateIpRestriction> ipRestrictions;
        private String name;
        private GetRancherCurrentStateNetworking networking;
        private String plan;
        private String region;
        private String url;
        private GetRancherCurrentStateUsage usage;
        private String version;
        public Builder() {}
        public Builder(GetRancherCurrentState defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bootstrapPassword = defaults.bootstrapPassword;
    	      this.ipRestrictions = defaults.ipRestrictions;
    	      this.name = defaults.name;
    	      this.networking = defaults.networking;
    	      this.plan = defaults.plan;
    	      this.region = defaults.region;
    	      this.url = defaults.url;
    	      this.usage = defaults.usage;
    	      this.version = defaults.version;
        }

        @CustomType.Setter
        public Builder bootstrapPassword(String bootstrapPassword) {
            if (bootstrapPassword == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "bootstrapPassword");
            }
            this.bootstrapPassword = bootstrapPassword;
            return this;
        }
        @CustomType.Setter
        public Builder ipRestrictions(List<GetRancherCurrentStateIpRestriction> ipRestrictions) {
            if (ipRestrictions == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "ipRestrictions");
            }
            this.ipRestrictions = ipRestrictions;
            return this;
        }
        public Builder ipRestrictions(GetRancherCurrentStateIpRestriction... ipRestrictions) {
            return ipRestrictions(List.of(ipRestrictions));
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder networking(GetRancherCurrentStateNetworking networking) {
            if (networking == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "networking");
            }
            this.networking = networking;
            return this;
        }
        @CustomType.Setter
        public Builder plan(String plan) {
            if (plan == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "plan");
            }
            this.plan = plan;
            return this;
        }
        @CustomType.Setter
        public Builder region(String region) {
            if (region == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "region");
            }
            this.region = region;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            if (url == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "url");
            }
            this.url = url;
            return this;
        }
        @CustomType.Setter
        public Builder usage(GetRancherCurrentStateUsage usage) {
            if (usage == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "usage");
            }
            this.usage = usage;
            return this;
        }
        @CustomType.Setter
        public Builder version(String version) {
            if (version == null) {
              throw new MissingRequiredPropertyException("GetRancherCurrentState", "version");
            }
            this.version = version;
            return this;
        }
        public GetRancherCurrentState build() {
            final var _resultValue = new GetRancherCurrentState();
            _resultValue.bootstrapPassword = bootstrapPassword;
            _resultValue.ipRestrictions = ipRestrictions;
            _resultValue.name = name;
            _resultValue.networking = networking;
            _resultValue.plan = plan;
            _resultValue.region = region;
            _resultValue.url = url;
            _resultValue.usage = usage;
            _resultValue.version = version;
            return _resultValue;
        }
    }
}
