// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::OpsWorks::Volume
 */
export function getVolume(args: GetVolumeArgs, opts?: pulumi.InvokeOptions): Promise<GetVolumeResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:opsworks:getVolume", {
        "id": args.id,
    }, opts);
}

export interface GetVolumeArgs {
    id: string;
}

export interface GetVolumeResult {
    readonly id?: string;
    readonly mountPoint?: string;
    readonly name?: string;
}

export function getVolumeOutput(args: GetVolumeOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVolumeResult> {
    return pulumi.output(args).apply(a => getVolume(a, opts))
}

export interface GetVolumeOutputArgs {
    id: pulumi.Input<string>;
}