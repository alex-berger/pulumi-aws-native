// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The `AWS::Redshift::ScheduledAction` resource creates an Amazon Redshift Scheduled Action.
 */
export function getScheduledAction(args: GetScheduledActionArgs, opts?: pulumi.InvokeOptions): Promise<GetScheduledActionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:redshift:getScheduledAction", {
        "scheduledActionName": args.scheduledActionName,
    }, opts);
}

export interface GetScheduledActionArgs {
    /**
     * The name of the scheduled action. The name must be unique within an account.
     */
    scheduledActionName: string;
}

export interface GetScheduledActionResult {
    /**
     * If true, the schedule is enabled. If false, the scheduled action does not trigger.
     */
    readonly enable?: boolean;
    /**
     * The end time in UTC of the scheduled action. After this time, the scheduled action does not trigger.
     */
    readonly endTime?: string;
    /**
     * The IAM role to assume to run the target action.
     */
    readonly iamRole?: string;
    /**
     * List of times when the scheduled action will run.
     */
    readonly nextInvocations?: string[];
    /**
     * The schedule in `at( )` or `cron( )` format.
     */
    readonly schedule?: string;
    /**
     * The description of the scheduled action.
     */
    readonly scheduledActionDescription?: string;
    /**
     * The start time in UTC of the scheduled action. Before this time, the scheduled action does not trigger.
     */
    readonly startTime?: string;
    /**
     * The state of the scheduled action.
     */
    readonly state?: enums.redshift.ScheduledActionState;
    /**
     * A JSON format string of the Amazon Redshift API operation with input parameters.
     */
    readonly targetAction?: outputs.redshift.ScheduledActionType;
}

export function getScheduledActionOutput(args: GetScheduledActionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScheduledActionResult> {
    return pulumi.output(args).apply(a => getScheduledAction(a, opts))
}

export interface GetScheduledActionOutputArgs {
    /**
     * The name of the scheduled action. The name must be unique within an account.
     */
    scheduledActionName: pulumi.Input<string>;
}