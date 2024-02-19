data "external" "public_files" {
  program     = ["python3", "scripts/list-files.py", "public"]
  working_dir = local.base_path
}

resource "google_storage_bucket_object" "public_files" {
  provider = google-beta

  for_each = data.external.public_files.result

  name   = each.value
  source = "${local.base_path}/${each.value}"

  bucket = local.firebase_bucket

  depends_on = [google_firebase_storage_bucket.default, data.external.public_files]
}

locals {
  public_dirs = distinct([for k, v in data.external.public_files.result : split("/", v)[1]])
}

output "frontend_public_files" {
  value = {
    for k, v in local.public_dirs : v => [for file in google_storage_bucket_object.public_files : file.name if split("/", file.name)[1] == v]
  }
}