// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::Connection.
 */
export function getConnection(args: GetConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetConnectionResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:events:getConnection", {
        "name": args.name,
    }, opts);
}

export interface GetConnectionArgs {
    /**
     * Name of the connection.
     */
    name: string;
}

export interface GetConnectionResult {
    /**
     * The arn of the connection resource.
     */
    readonly arn?: string;
    readonly authorizationType?: enums.events.ConnectionAuthorizationType;
    /**
     * Description of the connection.
     */
    readonly description?: string;
    /**
     * The arn of the secrets manager secret created in the customer account.
     */
    readonly secretArn?: string;
}

export function getConnectionOutput(args: GetConnectionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetConnectionResult> {
    return pulumi.output(args).apply(a => getConnection(a, opts))
}

export interface GetConnectionOutputArgs {
    /**
     * Name of the connection.
     */
    name: pulumi.Input<string>;
}
