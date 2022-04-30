// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Inspector::AssessmentTarget
 */
export function getAssessmentTarget(args: GetAssessmentTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetAssessmentTargetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:inspector:getAssessmentTarget", {
        "arn": args.arn,
    }, opts);
}

export interface GetAssessmentTargetArgs {
    arn: string;
}

export interface GetAssessmentTargetResult {
    readonly arn?: string;
    readonly resourceGroupArn?: string;
}

export function getAssessmentTargetOutput(args: GetAssessmentTargetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssessmentTargetResult> {
    return pulumi.output(args).apply(a => getAssessmentTarget(a, opts))
}

export interface GetAssessmentTargetOutputArgs {
    arn: pulumi.Input<string>;
}
