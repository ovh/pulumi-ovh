FROM gitpod/workspace-full:latest

RUN go install github.com/pulumi/upgrade-provider@main

RUN brew install gh

RUN brew install pulumi/tap/pulumi