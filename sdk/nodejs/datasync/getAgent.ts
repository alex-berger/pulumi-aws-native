// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::DataSync::Agent.
 */
export function getAgent(args: GetAgentArgs, opts?: pulumi.InvokeOptions): Promise<GetAgentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:datasync:getAgent", {
        "agentArn": args.agentArn,
    }, opts);
}

export interface GetAgentArgs {
    /**
     * The DataSync Agent ARN.
     */
    agentArn: string;
}

export interface GetAgentResult {
    /**
     * The DataSync Agent ARN.
     */
    readonly agentArn?: string;
    /**
     * The name configured for the agent. Text reference used to identify the agent in the console.
     */
    readonly agentName?: string;
    /**
     * The service endpoints that the agent will connect to.
     */
    readonly endpointType?: enums.datasync.AgentEndpointType;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.datasync.AgentTag[];
}

export function getAgentOutput(args: GetAgentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAgentResult> {
    return pulumi.output(args).apply(a => getAgent(a, opts))
}

export interface GetAgentOutputArgs {
    /**
     * The DataSync Agent ARN.
     */
    agentArn: pulumi.Input<string>;
}
