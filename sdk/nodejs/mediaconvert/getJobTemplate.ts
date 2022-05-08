// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::MediaConvert::JobTemplate
 */
export function getJobTemplate(args: GetJobTemplateArgs, opts?: pulumi.InvokeOptions): Promise<GetJobTemplateResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:mediaconvert:getJobTemplate", {
        "id": args.id,
    }, opts);
}

export interface GetJobTemplateArgs {
    id: string;
}

export interface GetJobTemplateResult {
    readonly accelerationSettings?: outputs.mediaconvert.JobTemplateAccelerationSettings;
    readonly arn?: string;
    readonly category?: string;
    readonly description?: string;
    readonly hopDestinations?: outputs.mediaconvert.JobTemplateHopDestination[];
    readonly id?: string;
    readonly priority?: number;
    readonly queue?: string;
    readonly settingsJson?: any;
    readonly statusUpdateInterval?: string;
    readonly tags?: any;
}

export function getJobTemplateOutput(args: GetJobTemplateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobTemplateResult> {
    return pulumi.output(args).apply(a => getJobTemplate(a, opts))
}

export interface GetJobTemplateOutputArgs {
    id: pulumi.Input<string>;
}
