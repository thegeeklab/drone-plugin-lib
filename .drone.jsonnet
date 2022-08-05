local PipelineTest(deps=[],) = {
  kind: 'pipeline',
  name: 'test',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'deps',
      image: 'golang:1.19',
      commands: [
        'make deps',
      ],
      volumes: [
        {
          name: 'godeps',
          path: '/go',
        },
      ],
    },
    {
      name: 'lint',
      image: 'golang:1.19',
      commands: [
        'make lint',
      ],
      volumes: [
        {
          name: 'godeps',
          path: '/go',
        },
      ],
    },
    {
      name: 'test',
      image: 'golang:1.19',
      commands: [
        'make test',
      ],
      volumes: [
        {
          name: 'godeps',
          path: '/go',
        },
      ],
    },
  ],
  volumes: [
    {
      name: 'godeps',
      temp: {},
    },
  ],
  depends_on: deps,
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};

local PipelineRelease(deps=[],) = {
  kind: 'pipeline',
  image_pull_secrets: ['docker_config'],
  name: 'release',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'changelog-generate',
      image: 'thegeeklab/git-chglog',
      commands: [
        'git fetch -tq',
        'git-chglog --no-color --no-emoji -o CHANGELOG.md ${DRONE_TAG:---next-tag unreleased unreleased}',
      ],
    },
    {
      name: 'changelog-format',
      image: 'thegeeklab/alpine-tools',
      commands: [
        'prettier CHANGELOG.md',
        'prettier -w CHANGELOG.md',
      ],
    },
    {
      name: 'publish',
      image: 'plugins/github-release',
      settings: {
        overwrite: true,
        api_key: {
          from_secret: 'github_token',
        },
        title: '${DRONE_TAG}',
        note: 'CHANGELOG.md',
      },
      when: {
        ref: [
          'refs/tags/**',
        ],
      },
    },
  ],
  depends_on: deps,
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};

local PipelineDocs(deps=[],) = {
  kind: 'pipeline',
  name: 'docs',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'markdownlint',
      image: 'thegeeklab/markdownlint-cli',
      commands: [
        "markdownlint 'README.md' 'CONTRIBUTING.md'",
      ],
    },
    {
      name: 'spellcheck',
      image: 'thegeeklab/alpine-tools',
      commands: [
        "spellchecker --files 'README.md' 'CONTRIBUTING.md' -d .dictionary -p spell indefinite-article syntax-urls --no-suggestions",
      ],
      environment: {
        FORCE_COLOR: true,
        NPM_CONFIG_LOGLEVEL: 'error',
      },
    },
  ],
  depends_on: deps,
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**', 'refs/pull/**'],
  },
};

local PipelineNotifications(deps=[],) = {
  kind: 'pipeline',
  name: 'notifications',
  platform: {
    os: 'linux',
    arch: 'amd64',
  },
  steps: [
    {
      name: 'matrix',
      image: 'thegeeklab/drone-matrix',
      settings: {
        homeserver: { from_secret: 'matrix_homeserver' },
        roomid: { from_secret: 'matrix_roomid' },
        template: 'Status: **{{ build.Status }}**<br/> Build: [{{ repo.Owner }}/{{ repo.Name }}]({{ build.Link }}){{#if build.Branch}} ({{ build.Branch }}){{/if}} by {{ commit.Author }}<br/> Message: {{ commit.Message.Title }}',
        username: { from_secret: 'matrix_username' },
        password: { from_secret: 'matrix_password' },
      },
      when: {
        status: ['success', 'failure'],
      },
    },
  ],
  depends_on: deps,
  trigger: {
    ref: ['refs/heads/main', 'refs/tags/**'],
    status: ['success', 'failure'],
  },
};

[
  PipelineTest(),
  PipelineRelease(deps=['test'],),
  PipelineDocs(deps=['release']),
  PipelineNotifications(deps=['docs']),
]
