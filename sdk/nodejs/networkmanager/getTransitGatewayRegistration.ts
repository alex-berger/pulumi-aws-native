// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::NetworkManager::TransitGatewayRegistration type registers a transit gateway in your global network. The transit gateway can be in any AWS Region, but it must be owned by the same AWS account that owns the global network. You cannot register a transit gateway in more than one global network.
 */
export function getTransitGatewayRegistration(args: GetTransitGatewayRegistrationArgs, opts?: pulumi.InvokeOptions): Promise<GetTransitGatewayRegistrationResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:networkmanager:getTransitGatewayRegistration", {
        "globalNetworkId": args.globalNetworkId,
        "transitGatewayArn": args.transitGatewayArn,
    }, opts);
}

export interface GetTransitGatewayRegistrationArgs {
    /**
     * The ID of the global network.
     */
    globalNetworkId: string;
    /**
     * The Amazon Resource Name (ARN) of the transit gateway.
     */
    transitGatewayArn: string;
}

export interface GetTransitGatewayRegistrationResult {
}

export function getTransitGatewayRegistrationOutput(args: GetTransitGatewayRegistrationOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTransitGatewayRegistrationResult> {
    return pulumi.output(args).apply(a => getTransitGatewayRegistration(a, opts))
}

export interface GetTransitGatewayRegistrationOutputArgs {
    /**
     * The ID of the global network.
     */
    globalNetworkId: pulumi.Input<string>;
    /**
     * The Amazon Resource Name (ARN) of the transit gateway.
     */
    transitGatewayArn: pulumi.Input<string>;
}
