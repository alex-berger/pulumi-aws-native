// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaPackage::Asset
 */
export function getAsset(args: GetAssetArgs, opts?: pulumi.InvokeOptions): Promise<GetAssetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:mediapackage:getAsset", {
        "id": args.id,
    }, opts);
}

export interface GetAssetArgs {
    /**
     * The unique identifier for the Asset.
     */
    id: string;
}

export interface GetAssetResult {
    /**
     * The ARN of the Asset.
     */
    readonly arn?: string;
    /**
     * The time the Asset was initially submitted for Ingest.
     */
    readonly createdAt?: string;
    /**
     * The list of egress endpoints available for the Asset.
     */
    readonly egressEndpoints?: outputs.mediapackage.AssetEgressEndpoint[];
    /**
     * The unique identifier for the Asset.
     */
    readonly id?: string;
    /**
     * The ID of the PackagingGroup for the Asset.
     */
    readonly packagingGroupId?: string;
    /**
     * The resource ID to include in SPEKE key requests.
     */
    readonly resourceId?: string;
    /**
     * ARN of the source object in S3.
     */
    readonly sourceArn?: string;
    /**
     * The IAM role_arn used to access the source S3 bucket.
     */
    readonly sourceRoleArn?: string;
    /**
     * A collection of tags associated with a resource
     */
    readonly tags?: outputs.mediapackage.AssetTag[];
}

export function getAssetOutput(args: GetAssetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssetResult> {
    return pulumi.output(args).apply(a => getAsset(a, opts))
}

export interface GetAssetOutputArgs {
    /**
     * The unique identifier for the Asset.
     */
    id: pulumi.Input<string>;
}