// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AwsNative.Lex.Outputs
{

    /// <summary>
    /// A custom vocabulary item that contains the phrase to recognize and a weight to give the boost.
    /// </summary>
    [OutputType]
    public sealed class BotCustomVocabularyItem
    {
        /// <summary>
        /// Phrase that should be recognized.
        /// </summary>
        public readonly string Phrase;
        /// <summary>
        /// The degree to which the phrase recognition is boosted.
        /// </summary>
        public readonly int? Weight;

        [OutputConstructor]
        private BotCustomVocabularyItem(
            string phrase,

            int? weight)
        {
            Phrase = phrase;
            Weight = weight;
        }
    }
}
