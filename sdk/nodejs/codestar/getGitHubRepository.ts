// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::CodeStar::GitHubRepository
 */
export function getGitHubRepository(args: GetGitHubRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetGitHubRepositoryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:codestar:getGitHubRepository", {
        "id": args.id,
    }, opts);
}

export interface GetGitHubRepositoryArgs {
    id: string;
}

export interface GetGitHubRepositoryResult {
    readonly code?: outputs.codestar.GitHubRepositoryCode;
    readonly connectionArn?: string;
    readonly enableIssues?: boolean;
    readonly id?: string;
    readonly isPrivate?: boolean;
    readonly repositoryAccessToken?: string;
    readonly repositoryDescription?: string;
    readonly repositoryName?: string;
    readonly repositoryOwner?: string;
}

export function getGitHubRepositoryOutput(args: GetGitHubRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetGitHubRepositoryResult> {
    return pulumi.output(args).apply(a => getGitHubRepository(a, opts))
}

export interface GetGitHubRepositoryOutputArgs {
    id: pulumi.Input<string>;
}
