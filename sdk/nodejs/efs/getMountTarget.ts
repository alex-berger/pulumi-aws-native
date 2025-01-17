// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EFS::MountTarget
 */
export function getMountTarget(args: GetMountTargetArgs, opts?: pulumi.InvokeOptions): Promise<GetMountTargetResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:efs:getMountTarget", {
        "id": args.id,
    }, opts);
}

export interface GetMountTargetArgs {
    id: string;
}

export interface GetMountTargetResult {
    readonly id?: string;
    readonly securityGroups?: string[];
}

export function getMountTargetOutput(args: GetMountTargetOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMountTargetResult> {
    return pulumi.output(args).apply(a => getMountTarget(a, opts))
}

export interface GetMountTargetOutputArgs {
    id: pulumi.Input<string>;
}
