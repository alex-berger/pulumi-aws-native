// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::GlobalAccelerator::Accelerator
 */
export function getAccelerator(args: GetAcceleratorArgs, opts?: pulumi.InvokeOptions): Promise<GetAcceleratorResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:globalaccelerator:getAccelerator", {
        "acceleratorArn": args.acceleratorArn,
    }, opts);
}

export interface GetAcceleratorArgs {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    acceleratorArn: string;
}

export interface GetAcceleratorResult {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    readonly acceleratorArn?: string;
    /**
     * The Domain Name System (DNS) name that Global Accelerator creates that points to your accelerator's static IP addresses.
     */
    readonly dnsName?: string;
    /**
     * Indicates whether an accelerator is enabled. The value is true or false.
     */
    readonly enabled?: boolean;
    /**
     * IP Address type.
     */
    readonly ipAddressType?: enums.globalaccelerator.AcceleratorIpAddressType;
    /**
     * The IP addresses from BYOIP Prefix pool.
     */
    readonly ipAddresses?: string[];
    /**
     * Name of accelerator.
     */
    readonly name?: string;
    readonly tags?: outputs.globalaccelerator.AcceleratorTag[];
}

export function getAcceleratorOutput(args: GetAcceleratorOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAcceleratorResult> {
    return pulumi.output(args).apply(a => getAccelerator(a, opts))
}

export interface GetAcceleratorOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the accelerator.
     */
    acceleratorArn: pulumi.Input<string>;
}