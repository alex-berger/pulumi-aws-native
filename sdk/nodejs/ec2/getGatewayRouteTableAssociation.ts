// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Associates a gateway with a route table. The gateway and route table must be in the same VPC. This association causes the incoming traffic to the gateway to be routed according to the routes in the route table.
 */
export function getGatewayRouteTableAssociation(args: GetGatewayRouteTableAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetGatewayRouteTableAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ec2:getGatewayRouteTableAssociation", {
        "gatewayId": args.gatewayId,
    }, opts);
}

export interface GetGatewayRouteTableAssociationArgs {
    /**
     * The ID of the gateway.
     */
    gatewayId: string;
}

export interface GetGatewayRouteTableAssociationResult {
    /**
     * The route table association ID.
     */
    readonly associationId?: string;
    /**
     * The ID of the route table.
     */
    readonly routeTableId?: string;
}

export function getGatewayRouteTableAssociationOutput(args: GetGatewayRouteTableAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGatewayRouteTableAssociationResult> {
    return pulumi.output(args).apply(a => getGatewayRouteTableAssociation(a, opts))
}

export interface GetGatewayRouteTableAssociationOutputArgs {
    /**
     * The ID of the gateway.
     */
    gatewayId: pulumi.Input<string>;
}
