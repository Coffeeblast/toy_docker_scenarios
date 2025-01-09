import sys
import time

def main():
    name = sys.argv[1]
    counter = 0
    while True:
        print(f"{name}: {counter}")
        time.sleep(1)
        counter += 1
        if counter > 5: break

if __name__ == '__main__':
    main()