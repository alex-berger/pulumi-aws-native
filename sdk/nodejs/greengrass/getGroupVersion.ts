// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::GroupVersion
 */
export function getGroupVersion(args: GetGroupVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetGroupVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:greengrass:getGroupVersion", {
        "id": args.id,
    }, opts);
}

export interface GetGroupVersionArgs {
    id: string;
}

export interface GetGroupVersionResult {
    readonly id?: string;
}

export function getGroupVersionOutput(args: GetGroupVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGroupVersionResult> {
    return pulumi.output(args).apply(a => getGroupVersion(a, opts))
}

export interface GetGroupVersionOutputArgs {
    id: pulumi.Input<string>;
}
