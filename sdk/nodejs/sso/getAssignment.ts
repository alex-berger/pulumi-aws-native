// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for SSO assignmet
 */
export function getAssignment(args: GetAssignmentArgs, opts?: pulumi.InvokeOptions): Promise<GetAssignmentResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:sso:getAssignment", {
        "instanceArn": args.instanceArn,
        "permissionSetArn": args.permissionSetArn,
        "principalId": args.principalId,
        "principalType": args.principalType,
        "targetId": args.targetId,
        "targetType": args.targetType,
    }, opts);
}

export interface GetAssignmentArgs {
    /**
     * The sso instance that the permission set is owned.
     */
    instanceArn: string;
    /**
     * The permission set that the assignemt will be assigned
     */
    permissionSetArn: string;
    /**
     * The assignee's identifier, user id/group id
     */
    principalId: string;
    /**
     * The assignee's type, user/group
     */
    principalType: enums.sso.AssignmentPrincipalType;
    /**
     * The account id to be provisioned.
     */
    targetId: string;
    /**
     * The type of resource to be provsioned to, only aws account now
     */
    targetType: enums.sso.AssignmentTargetType;
}

export interface GetAssignmentResult {
}

export function getAssignmentOutput(args: GetAssignmentOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAssignmentResult> {
    return pulumi.output(args).apply(a => getAssignment(a, opts))
}

export interface GetAssignmentOutputArgs {
    /**
     * The sso instance that the permission set is owned.
     */
    instanceArn: pulumi.Input<string>;
    /**
     * The permission set that the assignemt will be assigned
     */
    permissionSetArn: pulumi.Input<string>;
    /**
     * The assignee's identifier, user id/group id
     */
    principalId: pulumi.Input<string>;
    /**
     * The assignee's type, user/group
     */
    principalType: pulumi.Input<enums.sso.AssignmentPrincipalType>;
    /**
     * The account id to be provisioned.
     */
    targetId: pulumi.Input<string>;
    /**
     * The type of resource to be provsioned to, only aws account now
     */
    targetType: pulumi.Input<enums.sso.AssignmentTargetType>;
}
