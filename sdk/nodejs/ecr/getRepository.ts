// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::ECR::Repository resource specifies an Amazon Elastic Container Registry (Amazon ECR) repository, where users can push and pull Docker images. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/Repositories.html
 */
export function getRepository(args: GetRepositoryArgs, opts?: pulumi.InvokeOptions): Promise<GetRepositoryResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws-native:ecr:getRepository", {
        "repositoryName": args.repositoryName,
    }, opts);
}

export interface GetRepositoryArgs {
    /**
     * The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
     */
    repositoryName: string;
}

export interface GetRepositoryResult {
    readonly arn?: string;
    readonly imageScanningConfiguration?: outputs.ecr.RepositoryImageScanningConfiguration;
    /**
     * The image tag mutability setting for the repository.
     */
    readonly imageTagMutability?: enums.ecr.RepositoryImageTagMutability;
    readonly lifecyclePolicy?: outputs.ecr.RepositoryLifecyclePolicy;
    /**
     * The JSON repository policy text to apply to the repository. For more information, see https://docs.aws.amazon.com/AmazonECR/latest/userguide/RepositoryPolicyExamples.html in the Amazon Elastic Container Registry User Guide. 
     */
    readonly repositoryPolicyText?: any;
    readonly repositoryUri?: string;
    /**
     * An array of key-value pairs to apply to this resource.
     */
    readonly tags?: outputs.ecr.RepositoryTag[];
}

export function getRepositoryOutput(args: GetRepositoryOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRepositoryResult> {
    return pulumi.output(args).apply(a => getRepository(a, opts))
}

export interface GetRepositoryOutputArgs {
    /**
     * The name to use for the repository. The repository name may be specified on its own (such as nginx-web-app) or it can be prepended with a namespace to group the repository into a category (such as project-a/nginx-web-app). If you don't specify a name, AWS CloudFormation generates a unique physical ID and uses that ID for the repository name. For more information, see https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html.
     */
    repositoryName: pulumi.Input<string>;
}