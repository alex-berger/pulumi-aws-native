// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * The AWS::Amplify::Branch resource creates a new branch within an app.
 */
export class Branch extends pulumi.CustomResource {
    /**
     * Get an existing Branch resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Branch {
        return new Branch(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:amplify:Branch';

    /**
     * Returns true if the given object is an instance of Branch.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Branch {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Branch.__pulumiType;
    }

    public readonly appId!: pulumi.Output<string>;
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly basicAuthConfig!: pulumi.Output<outputs.amplify.BranchBasicAuthConfig | undefined>;
    public readonly branchName!: pulumi.Output<string>;
    public readonly buildSpec!: pulumi.Output<string | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enableAutoBuild!: pulumi.Output<boolean | undefined>;
    public readonly enablePerformanceMode!: pulumi.Output<boolean | undefined>;
    public readonly enablePullRequestPreview!: pulumi.Output<boolean | undefined>;
    public readonly environmentVariables!: pulumi.Output<outputs.amplify.BranchEnvironmentVariable[] | undefined>;
    public readonly pullRequestEnvironmentName!: pulumi.Output<string | undefined>;
    public readonly stage!: pulumi.Output<enums.amplify.BranchStage | undefined>;
    public readonly tags!: pulumi.Output<outputs.amplify.BranchTag[] | undefined>;

    /**
     * Create a Branch resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BranchArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["basicAuthConfig"] = args ? args.basicAuthConfig : undefined;
            resourceInputs["branchName"] = args ? args.branchName : undefined;
            resourceInputs["buildSpec"] = args ? args.buildSpec : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enableAutoBuild"] = args ? args.enableAutoBuild : undefined;
            resourceInputs["enablePerformanceMode"] = args ? args.enablePerformanceMode : undefined;
            resourceInputs["enablePullRequestPreview"] = args ? args.enablePullRequestPreview : undefined;
            resourceInputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            resourceInputs["pullRequestEnvironmentName"] = args ? args.pullRequestEnvironmentName : undefined;
            resourceInputs["stage"] = args ? args.stage : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
        } else {
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["basicAuthConfig"] = undefined /*out*/;
            resourceInputs["branchName"] = undefined /*out*/;
            resourceInputs["buildSpec"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["enableAutoBuild"] = undefined /*out*/;
            resourceInputs["enablePerformanceMode"] = undefined /*out*/;
            resourceInputs["enablePullRequestPreview"] = undefined /*out*/;
            resourceInputs["environmentVariables"] = undefined /*out*/;
            resourceInputs["pullRequestEnvironmentName"] = undefined /*out*/;
            resourceInputs["stage"] = undefined /*out*/;
            resourceInputs["tags"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Branch.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Branch resource.
 */
export interface BranchArgs {
    appId: pulumi.Input<string>;
    basicAuthConfig?: pulumi.Input<inputs.amplify.BranchBasicAuthConfigArgs>;
    branchName?: pulumi.Input<string>;
    buildSpec?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    enableAutoBuild?: pulumi.Input<boolean>;
    enablePerformanceMode?: pulumi.Input<boolean>;
    enablePullRequestPreview?: pulumi.Input<boolean>;
    environmentVariables?: pulumi.Input<pulumi.Input<inputs.amplify.BranchEnvironmentVariableArgs>[]>;
    pullRequestEnvironmentName?: pulumi.Input<string>;
    stage?: pulumi.Input<enums.amplify.BranchStage>;
    tags?: pulumi.Input<pulumi.Input<inputs.amplify.BranchTagArgs>[]>;
}
