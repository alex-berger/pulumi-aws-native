// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Evidently::Feature.
 */
export function getFeature(args: GetFeatureArgs, opts?: pulumi.InvokeOptions): Promise<GetFeatureResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:evidently:getFeature", {
        "arn": args.arn,
    }, opts);
}

export interface GetFeatureArgs {
    arn: string;
}

export interface GetFeatureResult {
    readonly arn?: string;
    readonly defaultVariation?: string;
    readonly description?: string;
    readonly entityOverrides?: outputs.evidently.FeatureEntityOverride[];
    readonly evaluationStrategy?: enums.evidently.FeatureEvaluationStrategy;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.evidently.FeatureTag[];
    readonly variations?: outputs.evidently.FeatureVariationObject[];
}

export function getFeatureOutput(args: GetFeatureOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFeatureResult> {
    return pulumi.output(args).apply(a => getFeature(a, opts))
}

export interface GetFeatureOutputArgs {
    arn: pulumi.Input<string>;
}