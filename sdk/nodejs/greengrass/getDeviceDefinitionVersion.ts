// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::DeviceDefinitionVersion
 */
export function getDeviceDefinitionVersion(args: GetDeviceDefinitionVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceDefinitionVersionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:greengrass:getDeviceDefinitionVersion", {
        "id": args.id,
    }, opts);
}

export interface GetDeviceDefinitionVersionArgs {
    id: string;
}

export interface GetDeviceDefinitionVersionResult {
    readonly id?: string;
}

export function getDeviceDefinitionVersionOutput(args: GetDeviceDefinitionVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceDefinitionVersionResult> {
    return pulumi.output(args).apply(a => getDeviceDefinitionVersion(a, opts))
}

export interface GetDeviceDefinitionVersionOutputArgs {
    id: pulumi.Input<string>;
}