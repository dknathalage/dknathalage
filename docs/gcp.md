https://gist.github.com/palewire/12c4b2b974ef735d22da7493cf7f4d37

```bash
export PROJECT_ID=dknathalage
export SERVICE_ACCOUNT=gha-ci-sa

gcloud iam service-accounts create "${SERVICE_ACCOUNT}" --project "${PROJECT_ID}"
```

```bash
gcloud services enable iamcredentials.googleapis.com --project "${PROJECT_ID}"
```

```bash
export WORKLOAD_IDENTITY_POOL=dknathalage-identity-pool

gcloud iam workload-identity-pools create "${WORKLOAD_IDENTITY_POOL}" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --display-name="${WORKLOAD_IDENTITY_POOL}"
```

```bash
export WORKLOAD_IDENTITY_POOL_ID=$(gcloud iam workload-identity-pools describe "${WORKLOAD_IDENTITY_POOL}" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --format="value(name)")
```

```bash
export WORKLOAD_PROVIDER=dknathalage-id-provider

gcloud iam workload-identity-pools providers create-oidc "${WORKLOAD_PROVIDER}" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --workload-identity-pool="${WORKLOAD_IDENTITY_POOL}" \
  --display-name="${WORKLOAD_PROVIDER}" \
  --attribute-mapping="google.subject=assertion.sub,attribute.actor=assertion.actor,attribute.repository=assertion.repository" \
  --issuer-uri="https://token.actions.githubusercontent.com" \
  --attribute-condition="assertion.repository_owner=='dknathalage'"
```

```bash
export REPO=dknathalage/dknathalage

gcloud iam service-accounts add-iam-policy-binding "${SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
  --project="${PROJECT_ID}" \
  --role="roles/iam.workloadIdentityUser" \
  --member="principalSet://iam.googleapis.com/${WORKLOAD_IDENTITY_POOL_ID}/attribute.repository/${REPO}"
```

```bash
gcloud iam workload-identity-pools providers describe "${WORKLOAD_PROVIDER}" \
  --project="${PROJECT_ID}" \
  --location="global" \
  --workload-identity-pool="${WORKLOAD_IDENTITY_POOL}" \
  --format="value(name)"
```

```bash
gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member="serviceAccount:${SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com" \
    --role="roles/artifactregistry.admin"
```

```bash
gcloud projects get-iam-policy $PROJECT_ID \
    --flatten="bindings[].members" \
    --format='table(bindings.role)' \
    --filter="bindings.members:${SERVICE_ACCOUNT}@${PROJECT_ID}.iam.gserviceaccount.com"
```
