// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Glue::Trigger
 */
export function getTrigger(args: GetTriggerArgs, opts?: pulumi.InvokeOptions): Promise<GetTriggerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:glue:getTrigger", {
        "id": args.id,
    }, opts);
}

export interface GetTriggerArgs {
    id: string;
}

export interface GetTriggerResult {
    readonly actions?: outputs.glue.TriggerAction[];
    readonly description?: string;
    readonly id?: string;
    readonly predicate?: outputs.glue.TriggerPredicate;
    readonly schedule?: string;
    readonly startOnCreation?: boolean;
    readonly tags?: any;
    readonly type?: string;
}

export function getTriggerOutput(args: GetTriggerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTriggerResult> {
    return pulumi.output(args).apply(a => getTrigger(a, opts))
}

export interface GetTriggerOutputArgs {
    id: pulumi.Input<string>;
}
