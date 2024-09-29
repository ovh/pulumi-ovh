// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOkmsServiceKeyResult {
    private String createdAt;
    private String curve;
    private String id;
    private String name;
    private String okmsId;
    private List<String> operations;
    private Double size;
    private String state;
    private String type;

    private GetOkmsServiceKeyResult() {}
    public String createdAt() {
        return this.createdAt;
    }
    public String curve() {
        return this.curve;
    }
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public String okmsId() {
        return this.okmsId;
    }
    public List<String> operations() {
        return this.operations;
    }
    public Double size() {
        return this.size;
    }
    public String state() {
        return this.state;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetOkmsServiceKeyResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String curve;
        private String id;
        private String name;
        private String okmsId;
        private List<String> operations;
        private Double size;
        private String state;
        private String type;
        public Builder() {}
        public Builder(GetOkmsServiceKeyResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.curve = defaults.curve;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.okmsId = defaults.okmsId;
    	      this.operations = defaults.operations;
    	      this.size = defaults.size;
    	      this.state = defaults.state;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder curve(String curve) {
            if (curve == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "curve");
            }
            this.curve = curve;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder okmsId(String okmsId) {
            if (okmsId == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "okmsId");
            }
            this.okmsId = okmsId;
            return this;
        }
        @CustomType.Setter
        public Builder operations(List<String> operations) {
            if (operations == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "operations");
            }
            this.operations = operations;
            return this;
        }
        public Builder operations(String... operations) {
            return operations(List.of(operations));
        }
        @CustomType.Setter
        public Builder size(Double size) {
            if (size == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "size");
            }
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetOkmsServiceKeyResult build() {
            final var _resultValue = new GetOkmsServiceKeyResult();
            _resultValue.createdAt = createdAt;
            _resultValue.curve = curve;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.okmsId = okmsId;
            _resultValue.operations = operations;
            _resultValue.size = size;
            _resultValue.state = state;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}