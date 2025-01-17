// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::LocationHDFS.
 */
export function getLocationHDFS(args: GetLocationHDFSArgs, opts?: pulumi.InvokeOptions): Promise<GetLocationHDFSResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getLocationHDFS", {
        "locationArn": args.locationArn,
    }, opts);
}

export interface GetLocationHDFSArgs {
    /**
     * The Amazon Resource Name (ARN) of the HDFS location.
     */
    locationArn: string;
}

export interface GetLocationHDFSResult {
    /**
     * ARN(s) of the agent(s) to use for an HDFS location.
     */
    readonly agentArns?: string[];
    /**
     * The authentication mode used to determine identity of user.
     */
    readonly authenticationType?: enums.datasync.LocationHDFSAuthenticationType;
    /**
     * Size of chunks (blocks) in bytes that the data is divided into when stored in the HDFS cluster.
     */
    readonly blockSize?: number;
    /**
     * The unique identity, or principal, to which Kerberos can assign tickets.
     */
    readonly kerberosPrincipal?: string;
    /**
     * The identifier for the Key Management Server where the encryption keys that encrypt data inside HDFS clusters are stored.
     */
    readonly kmsKeyProviderUri?: string;
    /**
     * The Amazon Resource Name (ARN) of the HDFS location.
     */
    readonly locationArn?: string;
    /**
     * The URL of the HDFS location that was described.
     */
    readonly locationUri?: string;
    /**
     * An array of Name Node(s) of the HDFS location.
     */
    readonly nameNodes?: outputs.datasync.LocationHDFSNameNode[];
    readonly qopConfiguration?: outputs.datasync.LocationHDFSQopConfiguration;
    /**
     * Number of copies of each block that exists inside the HDFS cluster.
     */
    readonly replicationFactor?: number;
    /**
     * The user name that has read and write permissions on the specified HDFS cluster.
     */
    readonly simpleUser?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.LocationHDFSTag[];
}

export function getLocationHDFSOutput(args: GetLocationHDFSOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetLocationHDFSResult> {
    return pulumi.output(args).apply(a => getLocationHDFS(a, opts))
}

export interface GetLocationHDFSOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the HDFS location.
     */
    locationArn: pulumi.Input<string>;
}
