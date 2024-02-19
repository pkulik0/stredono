import subprocess
import json
import os
import sys

ALLOWED_EXT = (".webp", ".png", ".jpg", ".jpeg")


# The bizarre format (%wx%h.) is due to the fact that *sometimes* the output is repeated N times.
# The dot is used to get only the first occurrence.
def get_img_dimensions(file_path) -> tuple[int, ...]:
    output = subprocess.check_output(["identify", "-format", "%wx%h.", file_path])
    output = output.decode("utf-8").split(".")[0]
    return tuple(map(int, output.split("x")))


def main():
    app_path = sys.argv[1]
    img_dir_name = sys.argv[2]
    if not os.path.exists(app_path):
        print(f"Error: app directory \"{app_path}\" does not exist")
        sys.exit(1)

    output_file = f"{app_path}/src/lib/{img_dir_name}.json"
    img_path = f"{app_path}/static/{img_dir_name}"

    if not os.path.exists(img_path):
        print("Error: image directory does not exist")
        sys.exit(1)

    img_dir_name = []

    for file in os.listdir(img_path):
        if not file.endswith(ALLOWED_EXT):
            continue

        size = get_img_dimensions(f"{img_path}/{file}")
        print(f"Found {file} with dimensions {size}")
        img_dir_name.append({"src": f"/emotes/{file}", "width": size[0], "height": size[1]})

    with open(output_file, "w") as f:
        json.dump(img_dir_name, f, indent=4)

    print(f"Generated {output_file}")


if __name__ == "__main__":
    main()