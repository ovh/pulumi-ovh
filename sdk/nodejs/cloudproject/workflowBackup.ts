// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manage a worflow that schedules backups of public cloud instance. Note that upon deletion, the workflow is deleted but any backups that have been created by this workflow are not.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as ovh from "@ovhcloud/pulumi-ovh";
 *
 * const myBackup = new ovh.cloudproject.WorkflowBackup("myBackup", {
 *     cron: "50 4 * * *",
 *     instanceId: "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxx",
 *     maxExecutionCount: 0,
 *     regionName: "GRA11",
 *     rotation: 7,
 *     serviceName: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
 * });
 * ```
 */
export class WorkflowBackup extends pulumi.CustomResource {
    /**
     * Get an existing WorkflowBackup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: WorkflowBackupState, opts?: pulumi.CustomResourceOptions): WorkflowBackup {
        return new WorkflowBackup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'ovh:CloudProject/workflowBackup:WorkflowBackup';

    /**
     * Returns true if the given object is an instance of WorkflowBackup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is WorkflowBackup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === WorkflowBackup.__pulumiType;
    }

    /**
     * The name of the backup files that are created. If empty, the `name` attribute is used.
     */
    public readonly backupName!: pulumi.Output<string>;
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * The cron periodicity at which the backup workflow is scheduled
     *
     * * `instanceId` the id of the instance to back up
     */
    public readonly cron!: pulumi.Output<string>;
    public readonly instanceId!: pulumi.Output<string>;
    /**
     * The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
     */
    public readonly maxExecutionCount!: pulumi.Output<number | undefined>;
    /**
     * The worflow name that is used in the UI
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the openstack region.
     */
    public readonly regionName!: pulumi.Output<string>;
    /**
     * The number of backup that are retained.
     */
    public readonly rotation!: pulumi.Output<number>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    public readonly serviceName!: pulumi.Output<string>;

    /**
     * Create a WorkflowBackup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowBackupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowBackupArgs | WorkflowBackupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as WorkflowBackupState | undefined;
            resourceInputs["backupName"] = state ? state.backupName : undefined;
            resourceInputs["createdAt"] = state ? state.createdAt : undefined;
            resourceInputs["cron"] = state ? state.cron : undefined;
            resourceInputs["instanceId"] = state ? state.instanceId : undefined;
            resourceInputs["maxExecutionCount"] = state ? state.maxExecutionCount : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["regionName"] = state ? state.regionName : undefined;
            resourceInputs["rotation"] = state ? state.rotation : undefined;
            resourceInputs["serviceName"] = state ? state.serviceName : undefined;
        } else {
            const args = argsOrState as WorkflowBackupArgs | undefined;
            if ((!args || args.cron === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cron'");
            }
            if ((!args || args.instanceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'instanceId'");
            }
            if ((!args || args.regionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'regionName'");
            }
            if ((!args || args.rotation === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rotation'");
            }
            if ((!args || args.serviceName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceName'");
            }
            resourceInputs["backupName"] = args ? args.backupName : undefined;
            resourceInputs["cron"] = args ? args.cron : undefined;
            resourceInputs["instanceId"] = args ? args.instanceId : undefined;
            resourceInputs["maxExecutionCount"] = args ? args.maxExecutionCount : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["regionName"] = args ? args.regionName : undefined;
            resourceInputs["rotation"] = args ? args.rotation : undefined;
            resourceInputs["serviceName"] = args ? args.serviceName : undefined;
            resourceInputs["createdAt"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(WorkflowBackup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering WorkflowBackup resources.
 */
export interface WorkflowBackupState {
    /**
     * The name of the backup files that are created. If empty, the `name` attribute is used.
     */
    backupName?: pulumi.Input<string>;
    createdAt?: pulumi.Input<string>;
    /**
     * The cron periodicity at which the backup workflow is scheduled
     *
     * * `instanceId` the id of the instance to back up
     */
    cron?: pulumi.Input<string>;
    instanceId?: pulumi.Input<string>;
    /**
     * The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
     */
    maxExecutionCount?: pulumi.Input<number>;
    /**
     * The worflow name that is used in the UI
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the openstack region.
     */
    regionName?: pulumi.Input<string>;
    /**
     * The number of backup that are retained.
     */
    rotation?: pulumi.Input<number>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a WorkflowBackup resource.
 */
export interface WorkflowBackupArgs {
    /**
     * The name of the backup files that are created. If empty, the `name` attribute is used.
     */
    backupName?: pulumi.Input<string>;
    /**
     * The cron periodicity at which the backup workflow is scheduled
     *
     * * `instanceId` the id of the instance to back up
     */
    cron: pulumi.Input<string>;
    instanceId: pulumi.Input<string>;
    /**
     * The number of times the worflow is run. Default value is `0` which means that the workflow will be scheduled continously until its deletion
     */
    maxExecutionCount?: pulumi.Input<number>;
    /**
     * The worflow name that is used in the UI
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the openstack region.
     */
    regionName: pulumi.Input<string>;
    /**
     * The number of backup that are retained.
     */
    rotation: pulumi.Input<number>;
    /**
     * The id of the public cloud project. If omitted, the `OVH_CLOUD_PROJECT_SERVICE` environment variable is used.
     */
    serviceName: pulumi.Input<string>;
}
