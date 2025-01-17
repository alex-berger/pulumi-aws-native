// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::SSM::ResourceDataSync
 */
export function getResourceDataSync(args: GetResourceDataSyncArgs, opts?: pulumi.InvokeOptions): Promise<GetResourceDataSyncResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ssm:getResourceDataSync", {
        "syncName": args.syncName,
    }, opts);
}

export interface GetResourceDataSyncArgs {
    syncName: string;
}

export interface GetResourceDataSyncResult {
    readonly syncSource?: outputs.ssm.ResourceDataSyncSyncSource;
}

export function getResourceDataSyncOutput(args: GetResourceDataSyncOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetResourceDataSyncResult> {
    return pulumi.output(args).apply(a => getResourceDataSync(a, opts))
}

export interface GetResourceDataSyncOutputArgs {
    syncName: pulumi.Input<string>;
}
