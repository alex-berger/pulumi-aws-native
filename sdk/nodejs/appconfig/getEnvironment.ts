// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppConfig::Environment
 */
export function getEnvironment(args: GetEnvironmentArgs, opts?: pulumi.InvokeOptions): Promise<GetEnvironmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appconfig:getEnvironment", {
        "id": args.id,
    }, opts);
}

export interface GetEnvironmentArgs {
    id: string;
}

export interface GetEnvironmentResult {
    readonly description?: string;
    readonly id?: string;
    readonly monitors?: outputs.appconfig.EnvironmentMonitors[];
    readonly name?: string;
    readonly tags?: outputs.appconfig.EnvironmentTags[];
}

export function getEnvironmentOutput(args: GetEnvironmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEnvironmentResult> {
    return pulumi.output(args).apply(a => getEnvironment(a, opts))
}

export interface GetEnvironmentOutputArgs {
    id: pulumi.Input<string>;
}
