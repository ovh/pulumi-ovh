// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Okms.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetOkmsServiceKeyJwkPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOkmsServiceKeyJwkPlainArgs Empty = new GetOkmsServiceKeyJwkPlainArgs();

    /**
     * ID of the service key
     * 
     */
    @Import(name="id", required=true)
    private String id;

    /**
     * @return ID of the service key
     * 
     */
    public String id() {
        return this.id;
    }

    /**
     * ID of the KMS
     * 
     */
    @Import(name="okmsId", required=true)
    private String okmsId;

    /**
     * @return ID of the KMS
     * 
     */
    public String okmsId() {
        return this.okmsId;
    }

    private GetOkmsServiceKeyJwkPlainArgs() {}

    private GetOkmsServiceKeyJwkPlainArgs(GetOkmsServiceKeyJwkPlainArgs $) {
        this.id = $.id;
        this.okmsId = $.okmsId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOkmsServiceKeyJwkPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOkmsServiceKeyJwkPlainArgs $;

        public Builder() {
            $ = new GetOkmsServiceKeyJwkPlainArgs();
        }

        public Builder(GetOkmsServiceKeyJwkPlainArgs defaults) {
            $ = new GetOkmsServiceKeyJwkPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param id ID of the service key
         * 
         * @return builder
         * 
         */
        public Builder id(String id) {
            $.id = id;
            return this;
        }

        /**
         * @param okmsId ID of the KMS
         * 
         * @return builder
         * 
         */
        public Builder okmsId(String okmsId) {
            $.okmsId = okmsId;
            return this;
        }

        public GetOkmsServiceKeyJwkPlainArgs build() {
            if ($.id == null) {
                throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkPlainArgs", "id");
            }
            if ($.okmsId == null) {
                throw new MissingRequiredPropertyException("GetOkmsServiceKeyJwkPlainArgs", "okmsId");
            }
            return $;
        }
    }

}
