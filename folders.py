import os


def print_folder_structure(folder_path, indent=""):
    folder_name = os.path.basename(folder_path)
    if not folder_name.startswith("."):
        print(f"{indent}|-{folder_name}")
        indent += "| "

        for item in os.listdir(folder_path):
            item_path = os.path.join(folder_path, item)
            if os.path.isdir(item_path):
                print_folder_structure(item_path, indent)
            elif os.path.isfile(item_path):
                print(f"{indent}|-{item}")


# Replace 'folder_path' with the path of the root folder you want to print
folder_path = "/home/user/Documents/goVkBot"
print_folder_structure(folder_path)
