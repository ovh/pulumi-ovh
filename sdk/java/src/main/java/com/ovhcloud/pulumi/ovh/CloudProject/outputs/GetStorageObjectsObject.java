// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetStorageObjectsObject {
    /**
     * @return ETag
     * 
     */
    private String etag;
    /**
     * @return Whether this object is a delete marker
     * 
     */
    private Boolean isDeleteMarker;
    /**
     * @return Whether this is the latest version of the object
     * 
     */
    private Boolean isLatest;
    /**
     * @return Key
     * 
     */
    private String key;
    /**
     * @return Last modification date
     * 
     */
    private String lastModified;
    /**
     * @return Size (bytes)
     * 
     */
    private Double size;
    /**
     * @return Storage class
     * 
     */
    private String storageClass;
    /**
     * @return Version ID of the object
     * 
     */
    private String versionId;

    private GetStorageObjectsObject() {}
    /**
     * @return ETag
     * 
     */
    public String etag() {
        return this.etag;
    }
    /**
     * @return Whether this object is a delete marker
     * 
     */
    public Boolean isDeleteMarker() {
        return this.isDeleteMarker;
    }
    /**
     * @return Whether this is the latest version of the object
     * 
     */
    public Boolean isLatest() {
        return this.isLatest;
    }
    /**
     * @return Key
     * 
     */
    public String key() {
        return this.key;
    }
    /**
     * @return Last modification date
     * 
     */
    public String lastModified() {
        return this.lastModified;
    }
    /**
     * @return Size (bytes)
     * 
     */
    public Double size() {
        return this.size;
    }
    /**
     * @return Storage class
     * 
     */
    public String storageClass() {
        return this.storageClass;
    }
    /**
     * @return Version ID of the object
     * 
     */
    public String versionId() {
        return this.versionId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetStorageObjectsObject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String etag;
        private Boolean isDeleteMarker;
        private Boolean isLatest;
        private String key;
        private String lastModified;
        private Double size;
        private String storageClass;
        private String versionId;
        public Builder() {}
        public Builder(GetStorageObjectsObject defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.etag = defaults.etag;
    	      this.isDeleteMarker = defaults.isDeleteMarker;
    	      this.isLatest = defaults.isLatest;
    	      this.key = defaults.key;
    	      this.lastModified = defaults.lastModified;
    	      this.size = defaults.size;
    	      this.storageClass = defaults.storageClass;
    	      this.versionId = defaults.versionId;
        }

        @CustomType.Setter
        public Builder etag(String etag) {
            if (etag == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "etag");
            }
            this.etag = etag;
            return this;
        }
        @CustomType.Setter
        public Builder isDeleteMarker(Boolean isDeleteMarker) {
            if (isDeleteMarker == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "isDeleteMarker");
            }
            this.isDeleteMarker = isDeleteMarker;
            return this;
        }
        @CustomType.Setter
        public Builder isLatest(Boolean isLatest) {
            if (isLatest == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "isLatest");
            }
            this.isLatest = isLatest;
            return this;
        }
        @CustomType.Setter
        public Builder key(String key) {
            if (key == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "key");
            }
            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder lastModified(String lastModified) {
            if (lastModified == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "lastModified");
            }
            this.lastModified = lastModified;
            return this;
        }
        @CustomType.Setter
        public Builder size(Double size) {
            if (size == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "size");
            }
            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder storageClass(String storageClass) {
            if (storageClass == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "storageClass");
            }
            this.storageClass = storageClass;
            return this;
        }
        @CustomType.Setter
        public Builder versionId(String versionId) {
            if (versionId == null) {
              throw new MissingRequiredPropertyException("GetStorageObjectsObject", "versionId");
            }
            this.versionId = versionId;
            return this;
        }
        public GetStorageObjectsObject build() {
            final var _resultValue = new GetStorageObjectsObject();
            _resultValue.etag = etag;
            _resultValue.isDeleteMarker = isDeleteMarker;
            _resultValue.isLatest = isLatest;
            _resultValue.key = key;
            _resultValue.lastModified = lastModified;
            _resultValue.size = size;
            _resultValue.storageClass = storageClass;
            _resultValue.versionId = versionId;
            return _resultValue;
        }
    }
}
