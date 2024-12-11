import pathlib
from queue import LifoQueue
import shutil
   
def depth_first_recursive_folder_iterator(root_dir_path, max_depth=None):
    """
        Parent folder is returned before child items. After encountering a folder, iteration over it is started immediately
    """
    iterator_stack = LifoQueue()
    current_iterator = root_dir_path.iterdir()
    current_depth = 1
    while True:
        try:
            item = next(current_iterator)
            yield item, current_depth
        except StopIteration:
            if iterator_stack.empty():
                return
            else:
                current_iterator, current_depth = iterator_stack.get_nowait()
        except PermissionError as e:
            print(e)
        else:
            try: 
                if item.is_dir() and ((max_depth is None) or (current_depth < max_depth)):
                    iterator_stack.put_nowait((current_iterator, current_depth))
                    current_iterator = item.iterdir()
                    current_depth += 1
            except OSError as e:
                print(e)

def print_row(vals, widths, sep):
    print(sep.join(val.rjust(width, " ") for val, width in zip(vals, widths)))

def print_table(header, data, sep):
    row_size = len(data[0])
    widths = [max(len(row[idx]) for row in data) for idx in range(row_size)]
    for idx in range(row_size):
        widths[idx] = max(widths[idx], 1, len(header[idx]))
    print_row(header, widths, sep)
    print_row([w * "-" for w in widths], widths, sep)
    for row in data:
        print_row(row, widths, sep)

def print_stats(root_dir : str, max_depth : int):
    """Print statistics of files in the root_dir, according to their suffixes"""
    root_dir_path = pathlib.Path(root_dir).resolve()
    print(f"Getting stats from {root_dir_path} with max depth {max_depth}")
    stats = {}
    counter = 0
    count_print_step = 10000
    for item, depth in depth_first_recursive_folder_iterator(root_dir_path, max_depth):
        #print(depth * "  " + f"--> [{depth}] {item.name}")
        try:
            if item.is_file():
                suffix = item.suffix
                stats_dict = stats.get(suffix)
                if stats_dict is None:
                    stats_dict = {"count" : 0, "size" : 0}
                    stats[suffix] = stats_dict
                stats_dict["count"] += 1
                stats_dict["size"] += item.stat().st_size

        except OSError as e:
            print(e)

        counter += 1
        if counter % count_print_step == 0:
            print(f"Processed {counter} items.")

    keys_sorted = list(stats.keys())
    keys_sorted.sort(key=lambda x: stats[x]["size"], reverse=True)
    
    header = ["Extension", "Count", "Size[MiB]", "Size[B]"]
    MIB_TO_BYTES = 1024 * 1024
    data = [
        [k, str(stats[k]["count"]), f"{stats[k]['size'] / MIB_TO_BYTES:.2f}", str(stats[k]['size'])]
        for k in keys_sorted
    ]
    sep = " | "
    print_table(header, data, sep)

def copy_tree_with_filter(src_dir : str, tar_dir: str, suffixes: "list[str]"):
    """Copy src_dir tree into tar_dir, but only files with certain suffixes"""
    src_dir_path = pathlib.Path(src_dir).resolve()
    tar_dir_path = pathlib.Path(tar_dir).resolve()
    if tar_dir_path.exists():
        print(f"Warning: {tar_dir} already exists. Skipping copy_tree_with_filter")
        return
    suffix_set = set(x if (x == "" or x[0] == ".") else "." + x for x in suffixes)
    tar_dir_path.mkdir(parents=True)
    counter = 0
    count_print_step = 10000
    for item, depth in depth_first_recursive_folder_iterator(src_dir_path):
        #print(depth * "  " + f"--> [{depth}] {item.name}")
        try:
            if item.is_file() and item.suffix in suffix_set:
                rel_path = item.relative_to(src_dir_path)
                tar_path = tar_dir_path / rel_path
                print(f"Copying {item} to {tar_path}")
                if not tar_path.parent.exists():
                    tar_path.parent.mkdir(parents=True)
                shutil.copyfile(item, tar_path)
                
        except OSError as e:
            print(e)

        counter += 1
        if counter % count_print_step == 0:
            print(f"Processed {counter} items.")

def main():
    #print_stats("../test_gradio", None)
    copy_tree_with_filter("../test_gradio", "./test_gradio_2", [".js", ".c"])

if __name__ == "__main__":
    main()