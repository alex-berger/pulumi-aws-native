// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html
 */
export class Dashboard extends pulumi.CustomResource {
    /**
     * Get an existing Dashboard resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Dashboard {
        return new Dashboard(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:quicksight:Dashboard';

    /**
     * Returns true if the given object is an instance of Dashboard.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Dashboard {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Dashboard.__pulumiType;
    }

    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-awsaccountid
     */
    public readonly awsAccountId!: pulumi.Output<string>;
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-dashboardid
     */
    public readonly dashboardId!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-dashboardpublishoptions
     */
    public readonly dashboardPublishOptions!: pulumi.Output<outputs.quicksight.DashboardDashboardPublishOptions | undefined>;
    public /*out*/ readonly lastPublishedTime!: pulumi.Output<string>;
    public /*out*/ readonly lastUpdatedTime!: pulumi.Output<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-name
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-parameters
     */
    public readonly parameters!: pulumi.Output<outputs.quicksight.DashboardParameters | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-permissions
     */
    public readonly permissions!: pulumi.Output<outputs.quicksight.DashboardResourcePermission[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-sourceentity
     */
    public readonly sourceEntity!: pulumi.Output<outputs.quicksight.DashboardDashboardSourceEntity | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-tags
     */
    public readonly tags!: pulumi.Output<outputs.Tag[] | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-themearn
     */
    public readonly themeArn!: pulumi.Output<string | undefined>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-versiondescription
     */
    public readonly versionDescription!: pulumi.Output<string | undefined>;

    /**
     * Create a Dashboard resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.awsAccountId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'awsAccountId'");
            }
            if ((!args || args.dashboardId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dashboardId'");
            }
            inputs["awsAccountId"] = args ? args.awsAccountId : undefined;
            inputs["dashboardId"] = args ? args.dashboardId : undefined;
            inputs["dashboardPublishOptions"] = args ? args.dashboardPublishOptions : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["permissions"] = args ? args.permissions : undefined;
            inputs["sourceEntity"] = args ? args.sourceEntity : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["themeArn"] = args ? args.themeArn : undefined;
            inputs["versionDescription"] = args ? args.versionDescription : undefined;
            inputs["arn"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["lastPublishedTime"] = undefined /*out*/;
            inputs["lastUpdatedTime"] = undefined /*out*/;
        } else {
            inputs["arn"] = undefined /*out*/;
            inputs["awsAccountId"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["dashboardId"] = undefined /*out*/;
            inputs["dashboardPublishOptions"] = undefined /*out*/;
            inputs["lastPublishedTime"] = undefined /*out*/;
            inputs["lastUpdatedTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["permissions"] = undefined /*out*/;
            inputs["sourceEntity"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["themeArn"] = undefined /*out*/;
            inputs["versionDescription"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Dashboard.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Dashboard resource.
 */
export interface DashboardArgs {
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-awsaccountid
     */
    awsAccountId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-dashboardid
     */
    dashboardId: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-dashboardpublishoptions
     */
    dashboardPublishOptions?: pulumi.Input<inputs.quicksight.DashboardDashboardPublishOptionsArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-name
     */
    name?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-parameters
     */
    parameters?: pulumi.Input<inputs.quicksight.DashboardParametersArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-permissions
     */
    permissions?: pulumi.Input<pulumi.Input<inputs.quicksight.DashboardResourcePermissionArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-sourceentity
     */
    sourceEntity?: pulumi.Input<inputs.quicksight.DashboardDashboardSourceEntityArgs>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-tags
     */
    tags?: pulumi.Input<pulumi.Input<inputs.TagArgs>[]>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-themearn
     */
    themeArn?: pulumi.Input<string>;
    /**
     * http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-quicksight-dashboard.html#cfn-quicksight-dashboard-versiondescription
     */
    versionDescription?: pulumi.Input<string>;
}