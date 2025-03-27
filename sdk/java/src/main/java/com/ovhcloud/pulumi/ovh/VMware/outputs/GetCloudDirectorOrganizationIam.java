// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.VMware.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetCloudDirectorOrganizationIam {
    /**
     * @return Resource display name
     * 
     */
    private String displayName;
    /**
     * @return Unique identifier of the resource
     * 
     */
    private String id;
    /**
     * @return Resource tags. Tags that were internally computed are prefixed with ovh:
     * 
     */
    private Map<String,String> tags;
    /**
     * @return Unique resource name used in policies
     * 
     */
    private String urn;

    private GetCloudDirectorOrganizationIam() {}
    /**
     * @return Resource display name
     * 
     */
    public String displayName() {
        return this.displayName;
    }
    /**
     * @return Unique identifier of the resource
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Resource tags. Tags that were internally computed are prefixed with ovh:
     * 
     */
    public Map<String,String> tags() {
        return this.tags;
    }
    /**
     * @return Unique resource name used in policies
     * 
     */
    public String urn() {
        return this.urn;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetCloudDirectorOrganizationIam defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String displayName;
        private String id;
        private Map<String,String> tags;
        private String urn;
        public Builder() {}
        public Builder(GetCloudDirectorOrganizationIam defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.displayName = defaults.displayName;
    	      this.id = defaults.id;
    	      this.tags = defaults.tags;
    	      this.urn = defaults.urn;
        }

        @CustomType.Setter
        public Builder displayName(String displayName) {
            if (displayName == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorOrganizationIam", "displayName");
            }
            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorOrganizationIam", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder tags(Map<String,String> tags) {
            if (tags == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorOrganizationIam", "tags");
            }
            this.tags = tags;
            return this;
        }
        @CustomType.Setter
        public Builder urn(String urn) {
            if (urn == null) {
              throw new MissingRequiredPropertyException("GetCloudDirectorOrganizationIam", "urn");
            }
            this.urn = urn;
            return this;
        }
        public GetCloudDirectorOrganizationIam build() {
            final var _resultValue = new GetCloudDirectorOrganizationIam();
            _resultValue.displayName = displayName;
            _resultValue.id = id;
            _resultValue.tags = tags;
            _resultValue.urn = urn;
            return _resultValue;
        }
    }
}
