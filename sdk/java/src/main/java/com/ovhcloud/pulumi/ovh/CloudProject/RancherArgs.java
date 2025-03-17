// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.RancherTargetSpecArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class RancherArgs extends com.pulumi.resources.ResourceArgs {

    public static final RancherArgs Empty = new RancherArgs();

    /**
     * Project ID
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return Project ID
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * Target specification for the managed Rancher service
     * 
     */
    @Import(name="targetSpec", required=true)
    private Output<RancherTargetSpecArgs> targetSpec;

    /**
     * @return Target specification for the managed Rancher service
     * 
     */
    public Output<RancherTargetSpecArgs> targetSpec() {
        return this.targetSpec;
    }

    private RancherArgs() {}

    private RancherArgs(RancherArgs $) {
        this.projectId = $.projectId;
        this.targetSpec = $.targetSpec;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RancherArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RancherArgs $;

        public Builder() {
            $ = new RancherArgs();
        }

        public Builder(RancherArgs defaults) {
            $ = new RancherArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param targetSpec Target specification for the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder targetSpec(Output<RancherTargetSpecArgs> targetSpec) {
            $.targetSpec = targetSpec;
            return this;
        }

        /**
         * @param targetSpec Target specification for the managed Rancher service
         * 
         * @return builder
         * 
         */
        public Builder targetSpec(RancherTargetSpecArgs targetSpec) {
            return targetSpec(Output.of(targetSpec));
        }

        public RancherArgs build() {
            if ($.projectId == null) {
                throw new MissingRequiredPropertyException("RancherArgs", "projectId");
            }
            if ($.targetSpec == null) {
                throw new MissingRequiredPropertyException("RancherArgs", "targetSpec");
            }
            return $;
        }
    }

}
