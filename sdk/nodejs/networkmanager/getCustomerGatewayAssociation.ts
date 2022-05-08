// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::NetworkManager::CustomerGatewayAssociation type associates a customer gateway with a device and optionally, with a link.
 */
export function getCustomerGatewayAssociation(args: GetCustomerGatewayAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetCustomerGatewayAssociationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:networkmanager:getCustomerGatewayAssociation", {
        "customerGatewayArn": args.customerGatewayArn,
        "globalNetworkId": args.globalNetworkId,
    }, opts);
}

export interface GetCustomerGatewayAssociationArgs {
    /**
     * The Amazon Resource Name (ARN) of the customer gateway.
     */
    customerGatewayArn: string;
    /**
     * The ID of the global network.
     */
    globalNetworkId: string;
}

export interface GetCustomerGatewayAssociationResult {
}

export function getCustomerGatewayAssociationOutput(args: GetCustomerGatewayAssociationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetCustomerGatewayAssociationResult> {
    return pulumi.output(args).apply(a => getCustomerGatewayAssociation(a, opts))
}

export interface GetCustomerGatewayAssociationOutputArgs {
    /**
     * The Amazon Resource Name (ARN) of the customer gateway.
     */
    customerGatewayArn: pulumi.Input<string>;
    /**
     * The ID of the global network.
     */
    globalNetworkId: pulumi.Input<string>;
}
