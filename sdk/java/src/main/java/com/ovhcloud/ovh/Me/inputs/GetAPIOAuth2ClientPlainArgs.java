// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.ovh.Me.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetAPIOAuth2ClientPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAPIOAuth2ClientPlainArgs Empty = new GetAPIOAuth2ClientPlainArgs();

    /**
     * Client ID of an existing OAuth2 service account.
     * 
     */
    @Import(name="clientId", required=true)
    private String clientId;

    /**
     * @return Client ID of an existing OAuth2 service account.
     * 
     */
    public String clientId() {
        return this.clientId;
    }

    private GetAPIOAuth2ClientPlainArgs() {}

    private GetAPIOAuth2ClientPlainArgs(GetAPIOAuth2ClientPlainArgs $) {
        this.clientId = $.clientId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAPIOAuth2ClientPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAPIOAuth2ClientPlainArgs $;

        public Builder() {
            $ = new GetAPIOAuth2ClientPlainArgs();
        }

        public Builder(GetAPIOAuth2ClientPlainArgs defaults) {
            $ = new GetAPIOAuth2ClientPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId Client ID of an existing OAuth2 service account.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            $.clientId = clientId;
            return this;
        }

        public GetAPIOAuth2ClientPlainArgs build() {
            if ($.clientId == null) {
                throw new MissingRequiredPropertyException("GetAPIOAuth2ClientPlainArgs", "clientId");
            }
            return $;
        }
    }

}