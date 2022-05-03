// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Create and manage FUOTA tasks.
 */
export function getFuotaTask(args: GetFuotaTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetFuotaTaskResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:iotwireless:getFuotaTask", {
        "id": args.id,
    }, opts);
}

export interface GetFuotaTaskArgs {
    /**
     * FUOTA task id. Returned after successful create.
     */
    id: string;
}

export interface GetFuotaTaskResult {
    /**
     * FUOTA task arn. Returned after successful create.
     */
    readonly arn?: string;
    /**
     * Multicast group to associate. Only for update request.
     */
    readonly associateMulticastGroup?: string;
    /**
     * Wireless device to associate. Only for update request.
     */
    readonly associateWirelessDevice?: string;
    /**
     * FUOTA task description
     */
    readonly description?: string;
    /**
     * Multicast group to disassociate. Only for update request.
     */
    readonly disassociateMulticastGroup?: string;
    /**
     * Wireless device to disassociate. Only for update request.
     */
    readonly disassociateWirelessDevice?: string;
    /**
     * FUOTA task firmware update image binary S3 link
     */
    readonly firmwareUpdateImage?: string;
    /**
     * FUOTA task firmware IAM role for reading S3
     */
    readonly firmwareUpdateRole?: string;
    /**
     * FUOTA task status. Returned after successful read.
     */
    readonly fuotaTaskStatus?: string;
    /**
     * FUOTA task id. Returned after successful create.
     */
    readonly id?: string;
    /**
     * FUOTA task LoRaWAN
     */
    readonly loRaWAN?: outputs.iotwireless.FuotaTaskLoRaWAN;
    /**
     * Name of FUOTA task
     */
    readonly name?: string;
    /**
     * A list of key-value pairs that contain metadata for the FUOTA task.
     */
    readonly tags?: outputs.iotwireless.FuotaTaskTag[];
}

export function getFuotaTaskOutput(args: GetFuotaTaskOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFuotaTaskResult> {
    return pulumi.output(args).apply(a => getFuotaTask(a, opts))
}

export interface GetFuotaTaskOutputArgs {
    /**
     * FUOTA task id. Returned after successful create.
     */
    id: pulumi.Input<string>;
}
