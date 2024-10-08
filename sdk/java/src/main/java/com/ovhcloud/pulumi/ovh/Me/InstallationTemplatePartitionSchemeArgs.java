// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.Me;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class InstallationTemplatePartitionSchemeArgs extends com.pulumi.resources.ResourceArgs {

    public static final InstallationTemplatePartitionSchemeArgs Empty = new InstallationTemplatePartitionSchemeArgs();

    /**
     * (Required) This partition scheme name.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return (Required) This partition scheme name.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default, among all the compatible partitioning schemes (given the underlying hardware specifications).
     * 
     */
    @Import(name="priority", required=true)
    private Output<Integer> priority;

    /**
     * @return on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default, among all the compatible partitioning schemes (given the underlying hardware specifications).
     * 
     */
    public Output<Integer> priority() {
        return this.priority;
    }

    /**
     * The template name of the partition scheme.
     * 
     */
    @Import(name="templateName", required=true)
    private Output<String> templateName;

    /**
     * @return The template name of the partition scheme.
     * 
     */
    public Output<String> templateName() {
        return this.templateName;
    }

    private InstallationTemplatePartitionSchemeArgs() {}

    private InstallationTemplatePartitionSchemeArgs(InstallationTemplatePartitionSchemeArgs $) {
        this.name = $.name;
        this.priority = $.priority;
        this.templateName = $.templateName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(InstallationTemplatePartitionSchemeArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private InstallationTemplatePartitionSchemeArgs $;

        public Builder() {
            $ = new InstallationTemplatePartitionSchemeArgs();
        }

        public Builder(InstallationTemplatePartitionSchemeArgs defaults) {
            $ = new InstallationTemplatePartitionSchemeArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name (Required) This partition scheme name.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name (Required) This partition scheme name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param priority on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default, among all the compatible partitioning schemes (given the underlying hardware specifications).
         * 
         * @return builder
         * 
         */
        public Builder priority(Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority on a reinstall, if a partitioning scheme is not specified, the one with the higher priority will be used by default, among all the compatible partitioning schemes (given the underlying hardware specifications).
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param templateName The template name of the partition scheme.
         * 
         * @return builder
         * 
         */
        public Builder templateName(Output<String> templateName) {
            $.templateName = templateName;
            return this;
        }

        /**
         * @param templateName The template name of the partition scheme.
         * 
         * @return builder
         * 
         */
        public Builder templateName(String templateName) {
            return templateName(Output.of(templateName));
        }

        public InstallationTemplatePartitionSchemeArgs build() {
            if ($.priority == null) {
                throw new MissingRequiredPropertyException("InstallationTemplatePartitionSchemeArgs", "priority");
            }
            if ($.templateName == null) {
                throw new MissingRequiredPropertyException("InstallationTemplatePartitionSchemeArgs", "templateName");
            }
            return $;
        }
    }

}
