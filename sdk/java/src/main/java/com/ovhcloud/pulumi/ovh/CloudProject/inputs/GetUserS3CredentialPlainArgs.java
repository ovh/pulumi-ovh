// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetUserS3CredentialPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetUserS3CredentialPlainArgs Empty = new GetUserS3CredentialPlainArgs();

    /**
     * the Access Key ID
     * 
     */
    @Import(name="accessKeyId", required=true)
    private String accessKeyId;

    /**
     * @return the Access Key ID
     * 
     */
    public String accessKeyId() {
        return this.accessKeyId;
    }

    /**
     * The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    @Import(name="serviceName", required=true)
    private String serviceName;

    /**
     * @return The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     * 
     */
    public String serviceName() {
        return this.serviceName;
    }

    /**
     * The ID of a public cloud project&#39;s user.
     * 
     */
    @Import(name="userId", required=true)
    private String userId;

    /**
     * @return The ID of a public cloud project&#39;s user.
     * 
     */
    public String userId() {
        return this.userId;
    }

    private GetUserS3CredentialPlainArgs() {}

    private GetUserS3CredentialPlainArgs(GetUserS3CredentialPlainArgs $) {
        this.accessKeyId = $.accessKeyId;
        this.serviceName = $.serviceName;
        this.userId = $.userId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetUserS3CredentialPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetUserS3CredentialPlainArgs $;

        public Builder() {
            $ = new GetUserS3CredentialPlainArgs();
        }

        public Builder(GetUserS3CredentialPlainArgs defaults) {
            $ = new GetUserS3CredentialPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessKeyId the Access Key ID
         * 
         * @return builder
         * 
         */
        public Builder accessKeyId(String accessKeyId) {
            $.accessKeyId = accessKeyId;
            return this;
        }

        /**
         * @param serviceName The ID of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
         * 
         * @return builder
         * 
         */
        public Builder serviceName(String serviceName) {
            $.serviceName = serviceName;
            return this;
        }

        /**
         * @param userId The ID of a public cloud project&#39;s user.
         * 
         * @return builder
         * 
         */
        public Builder userId(String userId) {
            $.userId = userId;
            return this;
        }

        public GetUserS3CredentialPlainArgs build() {
            if ($.accessKeyId == null) {
                throw new MissingRequiredPropertyException("GetUserS3CredentialPlainArgs", "accessKeyId");
            }
            if ($.serviceName == null) {
                throw new MissingRequiredPropertyException("GetUserS3CredentialPlainArgs", "serviceName");
            }
            if ($.userId == null) {
                throw new MissingRequiredPropertyException("GetUserS3CredentialPlainArgs", "userId");
            }
            return $;
        }
    }

}
