// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::AppFlow::Flow.
 */
export function getFlow(args: GetFlowArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:appflow:getFlow", {
        "flowName": args.flowName,
    }, opts);
}

export interface GetFlowArgs {
    /**
     * Name of the flow.
     */
    flowName: string;
}

export interface GetFlowResult {
    /**
     * Description of the flow.
     */
    readonly description?: string;
    /**
     * List of Destination connectors of the flow.
     */
    readonly destinationFlowConfigList?: outputs.appflow.FlowDestinationFlowConfig[];
    /**
     * ARN identifier of the flow.
     */
    readonly flowArn?: string;
    /**
     * Configurations of Source connector of the flow.
     */
    readonly sourceFlowConfig?: outputs.appflow.FlowSourceFlowConfig;
    /**
     * List of Tags.
     */
    readonly tags?: outputs.appflow.FlowTag[];
    /**
     * List of tasks for the flow.
     */
    readonly tasks?: outputs.appflow.FlowTask[];
    /**
     * Trigger settings of the flow.
     */
    readonly triggerConfig?: outputs.appflow.FlowTriggerConfig;
}

export function getFlowOutput(args: GetFlowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlowResult> {
    return pulumi.output(args).apply(a => getFlow(a, opts))
}

export interface GetFlowOutputArgs {
    /**
     * Name of the flow.
     */
    flowName: pulumi.Input<string>;
}
