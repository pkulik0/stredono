import sys
import subprocess
import yaml
from __workspace import INFRA_PATH, CONFIG_PATH

desired_branch = sys.argv[1]

with open(CONFIG_PATH) as f:
    workspaces = yaml.safe_load(f)["workspaces"]

branches = {}

for workspace in workspaces:
    for branch in workspaces[workspace]["branches"]:
        if branch not in branches:
            branches[branch] = workspace
        else:
            print(f"Invalid configuration: branch {branch} is defined in multiple workspaces")
            sys.exit(1)

if desired_branch not in branches:
    print("Branch not found in workspaces configuration")
    sys.exit(1)

print(f"Selecting workspace {branches[desired_branch]}")

result = subprocess.run(["terraform", "workspace", "select", branches[desired_branch]], cwd=INFRA_PATH)
if result.returncode != 0:
    print("Error selecting workspace")
    sys.exit(1)