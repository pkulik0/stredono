import os
import sys


def find_dir_by_content_ext(x, location=os.curdir):
    found_dir = None
    for file_top in os.listdir(location):
        if not os.path.isdir(file_top):
            continue
        for file_nested in os.listdir(file_top):
            if not file_nested.endswith(x):
                continue
            if found_dir:
                print(f"Directory with {x} files not unique")
                sys.exit(1)
            found_dir = file_top
            break
    if found_dir:
        return found_dir

    print(f"Directory with {x} files not found")
    sys.exit(1)


OUTPUT_NAME = "terraform_output"
CONFIG_PATH = "workspaces.yml"
INFRA_PATH = find_dir_by_content_ext(".tf")
GO_PACKAGE = find_dir_by_content_ext(".go")
GO_FILE = f"{GO_PACKAGE}/{OUTPUT_NAME}.go"
JSON_FILE = f"app/src/lib/{OUTPUT_NAME}.json"
