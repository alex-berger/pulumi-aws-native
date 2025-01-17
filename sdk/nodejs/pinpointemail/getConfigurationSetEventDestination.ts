// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::PinpointEmail::ConfigurationSetEventDestination
 */
export function getConfigurationSetEventDestination(args: GetConfigurationSetEventDestinationArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationSetEventDestinationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:pinpointemail:getConfigurationSetEventDestination", {
        "id": args.id,
    }, opts);
}

export interface GetConfigurationSetEventDestinationArgs {
    id: string;
}

export interface GetConfigurationSetEventDestinationResult {
    readonly eventDestination?: outputs.pinpointemail.ConfigurationSetEventDestinationEventDestination;
    readonly id?: string;
}

export function getConfigurationSetEventDestinationOutput(args: GetConfigurationSetEventDestinationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConfigurationSetEventDestinationResult> {
    return pulumi.output(args).apply(a => getConfigurationSetEventDestination(a, opts))
}

export interface GetConfigurationSetEventDestinationOutputArgs {
    id: pulumi.Input<string>;
}
