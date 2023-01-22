#!/usr/bin/env python3

import os

test_files = {
    "create.navi",
    "insert.navi"
}

execute = "./src -s {0} >> /dev/null"

if not os.path.isfile("src"):
    print("Navi executable not found, trying to compile..")
    if os.system("go build ../") == 0:
        print("Navi compiled successfully")
    else:
        print("Couldn't build navi in ../src directory")
        exit(1)

for f in test_files:
    if os.system(execute.format(f)) == 0:
        print(f"NAVI SCRIPT EXECUTED: {f}")
    else:
        print(f"NAVI FAILED ON FILE {f}")

