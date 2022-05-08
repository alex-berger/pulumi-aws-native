// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Resource Type definition for AWS::AppFlow::ConnectorProfile
 */
export class ConnectorProfile extends pulumi.CustomResource {
    /**
     * Get an existing ConnectorProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectorProfile {
        return new ConnectorProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws-native:appflow:ConnectorProfile';

    /**
     * Returns true if the given object is an instance of ConnectorProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectorProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectorProfile.__pulumiType;
    }

    /**
     * Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
     */
    public readonly connectionMode!: pulumi.Output<enums.appflow.ConnectorProfileConnectionMode>;
    /**
     * Unique identifier for connector profile resources
     */
    public /*out*/ readonly connectorProfileArn!: pulumi.Output<string>;
    /**
     * Connector specific configurations needed to create connector profile
     */
    public readonly connectorProfileConfig!: pulumi.Output<outputs.appflow.ConnectorProfileConfig | undefined>;
    /**
     * The maximum number of items to retrieve in a single batch.
     */
    public readonly connectorProfileName!: pulumi.Output<string>;
    /**
     * List of Saas providers that need connector profile to be created
     */
    public readonly connectorType!: pulumi.Output<enums.appflow.ConnectorProfileConnectorType>;
    /**
     * A unique Arn for Connector-Profile resource
     */
    public /*out*/ readonly credentialsArn!: pulumi.Output<string>;
    /**
     * The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
     */
    public readonly kMSArn!: pulumi.Output<string | undefined>;

    /**
     * Create a ConnectorProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectorProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionMode === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionMode'");
            }
            if ((!args || args.connectorType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectorType'");
            }
            resourceInputs["connectionMode"] = args ? args.connectionMode : undefined;
            resourceInputs["connectorProfileConfig"] = args ? args.connectorProfileConfig : undefined;
            resourceInputs["connectorProfileName"] = args ? args.connectorProfileName : undefined;
            resourceInputs["connectorType"] = args ? args.connectorType : undefined;
            resourceInputs["kMSArn"] = args ? args.kMSArn : undefined;
            resourceInputs["connectorProfileArn"] = undefined /*out*/;
            resourceInputs["credentialsArn"] = undefined /*out*/;
        } else {
            resourceInputs["connectionMode"] = undefined /*out*/;
            resourceInputs["connectorProfileArn"] = undefined /*out*/;
            resourceInputs["connectorProfileConfig"] = undefined /*out*/;
            resourceInputs["connectorProfileName"] = undefined /*out*/;
            resourceInputs["connectorType"] = undefined /*out*/;
            resourceInputs["credentialsArn"] = undefined /*out*/;
            resourceInputs["kMSArn"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ConnectorProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectorProfile resource.
 */
export interface ConnectorProfileArgs {
    /**
     * Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
     */
    connectionMode: pulumi.Input<enums.appflow.ConnectorProfileConnectionMode>;
    /**
     * Connector specific configurations needed to create connector profile
     */
    connectorProfileConfig?: pulumi.Input<inputs.appflow.ConnectorProfileConfigArgs>;
    /**
     * The maximum number of items to retrieve in a single batch.
     */
    connectorProfileName?: pulumi.Input<string>;
    /**
     * List of Saas providers that need connector profile to be created
     */
    connectorType: pulumi.Input<enums.appflow.ConnectorProfileConnectorType>;
    /**
     * The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.
     */
    kMSArn?: pulumi.Input<string>;
}
