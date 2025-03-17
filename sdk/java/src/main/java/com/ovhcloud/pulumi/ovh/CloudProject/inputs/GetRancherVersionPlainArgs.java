// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetRancherVersionPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetRancherVersionPlainArgs Empty = new GetRancherVersionPlainArgs();

    /**
     * Project ID
     * 
     */
    @Import(name="projectId", required=true)
    private String projectId;

    /**
     * @return Project ID
     * 
     */
    public String projectId() {
        return this.projectId;
    }

    private GetRancherVersionPlainArgs() {}

    private GetRancherVersionPlainArgs(GetRancherVersionPlainArgs $) {
        this.projectId = $.projectId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRancherVersionPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRancherVersionPlainArgs $;

        public Builder() {
            $ = new GetRancherVersionPlainArgs();
        }

        public Builder(GetRancherVersionPlainArgs defaults) {
            $ = new GetRancherVersionPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            $.projectId = projectId;
            return this;
        }

        public GetRancherVersionPlainArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("GetRancherVersionPlainArgs", "projectId");
            }
            return $;
        }
    }

}
