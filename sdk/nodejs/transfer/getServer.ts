// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::Transfer::Server
 */
export function getServer(args: GetServerArgs, opts?: pulumi.InvokeOptions): Promise<GetServerResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:transfer:getServer", {
        "serverId": args.serverId,
    }, opts);
}

export interface GetServerArgs {
    serverId: string;
}

export interface GetServerResult {
    readonly arn?: string;
    readonly certificate?: string;
    readonly endpointDetails?: outputs.transfer.ServerEndpointDetails;
    readonly endpointType?: string;
    readonly identityProviderDetails?: outputs.transfer.ServerIdentityProviderDetails;
    readonly loggingRole?: string;
    readonly postAuthenticationLoginBanner?: string;
    readonly preAuthenticationLoginBanner?: string;
    readonly protocolDetails?: outputs.transfer.ServerProtocolDetails;
    readonly protocols?: outputs.transfer.ServerProtocol[];
    readonly securityPolicyName?: string;
    readonly serverId?: string;
    readonly tags?: outputs.transfer.ServerTag[];
    readonly workflowDetails?: outputs.transfer.ServerWorkflowDetails;
}

export function getServerOutput(args: GetServerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetServerResult> {
    return pulumi.output(args).apply(a => getServer(a, opts))
}

export interface GetServerOutputArgs {
    serverId: pulumi.Input<string>;
}
