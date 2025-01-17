// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GuardDuty::Master
 */
export function getMaster(args: GetMasterArgs, opts?: pulumi.InvokeOptions): Promise<GetMasterResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:guardduty:getMaster", {
        "masterId": args.masterId,
    }, opts);
}

export interface GetMasterArgs {
    masterId: string;
}

export interface GetMasterResult {
}

export function getMasterOutput(args: GetMasterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetMasterResult> {
    return pulumi.output(args).apply(a => getMaster(a, opts))
}

export interface GetMasterOutputArgs {
    masterId: pulumi.Input<string>;
}
