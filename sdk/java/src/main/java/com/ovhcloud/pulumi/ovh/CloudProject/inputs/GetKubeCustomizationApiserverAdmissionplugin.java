// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.ovhcloud.pulumi.ovh.CloudProject.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetKubeCustomizationApiserverAdmissionplugin extends com.pulumi.resources.InvokeArgs {

    public static final GetKubeCustomizationApiserverAdmissionplugin Empty = new GetKubeCustomizationApiserverAdmissionplugin();

    /**
     * Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
     * 
     */
    @Import(name="disableds", required=true)
    private List<String> disableds;

    /**
     * @return Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
     * 
     */
    public List<String> disableds() {
        return this.disableds;
    }

    /**
     * Array of admission plugins enabled, default is [&#34;NodeRestriction&#34;,&#34;AlwaysPulImages&#34;] and only these admission plugins can be enabled at this time.
     * 
     */
    @Import(name="enableds", required=true)
    private List<String> enableds;

    /**
     * @return Array of admission plugins enabled, default is [&#34;NodeRestriction&#34;,&#34;AlwaysPulImages&#34;] and only these admission plugins can be enabled at this time.
     * 
     */
    public List<String> enableds() {
        return this.enableds;
    }

    private GetKubeCustomizationApiserverAdmissionplugin() {}

    private GetKubeCustomizationApiserverAdmissionplugin(GetKubeCustomizationApiserverAdmissionplugin $) {
        this.disableds = $.disableds;
        this.enableds = $.enableds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetKubeCustomizationApiserverAdmissionplugin defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetKubeCustomizationApiserverAdmissionplugin $;

        public Builder() {
            $ = new GetKubeCustomizationApiserverAdmissionplugin();
        }

        public Builder(GetKubeCustomizationApiserverAdmissionplugin defaults) {
            $ = new GetKubeCustomizationApiserverAdmissionplugin(Objects.requireNonNull(defaults));
        }

        /**
         * @param disableds Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
         * 
         * @return builder
         * 
         */
        public Builder disableds(List<String> disableds) {
            $.disableds = disableds;
            return this;
        }

        /**
         * @param disableds Array of admission plugins disabled, default is [] and only AlwaysPulImages can be disabled at this time.
         * 
         * @return builder
         * 
         */
        public Builder disableds(String... disableds) {
            return disableds(List.of(disableds));
        }

        /**
         * @param enableds Array of admission plugins enabled, default is [&#34;NodeRestriction&#34;,&#34;AlwaysPulImages&#34;] and only these admission plugins can be enabled at this time.
         * 
         * @return builder
         * 
         */
        public Builder enableds(List<String> enableds) {
            $.enableds = enableds;
            return this;
        }

        /**
         * @param enableds Array of admission plugins enabled, default is [&#34;NodeRestriction&#34;,&#34;AlwaysPulImages&#34;] and only these admission plugins can be enabled at this time.
         * 
         * @return builder
         * 
         */
        public Builder enableds(String... enableds) {
            return enableds(List.of(enableds));
        }

        public GetKubeCustomizationApiserverAdmissionplugin build() {
            if ($.disableds == null) {
                throw new MissingRequiredPropertyException("GetKubeCustomizationApiserverAdmissionplugin", "disableds");
            }
            if ($.enableds == null) {
                throw new MissingRequiredPropertyException("GetKubeCustomizationApiserverAdmissionplugin", "enableds");
            }
            return $;
        }
    }

}
