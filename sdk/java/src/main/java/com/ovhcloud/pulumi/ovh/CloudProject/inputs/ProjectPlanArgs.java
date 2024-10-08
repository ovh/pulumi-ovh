// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.ovhcloud.pulumi.ovh.CloudProject.inputs.ProjectPlanConfigurationArgs;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectPlanArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectPlanArgs Empty = new ProjectPlanArgs();

    /**
     * Catalog name
     * 
     */
    @Import(name="catalogName")
    private @Nullable Output<String> catalogName;

    /**
     * @return Catalog name
     * 
     */
    public Optional<Output<String>> catalogName() {
        return Optional.ofNullable(this.catalogName);
    }

    /**
     * Representation of a configuration item for personalizing product
     * 
     */
    @Import(name="configurations")
    private @Nullable Output<List<ProjectPlanConfigurationArgs>> configurations;

    /**
     * @return Representation of a configuration item for personalizing product
     * 
     */
    public Optional<Output<List<ProjectPlanConfigurationArgs>>> configurations() {
        return Optional.ofNullable(this.configurations);
    }

    /**
     * duration
     * 
     */
    @Import(name="duration", required=true)
    private Output<String> duration;

    /**
     * @return duration
     * 
     */
    public Output<String> duration() {
        return this.duration;
    }

    /**
     * Plan code. This value must be adapted depending on your `OVH_ENDPOINT` value. It&#39;s `project.2018` for `ovh-{eu,ca}` and `project` when using `ovh-us`.
     * 
     */
    @Import(name="planCode", required=true)
    private Output<String> planCode;

    /**
     * @return Plan code. This value must be adapted depending on your `OVH_ENDPOINT` value. It&#39;s `project.2018` for `ovh-{eu,ca}` and `project` when using `ovh-us`.
     * 
     */
    public Output<String> planCode() {
        return this.planCode;
    }

    /**
     * Pricing model identifier
     * 
     */
    @Import(name="pricingMode", required=true)
    private Output<String> pricingMode;

    /**
     * @return Pricing model identifier
     * 
     */
    public Output<String> pricingMode() {
        return this.pricingMode;
    }

    private ProjectPlanArgs() {}

    private ProjectPlanArgs(ProjectPlanArgs $) {
        this.catalogName = $.catalogName;
        this.configurations = $.configurations;
        this.duration = $.duration;
        this.planCode = $.planCode;
        this.pricingMode = $.pricingMode;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectPlanArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectPlanArgs $;

        public Builder() {
            $ = new ProjectPlanArgs();
        }

        public Builder(ProjectPlanArgs defaults) {
            $ = new ProjectPlanArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param catalogName Catalog name
         * 
         * @return builder
         * 
         */
        public Builder catalogName(@Nullable Output<String> catalogName) {
            $.catalogName = catalogName;
            return this;
        }

        /**
         * @param catalogName Catalog name
         * 
         * @return builder
         * 
         */
        public Builder catalogName(String catalogName) {
            return catalogName(Output.of(catalogName));
        }

        /**
         * @param configurations Representation of a configuration item for personalizing product
         * 
         * @return builder
         * 
         */
        public Builder configurations(@Nullable Output<List<ProjectPlanConfigurationArgs>> configurations) {
            $.configurations = configurations;
            return this;
        }

        /**
         * @param configurations Representation of a configuration item for personalizing product
         * 
         * @return builder
         * 
         */
        public Builder configurations(List<ProjectPlanConfigurationArgs> configurations) {
            return configurations(Output.of(configurations));
        }

        /**
         * @param configurations Representation of a configuration item for personalizing product
         * 
         * @return builder
         * 
         */
        public Builder configurations(ProjectPlanConfigurationArgs... configurations) {
            return configurations(List.of(configurations));
        }

        /**
         * @param duration duration
         * 
         * @return builder
         * 
         */
        public Builder duration(Output<String> duration) {
            $.duration = duration;
            return this;
        }

        /**
         * @param duration duration
         * 
         * @return builder
         * 
         */
        public Builder duration(String duration) {
            return duration(Output.of(duration));
        }

        /**
         * @param planCode Plan code. This value must be adapted depending on your `OVH_ENDPOINT` value. It&#39;s `project.2018` for `ovh-{eu,ca}` and `project` when using `ovh-us`.
         * 
         * @return builder
         * 
         */
        public Builder planCode(Output<String> planCode) {
            $.planCode = planCode;
            return this;
        }

        /**
         * @param planCode Plan code. This value must be adapted depending on your `OVH_ENDPOINT` value. It&#39;s `project.2018` for `ovh-{eu,ca}` and `project` when using `ovh-us`.
         * 
         * @return builder
         * 
         */
        public Builder planCode(String planCode) {
            return planCode(Output.of(planCode));
        }

        /**
         * @param pricingMode Pricing model identifier
         * 
         * @return builder
         * 
         */
        public Builder pricingMode(Output<String> pricingMode) {
            $.pricingMode = pricingMode;
            return this;
        }

        /**
         * @param pricingMode Pricing model identifier
         * 
         * @return builder
         * 
         */
        public Builder pricingMode(String pricingMode) {
            return pricingMode(Output.of(pricingMode));
        }

        public ProjectPlanArgs build() {
            if ($.duration == null) {
                throw new MissingRequiredPropertyException("ProjectPlanArgs", "duration");
            }
            if ($.planCode == null) {
                throw new MissingRequiredPropertyException("ProjectPlanArgs", "planCode");
            }
            if ($.pricingMode == null) {
                throw new MissingRequiredPropertyException("ProjectPlanArgs", "pricingMode");
            }
            return $;
        }
    }

}
