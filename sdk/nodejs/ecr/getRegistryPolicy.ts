// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * The AWS::ECR::RegistryPolicy is used to specify permissions for another AWS account and is used when configuring cross-account replication. For more information, see Registry permissions in the Amazon Elastic Container Registry User Guide: https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html
 */
export function getRegistryPolicy(args: GetRegistryPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetRegistryPolicyResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ecr:getRegistryPolicy", {
        "registryId": args.registryId,
    }, opts);
}

export interface GetRegistryPolicyArgs {
    registryId: string;
}

export interface GetRegistryPolicyResult {
    /**
     * The JSON policy text to apply to your registry. The policy text follows the same format as IAM policy text. For more information, see Registry permissions (https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the Amazon Elastic Container Registry User Guide.
     */
    readonly policyText?: any;
    readonly registryId?: string;
}

export function getRegistryPolicyOutput(args: GetRegistryPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRegistryPolicyResult> {
    return pulumi.output(args).apply(a => getRegistryPolicy(a, opts))
}

export interface GetRegistryPolicyOutputArgs {
    registryId: pulumi.Input<string>;
}
