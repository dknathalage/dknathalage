local services = ['service'];

local workflow = {
  on: {
    push: {
      branches: ['main'],
    },
  },
  permissions: {
    contents: 'write',
    'id-token': 'write',
  },
  jobs: {
    test: {
      'runs-on': 'ubuntu-latest',
      steps: [
        {
          uses: 'actions/checkout@v4',
        },
        {
          uses: 'actions/setup-go@v2',
          with: {
            'go-version': '1.23',
          },
        },
        {
          run: 'go mod download',
        },
        {
          run: 'go test ./...',
        },
        {
          uses: 'gwatts/go-coverage-action@v2',
          with: { 'coverage-threshold': 80 },
        },
      ],
    },
    'docker-build':
      {
        needs: ['test'],
        'runs-on': 'ubuntu-latest',
        strategy: {
          matrix: {
            service: services,
          },
        },
        steps: [
          {
            uses: 'actions/checkout@v4',
          },
          {
            id: 'auth',
            uses: 'google-github-actions/auth@v2',
            with: {
              token_format: 'access_token',
              workload_identity_provider: 'projects/719430876063/locations/global/workloadIdentityPools/dknathalage-identity-pool/providers/dknathalage-id-provider',
              service_account: 'gha-ci-sa@dknathalage.iam.gserviceaccount.com',
              access_token_lifetime: '300s',
            },
          },
          {
            uses: 'docker/login-action@v1',
            with: {
              registry: 'australia-southeast2-docker.pkg.dev',
              username: 'oauth2accesstoken',
              password: '${{ steps.auth.outputs.access_token }}',
            },
          },
          {
            uses: 'docker/build-push-action@v2',
            with: {
              push: true,
              tags: 'australia-southeast2-docker.pkg.dev/dknathalage/dknathalage/${{matrix.service}}:latest',
              'build-args': 'APP_NAME=${{ matrix.service }}',
            },
          },
          {
            uses: 'aquasecurity/trivy-action@0.20.0',
            id: 'trivy',
            with: {
              'image-ref': 'australia-southeast2-docker.pkg.dev/dknathalage/dknathalage/${{matrix.service}}:latest',
              format: 'table',
              'exit-code': '1',
              'ignore-unfixed': false,
              'vuln-type': 'os,library',
              output: 'results.txt',
              severity: 'LOW,MEDIUM,HIGH,CRITICAL',
            },
            env: {
              TRIVY_DB_REPOSITORY: 'public.ecr.aws/aquasecurity/trivy-db',
            },
          },
          {
            run: 'cat results.txt >> $GITHUB_STEP_SUMMARY',
          },
        ],
      },
  },
};

workflow
