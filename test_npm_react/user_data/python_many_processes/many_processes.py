import subprocess
import queue

def main():
    names = ["process_1", "process_2"]
    process_queue = queue.Queue()
    for name in names:
        print(f"Starting: {name}")
        p = subprocess.Popen(["python3", "one_process.py", name])
        print(f"After start: {name}")
        print(f"Type of p: {type(p)}")
        process_queue.put_nowait(p)
    print("Processing process queue")
    while not process_queue.empty():
        p = process_queue.get_nowait()
        print(f"Starting querying process {p}")
        print(p.communicate())
        print(f"Finishing querying process {p}")
    print("No more processes remain")

if __name__ == '__main__':
    main()