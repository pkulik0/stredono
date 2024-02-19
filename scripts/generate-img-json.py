from PIL import Image
import json
import os
import sys

ALLOWED_EXT = (".webp", ".png", ".jpg", ".jpeg")


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

        with Image.open(f"{img_path}/{file}") as img:
            img_dir_name.append({"src": f"/emotes/{file}", "width": img.size[0], "height": img.size[1]})
            print(f"Found {file} with dimensions {img.size}")

    with open(output_file, "w") as f:
        json.dump(img_dir_name, f, indent=4)

    print(f"Generated {output_file}")


if __name__ == "__main__":
    main()