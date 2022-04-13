// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::Config::OrganizationConformancePack.
 */
export function getOrganizationConformancePack(args: GetOrganizationConformancePackArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationConformancePackResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:configuration:getOrganizationConformancePack", {
        "organizationConformancePackName": args.organizationConformancePackName,
    }, opts);
}

export interface GetOrganizationConformancePackArgs {
    /**
     * The name of the organization conformance pack.
     */
    organizationConformancePackName: string;
}

export interface GetOrganizationConformancePackResult {
    /**
     * A list of ConformancePackInputParameter objects.
     */
    readonly conformancePackInputParameters?: outputs.configuration.OrganizationConformancePackConformancePackInputParameter[];
    /**
     * AWS Config stores intermediate files while processing conformance pack template.
     */
    readonly deliveryS3Bucket?: string;
    /**
     * The prefix for the delivery S3 bucket.
     */
    readonly deliveryS3KeyPrefix?: string;
    /**
     * A list of AWS accounts to be excluded from an organization conformance pack while deploying a conformance pack.
     */
    readonly excludedAccounts?: string[];
}

export function getOrganizationConformancePackOutput(args: GetOrganizationConformancePackOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationConformancePackResult> {
    return pulumi.output(args).apply(a => getOrganizationConformancePack(a, opts))
}

export interface GetOrganizationConformancePackOutputArgs {
    /**
     * The name of the organization conformance pack.
     */
    organizationConformancePackName: pulumi.Input<string>;
}
