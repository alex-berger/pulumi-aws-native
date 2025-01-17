// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Batch::JobQueue
 */
export function getJobQueue(args: GetJobQueueArgs, opts?: pulumi.InvokeOptions): Promise<GetJobQueueResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:batch:getJobQueue", {
        "jobQueueArn": args.jobQueueArn,
    }, opts);
}

export interface GetJobQueueArgs {
    jobQueueArn: string;
}

export interface GetJobQueueResult {
    readonly computeEnvironmentOrder?: outputs.batch.JobQueueComputeEnvironmentOrder[];
    readonly jobQueueArn?: string;
    readonly priority?: number;
    readonly schedulingPolicyArn?: string;
    readonly state?: enums.batch.JobQueueState;
}

export function getJobQueueOutput(args: GetJobQueueOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetJobQueueResult> {
    return pulumi.output(args).apply(a => getJobQueue(a, opts))
}

export interface GetJobQueueOutputArgs {
    jobQueueArn: pulumi.Input<string>;
}
