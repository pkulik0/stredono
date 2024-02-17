data "external" "public_files" {
  program = ["python3", "infra/scripts/list_files.py", "public"]
  working_dir = local.base_path
}

resource "google_storage_bucket_object" "public_files" {
  provider = google-beta

  for_each = data.external.public_files.result

  name = each.value
  source = "${local.base_path}/${each.value}"

  bucket = local.firebase_bucket

  depends_on = [google_firebase_storage_bucket.default, data.external.public_files]
}