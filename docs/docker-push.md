```yaml
docker-build:
  runs-on: ubuntu-latest
  strategy:
    matrix:
      service:
        - service

  steps:
    - uses: actions/checkout@v4

    - id: auth
      uses: google-github-actions/auth@v2
      with:
        token_format: access_token
        workload_identity_provider: projects/719430876063/locations/global/workloadIdentityPools/dknathalage-identity-pool/providers/dknathalage-id-provider
        service_account: gha-ci-sa@dknathalage.iam.gserviceaccount.com
        access_token_lifetime: 300s

    - uses: docker/login-action@v1
      with:
        registry: australia-southeast2-docker.pkg.dev
        username: oauth2accesstoken
        password: ${{ steps.auth.outputs.access_token }}

    - id: docker-push-tagged
      uses: docker/build-push-action@v2
      with:
        push: true
        tags: |
          australia-southeast2-docker.pkg.dev/dknathalage/dknathalage/${{ matrix.service }}:latest
          australia-southeast2-docker.pkg.dev/dknathalage/dknathalage/${{ matrix.service }}:${{ github.sha }}
        build-args: |
          APP_NAME=${{ matrix.service }}

image-scans:
  runs-on: ubuntu-latest
  strategy:
    matrix:
      service:
        - service
  needs: docker-build
  steps:
    - uses: actions/checkout@v4
    - uses: ./.github/actions/image-scans
      with:
        image-name: australia-southeast2-docker.pkg.dev/dknathalage/dknathalage/${{ matrix.service }}
        image-tag: ${{ github.sha }}
```
