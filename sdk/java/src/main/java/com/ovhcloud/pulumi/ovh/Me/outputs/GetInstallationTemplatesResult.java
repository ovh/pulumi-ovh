// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetInstallationTemplatesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return The list of custom installation templates IDs available for dedicated servers.
     * 
     */
    private List<String> results;

    private GetInstallationTemplatesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return The list of custom installation templates IDs available for dedicated servers.
     * 
     */
    public List<String> results() {
        return this.results;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetInstallationTemplatesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private List<String> results;
        public Builder() {}
        public Builder(GetInstallationTemplatesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.results = defaults.results;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplatesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder results(List<String> results) {
            if (results == null) {
              throw new MissingRequiredPropertyException("GetInstallationTemplatesResult", "results");
            }
            this.results = results;
            return this;
        }
        public Builder results(String... results) {
            return results(List.of(results));
        }
        public GetInstallationTemplatesResult build() {
            final var _resultValue = new GetInstallationTemplatesResult();
            _resultValue.id = id;
            _resultValue.results = results;
            return _resultValue;
        }
    }
}
