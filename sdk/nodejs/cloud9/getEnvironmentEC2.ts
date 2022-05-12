// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Cloud9::EnvironmentEC2
 */
export function getEnvironmentEC2(args: GetEnvironmentEC2Args, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentEC2Result> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:cloud9:getEnvironmentEC2", {
        "id": args.id,
    }, opts);
}

export interface GetEnvironmentEC2Args {
    id: string;
}

export interface GetEnvironmentEC2Result {
    readonly arn?: string;
    readonly description?: string;
    readonly id?: string;
    readonly name?: string;
    readonly tags?: outputs.cloud9.EnvironmentEC2Tag[];
}

export function getEnvironmentEC2Output(args: GetEnvironmentEC2OutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentEC2Result> {
    return pulumi.output(args).apply(a => getEnvironmentEC2(a, opts))
}

export interface GetEnvironmentEC2OutputArgs {
    id: pulumi.Input<string>;
}
