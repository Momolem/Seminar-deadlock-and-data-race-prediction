FROM mcr.microsoft.com/devcontainers/go

USER root

RUN apt-get update \
    && apt-get install -y --no-install-recommends curl xdg-utils \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /run/user/1000 \
    && chown vscode:vscode /run/user/1000 \
    && chmod 700 /run/user/1000

USER vscode

# Install Go 1.18
RUN go install golang.org/dl/go1.15@latest
RUN go1.15 download

# Install OpenCode
RUN curl -fsSL https://opencode.ai/v2/install | bash

ENV PATH="/home/vscode/.opencode/bin:${PATH}"

# Pre-create OpenCode state directories so the named volumes mounted in
# devcontainer.json inherit vscode ownership on first use.
RUN mkdir -p /home/vscode/.config/opencode \
        /home/vscode/.local/share/opencode \
        /home/vscode/.local/state/opencode

# Pre-download chronos Go dependencies.
# These two files should be copied before the rest of the source so
# Docker can cache this layer.
WORKDIR /tmp/chronos

COPY --chown=vscode:vscode chronos/go.mod chronos/go.sum ./

RUN go1.15 mod download