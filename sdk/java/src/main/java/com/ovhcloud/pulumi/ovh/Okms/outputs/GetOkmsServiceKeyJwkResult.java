// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.outputs;

import com.ovhcloud.pulumi.ovh.Okms.outputs.GetOkmsServiceKeyJwkKey;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetOkmsServiceKeyJwkResult {
    private String createdAt;
    private String id;
    private List<GetOkmsServiceKeyJwkKey> keys;
    private String name;
    private String okmsId;
    private Double size;
    private String state;
    private String type;

    private GetOkmsServiceKeyJwkResult() {}
    public String createdAt() {
        return this.createdAt;
    }
    public String id() {
        return this.id;
    }
    public List<GetOkmsServiceKeyJwkKey> keys() {
        return this.keys;
    }
    public String name() {
        return this.name;
    }
    public String okmsId() {
        return this.okmsId;
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

    public static Builder builder(GetOkmsServiceKeyJwkResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String createdAt;
        private String id;
        private List<GetOkmsServiceKeyJwkKey> keys;
        private String name;
        private String okmsId;
        private Double size;
        private String state;
        private String type;
        public Builder() {}
        public Builder(GetOkmsServiceKeyJwkResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.createdAt = defaults.createdAt;
    	      this.id = defaults.id;
    	      this.keys = defaults.keys;
    	      this.name = defaults.name;
    	      this.okmsId = defaults.okmsId;
    	      this.size = defaults.size;
    	      this.state = defaults.state;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder createdAt(String createdAt) {
            if (createdAt == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "createdAt");
            }
            this.createdAt = createdAt;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder keys(List<GetOkmsServiceKeyJwkKey> keys) {
            if (keys == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "keys");
            }
            this.keys = keys;
            return this;
        }
        public Builder keys(GetOkmsServiceKeyJwkKey... keys) {
            return keys(List.of(keys));
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder okmsId(String okmsId) {
            if (okmsId == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "okmsId");
            }
            this.okmsId = okmsId;
            return this;
        }
        @CustomType.Setter
        public Builder size(Double size) {
            if (size == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "size");
            }
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder state(String state) {
            if (state == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "state");
            }
            this.state = state;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkResult", "type");
            }
            this.type = type;
            return this;
        }
        public GetOkmsServiceKeyJwkResult build() {
            final var _resultValue = new GetOkmsServiceKeyJwkResult();
            _resultValue.createdAt = createdAt;
            _resultValue.id = id;
            _resultValue.keys = keys;
            _resultValue.name = name;
            _resultValue.okmsId = okmsId;
            _resultValue.size = size;
            _resultValue.state = state;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
