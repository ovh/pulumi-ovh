// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetInstallationTemplatePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetInstallationTemplatePlainArgs Empty = new GetInstallationTemplatePlainArgs();

    /**
     * Template name.
     * 
     */
    @Import(name="templateName", required=true)
    private String templateName;

    /**
     * @return Template name.
     * 
     */
    public String templateName() {
        return this.templateName;
    }

    private GetInstallationTemplatePlainArgs() {}

    private GetInstallationTemplatePlainArgs(GetInstallationTemplatePlainArgs $) {
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetInstallationTemplatePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetInstallationTemplatePlainArgs $;

        public Builder() {
            $ = new GetInstallationTemplatePlainArgs();
        }

        public Builder(GetInstallationTemplatePlainArgs defaults) {
            $ = new GetInstallationTemplatePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param templateName Template name.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            $.templateName = templateName;
            return this;
        }

        public GetInstallationTemplatePlainArgs build() {
            if ($.templateName == null) {
                throw new MissingRequiredPropertyException("GetInstallationTemplatePlainArgs", "templateName");
            }
            return $;
        }
    }

}