// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Schema for PackageVersion Resource Type
 */
export function getPackageVersion(args: GetPackageVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetPackageVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:panorama:getPackageVersion", {
        "packageId": args.packageId,
        "packageVersion": args.packageVersion,
        "patchVersion": args.patchVersion,
    }, opts);
}

export interface GetPackageVersionArgs {
    packageId: string;
    packageVersion: string;
    patchVersion: string;
}

export interface GetPackageVersionResult {
    readonly isLatestPatch?: boolean;
    readonly markLatest?: boolean;
    readonly packageArn?: string;
    readonly packageName?: string;
    readonly registeredTime?: number;
    readonly status?: enums.panorama.PackageVersionStatus;
    readonly statusDescription?: string;
    readonly updatedLatestPatchVersion?: string;
}

export function getPackageVersionOutput(args: GetPackageVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetPackageVersionResult> {
    return pulumi.output(args).apply(a => getPackageVersion(a, opts))
}

export interface GetPackageVersionOutputArgs {
    packageId: pulumi.Input<string>;
    packageVersion: pulumi.Input<string>;
    patchVersion: pulumi.Input<string>;
}
