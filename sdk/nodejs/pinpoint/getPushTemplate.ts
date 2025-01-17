// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Pinpoint::PushTemplate
 */
export function getPushTemplate(args: GetPushTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetPushTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpoint:getPushTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetPushTemplateArgs {
    id: string;
}

export interface GetPushTemplateResult {
    readonly aDM?: outputs.pinpoint.PushTemplateAndroidPushNotificationTemplate;
    readonly aPNS?: outputs.pinpoint.PushTemplateAPNSPushNotificationTemplate;
    readonly arn?: string;
    readonly baidu?: outputs.pinpoint.PushTemplateAndroidPushNotificationTemplate;
    readonly default?: outputs.pinpoint.PushTemplateDefaultPushNotificationTemplate;
    readonly defaultSubstitutions?: string;
    readonly gCM?: outputs.pinpoint.PushTemplateAndroidPushNotificationTemplate;
    readonly id?: string;
    readonly tags?: any;
    readonly templateDescription?: string;
}

export function getPushTemplateOutput(args: GetPushTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPushTemplateResult> {
    return pulumi.output(args).apply(a => getPushTemplate(a, opts))
}

export interface GetPushTemplateOutputArgs {
    id: pulumi.Input<string>;
}
