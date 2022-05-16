// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::EC2::InternetGateway
 */
export function getInternetGateway(args: GetInternetGatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetInternetGatewayResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getInternetGateway", {
        "internetGatewayId": args.internetGatewayId,
    }, opts);
}

export interface GetInternetGatewayArgs {
    /**
     * ID of internet gateway.
     */
    internetGatewayId: string;
}

export interface GetInternetGatewayResult {
    /**
     * ID of internet gateway.
     */
    readonly internetGatewayId?: string;
    /**
     * Any tags to assign to the internet gateway.
     */
    readonly tags?: outputs.ec2.InternetGatewayTag[];
}

export function getInternetGatewayOutput(args: GetInternetGatewayOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetInternetGatewayResult> {
    return pulumi.output(args).apply(a => getInternetGateway(a, opts))
}

export interface GetInternetGatewayOutputArgs {
    /**
     * ID of internet gateway.
     */
    internetGatewayId: pulumi.Input<string>;
}
