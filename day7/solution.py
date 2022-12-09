import os
import shlex
from typing import Dict, List


class Directory:
    path: str
    files: Dict[str, int]
    subdirectories: Dict[str, "Directory"]

    def __init__(self, path: str) -> None:
        self.path = path
        self.files = {}
        self.subdirectories = {}

    @property
    def parent(self) -> str:
        return os.path.abspath(os.path.join(self.path, os.path.pardir))

    @property
    def size(self) -> int:
        return sum(self.files.values()) + sum(
            [directory.size for directory in self.subdirectories.values()]
        )

    def __repr__(self) -> str:
        return f"Directory('{self.path}', {self.size})"

    def __str__(self) -> str:
        return repr(self)


# initialize a list to hold global directories
known_directories: Dict[str, Directory] = {}

# start out with a root cwd
cwd: str = "/"


def handle_command(terminal_output: List[str], pos: int) -> int:
    # lex the incoming shell command
    command_lexed: List[str] = shlex.split(terminal_output[pos][1:])

    match command_lexed[0]:
        case "cd":
            # get the target directory from the command
            target_dir: str = os.path.join(*command_lexed[1:])

            # change directory to the provided target directory
            # pylint: disable=global-statement,invalid-name
            global cwd
            cwd = os.path.abspath(os.path.join(cwd, target_dir))

            # increment to the next command
            pos += 1

        case "ls":
            # check if there were any unexpected arguments to the ls command
            if len(command_lexed) > 1:
                raise ValueError(f"Unexpected arguments to 'ls': {command_lexed[1:]}")

            # iterate past the 'ls' command
            pos += 1

            # loop over all of the lines for this ls command
            while (
                pos < len(terminal_output)
                and shlex.split(terminal_output[pos])[0] != "$"
            ):
                # split the current line into its consituents
                ls_line_lexed: List[str] = terminal_output[pos].split(" ")

                match ls_line_lexed[0]:
                    case "dir":
                        # this is a directory. add it to the dictionary
                        known_directories[
                            os.path.join(cwd, ls_line_lexed[1])
                        ] = Directory(os.path.join(cwd, ls_line_lexed[1]))
                    case other:
                        # ensure that this directory is already in the dictionary
                        if cwd not in known_directories:
                            known_directories[cwd] = Directory(cwd)

                        # this is a file and should be added to this directory
                        known_directories[cwd].files[ls_line_lexed[1]] = int(
                            ls_line_lexed[0]
                        )

                pos += 1
        case other:
            raise ValueError(f"Unknown command '{command_lexed[0]}' encountered")

    return pos


# get all of the terminal output
terminal_output: List[str]
with open("./input.txt", "r") as input_file:
    terminal_output = [
        line.strip() for line in input_file.readlines() if len(line.strip()) > 0
    ]

# loop through all of the terminal output
n: int = 0
while n < len(terminal_output):
    output_line: str = terminal_output[n]

    # determine if this is a command which is what we'd expect
    match output_line[0]:
        case "$":
            n = handle_command(terminal_output, n)
        case other:
            raise ValueError(f"Unexpected line '{output_line}'")

    print(output_line)

# sort the directories into buckets based on their depth
directories_by_depth: Dict[int, List[Directory]] = {}
for directory in known_directories.values():
    # get the depth for this directory
    path_depth: int = len(
        [dir_name for dir_name in directory.path.split(os.sep) if dir_name != ""]
    )

    # ensure a list exists in this map for the depth mapping for this directory
    if path_depth not in directories_by_depth:
        directories_by_depth[path_depth] = []

    # map the directory by its depth
    directories_by_depth[path_depth].append(directory)

# go through each depth from 0 to the lowest level and populate subdirectories
for depth in range(max(directories_by_depth.keys()), 0, -1):
    # loop over all of the subdirectories at this depth:
    for subdirectory in directories_by_depth[depth]:
        # populate the known parent with this subdirectory as a child
        known_directories[subdirectory.parent].subdirectories[
            subdirectory.path
        ] = subdirectory

# get a list of directories with size of 100000 or less
filtered_dirs: List[str] = [
    directory for directory in known_directories.values() if directory.size <= 100_000
]
print(filtered_dirs)

# output the final sum
print(f"sum: {sum(directory.size for directory in filtered_dirs)}")
