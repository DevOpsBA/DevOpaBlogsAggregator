|             Article Data             ||
| ----------------- | ----------------- |
| URL               | https://kubernetes.io/blog/2018/11/08/kubernetes-docs-updates-international-edition/        |
| Tags              | [kubernetes]       |
| Date Create       | 2018-11-08 00:00:00 &#43;0000 UTC |
| Date Parse        | 2021-12-06 10:51:20.4645158 &#43;0300 MSK m=&#43;2.053550101  |

# Kubernetes Docs Updates, International Edition | Kubernetes

	
	
	
	
	**Author**: Zach Corleissen (Linux Foundation)
As a co-chair of SIG Docs, I&#39;m excited to share that Kubernetes docs have a fully mature workflow for localization (l10n).
L10n is an abbreviation for *localization*.
I18n is an abbreviation for *internationalization*.
I18n is [what you do](https://www.w3.org/International/questions/qa-i18n) to make l10n easier. L10n is a fuller, more comprehensive process than translation (*t9n*).
The goal of SIG Docs is to make Kubernetes easier to use for as many people as possible.
One year ago, we looked at whether it was possible to host the output of a Chinese team working independently to translate the Kubernetes docs. After many conversations (including experts on OpenStack l10n), [much transformation](https://kubernetes.io/blog/2018/05/05/hugo-migration/), and [renewed commitment to easier localization](https://github.com/kubernetes/website/pull/10485), we realized that open source documentation is, like open source software, an ongoing exercise at the edges of what&#39;s possible.
Consolidating workflows, language labels, and team-level ownership may seem like simple improvements, but these features make l10n scalable for increasing numbers of l10n teams. While SIG Docs continues to iterate improvements, we&#39;ve paid off a significant amount of technical debt and streamlined l10n in a single workflow. That&#39;s great for the future as well as the present.
Localization is now consolidated in the [kubernetes/website](https://github.com/kubernetes/website) repository. We&#39;ve configured the Kubernetes CI/CD system, [Prow](https://github.com/kubernetes/test-infra/tree/master/prow), to handle automatic language label assignment as well as team-level PR review and approval.
Prow automatically applies language labels based on file path. Thanks to SIG Docs contributor [June Yi](https://github.com/kubernetes/test-infra/pull/9835), folks can also manually assign language labels in pull request (PR) comments. For example, when left as a comment on an issue or PR, this command assigns the label ```language/ko``` (Korean).
```/language ko
```These repo labels let reviewers filter for PRs and issues by language. For example, you can now filter the k/website dashboard for [PRs with Chinese content](https://github.com/kubernetes/website/pulls?utf8=%E2%9C%93&amp;q=is%3Aopen&#43;is%3Apr&#43;label%3Alanguage%2Fzh).
L10n teams can now review and approve their own PRs. For example, review and approval permissions for English are [assigned in an OWNERS file](https://github.com/kubernetes/website/blob/master/content/en/OWNERS) in the top subfolder for English content.
Adding ```OWNERS``` files to subdirectories lets localization teams review and approve changes without requiring a rubber stamp approval from reviewers who may lack fluency.
We&#39;re looking forward to the [doc sprint in Shanghai](https://kccncchina2018english.sched.com/event/HVb2/contributor-summit-doc-sprint-additional-registration-required) to serve as a resource for the Chinese l10n team.
We&#39;re excited to continue supporting the Japanese and Korean l10n teams, who are making excellent progress.
If you&#39;re interested in localizing Kubernetes for your own language or region, check out our [guide to localizing Kubernetes docs](/docs/contribute/localization/) and reach out to a [SIG Docs chair](https://github.com/kubernetes/community/tree/master/sig-docs#leadership) for support.
If you&#39;re interested in Kubernetes documentation, come to a SIG Docs [weekly meeting](https://github.com/kubernetes/community/tree/master/sig-docs#meetings), or join [#sig-docs in Kubernetes Slack](https://kubernetes.slack.com/messages/C1J0BPD2M/details/).


	

	


