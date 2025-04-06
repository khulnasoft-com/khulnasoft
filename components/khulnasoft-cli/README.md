# Khulnasoft CLI

A CLI binary that comes pre-installed within the Khulnasoft workspace environments.

## Description

Here are a few of the things you can do with it:

- Open a file in the IDE
- Open a URL in the IDE's preview
- Watch the progress of a running task defined on `.khulnasoft.yml` by attaching the terminal to its process.
- Control user-defined environment variables
- Stop the current workspace
- Notify and wait for events to happen, to control the order of execution of terminal tasks
- Make a port available on 0.0.0.0 so that it can be exposed to the Internet
- Wait for a process to listen on a port
- Print the URL of the current workspace
- Take a snapshot of the current workspace
- Create a Khulnasoft configuration for the current project

Learn more about it by running `gp —-help` or checking the [documentation](https://www.khulnasoft.com/docs/command-line-interface/).

## Useful Links

- [Usage Docs](https://www.khulnasoft.com/docs/command-line-interface)
- [GitHub Issues](https://github.com/khulnasoft-com/khulnasoft/labels/component%3A%20gp%20cli)

## Contributing

- The Khulnasoft CLI is written in [Go](https://go.dev/).
- Most of the functionalities that the CLI implements, make use of the [Supervisor API](https://github.com/khulnasoft-com/khulnasoft/tree/main/components/supervisor-api).

If you would like to contribute to this component, check the [related GitHub issues](https://github.com/khulnasoft-com/khulnasoft/labels/component%3A%20gp%20cli) or start a discussion in the Discord [#contributing](https://discord.com/channels/816244985187008514/885406100436951080) channel.

## Ownership

To know which Khulnasoft Team owns this component, check the [CODEOWNERS](https://github.com/khulnasoft-com/khulnasoft/blob/main/.github/CODEOWNERS).
