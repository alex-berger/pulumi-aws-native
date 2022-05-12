// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource schema for AWS::MediaConnect::FlowVpcInterface
 */
export function getFlowVpcInterface(args: GetFlowVpcInterfaceArgs, opts?: pulumi.InvokeOptions): Promise<GetFlowVpcInterfaceResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:mediaconnect:getFlowVpcInterface", {
        "flowArn": args.flowArn,
        "name": args.name,
    }, opts);
}

export interface GetFlowVpcInterfaceArgs {
    /**
     * The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
     */
    flowArn: string;
    /**
     * Immutable and has to be a unique against other VpcInterfaces in this Flow.
     */
    name: string;
}

export interface GetFlowVpcInterfaceResult {
    /**
     * IDs of the network interfaces created in customer's account by MediaConnect.
     */
    readonly networkInterfaceIds?: string[];
    /**
     * Role Arn MediaConnect can assumes to create ENIs in customer's account.
     */
    readonly roleArn?: string;
    /**
     * Security Group IDs to be used on ENI.
     */
    readonly securityGroupIds?: string[];
    /**
     * Subnet must be in the AZ of the Flow
     */
    readonly subnetId?: string;
}

export function getFlowVpcInterfaceOutput(args: GetFlowVpcInterfaceOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetFlowVpcInterfaceResult> {
    return pulumi.output(args).apply(a => getFlowVpcInterface(a, opts))
}

export interface GetFlowVpcInterfaceOutputArgs {
    /**
     * The Amazon Resource Name (ARN), a unique identifier for any AWS resource, of the flow.
     */
    flowArn: pulumi.Input<string>;
    /**
     * Immutable and has to be a unique against other VpcInterfaces in this Flow.
     */
    name: pulumi.Input<string>;
}
