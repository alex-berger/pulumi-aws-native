// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CloudFront::Distribution
 */
export function getDistribution(args: GetDistributionArgs, opts?: pulumi.InvokeOptions): Promise<GetDistributionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloudfront:getDistribution", {
        "id": args.id,
    }, opts);
}

export interface GetDistributionArgs {
    id: string;
}

export interface GetDistributionResult {
    readonly distributionConfig?: outputs.cloudfront.DistributionConfig;
    readonly domainName?: string;
    readonly id?: string;
    readonly tags?: outputs.cloudfront.DistributionTag[];
}

export function getDistributionOutput(args: GetDistributionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDistributionResult> {
    return pulumi.output(args).apply(a => getDistribution(a, opts))
}

export interface GetDistributionOutputArgs {
    id: pulumi.Input<string>;
}
