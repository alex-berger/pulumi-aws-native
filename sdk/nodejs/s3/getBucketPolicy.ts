// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::S3::BucketPolicy
 */
export function getBucketPolicy(args: GetBucketPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetBucketPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:s3:getBucketPolicy", {
        "id": args.id,
    }, opts);
}

export interface GetBucketPolicyArgs {
    id: string;
}

export interface GetBucketPolicyResult {
    readonly id?: string;
    readonly policyDocument?: any;
}

export function getBucketPolicyOutput(args: GetBucketPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBucketPolicyResult> {
    return pulumi.output(args).apply(a => getBucketPolicy(a, opts))
}

export interface GetBucketPolicyOutputArgs {
    id: pulumi.Input<string>;
}
