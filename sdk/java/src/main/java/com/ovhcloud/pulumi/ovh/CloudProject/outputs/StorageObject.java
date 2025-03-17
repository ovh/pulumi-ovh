// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class StorageObject {
    /**
     * @return ETag
     * 
     */
    private @Nullable String etag;
    /**
     * @return Whether this object is a delete marker
     * 
     */
    private @Nullable Boolean isDeleteMarker;
    /**
     * @return Whether this is the latest version of the object
     * 
     */
    private @Nullable Boolean isLatest;
    /**
     * @return Key
     * 
     */
    private @Nullable String key;
    /**
     * @return Last modification date
     * 
     */
    private @Nullable String lastModified;
    /**
     * @return Size (bytes)
     * 
     */
    private @Nullable Double size;
    /**
     * @return Storage class
     * 
     */
    private @Nullable String storageClass;
    /**
     * @return Version ID of the object
     * 
     */
    private @Nullable String versionId;

    private StorageObject() {}
    /**
     * @return ETag
     * 
     */
    public Optional<String> etag() {
        return Optional.ofNullable(this.etag);
    }
    /**
     * @return Whether this object is a delete marker
     * 
     */
    public Optional<Boolean> isDeleteMarker() {
        return Optional.ofNullable(this.isDeleteMarker);
    }
    /**
     * @return Whether this is the latest version of the object
     * 
     */
    public Optional<Boolean> isLatest() {
        return Optional.ofNullable(this.isLatest);
    }
    /**
     * @return Key
     * 
     */
    public Optional<String> key() {
        return Optional.ofNullable(this.key);
    }
    /**
     * @return Last modification date
     * 
     */
    public Optional<String> lastModified() {
        return Optional.ofNullable(this.lastModified);
    }
    /**
     * @return Size (bytes)
     * 
     */
    public Optional<Double> size() {
        return Optional.ofNullable(this.size);
    }
    /**
     * @return Storage class
     * 
     */
    public Optional<String> storageClass() {
        return Optional.ofNullable(this.storageClass);
    }
    /**
     * @return Version ID of the object
     * 
     */
    public Optional<String> versionId() {
        return Optional.ofNullable(this.versionId);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(StorageObject defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String etag;
        private @Nullable Boolean isDeleteMarker;
        private @Nullable Boolean isLatest;
        private @Nullable String key;
        private @Nullable String lastModified;
        private @Nullable Double size;
        private @Nullable String storageClass;
        private @Nullable String versionId;
        public Builder() {}
        public Builder(StorageObject defaults) {
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
        public Builder etag(@Nullable String etag) {

            this.etag = etag;
            return this;
        }
        @CustomType.Setter
        public Builder isDeleteMarker(@Nullable Boolean isDeleteMarker) {

            this.isDeleteMarker = isDeleteMarker;
            return this;
        }
        @CustomType.Setter
        public Builder isLatest(@Nullable Boolean isLatest) {

            this.isLatest = isLatest;
            return this;
        }
        @CustomType.Setter
        public Builder key(@Nullable String key) {

            this.key = key;
            return this;
        }
        @CustomType.Setter
        public Builder lastModified(@Nullable String lastModified) {

            this.lastModified = lastModified;
            return this;
        }
        @CustomType.Setter
        public Builder size(@Nullable Double size) {

            this.size = size;
            return this;
        }
        @CustomType.Setter
        public Builder storageClass(@Nullable String storageClass) {

            this.storageClass = storageClass;
            return this;
        }
        @CustomType.Setter
        public Builder versionId(@Nullable String versionId) {

            this.versionId = versionId;
            return this;
        }
        public StorageObject build() {
            final var _resultValue = new StorageObject();
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
