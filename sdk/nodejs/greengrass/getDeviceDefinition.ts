// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Greengrass::DeviceDefinition
 */
export function getDeviceDefinition(args: GetDeviceDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetDeviceDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:greengrass:getDeviceDefinition", {
        "id": args.id,
    }, opts);
}

export interface GetDeviceDefinitionArgs {
    id: string;
}

export interface GetDeviceDefinitionResult {
    readonly arn?: string;
    readonly id?: string;
    readonly latestVersionArn?: string;
    readonly name?: string;
    readonly tags?: any;
}

export function getDeviceDefinitionOutput(args: GetDeviceDefinitionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDeviceDefinitionResult> {
    return pulumi.output(args).apply(a => getDeviceDefinition(a, opts))
}

export interface GetDeviceDefinitionOutputArgs {
    id: pulumi.Input<string>;
}
