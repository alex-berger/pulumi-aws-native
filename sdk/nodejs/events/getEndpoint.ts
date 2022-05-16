// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Events::Endpoint.
 */
export function getEndpoint(args: GetEndpointArgs, opts?: pulumi.InvokeOptions): Promise<GetEndpointResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:events:getEndpoint", {
        "name": args.name,
    }, opts);
}

export interface GetEndpointArgs {
    name: string;
}

export interface GetEndpointResult {
    readonly arn?: string;
    readonly description?: string;
    readonly endpointId?: string;
    readonly endpointUrl?: string;
    readonly eventBuses?: outputs.events.EndpointEventBus[];
    readonly replicationConfig?: outputs.events.EndpointReplicationConfig;
    readonly roleArn?: string;
    readonly routingConfig?: outputs.events.EndpointRoutingConfig;
    readonly state?: enums.events.EndpointState;
    readonly stateReason?: string;
}

export function getEndpointOutput(args: GetEndpointOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetEndpointResult> {
    return pulumi.output(args).apply(a => getEndpoint(a, opts))
}

export interface GetEndpointOutputArgs {
    name: pulumi.Input<string>;
}