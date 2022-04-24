// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::CloudWatch::CompositeAlarm type specifies an alarm which aggregates the states of other Alarms (Metric or Composite Alarms) as defined by the AlarmRule expression
 */
export function getCompositeAlarm(args: GetCompositeAlarmArgs, opts?: pulumi.InvokeOptions): Promise<GetCompositeAlarmResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudwatch:getCompositeAlarm", {
        "alarmName": args.alarmName,
    }, opts);
}

export interface GetCompositeAlarmArgs {
    /**
     * The name of the Composite Alarm
     */
    alarmName: string;
}

export interface GetCompositeAlarmResult {
    /**
     * Indicates whether actions should be executed during any changes to the alarm state. The default is TRUE.
     */
    readonly actionsEnabled?: boolean;
    /**
     * The list of actions to execute when this alarm transitions into an ALARM state from any other state. Specify each action as an Amazon Resource Name (ARN).
     */
    readonly alarmActions?: string[];
    /**
     * The description of the alarm
     */
    readonly alarmDescription?: string;
    /**
     * Expression which aggregates the state of other Alarms (Metric or Composite Alarms)
     */
    readonly alarmRule?: string;
    /**
     * Amazon Resource Name (ARN) of the alarm
     */
    readonly arn?: string;
    /**
     * The actions to execute when this alarm transitions to the INSUFFICIENT_DATA state from any other state. Each action is specified as an Amazon Resource Name (ARN).
     */
    readonly insufficientDataActions?: string[];
    /**
     * The actions to execute when this alarm transitions to the OK state from any other state. Each action is specified as an Amazon Resource Name (ARN).
     */
    readonly oKActions?: string[];
}

export function getCompositeAlarmOutput(args: GetCompositeAlarmOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCompositeAlarmResult> {
    return pulumi.output(args).apply(a => getCompositeAlarm(a, opts))
}

export interface GetCompositeAlarmOutputArgs {
    /**
     * The name of the Composite Alarm
     */
    alarmName: pulumi.Input<string>;
}
