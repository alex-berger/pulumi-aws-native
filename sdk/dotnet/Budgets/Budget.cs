// *** WARNING: this file was generated by pulumigen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Budgets
{
    /// <summary>
    /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html
    /// </summary>
    [AwsNativeResourceType("aws-native:budgets:Budget")]
    public partial class Budget : Pulumi.CustomResource
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
        /// </summary>
        [Output("budget")]
        public Output<Outputs.BudgetBudgetData> BudgetValue { get; private set; } = null!;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
        /// </summary>
        [Output("notificationsWithSubscribers")]
        public Output<ImmutableArray<Outputs.BudgetNotificationWithSubscribers>> NotificationsWithSubscribers { get; private set; } = null!;


        /// <summary>
        /// Create a Budget resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Budget(string name, BudgetArgs args, CustomResourceOptions? options = null)
            : base("aws-native:budgets:Budget", name, args ?? new BudgetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Budget(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("aws-native:budgets:Budget", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Budget resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Budget Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Budget(name, id, options);
        }
    }

    public sealed class BudgetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-budget
        /// </summary>
        [Input("budget", required: true)]
        public Input<Inputs.BudgetBudgetDataArgs> BudgetValue { get; set; } = null!;

        [Input("notificationsWithSubscribers")]
        private InputList<Inputs.BudgetNotificationWithSubscribersArgs>? _notificationsWithSubscribers;

        /// <summary>
        /// http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-budgets-budget.html#cfn-budgets-budget-notificationswithsubscribers
        /// </summary>
        public InputList<Inputs.BudgetNotificationWithSubscribersArgs> NotificationsWithSubscribers
        {
            get => _notificationsWithSubscribers ?? (_notificationsWithSubscribers = new InputList<Inputs.BudgetNotificationWithSubscribersArgs>());
            set => _notificationsWithSubscribers = value;
        }

        public BudgetArgs()
        {
        }
    }
}
