// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AutoScaling::ScalingPolicy
 */
export function getScalingPolicy(args: GetScalingPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetScalingPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:autoscaling:getScalingPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetScalingPolicyArgs {
    id: string;
}

export interface GetScalingPolicyResult {
    readonly adjustmentType?: string;
    readonly autoScalingGroupName?: string;
    readonly cooldown?: string;
    readonly estimatedInstanceWarmup?: number;
    readonly id?: string;
    readonly metricAggregationType?: string;
    readonly minAdjustmentMagnitude?: number;
    readonly policyType?: string;
    readonly predictiveScalingConfiguration?: outputs.autoscaling.ScalingPolicyPredictiveScalingConfiguration;
    readonly scalingAdjustment?: number;
    readonly stepAdjustments?: outputs.autoscaling.ScalingPolicyStepAdjustment[];
    readonly targetTrackingConfiguration?: outputs.autoscaling.ScalingPolicyTargetTrackingConfiguration;
}

export function getScalingPolicyOutput(args: GetScalingPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScalingPolicyResult> {
    return pulumi.output(args).apply(a => getScalingPolicy(a, opts))
}

export interface GetScalingPolicyOutputArgs {
    id: pulumi.Input<string>;
}
