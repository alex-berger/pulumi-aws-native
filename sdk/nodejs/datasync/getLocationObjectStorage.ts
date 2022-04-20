// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationObjectStorage.
 */
export function getLocationObjectStorage(args: GetLocationObjectStorageArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationObjectStorageResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getLocationObjectStorage", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationObjectStorageArgs {
    /**
     * The Amazon Resource Name (ARN) of the location that is created.
     */
    locationArn: string;
}

export interface GetLocationObjectStorageResult {
    /**
     * Optional. The access key is used if credentials are required to access the self-managed object storage server.
     */
    readonly accessKey?: string;
    /**
     * The Amazon Resource Name (ARN) of the agents associated with the self-managed object storage server location.
     */
    readonly agentArns?: string[];
    /**
     * The Amazon Resource Name (ARN) of the location that is created.
     */
    readonly locationArn?: string;
    /**
     * The URL of the object storage location that was described.
     */
    readonly locationUri?: string;
    /**
     * The port that your self-managed server accepts inbound network traffic on.
     */
    readonly serverPort?: number;
    /**
     * The protocol that the object storage server uses to communicate.
     */
    readonly serverProtocol?: enums.datasync.LocationObjectStorageServerProtocol;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationObjectStorageTag[];
}

export function getLocationObjectStorageOutput(args: GetLocationObjectStorageOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationObjectStorageResult> {
    return pulumi.output(args).apply(a => getLocationObjectStorage(a, opts))
}

export interface GetLocationObjectStorageOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the location that is created.
     */
    locationArn: pulumi.Input<string>;
}
