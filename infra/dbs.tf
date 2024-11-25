resource "google_firestore_database" "invoice-db" {
  for_each    = toset(["dev", "prod"])
  name        = "invoice-db-${each.key}"
  location_id = "australia-southeast2"
  type        = "FIRESTORE_NATIVE"
}
