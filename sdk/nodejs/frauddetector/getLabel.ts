// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * An label for fraud detector.
 */
export function getLabel(args: GetLabelArgs, opts?: pulumi.InvokeOptions): Promise<GetLabelResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:frauddetector:getLabel", {
        "arn": args.arn,
    }, opts);
}

export interface GetLabelArgs {
    /**
     * The label ARN.
     */
    arn: string;
}

export interface GetLabelResult {
    /**
     * The label ARN.
     */
    readonly arn?: string;
    /**
     * The timestamp when the label was created.
     */
    readonly createdTime?: string;
    /**
     * The label description.
     */
    readonly description?: string;
    /**
     * The timestamp when the label was last updated.
     */
    readonly lastUpdatedTime?: string;
    /**
     * Tags associated with this label.
     */
    readonly tags?: outputs.frauddetector.LabelTag[];
}

export function getLabelOutput(args: GetLabelOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLabelResult> {
    return pulumi.output(args).apply(a => getLabel(a, opts))
}

export interface GetLabelOutputArgs {
    /**
     * The label ARN.
     */
    arn: pulumi.Input<string>;
}