// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html
 */
export class Task extends pulumi.CustomResource {
    /**
     * Get an existing Task resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Task {
        return new Task(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:datasync:Task';

    /**
     * Returns true if the given object is an instance of Task.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Task {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Task.__pulumiType;
    }

    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-cloudwatchloggrouparn
     */
    public readonly cloudWatchLogGroupArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-destinationlocationarn
     */
    public readonly destinationLocationArn!: pulumi.Output<string>;
    public /*out*/ readonly destinationNetworkInterfaceArns!: pulumi.Output<string[]>;
    public /*out*/ readonly errorCode!: pulumi.Output<string>;
    public /*out*/ readonly errorDetail!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-excludes
     */
    public readonly excludes!: pulumi.Output<outputs.datasync.TaskFilterRule[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-options
     */
    public readonly options!: pulumi.Output<outputs.datasync.TaskOptions | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-schedule
     */
    public readonly schedule!: pulumi.Output<outputs.datasync.TaskTaskSchedule | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-sourcelocationarn
     */
    public readonly sourceLocationArn!: pulumi.Output<string>;
    public /*out*/ readonly sourceNetworkInterfaceArns!: pulumi.Output<string[]>;
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    public /*out*/ readonly taskArn!: pulumi.Output<string>;

    /**
     * Create a Task resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TaskArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.destinationLocationArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'destinationLocationArn'");
            }
            if ((!args || args.sourceLocationArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sourceLocationArn'");
            }
            inputs["cloudWatchLogGroupArn"] = args ? args.cloudWatchLogGroupArn : undefined;
            inputs["destinationLocationArn"] = args ? args.destinationLocationArn : undefined;
            inputs["excludes"] = args ? args.excludes : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["options"] = args ? args.options : undefined;
            inputs["schedule"] = args ? args.schedule : undefined;
            inputs["sourceLocationArn"] = args ? args.sourceLocationArn : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["destinationNetworkInterfaceArns"] = undefined /*out*/;
            inputs["errorCode"] = undefined /*out*/;
            inputs["errorDetail"] = undefined /*out*/;
            inputs["sourceNetworkInterfaceArns"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["taskArn"] = undefined /*out*/;
        } else {
            inputs["cloudWatchLogGroupArn"] = undefined /*out*/;
            inputs["destinationLocationArn"] = undefined /*out*/;
            inputs["destinationNetworkInterfaceArns"] = undefined /*out*/;
            inputs["errorCode"] = undefined /*out*/;
            inputs["errorDetail"] = undefined /*out*/;
            inputs["excludes"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["options"] = undefined /*out*/;
            inputs["schedule"] = undefined /*out*/;
            inputs["sourceLocationArn"] = undefined /*out*/;
            inputs["sourceNetworkInterfaceArns"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["taskArn"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Task.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Task resource.
 */
export interface TaskArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-cloudwatchloggrouparn
     */
    cloudWatchLogGroupArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-destinationlocationarn
     */
    destinationLocationArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-excludes
     */
    excludes?: pulumi.Input<pulumi.Input<inputs.datasync.TaskFilterRuleArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-name
     */
    name?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-options
     */
    options?: pulumi.Input<inputs.datasync.TaskOptionsArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-schedule
     */
    schedule?: pulumi.Input<inputs.datasync.TaskTaskScheduleArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-sourcelocationarn
     */
    sourceLocationArn: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-datasync-task.html#cfn-datasync-task-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
}