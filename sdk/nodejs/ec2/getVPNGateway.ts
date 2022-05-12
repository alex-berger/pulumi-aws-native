// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::VPNGateway
 */
export function getVPNGateway(args: GetVPNGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetVPNGatewayResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getVPNGateway", {
        "id": args.id,
    }, opts);
}

export interface GetVPNGatewayArgs {
    id: string;
}

export interface GetVPNGatewayResult {
    readonly id?: string;
    readonly tags?: outputs.ec2.VPNGatewayTag[];
}

export function getVPNGatewayOutput(args: GetVPNGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVPNGatewayResult> {
    return pulumi.output(args).apply(a => getVPNGateway(a, opts))
}

export interface GetVPNGatewayOutputArgs {
    id: pulumi.Input<string>;
}
