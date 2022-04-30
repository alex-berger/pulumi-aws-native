// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The policy to be attached to a Multi Region Access Point
 */
export function getMultiRegionAccessPointPolicy(args: GetMultiRegionAccessPointPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetMultiRegionAccessPointPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:s3:getMultiRegionAccessPointPolicy", {
        "mrapName": args.mrapName,
    }, opts);
}

export interface GetMultiRegionAccessPointPolicyArgs {
    /**
     * The name of the Multi Region Access Point to apply policy
     */
    mrapName: string;
}

export interface GetMultiRegionAccessPointPolicyResult {
    /**
     * Policy document to apply to a Multi Region Access Point
     */
    readonly policy?: any;
    /**
     * The Policy Status associated with this Multi Region Access Point
     */
    readonly policyStatus?: outputs.s3.PolicyStatusProperties;
}

export function getMultiRegionAccessPointPolicyOutput(args: GetMultiRegionAccessPointPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMultiRegionAccessPointPolicyResult> {
    return pulumi.output(args).apply(a => getMultiRegionAccessPointPolicy(a, opts))
}

export interface GetMultiRegionAccessPointPolicyOutputArgs {
    /**
     * The name of the Multi Region Access Point to apply policy
     */
    mrapName: pulumi.Input<string>;
}
