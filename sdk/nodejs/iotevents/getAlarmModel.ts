// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::IoTEvents::AlarmModel resource creates a alarm model. AWS IoT Events alarms help you monitor your data for changes. The data can be metrics that you measure for your equipment and processes. You can create alarms that send notifications when a threshold is breached. Alarms help you detect issues, streamline maintenance, and optimize performance of your equipment and processes.
 *
 * Alarms are instances of alarm models. The alarm model specifies what to detect, when to send notifications, who gets notified, and more. You can also specify one or more supported actions that occur when the alarm state changes. AWS IoT Events routes input attributes derived from your data to the appropriate alarms. If the data that you're monitoring is outside the specified range, the alarm is invoked. You can also acknowledge the alarms or set them to the snooze mode.
 */
export function getAlarmModel(args: GetAlarmModelArgs, opts?: pulumi.InvokeOptions): Promise<GetAlarmModelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotevents:getAlarmModel", {
        "alarmModelName": args.alarmModelName,
    }, opts);
}

export interface GetAlarmModelArgs {
    /**
     * The name of the alarm model.
     */
    alarmModelName: string;
}

export interface GetAlarmModelResult {
    readonly alarmCapabilities?: outputs.iotevents.AlarmModelAlarmCapabilities;
    readonly alarmEventActions?: outputs.iotevents.AlarmModelAlarmEventActions;
    /**
     * A brief description of the alarm model.
     */
    readonly alarmModelDescription?: string;
    readonly alarmRule?: outputs.iotevents.AlarmModelAlarmRule;
    /**
     * The ARN of the role that grants permission to AWS IoT Events to perform its operations.
     */
    readonly roleArn?: string;
    /**
     * A non-negative integer that reflects the severity level of the alarm.
     */
    readonly severity?: number;
    /**
     * An array of key-value pairs to apply to this resource.
     *
     * For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
     */
    readonly tags?: outputs.iotevents.AlarmModelTag[];
}

export function getAlarmModelOutput(args: GetAlarmModelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAlarmModelResult> {
    return pulumi.output(args).apply(a => getAlarmModel(a, opts))
}

export interface GetAlarmModelOutputArgs {
    /**
     * The name of the alarm model.
     */
    alarmModelName: pulumi.Input<string>;
}
