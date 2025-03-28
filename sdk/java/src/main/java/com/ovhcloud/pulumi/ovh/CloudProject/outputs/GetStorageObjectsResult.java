// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.ovhcloud.pulumi.ovh.CloudProject.outputs.GetStorageObjectsObject;
import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetStorageObjectsResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return Key to start with when listing objects
     * 
     */
    private @Nullable String keyMarker;
    /**
     * @return Limit the number of objects returned (1000 maximum, defaults to 1000)
     * 
     */
    private @Nullable Double limit;
    /**
     * @return Name
     * 
     */
    private String name;
    private List<GetStorageObjectsObject> objects;
    /**
     * @return List objects whose key begins with this prefix
     * 
     */
    private @Nullable String prefix;
    /**
     * @return Region name
     * 
     */
    private String regionName;
    /**
     * @return Service name
     * 
     */
    private String serviceName;
    /**
     * @return Version ID to start listing from
     * 
     */
    private @Nullable String versionIdMarker;
    /**
     * @return List object versions
     * 
     */
    private @Nullable Boolean withVersions;

    private GetStorageObjectsResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Key to start with when listing objects
     * 
     */
    public Optional<String> keyMarker() {
        return Optional.ofNullable(this.keyMarker);
    }
    /**
     * @return Limit the number of objects returned (1000 maximum, defaults to 1000)
     * 
     */
    public Optional<Double> limit() {
        return Optional.ofNullable(this.limit);
    }
    /**
     * @return Name
     * 
     */
    public String name() {
        return this.name;
    }
    public List<GetStorageObjectsObject> objects() {
        return this.objects;
    }
    /**
     * @return List objects whose key begins with this prefix
     * 
     */
    public Optional<String> prefix() {
        return Optional.ofNullable(this.prefix);
    }
    /**
     * @return Region name
     * 
     */
    public String regionName() {
        return this.regionName;
    }
    /**
     * @return Service name
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }
    /**
     * @return Version ID to start listing from
     * 
     */
    public Optional<String> versionIdMarker() {
        return Optional.ofNullable(this.versionIdMarker);
    }
    /**
     * @return List object versions
     * 
     */
    public Optional<Boolean> withVersions() {
        return Optional.ofNullable(this.withVersions);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStorageObjectsResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private @Nullable String keyMarker;
        private @Nullable Double limit;
        private String name;
        private List<GetStorageObjectsObject> objects;
        private @Nullable String prefix;
        private String regionName;
        private String serviceName;
        private @Nullable String versionIdMarker;
        private @Nullable Boolean withVersions;
        public Builder() {}
        public Builder(GetStorageObjectsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.keyMarker = defaults.keyMarker;
    	      this.limit = defaults.limit;
    	      this.name = defaults.name;
    	      this.objects = defaults.objects;
    	      this.prefix = defaults.prefix;
    	      this.regionName = defaults.regionName;
    	      this.serviceName = defaults.serviceName;
    	      this.versionIdMarker = defaults.versionIdMarker;
    	      this.withVersions = defaults.withVersions;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder keyMarker(@Nullable String keyMarker) {

            this.keyMarker = keyMarker;
            return this;
        }
        @CustomType.Setter
        public Builder limit(@Nullable Double limit) {

            this.limit = limit;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder objects(List<GetStorageObjectsObject> objects) {
            if (objects == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsResult", "objects");
            }
            this.objects = objects;
            return this;
        }
        public Builder objects(GetStorageObjectsObject... objects) {
            return objects(List.of(objects));
        }
        @CustomType.Setter
        public Builder prefix(@Nullable String prefix) {

            this.prefix = prefix;
            return this;
        }
        @CustomType.Setter
        public Builder regionName(String regionName) {
            if (regionName == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsResult", "regionName");
            }
            this.regionName = regionName;
            return this;
        }
        @CustomType.Setter
        public Builder serviceName(String serviceName) {
            if (serviceName == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsResult", "serviceName");
            }
            this.serviceName = serviceName;
            return this;
        }
        @CustomType.Setter
        public Builder versionIdMarker(@Nullable String versionIdMarker) {

            this.versionIdMarker = versionIdMarker;
            return this;
        }
        @CustomType.Setter
        public Builder withVersions(@Nullable Boolean withVersions) {

            this.withVersions = withVersions;
            return this;
        }
        public GetStorageObjectsResult build() {
            final var _resultValue = new GetStorageObjectsResult();
            _resultValue.id = id;
            _resultValue.keyMarker = keyMarker;
            _resultValue.limit = limit;
            _resultValue.name = name;
            _resultValue.objects = objects;
            _resultValue.prefix = prefix;
            _resultValue.regionName = regionName;
            _resultValue.serviceName = serviceName;
            _resultValue.versionIdMarker = versionIdMarker;
            _resultValue.withVersions = withVersions;
            return _resultValue;
        }
    }
}
