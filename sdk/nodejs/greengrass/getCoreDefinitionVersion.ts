// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::CoreDefinitionVersion
 */
export function getCoreDefinitionVersion(args: GetCoreDefinitionVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetCoreDefinitionVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:greengrass:getCoreDefinitionVersion", {
        "id": args.id,
    }, opts);
}

export interface GetCoreDefinitionVersionArgs {
    id: string;
}

export interface GetCoreDefinitionVersionResult {
    readonly id?: string;
}

export function getCoreDefinitionVersionOutput(args: GetCoreDefinitionVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCoreDefinitionVersionResult> {
    return pulumi.output(args).apply(a => getCoreDefinitionVersion(a, opts))
}

export interface GetCoreDefinitionVersionOutputArgs {
    id: pulumi.Input<string>;
}
