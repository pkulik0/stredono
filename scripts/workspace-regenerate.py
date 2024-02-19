import sys
import subprocess
import json
from __workspace import INFRA_PATH, GO_PACKAGE, GO_FILE, JSON_FILE


def get_tf_output() -> str:
    result = subprocess.run(["terraform", "output", "-json"], cwd=INFRA_PATH, capture_output=True, text=True)
    if result.returncode != 0:
        print("Error getting output")
        sys.exit(1)

    return result.stdout


def format_key(key: str) -> str:
    return "".join([part.capitalize() for part in key.split("_")][1:])


def output_to_go(json_dict: dict, package: str) -> str:
    go = f"package {package}\n\n"
    if len(json_dict) == 0:
        return go

    go += "const (\n"

    for key, value in json_dict.items():
        if not key.startswith("backend_"):
            continue

        key = format_key(key)

        is_string = value["type"] == "string"

        val = value["value"]
        if is_string:
            val = f'"{val}"'

        go += f"\t{key} = {val}\n"

    go += ")\n"
    return go


def output_to_json(json_dict: dict) -> dict:
    result = {}
    for key, value in json_dict.items():
        if not key.startswith("frontend_"):
            continue
        result[format_key(key)] = value["value"]

    return result


def main():
    output = json.loads(get_tf_output())

    with open(GO_FILE, "w") as f:
        f.write(output_to_go(output, GO_PACKAGE))
    print(f"Written to {GO_FILE}")

    with open(JSON_FILE, "w") as f:
        f.write(json.dumps(output_to_json(output), indent=2))
    print(f"Written to {JSON_FILE}")


if __name__ == "__main__":
    main()
