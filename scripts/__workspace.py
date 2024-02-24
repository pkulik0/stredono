import os
import sys

OUTPUT_NAME = "terraform_output"
CONFIG_PATH = "workspaces.yml"
INFRA_PATH = "infra"
GO_PACKAGE = "platform"
GO_FILE = f"cloud/platform/{OUTPUT_NAME}.go"
JSON_FILE = f"app/src/lib/{OUTPUT_NAME}.json"
