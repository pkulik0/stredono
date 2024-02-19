import sys
import subprocess
import json
from __workspace import INFRA_PATH, CONFIG_PATH

input_branch = sys.argv[1]

with open(CONFIG_PATH) as f:
    workspaces = json.load(f)["workspaces"]

branches = {}

for workspace in workspaces:
    name = workspace["name"]
    for branch in workspace["branches"]:
        if branch not in branches:
            branches[branch] = name
        else:
            print(f"Invalid configuration: branch {branch} is defined in multiple workspaces")
            sys.exit(1)

if input_branch not in branches:
    print("Branch not found in workspaces configuration")
    sys.exit(1)

print(f"Selecting workspace {branches[input_branch]}")

result = subprocess.run(["terraform", "workspace", "select", branches[input_branch]], cwd=INFRA_PATH)
if result.returncode != 0:
    print("Error selecting workspace")
    sys.exit(1)