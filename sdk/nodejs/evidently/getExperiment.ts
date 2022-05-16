// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Experiment.
 */
export function getExperiment(args: GetExperimentArgs, opts?: pulumi.InvokeOptions): Promise<GetExperimentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:evidently:getExperiment", {
        "arn": args.arn,
    }, opts);
}

export interface GetExperimentArgs {
    arn: string;
}

export interface GetExperimentResult {
    readonly arn?: string;
    readonly description?: string;
    readonly metricGoals?: outputs.evidently.ExperimentMetricGoalObject[];
    readonly onlineAbConfig?: outputs.evidently.ExperimentOnlineAbConfigObject;
    readonly randomizationSalt?: string;
    /**
     * Start Experiment. Default is False
     */
    readonly runningStatus?: outputs.evidently.ExperimentRunningStatusObject;
    readonly samplingRate?: number;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.evidently.ExperimentTag[];
    readonly treatments?: outputs.evidently.ExperimentTreatmentObject[];
}

export function getExperimentOutput(args: GetExperimentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetExperimentResult> {
    return pulumi.output(args).apply(a => getExperiment(a, opts))
}

export interface GetExperimentOutputArgs {
    arn: pulumi.Input<string>;
}